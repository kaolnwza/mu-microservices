package service

import (
	"context"
	"errors"

	"github.com/google/uuid"
	entity "github.com/kaolnwza/muniverse/feed/internal/application/core/entities"
	port "github.com/kaolnwza/muniverse/feed/internal/ports"
)

type postSvc struct {
	tx          port.Transactor
	postRepo    port.PostRepository
	storeRpcCli port.ImageStorer
}

func NewPostService(tx port.Transactor, postRepo port.PostRepository, storeRpcCli port.ImageStorer) port.PostService {
	return &postSvc{tx: tx, postRepo: postRepo, storeRpcCli: storeRpcCli}
}

func (s *postSvc) GetAllPosts(ctx context.Context, bottomTime string, userUUID uuid.UUID) (*[]*entity.PostResponse, error) {
	posts := make([]*entity.PostWithImage, 0)

	if err := s.postRepo.GetByTime(ctx, &posts, bottomTime, userUUID); err != nil {
		return nil, err
	}

	resp := make([]*entity.PostResponse, 0)
	postsImg := make([]*entity.PostImageResponse, 0)
	imgLength := 0
	for idx, item := range posts {
		postImg := &entity.PostImageResponse{
			UUID:       item.PostImage.UUID,
			Order:      int(item.PostImage.Order.Int32),
			UploadUUID: item.PostImage.UploadUUID.UUID,
		}

		post := &entity.PostResponse{
			UUID:          item.Post.UUID,
			UserUUID:      item.Post.UserUUID,
			Title:         item.Post.Title,
			Text:          item.Post.Text,
			CreatedAt:     item.Post.CreatedAt,
			LikeStatus:    item.Post.LikeStatus,
			LikeCount:     item.Post.LikeCount,
			CommentAmount: item.Post.CommentAmount,
			PostImage:     postsImg,
		}

		if idx != 0 {
			post = &entity.PostResponse{
				UUID:          posts[idx-1].Post.UUID,
				UserUUID:      posts[idx-1].Post.UserUUID,
				Title:         posts[idx-1].Post.Title,
				Text:          posts[idx-1].Post.Text,
				CreatedAt:     posts[idx-1].Post.CreatedAt,
				LikeStatus:    posts[idx-1].Post.LikeStatus,
				LikeCount:     posts[idx-1].Post.LikeCount,
				CommentAmount: posts[idx-1].Post.CommentAmount,
				PostImage:     postsImg,
			}

			if posts[idx].Post.UUID != posts[idx-1].Post.UUID {
				resp = append(resp, post)
				postsImg = nil
			}
		}

		if item.PostImage.UUID != uuid.Nil {
			postsImg = append(postsImg, postImg)
			imgLength++
		}

		if idx == len(posts)-1 {
			post.PostImage = postsImg
			resp = append(resp, post)
		}
	}

	imgCh := make(chan error)
	if imgLength != 0 {
		for _, post := range resp {
			for _, item := range post.PostImage {
				go func(item *entity.PostImageResponse) {
					url, err := s.storeRpcCli.GetURLByUploadUUID(ctx, item.UploadUUID)
					imgCh <- err
					if url != nil {
						item.Url = url
					}

				}(item)
			}
		}

		for i := 0; i < imgLength; i++ {
			if err := <-imgCh; err != nil {
				return nil, err
			}
		}
	}

	return &resp, nil
}

func (s *postSvc) GetPostByPostUUID(ctx context.Context, postUUID uuid.UUID, userUUID uuid.UUID) (*entity.PostResponse, error) {
	posts := make([]*entity.PostWithImage, 0)

	if err := s.postRepo.GetByPostUUID(ctx, &posts, postUUID, userUUID); err != nil {
		return nil, err
	}

	resp := entity.PostResponse{}
	postsImg := make([]*entity.PostImageResponse, 0)

	for idx, item := range posts {
		if item.PostImage.UploadUUID.Valid {
			postImg := &entity.PostImageResponse{
				Order:      int(item.PostImage.Order.Int32),
				UUID:       item.PostImage.UUID,
				UploadUUID: item.PostImage.UploadUUID.UUID,
			}

			postsImg = append(postsImg, postImg)
		}

		if ok := idx == len(posts)-1; ok {
			resp = entity.PostResponse{
				UUID:          item.Post.UUID,
				UserUUID:      item.Post.UserUUID,
				Title:         item.Post.Title,
				Text:          item.Post.Text,
				CreatedAt:     item.Post.CreatedAt,
				LikeStatus:    item.Post.LikeStatus,
				LikeCount:     item.Post.LikeCount,
				CommentAmount: item.Post.CommentAmount,
				PostImage:     postsImg,
			}
		}

	}

	imgCh := make(chan error)
	if len(resp.PostImage) > 0 {
		for _, item := range resp.PostImage {
			go func(item *entity.PostImageResponse) {
				url, err := s.storeRpcCli.GetURLByUploadUUID(ctx, item.UploadUUID)
				imgCh <- err
				if url != nil {
					item.Url = url
				}
			}(item)
		}

		for i := 0; i < len(resp.PostImage); i++ {
			if err := <-imgCh; err != nil {
				return nil, err
			}
		}
	}

	return &resp, nil
}

func (s *postSvc) CreatePost(ctx context.Context, title string, text string, userUUID uuid.UUID, images []*entity.PostImageRequest) error {
	var postUUID uuid.UUID

	if err := s.tx.WithinTransaction(ctx, func(tx context.Context) error {
		if err := s.postRepo.CreatePost(tx, &postUUID, title, text, userUUID); err != nil {
			return err
		}

		if err := s.postRepo.CreatePostImages(tx, &postUUID, images); err != nil {
			return err
		}

		return nil
	}); err != nil {
		return err
	}

	return nil
}

func (s *postSvc) DeletePost(ctx context.Context, postUUID uuid.UUID, userUUID uuid.UUID) error {
	post := make([]*entity.PostWithImage, 0)

	if err := s.postRepo.GetByPostUUID(ctx, &post, postUUID, userUUID); err != nil {
		return err
	}

	if post[0].Post.UserUUID != userUUID {
		return errors.New("not owner kub")
	}

	if post[0].Post.DeletedAt.Valid {
		return errors.New("post is already deleted")
	}

	if err := s.postRepo.DeletePost(ctx, postUUID); err != nil {
		return err
	}

	return nil
}

func (s *postSvc) UpdatePost(ctx context.Context, userUUID uuid.UUID, postUUID uuid.UUID, title string, text string) error {
	post := make([]*entity.PostWithImage, 0)

	if err := s.postRepo.GetByPostUUID(ctx, &post, postUUID, userUUID); err != nil {
		return err
	}

	if post[0].Post.UserUUID != userUUID {
		return errors.New("not owner kub")
	}

	if err := s.postRepo.UpdatePost(ctx, postUUID, title, text); err != nil {
		return err
	}

	return nil
}

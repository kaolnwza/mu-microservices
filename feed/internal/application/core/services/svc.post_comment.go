package service

import (
	"context"
	"errors"

	"github.com/google/uuid"
	entity "github.com/kaolnwza/muniverse/feed/internal/application/core/entities"
	port "github.com/kaolnwza/muniverse/feed/internal/ports"
	"github.com/kaolnwza/muniverse/feed/lib/helper"
)

type comntSvc struct {
	tx        port.Transactor
	comntRepo port.CommentRepository
}

func NewCommentService(tx port.Transactor, comntRepo port.CommentRepository) port.CommentService {
	return &comntSvc{tx: tx, comntRepo: comntRepo}
}

func (s *comntSvc) CreateComment(ctx context.Context, userUUID uuid.UUID, postUUID uuid.UUID, text string) (*entity.CommentResponse, error) {
	comment := entity.Comment{}
	if err := s.comntRepo.CreateComment(ctx, &comment, userUUID, postUUID, text); err != nil {
		return nil, err
	}

	resp := &entity.CommentResponse{
		UUID:      comment.UUID,
		UserUUID:  comment.UserUUID,
		Comment:   comment.Comment,
		CreatedAt: comment.CreatedAt,
	}

	return resp, nil
}

func (s *comntSvc) GetCommentByPostUUID(ctx context.Context, postUUID uuid.UUID) (*[]*entity.CommentResponse, error) {
	comment := make([]*entity.Comment, 0)
	if err := s.comntRepo.GetCommentByPostUUID(ctx, &comment, postUUID); err != nil {
		return nil, err
	}

	resp := make([]*entity.CommentResponse, 0)

	if err := helper.StructCopy(comment, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

func (s *comntSvc) DeleteCommentByUUID(ctx context.Context, userUUID uuid.UUID, commentUUID uuid.UUID) error {
	comment := entity.Comment{}
	if err := s.comntRepo.GetCommentByUUID(ctx, &comment, commentUUID); err != nil {
		return err
	}

	if comment.UserUUID != userUUID {
		return errors.New("not owner kub")
	}

	if err := s.comntRepo.DeleteComment(ctx, commentUUID); err != nil {
		return err
	}

	return nil
}

package service

import (
	"context"

	"github.com/google/uuid"
	port "github.com/kaolnwza/muniverse/feed/internal/ports"
)

type likeSvc struct {
	tx       port.Transactor
	likeRepo port.LikeRepository
}

func NewLikeService(tx port.Transactor, likeRepo port.LikeRepository) port.LikeService {
	return &likeSvc{tx: tx, likeRepo: likeRepo}
}

func (s *likeSvc) PostLike(ctx context.Context, userUUID uuid.UUID, postUUID uuid.UUID) (*bool, error) {
	var status bool = true
	if err := s.likeRepo.PostLike(ctx, &status, userUUID, postUUID); err != nil {
		return nil, err
	}

	return &status, nil

}

func (s *likeSvc) PostUnlike(ctx context.Context, userUUID uuid.UUID, postUUID uuid.UUID) (*bool, error) {
	var status bool = false
	if err := s.likeRepo.PostUnlike(ctx, &status, userUUID, postUUID); err != nil {
		return nil, err
	}

	return &status, nil
}

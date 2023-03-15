package service

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
	entity "github.com/kaolnwza/muniverse/seer/internal/application/core/entities"
	port "github.com/kaolnwza/muniverse/seer/internal/ports"
)

type seerSvc struct {
	repo      port.SeerRepository
	tx        port.Transactor
	userSvc   port.UserService
	storerSvc port.StorageService
}

func NewSeerService(r port.SeerRepository, tx port.Transactor, userSvc port.UserService, storer port.StorageService) port.SeerService {
	return &seerSvc{repo: r, tx: tx, userSvc: userSvc, storerSvc: storer}
}

func (s *seerSvc) GetSeerByUUID(ctx context.Context, seerUUID uuid.UUID) (*entity.SeerResponse, error) {
	seer := entity.Seer{}
	resp := entity.SeerResponse{}

	if err := s.repo.GetByUUID(ctx, &seer, seerUUID); err != nil {
		return nil, err
	}

	userCh := make(chan error)
	imgCh := make(chan error)

	go func() {
		user, err := s.userSvc.GetUserInfo(ctx, seer.UserUUID)
		resp.UUID = seerUUID
		resp.UserUUID = user.UUID
		resp.OnsiteAvailable = seer.OnsiteAvailable
		resp.ChatAvailable = seer.ChatAvailable
		resp.CallAvailable = seer.CallAvailable
		resp.VideoCallAvailable = seer.VideoCallAvailable
		resp.Major = seer.Major
		resp.MajorDescription = seer.MajorDescription.String
		resp.DescriptionProfile = seer.DescriptionProfile.String
		resp.MapCoordinate = seer.MapCoordinate.String
		resp.DisplayName = user.DisplayName

		userCh <- err
	}()

	go func() {
		img, err := s.storerSvc.GetUserProfileImage(seer.UserUUID)
		resp.ImageURL = *img

		imgCh <- err
	}()

	if err := <-userCh; err != nil {
		return nil, err
	}

	if err := <-imgCh; err != nil {
		return nil, err
	}

	return &resp, nil
}

func (s *seerSvc) GetSeerByUserUUID(ctx context.Context, userUUID uuid.UUID) (*entity.SeerResponse, error) {
	seer := entity.Seer{}
	resp := entity.SeerResponse{}

	if err := s.repo.GetByUserUUID(ctx, &seer, userUUID); err != nil {
		if err == sql.ErrNoRows {
			return &resp, nil
		}
		return nil, err
	}

	userCh := make(chan error)
	imgCh := make(chan error)

	go func() {
		user, err := s.userSvc.GetUserInfo(ctx, seer.UserUUID)
		userCh <- err

		resp.UUID = seer.UUID
		resp.UserUUID = user.UUID
		resp.OnsiteAvailable = seer.OnsiteAvailable
		resp.ChatAvailable = seer.ChatAvailable
		resp.CallAvailable = seer.CallAvailable
		resp.VideoCallAvailable = seer.VideoCallAvailable
		resp.Major = seer.Major
		resp.MajorDescription = seer.MajorDescription.String
		resp.DescriptionProfile = seer.DescriptionProfile.String
		resp.MapCoordinate = seer.MapCoordinate.String
		resp.DisplayName = user.DisplayName

	}()

	go func() {
		img, err := s.storerSvc.GetUserProfileImage(userUUID)
		imgCh <- err

		resp.ImageURL = *img

	}()

	if err := <-userCh; err != nil {
		return nil, err
	}

	if err := <-imgCh; err != nil {
		return nil, err
	}

	// resp := entity.SeerResponse{
	// 	UUID:               seerUUID,
	// 	UserUUID:           user.UUID,
	// 	OnsiteAvailable:    seer.OnsiteAvailable,
	// 	ChatAvailable:      seer.ChatAvailable,
	// 	CallAvailable:      seer.CallAvailable,
	// 	VideoCallAvailable: seer.VideoCallAvailable,
	// 	Major:              seer.Major,
	// 	MajorDescription:   seer.MajorDescription.String,
	// 	DescriptionProfile: seer.DescriptionProfile.String,
	// 	MapCoordinate:      seer.MapCoordinate.String,
	// 	ImageURL:           user.ProfilePictureURL,
	// 	DisplayName:        user.DisplayName,
	// }

	return &resp, nil
}

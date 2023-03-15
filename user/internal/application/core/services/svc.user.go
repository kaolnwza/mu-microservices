package service

import (
	"context"

	"github.com/google/uuid"
	entity "github.com/kaolnwza/muniverse/user/internal/application/core/entities"
	port "github.com/kaolnwza/muniverse/user/internal/ports"
	log "github.com/kaolnwza/muniverse/user/lib/logs"
)

type userSvc struct {
	userRepo port.UserRepository
	tx       port.Transactor
	imgStore port.ImageStorer
	walSvc   port.WalletService
}

func NewUserService(r port.UserRepository, tx port.Transactor, storer port.ImageStorer, walSvc port.WalletService) port.UserService {
	return &userSvc{
		userRepo: r,
		tx:       tx,
		imgStore: storer,
		walSvc:   walSvc,
	}
}

func (s *userSvc) GetUserByUUID(ctx context.Context, userUUID uuid.UUID) (*entity.UserResponse, error) {
	user := entity.User{}
	resp := entity.UserResponse{}

	dbErr := make(chan error)
	imgErr := make(chan error)
	go func() {
		err := s.userRepo.GetByUUID(ctx, &user, userUUID)
		dbErr <- err

		resp.UUID = user.UUID
		resp.DisplayName = user.DisplayName
		resp.Birthday = user.Birthday.Time.String()
		resp.Description = user.Description.String
		resp.TelNumber = user.TelNumber.String
		resp.Role = user.Role
	}()

	go func() {
		imgUrl, err := s.imgStore.GetUserProfileImage(userUUID)
		if err != nil {
			log.Error(err)
		}
		imgErr <- err
		resp.ProfilePictureURL = *imgUrl
	}()

	if err := <-dbErr; err != nil {
		return nil, err
	}

	if err := <-imgErr; err != nil {
		return nil, err
	}

	return &resp, nil
}

func (s *userSvc) GetUserWithoutImgByUUID(ctx context.Context, userUUID uuid.UUID) (*entity.UserResponse, error) {
	user := entity.User{}
	if err := s.userRepo.GetByUUID(ctx, &user, userUUID); err != nil {
		return nil, err
	}

	resp := entity.UserResponse{
		UUID:              user.UUID,
		DisplayName:       user.DisplayName,
		Birthday:          user.Birthday.Time.String(),
		Description:       user.Description.String,
		TelNumber:         user.TelNumber.String,
		Role:              user.Role,
		ProfilePictureURL: "",
	}

	return &resp, nil
}

func (s *userSvc) CreateUser(ctx context.Context, display_name string, dob string, desc string) (*entity.User, error) {
	user := entity.User{}

	if err := s.tx.WithinTransaction(ctx, func(tx context.Context) error {
		if err := s.userRepo.CreateUser(tx, &user, display_name, dob, desc); err != nil {
			return err
		}

		if err := s.walSvc.CreateNewUserWallet(tx, user.UUID); err != nil {
			log.Error(err)
			return err
		}

		return nil
	}); err != nil {
		return nil, err
	}

	return &user, nil
}

func (s *userSvc) UpdateUser(ctx context.Context, userUUID uuid.UUID, display_name string, dob string, desc string, uploadUUID uuid.UUID) error {
	if uploadUUID != uuid.Nil {
		if err := s.imgStore.UpdateUserProfileImage(userUUID, uploadUUID); err != nil {
			return err
		}
	}

	if display_name != "" && dob != "" && desc != "" {
		if err := s.userRepo.Update(ctx, userUUID, display_name, dob, desc); err != nil {
			return err
		}
	}

	return nil
}

package port

import "github.com/google/uuid"

type ImageStorer interface {
	GetUserProfileImage(userUUID uuid.UUID) (*string, error)
	UpdateUserProfileImage(userUUID uuid.UUID, uploadUUID uuid.UUID) error
}

package port

import "github.com/google/uuid"

type StorageService interface {
	GetUserProfileImage(userUUID uuid.UUID) (*string, error)
}

package port

import (
	"context"

	"github.com/google/uuid"
)

type ImageStorer interface {
	GetURLByUploadUUID(ctx context.Context, uploadUUID uuid.UUID) (*string, error)
}

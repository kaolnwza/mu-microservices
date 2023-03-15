package port

import (
	"context"

	"github.com/google/uuid"
	entity "github.com/kaolnwza/muniverse/order/internal/application/core/entities"
)

type ImageStorageService interface {
	GetHoroServiceImages(ctx context.Context, horoUUID uuid.UUID) (*[]*entity.HoroImages, error)
	NewHoroServiceImages(ctx context.Context, horoUUID uuid.UUID, images []*entity.HoroImagesRequest) error
}

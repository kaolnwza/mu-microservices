package entity

import "github.com/google/uuid"

type HoroServiceImage struct {
	UUID            uuid.UUID `db:"-"`
	HoroServiceUUID uuid.UUID `db:"-"`
	UploadUUID      uuid.UUID `db:"-"`
	ImageOrder      int32     `db:"image_order"`
	Bucket          string    `db:"bucket" json:"bucket"`
	Path            string    `db:"path" json:"path"`
}

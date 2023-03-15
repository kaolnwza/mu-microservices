package postgres

import (
	"context"

	"github.com/google/uuid"
	entity "github.com/kaolnwza/muniverse/user/internal/application/core/entities"
	port "github.com/kaolnwza/muniverse/user/internal/ports"
)

type userRepo struct {
	tx port.Transactor
}

func NewUserRepository(tx port.Transactor) port.UserRepository {
	return &userRepo{tx: tx}
}

func (r *userRepo) GetByUUID(ctx context.Context, dest *entity.User, userUUID uuid.UUID) error {
	query := `
		SELECT
			uuid,
			display_name,
			birthday,
			description,
			tel_number,
			role
			-- profile_picture,
			-- coin_amount
		FROM "user"
		WHERE uuid = $1
	`

	return r.tx.GetOne(ctx, dest, query, userUUID)
}

func (r *userRepo) CreateUser(ctx context.Context, dest *entity.User, display_name string, dob string, desc string) error {
	query := `
		INSERT INTO "user" (display_name, birthday, description)
		VALUES ($1, $2, $3)
		RETURNING uuid, display_name, birthday, description
`

	return r.tx.InsertWithReturningOne(ctx, dest, query, display_name, dob, desc)
}

func (r *userRepo) Update(ctx context.Context, userUUID uuid.UUID, display_name string, dob string, desc string) error {
	query := `
		UPDATE "user"
		SET 
			display_name = $2,
			birthday = $3,
			description = $4
		WHERE uuid = $1
	`

	return r.tx.Insert(ctx, query, userUUID, display_name, dob, desc)
}

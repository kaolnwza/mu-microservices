package repository

import (
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	entity "github.com/kaolnwza/muniverse/auth/entities"
)

// type pgConn struct {
// 	conn *sqlx.DB
// }

// type OAuthRepository interface {
// 	FetchByEmail(sqlx.Queryer, *entity.AuthGoogle, string) error
// 	CreateAuthGoogle(sqlx.Execer, uuid.UUID) error
// }

// func NewOAuthRepo(conn *sqlx.DB) OAuthRepository {
// 	return &pgConn{conn: conn}
// }

func FetchByEmail(tx sqlx.Queryer, dest *entity.AuthGoogle, email string) error {
	query := `
		SELECT
			uuid,
			user_uuid,
			email
		FROM auth_google
		WHERE email = $1
	`

	return sqlx.Get(tx, dest, query, email)

}

func CreateAuthGoogle(tx sqlx.Execer, userUUID uuid.UUID, email string) error {
	query := `
		INSERT INTO auth_google (user_uuid, email)
		VALUES ($1, $2)
	`

	_, err := tx.Exec(query, userUUID, email)
	return err
}

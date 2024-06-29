package postgres

import (
	"database/sql"
	"fmt"
	"time"
	"user_service/genproto/user_management"

	"github.com/pkg/errors"
)

type UserRepo struct {
	Db *sql.DB
}

func (u *UserRepo) Put(user *user_management.User) error {
	query := "update users set "
	params := []interface{}{}

	if user.Username != "" {
		query += fmt.Sprintf("username = $%d, ", len(params)+1)
		params = append(params, user.Username)
	}
	if user.Email != "" {
		query += fmt.Sprintf("email = $%d, ", len(params)+1)
		params = append(params, user.Email)
	}
	if user.Password != "" {
		query += fmt.Sprintf("password_hash = $%d, ", len(params)+1)
	}
	query += fmt.Sprintf("updated_at = $%d ", len(params)+1)
	params = append(params, time.Now())
	query += fmt.Sprintf("where id = $%d", len(params)+1)
	params = append(params, user.Id)

	_, err := u.Db.Exec(query, params...)
	if err != nil {
		return errors.Wrap(err, "failed to update user")
	}

	return nil
}
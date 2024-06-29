package postgres

import (
	"database/sql"
	pb "user_service/genproto/user"
)

type UserRepo struct {
	Db *sql.DB
}

func (u *UserRepo) GetUserById(userId int32) (*pb.User, error) {
	query := `select id, username, email, password_hash, 
	created_at, updated_at from users where id = $1`

	user := pb.User{Id: userId}
	err := u.Db.QueryRow(query, userId).Scan(&user.Username,
		&user.Email, &user.Password, &user.CreatedAt,
		&user.UpdatedAt)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

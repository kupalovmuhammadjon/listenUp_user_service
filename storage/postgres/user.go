package postgres

import (
	"database/sql"
	"fmt"
	pbu "user_service/genproto/user"
)

type UserRepo struct {
	Db *sql.DB
}

func NewUserRepo(db *sql.DB) *UserRepo {
	return &UserRepo{Db: db}
}

func (u *UserRepo) DeleteUser(id string) error {
	query := `
	update
	    users 
	set
	    deleted_at = now(),
	WHERE 
	    id = $1;
`
	res, err := u.Db.Exec(query, id)
	affectedRows, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if affectedRows == 0 {
		return fmt.Errorf("user %s not found", id)
	}

	return nil
}

func (u *UserRepo) UpdateUserProfile(profile pbu.Profile) error {
	query := `
	update
	    user_profiles 
	set
    	full_name = $1,
    	bio = $2,
   	 	role = $3,
    	location = $4,
    	avatar_image = $5,
    	website = $6,
	    updated_at = now(),
	WHERE 
	    id = $7;
`
	res, err := u.Db.Exec(query, profile.FullName, profile.Bio, profile.Location, profile.AvatarImage, profile.Website)
	affectedRows, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if affectedRows == 0 {
		return fmt.Errorf("user %s not found", profile.UserId)
	}

	return nil
}

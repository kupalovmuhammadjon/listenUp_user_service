package service

import (
	"user_service/storage/postgres"
)

type Users struct {
	Users *postgres.UserRepo
}

func NewUsers(users *postgres.UserRepo) *Users {
	return &Users{Users: users}
}

// func (u *Users) GetUserById(ctx *context.Context, req.)

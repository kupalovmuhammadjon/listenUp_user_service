package service

import (
	"context"
	"database/sql"
	"fmt"
	pb "user_service/genproto/user"
	"user_service/storage/postgres"
)

type UserManagement struct {
	pb.UnimplementedUserManagementServer
	Users *postgres.UserRepo
}

func NewUserManagement(db *sql.DB) *UserManagement {
	users := postgres.NewUserRepo(db)
	return &UserManagement{Users: users}
}

func (u *UserManagement) GetUserById(ctx *context.Context, req *pb.ID) (*pb.User, error) {
	if len(req.Id) != 32 {
		return nil, fmt.Errorf("id is not valid")
	}
	user, err := u.Users.GetUserById(req.Id)
	if err != nil {
		return nil, err
	}
	return &pb.User{
		Id:        user.Id,
		Username:  user.Username,
		Email:     user.Email,
		Password:  user.Password,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}, nil
}

func (u *UserManagement) UpdateProfile(ctx context.Context, profile *pb.Profile) (*pb.Status, error) {
	err := u.Users.UpdateUserProfile(profile)
	if err != nil {
		return &pb.Status{Status: "Failed"}, err
	}
	return &pb.Status{Status: "Success"}, nil
}

// func (u *Users) GetUserById(ctx *context.Context, req.)

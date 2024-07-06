package service

import (
	"context"
	"database/sql"
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

func (u *UserManagement) GetUserByID(ctx context.Context, req *pb.ID) (*pb.User, error) {
	user, err := u.Users.GetUserById(req.Id)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (u *UserManagement) GetUserProfile(ctx context.Context, req *pb.ID) (*pb.Profile, error) {

	userProfile, err := u.Users.GetUserProfile(req.Id)
	if err != nil {
		return nil, err
	}
	return userProfile, nil
}

func (u *UserManagement) UpdateUserProfile(ctx context.Context, profile *pb.Profile) (*pb.Void, error) {
	err := u.Users.UpdateUserProfile(profile)
	if err != nil {
		return &pb.Void{}, err
	}

	return &pb.Void{}, nil
}

func (u *UserManagement) UpdateUser(ctx context.Context, req *pb.User) (*pb.Void, error) {
	err := u.Users.UpdateUser(req)
	if err != nil {
		return &pb.Void{}, err
	}

	return &pb.Void{}, nil
}

func (u *UserManagement) DeleteUser(ctx context.Context, req *pb.ID) (*pb.Void, error) {
	err := u.Users.DeleteUser(req.Id)
	if err != nil {
		return &pb.Void{}, err
	}

	return &pb.Void{}, nil
}

func (u *UserManagement) ValidateUserId(ctx context.Context, req *pb.ID) (*pb.Success, error) {
	resp, err := u.Users.ValidateUserId(req.Id)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

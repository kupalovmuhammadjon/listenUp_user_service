package postgres

import (
	"log"
	"reflect"
	"testing"
	pb "user_service/genproto/user"
)

func DB() *UserRepo {
	db, err := ConnectDB()
	if err != nil {
		log.Fatal(err)
	}
	return NewUserRepo(db)
}

func TestGetUserById(t *testing.T) {
	u := DB()
	res, err := u.GetUserById("89d2d54e-005a-4ed9-b9fe-72633a90450b")
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	exp := pb.User{
		Id:        "89d2d54e-005a-4ed9-b9fe-72633a90450b",
		Username:  "johndoe",
		Email:     "john.doe@email.com",
		Password:  "hashed_password_1",
		CreatedAt: "2024-07-03 11:48:24.956541",
		UpdatedAt: "",
	}

	if !reflect.DeepEqual(res, &exp) {
		t.Errorf("error while getting user by id")
	}

	log.Println(res, &exp)
}

package main

import (
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
	"user_service/config"
	pb "user_service/genproto/user"
	"user_service/storage/postgres"
)

func main() {
	cfg := config.Load()
	listener, err := net.Listen("tcp", fmt.Sprintf(":%s", cfg.USER_SERVICE_PORT))
	if err != nil {
		log.Fatal(err)
	}
	defer listener.Close()
	db, err := postgres.ConnectDB()
	if err != nil {
		log.Fatal(err)
	}

	userService := service.NewUserService(db)
	s := grpc.NewServer()
	pb.RegisterUserManagementServer(s, userService)
	log.Printf("server listening at %v", listener.Addr())
	err = s.Serve(listener)
	if err != nil {
		log.Fatal(err)
	}
}

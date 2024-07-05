package handler

import (
	"database/sql"
	"user_service/storage/postgres"
)

type Handler struct {
	UserRepo *postgres.UserRepo
}

func NewHandler(db *sql.DB) *Handler {
	return &Handler{postgres.NewUserRepo(db)}
}

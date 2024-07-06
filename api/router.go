package api

import (
	"database/sql"
	"user_service/api/handler"

	"github.com/gin-gonic/gin"
)

func NewRouter(db *sql.DB) *gin.Engine {

	h := handler.NewHandler(db)

	router := gin.Default()

	// router.Use(middleware.JWTMiddleware())
	users := router.Group("/users")
	users.POST("/register", h.Register)
	users.POST("/login", h.Login)

	return router
}

package handler

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	// pbU "user_service/genproto/user"
	"encoding/json"
	pbAu "user_service/genproto/authentication"

	"golang.org/x/crypto/bcrypt"
)

func (h *Handler) Register(ctx *gin.Context) {
	req := pbAu.RegisterRequest{}

	err := json.NewDecoder(ctx.Request.Body).Decode(&req)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, err)
		log.Println(err)
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, err)
		log.Println(err)
		return
	}
	req.Password = string(hashedPassword)

	err = h.UserRepo.Register(&req)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, err)
		log.Println(err)
		return
	}
	
	ctx.JSON(http.StatusCreated, nil)
}

func (h *Handler) Login(ctx *gin.Context) {

}

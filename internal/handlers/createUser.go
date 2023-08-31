package handlers

import (
	"avitoTech/internal/model"
	"avitoTech/internal/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CreateUserRequest struct {
	ID string `json:"id" validate:"required"`
}
type CreateUserResponse struct {
	ID string `json:"id"`
}

// CreateUser — HTTP handler for creating user.
func CreateUser(ctx *gin.Context) {
	reqCtx := ctx.Request.Context()

	request := &CreateUserRequest{}

	if errBind := ctx.BindJSON(request); errBind != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"msg": errBind.Error()})
		return
	}

	service := services.Must(reqCtx).UserSegmentService
	user := &model.User{
		ID: request.ID,
	}

	newUser, createErr := service.CreateUser(reqCtx, user)
	if createErr != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"msg": createErr.Error()})
		return
	}

	ctx.JSON(http.StatusOK, &CreateUserResponse{
		ID: newUser.ID,
	})
}

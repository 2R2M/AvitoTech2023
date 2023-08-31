package handlers

import (
	"avitoTech/internal/model"
	"avitoTech/internal/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CreateSegmentRequest struct {
	Slug string `json:"slug" validate:"required,uppercase"`
}
type CreateSegmentResponse struct {
	Slug string `json:"slug" validate:"required,uppercase"`
}

// CreateSegment — HTTP handler for creating user.
func CreateSegment(ctx *gin.Context) {
	reqCtx := ctx.Request.Context()

	request := &CreateSegmentRequest{}

	if errBind := ctx.BindJSON(request); errBind != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"msg": errBind.Error()})
		return
	}

	service := services.Must(reqCtx).UserSegmentService
	segment := &model.Segment{
		Slug: request.Slug,
	}
	newSegment, createErr := service.CreateSegment(reqCtx, segment)
	if createErr != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"msg": createErr.Error()})
		return
	}

	ctx.JSON(http.StatusOK, &CreateSegmentResponse{
		Slug: newSegment.Slug,
	})
}

package handlers

import (
	"avitoTech/internal/model"
	"avitoTech/internal/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CreateDeleteOperationRequest struct {
	Slugs []string `json:"slugs" validate:"required"`
}

type CreateDeleteOperationResponse struct {
	ID     string   `json:"id"`
	Slugs  []string `json:"slugs"`
	UserId string   `json:"user_id"`
}

func DeleteUsersSegment(ctx *gin.Context) {
	id := ctx.Param("id")
	reqCtx := ctx.Request.Context()

	request := &CreateDeleteOperationResponse{}

	if errBind := ctx.BindJSON(request); errBind != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"msg": errBind.Error()})
		return
	}

	service := services.Must(reqCtx).UserSegmentService
	operation := &model.Operation{UserId: id, Segment: request.Slugs}
	operationDelete, errDelete := service.DeleteUsersSegment(reqCtx, operation)
	if errDelete != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"msg": errDelete.Error()})
		return
	}

	ctx.JSON(http.StatusOK, &CreateDeleteOperationResponse{
		ID: operationDelete.ID, UserId: operationDelete.UserId, Slugs: operationDelete.Segment,
	})
}

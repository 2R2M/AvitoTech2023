package handlers

import (
	"avitoTech/internal/model"
	"avitoTech/internal/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CreateOperationRequest struct {
	Slugs     []string `json:"slugs"`
	ExpiredAt string   `json:"expired_at"`
}

type CreateOperationResponse struct {
	ID        string   `json:"id"`
	Slugs     []string `json:"slugs"`
	UserId    string   `json:"user_id"`
	ExpiredAt string   `json:"expired_at"`
}

func AddUsersSegments(ctx *gin.Context) {
	id := ctx.Param("id")
	reqCtx := ctx.Request.Context()

	request := &CreateOperationRequest{}

	if errBind := ctx.BindJSON(request); errBind != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"msg": errBind.Error()})
		return
	}

	service := services.Must(reqCtx).UserSegmentService

	operation := &model.Operation{UserId: id, Segment: request.Slugs, ExpiredAt: request.ExpiredAt}
	operationAdd, errAdd := service.AddUsersSegment(reqCtx, operation)
	if errAdd != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"msg": errAdd.Error()})
		return
	}

	ctx.JSON(http.StatusOK, &CreateOperationResponse{
		UserId: operationAdd.UserId, Slugs: operationAdd.Segment, ExpiredAt: operationAdd.ExpiredAt})
}

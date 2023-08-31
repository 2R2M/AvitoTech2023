package handlers

import (
	"avitoTech/internal/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type GetUsersSlugsResponse struct {
	Segments []string `json:"segments"`
}

func GetUsersSlugs(ctx *gin.Context) {
	id := ctx.Param("id")
	reqCtx := ctx.Request.Context()
	service := services.Must(reqCtx).UserSegmentService
	segments, err := service.GetUsersSegments(reqCtx, id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
		return
	}
	var response GetUsersSlugsResponse
	for _, segment := range segments {
		response.Segments = append(response.Segments, segment.Slug)
	}
	ctx.JSON(http.StatusOK, response)
}

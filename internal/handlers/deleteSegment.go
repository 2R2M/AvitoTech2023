package handlers

import (
	"avitoTech/internal/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func DeleteSegment(ctx *gin.Context) {
	reqCtx := ctx.Request.Context()
	slug := ctx.Param("slug")
	service := services.Must(reqCtx).UserSegmentService

	err := service.DeleteSegment(reqCtx, slug)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
		return
	}

	ctx.Status(http.StatusOK)
}

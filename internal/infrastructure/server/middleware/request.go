package middleware

import (
	"avitoTech/internal/services"
	"context"

	"github.com/gin-gonic/gin"
)

func RequestMiddleware(c context.Context, userSegmentService services.Services) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		cNew := services.Set(c, userSegmentService)
		ctx.Request = ctx.Request.WithContext(cNew)
		ctx.Next()
	}
}

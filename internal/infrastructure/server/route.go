package server

import "github.com/gin-gonic/gin"

type Route struct {
	Method      string
	Path        string
	HandleFuncs []gin.HandlerFunc
}

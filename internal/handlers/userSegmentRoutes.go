package handlers

import (
	"avitoTech/internal/infrastructure/server"
	"fmt"

	"github.com/gin-gonic/gin"
)

const (
	Prefix     = "/api"
	apiVersion = "v1"
)

func InjectUserSegmentRoutes(s *Server) {
	s.AddRoutes(&server.Route{
		Method:      "POST",
		Path:        fmt.Sprintf("%s/%s/user", Prefix, apiVersion),
		HandleFuncs: []gin.HandlerFunc{CreateUser},
	})
	s.AddRoutes(&server.Route{
		Method:      "POST",
		Path:        fmt.Sprintf("%s/%s/segment", Prefix, apiVersion),
		HandleFuncs: []gin.HandlerFunc{CreateSegment},
	})
	s.AddRoutes(&server.Route{
		Method:      "POST",
		Path:        fmt.Sprintf("%s/%s/users/:id/segments", Prefix, apiVersion),
		HandleFuncs: []gin.HandlerFunc{AddUsersSegments},
	})
	s.AddRoutes(&server.Route{
		Method:      "DELETE",
		Path:        fmt.Sprintf("%s/%s/users/:id/segments", Prefix, apiVersion),
		HandleFuncs: []gin.HandlerFunc{DeleteUsersSegment},
	})
	s.AddRoutes(&server.Route{
		Method:      "GET",
		Path:        fmt.Sprintf("%s/%s/users/:id/segments", Prefix, apiVersion),
		HandleFuncs: []gin.HandlerFunc{GetUsersSlugs},
	})
	s.AddRoutes(&server.Route{
		Method:      "GET",
		Path:        "/ping",
		HandleFuncs: []gin.HandlerFunc{s.PingPong},
	})
	s.AddRoutes(&server.Route{
		Method:      "GET",
		Path:        fmt.Sprintf("%s/%s/report/:month/:year", Prefix, apiVersion),
		HandleFuncs: []gin.HandlerFunc{GetReport},
	})
	s.AddRoutes(&server.Route{
		Method:      "GET",
		Path:        fmt.Sprintf("/download/:filename"),
		HandleFuncs: []gin.HandlerFunc{FileDownload},
	})
	s.AddRoutes(&server.Route{
		Method:      "DELETE",
		Path:        fmt.Sprintf("%s/%s/delete/:slug", Prefix, apiVersion),
		HandleFuncs: []gin.HandlerFunc{DeleteSegment},
	})

}

package router

import (
	"github.com/discordomo/api/api"
	"github.com/gin-gonic/gin"
)

// Load returns the gin engine containing all endpoints
func Load(options ...gin.HandlerFunc) *gin.Engine {
	r := gin.New()

	// Badge endpoint
	r.GET("/health", api.Health)

	return r
}

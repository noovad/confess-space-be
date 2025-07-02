package router

import (
	"go_confess_space-project/api/controller"
	"go_confess_space-project/config/websocket"
	"net/http"
	"os"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetupRouter(hub *websocket.Hub) *gin.Engine {
	r := gin.Default()
	wsController := controller.NewWebSocketController(hub)

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{os.Getenv("FRONTEND_BASE_URL")},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization", "Refresh-token", "Signed-token", "Oauth-State"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "pong"})
	})

	r.GET("/ws/connect", wsController.HandleWebSocket)

	apiV1 := r.Group("/api/v1")

	SpaceRoutes(apiV1)
	UserSpaceRoutes(apiV1)
	MessageRoutes(apiV1, hub)
	UserSpaceLastSeenRoutes(apiV1)

	return r
}

package router

import (
	"chat-rooms-server/internal/user"
	"chat-rooms-server/internal/ws"
	"fmt"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var r *gin.Engine

func InitRouter(userHandler *user.Handler, wsHandler *ws.Handler) {
	r = gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000", "http://127.0.0.1:3000"},
		AllowMethods:     []string{"GET", "POST"},
		AllowHeaders:     []string{"Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			fmt.Println(origin)
			return origin == "http://localhost:3000" || origin == "http://127.0.0.1:3000"
		},
		MaxAge: 12 * time.Hour,
	}))
	r.POST("/signup", userHandler.CreateUser)
	r.POST("/login", userHandler.Login)
	r.POST("/logout", userHandler.Logout)
	r.GET("/user", userHandler.GetUserInfo)

	r.POST("/ws/createRoom", wsHandler.CreateRoom)
	r.GET("/ws/joinRoom/:roomId", wsHandler.JoinRoom)
	r.GET("/ws/getRooms", wsHandler.GetRooms)
	r.GET("/ws/getClients/:roomId", wsHandler.GetClients)
}

func Start(addr string) error {
	return r.Run(addr)
}
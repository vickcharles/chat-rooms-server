package main

import (
	"chat-rooms-server/db"
	"chat-rooms-server/internal/user"
	"chat-rooms-server/internal/ws"
	"chat-rooms-server/router"
	"log"
	"os"
)

func main() {
    dbConn, err := db.NewDatabase()

	if err != nil {
		log.Fatal(err)
	}

	port := os.Getenv("PORT")
    if port == "" {
        port = "8080"
    }

	userRep := user.NewRepository(dbConn.GetDB())
	userSvc := user.NewService(userRep)
	userHandler := user.NewHandler(userSvc)

	hub := ws.NewHub()
	wsHandler := ws.NewHandler(hub)
	go hub.Run()

	router.InitRouter(userHandler, wsHandler)
	router.Start("0.0.0.0:" + port)
}

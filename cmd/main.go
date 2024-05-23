package main

import (
	"log"
	"net/http"
	"twitter-clone/internal/handlers"
	"twitter-clone/internal/repository"
	"twitter-clone/internal/router"
	"twitter-clone/internal/services"
	"twitter-clone/pkg/db"
)

func main() {
	db.InitDB()

	userRepo := repository.NewUserRepository(db.DB)
	userService := services.NewUserService(userRepo)
	userHandler := handlers.NewUserHandler(userService)
	r := router.InitRouter(userHandler)

	log.Println("Server started ar := 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}

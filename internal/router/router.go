package router

import (
	"twitter-clone/internal/handlers"
	// "twitter-clone/internal/middleware"

	"github.com/gorilla/mux"
)

func InitRouter(userHandler *handlers.UserHandler) *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/register", userHandler.RegisterUser).Methods("POST")
	r.HandleFunc("/login", userHandler.LoginUser).Methods("POST")

	// r.Handle("/tweets", middleware.AuthMiddleware())
	return r
}

package router

import (
	"twitter-clone/internal/handlers"

	"github.com/gorilla/mux"
)

func InitRouter(userHandler *handlers.UserHandler) *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/register", userHandler.RegisterUser).Methods("POST")
	return r
}

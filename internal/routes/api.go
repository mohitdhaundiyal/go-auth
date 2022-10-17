package routes

import (
	"github.com/gorilla/mux"
	"github.com/mohitdhaundiyal/go-auth/internal/controllers"
)

func Routes(r *mux.Router) {
	r.HandleFunc("/", controllers.Home).Methods("GET")
	r.HandleFunc("/login", controllers.Login).Methods("POST")
	r.HandleFunc("/register", controllers.Register).Methods("POST")
}

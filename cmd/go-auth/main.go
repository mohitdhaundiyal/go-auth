package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mohitdhaundiyal/go-auth/internal/config"
	"github.com/mohitdhaundiyal/go-auth/internal/routes"
)

func main() {
	fmt.Println("JWT AUTH...")

	// Databse Migration
	config.DatabaseMigration()

	// go server -
	r := mux.NewRouter()
	routes.Routes(r)

	fmt.Println("server running on port 9080")
	log.Fatal(http.ListenAndServe(":9080", r))
}

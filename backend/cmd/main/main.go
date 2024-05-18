package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/PrathameshTheurkar/go-todoapp/backend/database"
	"github.com/PrathameshTheurkar/go-todoapp/backend/routes"
	"github.com/gorilla/mux"
)

func main() {

	database.Connect()
	router := mux.NewRouter()

	routes.AllRoutes(router)

	fmt.Println("Server is running on port: 4000")
	log.Fatal(http.ListenAndServe(":4000", router))
}

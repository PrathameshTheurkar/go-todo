package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/PrathameshTheurkar/go-todoapp/backend/database"
	"github.com/PrathameshTheurkar/go-todoapp/backend/routes"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		fmt.Println("Error while loading env")
		panic(err)
	}
	database.Connect()
	router := mux.NewRouter()

	routes.AllRoutes(router)

	fmt.Println("Server is running on port: 4000")
	log.Fatal(http.ListenAndServe(os.Getenv("PORT"), router))
}

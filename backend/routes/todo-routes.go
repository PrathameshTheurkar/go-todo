package routes

import (
	"github.com/PrathameshTheurkar/go-todoapp/backend/controllers"
	"github.com/gorilla/mux"
)

func AllRoutes(router *mux.Router) {
	router.HandleFunc("/signin", controllers.Signin).Methods("POST")
	router.HandleFunc("/signup", controllers.Signup).Methods("POST")
	router.HandleFunc("/todo/{id}", controllers.GetTodo).Methods("GET")
	router.HandleFunc("/todos", controllers.GetTodos).Methods("GET")
	router.HandleFunc("/addtodo", controllers.AddTodo).Methods("POST")
	router.HandleFunc("/todo/{id}", controllers.UpdateTodo).Methods("PUT")
	router.HandleFunc("/todo/{id}", controllers.DeleteTodo).Methods("DELETE")
}

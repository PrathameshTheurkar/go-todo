package controllers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/PrathameshTheurkar/go-todoapp/backend/database"
	"github.com/PrathameshTheurkar/go-todoapp/backend/middlewares"
	"github.com/PrathameshTheurkar/go-todoapp/backend/models"
	"github.com/gorilla/mux"
)

func Signup(w http.ResponseWriter, r *http.Request) {
	var user = &models.User{}
	if body, err := io.ReadAll(r.Body); err == nil {
		if err := json.Unmarshal([]byte(body), user); err != nil {
			panic(err)
		}
	}

	count := 0

	err := database.Db.Get(&count, `Select count(*) from users where username = ?`, user.Username)
	if err != nil {
		panic(err)
	}

	if count == 0 {
		tokenString, err := middlewares.CreateToken(user.Username)
		if err != nil {
			panic(err)
		}

		_, err = database.Db.Exec("insert into users(username, password) values(?, ?)", user.Username, user.Password)
		if err != nil {
			panic(err)
		}

		res, err := json.Marshal(tokenString)
		if err != nil {
			panic(err)
		}

		w.Header().Set("Content-Type", "pkglication/json")
		w.WriteHeader(http.StatusOK)
		w.Write(res)
	} else {
		w.Write([]byte("User already exist"))
	}

}

func Signin(w http.ResponseWriter, r *http.Request) {
	var user = &models.User{}

	if body, err := io.ReadAll(r.Body); err == nil {
		if err := json.Unmarshal([]byte(body), user); err != nil {
			panic(err)
		}
	}

	count := 0

	err := database.Db.Get(&count, "select count(*) from users where username = ?", user.Username)
	if err != nil {
		panic(err)
	}

	if count != 0 {

		authHeader := r.Header.Get("Authorization")
		arrHeader := strings.Split(authHeader, " ")
		tokenString := arrHeader[1]
		// fmt.Println("token: ", tokenString)
		err := middlewares.VerifyToken(tokenString)

		if err != nil {
			// w.Write([]byte(""))
			fmt.Println("from controllers")
			panic(err)
		}

		w.Write([]byte("User successfully login"))
	} else {
		w.Write([]byte("User does not exist"))
	}

	// w.WriteHeader(http.StatusOK)
	// w.Write([]byte())
}

func GetTodo(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	var todo = models.Todo{}
	err := database.Db.Get(&todo, "select * from todo where id = ?", id)

	if err != nil {
		panic(err)
	}

	res, err := json.Marshal(todo)
	if err != nil {
		panic(err)
	}

	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetTodos(w http.ResponseWriter, r *http.Request) {
	rows, err := database.Db.Query("Select * from todo")

	if err != nil {
		panic(err)
	}

	todos := []models.Todo{}

	for rows.Next() {
		var id int64
		var title string
		var description string

		err := rows.Scan(&title, &description, &id)

		if err != nil {
			panic(err)
		}

		todos = append(todos, models.Todo{id, title, description})
	}

	res, err := json.Marshal(todos)

	if err != nil {
		panic(err)
	}

	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func AddTodo(w http.ResponseWriter, r *http.Request) {
	var x = &models.Todo{}

	if body, err := io.ReadAll(r.Body); err == nil {
		if err := json.Unmarshal([]byte(body), x); err != nil {
			panic(err)
		}
	}

	_, err := database.Db.Exec("insert into todo(id, title, description) values(?, ?, ?)", x.Id, x.Title, x.Description)
	if err != nil {
		panic(err)
	}

	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Added Todo"))
}

func UpdateTodo(w http.ResponseWriter, r *http.Request) {
	var x = &models.Todo{}

	if body, err := io.ReadAll(r.Body); err == nil {
		if err := json.Unmarshal([]byte(body), x); err != nil {
			panic(err)
		}
	}

	params := mux.Vars(r)
	id := params["id"]

	_, err := database.Db.Exec("update todo set title = ?, description = ? where id = ?", x.Title, x.Description, id)
	if err != nil {
		panic(err)
	}

	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Updated Todo"))
}

func DeleteTodo(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]
	_, err := database.Db.Exec("Delete from todo where id = ?", id)

	if err != nil {
		panic(err)
	}

	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Deleted Todo"))
}

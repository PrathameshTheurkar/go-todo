package controllers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

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

		_, err = database.Db.Exec("insert into users(username, password) values(?, ?)", user.Username, user.Password)
		if err != nil {
			panic(err)
		}

		err = database.Db.QueryRow("select personId from users where username = ?", user.Username).Scan(&user.PersonId)
		if err != nil {
			panic(err)
		}

		err := middlewares.CreateToken(w, user.PersonId)
		if err != nil {
			panic(err)
		}
		w.Header().Set("Content-Type", "pkglication/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("User signup successfully"))
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
		database.Db.QueryRow("select personId from users where username = ?", user.Username).Scan(&user.PersonId)
		err := middlewares.CreateToken(w, user.PersonId)
		if err != nil {
			// w.Write([]byte(""))
			panic(err)
			return
		}

		w.Write([]byte("User successfully login"))
	} else {
		w.Write([]byte("User does not exist"))
	}

	// w.WriteHeader(http.StatusOK)
	// w.Write([]byte())
}

func GetTodo(w http.ResponseWriter, r *http.Request) {
	personId, err := middlewares.VerifyToken(r)
	if err != nil {
		w.Write([]byte("Plz Login or Signup first"))
		return
		// panic(err)
	}

	params := mux.Vars(r)
	id := params["id"]

	var todo = models.Todo{}
	err = database.Db.Get(&todo, "select * from todo where id = ? and personId = ?", id, personId)

	if err != nil {
		panic(err)
	}

	if todo.Title == "" && todo.Description == "" {
		w.Write([]byte("Todo does not present for this user"))
		return
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
	personId, err := middlewares.VerifyToken(r)
	if err != nil {
		w.Write([]byte("Plz Login or Signup first"))
		return
		// panic(err)
	}
	rows, err := database.Db.Query("Select * from todo where personId = ?", personId)

	if err != nil {
		panic(err)
	}

	todos := []models.Todo{}

	for rows.Next() {
		var id int64
		var pid int64
		var title string
		var description string

		err := rows.Scan(&id, &pid, &title, &description)

		if err != nil {
			panic(err)
		}

		todos = append(todos, models.Todo{id, pid, title, description})
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
	personId, err := middlewares.VerifyToken(r)
	if err != nil {
		w.Write([]byte("Plz Login or Signup first"))
		return
		// panic(err)
	}
	var x = &models.Todo{}

	if body, err := io.ReadAll(r.Body); err == nil {
		if err := json.Unmarshal([]byte(body), x); err != nil {
			panic(err)
		}
	}

	_, err = database.Db.Exec("insert into todo(personId, title, description) values(?, ?, ?)", personId, x.Title, x.Description)
	if err != nil {
		fmt.Println("error from addtodo exec")
		panic(err)
	}

	// var personId int64
	// _, err = database.Db.Select(&personId, "select personId where ")

	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Added Todo"))
}

func UpdateTodo(w http.ResponseWriter, r *http.Request) {
	personId, err := middlewares.VerifyToken(r)
	fmt.Println(personId)
	if err != nil {
		w.Write([]byte("Plz Login or Signup first"))
		return
		// panic(err)
	}
	var x = &models.Todo{}

	if body, err := io.ReadAll(r.Body); err == nil {
		if err := json.Unmarshal([]byte(body), x); err != nil {
			panic(err)
		}
	}

	params := mux.Vars(r)
	id := params["id"]

	result, err := database.Db.Exec("update todo set title = ?, description = ? where id = ? and personId = ?", x.Title, x.Description, id, personId)
	if err != nil {
		panic(err)
	}

	affected, err := result.RowsAffected()

	if err != nil {
		panic(err)
	}

	if affected == 0 {
		w.Write([]byte("Todo with this id is not present in this user"))
		return
	}

	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Updated Todo"))
}

func DeleteTodo(w http.ResponseWriter, r *http.Request) {
	personId, err := middlewares.VerifyToken(r)
	if err != nil {
		w.Write([]byte("Plz Login or Signup first"))
		return
		// panic(err)
	}
	params := mux.Vars(r)
	id := params["id"]
	result, err := database.Db.Exec("Delete from todo where id = ? and personId = ?", id, personId)

	if err != nil {
		panic(err)
	}

	affected, err := result.RowsAffected()
	if err != nil {
		panic(err)
	}

	if affected == 0 {
		w.Write([]byte("Todo with this id is not present in this user"))
		return
	}

	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Deleted Todo"))
}

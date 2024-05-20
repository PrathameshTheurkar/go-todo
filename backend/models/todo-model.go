package models

type Todo struct {
	Id          int64  `json:"id" db:"id"`
	PersonId    int64  `json:"personId" db:"personId"`
	Title       string `json:"title" db:"title"`
	Description string `json:"description" db:"description"`
}

type User struct {
	PersonId int64  `json:"personId db:"personId"`
	Username string `json:"username" db:"username"`
	Password string `json:"password" db:"password"`
}

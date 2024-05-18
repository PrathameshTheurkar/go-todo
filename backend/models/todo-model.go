package models

type Todo struct {
	Id          int64  `json:"id" db:"id"`
	Title       string `json:"title" db:"title"`
	Description string `json:"description" db:"description"`
}

type User struct {
	Username string `json:"username" db:"username"`
	Password string `json:"password" db:"password"`
}

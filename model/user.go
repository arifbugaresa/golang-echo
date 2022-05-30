package model

type UserExample struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type UserBind struct {
	ID int `json:"id" form:"id"`
}

/*
Migration ke database
*/
type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

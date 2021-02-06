package entity

type User struct {
	Account  string `json:"account" binding:"required"`
	Password string `json:"password" binding:"required"`
	Company  string `json:"company" binding:"required"`
}

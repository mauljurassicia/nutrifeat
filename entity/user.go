package entity

type User struct {
	ID int `json:"id"`
	Username string `json:"username"`
	FullName string `json:"fullname"`	
	Email string `json:"email"`
	Password string `json:"password"`
}
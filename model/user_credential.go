package model

type UserCredential struct {
	Username string `json:"username"`
	Password string `json:"userpassword"`
	Email    string
}
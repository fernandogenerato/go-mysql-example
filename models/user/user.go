package user

import (
	"encoding/json"
)

type User struct {
	Id       int    `json:"id"	`
	Username string `json:"username"`
	Email    string `json:"email"`
}

func ToUser(data []byte) User {
	var user User
	err := json.Unmarshal(data, &user)
	if err != nil {
		panic(err)
	}
	return user
}

package domain

import "github.com/google/uuid"

type User struct {
	Email    string    `json:"email"`
	Id       uuid.UUID `json:"id"`
	Name     string    `json:"name"`
	Username string    `json:"username"`
}

package user

import "github.com/google/uuid"

type User struct {
	ID        uuid.UUID `db:"id"`
	Username  string    `db:"username"`
	FirstName string    `db:"first_name"`
	LastName  string    `db:"last_name"`
	Email     string    `db:"email"`
}

type UserDTO struct {
	ID        uuid.UUID `json:"id"`
	Username  string    `json:"username" binding:"required"`
	FirstName string    `json:"firstName" binding:"required"`
	LastName  string    `json:"lastName" binding:"required"`
	Email     string    `json:"email" binding:"required"`
}

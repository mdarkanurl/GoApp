package main

import (
	"time"
	"uuid"

	"github.com/mdarkanurl/GoApp/internal/database"
)

type User struct {
	ID       uuid.UUID `json:"id"`
	Name     string    `json:"name"`
	CreateAt time.Time `json:"created_at"`
	UpdateAt time.Time `json:"updated_at"`
}

func databaseUserToUser(dbUser database.User) User {
	return User{
		ID:       uuid.UUID(dbUser.ID),
		Name:     dbUser.Name,
		CreateAt: dbUser.CreateAt,
		UpdateAt: dbUser.UpdateAt,
	}
}

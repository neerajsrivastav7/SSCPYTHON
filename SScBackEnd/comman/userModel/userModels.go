package userModel

import (
	"time"
)

type UserDetails struct {
	ID           int       `json:"id"`         // Unique identifier for the user
	FirstName    string    `json:"first_name"` // User's first name
	LastName     string    `json:"last_name"`  // User's last name
	Email        string    `json:"email"`      // User's email address
	PasswordHash []byte    `json:"-"`          // User's hashed password (not exposed in JSON)
	CreatedAt    time.Time `json:"created_at"` // Timestamp when the user was created
	UpdatedAt    time.Time `json:"updated_at"` // Timestamp when the user was last updated
}

type Registration struct {
	FirstName string `json:"first_name"` // User's first name
	LastName  string `json:"last_name"`  // User's last name
	Email     string `json:"email"`      // User's email address
	Password  string `json:"password"`   // User's password (plaintext, should be hashed before storage)
}

type Login struct {
	Email    string `json:"email"`    // User's email address
	Password string `json:"password"` // User's password (plaintext, should be checked against the stored hash)
    LoginType string `json:"logintype"`
}

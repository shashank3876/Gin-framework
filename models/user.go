package models

import (
	"errors"
	"go-project/db"
	"go-project/utils"
)

// User represents a user in the system
type User struct {
	ID       int64  `json:"id"`
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

// Save inserts a new user into the database
func (u *User) Save() error {
	// Prepare insert query
	query := "INSERT INTO users(name, email, password) VALUES (?, ?, ?)"
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	// Hash password before saving
	hashedPassword, err := utils.HashPassword(u.Password)
	if err != nil {
		return err
	}

	// Execute insert
	result, err := stmt.Exec(u.Name, u.Email, hashedPassword)
	if err != nil {
		return err
	}

	// Get last inserted ID
	userId, err := result.LastInsertId()
	if err != nil {
		return err
		/*************  ✨ Windsurf Command ⭐  *************/
		// ValidateCredentials checks if the provided password matches the stored hashed password
		// for the corresponding user email. If the credentials are invalid, it returns an error.
		// If the credentials are valid, it returns nil.
		/*******  0c6327b8-9a69-4350-9997-616ef503ca4d  *******/
	}

	u.ID = userId
	return nil
}

// ValidateCredentials checks if the user's email and password are correct
func (u *User) ValidateCredentials() error {
	query := "SELECT id, name, password FROM users WHERE email = ?"
	row := db.DB.QueryRow(query, u.Email)

	var retrievedPassword string
	var retrievedName string
	err := row.Scan(&u.ID, &retrievedName, &retrievedPassword)
	if err != nil {
		return errors.New("credentials invalid")
	}

	// Update name from DB
	u.Name = retrievedName

	// Verify password
	if !utils.CheckPasswordHash(u.Password, retrievedPassword) {
		return errors.New("credentials invalid")
	}

	return nil
}

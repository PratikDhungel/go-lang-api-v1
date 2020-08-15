package services

import (
	"encoding/json"
	"fmt"
	"github.com/quality-dashboard-api/api/config"
)

// User - User Info Structure
type User struct{

	UserID			int 	`json:"user_id"`
	FirstName		string	`json:"first_name"`
	LastName		string	`json:"last_name"`
	Email			string	`json:"email"`
	Address			string	`json:"address"`
}

// GetAllUsers - Returns list of all Users
func GetAllUsers() []byte{

	fmt.Println("Inside GetAllUsers method")

	results, err := config.DB.Query("SELECT id, first_name, last_name, email, address FROM users WHERE id = 1")

	if err != nil{

		panic(err.Error()) 
	}

	// var users []User
	// var user User
	var usersBytes []byte

	for results.Next(){

		var user User

		var userID int
		var firstName, lastName, email, address string

		results.Scan(&userID, &firstName, &lastName, &email, &address)
		user = User{userID, firstName, lastName, email, address}
		// users = append(users, User{user_id, first_name, last_name, email, address})

		usersBytes, _ = json.Marshal(user)

	}
	return usersBytes
}

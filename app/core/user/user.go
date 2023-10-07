package user

import (
	"errors"
	"reflect"
	"time"
)

func (u *User) GetUserName() string {
	return u.FirstName + " " + u.LastName
}

// CreateUser creates a new user and returns the user instance.
func CreateUser(u *CreateUserDTO) *User {

	user := User{
		ID:        1,
		FirstName: "John",
		LastName:  "Doe",
		Email:     "john.doe@example.com",
		Password:  "mypassword",
		Role:      ADMINISTRATOR,
		Image:     "profile.jpg",
		UpdatedAt: time.Now(),
		CreatedAt: time.Now(),
	}
	return &user
}

func GetUser(userID int) *User {
	// Simulate fetching user data from MongoDB or storage.
	// Replace this with your actual data retrieval logic.
	u := User{
		ID:        1,
		FirstName: "John",
		LastName:  "Doe",
		Email:     "john.doe@example.com",
		Password:  "mypassword",
		Role:      ADMINISTRATOR,
		Image:     "profile.jpg",
		UpdatedAt: time.Now(),
		CreatedAt: time.Now(),
	}
	return &u
}

func GetUserByName(name string) (bool, error) {
	if reflect.TypeOf(name).Kind() == reflect.String {
		return false, errors.New("Name must be a string")
	}
	return false, nil
}

// ModifyUser updates the user's information.
func (u *User) ModifyUser(firstName, lastName, email string) {
	// Update the user's information.
	u.FirstName = firstName
	u.LastName = lastName
	u.Email = email
}

// DeleteUser deletes the user.
func (u *User) DeleteUser() {
	// Simulate deleting the user from MongoDB or storage.
	// Replace this with your actual user deletion logic.
	//fmt.Printf("User %s (ID: %s) has been deleted.\n", u.GetFullName(), u.ID.String())
}

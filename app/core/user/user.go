package user

import (
	"time"
)

type createUserDTO struct {
	FirstName string   `json:"firstName"`
	LastName  string   `json:"lastName"`
	Email     string   `json:"email"`
	Password  string   `json:"password"`
	Role      userRole `json:"role"`
	Image     string   `json:"image"`
}

type editUserDTO struct {
	FirstName *string   `json:"firstName"`
	LastName  *string   `json:"lastName"`
	Email     *string   `json:"email"`
	Password  *string   `json:"password"`
	Role      *userRole `json:"role"`
	Image     *string   `json:"image"`
}

type User struct {
	//uuid        uuid.UUID
	ID        int
	FirstName string    `json:"firstName"`
	LastName  string    `json:"lastName"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	Role      userRole  `json:"role"`
	Image     string    `json:"image"`
	UpdatedAt time.Time `json:"updateddAt"`
	CreatedAt time.Time `json:"createdAt"`
}

// NewUser creates a new User instance with a randomly generated UUID.
func NewUser(User) *User {
	return &User{
		ID:        id,
		FirstName: firstName,
		LastName:  lastName,
		Email:     email,
	}
}

func (u *User) GetUserName() string {
	return u.FirstName + " " + u.LastName
}

// CreateUser creates a new user and returns the user instance.
func CreateUser(u *createUserDTO) *User {

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
	return NewUser("John", "Doe", "john@example.com")
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

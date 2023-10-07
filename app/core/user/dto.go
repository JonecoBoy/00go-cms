package user

import "time"

type CreateUserDTO struct {
	FirstName string   `json:"firstName"`
	LastName  string   `json:"lastName"`
	Email     string   `json:"email"`
	Password  string   `json:"password"`
	Role      userRole `json:"role"`
	Image     string   `json:"image"`
}

type EditUserDTO struct {
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

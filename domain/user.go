package domain

type UserRole int

const (
	Admin   UserRole = 0
	Teacher UserRole = 1
	Student UserRole = 2
)

type User struct {
	ID       int64    `json:"id" gorm:"primary_key"`
	Name     string   `json:"name"`
	Email    string   `json:"email"`
	Password string   `json:"password"`
	Role     UserRole `json:"role"`
}

type UserResponse struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

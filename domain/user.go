package domain

type UserRole int

const (
	Admin   UserRole = -1
	Teacher UserRole = 1
	Student UserRole = 2
)

type User struct {
	ID       int64    `json:"id" gorm:"primary_key"`
	Name     string   `json:"name"`
	Email    string   `json:"email" gorm:"unique_index"`
	Password string   `json:"password"`
	Role     UserRole `json:"role"`
}

type CreateUserRequest struct {
	Name     string   `json:"name" binding:"required"`
	Email    string   `json:"email" binding:"required"`
	Password string   `json:"password" binding:"required"`
	Role     UserRole `json:"role"`
}

type UpdateUserRequest struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type ChangePasswordRequest struct {
	OldPassword string `json:"old_password" binding:"required"`
	NewPassword string `json:"new_password" binding:"required"`
}

type ChangeRoleRequest struct {
	Role UserRole `json:"role" binding:"required"`
}

type UserResponse struct {
	ID    int64    `json:"id"`
	Name  string   `json:"name"`
	Email string   `json:"email"`
	Role  UserRole `json:"role"`
}

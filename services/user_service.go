package services

import (
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/impal-lms/lms-backend/domain"
	"github.com/impal-lms/lms-backend/repository"
)

type CreateUserRequest struct {
	Name     string
	Email    string
	Password string
	Role     domain.UserRole
}

type LoginRequest struct {
	Email    string
	Password string
}

type UserResponse struct {
	Name  string
	Email string
	Role  domain.UserRole
}

type UserService struct {
	Repository repository.UserRepository
}

func (service UserService) GenerateToken(user domain.User) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	// Set claims
	claims := token.Claims.(jwt.MapClaims)
	claims["user"] = strconv.Itoa(int(user.Id))
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	// Generate encoded token and send it as response.
	return token.SignedString([]byte("secret"))
}

func (service *UserService) CreateUser(requesterId int64, req CreateUserRequest) (userId int64, err error) {
	requester, err := service.Repository.GetUserByID(requesterId)
	if err != nil || requester.Role != domain.Admin {
		return 401, err
	}

	user, err := domain.NewUser(req.Name, req.Email, req.Password, req.Role)
	if err != nil {
		return 500, err
	}

	return user.Id, service.Repository.Save(*user)
}

func (service UserService) GetUserByID(userId int64) (*UserResponse, error) {
	user, err := service.Repository.GetUserByID(userId)
	if err != nil {
		return nil, err
	}

	return &UserResponse{
		user.Name,
		user.Email,
		user.Role,
	}, nil
}

func (service UserService) ValidateToken(tokenString string) (userId int64, err error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if method, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		} else if method != jwt.SigningMethodHS256 {
			return nil, errors.New("Signing method invalid")
		}

		return []byte("secret"), nil
	})

	if err != nil {
		return -1, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return -1, errors.New("not ok when when converting token")
	}

	if claims["user"] != nil {
		fmt.Println(claims["user"])
		id, err := strconv.Atoi(claims["user"].(string))
		if err != nil {
			return -1, err
		}

		return int64(id), nil
	}

	return -1, errors.New("invalid token")
}

func (service UserService) Login(req LoginRequest) (token string, id int64, err error) {
	user, err := service.Repository.Authenticate(req.Email, req.Password)
	if err != nil {
		return "", -1, err
	}

	token, err = service.GenerateToken(*user)
	id = user.Id
	return
}

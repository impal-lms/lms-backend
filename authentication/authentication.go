package authentication

import (
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/impal-lms/lms-backend/domain"
	"golang.org/x/crypto/bcrypt"
)

type Authentication interface {
	GenerateToken(user domain.User) (string, error)
	ValidateToken(tokenString string) (userID int64, err error)
	HashPassword(password string) (string, error)
	CheckPasswordHash(hash, password string) bool
}

type JWTAuthentication struct {
}

func NewAuthentication() *JWTAuthentication {
	return &JWTAuthentication{}
}

func (auth *JWTAuthentication) GenerateToken(user domain.User) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	// Set claims
	claims := token.Claims.(jwt.MapClaims)
	claims["user"] = strconv.Itoa(int(user.ID))
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	// Generate encoded token and send it as response.
	return token.SignedString([]byte("secret"))
}

func (auth *JWTAuthentication) ValidateToken(tokenString string) (userID int64, err error) {
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

func (auth *JWTAuthentication) HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func (auth *JWTAuthentication) CheckPasswordHash(hash, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

package service

import (
	"fmt"
	"log"

	model "github.com/ICOMP-UNC/2023---soii---laboratorio-6-FrancoNB/models"
	repository "github.com/ICOMP-UNC/2023---soii---laboratorio-6-FrancoNB/repositories"
	"github.com/dgrijalva/jwt-go"
)

func GenerateToken(userID string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)

	claims["user_id"] = userID

	tokenString, err := token.SignedString([]byte("secret_key"))

	if err != nil {
		log.Println(err)
		return "", err
	}

	return tokenString, nil
}

func ValidateUser(user model.User) (string, error) {
	sysUser, err := repository.UserGet(user.Username)

	if err != nil {
		return "", err
	}

	err = repository.UserValidateCredentials(user.Username, user.Password)

	if err != nil {
		return "", err
	}

	token, err := GenerateToken(fmt.Sprint(sysUser.Id))

	if err != nil {
		return "", err
	}

	return token, nil
}

func NewUser(user model.User) (model.User, error) {
	var result model.User

	_, err := repository.UserGet(user.Username)

	if err == nil {
		return result, fmt.Errorf("User '%s' already exists", user.Username)
	}

	result, err = repository.UserCreate(user)

	if err != nil {
		repository.UserDelete(user.Username)
		return result, err
	}

	return result, nil
}

func ListAllUsers() ([]model.User, error) {
	return repository.UserGetAll()
}

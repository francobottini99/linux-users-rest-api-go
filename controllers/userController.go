package controller

import (
	"net/http"

	model "github.com/ICOMP-UNC/2023---soii---laboratorio-6-FrancoNB/models"
	service "github.com/ICOMP-UNC/2023---soii---laboratorio-6-FrancoNB/services"
	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	var userLogin model.User

	if err := c.ShouldBindJSON(&userLogin); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Bad request",
		})
		return
	}

	token, err := service.ValidateUser(userLogin)

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"message": "Login successful !",
			"token":   token,
		})
	}
}

func Register(c *gin.Context) {
	var userRegister model.User

	if err := c.ShouldBindJSON(&userRegister); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Bad request",
		})
		return
	}

	result, err := service.NewUser(userRegister)

	if err != nil {
		c.JSON(http.StatusConflict, gin.H{
			"message": err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, result)
	}
}

func ListAll(c *gin.Context) {
	users, err := service.ListAllUsers()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
	} else {
		if len(users) == 0 {
			c.JSON(http.StatusOK, gin.H{
				"message": "No users found !",
			})
		} else {
			c.JSON(http.StatusOK, users)
		}
	}
}

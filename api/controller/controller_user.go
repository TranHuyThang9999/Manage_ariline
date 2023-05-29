package controller

import (
	"btl/api/middleware"
	"btl/infrastructure/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (t *RepositoryController) UpdateProflie(c *gin.Context) {
	var user *model.UserUpdate
	name := c.Param("user_name")
	phone := c.Param("phone_number")
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error 1": err.Error()})
		return
	}
	status, err := t.ctrl.UpdateUser(c, user, name, phone)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error 2": err.Error()})
		return
	}
	t.Success(c, map[string]bool{
		"is_update profile": status,
	})
}
func (t *RepositoryController) UpdatePassword(c *gin.Context) {
	phone := c.Param("phone_number")
	oldPassword := c.Param("oldPassword")
	newPassword := c.Param("newPassword")
	status, err := t.ctrl.UpdatePassword(c, phone, oldPassword, newPassword)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error 2": err.Error()})
		return
	}
	t.Success(c, map[string]bool{
		"is_update password": status,
	})

}

func (t *RepositoryController) Login(c *gin.Context) {
	var user model.UserLogin
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}
	status, err := t.ctrl.LoginUser(c, &user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}
	if !status {
		c.JSON(http.StatusOK, gin.H{"error": "Incorrect phone number or password"})
		return
	}
	token, err := middleware.GenerateJWT(user.PhoneNumber, "lkk")

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}
	t.saveSession(c, user.PhoneNumber, token)
	t.SuccessToken(c, token)

}

func (t *RepositoryController) CreateUser(c *gin.Context) {
	var user model.User
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error 1": err.Error()})
		return
	}
	user_s, _ := t.ctrl.FindByPhoneNumber(c, user.PhoneNumber)
	if user_s != nil {
		c.JSON(http.StatusConflict, gin.H{"error": "user already exists"})
		return
	}
	status, err := t.ctrl.CreateAccountUser(c.Request.Context(), &user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error 2": err.Error()})
		return
	}

	if !status {
		c.JSON(http.StatusConflict, gin.H{"error": "create error"})
		return
	}
	t.Success(c, map[string]bool{
		"is_create": status,
	})
}

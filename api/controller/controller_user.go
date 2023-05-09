package controller

import (
	"btl/api/middware"
	"btl/core/user_case"
	"btl/infra/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

type RepositoryControoler struct {
	ctrl *user_case.RepositoryUserCase
}

func NewController(ctl *user_case.RepositoryUserCase) *RepositoryControoler {
	return &RepositoryControoler{
		ctrl: ctl,
	}
}

//func (t *RepositoryControoler) FindAll(c *gin.Context) {
//	user, err := t.ctrl.FindALlUser(c)
//	if err != nil {
//		c.JSON(http.StatusInternalServerError, gin.H{"error 2": err.Error()})
//		return
//	}
//	c.JSON(http.StatusOK, gin.H{"Infor ": user})
//
//}

func (t *RepositoryControoler) UpdateProflie(c *gin.Context) {
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
func (t *RepositoryControoler) UpdatePassword(c *gin.Context) {
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

func (t *RepositoryControoler) Login(c *gin.Context) {
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
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Incorrect phone number or password"})
		return
	}
	token, err := middware.GenerateJWT(user.Password, user.Password)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	t.Success(c, token)

}

func (t *RepositoryControoler) CreateUser(c *gin.Context) {
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

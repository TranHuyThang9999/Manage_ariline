package controller

import (
	"btl/api/middleware"
	"btl/infrastructure/model"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func (ctxadmin *RepositoryController) CreateAccountAdmin(c *gin.Context) {
	var user model.Admin
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error 1": err.Error()})
		return
	}
	user_s, _ := ctxadmin.ctrl.FindByPhoneNumberAdmin(c, user.PhoneNumber)
	if user_s != nil {
		c.JSON(http.StatusConflict, gin.H{"error 2": "user already exists"})
		return
	}
	status, err := ctxadmin.ctrl.CreateAccountAdmin(c, &user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error 3": err.Error()})
		return
	}
	ctxadmin.Success(c, map[string]bool{
		"is_create": status,
	})
}
func (ctxadmin *RepositoryController) LoginAdmin(c *gin.Context) {
	var user model.UserLogin

	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error 1": err.Error()})
		return
	}
	status, err := ctxadmin.ctrl.LoginAdmin(c, &user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error 2": err.Error()})
		return
	}
	if !status {
		c.JSON(http.StatusOK, gin.H{"error": "Incorrect phone number or password"})
		return
	}
	AccessExpire := time.Now().Unix()
	access_token, err := middleware.GenerateJWT(AccessExpire)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error 3": status})
	}
	c.SetCookie("access_token", access_token, int(time.Hour)*60, "/", "localhost", false, true)
	ctxadmin.SuccessToken(c, access_token)
}
func (ctxadmin *RepositoryController) FindByFormAccount(c *gin.Context) {
	var user model.UserByForm
	if err := c.ShouldBind(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error 1": err.Error()})
		return
	}
	users, err := ctxadmin.ctrl.FindByForm(c, &user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error 2": err.Error()})
		return
	}
	ctxadmin.Success(c, users)
}

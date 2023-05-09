package controller

import (
	"btl/api/middware"
	"btl/infra/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (ctxadmin *RepositoryControoler) CreateAccountAdmin(c *gin.Context) {
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
	c.JSON(http.StatusOK, gin.H{"create account sucess": status})
}
func (ctxadmin *RepositoryControoler) LoginAdmin(c *gin.Context) {
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

	token, err := middware.GenerateJWT(user.Password, user.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error 3": status})
	}
	c.JSON(http.StatusOK, gin.H{"login sucess ": token})
}
func (ctxadmin *RepositoryControoler) FindByFormAccount(c *gin.Context) {
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
	c.JSON(http.StatusOK, gin.H{"Info": users})
}

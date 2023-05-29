package controller

import (
	"btl/api/resources"
	"btl/core/enums"
	"btl/core/user_case"
	"github.com/gin-contrib/sessions"
	"net/http"

	"github.com/gin-gonic/gin"
)

type RepositoryController struct {
	ctrl *user_case.RepositoryUserCase
}

func NewController(ctl *user_case.RepositoryUserCase) *RepositoryController {
	return &RepositoryController{
		ctrl: ctl,
	}
}

func (b *RepositoryController) Success(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, resources.NewResponseResource(enums.SuccessCode, "Success", data))
}

func (b *RepositoryController) SuccessToken(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, resources.NewResponseResourceToken(enums.SuccessCode, "Success", data))
}
func (t *RepositoryController) saveSession(c *gin.Context, phoneNumber, token string) {
	// Lưu thông tin phiên đăng nhập vào cookie
	session := sessions.Default(c)
	session.Set("phoneNumber", phoneNumber)
	session.Set("token", token)
	session.Save()
}

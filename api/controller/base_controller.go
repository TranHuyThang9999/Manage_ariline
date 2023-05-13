package controller

import (
	"btl/api/resources"
	"btl/core/enums"
	"btl/core/user_case"
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

func (b *RepositoryControoler) Success(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, resources.NewResponseResource(enums.SuccessCode, "Success", data))
}

func (b *RepositoryControoler) SuccessToken(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, resources.NewResponseResourceToken(enums.SuccessCode, "Success", data))
}

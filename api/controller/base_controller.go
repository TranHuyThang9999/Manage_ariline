package controller

import (
	"btl/api/resources"
	"btl/core/enums"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (b *RepositoryControoler) Success(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, resources.NewResponseResource(enums.SuccessCode, "Success", data))
}

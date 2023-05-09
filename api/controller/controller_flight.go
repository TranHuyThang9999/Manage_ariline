package controller

import (
	"btl/infra/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (ft *RepositoryControoler) CreateFlight(c *gin.Context) {
	var flight model.Flight
	if err := c.BindJSON(&flight); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error 1": err.Error()})
		return
	}
	status, err := ft.ctrl.CreateFilght(c, &flight)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error 2": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"Create sucessful": status})
}
func (ft *RepositoryControoler) UpdateFlight(c *gin.Context) {
	id := c.Param("flight_id")
	name := c.Param("name_flight")
	var flight model.Flight
	if err := c.BindJSON(&flight); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error 1": err.Error()})
		return
	}
	status, err := ft.ctrl.UpdateFlight(c, id, name, &flight)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error 2": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"Status Update": status})
}
func (ft *RepositoryControoler) DeleteFlight(c *gin.Context) {
	id := c.Param("flight_id")
	name := c.Param("name_flight")
	status, err := ft.ctrl.DeleteFlight(c, id, name)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error 2": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"Status Update": status})
}
func (ft *RepositoryControoler) FindByFormFlight(c *gin.Context) {
	var flight model.FlightByForm
	err := c.ShouldBind(&flight)
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"status": "fail", "message": err.Error()})
		return
	}
	users, err := ft.ctrl.FindBFlightByForm(c, flight)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error 2": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"Info": users})
}

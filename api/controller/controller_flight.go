package controller

import (
	"btl/core/cache/helper"
	"btl/infrastructure/model"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func (ft *RepositoryController) CreateFlight(c *gin.Context) {
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
func (ft *RepositoryController) UpdateFlight(c *gin.Context) {
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
func (ft *RepositoryController) DeleteFlight(c *gin.Context) {
	id := c.Param("flight_id")
	name := c.Param("name_flight")
	status, err := ft.ctrl.DeleteFlight(c, id, name)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error 2": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"Status Update": status})
}
func (ft *RepositoryController) FindByFormFlight(c *gin.Context) {
	var flight model.FlightByForm

	err := c.ShouldBind(&flight)
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"status": "fail", "message": err.Error()})
		return
	}
	cachedData, err := helper.GetValueCache("infor flights")
	if err == nil && cachedData != nil {
		var filghts []*model.Flight
		err := json.Unmarshal(cachedData, &filghts)
		if err != nil {
			fmt.Println(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
			return
		}
		ft.Success(c, filghts)
		return
	}

	info_flights, err := ft.ctrl.FindBFlightByForm(c, flight)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error 2": err.Error()})
		return
	}
	data := info_flights
	cachedData, err = helper.SetValueCache(c, "infor flights", data, time.Minute*1)
	if err != nil {
		fmt.Println(err)
	}
	//c.JSON(http.StatusOK, gin.H{"data": data})
	ft.Success(c, data)
}

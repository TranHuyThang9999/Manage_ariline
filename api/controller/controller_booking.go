package controller

import (
	"btl/config"
	"btl/core/cache/helper"
	"btl/infrastructure/model"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func (tck *RepositoryController) RegisterTicket(c *gin.Context) {
	var booking model.BookingRequest
	if err := c.BindJSON(&booking); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error 1": err.Error()})
		return
	}
	status, err := tck.ctrl.RegisterTicket(c, &booking)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error 2": err.Error()})
		return
	}
	tck.Success(c, map[string]bool{
		"is_register tocket": status,
	})
}
func (tck *RepositoryController) CanCelTicket(c *gin.Context) {
	phone_number := c.Param("phone_number")
	booking_id := c.Param("booking_id")

	ticket, err := tck.ctrl.GetStatusTicket(c, phone_number, booking_id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error 1": err.Error()})
		return
	}
	if ticket.Status == "ticket canceled" {
		c.JSON(http.StatusBadRequest, gin.H{"Tick canceled:": ticket.Status})
		return
	}
	status, err := tck.ctrl.CanCelTicket(c, phone_number, booking_id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error 2": err.Error()})
		return
	}
	tck.Success(c, map[string]bool{
		"is_cancel": status,
	})
}
func (tck *RepositoryController) GetAllTicket(c *gin.Context) {

	tickets, err := tck.ctrl.GetAllTicket(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error 2": err.Error()})
		return
	}
	tck.Success(c, tickets)
}
func (tck *RepositoryController) GetTicketByPhoneNumber(c *gin.Context) {
	phone_number := c.Param("phone_number")
	tickets, err := tck.ctrl.GetTicketByPhoneNumber(c, phone_number)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error 2": err.Error()})
		return
	}
	tck.Success(c, tickets)
}
func (tck *RepositoryController) GetAllTicketByForm(c *gin.Context) {

	var ticketRequest model.BookingByForm

	config, err := config.LoadConfig("config/config.yaml")
	if err != nil {
		log.Fatal(err)
		return
	}

	err = c.ShouldBind(&ticketRequest)
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"status": "fail", "message": err.Error()})
		return
	}
	cacheData, err := helper.GetValueCache("infor ticketed")
	if err != nil && cacheData != nil {
		var booking []*model.Booking
		err := json.Unmarshal(cacheData, &booking)
		if err != nil {
			fmt.Println(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"data": booking})
		return
	}
	tickets, err := tck.ctrl.GetAllTicketByForm(c, ticketRequest)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error 2": err.Error()})
		return
	}
	data := tickets
	cacheData, err = helper.SetValueCache(c, "infor ticketed", data, time.Minute*time.Duration(config.Expiretime.Expiration))
	if err != nil {
		fmt.Println(err)
	}
	tck.Success(c, data)
}

package controller

import (
	"btl/infra/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (tck *RepositoryControoler) RegisterTicket(c *gin.Context) {
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
	c.JSON(http.StatusOK, gin.H{"RegisTicket sucess": status})
}
func (tck *RepositoryControoler) CanCelTicket(c *gin.Context) {
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
	c.JSON(http.StatusOK, gin.H{"Calcel sucess": status})
}
func (tck *RepositoryControoler) GetAllTicket(c *gin.Context) {
	data, err := tck.ctrl.GetAllTicket(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error 2": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"Calcel sucess": data})
}
func (tck *RepositoryControoler) GetTicketByPhoneNumber(c *gin.Context) {
	phone_number := c.Param("phone_number")
	data, err := tck.ctrl.GetTicketByPhoneNumber(c, phone_number)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error 2": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"Code 0 :": data})
}

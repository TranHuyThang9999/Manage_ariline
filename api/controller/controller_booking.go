package controller

import (
	"btl/infra/model"
	"net/http"

	"github.com/gin-gonic/gin"
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
	tck.Success(c, map[string]bool{
		"is_register tocket": status,
	})
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
	tck.Success(c, map[string]bool{
		"is_cancel": status,
	})
}
func (tck *RepositoryControoler) GetAllTicket(c *gin.Context) {
	tickets, err := tck.ctrl.GetAllTicket(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error 2": err.Error()})
		return
	}
	tck.Success(c, tickets)
}
func (tck *RepositoryControoler) GetTicketByPhoneNumber(c *gin.Context) {
	phone_number := c.Param("phone_number")
	tickets, err := tck.ctrl.GetTicketByPhoneNumber(c, phone_number)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error 2": err.Error()})
		return
	}
	tck.Success(c, tickets)
}

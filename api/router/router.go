package router

import (
	"btl/api/controller"
	"btl/api/middleware"
	"btl/config"
	"btl/core/user_case"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func NewRouter() (*gin.Engine, error) {
	r := gin.Default()

	//r.Use(cors.Default())
	//	r.Use(middleware.CORS())
	// store := cookie.NewStore([]byte("yangyaning"))
	// r.Use(sessions.Sessions("mysession", store))

	//r.Use(cors.Default())
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"PUT", "PATCH", "GET", "POST", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))
	config, err := config.Connect("config/config.yaml")
	if err != nil {
		return nil, err
	}
	user_case := user_case.NewPorts(config)
	controller_user := controller.NewController(user_case)

	r.POST("/user/create", controller_user.CreateUser)
	r.POST("/user/login", controller_user.Login)

	//user
	api_user := r.Group("/user", middleware.Auth())
	{

		api_user.POST("/register/ticket", controller_user.RegisterTicket)
		api_user.PATCH("/cancel/ticket/:phone_number/:booking_id", controller_user.CanCelTicket)
		api_user.PATCH("/reset/:phone_number/:oldPassword/:newPassword", controller_user.UpdatePassword)

		api_user.GET("/info/flight", controller_user.FindByFormFlight)
		api_user.GET("/info/ticket/phone_number/:phone_number", controller_user.GetTicketByPhoneNumber)
		api_user.GET("/logout", controller_user.Logout)
	}

	/// admin
	r.POST("/admin/create", controller_user.CreateAccountAdmin)
	r.POST("/admin/login", controller_user.LoginAdmin)

	api_admin := r.Group("/admin", middleware.Auth())
	{

		api_admin.POST("/create/flight", controller_user.CreateFlight)
		api_admin.PATCH("/update/flight/:flight_id/:name_flight", controller_user.UpdateFlight)
		api_admin.DELETE("/delete/flight/:flight_id/:name_flight", controller_user.DeleteFlight)

		api_admin.GET("/info/user", controller_user.FindByFormAccount)
		api_admin.GET("/info/flight", controller_user.FindByFormFlight)
		api_admin.GET("/infor/tickets", controller_user.GetAllTicketByForm)
		api_admin.GET("/logout", controller_user.Logout)
	}

	return r, nil
}

package user_case

import (
	"btl/core/port"
	"btl/infrastructure/database/postGresql"

	"gorm.io/gorm"
)

type RepositoryUserCase struct {
	user_case       port.RepositoryUser
	admin_user_case port.RepositoryAdmin
	flight_use_case port.RepositoryFlight
	booking         port.RepositoryBooking
}

func NewPorts(db *gorm.DB) *RepositoryUserCase {
	pgRepo := postGresql.NewPostGresql(db)
	return &RepositoryUserCase{
		user_case:       pgRepo,
		admin_user_case: pgRepo,
		flight_use_case: pgRepo,
		booking:         pgRepo,
	}
}

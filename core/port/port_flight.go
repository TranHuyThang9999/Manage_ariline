package port

import (
	"btl/infrastructure/model"
	"context"
)

// admin
type RepositoryFlight interface {
	CreateFlight(ctx context.Context, flight *model.Flight) (bool, error)
	UpdateFlight(ctx context.Context, flightID string, nameFlight string, updateFlight *model.Flight) (bool, error)
	DeleteFlight(ctx context.Context, flight_id string, name string) (bool, error)
	FindBFlightByForm(ctx context.Context, flight model.FlightByForm) ([]*model.Flight, error)
}

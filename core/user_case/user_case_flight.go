package user_case

import (
	"btl/infra/model"
	"context"
)

func (fl *RepositoryUserCase) CreateFilght(ctx context.Context, flight *model.Flight) (bool, error) {
	status, err := fl.flight_use_case.CreateFlight(ctx, flight)
	if err != nil {
		return false, err
	}
	return status, nil
}
func (fl *RepositoryUserCase) UpdateFlight(ctx context.Context, flightID string, nameFlight string, updateFlight *model.Flight) (bool, error) {
	status, err := fl.flight_use_case.UpdateFlight(ctx, flightID, nameFlight, updateFlight)
	if err != nil {
		return false, err
	}
	return status, nil
}
func (fl *RepositoryUserCase) DeleteFlight(ctx context.Context, flight_id string, name string) (bool, error) {
	status, err := fl.flight_use_case.DeleteFlight(ctx, flight_id, name)
	if err != nil {
		return false, err
	}
	return status, nil
}

// func (fl *RepositoryUserCase) FindAllFlight(ctx context.Context) ([]*model.Flight, error) {
//
//	return nil, nil
//
// }
func (fl *RepositoryUserCase) FindBFlightByForm(ctx context.Context, flight model.FlightByForm) ([]*model.Flight, error) {
	info_flight, err := fl.flight_use_case.FindBFlightByForm(ctx, flight)
	if err != nil {
		return nil, err
	}
	return info_flight, nil

}

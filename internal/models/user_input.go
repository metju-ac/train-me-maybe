package models

import openapiclient "github.com/metju-ac/train-me-maybe/openapi"

type UserInput struct {
	DepartingStation *StationModel
	ArrivingStation  *StationModel
	SelectedRoute    *openapiclient.SimpleRoute
	RouteDetail      *openapiclient.Route
	SeatClasses      []string
	Tariff           *openapiclient.Tariff
	Section          *openapiclient.SimpleSection
}

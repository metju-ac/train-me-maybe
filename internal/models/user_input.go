package models

import openapiclient "github.com/metju-ac/train-me-maybe/openapi"

type UserInput struct {
	DepartingStation   *StationModel
	ArrivingStation    *StationModel
	SelectedRouteIds   string
	RouteDetail        *openapiclient.Route
	SeatClasses        []string
	TariffKey          string
	Section            *openapiclient.SimpleSection
	CreditUserNumber   string
	CreditUserPassword string
	UserEmail          string
}

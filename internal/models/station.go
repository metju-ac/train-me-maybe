package models

type StationModel struct {
	Country        string `json:"country"`        // Country name
	City           string `json:"city"`           // City name
	StationID      int64  `json:"stationId"`      // Station ID
	StationName    string `json:"stationName"`    // Station name
	IsTrainStation bool   `json:"isTrainStation"` // Indicates if the station is a train station
	IsBusStation   bool   `json:"isBusStation"`   // Indicates if the station is a bus station
}

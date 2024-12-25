package models

type StationModel struct {
	Country        string // Country name
	City           string // City name
	StationID      int64  // Station ID
	StationName    string // Station name
	IsTrainStation bool   // Indicates if the station is a train station
	IsBusStation   bool   // Indicates if the station is a bus station
}

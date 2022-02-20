package models

type CourierLocation struct {
	CourierId       int64   `json:"courierId" db:"driver_id"`
	Latitude        float64 `json:"latitude" db:"latitude"`
	Longitude       float64 `json:"longitude" db:"longitude"`
	Speed           float64 `json:"speed" db:"speed"`
	Accuracy        float64 `json:"accuracy" db:"accuracy"`
	Azimuth         float64 `json:"azimuth" db:"azimuth"`
	CreatedDateTime float64 `db:"created_datetime"`
}

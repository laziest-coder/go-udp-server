package models

type CourierLocation struct {
	CourierId       int64   `db:"driver_id"`
	Latitude        float64 `db:"latitude"`
	Longitude       float64 `db:"longitude"`
	Speed           float64 `db:"speed"`
	Accuracy        float64 `db:"accuracy"`
	Azimuth         float64 `db:"azimuth"`
	CreatedDateTime float64 `db:"created_datetime"`
}

package models

import "errors"

type CourierLocation struct {
	CourierId       int64   `json:"courier_id" db:"driver_id"`
	Latitude        float64 `json:"latitude" db:"latitude"`
	Longitude       float64 `json:"longitude" db:"longitude"`
	Speed           float64 `json:"speed" db:"speed"`
	Accuracy        float64 `json:"accuracy" db:"accuracy"`
	Azimuth         float64 `json:"azimuth" db:"azimuth"`
	CreatedDateTime float64 `db:"created_datetime"`
}

func (c CourierLocation) ValidateData() error {
	if !(c.CourierId > 0) {
		return errors.New("missing courier id field")
	}
	if !(c.Latitude > 0) {
		return errors.New("missing latitude field")
	}
	if !(c.Longitude > 0) {
		return errors.New("missing longitude field")
	}
	return nil
}

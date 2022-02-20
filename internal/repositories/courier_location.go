package repositories

import (
	"fmt"
	"github.com/Express-24/courier-location-tracker/internal/factory"
	"github.com/Express-24/courier-location-tracker/internal/models"
)

type CourierLocationRepository interface {
	InsertLocation(
		courierId int, latitude float64, longitude float64,
		speed float64, accuracy float64, azimuth float64) (cl models.CourierLocation)
}

type SqlRepository struct {
	db factory.DB
}

func (d SqlRepository) InsertLocation(
	courierId int, latitude float64, longitude float64,
	speed float64, accuracy float64, azimuth float64) error {

	var query = d.queryInsertLocation(courierId, latitude, longitude, speed, accuracy, azimuth)

	return d.db.Exec(query)
}

func (d SqlRepository) queryInsertLocation(
	courierId int, latitude float64, longitude float64,
	speed float64, accuracy float64, azimuth float64) string {
	return "INSERT INTO ex24_drivers_location (" +
		fmt.Sprint(courierId) + "," +
		fmt.Sprint(latitude) + "," +
		fmt.Sprint(longitude) + "," +
		fmt.Sprint(speed) + "," +
		fmt.Sprint(accuracy) + "," +
		fmt.Sprint(azimuth) + "," +
		")"
}

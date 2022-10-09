package repositories

import (
	"fmt"
	"github.com/Express-24/courier-location-tracker/internal/factory"
)

type CourierLocation interface {
	InsertLocation(
		courierId int, latitude float64, longitude float64,
		speed float64, accuracy float64, azimuth float64) error
}

type SqlRepository struct {
	db factory.DB
}

func CourierLocationRepository(db factory.DB) CourierLocation {
	return &SqlRepository{db: db}
}

func (d SqlRepository) InsertLocation(
	courierId int, latitude float64, longitude float64,
	speed float64, accuracy float64, azimuth float64) error {

	var query = d.queryInsertLocationQuery(courierId, latitude, longitude, speed, accuracy, azimuth)

	return d.db.Exec(query)
}

func (d SqlRepository) queryInsertLocationQuery(
	courierId int, latitude float64, longitude float64,
	speed float64, accuracy float64, azimuth float64) string {
	return "INSERT INTO couriers_location (courier_id, latitude, longitude, speed, accuracy, azimuth) VALUES (" +
		fmt.Sprint(courierId) + "," +
		fmt.Sprint(latitude) + "," +
		fmt.Sprint(longitude) + "," +
		fmt.Sprint(speed) + "," +
		fmt.Sprint(accuracy) + "," +
		fmt.Sprint(azimuth) +
		")"
}

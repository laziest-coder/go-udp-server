package main

import (
	"encoding/json"
	"fmt"
	"github.com/Express-24/courier-location-tracker/internal/config"
	"github.com/Express-24/courier-location-tracker/internal/factory"
	"github.com/Express-24/courier-location-tracker/internal/models"
	"github.com/Express-24/courier-location-tracker/internal/repositories"
	"net"
	"strings"
)

const MaxBytesSize = 1024
const LocalIp = "0.0.0.0"

func main() {

	config.Initialize()

	conn, err := net.ListenUDP("udp", &net.UDPAddr{
		Port: config.GetConfig().App.Port,
		IP:   net.ParseIP(LocalIp),
	})
	if err != nil {
		panic(err)
	}

	defer conn.Close()

	for {
		message := make([]byte, MaxBytesSize)
		dataLength, _, err := conn.ReadFromUDP(message[:])
		if err != nil {
			panic(err)
		}

		data := strings.TrimSpace(string(message[:dataLength]))

		cl := models.CourierLocation{}
		err = json.Unmarshal([]byte(data), &cl)

		if err != nil {
			fmt.Println(err) // TODO LOG HERE
		}

		repo := repositories.CourierLocationRepository(factory.GetDBInstance())

		err = repo.InsertLocation(int(cl.CourierId), cl.Latitude, cl.Longitude, cl.Speed, cl.Accuracy, cl.Azimuth)

		if err != nil {
			fmt.Println(err) // TODO LOG HERE
		}
	}
}

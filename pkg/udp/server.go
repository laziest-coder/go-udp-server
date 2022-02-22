package udp

import (
	"encoding/json"
	"github.com/Express-24/courier-location-tracker/internal/factory"
	"github.com/Express-24/courier-location-tracker/internal/models"
	"github.com/Express-24/courier-location-tracker/internal/repositories"
	"net"
	"strings"
)

type Server struct {
	Port          int
	Ip            string
	MaxBufferSize int
}

func (s Server) Start() {
	log := factory.GetLogger()
	conn, err := net.ListenUDP("udp", &net.UDPAddr{
		Port: s.Port,
		IP:   net.ParseIP(s.Ip),
	})
	if err != nil {
		log.Error("Error while creating udp server: ", err)
		panic(err)
	}

	defer conn.Close()

	for {
		message := make([]byte, s.MaxBufferSize)
		dataLength, _, err := conn.ReadFromUDP(message[:])
		if err != nil {
			log.Error("Error while reading data from udp: ", err)
			panic(err)
		}
		if dataLength > s.MaxBufferSize {
			log.WithField("data", dataLength).Error("Data length exceeded max buffer size")
			continue
		}

		data := strings.TrimSpace(string(message[:dataLength]))

		cl := models.CourierLocation{}
		err = json.Unmarshal([]byte(data), &cl)

		if err != nil {
			log.Error("Error while unmarshalling incoming json data: ", err)
			continue
		}

		err = cl.ValidateData()
		if err != nil {
			log.WithField("data", cl).Error("Error while validating incoming json data: ", err)
			continue
		}

		repo := repositories.CourierLocationRepository(factory.GetDBInstance())
		log.WithField("data", cl).Debug("Before inserting courier location to db")
		err = repo.InsertLocation(int(cl.CourierId), cl.Latitude, cl.Longitude, cl.Speed, cl.Accuracy, cl.Azimuth)

		if err != nil {
			log.WithField("data", cl).Error("Error while inserting courier location: ", err)
		}
	}
}

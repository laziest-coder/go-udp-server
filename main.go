package main

import (
	"github.com/Express-24/courier-location-tracker/internal/config"
	"github.com/Express-24/courier-location-tracker/pkg/udp"
)

const MaxBytesSize = 1024
const LocalIp = "0.0.0.0"

func main() {

	config.Initialize()

	server := udp.Server{Port: config.GetConfig().App.Port, Ip: LocalIp, MaxBufferSize: MaxBytesSize}
	server.Start()
}

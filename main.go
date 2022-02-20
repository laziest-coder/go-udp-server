package main

import (
	"fmt"
	"github.com/Express-24/courier-location-tracker/internal/config"
	"net"
	"strings"
)

func main() {

	config.Initialize()

	conn, err := net.ListenUDP("udp", &net.UDPAddr{
		Port: config.GetConfig().App.Port,
		IP:   net.ParseIP("0.0.0.0"),
	})
	if err != nil {
		panic(err)
	}

	defer conn.Close()
	fmt.Printf("server listening %s\n", conn.LocalAddr().String())

	for {
		message := make([]byte, 1024)
		rlen, remote, err := conn.ReadFromUDP(message[:])
		if err != nil {
			panic(err)
		}

		data := strings.TrimSpace(string(message[:rlen]))
		fmt.Printf("received: %s from %s\n", data, remote)
	}
}

package main

import (
	"os"

	eeureka "github.com/puzzledvacuum/backing-fulfillment/eeureka"
	service "github.com/puzzledvacuum/backing-fulfillment/service"
)

func main() {
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "3001"
	}

	eeureka.RegisterAt("http://localhost:8081", "backing-fulfillment", port, "8443")
	// Ordinarily we'd use a CF environment here, but we don't need it for
	// the fake data we're returning.
	server := service.NewServer()
	server.Run(":" + port)
}

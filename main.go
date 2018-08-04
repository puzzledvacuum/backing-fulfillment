package main

import (
	"fmt"
	"os"
	"time"

	eeureka "github.com/puzzledvacuum/backing-fulfillment/eeureka"
	service "github.com/puzzledvacuum/backing-fulfillment/service"
)

func main() {
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "3001"
	}

	eeureka.RegisterAt("http://localhost:8081", "backing-fulfillment", port, "8443")

	ticker := time.NewTicker(500 * time.Millisecond)
	count := 1

	for t := range ticker.C {
		frontend, err := eeureka.GetServiceInstances("backing-catalog")
		if count%4 == 0 {
			fmt.Println("...Waiting for message:", t)
		}
		if err != nil {
			fmt.Printf("Service not available, %v\n", err)
			// return
		}
		if len(frontend) > 0 {
			fmt.Println("HostName:", frontend[0].HostName)
			fmt.Println("Port:", frontend[0].Port.Port)
			break
		}
		count++
		// if count%60 == 0 {
		// 	break
		// }
	}

	// Ordinarily we'd use a CF environment here, but we don't need it for
	// the fake data we're returning.
	server := service.NewServer()
	server.Run(":" + port)
}

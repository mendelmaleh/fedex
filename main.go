package main

import (
	"log"
	"os"

	"github.com/kr/pretty"
	"github.com/pelletier/go-toml/v2"
)

type TomlConfig struct {
	Fedex struct {
		Config
		TrackingNumber string
	}
}

func main() {
	// config
	data, err := os.ReadFile("config.toml")
	if err != nil {
		log.Fatal(err)
	}

	var config TomlConfig
	if err = toml.Unmarshal(data, &config); err != nil {
		log.Fatal(err)
	}

	cl := Client{Config: config.Fedex.Config}

	tracking, err := cl.Track(config.Fedex.TrackingNumber)
	if err != nil {
		log.Fatal(err)
	}

	pretty.Println(tracking)
}

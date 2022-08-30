package main

import (
	"log"
	"os"

	"git.sr.ht/~mendelmaleh/fedex"
	"github.com/kr/pretty"
	"github.com/pelletier/go-toml/v2"
)

type Config struct {
	Fedex struct {
		fedex.Config
		TrackingNumber string
	}
}

func main() {
	// config
	data, err := os.ReadFile("config.toml")
	if err != nil {
		log.Fatal(err)
	}

	var config Config
	if err = toml.Unmarshal(data, &config); err != nil {
		log.Fatal(err)
	}

	// fedex
	cl := fedex.Client{Config: config.Fedex.Config}

	tracking, err := cl.Track(config.Fedex.TrackingNumber)
	if err != nil {
		log.Fatal(err)
	}

	pretty.Println(tracking)
}

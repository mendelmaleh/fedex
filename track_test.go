package main

import (
	"encoding/json"
	"os"
	"testing"
)

func TestTrackingNumbers(t *testing.T) {
	file, err := os.Open("testdata/trackingnumbers.json")
	if err != nil {
		t.Error(err)
	}

	var resp TrackResponse
	if err := json.NewDecoder(file).Decode(&resp); err != nil {
		t.Error(err)
	}

	// t.Errorf("%#+v\n", resp)
}

package state

import (
	"encoding/json"
	"fmt"
	"github.com/rs/zerolog/log"
	"github.com/tlj/tesla"
	"io/ioutil"
	"os"
)

var Current *State

type VehiclesData []tesla.VehicleData

func (v VehiclesData) ByID(id int64) *tesla.VehicleData {
	for idx, _ := range v {
		if v[idx].ID == id {
			return &v[idx]
		}
	}

	return nil
}

func (v VehiclesData) ByVehicleID(vehicleID int64) *tesla.VehicleData {
	for idx, _ := range v {
		if v[idx].VehicleID == vehicleID {
			return &v[idx]
		}
	}

	return nil
}

func (v VehiclesData) ByIDs(ids string) *tesla.VehicleData {
	for idx, _ := range v {
		if fmt.Sprintf("%d", v[idx].VehicleID) == ids {
			return &v[idx]
		}
	}

	return nil
}

type State struct {
	VehiclesData   VehiclesData                   `json:"vehicles_data"`
	VehiclesConfig map[string]tesla.VehicleConfig `json:"vehicles_config"`
}

func InitState() {
	filename := "cmd/mock/fixtures/state.json"
	jsonFile, err := os.Open(filename)
	if err != nil {
		log.Fatal().Msgf("Error opening %s: %s", filename, err.Error())
	}
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)
	err = json.Unmarshal(byteValue, &Current)
	if err != nil {
		log.Fatal().Msgf("Error parsing %s: %s", filename, err.Error())
	}
}

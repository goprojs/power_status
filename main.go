package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/distatus/battery"
)

type Config struct {
	LocationID   string `json:"location_id"`
	LocationName string `json:"location_name"`
}

type indicator struct {
    ElectricityStatus  bool  `json:"estatus"`
    LocationName  string  `json:"location_name"`
    LocationID string  `json:"location_id"`
    CurrentTime  string `json:"timestamp"`
}


func batteryHasPowerSupply() (bool, error) {
	batteries, err := battery.GetAll()
	if err != nil {
		return false, fmt.Errorf("error getting battery info: %v", err)
	}
	for i, battery := range batteries {
		bState := battery.State.String()
		fmt.Printf("Bat: %d has state: %s\n", i, bState)
		if strings.Contains(bState, "Full") {
			return true, nil
		}
	}
	return false, nil
}

func sendDataToServer(config Config, serverURL string) error {
	//struct to json
	jsonData, err := json.Marshal(config)
	if err != nil {
		fmt.Errorf("%v", err)
	}
	resp, err := http.Post(serverURL, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Errorf("%v", err)
	}
	defer resp.Body.Close()

	//checking response status
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected response status: %s", resp.Status)
	}
	return nil

}

func main() {
	hasPower, err := batteryHasPowerSupply()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(hasPower)

	//reading the config.json file
	configFile, err := ioutil.ReadFile("config.json")
	if err != nil {
		fmt.Println(err)
		return
	}

	//json from config
	var config Config
	err = json.Unmarshal(configFile, &config)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Location ID: %s Location Name: %s\n", config.LocationID, config.LocationName)

	serverURL := "http://localhost:8080/status"
	err = sendDataToServer(config, serverURL)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Config sent successfully!!")
	}

	// fmt.Println(config.LocationID)
	// fmt.Println(config.LocationName)

}

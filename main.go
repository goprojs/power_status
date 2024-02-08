package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/distatus/battery"
)

type Config struct {
	LocationID   string `json:"location_id"`
	LocationName string `json:"location_name"`
	ServerURL    string `json:"server_url"`
}

type indicator struct {
	ElectricityStatus bool   `json:"estatus"`
	LocationName      string `json:"location_name"`
	LocationID        string `json:"location_id"`
	CurrentTime       string `json:"timestamp"`
}

func batteryHasPowerSupply() (bool, error) {
	batteries, err := battery.GetAll()
	if err != nil {
		return false, fmt.Errorf("error getting battery info: %v", err)
	}
	for i, battery := range batteries {
		bState := battery.State.String()
		fmt.Printf("Bat: %d has state: %s\n", i, bState)
		if strings.Contains(bState, "Full") || strings.Contains(bState, "Charging") {
			return true, nil
		}
	}
	return false, nil
}

func sendDataToServer(data indicator, serverURL string) error {
	//struct to json
	jsonData, err := json.Marshal(data)
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
	currentState, _ := batteryHasPowerSupply()
	for {
		hasPower, err := batteryHasPowerSupply()
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(hasPower)

		//reading the config.json file
		configFile, err := os.ReadFile("config.json")
		if err != nil {
			fmt.Println(err)
			return
		}
		if hasPower != currentState {
			currentState = hasPower
			//json from config
			var config Config
			err = json.Unmarshal(configFile, &config)
			if err != nil {
				fmt.Println(err)
				return
				// time.Sleep(1 * time.Second)
				// continue
			}
			fmt.Printf("Location ID: %s Location Name: %s\n", config.LocationID, config.LocationName)

			currentTime := time.Now()
			timeString := currentTime.Format("2006-01-02 15:04")

			indicatorData := indicator{
				ElectricityStatus: hasPower,
				LocationName:      config.LocationName,
				LocationID:        config.LocationID,
				CurrentTime:       timeString,
			}

			serverURL := config.ServerURL
			// serverURL := "http://localhost:8080/status"
			err = sendDataToServer(indicatorData, serverURL)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println("indicator sent successfully!!")
			}
			time.Sleep(1 * time.Second)
		}
		// fmt.Println(config.LocationID)
		// fmt.Println(config.LocationName)
	}
}

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
	statusCode := resp.StatusCode
	fmt.Println(statusCode)
	if (statusCode != 200) && (statusCode != 201) {
		return fmt.Errorf("unexpected response status: %s", resp.Status)
	}
	return nil
}

func getAndSend() (bool, error) {
	// Read config file
	configFileName := "config.json"
	configFile, err := os.ReadFile(configFileName)
	if err != nil {
		fmt.Println(err)
		return false, err
	}
	var config Config
	err = json.Unmarshal(configFile, &config)
	if err != nil {
		fmt.Println(err)
		return false, err
	}


	currentTime := time.Now()
	timeString := currentTime.Format("2006-01-02 15:04:05")


	currentState, _ := batteryHasPowerSupply()
	indicatorData := indicator{
		ElectricityStatus: currentState,
		LocationName:      config.LocationName,
		LocationID:        config.LocationID,
		CurrentTime:       timeString,
	}

	serverURL := config.ServerURL
	err = sendDataToServer(indicatorData, serverURL)
	if err != nil {
		fmt.Println(err)
	}
	return currentState, nil
}

func main() {
	initialState, err := getAndSend()
	// Send initial status to server
	if err != nil {
		fmt.Println("Error sending initial data:", err)
	}
	// Print initial battery status
	fmt.Println("Initial battery status:", initialState)
	for {
		currentState, err := batteryHasPowerSupply()
		if err != nil {
			fmt.Println(err)
		}
		//fmt.Println(hasPower)
		if currentState != initialState {
			initialState = currentState
			_, err := getAndSend()
			if err != nil {
				fmt.Println("Error sending initial data:", err)
			}
			time.Sleep(1 * time.Second)
		}
	}
}

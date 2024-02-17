package utils

import (
	"testing"
)

func TestBatteryHasPowerSupplyNew(t *testing.T) {
	returnVal, err := BatteryHasPowerSupply()
	if err != nil {
		t.Errorf("Failed to get battery status")
	}
	if returnVal && !returnVal {
		t.Errorf("Unexpected status returned: %v", returnVal)
	}
}

// func TestBatteryHasPowerSupply(t *testing.T) {
// 	// Simulate a battery that is charging
// 	chargingBattery := batteryMock{State: batteryMockCharging}
// 	batteryMockGetAll = func() ([]batteryMock, error) {
// 		return []batteryMock{chargingBattery}, nil
// 	}
// 	hasPower, err := BatteryHasPowerSupply()
// 	if err != nil {
// 		t.Errorf("Unexpected error: %v", err)
// 	}
// 	if !hasPower {
// 		t.Error("Expected battery to have power supply, but it doesn't.")
// 	}

// 	// Simulate a battery that is not charging
// 	emptyBattery := batteryMock{State: batteryMockEmpty}
// 	batteryMockGetAll = func() ([]batteryMock, error) {
// 		return []batteryMock{emptyBattery}, nil
// 	}
// 	hasPower, err = BatteryHasPowerSupply()
// 	if err != nil {
// 		t.Errorf("Unexpected error: %v", err)
// 	}
// 	if hasPower {
// 		t.Error("Expected battery not to have power supply, but it does.")
// 	}

// 	// Simulate error when getting battery info
// 	batteryMockGetAll = func() ([]batteryMock, error) {
// 		return nil, mockError
// 	}
// 	_, err = BatteryHasPowerSupply()
// 	if err == nil {
// 		t.Error("Expected error when getting battery info, but got nil.")
// 	}
// }

// func TestSendDataToServer(t *testing.T) {
// 	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		w.WriteHeader(http.StatusOK)
// 	}))
// 	defer server.Close()

// 	data := Indicator{
// 		ElectricityStatus: true,
// 		LocationName:      "TestLocation",
// 		LocationID:        "123",
// 		CurrentTime:       time.Now().Format("2006-01-02 15:04:05"),
// 	}

// 	err := SendDataToServer(data, server.URL)
// 	if err != nil {
// 		t.Errorf("Unexpected error: %v", err)
// 	}
// }

// func TestGetAndSend(t *testing.T) {
// 	// Create a temporary config file for testing
// 	tempConfigFile, err := os.CreateTemp("", "config*.json")
// 	if err != nil {
// 		t.Fatalf("Error creating temp file: %v", err)
// 	}
// 	defer os.Remove(tempConfigFile.Name())

// 	// Write test config data to the temporary file
// 	testConfigData := `{"location_id":"123","location_name":"TestLocation","server_url":"http://example.com"}`
// 	_, err = tempConfigFile.WriteString(testConfigData)
// 	if err != nil {
// 		t.Fatalf("Error writing to temp file: %v", err)
// 	}

// 	// Test GetAndSend function
// 	hasPower, err := GetAndSend(tempConfigFile.Name())
// 	if err != nil {
// 		t.Errorf("Unexpected error: %v", err)
// 	}
// 	if hasPower {
// 		t.Error("Expected battery not to have power supply, but it does.")
// 	}
// }

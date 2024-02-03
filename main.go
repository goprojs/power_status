package main

import (
	"fmt"
	"strings"

	"github.com/distatus/battery"
)

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

func main() {
	hasPower, err := batteryHasPowerSupply()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(hasPower)
	//editing for push
}

package main

import (
	"fmt"
	"strconv"
	"strings"
)

// Data struct represents the data for each vehicle
type Data struct {
	Name string
	Plat string
	Type string
	Rute []string
}

// kenaRazia function checks for possible violations
func kenaRazia(date int, data []Data) []map[string]interface{} {
	// List of routes where ganjil-genap is enforced
	enforcedRoutes := []string{
		"Gajah Mada",
		"Hayam Wuruk",
		"Sisingamangaraja",
		"Panglima Polim",
		"Fatmawati",
		"Tomang Raya",
	}

	var violations []map[string]interface{}

	// Loop through each vehicle data
	for _, vehicle := range data {
		if vehicle.Type != "Mobil" {
			continue
		}

		// Get the last digit of the plate number
		plateParts := strings.Split(vehicle.Plat, " ")
		plateNumber := plateParts[1]
		lastDigit, _ := strconv.Atoi(string(plateNumber[len(plateNumber)-1]))

		// Determine if the date is odd or even
		dateOdd := date%2 != 0
		plateOdd := lastDigit%2 != 0

		if dateOdd != plateOdd {
			// Check if the vehicle is on an enforced route
			violationCount := 0
			for _, route := range vehicle.Rute {
				for _, enforcedRoute := range enforcedRoutes {
					if route == enforcedRoute {
						violationCount++
					}
				}
			}

			if violationCount > 0 {
				violation := map[string]interface{}{
					"name":   vehicle.Name,
					"tilang": violationCount,
				}
				violations = append(violations, violation)
			}
		}
	}

	return violations
}

func main() {
	data := []Data{
		{
			Name: "Denver",
			Plat: "B 2791 KDS",
			Type: "Mobil",
			Rute: []string{"TB Simatupang", "Panglima Polim", "Depok", "Senen Raya"},
		},
		{
			Name: "Toni",
			Plat: "B 1212 JBB",
			Type: "Mobil",
			Rute: []string{"Pintu Besar Selatan", "Panglima Polim", "Depok", "Senen Raya", "Kemang"},
		},
		{
			Name: "Stark",
			Plat: "B 444 XSX",
			Type: "Motor",
			Rute: []string{"Pondok Indah", "Depok", "Senen Raya", "Kemang"},
		},
		{
			Name: "Anna",
			Plat: "B 678 DD",
			Type: "Mobil",
			Rute: []string{"Fatmawati", "Panglima Polim", "Depok", "Senen Raya", "Kemang", "Gajah Mada"},
		},
	}

	result := kenaRazia(27, data)
	fmt.Println(result) // Expected output: [{name: "Toni", tilang: 1}, {name: "Anna", tilang: 3}]
}

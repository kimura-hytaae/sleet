package main

import (
	_ "embed"
	"encoding/json"
	"fmt"
	"strconv"
)

//go:embed "cities.json"
var citiesJson []byte

type City struct {
	Country   string  `json:"country"`
	Name      string  `json:"name"`
	Lat       string  `json:"lat"`
	Lng       string  `json:"lng"`
	Latitude  float64 `json:"-"`
	Longitude float64 `json:"-"`
}

func goMain2(args []string) int {
	cities := []City{}
	err := json.Unmarshal(citiesJson, &cities)
	if err != nil {
		fmt.Println(err.Error())
		return 1
	}
	for _, city := range cities {
		city.Latitude, _ = strconv.ParseFloat(city.Lat, 64)
		city.Longitude, _ = strconv.ParseFloat(city.Lng, 64)
	}
	fmt.Printf("read %d entries\n", len(cities))
	/*for i := 0; i < len(cities); i++ {
		if args[1] == cities[i].Name {
			fmt.Println(cities[i])
			return 0
		}
	}*/
	return 2
}

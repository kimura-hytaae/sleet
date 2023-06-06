package sleet

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

func GoMain2(args []string) int {
	cities := []*City{}
	err := json.Unmarshal(citiesJson, &cities)
	if err != nil {
		fmt.Println(err.Error())
		return 1
	}
	for _, city := range cities {
		city.Latitude, _ = strconv.ParseFloat(city.Lat, 64)
		city.Longitude, _ = strconv.ParseFloat(city.Lng, 64)
		// fmt.Printf("%s %s (%s, %s), (%f, %f)\n", city.Country, city.Name, city.Lat, city.Lng, city.Latitude, city.Longitude) // 「cities := []*City{}」をポインタにすることで解消, ポインタにしない場合 一つ一つの配列に書く必要性あり
	}
	fmt.Printf("read %d entries\n", len(cities))  // いずれ消す
	/*
	for i := 0; i < len(cities); i++ {
		if args[1] == cities[i].Name {
			fmt.Printf("%sの緯度：%f, 経度：%f\n", cities[i].Name, cities[i].Latitude, cities[i].Longitude)
			return 0
		}
	} */
	for _, citiy := range cities {
		if args[1] == citiy.Name {
			fmt.Printf("%sの緯度：%f, 経度：%f\n", citiy.Name, citiy.Latitude, citiy.Longitude)
		}
	}
	return 2
}

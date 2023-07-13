package sleet

import (
	_ "embed"
	"encoding/json"
	"fmt"
	//"strconv"
)

//go:embed "weather.json"
var weatherJson []byte

type Weather struct {
	Code int    `json:"code"`
	Jp   string `json:"jp"`
	En   string `json:"en"`
}

// func Weathers(weather_code string) {
func Weathers(url_weathers *Result) {
	weathers := []*Weather{}
	err := json.Unmarshal(weatherJson, &weathers)
	if err != nil {
		//return nil, err
		fmt.Println("error")
	}
	fmt.Println("日付：天気")
	for index, url_weather := range url_weathers.Day.Whetercode {
		weather_name := weathers[url_weather]
		weather_days := url_weathers.Day.Time[index]
		weather_month := weather_days[5:7]
		weather_day := weather_days[8:10]
		fmt.Printf("%s月%s日：%s\n", weather_month, weather_day, weather_name.Jp) //wehathers.Code[url_weather]
	}
}

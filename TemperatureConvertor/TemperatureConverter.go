package main

import "fmt"

func main() {
	var celciusVal float64 = 41
	fahrConv := tempConvToCelcius(float64(celciusVal))
	fmt.Println("Fahrenheit Value: " + fmt.Sprint(celciusVal))
	fmt.Println("Celcius Value: " + fmt.Sprint(fahrConv))

	var fahrVal float64 = 41
	celciusConv := tempConvToFahrenheit(float64(fahrVal))
	fmt.Println("Fahrenheit Value: " + fmt.Sprint(fahrVal))
	fmt.Println("Celcius Value: " + fmt.Sprint(celciusConv))

}

func tempConvToCelcius(fahrenheit float64) float64 {
	var celcius float64 = (fahrenheit - 32) * (float64(5) / float64(9))
	return celcius
}

func tempConvToFahrenheit(celcius float64) float64 {
	fahrenheit := (celcius * 9 / 5) + 32
	return fahrenheit
}

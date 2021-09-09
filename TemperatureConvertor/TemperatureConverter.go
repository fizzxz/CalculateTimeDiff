package main

import "fmt"

func main() {
	var celciusVal float64 = 0
	fahrConvFromCelc := tempConvFromCelciusToFahrenheit(celciusVal)
	kelvinConvFromCelc := tempConvFromCelciusToKelvin(celciusVal)
	fmt.Println("Original Celcius Value: " + fmt.Sprint(celciusVal))
	fmt.Println("Fahrenheit Value: " + fmt.Sprint(fahrConvFromCelc))
	fmt.Println("Kelvin Value: " + fmt.Sprint(kelvinConvFromCelc))
	fmt.Println(" ")
	var fahrVal float64 = 0
	celciusConvFromFahr := tempConvFromFahrenheitToCelcius(fahrVal)
	kelvinConvFromFahr := tempConvFromFahrenheitToKelvin(fahrVal)
	fmt.Println("Original Fahrenheit Value: " + fmt.Sprint(fahrVal))
	fmt.Println("Celcius Value: " + fmt.Sprint(celciusConvFromFahr))
	fmt.Println("Kelvin Value: " + fmt.Sprint(kelvinConvFromFahr))
	fmt.Println(" ")
	var kelvinVal float64 = 0
	fahrConvFromKelv := tempConvFromKelvinToFahrenheit(kelvinVal)
	celcConvFromKelv := tempConvFromKelvinToCelcius(kelvinVal)
	fmt.Println("Original Kelvin Value: " + fmt.Sprint(kelvinVal))
	fmt.Println("Fahrenheit Value: " + fmt.Sprint(fahrConvFromKelv))
	fmt.Println("Celcius Value: " + fmt.Sprint(celcConvFromKelv))

}

func tempConvFromFahrenheitToCelcius(fahrenheit float64) float64 {
	var celcius float64 = (fahrenheit - 32) * (float64(5) / float64(9))
	return celcius
}
func tempConvFromKelvinToCelcius(kelvin float64) float64 {
	var celcius float64 = kelvin - 273.15
	return celcius
}

func tempConvFromCelciusToFahrenheit(celcius float64) float64 {
	fahrenheit := (celcius * 9 / 5) + 32
	return fahrenheit
}

func tempConvFromKelvinToFahrenheit(kelvin float64) float64 {
	fahrenheit := tempConvFromCelciusToFahrenheit(
		tempConvFromKelvinToCelcius(kelvin))
	return fahrenheit
}

func tempConvFromCelciusToKelvin(celcius float64) float64 {
	var kelvin float64 = celcius + 273.15
	return kelvin
}

func tempConvFromFahrenheitToKelvin(fahrenheit float64) float64 {
	kelvin := tempConvFromCelciusToKelvin(
		tempConvFromFahrenheitToCelcius(fahrenheit))
	return kelvin
}

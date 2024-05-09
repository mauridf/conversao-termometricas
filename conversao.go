package main

import "fmt"

// Função para converter de Celsius para Fahrenheit
func celsiusParaFahrenheit(celsius float64) float64 {
	return (celsius * 9 / 5) + 32
}

// Função para converter de Celsius para Kelvin
func celsiusParaKelvin(celsius float64) float64 {
	return celsius + 273.15
}

func main() {
	// Definindo uma temperatura em Celsius
	temperaturaCelsius := 100.0

	// Convertendo de Celsius para Fahrenheit
	temperaturaFahrenheit := celsiusParaFahrenheit(temperaturaCelsius)
	fmt.Printf("%.2f°C é equivalente a %.2f°F\n", temperaturaCelsius, temperaturaFahrenheit)

	// Convertendo de Celsius para Kelvin
	temperaturaKelvin := celsiusParaKelvin(temperaturaCelsius)
	fmt.Printf("%.2f°C é equivalente a %.2fK\n", temperaturaCelsius, temperaturaKelvin)
}

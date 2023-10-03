package humidity

const (
	d = 273.15 // Conversion from Celsius to Kelvin
)

func CelsiusToFahrenheit(celsius float64) (fahrenheit float64) {
	fahrenheit = (celsius * 9 / 5) + 32
	return fahrenheit
}

func FahrenheitToCelsius(fahrenheit float64) (celsius float64) {
	celsius = (fahrenheit - 32) * 5 / 9
	return celsius
}

func CelsiusToKelvin(celsius float64) (kelvin float64) {
	kelvin = celsius + d
	return kelvin
}

func KelvinToCelsiu(kelvin float64) (celsius float64) {
	celsius = kelvin - d
	return celsius
}

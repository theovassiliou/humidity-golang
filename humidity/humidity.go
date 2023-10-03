package humidity

import (
	"math"
)

// Sources: 	https://www.schweizer-fn.de/lueftung/feuchte/feuchte.php
// 				https://bauweise.net/grundlagen/feuchte/feucht.htm
// 				https://de.wikipedia.org/wiki/Taupunkt#Abh%C3%A4ngigkeit_der_Taupunkttemperatur_von_relativer_Luftfeuchtigkeit_und_Lufttemperatur

// RelativeToAbsolute converts relative humidity in % at a given tempreture (in °C) to absolute humidity in g/m^3.
func RelativeToAbsolute(rh float64, tempC float64) float64 {
	// Convert temperature to Kelvin
	tempK := tempC + d

	// Convert relative humidity to absolute humidity
	ah := rh * MagnusEquation(tempC) / (gasConst * tempK)

	return ah * 10
}

// AbsoluteToRelative converts absolute humidity (in g/m^3) at a given temperature (in °C) to relative humidity in %.
func AbsoluteToRelative(ah float64, tempC float64) float64 {
	// Convert temperature to Kelvin
	tempK := tempC + d

	// Convert absolute humidity to relative humidity
	rh := (ah / 10 * gasConst * tempK) / MagnusEquation(tempC)

	return rh
}

// Humidity represents the humidity of air for a given temperature
// Humidity can be expressed in a variety of ways, mainly absolute humidity
// defined by the amount of water in a volume in the air at a given temperature.
// Alternatively humidity can be expressed as relative humidity in %, which is the
// amount of water in relation to the maximum amount of water that air can
// contain for a given temperature.
type Humidity struct {
	tempC float64
	absH  float64
}

// RelativeHumidity returns the relative humidity in %
func (h Humidity) RelativeHumidity() (rh float64) {
	return AbsoluteToRelative(h.absH, h.tempC)
}

// AbsoluteHumidity returns the absolut humidity in g/m^3
func (h Humidity) AbsoluteHumidity() (ah float64) {
	return h.absH
}

// Fahrenheit returns the temperatur in Fahrenheit
func (h Humidity) Fahrenheit() (t float64) {
	return CelsiusToFahrenheit(h.tempC)
}

// Celsius returns the temperatur in Celsius
func (h Humidity) Celsius() (t float64) {
	return h.tempC
}

// DewPointTemperatur returns the dew point temperatur in °C
// The dew point temperature is the temperature at which the
// water vapour saturation concentration or the water vapour saturation pressure of the air is reached.
// The relative humidity in this state is 100%.
// If the humid air is cooled below the dew point temperature, a phase change from
// gaseous to liquid occurs and part of the water vapour
// contained in the air is excreted as excess moisture in liquid form as condensation.
func (h Humidity) DewPointTemperatur() (dewPointTemp float64) {
	dewPointTemp = DewPointTemperature(h.RelativeHumidity(), h.tempC)
	return dewPointTemp
}

func NewRelativeHumidityCelsius(relH, tempC float64) (h Humidity) {
	return Humidity{
		tempC: tempC,
		absH:  RelativeToAbsolute(relH, tempC),
	}
}

func NewRelativeHumidityFahrenheit(relH, tempF float64) (h Humidity) {
	tempC := FahrenheitToCelsius(tempF)
	return Humidity{
		tempC: tempC,
		absH:  RelativeToAbsolute(relH, tempC),
	}
}

func NewHumidity(ah, tempC float64) (h Humidity) {
	return Humidity{
		tempC: tempC,
		absH:  ah,
	}
}

// Constants for the Magnus equation
const (
	K1       = 611.2
	K2       = 17.62
	K3       = 243.12
	gasConst = 461.51
)

// Water vapour partial pressure at a given temperatur in °C according the Magnus Equation
func MagnusEquation(tempC float64) float64 {
	if tempC < -45.0 || tempC > 60.0 {
		panic("Magnugs equation only valid between -45°C and 60°C")
	}

	return K1 * math.Exp((K2*tempC)/(K3+tempC))
}

// DewPointTemperature return for a given relative humidity (in %)
// and temperature the dew point temperatur in °C
func DewPointTemperature(rH float64, tempC float64) (dpt float64) {
	if tempC < -45.0 || tempC > 60.0 {
		panic("Dew Point calculation via Magnus equation only valid between -45°C and 60°C")
	}
	rH = rH / 100.0
	dpt = K3 *
		((K2*tempC)/(K3+tempC) + math.Log(rH)) /
		((K2*K3)/(K3+tempC) - math.Log(rH))
	return dpt
}

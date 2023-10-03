package main

import (
	"fmt"

	"github.com/theovassiliou/humidity-golang/humidity"
)

// Example usage
func main() {

	// Room temperature in °C
	tempIn := 22.7

	// relative humidity in the room in %
	rHumIn := 66.0

	// Outside temp and relative humidity
	tempOut, rHumOut := 18.0, 74.0

	if CanDehumidfyInterior(rHumIn, tempIn, rHumOut, tempOut) {
		fmt.Println("When opening the windows humidity in the room will decrease")
	} else {
		fmt.Println("When opening the windows humidity in the room will increase")
	}

}

// CanDehumidfyInterior returns true if the absolute humidity outside is sufficiently enough lower
// so that by opening the windows you could dehumidify the interiour
// Background: Only if air, that comes from the outside has a substantially lower absolute humidity
// it can, when getting warm inside absorb additional water and thus demhumidfy a room.
// If the difference if not siginifcant enought, or if the absolute outside humidity is
// higher it wont dehumidify the inside of a room, or even transport humidity inside.
func CanDehumidfyInterior(rHin, tempIn, rHout, tempOut float64, minDifInPercent ...float64) bool {

	ahIn := humidity.RelativeToAbsolute(rHin, tempIn)
	ahOut := humidity.RelativeToAbsolute(rHout, tempOut)

	margin := 0.8
	if len(minDifInPercent) > 0 {
		diff := minDifInPercent[0]
		if diff < 1 {
			margin = 1 - diff

		}
	}
	return (ahOut < ahIn*margin)

}

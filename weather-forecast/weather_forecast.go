// Package weather provides a function to forecast.
package weather

// CurrentCondition is current condition.
var CurrentCondition string

// CurrentLocation is current location.
var CurrentLocation string

// Forecast returns a string which tells the weather at the given location and condition.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}

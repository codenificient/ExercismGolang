// Package weather provides tools to provide current weather conditions in a particular city.
package weather

// CurrentCondition stores the weather conditions of the city.
var CurrentCondition string

// CurrentLocation stores the name of the city.
var CurrentLocation string

// Forecast function allows us to provide the weather conditions at a given location or city.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}

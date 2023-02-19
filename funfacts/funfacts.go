package funfacts

import (
	"fmt"

	conv "github.com/jonaslefdal/funtemps/conv"
)

type FunFacts struct {
	Sun   []string
	Luna  []string
	Terra []string
}

func GetFunFacts(about string, value int, unit string) []string {
	var convFunc func(float64) float64
	switch unit {
	case "C":
		convFunc = func(t float64) float64 { return t }
	case "F":
		convFunc = conv.CelsiusToFahrenheit
	case "K":
		convFunc = conv.CelsiusToKelvin
	default:
		return []string{"Invalid unit. Please enter 'C', 'F', or 'K'."}
	}

	facts := FunFacts{
		Sun: []string{
			fmt.Sprintf("Temperatur i Solens kjerne: %.2f°%s", convFunc(15000000), unit),
			fmt.Sprintf("Temperatur på ytre lag av Solen: %.2f°%s", convFunc(5504.85), unit),
		},
		Luna: []string{
			fmt.Sprintf("Temperatur på Månens overflate om natten: %.2f°%s", convFunc(-183), unit),
			fmt.Sprintf("Temperatur på Månens overflate om dagen: %.2f°%s", convFunc(106), unit),
		},
		Terra: []string{
			fmt.Sprintf("Høyeste temperatur målt på Jordens overflate: %.2f°%s", convFunc(56.7), unit),
			fmt.Sprintf("Laveste temperatur målt på Jordens overflate: %.2f°%s", convFunc(-89.4), unit),
			fmt.Sprintf("Temperatur i Jordens indre kjerne: %.2f°%s", convFunc(9118.85), unit),
		},
	}

	switch about {
	case "sun":
		return facts.Sun
	case "luna":
		return facts.Luna
	case "terra":
		return facts.Terra
	default:
		return []string{"Invalid input. Please enter 'sun', 'luna', or 'terra'."}
	}
}

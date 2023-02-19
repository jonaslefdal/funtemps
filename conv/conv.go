package conv

/*
// Måtte importe math for funksjonen
import (
  "math"
)*/

/* //https://gosamples.dev/round-float/
// Denne funksjonen ble brukt under testing for at tallene skulle rundes opp til max 2 desimaler
   og ikke feile. Feilmeldingen ble
   --- FAIL: TestCelsiusToKelvin (0.00s)
    conv_test.go:81: expected: 292.47, got: 292.46999999999997
// Etter testing ble koden endret grunnet oppgavene tilsa formatering av float64 ikke skulle gjøres her, men i main.go

func roundFloat(val float64, precision uint) float64 {
  ratio := math.Pow(10, float64(precision))
  return math.Round(val*ratio) / ratio
}
*/

// De kommenterte return kodene ble brukt under testing for å ha max 2 desimaler

//OBS! Koden skal fungere som normalt uten det nevnt over da jeg nå bruker "withintolerance i test filen"

func FahrenheitToCelsius(value float64) float64 {
	//return roundFloat(((value - 32) * 5/9), 2)
	return (value - 32) * 5 / 9
}
func FahrenheitToKelvin(value float64) float64 {
	//return roundFloat(((value - 32) * 5/9 + 273.15), 2)
	return (value-32)*5/9 + 273.15
}
func CelsiusToFahrenheit(value float64) float64 {
	//return roundFloat((value*(9/5) + 32), 2)
	return (value*9/5 + 32)
}
func CelsiusToKelvin(value float64) float64 {
	//return roundFloat(value + 273.15, 2)
	return (value + 273.15)
}
func KelvinToFarhenheit(value float64) float64 {
	//return roundFloat((value - 273.15)*(9/5) + 32, 2)
	return (value-273.15)*9/5 + 32
}
func KelvinToCelsius(value float64) float64 {
	//return roundFloat(value - 273.15, 2)
	return (value - 273.15)
}

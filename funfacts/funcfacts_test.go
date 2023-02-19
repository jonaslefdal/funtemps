package funfacts

import (
	"reflect"
	"testing"
)

/*
*

	Mal for TestGetFunFacts funksjonen.
	Definer korrekte typer for input og want,
	og sette korrekte testverdier i slice tests.
*/
func TestGetFunFacts(t *testing.T) {
	type test struct {
		funfact string
		value   int
		unit    string
		want    []string
	}

	tests := []test{
		{
			funfact: "sun",
			value:   0,
			unit:    "C",
			want: []string{"Temperatur i Solens kjerne: 15000000.00°C",
				"Temperatur på ytre lag av Solen: 5504.85°C",
			},
		},
		{
			funfact: "luna",
			value:   0,
			unit:    "C",
			want: []string{"Temperatur på Månens overflate om natten: -183.00°C",
				"Temperatur på Månens overflate om dagen: 106.00°C",
			},
		},
		{
			funfact: "terra",
			value:   0,
			unit:    "C",
			want: []string{"Høyeste temperatur målt på Jordens overflate: 56.70°C",
				"Laveste temperatur målt på Jordens overflate: -89.40°C",
				"Temperatur i Jordens indre kjerne: 9118.85°C"},
		},
	}

	for _, tc := range tests {
		got := GetFunFacts(tc.funfact, tc.value, tc.unit)
		if !reflect.DeepEqual(tc.want, got) {
			t.Errorf("expected: %v, got: %v", tc.want, got)
		}
	}
}

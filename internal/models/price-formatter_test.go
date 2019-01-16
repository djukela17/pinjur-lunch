package models

import (
	"testing"
)

func TestParsePrice(t *testing.T) {
	cases := []struct {
		in   string
		want int
	}{
		{"10,00 kn", 1000},
		{"1,00 kn", 100},
		{"0,50 kn", 50},
		{"0,00 kn", 0},
		{"0,0 kn", 0},

		{"od 52,00 kn", 5200},
	}

	for _, c := range cases {
		got, err := ParsePrice(c.in)

		if err != nil || got != c.want {
			t.Errorf("ParsePrice(%v) == %v, expected: %v\n", c.in, got, c.want)
		}
	}
}

func TestFormatPrice(t *testing.T) {
	cases := []struct {
		inAmount   int
		inCurrency string
		want       string
	}{
		{inAmount: 0, inCurrency: "Kn", want: "0,00 Kn"},
		{inAmount: 100, inCurrency: "Kn", want: "1,00 Kn"},
		{inAmount: 50, inCurrency: "Kn", want: "0,50 Kn"},
		{inAmount: 5, inCurrency: "Kn", want: "0,05 Kn"},
		{inAmount: 5200, inCurrency: "Kn", want: "52,00 Kn"},
	}

	for _, c := range cases {
		got := FormatPrice(c.inAmount, c.inCurrency)

		if got != c.want {
			t.Errorf("FormatPrice(%v, %v) == %v, expected: %v \n", c.inAmount, c.inCurrency, got, c.want)
		}
	}
}

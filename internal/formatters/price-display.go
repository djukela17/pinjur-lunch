package formatters

import "strconv"

func DisplayPrice(price int) string {
	currency := "Kn"
	whole := price / 100
	fraction := price % 100

	fractionTxt := strconv.Itoa(fraction)
	if fraction < 10 && fraction > -10 {
		fractionTxt = "0" + fractionTxt
	}
	return strconv.Itoa(whole) + "," + fractionTxt + " " + currency
}

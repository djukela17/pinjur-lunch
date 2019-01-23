package models

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

type Dish struct {
	Name       string `json:"name"`
	Type       string `json:"type"`
	Price      int    `json:"price"`
	DefaultOn  bool   `json:"default"`
	SideDishes MealAdditionsCollection
	Extras     MealAdditionsCollection
}

func (d *Dish) DisplayPrice() string {
	currency := "Kn"
	whole := d.Price / 100
	fraction := d.Price % 100

	fractionTxt := strconv.Itoa(fraction)
	if fraction < 10 && fraction > -10 {
		fractionTxt = "0" + fractionTxt
	}
	return strconv.Itoa(whole) + "," + fractionTxt + " " + currency
}

func (d *Dish) ParsePrice(price string) (int, error) {

	priceParts := strings.Split(price, ",")

	if len(priceParts) > 1 {
		wholePart := strings.Split(priceParts[0], " ")

		idx := len(wholePart) - 1

		whole, err := strconv.Atoi(wholePart[idx])

		if err != nil {
			return 0, err
		}
		fractionPart := strings.Split(priceParts[1], " ")

		fraction, err := strconv.Atoi(fractionPart[0])
		if err != nil {
			return 0, err
		}

		return whole*100 + fraction, nil
	}

	return 0, fmt.Errorf("the price could not be divided by the provided separator")
}

func LoadDishList(path string) ([]Dish, error) {
	jsonDishData, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	byteData, _ := ioutil.ReadAll(jsonDishData)
	defer jsonDishData.Close()

	var dishes []Dish
	if err := json.Unmarshal(byteData, &dishes); err != nil {
		return nil, err
	}

	return dishes, nil
}

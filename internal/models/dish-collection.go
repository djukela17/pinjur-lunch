package models

import "fmt"

type DishCollection struct {
	Dishes []Dish
}

func (c *DishCollection) GetDish(name string) (Dish, error) {

	for _, d := range c.Dishes {
		if d.Name == name {
			return d, nil
		}
	}

	return Dish{}, fmt.Errorf("no dish found")
}

func (c *DishCollection) GetAll() []Dish {
	return c.Dishes
}

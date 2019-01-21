package models

import "fmt"

type DishCollection struct {
	Dishes []Dish
}

func NewDishCollection(dishes []Dish) DishCollection {
	return DishCollection{Dishes: dishes}
}

func (c *DishCollection) AddDish(dish Dish) {
	c.Dishes = append(c.Dishes, dish)
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

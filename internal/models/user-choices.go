package models

import (
	"fmt"
	"strconv"
)

type UserChoice struct {
	Username     string
	ChosenDish   Dish
	OptionalNote string
}

type UserChoices struct {
	Choices []UserChoice
}

func (u *UserChoices) GetUserChoices() []UserChoice {
	return u.Choices
}

func (u *UserChoices) AddDish(dish Dish, name, optionalNote string) error {
	choice := UserChoice{Username: name, ChosenDish: dish, OptionalNote: optionalNote}
	u.Choices = append(u.Choices, choice)
	fmt.Println(u.Choices)
	return nil
}

func (u *UserChoices) CreateCompressedList() []string {
	var ret []string
	stacked := make(map[string]int)

	for _, choice := range u.Choices {
		fullChoice := choice.ChosenDish.Name
		if choice.OptionalNote != "" {
			fullChoice += "(" + choice.OptionalNote + ")"
		}
		stacked[fullChoice] += 1
	}

	for k, v := range stacked {
		ret = append(ret, strconv.Itoa(v)+"x "+k)
	}
	fmt.Println(stacked)

	return ret
}

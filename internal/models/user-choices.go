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

func (u *UserChoices) AddDish(dishList []Dish, dishName, username, optionalNote string) error {

	// find the corresponding dish from the list
	for _, dish := range dishList {
		if dish.Name == dishName {
			chosenDish := dish
			choice := UserChoice{Username: username, ChosenDish: chosenDish, OptionalNote: optionalNote}
			u.Choices = append(u.Choices, choice)
			fmt.Println(u.Choices)
			return nil
		}
	}

	return fmt.Errorf("could not find the provided dish in the available dish list")
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

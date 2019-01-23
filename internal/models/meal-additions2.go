package models

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

type MealAdditions2 struct {
	Type  string   `json:"type"`
	Items []string `json:"list"`
}

type MealAdditions2Collection struct {
	Additions []MealAdditions2
}

func NewMealAdditions2Collection(adds []MealAdditions2) MealAdditions2Collection {
	return MealAdditions2Collection{
		Additions: adds,
	}
}

func LoadFromFile(path string) ([]MealAdditions2, error) {
	jsonDishData, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	byteData, _ := ioutil.ReadAll(jsonDishData)
	defer jsonDishData.Close()

	var adds []MealAdditions2
	if err := json.Unmarshal(byteData, &adds); err != nil {
		return nil, err
	}

	return adds, nil
}

func (m *MealAdditions2Collection) GetAll() []MealAdditions2 {
	return m.Additions
}

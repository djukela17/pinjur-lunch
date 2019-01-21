package models

import (
	"fmt"
	"testing"
)

func TestDishCollection_GetDish(t *testing.T) {
	collection1 := DishCollection{
		Dishes: []Dish{
			{Name: "Piletina",
				Price: 3000,
			},
			{
				Name:  "Hamburger",
				Price: 2400,
			},
		},
	}
	cases := []struct {
		in      string
		wantRes Dish
		wantErr error
	}{
		{in: "Piletina", wantRes: Dish{Name: "Piletina", Price: 3000}, wantErr: nil},
		{in: "Hamburger", wantRes: Dish{Name: "Piletina", Price: 2400}, wantErr: nil},

		// should fail
		{in: "Cheesburger", wantRes: Dish{}, wantErr: fmt.Errorf("no dish found")},
	}

	for _, c := range cases {
		got, err := collection1.GetDish(c.in)

		if got.Name != c.wantRes.Name && err != c.wantErr {
			t.Errorf("DishCollection_GetDish(%v) == %v, %v. Expected: %v, %v", c.in, got, err, c.wantRes, c.wantErr)
		}

	}
}

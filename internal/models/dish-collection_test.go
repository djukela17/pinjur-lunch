package models

import "testing"

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
		{in: "Cheesburger", wantRes: Dish{Name: "Cheesburger", Price: 2500}, wantErr: nil},
	}

	for _, c := range cases {
		got, err := collection1.GetDish(c.in)

		if got.Name != c.wantRes.Name && err != c.wantErr {
			t.Errorf("DishCollection_GetDish(%v) == %v, %v. Expected: %v, %v", c.in, got, err, c.wantRes, c.wantErr)
		}

	}
}

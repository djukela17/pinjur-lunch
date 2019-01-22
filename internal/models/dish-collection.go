package models

import (
	"context"
	"fmt"
	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/mongodb/mongo-go-driver/mongo"
	"github.com/mongodb/mongo-go-driver/mongo/readpref"
	"time"
)

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

func (c *DishCollection) InsertAll(client *mongo.Client, databaseName, collectionName string, deleteOld bool) error {
	// ping to check if server is available
	// this should not be necessary every time
	ctx, _ := context.WithTimeout(context.Background(), 2*time.Second)
	err := client.Ping(ctx, readpref.Primary())
	if err != nil {
		return err
	}

	interfaceSlice := make([]interface{}, len(c.Dishes))
	for i, d := range c.Dishes {
		interfaceSlice[i] = d
	}

	collection := client.Database(databaseName).Collection(collectionName)
	ctx, _ = context.WithTimeout(context.Background(), 5*time.Second)

	if deleteOld {
		res, err := collection.DeleteMany(ctx, bson.M{})
		if err != nil {
			return err
		}
		fmt.Println(res)
	}

	ctx, _ = context.WithTimeout(context.Background(), 5*time.Second)
	res, err := collection.InsertMany(ctx, interfaceSlice)

	fmt.Println(res.InsertedIDs)
	fmt.Println(len(res.InsertedIDs))

	return nil
}

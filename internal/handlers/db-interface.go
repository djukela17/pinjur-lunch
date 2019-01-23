package handlers

import (
	"context"
	"github.com/djukela17/pinjur-lunch/internal/models"
	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/mongodb/mongo-go-driver/mongo/readpref"
	"time"
)

func (h *MainHandler) GetAllDishes() ([]models.Dish, error) {
	var dishes []models.Dish

	// ping to check if server is available
	// this should not be necessary every time
	ctx, _ := context.WithTimeout(context.Background(), 2*time.Second)
	if err := h.MongoClient.Ping(ctx, readpref.Primary()); err != nil {
		return nil, err
	}

	collection := h.MongoClient.Database(h.DatabaseName).Collection(h.AllDishCollectionName)
	ctx, _ = context.WithTimeout(context.Background(), 10*time.Second)

	filter := bson.M{}

	cur, err := collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	for cur.Next(ctx) {
		dish := models.Dish{}

		if err := cur.Decode(&dish); err != nil {
			// ? skip errors while decoding ?
			continue
		}
		dishes = append(dishes, dish)
	}
	if err := cur.Err(); err != nil {
		return nil, err
	}
	return dishes, nil
}

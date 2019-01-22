package models

import (
	"context"
	"fmt"
	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/mongodb/mongo-go-driver/mongo"
	"github.com/mongodb/mongo-go-driver/mongo/readpref"
	"time"
)

type MealAdditions struct {
	Name string `json:"name"`
	Tag  string `json:"tag"`
}

type MealAdditionsCollection struct {
	// Mongo db constants
	DatabaseName   string
	CollectionName string

	SideDishes []MealAdditions
}

func NewAdditionsCollection(databaseName, collectionName string) MealAdditionsCollection {
	//jsonData, err := os.Open(path)
	//if err != nil {
	//	return MealAdditionsCollection{}, err
	//}
	//
	//byteData, _ := ioutil.ReadAll(jsonData)
	//defer jsonData.Close()
	//
	//var sideDishes []MealAdditions
	//if err := json.Unmarshal(byteData, &sideDishes); err != nil {
	//	return MealAdditionsCollection{}, err
	//}

	adds := MealAdditionsCollection{
		DatabaseName:   databaseName,
		CollectionName: collectionName,

		//SideDishes: sideDishes,
	}
	return adds
}

func (m *MealAdditionsCollection) LoadAll(client *mongo.Client) error {
	var additions []MealAdditions
	// ping to check if server is available
	// this should not be necessary every time
	ctx, _ := context.WithTimeout(context.Background(), 2*time.Second)
	err := client.Ping(ctx, readpref.Primary())
	if err != nil {
		return err
	}

	collection := client.Database(m.DatabaseName).Collection(m.CollectionName)
	ctx, _ = context.WithTimeout(context.Background(), 10*time.Second)

	filter := bson.M{}

	cur, err := collection.Find(ctx, filter)
	if err != nil {
		return err
	}
	for cur.Next(ctx) {
		addition := MealAdditions{}

		if err := cur.Decode(&addition); err != nil {
			// ? skip errors while decoding ?
			continue
		}
		additions = append(additions, addition)
		fmt.Println(addition)
	}
	if err := cur.Err(); err != nil {
		return err
	}

	return nil
}

func (m *MealAdditionsCollection) InsertAll(client *mongo.Client, dbName, collName string) error {
	// ping to check if server is available
	// this should not be necessary every time
	ctx, _ := context.WithTimeout(context.Background(), 2*time.Second)
	err := client.Ping(ctx, readpref.Primary())
	if err != nil {
		return err
	}

	interfaceSlice := make([]interface{}, len(m.SideDishes))
	for i, d := range m.SideDishes {
		interfaceSlice[i] = d
	}

	collection := client.Database(dbName).Collection(collName)
	ctx, _ = context.WithTimeout(context.Background(), 5*time.Second)
	res, err := collection.InsertMany(ctx, interfaceSlice)

	fmt.Println(res.InsertedIDs)
	fmt.Println(len(res.InsertedIDs))

	return nil
}

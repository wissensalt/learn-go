package main

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Person struct {
	Name string
	Age  int
	City string
}

func main() {
	fmt.Println("Running App")
	clientOptions := options.Client().ApplyURI("mongodb://local.mokapos.com:27017")
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")

	ruan := Person{"Ruan", 34, "Cape Town"}

	collection := client.Database("db-one").Collection("person")
	insertResult, err := collection.InsertOne(context.TODO(), ruan)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted a Single Document: ", insertResult.InsertedID)


	james := Person{"James", 32, "Nairobi"}
	frankie := Person{"Frankie", 31, "Nairobi"}

	trainers := []interface{}{james, frankie}

	insertManyResult, err := collection.InsertMany(context.TODO(), trainers)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted multiple documents: ", insertManyResult.InsertedIDs)


	filter := bson.D
	update := bson.D {
		{"$inc", bson.D{
			{"age", 1},
		}},
	}

	updateResult, err := collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Matched %v documents and updated %v documents.\n", updateResult.MatchedCount, updateResult.ModifiedCount)
}

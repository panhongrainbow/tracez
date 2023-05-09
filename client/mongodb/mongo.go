package main

import (
	"context"
	"fmt"
	"github.com/panhongrainbow/tracez/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

func main() {
	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		panic(err)
	}
	defer func() {
		_ = client.Disconnect(context.TODO())
	}()

	// Select database and collection
	db := client.Database("openTelemetry2mongodb")
	ctx := context.TODO()
	collection := db.Collection("openTelemetry2mongodb")

	// Query for documents in which the instrumentationlibrary.name field equals mainTracer
	now := time.Now()
	past := now.Add(-30 * 24 * time.Hour)
	pipeline := mongo.Pipeline{
		bson.D{
			{"$match", bson.D{{"instrumentationlibrary.name", "mainTracer"},
				{"starttime", bson.D{{"$gte", past}}},
			}}},
		{{"$sort", bson.D{{"starttime", 1}}}},
		{{"$limit", 1}},
	}

	// Aggregate the documents
	cursor, err := collection.Aggregate(ctx, pipeline)
	if err != nil {
		panic(err)
	}

	// Make a slice of slices of group documents
	var temps = make([]model.TracingData, 0, 5)
	var i = 0
	for cursor.Next(ctx) {
		var t model.TracingData
		err = cursor.Decode(&t)
		if err != nil {
			_ = cursor.Close(ctx)
			panic(err)
		}
		temps = append(temps, t)
		i++
	}
	if err = cursor.Err(); err != nil {
		_ = cursor.Close(ctx)
		panic(err)
	}
	if err = cursor.Close(ctx); err != nil {
		panic(err)
	}

	pipeline = mongo.Pipeline{
		bson.D{
			{"$match", bson.D{{"parent.spanid", temps[0].SpanContext.SpanID}}}},
	}

	// Print the documents in each bucket
	fmt.Println()
	for i := range temps {
		fmt.Println(temps[i])
	}
}

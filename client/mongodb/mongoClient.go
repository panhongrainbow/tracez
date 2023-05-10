package mongodb

import (
	"context"
	"github.com/panhongrainbow/tracez/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

// TracingMongo is a mongodb client dedicating to organizing tracing data in MongoDB.
type TracingMongo struct {
	client *mongo.Client
}

// New is to connect MongoDB, return TracingMongo struct or error.
// The uri parameter could be input "mongodb://localhost:27017" for example.
func New(uri string) (tm *TracingMongo, err error) {
	// Connect to MongoDB
	tm = new(TracingMongo)
	tm.client, err = mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	return
}

// Close is to disconnect MongoDB client, return error or nil.
func (tm *TracingMongo) Close() (err error) {
	err = tm.client.Disconnect(context.TODO())
	return
}

// TracingCollection is a tracing collection in MongoDB.
type TracingCollection mongo.Collection

// PickUpDocument is to pick up a document from MongoDB, return TracingCollection struct or error.
// The database parameter could be input "openTelemetry2mongodb" for example.
// The collection parameter could be input "openTelemetry2mongodb" for example.
func (tm *TracingMongo) PickUpDocument(database, collection string) (tc *TracingCollection, err error) {
	mongoCol := tm.client.Database(database).Collection(collection)
	tc = (*TracingCollection)(mongoCol)
	return
}

type TracingFilter struct {
	key   string
	value string
}

// Search (function) searches data, gets parameters, constructs filters, aggregates, parses, handles errors, and returns tracing data.
// The key parameter could be input "instrumentationlibrary.name" for example.
// The value parameter could be input "mainTracer" for example.
func (tc *TracingCollection) Search(start, end time.Time, limit int, pair ...TracingFilter) (tracingData []model.TracingData) {
	// Construct multiple filters for the next aggregation pipeline
	var filters = make([]bson.M, len(pair)+2)
	// Construct filters from matching key-value pairs
	for i := 0; i < len(pair); i++ {
		filters[i] = bson.M{pair[i].key: pair[i].value}
	}
	// Add time range filters
	filters[len(pair)] = bson.M{"starttime": bson.M{"$gte": start}}
	filters[len(pair)+1] = bson.M{"starttime": bson.M{"$lte": end}}

	// Create aggregation pipeline
	ctx := context.TODO()
	pipeline := mongo.Pipeline{
		bson.D{{"$match", bson.M{"$and": filters}}},
		{{"$sort", bson.D{{"starttime", -1}}}},
		{{"$limit", limit}},
	}

	// Aggregate the documents
	cursor, err := (*mongo.Collection)(tc).Aggregate(ctx, pipeline)
	if err != nil {
		return
	}

	// Make a slice of slices of group documents
	tracingData = make([]model.TracingData, 0, limit)
	for cursor.Next(ctx) {
		var t model.TracingData
		err = cursor.Decode(&t)
		if err != nil {
			_ = cursor.Close(ctx)
			return
		}
		tracingData = append(tracingData, t)
	}

	// Error handling
	if err = cursor.Err(); err != nil {
		_ = cursor.Close(ctx)
		return
	}
	if err = cursor.Close(ctx); err != nil {
		return
	}
	return
}

/*
	pipeline = mongo.Pipeline{
		bson.D{
			{"$match", bson.D{{"parent.spanid", temps[0].SpanContext.SpanID}}}},
	}
*/

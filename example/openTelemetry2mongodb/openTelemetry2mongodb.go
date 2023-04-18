package main

import (
	"context"
	"io"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/stdout/stdouttrace"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
)

// mongowriter implements the io.Writer interface
type mongowriter struct {
	// MongoDB collection
	collection *mongo.Collection
}

// NewMongowriter returns a new Mongowriter instance
func NewMongowriter(collection *mongo.Collection) io.Writer {
	return &mongowriter{collection: collection}
}

// Write implements the io.Writer interface
func (w *mongowriter) Write(p []byte) (int, error) {
	// Insert document into MongoDB collection
	_, err := w.collection.InsertOne(context.TODO(), bson.M{
		"data": string(p),
	})
	if err != nil {
		return 0, err
	}
	return len(p), nil
}

// functionB is a function that is called by functionA
func functionB(ctx context.Context) {
	tracer := otel.Tracer("functionBTracer")
	ctx, span := tracer.Start(ctx, "functionBSpan")
	defer span.End()

	time.Sleep(time.Second)
}

// functionA calls functionB
func functionA(ctx context.Context) {
	tracer := otel.Tracer("functionATracer")
	ctx, span := tracer.Start(ctx, "functionASpan")
	defer span.End()

	functionB(ctx)
}

// main is the entry point for this example
func main() {
	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://127.0.0.1:27017"))
	if err != nil {
		log.Fatal(err)
	}
	collection := client.Database("openTelemetry2mongodb").Collection("openTelemetry2mongodb")

	// Create Mongowriter
	writer := NewMongowriter(collection)

	// Create stdout exporter
	exporter, err := stdouttrace.New(
		stdouttrace.WithWriter(writer),
		stdouttrace.WithPrettyPrint(),
	)
	if err != nil {
		log.Fatal(err)
	}

	// Create TracerProvider
	tp := sdktrace.NewTracerProvider(
		sdktrace.WithSampler(sdktrace.AlwaysSample()),
		sdktrace.WithBatcher(exporter),
	)
	// Close TracerProvider when main() returns
	defer tp.Shutdown(context.Background())

	// Set global TracerProvider
	otel.SetTracerProvider(tp)

	// Use OpenTelemetry Tracer to create spans
	tracer := otel.Tracer("mainTracer")
	ctx, span := tracer.Start(context.Background(), "mainSpan")
	// Close span when function returns
	defer span.End()

	// Call functionA
	functionA(ctx)

	// Add event to span
	span.AddEvent("mainEvent")
}

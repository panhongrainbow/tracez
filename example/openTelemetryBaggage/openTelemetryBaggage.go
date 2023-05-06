package main

import (
	"fmt"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	"os"
	"time"
)

/*
When using context in Go language, you need to pay specific attention to the problem of data race
I'm considering using baggage to handle it
https://opentelemetry.io/docs/concepts/signals/baggage/
*/

import (
	"context"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/baggage"
	"go.opentelemetry.io/otel/exporters/stdout/stdouttrace"
	"log"
)

// functionB is a function that is called by functionA
func functionB(ctx context.Context) {
	// List baggage
	ListBaggage(ctx)

	// Create a tracer and a span
	tracer := otel.Tracer("functionBTracer")
	ctx, span := tracer.Start(ctx, "functionBSpan")
	defer span.End()

	// Add baggage to context
	var err error
	ctx, err = WithBaggage(ctx, "3")
	if err != nil {
		return
	}

	// List baggage
	ListBaggage(ctx)

	// Sleep for 1 second
	time.Sleep(time.Second)
}

// functionA calls functionB
func functionA(ctx context.Context) {
	// Create a tracer and a span
	tracer := otel.Tracer("functionATracer")
	ctx, span := tracer.Start(ctx, "functionASpan")
	defer span.End()

	// Add baggage to context
	var err error
	ctx, err = WithBaggage(ctx, "2")
	if err != nil {
		return
	}

	// Call functionB
	functionB(ctx)
}

// main creates a span and calls functionA
func main() {
	// Create log file
	file, err := os.OpenFile("openTelemetryBaggage.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}
	// Close file when main() returns
	defer file.Close()

	// Create stdout exporter
	exporter, err := stdouttrace.New(
		stdouttrace.WithWriter(file),
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

	// Add baggage to context
	ctx, err = WithBaggage(ctx, "1")
	if err != nil {
		return
	}

	// end span when main() returns
	defer span.End()

	// Call functionA
	functionA(ctx)

	// Add event to span
	span.AddEvent("mainEvent")
}

// WithBaggage adds baggage to context
func WithBaggage(ctx context.Context, value string) (ret context.Context, err error) {
	// Create baggage member
	var member baggage.Member
	member, err = baggage.NewMember("key", value)
	if err != nil {
		return
	}

	// Create baggage
	var bag baggage.Baggage
	bag, err = baggage.New(member)
	if err != nil {
		return
	}

	// Add baggage to context
	ret = baggage.ContextWithBaggage(ctx, bag)

	// Return context
	return
}

// ListBaggage lists baggage
func ListBaggage(ctx context.Context) {
	// Get baggage from context
	bag := baggage.FromContext(ctx)
	if bag.Len() > 0 {
		for _, kv := range bag.Members() {
			// Print baggage
			fmt.Println(kv.Key(), kv.Value())
		}
	}
}

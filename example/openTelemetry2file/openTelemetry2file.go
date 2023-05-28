package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/stdout/stdouttrace"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
)

// functionB is a function that is called by functionA
func functionB(ctx context.Context) {
	tracer := otel.Tracer("functionBTracer")
	ctx, span := tracer.Start(ctx, "functionBSpan")
	span.SetAttributes(attribute.KeyValue{
		Key:   "ParameterB",
		Value: attribute.StringValue("ValueB"),
	})
	span.RecordError(fmt.Errorf("error"))
	defer span.End()

	time.Sleep(time.Second)
}

// functionA calls functionB
func functionA(ctx context.Context) {
	tracer := otel.Tracer("functionATracer")
	ctx, span := tracer.Start(ctx, "functionASpan")
	span.SetAttributes(attribute.KeyValue{
		Key:   "ParameterA",
		Value: attribute.StringValue("ValueA"),
	})
	defer span.End()

	functionB(ctx)
}

// main creates a span and calls functionA
func main() {
	// Create log file
	file, err := os.OpenFile("openTelemetry2fileLogs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
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
	span.SetAttributes(attribute.KeyValue{
		Key:   "ParameterMain",
		Value: attribute.StringValue("ValueMain"),
	})

	// end span when main() returns
	defer span.End()

	// Call functionA
	functionA(ctx)

	// Add event to span
	span.AddEvent("mainEvent")
}

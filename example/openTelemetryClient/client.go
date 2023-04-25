package main

import (
	"context"
	"fmt"
	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/stdout/stdouttrace"
	"go.opentelemetry.io/otel/sdk/trace"
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	// Create log file
	file, err := os.OpenFile("openTelemetryInClient.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
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

	// Create a new trace provider with the OTLP exporter
	tp := trace.NewTracerProvider(
		trace.WithSyncer(exporter),
		trace.WithSampler(trace.AlwaysSample()),
	)

	// Set the global trace provider
	otel.SetTracerProvider(tp)

	// Use the tracer provider to create a new tracer
	tracer := otel.Tracer("example")

	// Create a new HTTP client
	client := &http.Client{
		Transport: otelhttp.NewTransport(http.DefaultTransport),
	}

	// Use the tracer to create a new span
	ctx, span := tracer.Start(context.Background(), "example-client")
	defer span.End()

	// Set attributes on the span
	span.SetAttributes(
		attribute.KeyValue{
			Key:   "http.method",
			Value: attribute.StringValue("GET"),
		},
		attribute.KeyValue{
			Key:   "http.url",
			Value: attribute.StringValue("http://localhost:8080/"),
		},
	)

	// Make an HTTP request
	req, err := http.NewRequestWithContext(ctx, "GET", "http://localhost:8080/", nil)
	if err != nil {
		panic(err)
	}

	req.Header.Set("trace-id", span.SpanContext().TraceID().String())
	req.Header.Set("span-id", span.SpanContext().SpanID().String())

	res, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	// Read the response
	body, err := io.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(body))
}

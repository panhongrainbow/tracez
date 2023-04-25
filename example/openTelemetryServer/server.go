package main

import (
	"context"
	"fmt"
	ot "github.com/opentracing/opentracing-go"
	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/stdout/stdouttrace"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	contextTrace "go.opentelemetry.io/otel/trace"
	"time"

	"log"
	"net/http"
	"os"
)

type spanPropagator struct{}

func (p *spanPropagator) SpanContextFromRequest(r *http.Request) (res bool, sc contextTrace.SpanContext) {
	// 从 HTTP Request 中提取 SpanContext
	fmt.Println("1>>>", r.Header.Get("trace-id"))
	fmt.Println("2>>>", r.Header.Get("span-id"))
	traceID, _ := contextTrace.TraceIDFromHex(r.Header.Get("trace-id"))
	sc = sc.WithTraceID(traceID)
	spanID, _ := contextTrace.SpanIDFromHex(r.Header.Get("span-id"))
	sc = sc.WithSpanID(spanID)
	return
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

// main creates a span and calls functionA
func main() {
	// Create log file
	file, err := os.OpenFile("openTelemetryInServer.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
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

	// end span when main() returns
	defer span.End()

	// Add event to span
	span.AddEvent("mainEvent")

	// Create a new HTTP handler
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		sp := spanPropagator{}
		_, sc := sp.SpanContextFromRequest(r)

		// Use the tracer to create a new span
		ctx = contextTrace.ContextWithRemoteSpanContext(ctx, sc)
		span, _ := ot.StartSpanFromContext(ctx, "example-handler")
		defer span.Finish()
		// _, span := tracer.Start(r.Context(), "example-handler")
		// defer span.End()

		// Set attributes on the span
		/*span.SetAttributes(
			attribute.KeyValue{
				Key:   "http.method",
				Value: attribute.StringValue(r.Method),
			},
			attribute.KeyValue{
				Key:   "http.url",
				Value: attribute.StringValue(r.URL.String()),
			},
		)*/

		// Call functionA
		functionA(ctx)

		// Write a response
		fmt.Fprintln(w, "Hello, world!")
	})

	if err := http.ListenAndServe(":8080",
		otelhttp.NewHandler(&handler, "server",
			otelhttp.WithMessageEvents(otelhttp.ReadEvents, otelhttp.WriteEvents),
		),
	); err != nil {
		log.Fatal(err)
	}

}

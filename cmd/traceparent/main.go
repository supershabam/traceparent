package main

import (
	"context"
	"fmt"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
)

func main() {
	ctx := context.Background()
	tp := sdktrace.NewTracerProvider()

	otel.SetTracerProvider(tp)
	ctx, _ = otel.Tracer("traceparent").Start(ctx, "main")

	tc := propagation.TraceContext{}
	hc := propagation.HeaderCarrier{}

	tc.Inject(ctx, hc)

	traceparent := hc.Get("Traceparent")

	fmt.Printf("%s", traceparent)
}

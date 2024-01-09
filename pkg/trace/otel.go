package trace

import (
	"context"
	"os"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/stdout/stdouttrace"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	"go.opentelemetry.io/otel/sdk/trace"
)

func SetupTracing() (func(context.Context), error) {
	// Set up propagator
	prop := propagation.NewCompositeTextMapPropagator(
		propagation.TraceContext{},
		propagation.Baggage{},
	)
	otel.SetTextMapPropagator(prop)

	// Setup resource
	res := resource.Default()

	// Set up trace file
	tf, err := os.Create("/tmp/trace")
	if err != nil {
		return nil, err
	}

	exporter, err := stdouttrace.New(
		stdouttrace.WithWriter(tf),
	)
	if err != nil {
		return nil, err
	}

	// Set up trace provider
	provider := trace.NewTracerProvider(
		trace.WithBatcher(exporter),
		trace.WithResource(res),
	)
	if err != nil {
		return nil, err
	}
	otel.SetTracerProvider(provider)

	return func(ctx context.Context) {
		_ = provider.Shutdown(ctx)
		_ = tf.Close()
	}, nil
}

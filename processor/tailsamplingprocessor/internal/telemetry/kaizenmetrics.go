package telemetry

import (
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/otel/metric"
)

var SampledSpanSizePerService metric.Int64Counter

//var SampledSpanSizePerServiceHistogram metric.Int64Histogram

func ConfigureKaizenMetrics(telemetry *component.TelemetrySettings) error {
	meter := telemetry.MeterProvider.Meter("tailsamplingprocessorkaizenmetrics")

	var err error

	SampledSpanSizePerService, err = meter.Int64Counter("otelcol_processor_tail_sampling_sampled_spans_size",
		metric.WithDescription("The size of sampled spans in bytes per service"),
		metric.WithUnit("{bytes}"))

	//SampledSpanSizePerServiceHistogram, err = meter.Int64Histogram("span-size-per-service",
	//	metric.WithDescription("The size of spans in bytes per service"),
	//	metric.WithUnit("{bytes}"),
	//	metric.WithExplicitBucketBoundaries(50, 100, 150, 200, 250, 300, 350, 400, 450, 500, 700, 1000, 2000, 4000, 8000))

	if err != nil {
		return err
	}

	return nil
}

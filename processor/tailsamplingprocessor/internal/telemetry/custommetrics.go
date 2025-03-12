package telemetry

import (
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/otel/metric"
)

var SampledSpanSizePerService metric.Int64Histogram
var MissingAttributesPerService metric.Int64Gauge

func ConfigureCustomMetrics(telemetry *component.TelemetrySettings) error {
	meter := telemetry.MeterProvider.Meter("tailsamplingprocessorcustommetrics")

	var err error

	SampledSpanSizePerService, err = meter.Int64Histogram("otelcol_processor_tail_sampling_sampled_spans_size",
		metric.WithDescription("The size of sampled spans in bytes per service"),
		metric.WithUnit("{bytes}"),
		metric.WithExplicitBucketBoundaries(0))

	if err != nil {
		return err
	}

	MissingAttributesPerService, err = meter.Int64Gauge("otelcol_processor_tail_sampling_missing_resource_attributes",
		metric.WithDescription("Missing resource attributes per service"),
		metric.WithUnit("1"))

	return nil
}

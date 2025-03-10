package telemetry

import (
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/otel/metric"
)

var SampledSpanSizePerService metric.Int64Counter

func ConfigureCustomMetrics(telemetry *component.TelemetrySettings) error {
	meter := telemetry.MeterProvider.Meter("tailsamplingprocessorcustommetrics")

	var err error

	SampledSpanSizePerService, err = meter.Int64Counter("otelcol_processor_tail_sampling_sampled_spans_size",
		metric.WithDescription("The size of sampled spans in bytes per service"),
		metric.WithUnit("{bytes}"))

	if err != nil {
		return err
	}

	return nil
}

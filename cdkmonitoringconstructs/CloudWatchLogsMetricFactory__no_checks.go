//go:build no_runtime_type_checking

package cdkmonitoringconstructs

// Building without runtime type checking enabled, so all the below just return nil

func validateNewCloudWatchLogsMetricFactoryParameters(metricFactory MetricFactory, props *CloudWatchLogsMetricFactoryProps) error {
	return nil
}

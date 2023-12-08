//go:build no_runtime_type_checking

package cdkmonitoringconstructs

// Building without runtime type checking enabled, so all the below just return nil

func validateNewS3BucketMetricFactoryParameters(metricFactory MetricFactory, props *S3BucketMetricFactoryProps) error {
	return nil
}


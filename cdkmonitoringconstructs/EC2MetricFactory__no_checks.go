//go:build no_runtime_type_checking

package cdkmonitoringconstructs

// Building without runtime type checking enabled, so all the below just return nil

func validateNewEC2MetricFactoryParameters(metricFactory MetricFactory, props *EC2MetricFactoryProps) error {
	return nil
}


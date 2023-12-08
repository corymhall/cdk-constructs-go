//go:build no_runtime_type_checking

package cdkmonitoringconstructs

// Building without runtime type checking enabled, so all the below just return nil

func validateNewWafV2MetricFactoryParameters(metricFactory MetricFactory, props *WafV2MetricFactoryProps) error {
	return nil
}


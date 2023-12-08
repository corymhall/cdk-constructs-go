//go:build no_runtime_type_checking

package cdkmonitoringconstructs

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_ApiGatewayMetricFactory) validateMetricLatencyInMillisParameters(latencyType LatencyType) error {
	return nil
}

func validateNewApiGatewayMetricFactoryParameters(metricFactory MetricFactory, props *ApiGatewayMetricFactoryProps) error {
	return nil
}


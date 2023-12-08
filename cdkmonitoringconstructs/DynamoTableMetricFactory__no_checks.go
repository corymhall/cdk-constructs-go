//go:build no_runtime_type_checking

package cdkmonitoringconstructs

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DynamoTableMetricFactory) validateMetricAverageSuccessfulRequestLatencyInMillisParameters(operation awsdynamodb.Operation) error {
	return nil
}

func validateNewDynamoTableMetricFactoryParameters(metricFactory MetricFactory, props *DynamoTableMetricFactoryProps) error {
	return nil
}


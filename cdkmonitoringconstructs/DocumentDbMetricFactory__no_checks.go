//go:build no_runtime_type_checking

package cdkmonitoringconstructs

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DocumentDbMetricFactory) validateMetricReadLatencyInMillisParameters(latencyType LatencyType) error {
	return nil
}

func (d *jsiiProxy_DocumentDbMetricFactory) validateMetricWriteLatencyInMillisParameters(latencyType LatencyType) error {
	return nil
}

func validateNewDocumentDbMetricFactoryParameters(metricFactory MetricFactory, props *DocumentDbMetricFactoryProps) error {
	return nil
}


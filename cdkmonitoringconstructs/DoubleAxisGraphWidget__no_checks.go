//go:build no_runtime_type_checking

package cdkmonitoringconstructs

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DoubleAxisGraphWidget) validateAddLeftMetricParameters(metric awscloudwatch.IMetric) error {
	return nil
}

func (d *jsiiProxy_DoubleAxisGraphWidget) validateAddRightMetricParameters(metric awscloudwatch.IMetric) error {
	return nil
}

func (d *jsiiProxy_DoubleAxisGraphWidget) validatePositionParameters(x *float64, y *float64) error {
	return nil
}

func validateNewDoubleAxisGraphWidgetParameters(props *DoubleAxisGraphWidgetProps) error {
	return nil
}


//go:build no_runtime_type_checking

package cdkmonitoringconstructs

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_SingleAxisGraphWidget) validateAddLeftMetricParameters(metric awscloudwatch.IMetric) error {
	return nil
}

func (s *jsiiProxy_SingleAxisGraphWidget) validateAddRightMetricParameters(metric awscloudwatch.IMetric) error {
	return nil
}

func (s *jsiiProxy_SingleAxisGraphWidget) validatePositionParameters(x *float64, y *float64) error {
	return nil
}

func validateNewSingleAxisGraphWidgetParameters(props *SingleAxisGraphWidgetProps) error {
	return nil
}


//go:build no_runtime_type_checking

package cdkmonitoringconstructs

// Building without runtime type checking enabled, so all the below just return nil

func (m *jsiiProxy_MetricFactory) validateAdaptMetricParameters(metric interface{}) error {
	return nil
}

func (m *jsiiProxy_MetricFactory) validateAdaptMetricPreservingPeriodParameters(metric interface{}) error {
	return nil
}

func (m *jsiiProxy_MetricFactory) validateAddAdditionalDimensionsParameters(target *map[string]*string, additionalDimensions *map[string]*string) error {
	return nil
}

func (m *jsiiProxy_MetricFactory) validateCreateMetricParameters(metricName *string, statistic MetricStatistic) error {
	return nil
}

func (m *jsiiProxy_MetricFactory) validateCreateMetricAnomalyDetectionParameters(metric awscloudwatch.IMetric, stdev *float64, label *string) error {
	return nil
}

func (m *jsiiProxy_MetricFactory) validateCreateMetricMathParameters(expression *string, usingMetrics *map[string]awscloudwatch.IMetric, label *string) error {
	return nil
}

func (m *jsiiProxy_MetricFactory) validateCreateMetricSearchParameters(query *string, dimensionsMap *map[string]*string, statistic MetricStatistic) error {
	return nil
}

func (m *jsiiProxy_MetricFactory) validateDivideMetricParameters(metric interface{}, divisor *float64, label *string) error {
	return nil
}

func (m *jsiiProxy_MetricFactory) validateMultiplyMetricParameters(metric interface{}, multiplier *float64, label *string) error {
	return nil
}

func (m *jsiiProxy_MetricFactory) validateSanitizeMetricExpressionIdSuffixParameters(expressionId *string) error {
	return nil
}

func (m *jsiiProxy_MetricFactory) validateToRateParameters(metric interface{}, method RateComputationMethod) error {
	return nil
}

func validateNewMetricFactoryParameters(props *MetricFactoryProps) error {
	return nil
}


//go:build !no_runtime_type_checking

package cdkmonitoringconstructs

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudwatch"
)

func (m *jsiiProxy_MetricFactory) validateAdaptMetricParameters(metric interface{}) error {
	if metric == nil {
		return fmt.Errorf("parameter metric is required, but nil was provided")
	}
	switch metric.(type) {
	case awscloudwatch.Metric:
		// ok
	case awscloudwatch.MathExpression:
		// ok
	default:
		if !_jsii_.IsAnonymousProxy(metric) {
			return fmt.Errorf("parameter metric must be one of the allowed types: awscloudwatch.Metric, awscloudwatch.MathExpression; received %#v (a %T)", metric, metric)
		}
	}

	return nil
}

func (m *jsiiProxy_MetricFactory) validateAdaptMetricPreservingPeriodParameters(metric interface{}) error {
	if metric == nil {
		return fmt.Errorf("parameter metric is required, but nil was provided")
	}
	switch metric.(type) {
	case awscloudwatch.Metric:
		// ok
	case awscloudwatch.MathExpression:
		// ok
	default:
		if !_jsii_.IsAnonymousProxy(metric) {
			return fmt.Errorf("parameter metric must be one of the allowed types: awscloudwatch.Metric, awscloudwatch.MathExpression; received %#v (a %T)", metric, metric)
		}
	}

	return nil
}

func (m *jsiiProxy_MetricFactory) validateAddAdditionalDimensionsParameters(target *map[string]*string, additionalDimensions *map[string]*string) error {
	if target == nil {
		return fmt.Errorf("parameter target is required, but nil was provided")
	}

	if additionalDimensions == nil {
		return fmt.Errorf("parameter additionalDimensions is required, but nil was provided")
	}

	return nil
}

func (m *jsiiProxy_MetricFactory) validateCreateMetricParameters(metricName *string, statistic MetricStatistic) error {
	if metricName == nil {
		return fmt.Errorf("parameter metricName is required, but nil was provided")
	}

	if statistic == "" {
		return fmt.Errorf("parameter statistic is required, but nil was provided")
	}

	return nil
}

func (m *jsiiProxy_MetricFactory) validateCreateMetricAnomalyDetectionParameters(metric awscloudwatch.IMetric, stdev *float64, label *string) error {
	if metric == nil {
		return fmt.Errorf("parameter metric is required, but nil was provided")
	}

	if stdev == nil {
		return fmt.Errorf("parameter stdev is required, but nil was provided")
	}

	if label == nil {
		return fmt.Errorf("parameter label is required, but nil was provided")
	}

	return nil
}

func (m *jsiiProxy_MetricFactory) validateCreateMetricMathParameters(expression *string, usingMetrics *map[string]awscloudwatch.IMetric, label *string) error {
	if expression == nil {
		return fmt.Errorf("parameter expression is required, but nil was provided")
	}

	if usingMetrics == nil {
		return fmt.Errorf("parameter usingMetrics is required, but nil was provided")
	}

	if label == nil {
		return fmt.Errorf("parameter label is required, but nil was provided")
	}

	return nil
}

func (m *jsiiProxy_MetricFactory) validateCreateMetricSearchParameters(query *string, dimensionsMap *map[string]*string, statistic MetricStatistic) error {
	if query == nil {
		return fmt.Errorf("parameter query is required, but nil was provided")
	}

	if dimensionsMap == nil {
		return fmt.Errorf("parameter dimensionsMap is required, but nil was provided")
	}

	if statistic == "" {
		return fmt.Errorf("parameter statistic is required, but nil was provided")
	}

	return nil
}

func (m *jsiiProxy_MetricFactory) validateDivideMetricParameters(metric interface{}, divisor *float64, label *string) error {
	if metric == nil {
		return fmt.Errorf("parameter metric is required, but nil was provided")
	}
	switch metric.(type) {
	case awscloudwatch.Metric:
		// ok
	case awscloudwatch.MathExpression:
		// ok
	default:
		if !_jsii_.IsAnonymousProxy(metric) {
			return fmt.Errorf("parameter metric must be one of the allowed types: awscloudwatch.Metric, awscloudwatch.MathExpression; received %#v (a %T)", metric, metric)
		}
	}

	if divisor == nil {
		return fmt.Errorf("parameter divisor is required, but nil was provided")
	}

	if label == nil {
		return fmt.Errorf("parameter label is required, but nil was provided")
	}

	return nil
}

func (m *jsiiProxy_MetricFactory) validateMultiplyMetricParameters(metric interface{}, multiplier *float64, label *string) error {
	if metric == nil {
		return fmt.Errorf("parameter metric is required, but nil was provided")
	}
	switch metric.(type) {
	case awscloudwatch.Metric:
		// ok
	case awscloudwatch.MathExpression:
		// ok
	default:
		if !_jsii_.IsAnonymousProxy(metric) {
			return fmt.Errorf("parameter metric must be one of the allowed types: awscloudwatch.Metric, awscloudwatch.MathExpression; received %#v (a %T)", metric, metric)
		}
	}

	if multiplier == nil {
		return fmt.Errorf("parameter multiplier is required, but nil was provided")
	}

	if label == nil {
		return fmt.Errorf("parameter label is required, but nil was provided")
	}

	return nil
}

func (m *jsiiProxy_MetricFactory) validateSanitizeMetricExpressionIdSuffixParameters(expressionId *string) error {
	if expressionId == nil {
		return fmt.Errorf("parameter expressionId is required, but nil was provided")
	}

	return nil
}

func (m *jsiiProxy_MetricFactory) validateToRateParameters(metric interface{}, method RateComputationMethod) error {
	if metric == nil {
		return fmt.Errorf("parameter metric is required, but nil was provided")
	}
	switch metric.(type) {
	case awscloudwatch.Metric:
		// ok
	case awscloudwatch.MathExpression:
		// ok
	default:
		if !_jsii_.IsAnonymousProxy(metric) {
			return fmt.Errorf("parameter metric must be one of the allowed types: awscloudwatch.Metric, awscloudwatch.MathExpression; received %#v (a %T)", metric, metric)
		}
	}

	if method == "" {
		return fmt.Errorf("parameter method is required, but nil was provided")
	}

	return nil
}

func validateNewMetricFactoryParameters(props *MetricFactoryProps) error {
	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}


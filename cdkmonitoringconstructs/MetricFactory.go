package cdkmonitoringconstructs

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/corymhall/cdk-constructs-go/cdkmonitoringconstructs/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudwatch"
)

// Experimental.
type MetricFactory interface {
	// Experimental.
	GlobalDefaults() *MetricFactoryDefaults
	// Adapts properties of a foreign metric (metric created outside of this metric factory) to comply with the global defaults.
	//
	// Might modify namespace and metric period.
	// Experimental.
	AdaptMetric(metric interface{}) interface{}
	// Adapts properties of a foreign metric (metric created outside of this metric factory) to comply with the global defaults.
	//
	// Might modify namespace. Preserves metric period.
	// Experimental.
	AdaptMetricPreservingPeriod(metric interface{}) interface{}
	// Merges the given additional dimensions to the given target dimension hash.
	//
	// All existing dimensions with the same key are replaced.
	// Experimental.
	AddAdditionalDimensions(target *map[string]*string, additionalDimensions *map[string]*string)
	// Factory method that creates a metric.
	//
	// The metric properties will already be updated to comply with the global defaults.
	// Experimental.
	CreateMetric(metricName *string, statistic MetricStatistic, label *string, dimensionsMap *map[string]*string, color *string, namespace *string, period awscdk.Duration, region *string, account *string) interface{}
	// Factory method that creates anomaly detection on a metric.
	//
	// Anomaly occurs whenever a metric value falls outside of a precomputed range of predicted values.
	// The detection does not need any setup. The model will start learning automatically and should be ready in a few minutes.
	// Usually, the anomaly detection is paired with an alarm.
	// See: https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/CloudWatch_Anomaly_Detection.html
	//
	// Experimental.
	CreateMetricAnomalyDetection(metric awscloudwatch.IMetric, stdev *float64, label *string, color *string, expressionId *string, period awscdk.Duration, region *string, account *string) interface{}
	// Factory method that creates a metric math expression.
	//
	// The metric properties will already be updated to comply with the global defaults.
	// Experimental.
	CreateMetricMath(expression *string, usingMetrics *map[string]awscloudwatch.IMetric, label *string, color *string, period awscdk.Duration, region *string, account *string) interface{}
	// Factory method that creates a metric search query.
	//
	// The metric properties will already be updated to comply with the global defaults.
	// If you want to match "any value" of a specific dimension, please use `undefined` value representation in your consumer language.
	// (For example, `undefined as any as string` in TypeScript, due to JSII typing quirks.)
	// Experimental.
	CreateMetricSearch(query *string, dimensionsMap *map[string]*string, statistic MetricStatistic, namespace *string, label *string, period awscdk.Duration, region *string, account *string) awscloudwatch.IMetric
	// Creates a metric math expression that divides the given metric by given coefficient.
	//
	// Does nothing if the divisor is one. Preserves the metric period.
	// Experimental.
	DivideMetric(metric interface{}, divisor *float64, label *string, expressionId *string) interface{}
	// Returns the given namespace (if defined) or the global namespace as a fallback.
	//
	// If there is no namespace to fallback to (neither the custom or the default one), it will fail.
	// Experimental.
	GetNamespaceWithFallback(value *string) *string
	// Creates a metric math expression that multiplies the given metric by given coefficient.
	//
	// Does nothing if the multiplier is one. Preserves the metric period.
	// Experimental.
	MultiplyMetric(metric interface{}, multiplier *float64, label *string, expressionId *string) interface{}
	// Helper method that helps to sanitize the given expression ID and removes all invalid characters.
	//
	// Valid expression ID regexp is the following: ^[a-z][a-zA-Z0-9_]*$
	// As this is just to validate a suffix and not the whole ID, we do not have to verify the first lower case letter.
	// Experimental.
	SanitizeMetricExpressionIdSuffix(expressionId *string) *string
	// Creates a metric math expression that computes a rate from a regular metric.
	//
	// For example, it allows you to compute rate per second (TPS), per minute, or just an average of your transactions.
	// Experimental.
	ToRate(metric interface{}, method RateComputationMethod, addStatsToLabel *bool, expressionId *string, fillWithZeroes *bool) interface{}
}

// The jsii proxy struct for MetricFactory
type jsiiProxy_MetricFactory struct {
	_ byte // padding
}

func (j *jsiiProxy_MetricFactory) GlobalDefaults() *MetricFactoryDefaults {
	var returns *MetricFactoryDefaults
	_jsii_.Get(
		j,
		"globalDefaults",
		&returns,
	)
	return returns
}


// Experimental.
func NewMetricFactory(props *MetricFactoryProps) MetricFactory {
	_init_.Initialize()

	if err := validateNewMetricFactoryParameters(props); err != nil {
		panic(err)
	}
	j := jsiiProxy_MetricFactory{}

	_jsii_.Create(
		"cdk-monitoring-constructs.MetricFactory",
		[]interface{}{props},
		&j,
	)

	return &j
}

// Experimental.
func NewMetricFactory_Override(m MetricFactory, props *MetricFactoryProps) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk-monitoring-constructs.MetricFactory",
		[]interface{}{props},
		m,
	)
}

func (m *jsiiProxy_MetricFactory) AdaptMetric(metric interface{}) interface{} {
	if err := m.validateAdaptMetricParameters(metric); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		m,
		"adaptMetric",
		[]interface{}{metric},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MetricFactory) AdaptMetricPreservingPeriod(metric interface{}) interface{} {
	if err := m.validateAdaptMetricPreservingPeriodParameters(metric); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		m,
		"adaptMetricPreservingPeriod",
		[]interface{}{metric},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MetricFactory) AddAdditionalDimensions(target *map[string]*string, additionalDimensions *map[string]*string) {
	if err := m.validateAddAdditionalDimensionsParameters(target, additionalDimensions); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"addAdditionalDimensions",
		[]interface{}{target, additionalDimensions},
	)
}

func (m *jsiiProxy_MetricFactory) CreateMetric(metricName *string, statistic MetricStatistic, label *string, dimensionsMap *map[string]*string, color *string, namespace *string, period awscdk.Duration, region *string, account *string) interface{} {
	if err := m.validateCreateMetricParameters(metricName, statistic); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		m,
		"createMetric",
		[]interface{}{metricName, statistic, label, dimensionsMap, color, namespace, period, region, account},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MetricFactory) CreateMetricAnomalyDetection(metric awscloudwatch.IMetric, stdev *float64, label *string, color *string, expressionId *string, period awscdk.Duration, region *string, account *string) interface{} {
	if err := m.validateCreateMetricAnomalyDetectionParameters(metric, stdev, label); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		m,
		"createMetricAnomalyDetection",
		[]interface{}{metric, stdev, label, color, expressionId, period, region, account},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MetricFactory) CreateMetricMath(expression *string, usingMetrics *map[string]awscloudwatch.IMetric, label *string, color *string, period awscdk.Duration, region *string, account *string) interface{} {
	if err := m.validateCreateMetricMathParameters(expression, usingMetrics, label); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		m,
		"createMetricMath",
		[]interface{}{expression, usingMetrics, label, color, period, region, account},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MetricFactory) CreateMetricSearch(query *string, dimensionsMap *map[string]*string, statistic MetricStatistic, namespace *string, label *string, period awscdk.Duration, region *string, account *string) awscloudwatch.IMetric {
	if err := m.validateCreateMetricSearchParameters(query, dimensionsMap, statistic); err != nil {
		panic(err)
	}
	var returns awscloudwatch.IMetric

	_jsii_.Invoke(
		m,
		"createMetricSearch",
		[]interface{}{query, dimensionsMap, statistic, namespace, label, period, region, account},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MetricFactory) DivideMetric(metric interface{}, divisor *float64, label *string, expressionId *string) interface{} {
	if err := m.validateDivideMetricParameters(metric, divisor, label); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		m,
		"divideMetric",
		[]interface{}{metric, divisor, label, expressionId},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MetricFactory) GetNamespaceWithFallback(value *string) *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"getNamespaceWithFallback",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MetricFactory) MultiplyMetric(metric interface{}, multiplier *float64, label *string, expressionId *string) interface{} {
	if err := m.validateMultiplyMetricParameters(metric, multiplier, label); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		m,
		"multiplyMetric",
		[]interface{}{metric, multiplier, label, expressionId},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MetricFactory) SanitizeMetricExpressionIdSuffix(expressionId *string) *string {
	if err := m.validateSanitizeMetricExpressionIdSuffixParameters(expressionId); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		m,
		"sanitizeMetricExpressionIdSuffix",
		[]interface{}{expressionId},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MetricFactory) ToRate(metric interface{}, method RateComputationMethod, addStatsToLabel *bool, expressionId *string, fillWithZeroes *bool) interface{} {
	if err := m.validateToRateParameters(metric, method); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		m,
		"toRate",
		[]interface{}{metric, method, addStatsToLabel, expressionId, fillWithZeroes},
		&returns,
	)

	return returns
}


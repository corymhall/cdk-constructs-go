package cdkmonitoringconstructs

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/corymhall/cdk-constructs-go/cdkmonitoringconstructs/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudwatch"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/corymhall/cdk-constructs-go/cdkmonitoringconstructs/internal"
)

// Captures specific MathExpression for anomaly detection, for which alarm generation is different.
//
// Added to overcome certain CDK limitations at the time of writing.
// See: https://github.com/aws/aws-cdk/issues/10540
//
// Experimental.
type AnomalyDetectionMathExpression interface {
	awscloudwatch.MathExpression
	// The hex color code, prefixed with '#' (e.g. '#00ff00'), to use when this metric is rendered on a graph. The `Color` class has a set of standard colors that can be used here.
	// Experimental.
	Color() *string
	// The expression defining the metric.
	// Experimental.
	Expression() *string
	// Label for this metric when added to a Graph.
	// Experimental.
	Label() *string
	// Aggregation period of this metric.
	// Experimental.
	Period() awscdk.Duration
	// Account to evaluate search expressions within.
	// Experimental.
	SearchAccount() *string
	// Region to evaluate search expressions within.
	// Experimental.
	SearchRegion() *string
	// The metrics used in the expression as KeyValuePair <id, metric>.
	// Experimental.
	UsingMetrics() *map[string]awscloudwatch.IMetric
	// Warnings generated by this math expression.
	// Deprecated: - use warningsV2.
	Warnings() *[]*string
	// Warnings generated by this math expression.
	// Experimental.
	WarningsV2() *map[string]*string
	// Make a new Alarm for this metric.
	//
	// Combines both properties that may adjust the metric (aggregation) as well
	// as alarm properties.
	// Experimental.
	CreateAlarm(scope constructs.Construct, id *string, props *awscloudwatch.CreateAlarmOptions) awscloudwatch.Alarm
	// Inspect the details of the metric object.
	// Experimental.
	ToMetricConfig() *awscloudwatch.MetricConfig
	// Returns a string representation of an object.
	// Experimental.
	ToString() *string
	// Return a copy of Metric with properties changed.
	//
	// All properties except namespace and metricName can be changed.
	// Experimental.
	With(props *awscloudwatch.MathExpressionOptions) awscloudwatch.MathExpression
}

// The jsii proxy struct for AnomalyDetectionMathExpression
type jsiiProxy_AnomalyDetectionMathExpression struct {
	internal.Type__awscloudwatchMathExpression
}

func (j *jsiiProxy_AnomalyDetectionMathExpression) Color() *string {
	var returns *string
	_jsii_.Get(
		j,
		"color",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AnomalyDetectionMathExpression) Expression() *string {
	var returns *string
	_jsii_.Get(
		j,
		"expression",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AnomalyDetectionMathExpression) Label() *string {
	var returns *string
	_jsii_.Get(
		j,
		"label",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AnomalyDetectionMathExpression) Period() awscdk.Duration {
	var returns awscdk.Duration
	_jsii_.Get(
		j,
		"period",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AnomalyDetectionMathExpression) SearchAccount() *string {
	var returns *string
	_jsii_.Get(
		j,
		"searchAccount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AnomalyDetectionMathExpression) SearchRegion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"searchRegion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AnomalyDetectionMathExpression) UsingMetrics() *map[string]awscloudwatch.IMetric {
	var returns *map[string]awscloudwatch.IMetric
	_jsii_.Get(
		j,
		"usingMetrics",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AnomalyDetectionMathExpression) Warnings() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"warnings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AnomalyDetectionMathExpression) WarningsV2() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"warningsV2",
		&returns,
	)
	return returns
}


// Experimental.
func NewAnomalyDetectionMathExpression(props *awscloudwatch.MathExpressionProps) AnomalyDetectionMathExpression {
	_init_.Initialize()

	if err := validateNewAnomalyDetectionMathExpressionParameters(props); err != nil {
		panic(err)
	}
	j := jsiiProxy_AnomalyDetectionMathExpression{}

	_jsii_.Create(
		"cdk-monitoring-constructs.AnomalyDetectionMathExpression",
		[]interface{}{props},
		&j,
	)

	return &j
}

// Experimental.
func NewAnomalyDetectionMathExpression_Override(a AnomalyDetectionMathExpression, props *awscloudwatch.MathExpressionProps) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk-monitoring-constructs.AnomalyDetectionMathExpression",
		[]interface{}{props},
		a,
	)
}

func (a *jsiiProxy_AnomalyDetectionMathExpression) CreateAlarm(scope constructs.Construct, id *string, props *awscloudwatch.CreateAlarmOptions) awscloudwatch.Alarm {
	if err := a.validateCreateAlarmParameters(scope, id, props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Alarm

	_jsii_.Invoke(
		a,
		"createAlarm",
		[]interface{}{scope, id, props},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AnomalyDetectionMathExpression) ToMetricConfig() *awscloudwatch.MetricConfig {
	var returns *awscloudwatch.MetricConfig

	_jsii_.Invoke(
		a,
		"toMetricConfig",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AnomalyDetectionMathExpression) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AnomalyDetectionMathExpression) With(props *awscloudwatch.MathExpressionOptions) awscloudwatch.MathExpression {
	if err := a.validateWithParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.MathExpression

	_jsii_.Invoke(
		a,
		"with",
		[]interface{}{props},
		&returns,
	)

	return returns
}


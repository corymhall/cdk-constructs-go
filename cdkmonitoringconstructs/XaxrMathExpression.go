package cdkmonitoringconstructs

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/corymhall/cdk-constructs-go/cdkmonitoringconstructs/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudwatch"
	"github.com/corymhall/cdk-constructs-go/cdkmonitoringconstructs/internal"
)

// Custom wrapper class for MathExpression that supports account and region specification.
// See: https://github.com/aws/aws-cdk/issues/9039
//
// Deprecated: Use MathExpression from aws-cdk-lib/aws-cloudwatch instead.
type XaxrMathExpression interface {
	awscloudwatch.IMetric
	// Inspect the details of the metric object.
	// Deprecated: Use MathExpression from aws-cdk-lib/aws-cloudwatch instead.
	ToMetricConfig() *awscloudwatch.MetricConfig
	// Deprecated: Use MathExpression from aws-cdk-lib/aws-cloudwatch instead.
	With(options *awscloudwatch.MathExpressionOptions) awscloudwatch.IMetric
}

// The jsii proxy struct for XaxrMathExpression
type jsiiProxy_XaxrMathExpression struct {
	internal.Type__awscloudwatchIMetric
}

// Deprecated: Use MathExpression from aws-cdk-lib/aws-cloudwatch instead.
func NewXaxrMathExpression(props *XaxrMathExpressionProps) XaxrMathExpression {
	_init_.Initialize()

	if err := validateNewXaxrMathExpressionParameters(props); err != nil {
		panic(err)
	}
	j := jsiiProxy_XaxrMathExpression{}

	_jsii_.Create(
		"cdk-monitoring-constructs.XaxrMathExpression",
		[]interface{}{props},
		&j,
	)

	return &j
}

// Deprecated: Use MathExpression from aws-cdk-lib/aws-cloudwatch instead.
func NewXaxrMathExpression_Override(x XaxrMathExpression, props *XaxrMathExpressionProps) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk-monitoring-constructs.XaxrMathExpression",
		[]interface{}{props},
		x,
	)
}

func (x *jsiiProxy_XaxrMathExpression) ToMetricConfig() *awscloudwatch.MetricConfig {
	var returns *awscloudwatch.MetricConfig

	_jsii_.Invoke(
		x,
		"toMetricConfig",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (x *jsiiProxy_XaxrMathExpression) With(options *awscloudwatch.MathExpressionOptions) awscloudwatch.IMetric {
	if err := x.validateWithParameters(options); err != nil {
		panic(err)
	}
	var returns awscloudwatch.IMetric

	_jsii_.Invoke(
		x,
		"with",
		[]interface{}{options},
		&returns,
	)

	return returns
}


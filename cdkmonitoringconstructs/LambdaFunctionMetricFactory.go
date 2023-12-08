package cdkmonitoringconstructs

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/corymhall/cdk-constructs-go/cdkmonitoringconstructs/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
)

// Experimental.
type LambdaFunctionMetricFactory interface {
	// Experimental.
	FillTpsWithZeroes() *bool
	// Experimental.
	LambdaFunction() awslambda.IFunction
	// Experimental.
	MetricFactory() MetricFactory
	// Experimental.
	RateComputationMethod() RateComputationMethod
	// Experimental.
	MetricConcurrentExecutions() interface{}
	// Experimental.
	MetricFaultCount() interface{}
	// Experimental.
	MetricFaultRate() interface{}
	// Experimental.
	MetricInvocationCount() interface{}
	// Experimental.
	MetricInvocationRate() interface{}
	// Experimental.
	MetricLatencyP50InMillis() interface{}
	// Experimental.
	MetricLatencyP90InMillis() interface{}
	// Experimental.
	MetricLatencyP99InMillis() interface{}
	// Experimental.
	MetricMaxIteratorAgeInMillis() interface{}
	// Experimental.
	MetricProvisionedConcurrencySpilloverInvocations() interface{}
	// Experimental.
	MetricProvisionedConcurrencySpilloverRate() interface{}
	// Experimental.
	MetricThrottlesCount() interface{}
	// Experimental.
	MetricThrottlesRate() interface{}
	// Deprecated: use metricInvocationRate.
	MetricTps() interface{}
}

// The jsii proxy struct for LambdaFunctionMetricFactory
type jsiiProxy_LambdaFunctionMetricFactory struct {
	_ byte // padding
}

func (j *jsiiProxy_LambdaFunctionMetricFactory) FillTpsWithZeroes() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"fillTpsWithZeroes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunctionMetricFactory) LambdaFunction() awslambda.IFunction {
	var returns awslambda.IFunction
	_jsii_.Get(
		j,
		"lambdaFunction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunctionMetricFactory) MetricFactory() MetricFactory {
	var returns MetricFactory
	_jsii_.Get(
		j,
		"metricFactory",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunctionMetricFactory) RateComputationMethod() RateComputationMethod {
	var returns RateComputationMethod
	_jsii_.Get(
		j,
		"rateComputationMethod",
		&returns,
	)
	return returns
}


// Experimental.
func NewLambdaFunctionMetricFactory(metricFactory MetricFactory, props *LambdaFunctionMetricFactoryProps) LambdaFunctionMetricFactory {
	_init_.Initialize()

	if err := validateNewLambdaFunctionMetricFactoryParameters(metricFactory, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_LambdaFunctionMetricFactory{}

	_jsii_.Create(
		"cdk-monitoring-constructs.LambdaFunctionMetricFactory",
		[]interface{}{metricFactory, props},
		&j,
	)

	return &j
}

// Experimental.
func NewLambdaFunctionMetricFactory_Override(l LambdaFunctionMetricFactory, metricFactory MetricFactory, props *LambdaFunctionMetricFactoryProps) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk-monitoring-constructs.LambdaFunctionMetricFactory",
		[]interface{}{metricFactory, props},
		l,
	)
}

func (l *jsiiProxy_LambdaFunctionMetricFactory) MetricConcurrentExecutions() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		l,
		"metricConcurrentExecutions",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LambdaFunctionMetricFactory) MetricFaultCount() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		l,
		"metricFaultCount",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LambdaFunctionMetricFactory) MetricFaultRate() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		l,
		"metricFaultRate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LambdaFunctionMetricFactory) MetricInvocationCount() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		l,
		"metricInvocationCount",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LambdaFunctionMetricFactory) MetricInvocationRate() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		l,
		"metricInvocationRate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LambdaFunctionMetricFactory) MetricLatencyP50InMillis() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		l,
		"metricLatencyP50InMillis",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LambdaFunctionMetricFactory) MetricLatencyP90InMillis() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		l,
		"metricLatencyP90InMillis",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LambdaFunctionMetricFactory) MetricLatencyP99InMillis() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		l,
		"metricLatencyP99InMillis",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LambdaFunctionMetricFactory) MetricMaxIteratorAgeInMillis() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		l,
		"metricMaxIteratorAgeInMillis",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LambdaFunctionMetricFactory) MetricProvisionedConcurrencySpilloverInvocations() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		l,
		"metricProvisionedConcurrencySpilloverInvocations",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LambdaFunctionMetricFactory) MetricProvisionedConcurrencySpilloverRate() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		l,
		"metricProvisionedConcurrencySpilloverRate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LambdaFunctionMetricFactory) MetricThrottlesCount() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		l,
		"metricThrottlesCount",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LambdaFunctionMetricFactory) MetricThrottlesRate() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		l,
		"metricThrottlesRate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LambdaFunctionMetricFactory) MetricTps() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		l,
		"metricTps",
		nil, // no parameters
		&returns,
	)

	return returns
}


package cdkmonitoringconstructs

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/corymhall/cdk-constructs-go/cdkmonitoringconstructs/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
)

// Experimental.
type LambdaFunctionEnhancedMetricFactory interface {
	// Experimental.
	LambdaFunction() awslambda.IFunction
	// Experimental.
	MetricFactory() MetricFactory
	// Experimental.
	EnhancedMetricAvgCpuTotalTime() interface{}
	// Experimental.
	EnhancedMetricAvgMemoryUtilization() interface{}
	// Experimental.
	EnhancedMetricFunctionCost() interface{}
	// Experimental.
	EnhancedMetricMaxCpuTotalTime() interface{}
	// Experimental.
	EnhancedMetricMaxMemoryUtilization() interface{}
	// Experimental.
	EnhancedMetricP90CpuTotalTime() interface{}
	// Experimental.
	EnhancedMetricP90MemoryUtilization() interface{}
}

// The jsii proxy struct for LambdaFunctionEnhancedMetricFactory
type jsiiProxy_LambdaFunctionEnhancedMetricFactory struct {
	_ byte // padding
}

func (j *jsiiProxy_LambdaFunctionEnhancedMetricFactory) LambdaFunction() awslambda.IFunction {
	var returns awslambda.IFunction
	_jsii_.Get(
		j,
		"lambdaFunction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunctionEnhancedMetricFactory) MetricFactory() MetricFactory {
	var returns MetricFactory
	_jsii_.Get(
		j,
		"metricFactory",
		&returns,
	)
	return returns
}


// Experimental.
func NewLambdaFunctionEnhancedMetricFactory(metricFactory MetricFactory, lambdaFunction awslambda.IFunction) LambdaFunctionEnhancedMetricFactory {
	_init_.Initialize()

	if err := validateNewLambdaFunctionEnhancedMetricFactoryParameters(metricFactory, lambdaFunction); err != nil {
		panic(err)
	}
	j := jsiiProxy_LambdaFunctionEnhancedMetricFactory{}

	_jsii_.Create(
		"cdk-monitoring-constructs.LambdaFunctionEnhancedMetricFactory",
		[]interface{}{metricFactory, lambdaFunction},
		&j,
	)

	return &j
}

// Experimental.
func NewLambdaFunctionEnhancedMetricFactory_Override(l LambdaFunctionEnhancedMetricFactory, metricFactory MetricFactory, lambdaFunction awslambda.IFunction) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk-monitoring-constructs.LambdaFunctionEnhancedMetricFactory",
		[]interface{}{metricFactory, lambdaFunction},
		l,
	)
}

func (l *jsiiProxy_LambdaFunctionEnhancedMetricFactory) EnhancedMetricAvgCpuTotalTime() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		l,
		"enhancedMetricAvgCpuTotalTime",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LambdaFunctionEnhancedMetricFactory) EnhancedMetricAvgMemoryUtilization() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		l,
		"enhancedMetricAvgMemoryUtilization",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LambdaFunctionEnhancedMetricFactory) EnhancedMetricFunctionCost() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		l,
		"enhancedMetricFunctionCost",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LambdaFunctionEnhancedMetricFactory) EnhancedMetricMaxCpuTotalTime() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		l,
		"enhancedMetricMaxCpuTotalTime",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LambdaFunctionEnhancedMetricFactory) EnhancedMetricMaxMemoryUtilization() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		l,
		"enhancedMetricMaxMemoryUtilization",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LambdaFunctionEnhancedMetricFactory) EnhancedMetricP90CpuTotalTime() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		l,
		"enhancedMetricP90CpuTotalTime",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LambdaFunctionEnhancedMetricFactory) EnhancedMetricP90MemoryUtilization() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		l,
		"enhancedMetricP90MemoryUtilization",
		nil, // no parameters
		&returns,
	)

	return returns
}


package cdkmonitoringconstructs

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/corymhall/cdk-constructs-go/cdkmonitoringconstructs/jsii"
)

// Experimental.
type StepFunctionMetricFactory interface {
	// Experimental.
	DimensionsMap() *map[string]*string
	// Experimental.
	MetricFactory() MetricFactory
	// Experimental.
	RateComputationMethod() RateComputationMethod
	// Experimental.
	MetricExecutionsAborted() interface{}
	// Experimental.
	MetricExecutionsFailed() interface{}
	// Experimental.
	MetricExecutionsFailedRate() interface{}
	// Experimental.
	MetricExecutionsStarted() interface{}
	// Experimental.
	MetricExecutionsSucceeded() interface{}
	// Experimental.
	MetricExecutionsTimedOut() interface{}
	// Experimental.
	MetricExecutionThrottled() interface{}
	// Experimental.
	MetricExecutionTimeP50InMillis() interface{}
	// Experimental.
	MetricExecutionTimeP90InMillis() interface{}
	// Experimental.
	MetricExecutionTimeP99InMillis() interface{}
}

// The jsii proxy struct for StepFunctionMetricFactory
type jsiiProxy_StepFunctionMetricFactory struct {
	_ byte // padding
}

func (j *jsiiProxy_StepFunctionMetricFactory) DimensionsMap() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"dimensionsMap",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StepFunctionMetricFactory) MetricFactory() MetricFactory {
	var returns MetricFactory
	_jsii_.Get(
		j,
		"metricFactory",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StepFunctionMetricFactory) RateComputationMethod() RateComputationMethod {
	var returns RateComputationMethod
	_jsii_.Get(
		j,
		"rateComputationMethod",
		&returns,
	)
	return returns
}


// Experimental.
func NewStepFunctionMetricFactory(metricFactory MetricFactory, props *StepFunctionMetricFactoryProps) StepFunctionMetricFactory {
	_init_.Initialize()

	if err := validateNewStepFunctionMetricFactoryParameters(metricFactory, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_StepFunctionMetricFactory{}

	_jsii_.Create(
		"cdk-monitoring-constructs.StepFunctionMetricFactory",
		[]interface{}{metricFactory, props},
		&j,
	)

	return &j
}

// Experimental.
func NewStepFunctionMetricFactory_Override(s StepFunctionMetricFactory, metricFactory MetricFactory, props *StepFunctionMetricFactoryProps) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk-monitoring-constructs.StepFunctionMetricFactory",
		[]interface{}{metricFactory, props},
		s,
	)
}

func (s *jsiiProxy_StepFunctionMetricFactory) MetricExecutionsAborted() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"metricExecutionsAborted",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StepFunctionMetricFactory) MetricExecutionsFailed() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"metricExecutionsFailed",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StepFunctionMetricFactory) MetricExecutionsFailedRate() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"metricExecutionsFailedRate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StepFunctionMetricFactory) MetricExecutionsStarted() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"metricExecutionsStarted",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StepFunctionMetricFactory) MetricExecutionsSucceeded() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"metricExecutionsSucceeded",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StepFunctionMetricFactory) MetricExecutionsTimedOut() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"metricExecutionsTimedOut",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StepFunctionMetricFactory) MetricExecutionThrottled() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"metricExecutionThrottled",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StepFunctionMetricFactory) MetricExecutionTimeP50InMillis() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"metricExecutionTimeP50InMillis",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StepFunctionMetricFactory) MetricExecutionTimeP90InMillis() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"metricExecutionTimeP90InMillis",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StepFunctionMetricFactory) MetricExecutionTimeP99InMillis() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"metricExecutionTimeP99InMillis",
		nil, // no parameters
		&returns,
	)

	return returns
}


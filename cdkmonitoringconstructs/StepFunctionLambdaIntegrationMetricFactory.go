package cdkmonitoringconstructs

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/corymhall/cdk-constructs-go/cdkmonitoringconstructs/jsii"
)

// Experimental.
type StepFunctionLambdaIntegrationMetricFactory interface {
	// Experimental.
	DimensionsMap() *map[string]*string
	// Experimental.
	MetricFactory() MetricFactory
	// Experimental.
	RateComputationMethod() RateComputationMethod
	// Experimental.
	MetricFunctionRunTimeP50InMillis() interface{}
	// Experimental.
	MetricFunctionRunTimeP90InMillis() interface{}
	// Experimental.
	MetricFunctionRunTimeP99InMillis() interface{}
	// Experimental.
	MetricFunctionScheduleTimeP50InMillis() interface{}
	// Experimental.
	MetricFunctionScheduleTimeP90InMillis() interface{}
	// Experimental.
	MetricFunctionScheduleTimeP99InMillis() interface{}
	// Experimental.
	MetricFunctionsFailed() interface{}
	// Experimental.
	MetricFunctionsFailedRate() interface{}
	// Experimental.
	MetricFunctionsScheduled() interface{}
	// Experimental.
	MetricFunctionsStarted() interface{}
	// Experimental.
	MetricFunctionsSucceeded() interface{}
	// Experimental.
	MetricFunctionsTimedOut() interface{}
	// Experimental.
	MetricFunctionTimeP50InMillis() interface{}
	// Experimental.
	MetricFunctionTimeP90InMillis() interface{}
	// Experimental.
	MetricFunctionTimeP99InMillis() interface{}
}

// The jsii proxy struct for StepFunctionLambdaIntegrationMetricFactory
type jsiiProxy_StepFunctionLambdaIntegrationMetricFactory struct {
	_ byte // padding
}

func (j *jsiiProxy_StepFunctionLambdaIntegrationMetricFactory) DimensionsMap() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"dimensionsMap",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StepFunctionLambdaIntegrationMetricFactory) MetricFactory() MetricFactory {
	var returns MetricFactory
	_jsii_.Get(
		j,
		"metricFactory",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StepFunctionLambdaIntegrationMetricFactory) RateComputationMethod() RateComputationMethod {
	var returns RateComputationMethod
	_jsii_.Get(
		j,
		"rateComputationMethod",
		&returns,
	)
	return returns
}


// Experimental.
func NewStepFunctionLambdaIntegrationMetricFactory(metricFactory MetricFactory, props *StepFunctionLambdaIntegrationMetricFactoryProps) StepFunctionLambdaIntegrationMetricFactory {
	_init_.Initialize()

	if err := validateNewStepFunctionLambdaIntegrationMetricFactoryParameters(metricFactory, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_StepFunctionLambdaIntegrationMetricFactory{}

	_jsii_.Create(
		"cdk-monitoring-constructs.StepFunctionLambdaIntegrationMetricFactory",
		[]interface{}{metricFactory, props},
		&j,
	)

	return &j
}

// Experimental.
func NewStepFunctionLambdaIntegrationMetricFactory_Override(s StepFunctionLambdaIntegrationMetricFactory, metricFactory MetricFactory, props *StepFunctionLambdaIntegrationMetricFactoryProps) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk-monitoring-constructs.StepFunctionLambdaIntegrationMetricFactory",
		[]interface{}{metricFactory, props},
		s,
	)
}

func (s *jsiiProxy_StepFunctionLambdaIntegrationMetricFactory) MetricFunctionRunTimeP50InMillis() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"metricFunctionRunTimeP50InMillis",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StepFunctionLambdaIntegrationMetricFactory) MetricFunctionRunTimeP90InMillis() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"metricFunctionRunTimeP90InMillis",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StepFunctionLambdaIntegrationMetricFactory) MetricFunctionRunTimeP99InMillis() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"metricFunctionRunTimeP99InMillis",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StepFunctionLambdaIntegrationMetricFactory) MetricFunctionScheduleTimeP50InMillis() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"metricFunctionScheduleTimeP50InMillis",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StepFunctionLambdaIntegrationMetricFactory) MetricFunctionScheduleTimeP90InMillis() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"metricFunctionScheduleTimeP90InMillis",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StepFunctionLambdaIntegrationMetricFactory) MetricFunctionScheduleTimeP99InMillis() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"metricFunctionScheduleTimeP99InMillis",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StepFunctionLambdaIntegrationMetricFactory) MetricFunctionsFailed() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"metricFunctionsFailed",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StepFunctionLambdaIntegrationMetricFactory) MetricFunctionsFailedRate() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"metricFunctionsFailedRate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StepFunctionLambdaIntegrationMetricFactory) MetricFunctionsScheduled() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"metricFunctionsScheduled",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StepFunctionLambdaIntegrationMetricFactory) MetricFunctionsStarted() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"metricFunctionsStarted",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StepFunctionLambdaIntegrationMetricFactory) MetricFunctionsSucceeded() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"metricFunctionsSucceeded",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StepFunctionLambdaIntegrationMetricFactory) MetricFunctionsTimedOut() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"metricFunctionsTimedOut",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StepFunctionLambdaIntegrationMetricFactory) MetricFunctionTimeP50InMillis() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"metricFunctionTimeP50InMillis",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StepFunctionLambdaIntegrationMetricFactory) MetricFunctionTimeP90InMillis() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"metricFunctionTimeP90InMillis",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StepFunctionLambdaIntegrationMetricFactory) MetricFunctionTimeP99InMillis() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"metricFunctionTimeP99InMillis",
		nil, // no parameters
		&returns,
	)

	return returns
}


package cdkmonitoringconstructs

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/corymhall/cdk-constructs-go/cdkmonitoringconstructs/jsii"
)

// Experimental.
type StepFunctionServiceIntegrationMetricFactory interface {
	// Experimental.
	DimensionsMap() *map[string]*string
	// Experimental.
	MetricFactory() MetricFactory
	// Experimental.
	RateComputationMethod() RateComputationMethod
	// Experimental.
	MetricServiceIntegrationRunTimeP50InMillis() interface{}
	// Experimental.
	MetricServiceIntegrationRunTimeP90InMillis() interface{}
	// Experimental.
	MetricServiceIntegrationRunTimeP99InMillis() interface{}
	// Experimental.
	MetricServiceIntegrationScheduleTimeP50InMillis() interface{}
	// Experimental.
	MetricServiceIntegrationScheduleTimeP90InMillis() interface{}
	// Experimental.
	MetricServiceIntegrationScheduleTimeP99InMillis() interface{}
	// Experimental.
	MetricServiceIntegrationsFailed() interface{}
	// Experimental.
	MetricServiceIntegrationsFailedRate() interface{}
	// Experimental.
	MetricServiceIntegrationsScheduled() interface{}
	// Experimental.
	MetricServiceIntegrationsStarted() interface{}
	// Experimental.
	MetricServiceIntegrationsSucceeded() interface{}
	// Experimental.
	MetricServiceIntegrationsTimedOut() interface{}
	// Experimental.
	MetricServiceIntegrationTimeP50InMillis() interface{}
	// Experimental.
	MetricServiceIntegrationTimeP90InMillis() interface{}
	// Experimental.
	MetricServiceIntegrationTimeP99InMillis() interface{}
}

// The jsii proxy struct for StepFunctionServiceIntegrationMetricFactory
type jsiiProxy_StepFunctionServiceIntegrationMetricFactory struct {
	_ byte // padding
}

func (j *jsiiProxy_StepFunctionServiceIntegrationMetricFactory) DimensionsMap() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"dimensionsMap",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StepFunctionServiceIntegrationMetricFactory) MetricFactory() MetricFactory {
	var returns MetricFactory
	_jsii_.Get(
		j,
		"metricFactory",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StepFunctionServiceIntegrationMetricFactory) RateComputationMethod() RateComputationMethod {
	var returns RateComputationMethod
	_jsii_.Get(
		j,
		"rateComputationMethod",
		&returns,
	)
	return returns
}


// Experimental.
func NewStepFunctionServiceIntegrationMetricFactory(metricFactory MetricFactory, props *StepFunctionServiceIntegrationMetricFactoryProps) StepFunctionServiceIntegrationMetricFactory {
	_init_.Initialize()

	if err := validateNewStepFunctionServiceIntegrationMetricFactoryParameters(metricFactory, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_StepFunctionServiceIntegrationMetricFactory{}

	_jsii_.Create(
		"cdk-monitoring-constructs.StepFunctionServiceIntegrationMetricFactory",
		[]interface{}{metricFactory, props},
		&j,
	)

	return &j
}

// Experimental.
func NewStepFunctionServiceIntegrationMetricFactory_Override(s StepFunctionServiceIntegrationMetricFactory, metricFactory MetricFactory, props *StepFunctionServiceIntegrationMetricFactoryProps) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk-monitoring-constructs.StepFunctionServiceIntegrationMetricFactory",
		[]interface{}{metricFactory, props},
		s,
	)
}

func (s *jsiiProxy_StepFunctionServiceIntegrationMetricFactory) MetricServiceIntegrationRunTimeP50InMillis() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"metricServiceIntegrationRunTimeP50InMillis",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StepFunctionServiceIntegrationMetricFactory) MetricServiceIntegrationRunTimeP90InMillis() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"metricServiceIntegrationRunTimeP90InMillis",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StepFunctionServiceIntegrationMetricFactory) MetricServiceIntegrationRunTimeP99InMillis() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"metricServiceIntegrationRunTimeP99InMillis",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StepFunctionServiceIntegrationMetricFactory) MetricServiceIntegrationScheduleTimeP50InMillis() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"metricServiceIntegrationScheduleTimeP50InMillis",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StepFunctionServiceIntegrationMetricFactory) MetricServiceIntegrationScheduleTimeP90InMillis() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"metricServiceIntegrationScheduleTimeP90InMillis",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StepFunctionServiceIntegrationMetricFactory) MetricServiceIntegrationScheduleTimeP99InMillis() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"metricServiceIntegrationScheduleTimeP99InMillis",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StepFunctionServiceIntegrationMetricFactory) MetricServiceIntegrationsFailed() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"metricServiceIntegrationsFailed",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StepFunctionServiceIntegrationMetricFactory) MetricServiceIntegrationsFailedRate() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"metricServiceIntegrationsFailedRate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StepFunctionServiceIntegrationMetricFactory) MetricServiceIntegrationsScheduled() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"metricServiceIntegrationsScheduled",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StepFunctionServiceIntegrationMetricFactory) MetricServiceIntegrationsStarted() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"metricServiceIntegrationsStarted",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StepFunctionServiceIntegrationMetricFactory) MetricServiceIntegrationsSucceeded() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"metricServiceIntegrationsSucceeded",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StepFunctionServiceIntegrationMetricFactory) MetricServiceIntegrationsTimedOut() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"metricServiceIntegrationsTimedOut",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StepFunctionServiceIntegrationMetricFactory) MetricServiceIntegrationTimeP50InMillis() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"metricServiceIntegrationTimeP50InMillis",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StepFunctionServiceIntegrationMetricFactory) MetricServiceIntegrationTimeP90InMillis() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"metricServiceIntegrationTimeP90InMillis",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StepFunctionServiceIntegrationMetricFactory) MetricServiceIntegrationTimeP99InMillis() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"metricServiceIntegrationTimeP99InMillis",
		nil, // no parameters
		&returns,
	)

	return returns
}


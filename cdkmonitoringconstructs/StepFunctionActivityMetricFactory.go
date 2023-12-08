package cdkmonitoringconstructs

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/corymhall/cdk-constructs-go/cdkmonitoringconstructs/jsii"
)

// Experimental.
type StepFunctionActivityMetricFactory interface {
	// Experimental.
	DimensionsMap() *map[string]*string
	// Experimental.
	MetricFactory() MetricFactory
	// Experimental.
	RateComputationMethod() RateComputationMethod
	// Experimental.
	MetricActivitiesFailed() interface{}
	// Experimental.
	MetricActivitiesFailedRate() interface{}
	// Experimental.
	MetricActivitiesHeartbeatTimedOut() interface{}
	// Experimental.
	MetricActivitiesScheduled() interface{}
	// Experimental.
	MetricActivitiesStarted() interface{}
	// Experimental.
	MetricActivitiesSucceeded() interface{}
	// Experimental.
	MetricActivitiesTimedOut() interface{}
	// Experimental.
	MetricActivityRunTimeP50InMillis() interface{}
	// Experimental.
	MetricActivityRunTimeP90InMillis() interface{}
	// Experimental.
	MetricActivityRunTimeP99InMillis() interface{}
	// Experimental.
	MetricActivityScheduleTimeP50InMillis() interface{}
	// Experimental.
	MetricActivityScheduleTimeP90InMillis() interface{}
	// Experimental.
	MetricActivityScheduleTimeP99InMillis() interface{}
	// Experimental.
	MetricActivityTimeP50InMillis() interface{}
	// Experimental.
	MetricActivityTimeP90InMillis() interface{}
	// Experimental.
	MetricActivityTimeP99InMillis() interface{}
}

// The jsii proxy struct for StepFunctionActivityMetricFactory
type jsiiProxy_StepFunctionActivityMetricFactory struct {
	_ byte // padding
}

func (j *jsiiProxy_StepFunctionActivityMetricFactory) DimensionsMap() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"dimensionsMap",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StepFunctionActivityMetricFactory) MetricFactory() MetricFactory {
	var returns MetricFactory
	_jsii_.Get(
		j,
		"metricFactory",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StepFunctionActivityMetricFactory) RateComputationMethod() RateComputationMethod {
	var returns RateComputationMethod
	_jsii_.Get(
		j,
		"rateComputationMethod",
		&returns,
	)
	return returns
}


// Experimental.
func NewStepFunctionActivityMetricFactory(metricFactory MetricFactory, props *StepFunctionActivityMetricFactoryProps) StepFunctionActivityMetricFactory {
	_init_.Initialize()

	if err := validateNewStepFunctionActivityMetricFactoryParameters(metricFactory, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_StepFunctionActivityMetricFactory{}

	_jsii_.Create(
		"cdk-monitoring-constructs.StepFunctionActivityMetricFactory",
		[]interface{}{metricFactory, props},
		&j,
	)

	return &j
}

// Experimental.
func NewStepFunctionActivityMetricFactory_Override(s StepFunctionActivityMetricFactory, metricFactory MetricFactory, props *StepFunctionActivityMetricFactoryProps) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk-monitoring-constructs.StepFunctionActivityMetricFactory",
		[]interface{}{metricFactory, props},
		s,
	)
}

func (s *jsiiProxy_StepFunctionActivityMetricFactory) MetricActivitiesFailed() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"metricActivitiesFailed",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StepFunctionActivityMetricFactory) MetricActivitiesFailedRate() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"metricActivitiesFailedRate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StepFunctionActivityMetricFactory) MetricActivitiesHeartbeatTimedOut() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"metricActivitiesHeartbeatTimedOut",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StepFunctionActivityMetricFactory) MetricActivitiesScheduled() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"metricActivitiesScheduled",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StepFunctionActivityMetricFactory) MetricActivitiesStarted() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"metricActivitiesStarted",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StepFunctionActivityMetricFactory) MetricActivitiesSucceeded() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"metricActivitiesSucceeded",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StepFunctionActivityMetricFactory) MetricActivitiesTimedOut() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"metricActivitiesTimedOut",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StepFunctionActivityMetricFactory) MetricActivityRunTimeP50InMillis() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"metricActivityRunTimeP50InMillis",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StepFunctionActivityMetricFactory) MetricActivityRunTimeP90InMillis() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"metricActivityRunTimeP90InMillis",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StepFunctionActivityMetricFactory) MetricActivityRunTimeP99InMillis() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"metricActivityRunTimeP99InMillis",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StepFunctionActivityMetricFactory) MetricActivityScheduleTimeP50InMillis() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"metricActivityScheduleTimeP50InMillis",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StepFunctionActivityMetricFactory) MetricActivityScheduleTimeP90InMillis() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"metricActivityScheduleTimeP90InMillis",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StepFunctionActivityMetricFactory) MetricActivityScheduleTimeP99InMillis() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"metricActivityScheduleTimeP99InMillis",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StepFunctionActivityMetricFactory) MetricActivityTimeP50InMillis() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"metricActivityTimeP50InMillis",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StepFunctionActivityMetricFactory) MetricActivityTimeP90InMillis() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"metricActivityTimeP90InMillis",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StepFunctionActivityMetricFactory) MetricActivityTimeP99InMillis() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"metricActivityTimeP99InMillis",
		nil, // no parameters
		&returns,
	)

	return returns
}


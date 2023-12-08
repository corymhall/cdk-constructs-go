package cdkmonitoringconstructs

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/corymhall/cdk-constructs-go/cdkmonitoringconstructs/jsii"
)

// Experimental.
type AppSyncMetricFactory interface {
	// Experimental.
	DimensionsMap() *map[string]*string
	// Experimental.
	FillTpsWithZeroes() *bool
	// Experimental.
	MetricFactory() MetricFactory
	// Experimental.
	RateComputationMethod() RateComputationMethod
	// Experimental.
	Metric4XXErrorCount() interface{}
	// Experimental.
	Metric4XXErrorRate() interface{}
	// Experimental.
	Metric5XXFaultCount() interface{}
	// Experimental.
	Metric5XXFaultRate() interface{}
	// Experimental.
	MetricLatencyP50InMillis() interface{}
	// Experimental.
	MetricLatencyP90InMillis() interface{}
	// Experimental.
	MetricLatencyP99InMillis() interface{}
	// Experimental.
	MetricRequestCount() interface{}
	// Experimental.
	MetricRequestRate() interface{}
	// Deprecated: use metricRequestRate.
	MetricTps() interface{}
}

// The jsii proxy struct for AppSyncMetricFactory
type jsiiProxy_AppSyncMetricFactory struct {
	_ byte // padding
}

func (j *jsiiProxy_AppSyncMetricFactory) DimensionsMap() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"dimensionsMap",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSyncMetricFactory) FillTpsWithZeroes() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"fillTpsWithZeroes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSyncMetricFactory) MetricFactory() MetricFactory {
	var returns MetricFactory
	_jsii_.Get(
		j,
		"metricFactory",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSyncMetricFactory) RateComputationMethod() RateComputationMethod {
	var returns RateComputationMethod
	_jsii_.Get(
		j,
		"rateComputationMethod",
		&returns,
	)
	return returns
}


// Experimental.
func NewAppSyncMetricFactory(metricFactory MetricFactory, props *AppSyncMetricFactoryProps) AppSyncMetricFactory {
	_init_.Initialize()

	if err := validateNewAppSyncMetricFactoryParameters(metricFactory, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_AppSyncMetricFactory{}

	_jsii_.Create(
		"cdk-monitoring-constructs.AppSyncMetricFactory",
		[]interface{}{metricFactory, props},
		&j,
	)

	return &j
}

// Experimental.
func NewAppSyncMetricFactory_Override(a AppSyncMetricFactory, metricFactory MetricFactory, props *AppSyncMetricFactoryProps) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk-monitoring-constructs.AppSyncMetricFactory",
		[]interface{}{metricFactory, props},
		a,
	)
}

func (a *jsiiProxy_AppSyncMetricFactory) Metric4XXErrorCount() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"metric4XXErrorCount",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppSyncMetricFactory) Metric4XXErrorRate() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"metric4XXErrorRate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppSyncMetricFactory) Metric5XXFaultCount() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"metric5XXFaultCount",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppSyncMetricFactory) Metric5XXFaultRate() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"metric5XXFaultRate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppSyncMetricFactory) MetricLatencyP50InMillis() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"metricLatencyP50InMillis",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppSyncMetricFactory) MetricLatencyP90InMillis() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"metricLatencyP90InMillis",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppSyncMetricFactory) MetricLatencyP99InMillis() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"metricLatencyP99InMillis",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppSyncMetricFactory) MetricRequestCount() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"metricRequestCount",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppSyncMetricFactory) MetricRequestRate() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"metricRequestRate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppSyncMetricFactory) MetricTps() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"metricTps",
		nil, // no parameters
		&returns,
	)

	return returns
}


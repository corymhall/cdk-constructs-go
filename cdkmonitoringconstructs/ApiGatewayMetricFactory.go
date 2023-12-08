package cdkmonitoringconstructs

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/corymhall/cdk-constructs-go/cdkmonitoringconstructs/jsii"
)

// Experimental.
type ApiGatewayMetricFactory interface {
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
	MetricInvocationCount() interface{}
	// Experimental.
	MetricInvocationRate() interface{}
	// Experimental.
	MetricLatencyInMillis(latencyType LatencyType) interface{}
	// Deprecated: use metricLatencyInMillis instead.
	MetricLatencyP50InMillis() interface{}
	// Deprecated: use metricLatencyInMillis instead.
	MetricLatencyP90InMillis() interface{}
	// Deprecated: use metricLatencyInMillis instead.
	MetricLatencyP99InMillis() interface{}
	// Deprecated: use metricInvocationRate.
	MetricTps() interface{}
}

// The jsii proxy struct for ApiGatewayMetricFactory
type jsiiProxy_ApiGatewayMetricFactory struct {
	_ byte // padding
}

func (j *jsiiProxy_ApiGatewayMetricFactory) DimensionsMap() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"dimensionsMap",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayMetricFactory) FillTpsWithZeroes() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"fillTpsWithZeroes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayMetricFactory) MetricFactory() MetricFactory {
	var returns MetricFactory
	_jsii_.Get(
		j,
		"metricFactory",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayMetricFactory) RateComputationMethod() RateComputationMethod {
	var returns RateComputationMethod
	_jsii_.Get(
		j,
		"rateComputationMethod",
		&returns,
	)
	return returns
}


// Experimental.
func NewApiGatewayMetricFactory(metricFactory MetricFactory, props *ApiGatewayMetricFactoryProps) ApiGatewayMetricFactory {
	_init_.Initialize()

	if err := validateNewApiGatewayMetricFactoryParameters(metricFactory, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_ApiGatewayMetricFactory{}

	_jsii_.Create(
		"cdk-monitoring-constructs.ApiGatewayMetricFactory",
		[]interface{}{metricFactory, props},
		&j,
	)

	return &j
}

// Experimental.
func NewApiGatewayMetricFactory_Override(a ApiGatewayMetricFactory, metricFactory MetricFactory, props *ApiGatewayMetricFactoryProps) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk-monitoring-constructs.ApiGatewayMetricFactory",
		[]interface{}{metricFactory, props},
		a,
	)
}

func (a *jsiiProxy_ApiGatewayMetricFactory) Metric4XXErrorCount() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"metric4XXErrorCount",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiGatewayMetricFactory) Metric4XXErrorRate() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"metric4XXErrorRate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiGatewayMetricFactory) Metric5XXFaultCount() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"metric5XXFaultCount",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiGatewayMetricFactory) Metric5XXFaultRate() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"metric5XXFaultRate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiGatewayMetricFactory) MetricInvocationCount() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"metricInvocationCount",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiGatewayMetricFactory) MetricInvocationRate() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"metricInvocationRate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiGatewayMetricFactory) MetricLatencyInMillis(latencyType LatencyType) interface{} {
	if err := a.validateMetricLatencyInMillisParameters(latencyType); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		a,
		"metricLatencyInMillis",
		[]interface{}{latencyType},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiGatewayMetricFactory) MetricLatencyP50InMillis() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"metricLatencyP50InMillis",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiGatewayMetricFactory) MetricLatencyP90InMillis() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"metricLatencyP90InMillis",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiGatewayMetricFactory) MetricLatencyP99InMillis() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"metricLatencyP99InMillis",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiGatewayMetricFactory) MetricTps() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"metricTps",
		nil, // no parameters
		&returns,
	)

	return returns
}


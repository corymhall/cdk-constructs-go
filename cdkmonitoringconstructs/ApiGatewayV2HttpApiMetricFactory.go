package cdkmonitoringconstructs

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/corymhall/cdk-constructs-go/cdkmonitoringconstructs/jsii"
)

// Experimental.
type ApiGatewayV2HttpApiMetricFactory interface {
	// Experimental.
	DimensionsMap() *map[string]*string
	// Experimental.
	FillTpsWithZeroes() *bool
	// Experimental.
	MetricFactory() MetricFactory
	// Experimental.
	RateComputationMethod() RateComputationMethod
	// Experimental.
	Metric4xxCount() interface{}
	// Experimental.
	Metric4xxRate() interface{}
	// Experimental.
	Metric5xxCount() interface{}
	// Experimental.
	Metric5xxRate() interface{}
	// Experimental.
	MetricIntegrationLatencyInMillis(latencyType LatencyType) interface{}
	// Deprecated: use metricIntegrationLatencyInMillis instead.
	MetricIntegrationLatencyP50InMillis() interface{}
	// Deprecated: use metricIntegrationLatencyInMillis instead.
	MetricIntegrationLatencyP90InMillis() interface{}
	// Deprecated: use metricIntegrationLatencyInMillis instead.
	MetricIntegrationLatencyP99InMillis() interface{}
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

// The jsii proxy struct for ApiGatewayV2HttpApiMetricFactory
type jsiiProxy_ApiGatewayV2HttpApiMetricFactory struct {
	_ byte // padding
}

func (j *jsiiProxy_ApiGatewayV2HttpApiMetricFactory) DimensionsMap() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"dimensionsMap",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayV2HttpApiMetricFactory) FillTpsWithZeroes() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"fillTpsWithZeroes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayV2HttpApiMetricFactory) MetricFactory() MetricFactory {
	var returns MetricFactory
	_jsii_.Get(
		j,
		"metricFactory",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayV2HttpApiMetricFactory) RateComputationMethod() RateComputationMethod {
	var returns RateComputationMethod
	_jsii_.Get(
		j,
		"rateComputationMethod",
		&returns,
	)
	return returns
}


// Experimental.
func NewApiGatewayV2HttpApiMetricFactory(metricFactory MetricFactory, props *ApiGatewayV2HttpApiMetricFactoryProps) ApiGatewayV2HttpApiMetricFactory {
	_init_.Initialize()

	if err := validateNewApiGatewayV2HttpApiMetricFactoryParameters(metricFactory, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_ApiGatewayV2HttpApiMetricFactory{}

	_jsii_.Create(
		"cdk-monitoring-constructs.ApiGatewayV2HttpApiMetricFactory",
		[]interface{}{metricFactory, props},
		&j,
	)

	return &j
}

// Experimental.
func NewApiGatewayV2HttpApiMetricFactory_Override(a ApiGatewayV2HttpApiMetricFactory, metricFactory MetricFactory, props *ApiGatewayV2HttpApiMetricFactoryProps) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk-monitoring-constructs.ApiGatewayV2HttpApiMetricFactory",
		[]interface{}{metricFactory, props},
		a,
	)
}

func (a *jsiiProxy_ApiGatewayV2HttpApiMetricFactory) Metric4xxCount() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"metric4xxCount",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiGatewayV2HttpApiMetricFactory) Metric4xxRate() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"metric4xxRate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiGatewayV2HttpApiMetricFactory) Metric5xxCount() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"metric5xxCount",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiGatewayV2HttpApiMetricFactory) Metric5xxRate() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"metric5xxRate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiGatewayV2HttpApiMetricFactory) MetricIntegrationLatencyInMillis(latencyType LatencyType) interface{} {
	if err := a.validateMetricIntegrationLatencyInMillisParameters(latencyType); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		a,
		"metricIntegrationLatencyInMillis",
		[]interface{}{latencyType},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiGatewayV2HttpApiMetricFactory) MetricIntegrationLatencyP50InMillis() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"metricIntegrationLatencyP50InMillis",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiGatewayV2HttpApiMetricFactory) MetricIntegrationLatencyP90InMillis() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"metricIntegrationLatencyP90InMillis",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiGatewayV2HttpApiMetricFactory) MetricIntegrationLatencyP99InMillis() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"metricIntegrationLatencyP99InMillis",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiGatewayV2HttpApiMetricFactory) MetricInvocationCount() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"metricInvocationCount",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiGatewayV2HttpApiMetricFactory) MetricInvocationRate() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"metricInvocationRate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiGatewayV2HttpApiMetricFactory) MetricLatencyInMillis(latencyType LatencyType) interface{} {
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

func (a *jsiiProxy_ApiGatewayV2HttpApiMetricFactory) MetricLatencyP50InMillis() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"metricLatencyP50InMillis",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiGatewayV2HttpApiMetricFactory) MetricLatencyP90InMillis() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"metricLatencyP90InMillis",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiGatewayV2HttpApiMetricFactory) MetricLatencyP99InMillis() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"metricLatencyP99InMillis",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiGatewayV2HttpApiMetricFactory) MetricTps() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"metricTps",
		nil, // no parameters
		&returns,
	)

	return returns
}


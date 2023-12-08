package cdkmonitoringconstructs

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/corymhall/cdk-constructs-go/cdkmonitoringconstructs/jsii"
)

// https://docs.aws.amazon.com/waf/latest/developerguide/monitoring-cloudwatch.html.
// Experimental.
type WafV2MetricFactory interface {
	// Experimental.
	Dimensions() *map[string]interface{}
	// Experimental.
	MetricFactory() MetricFactory
	// Experimental.
	MetricAllowedRequests() interface{}
	// Experimental.
	MetricBlockedRequests() interface{}
	// Experimental.
	MetricBlockedRequestsRate() interface{}
}

// The jsii proxy struct for WafV2MetricFactory
type jsiiProxy_WafV2MetricFactory struct {
	_ byte // padding
}

func (j *jsiiProxy_WafV2MetricFactory) Dimensions() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"dimensions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WafV2MetricFactory) MetricFactory() MetricFactory {
	var returns MetricFactory
	_jsii_.Get(
		j,
		"metricFactory",
		&returns,
	)
	return returns
}


// Experimental.
func NewWafV2MetricFactory(metricFactory MetricFactory, props *WafV2MetricFactoryProps) WafV2MetricFactory {
	_init_.Initialize()

	if err := validateNewWafV2MetricFactoryParameters(metricFactory, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_WafV2MetricFactory{}

	_jsii_.Create(
		"cdk-monitoring-constructs.WafV2MetricFactory",
		[]interface{}{metricFactory, props},
		&j,
	)

	return &j
}

// Experimental.
func NewWafV2MetricFactory_Override(w WafV2MetricFactory, metricFactory MetricFactory, props *WafV2MetricFactoryProps) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk-monitoring-constructs.WafV2MetricFactory",
		[]interface{}{metricFactory, props},
		w,
	)
}

func (w *jsiiProxy_WafV2MetricFactory) MetricAllowedRequests() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		w,
		"metricAllowedRequests",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WafV2MetricFactory) MetricBlockedRequests() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		w,
		"metricBlockedRequests",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WafV2MetricFactory) MetricBlockedRequestsRate() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		w,
		"metricBlockedRequestsRate",
		nil, // no parameters
		&returns,
	)

	return returns
}


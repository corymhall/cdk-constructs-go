package cdkmonitoringconstructs

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Common interface for load-balancer based service metric factories.
// Experimental.
type ILoadBalancerMetricFactory interface {
	// Experimental.
	MetricActiveConnectionCount() interface{}
	// Experimental.
	MetricHealthyTaskCount() interface{}
	// Experimental.
	MetricHealthyTaskInPercent() interface{}
	// Experimental.
	MetricNewConnectionCount() interface{}
	// Experimental.
	MetricProcessedBytesMin() interface{}
	// Experimental.
	MetricUnhealthyTaskCount() interface{}
}

// The jsii proxy for ILoadBalancerMetricFactory
type jsiiProxy_ILoadBalancerMetricFactory struct {
	_ byte // padding
}

func (i *jsiiProxy_ILoadBalancerMetricFactory) MetricActiveConnectionCount() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		i,
		"metricActiveConnectionCount",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_ILoadBalancerMetricFactory) MetricHealthyTaskCount() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		i,
		"metricHealthyTaskCount",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_ILoadBalancerMetricFactory) MetricHealthyTaskInPercent() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		i,
		"metricHealthyTaskInPercent",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_ILoadBalancerMetricFactory) MetricNewConnectionCount() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		i,
		"metricNewConnectionCount",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_ILoadBalancerMetricFactory) MetricProcessedBytesMin() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		i,
		"metricProcessedBytesMin",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_ILoadBalancerMetricFactory) MetricUnhealthyTaskCount() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		i,
		"metricUnhealthyTaskCount",
		nil, // no parameters
		&returns,
	)

	return returns
}


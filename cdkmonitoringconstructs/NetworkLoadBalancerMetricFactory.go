package cdkmonitoringconstructs

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/corymhall/cdk-constructs-go/cdkmonitoringconstructs/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2/awselasticloadbalancingv2"
)

// Metric factory to create metrics for network load-balanced service.
// Experimental.
type NetworkLoadBalancerMetricFactory interface {
	ILoadBalancerMetricFactory
	// Experimental.
	InvertStatisticsOfTaskCountEnabled() *bool
	// Experimental.
	MetricFactory() MetricFactory
	// Experimental.
	NetworkLoadBalancer() awselasticloadbalancingv2.INetworkLoadBalancer
	// Experimental.
	NetworkTargetGroup() awselasticloadbalancingv2.INetworkTargetGroup
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

// The jsii proxy struct for NetworkLoadBalancerMetricFactory
type jsiiProxy_NetworkLoadBalancerMetricFactory struct {
	jsiiProxy_ILoadBalancerMetricFactory
}

func (j *jsiiProxy_NetworkLoadBalancerMetricFactory) InvertStatisticsOfTaskCountEnabled() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"invertStatisticsOfTaskCountEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkLoadBalancerMetricFactory) MetricFactory() MetricFactory {
	var returns MetricFactory
	_jsii_.Get(
		j,
		"metricFactory",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkLoadBalancerMetricFactory) NetworkLoadBalancer() awselasticloadbalancingv2.INetworkLoadBalancer {
	var returns awselasticloadbalancingv2.INetworkLoadBalancer
	_jsii_.Get(
		j,
		"networkLoadBalancer",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkLoadBalancerMetricFactory) NetworkTargetGroup() awselasticloadbalancingv2.INetworkTargetGroup {
	var returns awselasticloadbalancingv2.INetworkTargetGroup
	_jsii_.Get(
		j,
		"networkTargetGroup",
		&returns,
	)
	return returns
}


// Experimental.
func NewNetworkLoadBalancerMetricFactory(metricFactory MetricFactory, props *NetworkLoadBalancerMetricFactoryProps) NetworkLoadBalancerMetricFactory {
	_init_.Initialize()

	if err := validateNewNetworkLoadBalancerMetricFactoryParameters(metricFactory, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_NetworkLoadBalancerMetricFactory{}

	_jsii_.Create(
		"cdk-monitoring-constructs.NetworkLoadBalancerMetricFactory",
		[]interface{}{metricFactory, props},
		&j,
	)

	return &j
}

// Experimental.
func NewNetworkLoadBalancerMetricFactory_Override(n NetworkLoadBalancerMetricFactory, metricFactory MetricFactory, props *NetworkLoadBalancerMetricFactoryProps) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk-monitoring-constructs.NetworkLoadBalancerMetricFactory",
		[]interface{}{metricFactory, props},
		n,
	)
}

func (n *jsiiProxy_NetworkLoadBalancerMetricFactory) MetricActiveConnectionCount() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		n,
		"metricActiveConnectionCount",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkLoadBalancerMetricFactory) MetricHealthyTaskCount() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		n,
		"metricHealthyTaskCount",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkLoadBalancerMetricFactory) MetricHealthyTaskInPercent() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		n,
		"metricHealthyTaskInPercent",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkLoadBalancerMetricFactory) MetricNewConnectionCount() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		n,
		"metricNewConnectionCount",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkLoadBalancerMetricFactory) MetricProcessedBytesMin() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		n,
		"metricProcessedBytesMin",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkLoadBalancerMetricFactory) MetricUnhealthyTaskCount() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		n,
		"metricUnhealthyTaskCount",
		nil, // no parameters
		&returns,
	)

	return returns
}


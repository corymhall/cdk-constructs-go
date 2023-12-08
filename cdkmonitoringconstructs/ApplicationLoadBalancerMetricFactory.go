package cdkmonitoringconstructs

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/corymhall/cdk-constructs-go/cdkmonitoringconstructs/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2/awselasticloadbalancingv2"
)

// Metric factory to create metrics for application load-balanced service.
// Experimental.
type ApplicationLoadBalancerMetricFactory interface {
	ILoadBalancerMetricFactory
	// Experimental.
	ApplicationLoadBalancer() awselasticloadbalancingv2.IApplicationLoadBalancer
	// Experimental.
	ApplicationTargetGroup() awselasticloadbalancingv2.IApplicationTargetGroup
	// Experimental.
	InvertStatisticsOfTaskCountEnabled() *bool
	// Experimental.
	MetricFactory() MetricFactory
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

// The jsii proxy struct for ApplicationLoadBalancerMetricFactory
type jsiiProxy_ApplicationLoadBalancerMetricFactory struct {
	jsiiProxy_ILoadBalancerMetricFactory
}

func (j *jsiiProxy_ApplicationLoadBalancerMetricFactory) ApplicationLoadBalancer() awselasticloadbalancingv2.IApplicationLoadBalancer {
	var returns awselasticloadbalancingv2.IApplicationLoadBalancer
	_jsii_.Get(
		j,
		"applicationLoadBalancer",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationLoadBalancerMetricFactory) ApplicationTargetGroup() awselasticloadbalancingv2.IApplicationTargetGroup {
	var returns awselasticloadbalancingv2.IApplicationTargetGroup
	_jsii_.Get(
		j,
		"applicationTargetGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationLoadBalancerMetricFactory) InvertStatisticsOfTaskCountEnabled() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"invertStatisticsOfTaskCountEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationLoadBalancerMetricFactory) MetricFactory() MetricFactory {
	var returns MetricFactory
	_jsii_.Get(
		j,
		"metricFactory",
		&returns,
	)
	return returns
}


// Experimental.
func NewApplicationLoadBalancerMetricFactory(metricFactory MetricFactory, props *ApplicationLoadBalancerMetricFactoryProps) ApplicationLoadBalancerMetricFactory {
	_init_.Initialize()

	if err := validateNewApplicationLoadBalancerMetricFactoryParameters(metricFactory, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_ApplicationLoadBalancerMetricFactory{}

	_jsii_.Create(
		"cdk-monitoring-constructs.ApplicationLoadBalancerMetricFactory",
		[]interface{}{metricFactory, props},
		&j,
	)

	return &j
}

// Experimental.
func NewApplicationLoadBalancerMetricFactory_Override(a ApplicationLoadBalancerMetricFactory, metricFactory MetricFactory, props *ApplicationLoadBalancerMetricFactoryProps) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk-monitoring-constructs.ApplicationLoadBalancerMetricFactory",
		[]interface{}{metricFactory, props},
		a,
	)
}

func (a *jsiiProxy_ApplicationLoadBalancerMetricFactory) MetricActiveConnectionCount() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"metricActiveConnectionCount",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApplicationLoadBalancerMetricFactory) MetricHealthyTaskCount() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"metricHealthyTaskCount",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApplicationLoadBalancerMetricFactory) MetricHealthyTaskInPercent() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"metricHealthyTaskInPercent",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApplicationLoadBalancerMetricFactory) MetricNewConnectionCount() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"metricNewConnectionCount",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApplicationLoadBalancerMetricFactory) MetricProcessedBytesMin() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"metricProcessedBytesMin",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApplicationLoadBalancerMetricFactory) MetricUnhealthyTaskCount() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"metricUnhealthyTaskCount",
		nil, // no parameters
		&returns,
	)

	return returns
}


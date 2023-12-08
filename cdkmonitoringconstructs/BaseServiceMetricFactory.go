package cdkmonitoringconstructs

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/corymhall/cdk-constructs-go/cdkmonitoringconstructs/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsecs"
)

// Metric factory for a base service (parent class for e.g. Fargate and EC2 services).
// Experimental.
type BaseServiceMetricFactory interface {
	// Experimental.
	MetricFactory() MetricFactory
	// Experimental.
	Service() awsecs.BaseService
	// Experimental.
	MetricClusterCpuUtilisationInPercent() interface{}
	// Experimental.
	MetricClusterMemoryUtilisationInPercent() interface{}
	// Experimental.
	MetricRunningTaskCount() interface{}
}

// The jsii proxy struct for BaseServiceMetricFactory
type jsiiProxy_BaseServiceMetricFactory struct {
	_ byte // padding
}

func (j *jsiiProxy_BaseServiceMetricFactory) MetricFactory() MetricFactory {
	var returns MetricFactory
	_jsii_.Get(
		j,
		"metricFactory",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BaseServiceMetricFactory) Service() awsecs.BaseService {
	var returns awsecs.BaseService
	_jsii_.Get(
		j,
		"service",
		&returns,
	)
	return returns
}


// Experimental.
func NewBaseServiceMetricFactory(metricFactory MetricFactory, props *BaseServiceMetricFactoryProps) BaseServiceMetricFactory {
	_init_.Initialize()

	if err := validateNewBaseServiceMetricFactoryParameters(metricFactory, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_BaseServiceMetricFactory{}

	_jsii_.Create(
		"cdk-monitoring-constructs.BaseServiceMetricFactory",
		[]interface{}{metricFactory, props},
		&j,
	)

	return &j
}

// Experimental.
func NewBaseServiceMetricFactory_Override(b BaseServiceMetricFactory, metricFactory MetricFactory, props *BaseServiceMetricFactoryProps) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk-monitoring-constructs.BaseServiceMetricFactory",
		[]interface{}{metricFactory, props},
		b,
	)
}

func (b *jsiiProxy_BaseServiceMetricFactory) MetricClusterCpuUtilisationInPercent() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		b,
		"metricClusterCpuUtilisationInPercent",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BaseServiceMetricFactory) MetricClusterMemoryUtilisationInPercent() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		b,
		"metricClusterMemoryUtilisationInPercent",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BaseServiceMetricFactory) MetricRunningTaskCount() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		b,
		"metricRunningTaskCount",
		nil, // no parameters
		&returns,
	)

	return returns
}


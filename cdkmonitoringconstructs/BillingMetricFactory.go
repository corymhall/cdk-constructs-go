package cdkmonitoringconstructs

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/corymhall/cdk-constructs-go/cdkmonitoringconstructs/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudwatch"
)

// Experimental.
type BillingMetricFactory interface {
	// Experimental.
	MetricSearchTopCostByServiceInUsd() awscloudwatch.IMetric
	// Experimental.
	MetricTotalCostInUsd() interface{}
}

// The jsii proxy struct for BillingMetricFactory
type jsiiProxy_BillingMetricFactory struct {
	_ byte // padding
}

// Experimental.
func NewBillingMetricFactory() BillingMetricFactory {
	_init_.Initialize()

	j := jsiiProxy_BillingMetricFactory{}

	_jsii_.Create(
		"cdk-monitoring-constructs.BillingMetricFactory",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewBillingMetricFactory_Override(b BillingMetricFactory) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk-monitoring-constructs.BillingMetricFactory",
		nil, // no parameters
		b,
	)
}

func (b *jsiiProxy_BillingMetricFactory) MetricSearchTopCostByServiceInUsd() awscloudwatch.IMetric {
	var returns awscloudwatch.IMetric

	_jsii_.Invoke(
		b,
		"metricSearchTopCostByServiceInUsd",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BillingMetricFactory) MetricTotalCostInUsd() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		b,
		"metricTotalCostInUsd",
		nil, // no parameters
		&returns,
	)

	return returns
}


package cdkmonitoringconstructs

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/corymhall/cdk-constructs-go/cdkmonitoringconstructs/jsii"
)

// Experimental.
type CloudWatchLogsMetricFactory interface {
	// Experimental.
	MetricIncomingLogEvents() interface{}
}

// The jsii proxy struct for CloudWatchLogsMetricFactory
type jsiiProxy_CloudWatchLogsMetricFactory struct {
	_ byte // padding
}

// Experimental.
func NewCloudWatchLogsMetricFactory(metricFactory MetricFactory, props *CloudWatchLogsMetricFactoryProps) CloudWatchLogsMetricFactory {
	_init_.Initialize()

	if err := validateNewCloudWatchLogsMetricFactoryParameters(metricFactory, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_CloudWatchLogsMetricFactory{}

	_jsii_.Create(
		"cdk-monitoring-constructs.CloudWatchLogsMetricFactory",
		[]interface{}{metricFactory, props},
		&j,
	)

	return &j
}

// Experimental.
func NewCloudWatchLogsMetricFactory_Override(c CloudWatchLogsMetricFactory, metricFactory MetricFactory, props *CloudWatchLogsMetricFactoryProps) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk-monitoring-constructs.CloudWatchLogsMetricFactory",
		[]interface{}{metricFactory, props},
		c,
	)
}

func (c *jsiiProxy_CloudWatchLogsMetricFactory) MetricIncomingLogEvents() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"metricIncomingLogEvents",
		nil, // no parameters
		&returns,
	)

	return returns
}


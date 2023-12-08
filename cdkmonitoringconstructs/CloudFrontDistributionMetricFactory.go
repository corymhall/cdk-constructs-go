package cdkmonitoringconstructs

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/corymhall/cdk-constructs-go/cdkmonitoringconstructs/jsii"
)

// To get the CloudFront metrics from the CloudWatch API, you must use the US East (N.
//
// Virginia) Region (us-east-1).
// https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/programming-cloudwatch-metrics.html
// Experimental.
type CloudFrontDistributionMetricFactory interface {
	// Experimental.
	Metric4xxErrorRateAverage() interface{}
	// Experimental.
	Metric5xxErrorRateAverage() interface{}
	// Cache hit rate metric.
	//
	// This is an additional metric that needs to be explicitly enabled for an additional cost.
	// See: https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/viewing-cloudfront-metrics.html#monitoring-console.distributions-additional
	//
	// Experimental.
	MetricCacheHitRateAverageInPercent() interface{}
	// Experimental.
	MetricRequestCount() interface{}
	// Experimental.
	MetricRequestRate() interface{}
	// Deprecated: use metricRequestRate.
	MetricRequestTps() interface{}
	// Experimental.
	MetricTotalBytesDownloaded() interface{}
	// Experimental.
	MetricTotalBytesUploaded() interface{}
	// Experimental.
	MetricTotalErrorRateAverage() interface{}
}

// The jsii proxy struct for CloudFrontDistributionMetricFactory
type jsiiProxy_CloudFrontDistributionMetricFactory struct {
	_ byte // padding
}

// Experimental.
func NewCloudFrontDistributionMetricFactory(metricFactory MetricFactory, props *CloudFrontDistributionMetricFactoryProps) CloudFrontDistributionMetricFactory {
	_init_.Initialize()

	if err := validateNewCloudFrontDistributionMetricFactoryParameters(metricFactory, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_CloudFrontDistributionMetricFactory{}

	_jsii_.Create(
		"cdk-monitoring-constructs.CloudFrontDistributionMetricFactory",
		[]interface{}{metricFactory, props},
		&j,
	)

	return &j
}

// Experimental.
func NewCloudFrontDistributionMetricFactory_Override(c CloudFrontDistributionMetricFactory, metricFactory MetricFactory, props *CloudFrontDistributionMetricFactoryProps) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk-monitoring-constructs.CloudFrontDistributionMetricFactory",
		[]interface{}{metricFactory, props},
		c,
	)
}

func (c *jsiiProxy_CloudFrontDistributionMetricFactory) Metric4xxErrorRateAverage() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"metric4xxErrorRateAverage",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudFrontDistributionMetricFactory) Metric5xxErrorRateAverage() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"metric5xxErrorRateAverage",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudFrontDistributionMetricFactory) MetricCacheHitRateAverageInPercent() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"metricCacheHitRateAverageInPercent",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudFrontDistributionMetricFactory) MetricRequestCount() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"metricRequestCount",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudFrontDistributionMetricFactory) MetricRequestRate() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"metricRequestRate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudFrontDistributionMetricFactory) MetricRequestTps() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"metricRequestTps",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudFrontDistributionMetricFactory) MetricTotalBytesDownloaded() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"metricTotalBytesDownloaded",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudFrontDistributionMetricFactory) MetricTotalBytesUploaded() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"metricTotalBytesUploaded",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudFrontDistributionMetricFactory) MetricTotalErrorRateAverage() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"metricTotalErrorRateAverage",
		nil, // no parameters
		&returns,
	)

	return returns
}


package cdkmonitoringconstructs

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/corymhall/cdk-constructs-go/cdkmonitoringconstructs/jsii"
)

// Experimental.
type S3BucketMetricFactory interface {
	// Experimental.
	MetricFactory() MetricFactory
	// Experimental.
	Props() *S3BucketMetricFactoryProps
	// Experimental.
	MetricBucketSizeBytes() interface{}
	// Experimental.
	MetricNumberOfObjects() interface{}
}

// The jsii proxy struct for S3BucketMetricFactory
type jsiiProxy_S3BucketMetricFactory struct {
	_ byte // padding
}

func (j *jsiiProxy_S3BucketMetricFactory) MetricFactory() MetricFactory {
	var returns MetricFactory
	_jsii_.Get(
		j,
		"metricFactory",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3BucketMetricFactory) Props() *S3BucketMetricFactoryProps {
	var returns *S3BucketMetricFactoryProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}


// Experimental.
func NewS3BucketMetricFactory(metricFactory MetricFactory, props *S3BucketMetricFactoryProps) S3BucketMetricFactory {
	_init_.Initialize()

	if err := validateNewS3BucketMetricFactoryParameters(metricFactory, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_S3BucketMetricFactory{}

	_jsii_.Create(
		"cdk-monitoring-constructs.S3BucketMetricFactory",
		[]interface{}{metricFactory, props},
		&j,
	)

	return &j
}

// Experimental.
func NewS3BucketMetricFactory_Override(s S3BucketMetricFactory, metricFactory MetricFactory, props *S3BucketMetricFactoryProps) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk-monitoring-constructs.S3BucketMetricFactory",
		[]interface{}{metricFactory, props},
		s,
	)
}

func (s *jsiiProxy_S3BucketMetricFactory) MetricBucketSizeBytes() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"metricBucketSizeBytes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_S3BucketMetricFactory) MetricNumberOfObjects() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"metricNumberOfObjects",
		nil, // no parameters
		&returns,
	)

	return returns
}


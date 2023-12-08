package cdkmonitoringconstructs

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/corymhall/cdk-constructs-go/cdkmonitoringconstructs/jsii"
)

// See: https://docs.aws.amazon.com/firehose/latest/dev/monitoring-with-cloudwatch-metrics.html
//
// Experimental.
type KinesisFirehoseMetricFactory interface {
	// Experimental.
	DimensionsMap() *map[string]*string
	// Experimental.
	MetricFactory() MetricFactory
	// Experimental.
	MetricBytesPerSecondLimit() interface{}
	// Experimental.
	MetricFailedConversionCount() interface{}
	// Experimental.
	MetricIncomingBytes() interface{}
	// Experimental.
	MetricIncomingBytesToLimitRate() interface{}
	// Experimental.
	MetricIncomingPutRequests() interface{}
	// Experimental.
	MetricIncomingPutRequestsToLimitRate() interface{}
	// Experimental.
	MetricIncomingRecordCount() interface{}
	// Experimental.
	MetricIncomingRecordsToLimitRate() interface{}
	// Experimental.
	MetricPutRecordBatchLatencyP90InMillis() interface{}
	// Experimental.
	MetricPutRecordLatencyP90InMillis() interface{}
	// Experimental.
	MetricPutRequestsPerSecondLimit() interface{}
	// Experimental.
	MetricRecordsPerSecondLimit() interface{}
	// Experimental.
	MetricSuccessfulConversionCount() interface{}
	// Experimental.
	MetricThrottledRecordCount() interface{}
}

// The jsii proxy struct for KinesisFirehoseMetricFactory
type jsiiProxy_KinesisFirehoseMetricFactory struct {
	_ byte // padding
}

func (j *jsiiProxy_KinesisFirehoseMetricFactory) DimensionsMap() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"dimensionsMap",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisFirehoseMetricFactory) MetricFactory() MetricFactory {
	var returns MetricFactory
	_jsii_.Get(
		j,
		"metricFactory",
		&returns,
	)
	return returns
}


// Experimental.
func NewKinesisFirehoseMetricFactory(metricFactory MetricFactory, props *KinesisFirehoseMetricFactoryProps) KinesisFirehoseMetricFactory {
	_init_.Initialize()

	if err := validateNewKinesisFirehoseMetricFactoryParameters(metricFactory, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_KinesisFirehoseMetricFactory{}

	_jsii_.Create(
		"cdk-monitoring-constructs.KinesisFirehoseMetricFactory",
		[]interface{}{metricFactory, props},
		&j,
	)

	return &j
}

// Experimental.
func NewKinesisFirehoseMetricFactory_Override(k KinesisFirehoseMetricFactory, metricFactory MetricFactory, props *KinesisFirehoseMetricFactoryProps) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk-monitoring-constructs.KinesisFirehoseMetricFactory",
		[]interface{}{metricFactory, props},
		k,
	)
}

func (k *jsiiProxy_KinesisFirehoseMetricFactory) MetricBytesPerSecondLimit() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		k,
		"metricBytesPerSecondLimit",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KinesisFirehoseMetricFactory) MetricFailedConversionCount() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		k,
		"metricFailedConversionCount",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KinesisFirehoseMetricFactory) MetricIncomingBytes() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		k,
		"metricIncomingBytes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KinesisFirehoseMetricFactory) MetricIncomingBytesToLimitRate() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		k,
		"metricIncomingBytesToLimitRate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KinesisFirehoseMetricFactory) MetricIncomingPutRequests() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		k,
		"metricIncomingPutRequests",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KinesisFirehoseMetricFactory) MetricIncomingPutRequestsToLimitRate() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		k,
		"metricIncomingPutRequestsToLimitRate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KinesisFirehoseMetricFactory) MetricIncomingRecordCount() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		k,
		"metricIncomingRecordCount",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KinesisFirehoseMetricFactory) MetricIncomingRecordsToLimitRate() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		k,
		"metricIncomingRecordsToLimitRate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KinesisFirehoseMetricFactory) MetricPutRecordBatchLatencyP90InMillis() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		k,
		"metricPutRecordBatchLatencyP90InMillis",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KinesisFirehoseMetricFactory) MetricPutRecordLatencyP90InMillis() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		k,
		"metricPutRecordLatencyP90InMillis",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KinesisFirehoseMetricFactory) MetricPutRequestsPerSecondLimit() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		k,
		"metricPutRequestsPerSecondLimit",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KinesisFirehoseMetricFactory) MetricRecordsPerSecondLimit() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		k,
		"metricRecordsPerSecondLimit",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KinesisFirehoseMetricFactory) MetricSuccessfulConversionCount() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		k,
		"metricSuccessfulConversionCount",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KinesisFirehoseMetricFactory) MetricThrottledRecordCount() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		k,
		"metricThrottledRecordCount",
		nil, // no parameters
		&returns,
	)

	return returns
}


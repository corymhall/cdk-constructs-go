package cdkmonitoringconstructs

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/corymhall/cdk-constructs-go/cdkmonitoringconstructs/jsii"
)

// See: https://docs.aws.amazon.com/streams/latest/dev/monitoring-with-cloudwatch.html
//
// Experimental.
type KinesisDataStreamMetricFactory interface {
	// Experimental.
	DimensionsMap() *map[string]*string
	// Experimental.
	MetricFactory() MetricFactory
	// Experimental.
	MetricGetRecordsIteratorAgeMaxMs() interface{}
	// Experimental.
	MetricGetRecordsLatencyAverageMs() interface{}
	// Experimental.
	MetricGetRecordsSuccessCount() interface{}
	// Experimental.
	MetricGetRecordsSumBytes() interface{}
	// Experimental.
	MetricGetRecordsSumCount() interface{}
	// Experimental.
	MetricIncomingDataSumBytes() interface{}
	// Experimental.
	MetricIncomingDataSumCount() interface{}
	// Experimental.
	MetricPutRecordLatencyAverageMs() interface{}
	// Experimental.
	MetricPutRecordsFailedRecordsCount() interface{}
	// Experimental.
	MetricPutRecordsLatencyAverageMs() interface{}
	// Experimental.
	MetricPutRecordsSuccessCount() interface{}
	// Experimental.
	MetricPutRecordsSuccessfulRecordsCount() interface{}
	// Experimental.
	MetricPutRecordsSumBytes() interface{}
	// Experimental.
	MetricPutRecordsThrottledRecordsCount() interface{}
	// Experimental.
	MetricPutRecordsTotalRecordsCount() interface{}
	// Experimental.
	MetricPutRecordSuccessCount() interface{}
	// Experimental.
	MetricPutRecordSumBytes() interface{}
	// Experimental.
	MetricReadProvisionedThroughputExceeded() interface{}
	// Deprecated: please use `metricReadProvisionedThroughputExceeded` instead.
	MetricReadProvisionedThroughputExceededPercent() interface{}
	// Experimental.
	MetricWriteProvisionedThroughputExceeded() interface{}
	// Deprecated: please use `metricWriteProvisionedThroughputExceeded` instead.
	MetricWriteProvisionedThroughputExceededPercent() interface{}
}

// The jsii proxy struct for KinesisDataStreamMetricFactory
type jsiiProxy_KinesisDataStreamMetricFactory struct {
	_ byte // padding
}

func (j *jsiiProxy_KinesisDataStreamMetricFactory) DimensionsMap() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"dimensionsMap",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisDataStreamMetricFactory) MetricFactory() MetricFactory {
	var returns MetricFactory
	_jsii_.Get(
		j,
		"metricFactory",
		&returns,
	)
	return returns
}


// Experimental.
func NewKinesisDataStreamMetricFactory(metricFactory MetricFactory, props *KinesisDataStreamMetricFactoryProps) KinesisDataStreamMetricFactory {
	_init_.Initialize()

	if err := validateNewKinesisDataStreamMetricFactoryParameters(metricFactory, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_KinesisDataStreamMetricFactory{}

	_jsii_.Create(
		"cdk-monitoring-constructs.KinesisDataStreamMetricFactory",
		[]interface{}{metricFactory, props},
		&j,
	)

	return &j
}

// Experimental.
func NewKinesisDataStreamMetricFactory_Override(k KinesisDataStreamMetricFactory, metricFactory MetricFactory, props *KinesisDataStreamMetricFactoryProps) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk-monitoring-constructs.KinesisDataStreamMetricFactory",
		[]interface{}{metricFactory, props},
		k,
	)
}

func (k *jsiiProxy_KinesisDataStreamMetricFactory) MetricGetRecordsIteratorAgeMaxMs() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		k,
		"metricGetRecordsIteratorAgeMaxMs",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KinesisDataStreamMetricFactory) MetricGetRecordsLatencyAverageMs() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		k,
		"metricGetRecordsLatencyAverageMs",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KinesisDataStreamMetricFactory) MetricGetRecordsSuccessCount() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		k,
		"metricGetRecordsSuccessCount",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KinesisDataStreamMetricFactory) MetricGetRecordsSumBytes() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		k,
		"metricGetRecordsSumBytes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KinesisDataStreamMetricFactory) MetricGetRecordsSumCount() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		k,
		"metricGetRecordsSumCount",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KinesisDataStreamMetricFactory) MetricIncomingDataSumBytes() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		k,
		"metricIncomingDataSumBytes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KinesisDataStreamMetricFactory) MetricIncomingDataSumCount() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		k,
		"metricIncomingDataSumCount",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KinesisDataStreamMetricFactory) MetricPutRecordLatencyAverageMs() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		k,
		"metricPutRecordLatencyAverageMs",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KinesisDataStreamMetricFactory) MetricPutRecordsFailedRecordsCount() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		k,
		"metricPutRecordsFailedRecordsCount",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KinesisDataStreamMetricFactory) MetricPutRecordsLatencyAverageMs() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		k,
		"metricPutRecordsLatencyAverageMs",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KinesisDataStreamMetricFactory) MetricPutRecordsSuccessCount() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		k,
		"metricPutRecordsSuccessCount",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KinesisDataStreamMetricFactory) MetricPutRecordsSuccessfulRecordsCount() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		k,
		"metricPutRecordsSuccessfulRecordsCount",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KinesisDataStreamMetricFactory) MetricPutRecordsSumBytes() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		k,
		"metricPutRecordsSumBytes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KinesisDataStreamMetricFactory) MetricPutRecordsThrottledRecordsCount() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		k,
		"metricPutRecordsThrottledRecordsCount",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KinesisDataStreamMetricFactory) MetricPutRecordsTotalRecordsCount() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		k,
		"metricPutRecordsTotalRecordsCount",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KinesisDataStreamMetricFactory) MetricPutRecordSuccessCount() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		k,
		"metricPutRecordSuccessCount",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KinesisDataStreamMetricFactory) MetricPutRecordSumBytes() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		k,
		"metricPutRecordSumBytes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KinesisDataStreamMetricFactory) MetricReadProvisionedThroughputExceeded() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		k,
		"metricReadProvisionedThroughputExceeded",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KinesisDataStreamMetricFactory) MetricReadProvisionedThroughputExceededPercent() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		k,
		"metricReadProvisionedThroughputExceededPercent",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KinesisDataStreamMetricFactory) MetricWriteProvisionedThroughputExceeded() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		k,
		"metricWriteProvisionedThroughputExceeded",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KinesisDataStreamMetricFactory) MetricWriteProvisionedThroughputExceededPercent() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		k,
		"metricWriteProvisionedThroughputExceededPercent",
		nil, // no parameters
		&returns,
	)

	return returns
}


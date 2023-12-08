package cdkmonitoringconstructs

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/corymhall/cdk-constructs-go/cdkmonitoringconstructs/jsii"
)

// See: https://docs.aws.amazon.com/kinesisanalytics/latest/java/metrics-dimensions.html
//
// Experimental.
type KinesisDataAnalyticsMetricFactory interface {
	// Experimental.
	DimensionsMap() *map[string]*string
	// Experimental.
	MetricFactory() MetricFactory
	// Experimental.
	MetricCpuUtilizationPercent() interface{}
	// Experimental.
	MetricDowntimeMs() interface{}
	// Experimental.
	MetricFullRestartsCount() interface{}
	// Experimental.
	MetricHeapMemoryUtilizationPercent() interface{}
	// Experimental.
	MetricKPUsCount() interface{}
	// Experimental.
	MetricLastCheckpointDurationMs() interface{}
	// Experimental.
	MetricLastCheckpointSizeBytes() interface{}
	// Experimental.
	MetricNumberOfFailedCheckpointsCount() interface{}
	// Experimental.
	MetricOldGenerationGCCount() interface{}
	// Experimental.
	MetricOldGenerationGCTimeMs() interface{}
	// Experimental.
	MetricUptimeMs() interface{}
}

// The jsii proxy struct for KinesisDataAnalyticsMetricFactory
type jsiiProxy_KinesisDataAnalyticsMetricFactory struct {
	_ byte // padding
}

func (j *jsiiProxy_KinesisDataAnalyticsMetricFactory) DimensionsMap() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"dimensionsMap",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisDataAnalyticsMetricFactory) MetricFactory() MetricFactory {
	var returns MetricFactory
	_jsii_.Get(
		j,
		"metricFactory",
		&returns,
	)
	return returns
}


// Experimental.
func NewKinesisDataAnalyticsMetricFactory(metricFactory MetricFactory, props *KinesisDataAnalyticsMetricFactoryProps) KinesisDataAnalyticsMetricFactory {
	_init_.Initialize()

	if err := validateNewKinesisDataAnalyticsMetricFactoryParameters(metricFactory, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_KinesisDataAnalyticsMetricFactory{}

	_jsii_.Create(
		"cdk-monitoring-constructs.KinesisDataAnalyticsMetricFactory",
		[]interface{}{metricFactory, props},
		&j,
	)

	return &j
}

// Experimental.
func NewKinesisDataAnalyticsMetricFactory_Override(k KinesisDataAnalyticsMetricFactory, metricFactory MetricFactory, props *KinesisDataAnalyticsMetricFactoryProps) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk-monitoring-constructs.KinesisDataAnalyticsMetricFactory",
		[]interface{}{metricFactory, props},
		k,
	)
}

func (k *jsiiProxy_KinesisDataAnalyticsMetricFactory) MetricCpuUtilizationPercent() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		k,
		"metricCpuUtilizationPercent",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KinesisDataAnalyticsMetricFactory) MetricDowntimeMs() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		k,
		"metricDowntimeMs",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KinesisDataAnalyticsMetricFactory) MetricFullRestartsCount() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		k,
		"metricFullRestartsCount",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KinesisDataAnalyticsMetricFactory) MetricHeapMemoryUtilizationPercent() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		k,
		"metricHeapMemoryUtilizationPercent",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KinesisDataAnalyticsMetricFactory) MetricKPUsCount() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		k,
		"metricKPUsCount",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KinesisDataAnalyticsMetricFactory) MetricLastCheckpointDurationMs() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		k,
		"metricLastCheckpointDurationMs",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KinesisDataAnalyticsMetricFactory) MetricLastCheckpointSizeBytes() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		k,
		"metricLastCheckpointSizeBytes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KinesisDataAnalyticsMetricFactory) MetricNumberOfFailedCheckpointsCount() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		k,
		"metricNumberOfFailedCheckpointsCount",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KinesisDataAnalyticsMetricFactory) MetricOldGenerationGCCount() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		k,
		"metricOldGenerationGCCount",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KinesisDataAnalyticsMetricFactory) MetricOldGenerationGCTimeMs() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		k,
		"metricOldGenerationGCTimeMs",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KinesisDataAnalyticsMetricFactory) MetricUptimeMs() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		k,
		"metricUptimeMs",
		nil, // no parameters
		&returns,
	)

	return returns
}


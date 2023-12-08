package cdkmonitoringconstructs

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/corymhall/cdk-constructs-go/cdkmonitoringconstructs/jsii"
)

// Experimental.
type GlueJobMetricFactory interface {
	// Experimental.
	DimensionsMap() *map[string]*string
	// Experimental.
	MetricFactory() MetricFactory
	// Experimental.
	RateComputationMethod() RateComputationMethod
	// Experimental.
	TypeCountDimensionsMap() *map[string]*string
	// Experimental.
	MetricActiveExecutorsAverage() interface{}
	// Experimental.
	MetricAverageExecutorCpuUsagePercentage() interface{}
	// Experimental.
	MetricAverageExecutorMemoryUsagePercentage() interface{}
	// Experimental.
	MetricCompletedStagesSum() interface{}
	// Experimental.
	MetricCompletedTasksSum() interface{}
	// Experimental.
	MetricFailedTasksRate() interface{}
	// Experimental.
	MetricFailedTasksSum() interface{}
	// Experimental.
	MetricKilledTasksRate() interface{}
	// Experimental.
	MetricKilledTasksSum() interface{}
	// Experimental.
	MetricMaximumNeededExecutors() interface{}
	// Experimental.
	MetricTotalReadBytesFromS3() interface{}
	// Experimental.
	MetricTotalWrittenBytesToS3() interface{}
}

// The jsii proxy struct for GlueJobMetricFactory
type jsiiProxy_GlueJobMetricFactory struct {
	_ byte // padding
}

func (j *jsiiProxy_GlueJobMetricFactory) DimensionsMap() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"dimensionsMap",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueJobMetricFactory) MetricFactory() MetricFactory {
	var returns MetricFactory
	_jsii_.Get(
		j,
		"metricFactory",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueJobMetricFactory) RateComputationMethod() RateComputationMethod {
	var returns RateComputationMethod
	_jsii_.Get(
		j,
		"rateComputationMethod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueJobMetricFactory) TypeCountDimensionsMap() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"typeCountDimensionsMap",
		&returns,
	)
	return returns
}


// Experimental.
func NewGlueJobMetricFactory(metricFactory MetricFactory, props *GlueJobMetricFactoryProps) GlueJobMetricFactory {
	_init_.Initialize()

	if err := validateNewGlueJobMetricFactoryParameters(metricFactory, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_GlueJobMetricFactory{}

	_jsii_.Create(
		"cdk-monitoring-constructs.GlueJobMetricFactory",
		[]interface{}{metricFactory, props},
		&j,
	)

	return &j
}

// Experimental.
func NewGlueJobMetricFactory_Override(g GlueJobMetricFactory, metricFactory MetricFactory, props *GlueJobMetricFactoryProps) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk-monitoring-constructs.GlueJobMetricFactory",
		[]interface{}{metricFactory, props},
		g,
	)
}

func (g *jsiiProxy_GlueJobMetricFactory) MetricActiveExecutorsAverage() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"metricActiveExecutorsAverage",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GlueJobMetricFactory) MetricAverageExecutorCpuUsagePercentage() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"metricAverageExecutorCpuUsagePercentage",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GlueJobMetricFactory) MetricAverageExecutorMemoryUsagePercentage() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"metricAverageExecutorMemoryUsagePercentage",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GlueJobMetricFactory) MetricCompletedStagesSum() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"metricCompletedStagesSum",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GlueJobMetricFactory) MetricCompletedTasksSum() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"metricCompletedTasksSum",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GlueJobMetricFactory) MetricFailedTasksRate() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"metricFailedTasksRate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GlueJobMetricFactory) MetricFailedTasksSum() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"metricFailedTasksSum",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GlueJobMetricFactory) MetricKilledTasksRate() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"metricKilledTasksRate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GlueJobMetricFactory) MetricKilledTasksSum() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"metricKilledTasksSum",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GlueJobMetricFactory) MetricMaximumNeededExecutors() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"metricMaximumNeededExecutors",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GlueJobMetricFactory) MetricTotalReadBytesFromS3() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"metricTotalReadBytesFromS3",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GlueJobMetricFactory) MetricTotalWrittenBytesToS3() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"metricTotalWrittenBytesToS3",
		nil, // no parameters
		&returns,
	)

	return returns
}


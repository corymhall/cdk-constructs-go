package cdkmonitoringconstructs

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/corymhall/cdk-constructs-go/cdkmonitoringconstructs/jsii"
)

// Experimental.
type RedshiftClusterMetricFactory interface {
	// Experimental.
	DimensionsMap() *map[string]*string
	// Experimental.
	MetricFactory() MetricFactory
	// Experimental.
	MetricAverageCpuUsageInPercent() interface{}
	// Experimental.
	MetricAverageDiskSpaceUsageInPercent() interface{}
	// Experimental.
	MetricLongQueryDurationP90InMillis() interface{}
	// Experimental.
	MetricMaintenanceModeEnabled() interface{}
	// Experimental.
	MetricMediumQueryDurationP90InMillis() interface{}
	// Experimental.
	MetricReadLatencyP90InMillis() interface{}
	// Experimental.
	MetricShortQueryDurationP90InMillis() interface{}
	// Experimental.
	MetricTotalConnectionCount() interface{}
	// Experimental.
	MetricWriteLatencyP90InMillis() interface{}
}

// The jsii proxy struct for RedshiftClusterMetricFactory
type jsiiProxy_RedshiftClusterMetricFactory struct {
	_ byte // padding
}

func (j *jsiiProxy_RedshiftClusterMetricFactory) DimensionsMap() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"dimensionsMap",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedshiftClusterMetricFactory) MetricFactory() MetricFactory {
	var returns MetricFactory
	_jsii_.Get(
		j,
		"metricFactory",
		&returns,
	)
	return returns
}


// Experimental.
func NewRedshiftClusterMetricFactory(metricFactory MetricFactory, props *RedshiftClusterMetricFactoryProps) RedshiftClusterMetricFactory {
	_init_.Initialize()

	if err := validateNewRedshiftClusterMetricFactoryParameters(metricFactory, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_RedshiftClusterMetricFactory{}

	_jsii_.Create(
		"cdk-monitoring-constructs.RedshiftClusterMetricFactory",
		[]interface{}{metricFactory, props},
		&j,
	)

	return &j
}

// Experimental.
func NewRedshiftClusterMetricFactory_Override(r RedshiftClusterMetricFactory, metricFactory MetricFactory, props *RedshiftClusterMetricFactoryProps) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk-monitoring-constructs.RedshiftClusterMetricFactory",
		[]interface{}{metricFactory, props},
		r,
	)
}

func (r *jsiiProxy_RedshiftClusterMetricFactory) MetricAverageCpuUsageInPercent() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		r,
		"metricAverageCpuUsageInPercent",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RedshiftClusterMetricFactory) MetricAverageDiskSpaceUsageInPercent() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		r,
		"metricAverageDiskSpaceUsageInPercent",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RedshiftClusterMetricFactory) MetricLongQueryDurationP90InMillis() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		r,
		"metricLongQueryDurationP90InMillis",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RedshiftClusterMetricFactory) MetricMaintenanceModeEnabled() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		r,
		"metricMaintenanceModeEnabled",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RedshiftClusterMetricFactory) MetricMediumQueryDurationP90InMillis() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		r,
		"metricMediumQueryDurationP90InMillis",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RedshiftClusterMetricFactory) MetricReadLatencyP90InMillis() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		r,
		"metricReadLatencyP90InMillis",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RedshiftClusterMetricFactory) MetricShortQueryDurationP90InMillis() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		r,
		"metricShortQueryDurationP90InMillis",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RedshiftClusterMetricFactory) MetricTotalConnectionCount() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		r,
		"metricTotalConnectionCount",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RedshiftClusterMetricFactory) MetricWriteLatencyP90InMillis() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		r,
		"metricWriteLatencyP90InMillis",
		nil, // no parameters
		&returns,
	)

	return returns
}


package cdkmonitoringconstructs

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/corymhall/cdk-constructs-go/cdkmonitoringconstructs/jsii"
)

// Experimental.
type DocumentDbMetricFactory interface {
	// Experimental.
	ClusterIdentifier() *string
	// Experimental.
	DimensionsMap() *map[string]*string
	// Experimental.
	MetricFactory() MetricFactory
	// Experimental.
	MetricAverageCpuUsageInPercent() interface{}
	// Experimental.
	MetricMaxConnectionCount() interface{}
	// Experimental.
	MetricMaxCursorCount() interface{}
	// Experimental.
	MetricMaxTransactionOpenCount() interface{}
	// Experimental.
	MetricOperationsThrottledDueLowMemoryCount() interface{}
	// Experimental.
	MetricReadLatencyInMillis(latencyType LatencyType) interface{}
	// Experimental.
	MetricWriteLatencyInMillis(latencyType LatencyType) interface{}
}

// The jsii proxy struct for DocumentDbMetricFactory
type jsiiProxy_DocumentDbMetricFactory struct {
	_ byte // padding
}

func (j *jsiiProxy_DocumentDbMetricFactory) ClusterIdentifier() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterIdentifier",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DocumentDbMetricFactory) DimensionsMap() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"dimensionsMap",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DocumentDbMetricFactory) MetricFactory() MetricFactory {
	var returns MetricFactory
	_jsii_.Get(
		j,
		"metricFactory",
		&returns,
	)
	return returns
}


// Experimental.
func NewDocumentDbMetricFactory(metricFactory MetricFactory, props *DocumentDbMetricFactoryProps) DocumentDbMetricFactory {
	_init_.Initialize()

	if err := validateNewDocumentDbMetricFactoryParameters(metricFactory, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_DocumentDbMetricFactory{}

	_jsii_.Create(
		"cdk-monitoring-constructs.DocumentDbMetricFactory",
		[]interface{}{metricFactory, props},
		&j,
	)

	return &j
}

// Experimental.
func NewDocumentDbMetricFactory_Override(d DocumentDbMetricFactory, metricFactory MetricFactory, props *DocumentDbMetricFactoryProps) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk-monitoring-constructs.DocumentDbMetricFactory",
		[]interface{}{metricFactory, props},
		d,
	)
}

func (d *jsiiProxy_DocumentDbMetricFactory) MetricAverageCpuUsageInPercent() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"metricAverageCpuUsageInPercent",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DocumentDbMetricFactory) MetricMaxConnectionCount() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"metricMaxConnectionCount",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DocumentDbMetricFactory) MetricMaxCursorCount() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"metricMaxCursorCount",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DocumentDbMetricFactory) MetricMaxTransactionOpenCount() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"metricMaxTransactionOpenCount",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DocumentDbMetricFactory) MetricOperationsThrottledDueLowMemoryCount() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"metricOperationsThrottledDueLowMemoryCount",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DocumentDbMetricFactory) MetricReadLatencyInMillis(latencyType LatencyType) interface{} {
	if err := d.validateMetricReadLatencyInMillisParameters(latencyType); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		d,
		"metricReadLatencyInMillis",
		[]interface{}{latencyType},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DocumentDbMetricFactory) MetricWriteLatencyInMillis(latencyType LatencyType) interface{} {
	if err := d.validateMetricWriteLatencyInMillisParameters(latencyType); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		d,
		"metricWriteLatencyInMillis",
		[]interface{}{latencyType},
		&returns,
	)

	return returns
}


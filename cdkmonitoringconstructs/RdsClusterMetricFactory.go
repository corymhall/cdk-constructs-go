package cdkmonitoringconstructs

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/corymhall/cdk-constructs-go/cdkmonitoringconstructs/jsii"
)

// Experimental.
type RdsClusterMetricFactory interface {
	// Experimental.
	Cluster() interface{}
	// Experimental.
	ClusterIdentifier() *string
	// Experimental.
	DimensionsMap() *map[string]*string
	// Experimental.
	MetricFactory() MetricFactory
	// Experimental.
	MetricAverageCpuUsageInPercent() interface{}
	// Experimental.
	MetricCommitLatencyP90InMillis() interface{}
	// Experimental.
	MetricDeleteLatencyP90InMillis() interface{}
	// Experimental.
	MetricDiskSpaceUsageInPercent() interface{}
	// Experimental.
	MetricFreeStorageInBytes() interface{}
	// Experimental.
	MetricInsertLatencyP90InMillis() interface{}
	// Experimental.
	MetricSelectLatencyP90InMillis() interface{}
	// Experimental.
	MetricServerlessDatabaseCapacity() interface{}
	// Experimental.
	MetricTotalConnectionCount() interface{}
	// Experimental.
	MetricUpdateLatencyP90InMillis() interface{}
	// Experimental.
	MetricUsedStorageInBytes() interface{}
}

// The jsii proxy struct for RdsClusterMetricFactory
type jsiiProxy_RdsClusterMetricFactory struct {
	_ byte // padding
}

func (j *jsiiProxy_RdsClusterMetricFactory) Cluster() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"cluster",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsClusterMetricFactory) ClusterIdentifier() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterIdentifier",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsClusterMetricFactory) DimensionsMap() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"dimensionsMap",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsClusterMetricFactory) MetricFactory() MetricFactory {
	var returns MetricFactory
	_jsii_.Get(
		j,
		"metricFactory",
		&returns,
	)
	return returns
}


// Experimental.
func NewRdsClusterMetricFactory(metricFactory MetricFactory, props *RdsClusterMetricFactoryProps) RdsClusterMetricFactory {
	_init_.Initialize()

	if err := validateNewRdsClusterMetricFactoryParameters(metricFactory, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_RdsClusterMetricFactory{}

	_jsii_.Create(
		"cdk-monitoring-constructs.RdsClusterMetricFactory",
		[]interface{}{metricFactory, props},
		&j,
	)

	return &j
}

// Experimental.
func NewRdsClusterMetricFactory_Override(r RdsClusterMetricFactory, metricFactory MetricFactory, props *RdsClusterMetricFactoryProps) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk-monitoring-constructs.RdsClusterMetricFactory",
		[]interface{}{metricFactory, props},
		r,
	)
}

func (r *jsiiProxy_RdsClusterMetricFactory) MetricAverageCpuUsageInPercent() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		r,
		"metricAverageCpuUsageInPercent",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RdsClusterMetricFactory) MetricCommitLatencyP90InMillis() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		r,
		"metricCommitLatencyP90InMillis",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RdsClusterMetricFactory) MetricDeleteLatencyP90InMillis() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		r,
		"metricDeleteLatencyP90InMillis",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RdsClusterMetricFactory) MetricDiskSpaceUsageInPercent() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		r,
		"metricDiskSpaceUsageInPercent",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RdsClusterMetricFactory) MetricFreeStorageInBytes() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		r,
		"metricFreeStorageInBytes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RdsClusterMetricFactory) MetricInsertLatencyP90InMillis() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		r,
		"metricInsertLatencyP90InMillis",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RdsClusterMetricFactory) MetricSelectLatencyP90InMillis() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		r,
		"metricSelectLatencyP90InMillis",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RdsClusterMetricFactory) MetricServerlessDatabaseCapacity() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		r,
		"metricServerlessDatabaseCapacity",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RdsClusterMetricFactory) MetricTotalConnectionCount() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		r,
		"metricTotalConnectionCount",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RdsClusterMetricFactory) MetricUpdateLatencyP90InMillis() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		r,
		"metricUpdateLatencyP90InMillis",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RdsClusterMetricFactory) MetricUsedStorageInBytes() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		r,
		"metricUsedStorageInBytes",
		nil, // no parameters
		&returns,
	)

	return returns
}


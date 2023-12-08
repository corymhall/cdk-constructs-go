package cdkmonitoringconstructs

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/corymhall/cdk-constructs-go/cdkmonitoringconstructs/jsii"
)

// See: https://docs.aws.amazon.com/AmazonElastiCache/latest/mem-ug/CacheMetrics.html
//
// Experimental.
type ElastiCacheClusterMetricFactory interface {
	// Experimental.
	DimensionsMap() *map[string]*string
	// Experimental.
	MetricFactory() MetricFactory
	// Experimental.
	MetricAverageCachedItemsSizeInBytes() interface{}
	// Experimental.
	MetricAverageConnections() interface{}
	// Experimental.
	MetricAverageFreeableMemoryInBytes() interface{}
	// Experimental.
	MetricAverageSwapUsageInBytes() interface{}
	// Experimental.
	MetricAverageUnusedMemoryInBytes() interface{}
	// Experimental.
	MetricEvictions() interface{}
	// Experimental.
	MetricMaxCpuUtilizationInPercent() interface{}
	// Experimental.
	MetricMaxItemCount() interface{}
	// Because Redis is single-threaded, you can use this metric to analyze the load of the Redis process itself.
	//
	// Note that you may want to monitor both Engine CPU Utilization as well as CPU Utilization as background
	// processes can take up a significant portion of the CPU workload. This is especially important for
	// hosts with 2 vCPUs or less.
	// See: https://docs.aws.amazon.com/AmazonElastiCache/latest/red-ug/CacheMetrics.Redis.html
	//
	// Experimental.
	MetricMaxRedisEngineCpuUtilizationInPercent() interface{}
	// Experimental.
	MetricNetworkBytesIn() interface{}
	// Experimental.
	MetricNetworkBytesOut() interface{}
}

// The jsii proxy struct for ElastiCacheClusterMetricFactory
type jsiiProxy_ElastiCacheClusterMetricFactory struct {
	_ byte // padding
}

func (j *jsiiProxy_ElastiCacheClusterMetricFactory) DimensionsMap() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"dimensionsMap",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastiCacheClusterMetricFactory) MetricFactory() MetricFactory {
	var returns MetricFactory
	_jsii_.Get(
		j,
		"metricFactory",
		&returns,
	)
	return returns
}


// Experimental.
func NewElastiCacheClusterMetricFactory(metricFactory MetricFactory, props *ElastiCacheClusterMetricFactoryProps) ElastiCacheClusterMetricFactory {
	_init_.Initialize()

	if err := validateNewElastiCacheClusterMetricFactoryParameters(metricFactory, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_ElastiCacheClusterMetricFactory{}

	_jsii_.Create(
		"cdk-monitoring-constructs.ElastiCacheClusterMetricFactory",
		[]interface{}{metricFactory, props},
		&j,
	)

	return &j
}

// Experimental.
func NewElastiCacheClusterMetricFactory_Override(e ElastiCacheClusterMetricFactory, metricFactory MetricFactory, props *ElastiCacheClusterMetricFactoryProps) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk-monitoring-constructs.ElastiCacheClusterMetricFactory",
		[]interface{}{metricFactory, props},
		e,
	)
}

func (e *jsiiProxy_ElastiCacheClusterMetricFactory) MetricAverageCachedItemsSizeInBytes() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"metricAverageCachedItemsSizeInBytes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ElastiCacheClusterMetricFactory) MetricAverageConnections() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"metricAverageConnections",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ElastiCacheClusterMetricFactory) MetricAverageFreeableMemoryInBytes() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"metricAverageFreeableMemoryInBytes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ElastiCacheClusterMetricFactory) MetricAverageSwapUsageInBytes() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"metricAverageSwapUsageInBytes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ElastiCacheClusterMetricFactory) MetricAverageUnusedMemoryInBytes() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"metricAverageUnusedMemoryInBytes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ElastiCacheClusterMetricFactory) MetricEvictions() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"metricEvictions",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ElastiCacheClusterMetricFactory) MetricMaxCpuUtilizationInPercent() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"metricMaxCpuUtilizationInPercent",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ElastiCacheClusterMetricFactory) MetricMaxItemCount() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"metricMaxItemCount",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ElastiCacheClusterMetricFactory) MetricMaxRedisEngineCpuUtilizationInPercent() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"metricMaxRedisEngineCpuUtilizationInPercent",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ElastiCacheClusterMetricFactory) MetricNetworkBytesIn() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"metricNetworkBytesIn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ElastiCacheClusterMetricFactory) MetricNetworkBytesOut() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"metricNetworkBytesOut",
		nil, // no parameters
		&returns,
	)

	return returns
}


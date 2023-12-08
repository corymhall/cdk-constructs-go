package cdkmonitoringconstructs

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/corymhall/cdk-constructs-go/cdkmonitoringconstructs/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudwatch"
)

// Experimental.
type OpenSearchClusterMetricFactory interface {
	// Experimental.
	DomainMetrics() OpenSearchBackportedMetrics
	// Experimental.
	FillTpsWithZeroes() *bool
	// Experimental.
	MetricFactory() MetricFactory
	// Experimental.
	RateComputationMethod() RateComputationMethod
	// Experimental.
	MetricAutomatedSnapshotFailure() interface{}
	// Deprecated: use metricClusterIndexWritesBlocked instead.
	MetricClusterIndexWriteBlocked() interface{}
	// Experimental.
	MetricClusterIndexWritesBlocked() interface{}
	// Experimental.
	MetricClusterStatusRed() interface{}
	// Experimental.
	MetricClusterStatusYellow() interface{}
	// Experimental.
	MetricCpuUsage() interface{}
	// Experimental.
	MetricDiskSpaceUsageInPercent() interface{}
	// Experimental.
	MetricIndexingLatencyP50InMillis() interface{}
	// Experimental.
	MetricIndexingLatencyP90InMillis() interface{}
	// Experimental.
	MetricIndexingLatencyP99InMillis() interface{}
	// Experimental.
	MetricJvmMemoryPressure() interface{}
	// Experimental.
	MetricKmsKeyError() interface{}
	// Experimental.
	MetricKmsKeyInaccessible() interface{}
	// Experimental.
	MetricMasterCpuUsage() interface{}
	// Experimental.
	MetricMasterJvmMemoryPressure() interface{}
	// Experimental.
	MetricNodes() interface{}
	// Experimental.
	MetricSearchCount() awscloudwatch.Metric
	// Experimental.
	MetricSearchLatencyP50InMillis() interface{}
	// Experimental.
	MetricSearchLatencyP90InMillis() interface{}
	// Experimental.
	MetricSearchLatencyP99InMillis() interface{}
	// Experimental.
	MetricSearchRate() interface{}
	// Deprecated: use metricSearchRate.
	MetricTps() interface{}
}

// The jsii proxy struct for OpenSearchClusterMetricFactory
type jsiiProxy_OpenSearchClusterMetricFactory struct {
	_ byte // padding
}

func (j *jsiiProxy_OpenSearchClusterMetricFactory) DomainMetrics() OpenSearchBackportedMetrics {
	var returns OpenSearchBackportedMetrics
	_jsii_.Get(
		j,
		"domainMetrics",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpenSearchClusterMetricFactory) FillTpsWithZeroes() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"fillTpsWithZeroes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpenSearchClusterMetricFactory) MetricFactory() MetricFactory {
	var returns MetricFactory
	_jsii_.Get(
		j,
		"metricFactory",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpenSearchClusterMetricFactory) RateComputationMethod() RateComputationMethod {
	var returns RateComputationMethod
	_jsii_.Get(
		j,
		"rateComputationMethod",
		&returns,
	)
	return returns
}


// Experimental.
func NewOpenSearchClusterMetricFactory(metricFactory MetricFactory, props *OpenSearchClusterMetricFactoryProps) OpenSearchClusterMetricFactory {
	_init_.Initialize()

	if err := validateNewOpenSearchClusterMetricFactoryParameters(metricFactory, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_OpenSearchClusterMetricFactory{}

	_jsii_.Create(
		"cdk-monitoring-constructs.OpenSearchClusterMetricFactory",
		[]interface{}{metricFactory, props},
		&j,
	)

	return &j
}

// Experimental.
func NewOpenSearchClusterMetricFactory_Override(o OpenSearchClusterMetricFactory, metricFactory MetricFactory, props *OpenSearchClusterMetricFactoryProps) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk-monitoring-constructs.OpenSearchClusterMetricFactory",
		[]interface{}{metricFactory, props},
		o,
	)
}

func (o *jsiiProxy_OpenSearchClusterMetricFactory) MetricAutomatedSnapshotFailure() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		o,
		"metricAutomatedSnapshotFailure",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OpenSearchClusterMetricFactory) MetricClusterIndexWriteBlocked() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		o,
		"metricClusterIndexWriteBlocked",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OpenSearchClusterMetricFactory) MetricClusterIndexWritesBlocked() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		o,
		"metricClusterIndexWritesBlocked",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OpenSearchClusterMetricFactory) MetricClusterStatusRed() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		o,
		"metricClusterStatusRed",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OpenSearchClusterMetricFactory) MetricClusterStatusYellow() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		o,
		"metricClusterStatusYellow",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OpenSearchClusterMetricFactory) MetricCpuUsage() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		o,
		"metricCpuUsage",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OpenSearchClusterMetricFactory) MetricDiskSpaceUsageInPercent() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		o,
		"metricDiskSpaceUsageInPercent",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OpenSearchClusterMetricFactory) MetricIndexingLatencyP50InMillis() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		o,
		"metricIndexingLatencyP50InMillis",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OpenSearchClusterMetricFactory) MetricIndexingLatencyP90InMillis() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		o,
		"metricIndexingLatencyP90InMillis",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OpenSearchClusterMetricFactory) MetricIndexingLatencyP99InMillis() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		o,
		"metricIndexingLatencyP99InMillis",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OpenSearchClusterMetricFactory) MetricJvmMemoryPressure() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		o,
		"metricJvmMemoryPressure",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OpenSearchClusterMetricFactory) MetricKmsKeyError() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		o,
		"metricKmsKeyError",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OpenSearchClusterMetricFactory) MetricKmsKeyInaccessible() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		o,
		"metricKmsKeyInaccessible",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OpenSearchClusterMetricFactory) MetricMasterCpuUsage() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		o,
		"metricMasterCpuUsage",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OpenSearchClusterMetricFactory) MetricMasterJvmMemoryPressure() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		o,
		"metricMasterJvmMemoryPressure",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OpenSearchClusterMetricFactory) MetricNodes() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		o,
		"metricNodes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OpenSearchClusterMetricFactory) MetricSearchCount() awscloudwatch.Metric {
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		o,
		"metricSearchCount",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OpenSearchClusterMetricFactory) MetricSearchLatencyP50InMillis() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		o,
		"metricSearchLatencyP50InMillis",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OpenSearchClusterMetricFactory) MetricSearchLatencyP90InMillis() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		o,
		"metricSearchLatencyP90InMillis",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OpenSearchClusterMetricFactory) MetricSearchLatencyP99InMillis() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		o,
		"metricSearchLatencyP99InMillis",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OpenSearchClusterMetricFactory) MetricSearchRate() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		o,
		"metricSearchRate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OpenSearchClusterMetricFactory) MetricTps() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		o,
		"metricTps",
		nil, // no parameters
		&returns,
	)

	return returns
}


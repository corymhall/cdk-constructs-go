package cdkmonitoringconstructs

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/corymhall/cdk-constructs-go/cdkmonitoringconstructs/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudwatch"
)

// Backported set of metric functions added in @aws-cdk/aws-elasticsearch@1.65.0.
// See: https://github.com/aws/aws-cdk/releases/tag/v1.73.0
//
// Experimental.
type OpenSearchBackportedMetrics interface {
	// Experimental.
	DimensionsMap() *map[string]*string
	// Return the given named metric for this Domain.
	// Experimental.
	Metric(metricName *string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Metric for automated snapshot failures.
	// Default: - maximum over 5 minutes.
	//
	// Experimental.
	MetricAutomatedSnapshotFailure(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Metric for the cluster blocking index writes.
	// Default: - maximum over 1 minute.
	//
	// Deprecated: use metricClusterIndexWritesBlocked instead.
	MetricClusterIndexWriteBlocked(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Metric for the cluster blocking index writes.
	// Default: - maximum over 1 minute.
	//
	// Experimental.
	MetricClusterIndexWritesBlocked(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Metric for the time the cluster status is red.
	// Default: - maximum over 5 minutes.
	//
	// Experimental.
	MetricClusterStatusRed(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Metric for the time the cluster status is yellow.
	// Default: - maximum over 5 minutes.
	//
	// Experimental.
	MetricClusterStatusYellow(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Metric for CPU utilization.
	// Default: - maximum over 5 minutes.
	//
	// Experimental.
	MetricCPUUtilization(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Metric for the storage space of nodes in the cluster.
	// Default: - minimum over 5 minutes.
	//
	// Experimental.
	MetricFreeStorageSpace(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Metric for indexing latency.
	// Default: - p99 over 5 minutes.
	//
	// Experimental.
	MetricIndexingLatency(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Metric for JVM memory pressure.
	// Default: - maximum over 5 minutes.
	//
	// Experimental.
	MetricJVMMemoryPressure(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Metric for KMS key errors.
	// Default: - maximum over 5 minutes.
	//
	// Experimental.
	MetricKMSKeyError(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Metric for KMS key being inaccessible.
	// Default: - maximum over 5 minutes.
	//
	// Experimental.
	MetricKMSKeyInaccessible(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Metric for master CPU utilization.
	// Default: - maximum over 5 minutes.
	//
	// Experimental.
	MetricMasterCPUUtilization(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Metric for master JVM memory pressure.
	// Default: - maximum over 5 minutes.
	//
	// Experimental.
	MetricMasterJVMMemoryPressure(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Metric for the number of nodes.
	// Default: - minimum over 1 hour.
	//
	// Experimental.
	MetricNodes(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Metric for number of searchable documents.
	// Default: - maximum over 5 minutes.
	//
	// Experimental.
	MetricSearchableDocuments(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Metric for search latency.
	// Default: - p99 over 5 minutes.
	//
	// Experimental.
	MetricSearchLatency(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
}

// The jsii proxy struct for OpenSearchBackportedMetrics
type jsiiProxy_OpenSearchBackportedMetrics struct {
	_ byte // padding
}

func (j *jsiiProxy_OpenSearchBackportedMetrics) DimensionsMap() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"dimensionsMap",
		&returns,
	)
	return returns
}


// Experimental.
func NewOpenSearchBackportedMetrics(domain interface{}) OpenSearchBackportedMetrics {
	_init_.Initialize()

	if err := validateNewOpenSearchBackportedMetricsParameters(domain); err != nil {
		panic(err)
	}
	j := jsiiProxy_OpenSearchBackportedMetrics{}

	_jsii_.Create(
		"cdk-monitoring-constructs.OpenSearchBackportedMetrics",
		[]interface{}{domain},
		&j,
	)

	return &j
}

// Experimental.
func NewOpenSearchBackportedMetrics_Override(o OpenSearchBackportedMetrics, domain interface{}) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk-monitoring-constructs.OpenSearchBackportedMetrics",
		[]interface{}{domain},
		o,
	)
}

func (o *jsiiProxy_OpenSearchBackportedMetrics) Metric(metricName *string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := o.validateMetricParameters(metricName, props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		o,
		"metric",
		[]interface{}{metricName, props},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OpenSearchBackportedMetrics) MetricAutomatedSnapshotFailure(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := o.validateMetricAutomatedSnapshotFailureParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		o,
		"metricAutomatedSnapshotFailure",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OpenSearchBackportedMetrics) MetricClusterIndexWriteBlocked(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := o.validateMetricClusterIndexWriteBlockedParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		o,
		"metricClusterIndexWriteBlocked",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OpenSearchBackportedMetrics) MetricClusterIndexWritesBlocked(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := o.validateMetricClusterIndexWritesBlockedParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		o,
		"metricClusterIndexWritesBlocked",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OpenSearchBackportedMetrics) MetricClusterStatusRed(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := o.validateMetricClusterStatusRedParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		o,
		"metricClusterStatusRed",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OpenSearchBackportedMetrics) MetricClusterStatusYellow(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := o.validateMetricClusterStatusYellowParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		o,
		"metricClusterStatusYellow",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OpenSearchBackportedMetrics) MetricCPUUtilization(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := o.validateMetricCPUUtilizationParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		o,
		"metricCPUUtilization",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OpenSearchBackportedMetrics) MetricFreeStorageSpace(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := o.validateMetricFreeStorageSpaceParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		o,
		"metricFreeStorageSpace",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OpenSearchBackportedMetrics) MetricIndexingLatency(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := o.validateMetricIndexingLatencyParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		o,
		"metricIndexingLatency",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OpenSearchBackportedMetrics) MetricJVMMemoryPressure(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := o.validateMetricJVMMemoryPressureParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		o,
		"metricJVMMemoryPressure",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OpenSearchBackportedMetrics) MetricKMSKeyError(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := o.validateMetricKMSKeyErrorParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		o,
		"metricKMSKeyError",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OpenSearchBackportedMetrics) MetricKMSKeyInaccessible(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := o.validateMetricKMSKeyInaccessibleParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		o,
		"metricKMSKeyInaccessible",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OpenSearchBackportedMetrics) MetricMasterCPUUtilization(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := o.validateMetricMasterCPUUtilizationParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		o,
		"metricMasterCPUUtilization",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OpenSearchBackportedMetrics) MetricMasterJVMMemoryPressure(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := o.validateMetricMasterJVMMemoryPressureParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		o,
		"metricMasterJVMMemoryPressure",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OpenSearchBackportedMetrics) MetricNodes(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := o.validateMetricNodesParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		o,
		"metricNodes",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OpenSearchBackportedMetrics) MetricSearchableDocuments(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := o.validateMetricSearchableDocumentsParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		o,
		"metricSearchableDocuments",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OpenSearchBackportedMetrics) MetricSearchLatency(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := o.validateMetricSearchLatencyParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		o,
		"metricSearchLatency",
		[]interface{}{props},
		&returns,
	)

	return returns
}


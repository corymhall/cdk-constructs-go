package cdkmonitoringconstructs

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/corymhall/cdk-constructs-go/cdkmonitoringconstructs/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudwatch"
)

// Experimental.
type OpenSearchClusterMonitoring interface {
	Monitoring
	// Experimental.
	Alarms() *[]*AlarmWithAnnotation
	// Experimental.
	AutomatedSnapshotFailureMetric() interface{}
	// Experimental.
	ClusterAlarmFactory() OpenSearchClusterAlarmFactory
	// Experimental.
	ClusterAnnotations() *[]*awscloudwatch.HorizontalAnnotation
	// Experimental.
	ClusterStatusRedMetric() interface{}
	// Experimental.
	ClusterStatusYellowMetric() interface{}
	// Experimental.
	CpuUsageMetric() interface{}
	// Experimental.
	DiskSpaceUsageMetric() interface{}
	// Experimental.
	IndexingLatencyAlarmFactory() LatencyAlarmFactory
	// Experimental.
	IndexingLatencyAnnotations() *[]*awscloudwatch.HorizontalAnnotation
	// Experimental.
	IndexWriteBlockedMetric() interface{}
	// Experimental.
	JvmMemoryPressureMetric() interface{}
	// Experimental.
	KmsKeyErrorMetric() interface{}
	// Experimental.
	KmsKeyInaccessibleMetric() interface{}
	// Experimental.
	LocalAlarmNamePrefixOverride() *string
	// Experimental.
	MasterCpuUsageMetric() interface{}
	// Experimental.
	MasterJvmMemoryPressureMetric() interface{}
	// Experimental.
	MasterUsageAnnotations() *[]*awscloudwatch.HorizontalAnnotation
	// Experimental.
	NodeAnnotations() *[]*awscloudwatch.HorizontalAnnotation
	// Experimental.
	NodesMetric() interface{}
	// Experimental.
	P50IndexingLatencyMetric() interface{}
	// Experimental.
	P50SearchLatencyMetric() interface{}
	// Experimental.
	P90IndexingLatencyMetric() interface{}
	// Experimental.
	P90SearchLatencyMetric() interface{}
	// Experimental.
	P99IndexingLatencyMetric() interface{}
	// Experimental.
	P99SearchLatencyMetric() interface{}
	// Experimental.
	Scope() MonitoringScope
	// Experimental.
	SearchLatencyAlarmFactory() LatencyAlarmFactory
	// Experimental.
	SearchLatencyAnnotations() *[]*awscloudwatch.HorizontalAnnotation
	// Experimental.
	Title() *string
	// Experimental.
	TpsMetric() interface{}
	// Experimental.
	Url() *string
	// Experimental.
	UsageAlarmFactory() UsageAlarmFactory
	// Experimental.
	UsageAnnotations() *[]*awscloudwatch.HorizontalAnnotation
	// Adds an alarm.
	// Experimental.
	AddAlarm(alarm *AlarmWithAnnotation)
	// Returns widgets for all alarms.
	//
	// These can go to runbook or to service dashboard.
	// Experimental.
	AlarmWidgets() *[]awscloudwatch.IWidget
	// Creates a new alarm factory.
	//
	// Alarms created will be named with the given prefix, unless a local name override is present.
	// Experimental.
	CreateAlarmFactory(alarmNamePrefix *string) AlarmFactory
	// Returns all the alarms created.
	// Experimental.
	CreatedAlarms() *[]*AlarmWithAnnotation
	// Creates a new metric factory.
	// Experimental.
	CreateMetricFactory() MetricFactory
	// Creates a new widget factory.
	// Experimental.
	CreateWidgetFactory() IWidgetFactory
	// Returns widgets to be placed on the summary dashboard.
	// Experimental.
	SummaryWidgets() *[]awscloudwatch.IWidget
	// Returns widgets to be placed on the main dashboard.
	// Experimental.
	Widgets() *[]awscloudwatch.IWidget
	// Returns widgets for the requested dashboard type.
	// Experimental.
	WidgetsForDashboard(name *string) *[]awscloudwatch.IWidget
}

// The jsii proxy struct for OpenSearchClusterMonitoring
type jsiiProxy_OpenSearchClusterMonitoring struct {
	jsiiProxy_Monitoring
}

func (j *jsiiProxy_OpenSearchClusterMonitoring) Alarms() *[]*AlarmWithAnnotation {
	var returns *[]*AlarmWithAnnotation
	_jsii_.Get(
		j,
		"alarms",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpenSearchClusterMonitoring) AutomatedSnapshotFailureMetric() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"automatedSnapshotFailureMetric",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpenSearchClusterMonitoring) ClusterAlarmFactory() OpenSearchClusterAlarmFactory {
	var returns OpenSearchClusterAlarmFactory
	_jsii_.Get(
		j,
		"clusterAlarmFactory",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpenSearchClusterMonitoring) ClusterAnnotations() *[]*awscloudwatch.HorizontalAnnotation {
	var returns *[]*awscloudwatch.HorizontalAnnotation
	_jsii_.Get(
		j,
		"clusterAnnotations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpenSearchClusterMonitoring) ClusterStatusRedMetric() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"clusterStatusRedMetric",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpenSearchClusterMonitoring) ClusterStatusYellowMetric() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"clusterStatusYellowMetric",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpenSearchClusterMonitoring) CpuUsageMetric() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"cpuUsageMetric",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpenSearchClusterMonitoring) DiskSpaceUsageMetric() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"diskSpaceUsageMetric",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpenSearchClusterMonitoring) IndexingLatencyAlarmFactory() LatencyAlarmFactory {
	var returns LatencyAlarmFactory
	_jsii_.Get(
		j,
		"indexingLatencyAlarmFactory",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpenSearchClusterMonitoring) IndexingLatencyAnnotations() *[]*awscloudwatch.HorizontalAnnotation {
	var returns *[]*awscloudwatch.HorizontalAnnotation
	_jsii_.Get(
		j,
		"indexingLatencyAnnotations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpenSearchClusterMonitoring) IndexWriteBlockedMetric() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"indexWriteBlockedMetric",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpenSearchClusterMonitoring) JvmMemoryPressureMetric() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"jvmMemoryPressureMetric",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpenSearchClusterMonitoring) KmsKeyErrorMetric() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"kmsKeyErrorMetric",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpenSearchClusterMonitoring) KmsKeyInaccessibleMetric() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"kmsKeyInaccessibleMetric",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpenSearchClusterMonitoring) LocalAlarmNamePrefixOverride() *string {
	var returns *string
	_jsii_.Get(
		j,
		"localAlarmNamePrefixOverride",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpenSearchClusterMonitoring) MasterCpuUsageMetric() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"masterCpuUsageMetric",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpenSearchClusterMonitoring) MasterJvmMemoryPressureMetric() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"masterJvmMemoryPressureMetric",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpenSearchClusterMonitoring) MasterUsageAnnotations() *[]*awscloudwatch.HorizontalAnnotation {
	var returns *[]*awscloudwatch.HorizontalAnnotation
	_jsii_.Get(
		j,
		"masterUsageAnnotations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpenSearchClusterMonitoring) NodeAnnotations() *[]*awscloudwatch.HorizontalAnnotation {
	var returns *[]*awscloudwatch.HorizontalAnnotation
	_jsii_.Get(
		j,
		"nodeAnnotations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpenSearchClusterMonitoring) NodesMetric() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"nodesMetric",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpenSearchClusterMonitoring) P50IndexingLatencyMetric() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"p50IndexingLatencyMetric",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpenSearchClusterMonitoring) P50SearchLatencyMetric() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"p50SearchLatencyMetric",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpenSearchClusterMonitoring) P90IndexingLatencyMetric() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"p90IndexingLatencyMetric",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpenSearchClusterMonitoring) P90SearchLatencyMetric() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"p90SearchLatencyMetric",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpenSearchClusterMonitoring) P99IndexingLatencyMetric() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"p99IndexingLatencyMetric",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpenSearchClusterMonitoring) P99SearchLatencyMetric() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"p99SearchLatencyMetric",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpenSearchClusterMonitoring) Scope() MonitoringScope {
	var returns MonitoringScope
	_jsii_.Get(
		j,
		"scope",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpenSearchClusterMonitoring) SearchLatencyAlarmFactory() LatencyAlarmFactory {
	var returns LatencyAlarmFactory
	_jsii_.Get(
		j,
		"searchLatencyAlarmFactory",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpenSearchClusterMonitoring) SearchLatencyAnnotations() *[]*awscloudwatch.HorizontalAnnotation {
	var returns *[]*awscloudwatch.HorizontalAnnotation
	_jsii_.Get(
		j,
		"searchLatencyAnnotations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpenSearchClusterMonitoring) Title() *string {
	var returns *string
	_jsii_.Get(
		j,
		"title",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpenSearchClusterMonitoring) TpsMetric() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tpsMetric",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpenSearchClusterMonitoring) Url() *string {
	var returns *string
	_jsii_.Get(
		j,
		"url",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpenSearchClusterMonitoring) UsageAlarmFactory() UsageAlarmFactory {
	var returns UsageAlarmFactory
	_jsii_.Get(
		j,
		"usageAlarmFactory",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpenSearchClusterMonitoring) UsageAnnotations() *[]*awscloudwatch.HorizontalAnnotation {
	var returns *[]*awscloudwatch.HorizontalAnnotation
	_jsii_.Get(
		j,
		"usageAnnotations",
		&returns,
	)
	return returns
}


// Experimental.
func NewOpenSearchClusterMonitoring(scope MonitoringScope, props *OpenSearchClusterMonitoringProps) OpenSearchClusterMonitoring {
	_init_.Initialize()

	if err := validateNewOpenSearchClusterMonitoringParameters(scope, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_OpenSearchClusterMonitoring{}

	_jsii_.Create(
		"cdk-monitoring-constructs.OpenSearchClusterMonitoring",
		[]interface{}{scope, props},
		&j,
	)

	return &j
}

// Experimental.
func NewOpenSearchClusterMonitoring_Override(o OpenSearchClusterMonitoring, scope MonitoringScope, props *OpenSearchClusterMonitoringProps) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk-monitoring-constructs.OpenSearchClusterMonitoring",
		[]interface{}{scope, props},
		o,
	)
}

func (o *jsiiProxy_OpenSearchClusterMonitoring) AddAlarm(alarm *AlarmWithAnnotation) {
	if err := o.validateAddAlarmParameters(alarm); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"addAlarm",
		[]interface{}{alarm},
	)
}

func (o *jsiiProxy_OpenSearchClusterMonitoring) AlarmWidgets() *[]awscloudwatch.IWidget {
	var returns *[]awscloudwatch.IWidget

	_jsii_.Invoke(
		o,
		"alarmWidgets",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OpenSearchClusterMonitoring) CreateAlarmFactory(alarmNamePrefix *string) AlarmFactory {
	if err := o.validateCreateAlarmFactoryParameters(alarmNamePrefix); err != nil {
		panic(err)
	}
	var returns AlarmFactory

	_jsii_.Invoke(
		o,
		"createAlarmFactory",
		[]interface{}{alarmNamePrefix},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OpenSearchClusterMonitoring) CreatedAlarms() *[]*AlarmWithAnnotation {
	var returns *[]*AlarmWithAnnotation

	_jsii_.Invoke(
		o,
		"createdAlarms",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OpenSearchClusterMonitoring) CreateMetricFactory() MetricFactory {
	var returns MetricFactory

	_jsii_.Invoke(
		o,
		"createMetricFactory",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OpenSearchClusterMonitoring) CreateWidgetFactory() IWidgetFactory {
	var returns IWidgetFactory

	_jsii_.Invoke(
		o,
		"createWidgetFactory",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OpenSearchClusterMonitoring) SummaryWidgets() *[]awscloudwatch.IWidget {
	var returns *[]awscloudwatch.IWidget

	_jsii_.Invoke(
		o,
		"summaryWidgets",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OpenSearchClusterMonitoring) Widgets() *[]awscloudwatch.IWidget {
	var returns *[]awscloudwatch.IWidget

	_jsii_.Invoke(
		o,
		"widgets",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OpenSearchClusterMonitoring) WidgetsForDashboard(name *string) *[]awscloudwatch.IWidget {
	if err := o.validateWidgetsForDashboardParameters(name); err != nil {
		panic(err)
	}
	var returns *[]awscloudwatch.IWidget

	_jsii_.Invoke(
		o,
		"widgetsForDashboard",
		[]interface{}{name},
		&returns,
	)

	return returns
}


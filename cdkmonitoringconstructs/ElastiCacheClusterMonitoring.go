package cdkmonitoringconstructs

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/corymhall/cdk-constructs-go/cdkmonitoringconstructs/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudwatch"
)

// Experimental.
type ElastiCacheClusterMonitoring interface {
	Monitoring
	// Experimental.
	Alarms() *[]*AlarmWithAnnotation
	// Experimental.
	ClusterType() ElastiCacheClusterType
	// Experimental.
	ClusterUrl() *string
	// Experimental.
	ConnectionsMetric() interface{}
	// Experimental.
	CpuUsageAnnotations() *[]*awscloudwatch.HorizontalAnnotation
	// Experimental.
	CpuUsageMetric() interface{}
	// Experimental.
	ElastiCacheAlarmFactory() ElastiCacheAlarmFactory
	// Experimental.
	EvictedItemsCountAnnotations() *[]*awscloudwatch.HorizontalAnnotation
	// Experimental.
	FreeableMemoryMetric() interface{}
	// Experimental.
	ItemsCountAnnotations() *[]*awscloudwatch.HorizontalAnnotation
	// Experimental.
	ItemsCountMetrics() interface{}
	// Experimental.
	ItemsEvictedMetrics() interface{}
	// Experimental.
	ItemsMemoryMetric() interface{}
	// Experimental.
	LocalAlarmNamePrefixOverride() *string
	// Experimental.
	MemoryUsageAnnotations() *[]*awscloudwatch.HorizontalAnnotation
	// Experimental.
	RedisEngineCpuUsageAnnotations() *[]*awscloudwatch.HorizontalAnnotation
	// Experimental.
	RedisEngineCpuUsageMetric() interface{}
	// Experimental.
	Scope() MonitoringScope
	// Experimental.
	SwapMemoryMetric() interface{}
	// Experimental.
	Title() *string
	// Experimental.
	UnusedMemoryMetric() interface{}
	// Experimental.
	UsageAlarmFactory() UsageAlarmFactory
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
	// Experimental.
	CreateConnectionsWidget(width *float64, height *float64) awscloudwatch.GraphWidget
	// Experimental.
	CreateCpuUsageWidget(width *float64, height *float64) awscloudwatch.GraphWidget
	// Returns all the alarms created.
	// Experimental.
	CreatedAlarms() *[]*AlarmWithAnnotation
	// Experimental.
	CreateItemCountWidget(width *float64, height *float64) awscloudwatch.GraphWidget
	// Experimental.
	CreateMemoryUsageWidget(width *float64, height *float64) awscloudwatch.GraphWidget
	// Creates a new metric factory.
	// Experimental.
	CreateMetricFactory() MetricFactory
	// Experimental.
	CreateRedisEngineCpuUsageWidget(width *float64, height *float64) awscloudwatch.GraphWidget
	// Experimental.
	CreateTitleWidget() MonitoringHeaderWidget
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

// The jsii proxy struct for ElastiCacheClusterMonitoring
type jsiiProxy_ElastiCacheClusterMonitoring struct {
	jsiiProxy_Monitoring
}

func (j *jsiiProxy_ElastiCacheClusterMonitoring) Alarms() *[]*AlarmWithAnnotation {
	var returns *[]*AlarmWithAnnotation
	_jsii_.Get(
		j,
		"alarms",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastiCacheClusterMonitoring) ClusterType() ElastiCacheClusterType {
	var returns ElastiCacheClusterType
	_jsii_.Get(
		j,
		"clusterType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastiCacheClusterMonitoring) ClusterUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastiCacheClusterMonitoring) ConnectionsMetric() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connectionsMetric",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastiCacheClusterMonitoring) CpuUsageAnnotations() *[]*awscloudwatch.HorizontalAnnotation {
	var returns *[]*awscloudwatch.HorizontalAnnotation
	_jsii_.Get(
		j,
		"cpuUsageAnnotations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastiCacheClusterMonitoring) CpuUsageMetric() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"cpuUsageMetric",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastiCacheClusterMonitoring) ElastiCacheAlarmFactory() ElastiCacheAlarmFactory {
	var returns ElastiCacheAlarmFactory
	_jsii_.Get(
		j,
		"elastiCacheAlarmFactory",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastiCacheClusterMonitoring) EvictedItemsCountAnnotations() *[]*awscloudwatch.HorizontalAnnotation {
	var returns *[]*awscloudwatch.HorizontalAnnotation
	_jsii_.Get(
		j,
		"evictedItemsCountAnnotations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastiCacheClusterMonitoring) FreeableMemoryMetric() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"freeableMemoryMetric",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastiCacheClusterMonitoring) ItemsCountAnnotations() *[]*awscloudwatch.HorizontalAnnotation {
	var returns *[]*awscloudwatch.HorizontalAnnotation
	_jsii_.Get(
		j,
		"itemsCountAnnotations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastiCacheClusterMonitoring) ItemsCountMetrics() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"itemsCountMetrics",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastiCacheClusterMonitoring) ItemsEvictedMetrics() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"itemsEvictedMetrics",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastiCacheClusterMonitoring) ItemsMemoryMetric() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"itemsMemoryMetric",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastiCacheClusterMonitoring) LocalAlarmNamePrefixOverride() *string {
	var returns *string
	_jsii_.Get(
		j,
		"localAlarmNamePrefixOverride",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastiCacheClusterMonitoring) MemoryUsageAnnotations() *[]*awscloudwatch.HorizontalAnnotation {
	var returns *[]*awscloudwatch.HorizontalAnnotation
	_jsii_.Get(
		j,
		"memoryUsageAnnotations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastiCacheClusterMonitoring) RedisEngineCpuUsageAnnotations() *[]*awscloudwatch.HorizontalAnnotation {
	var returns *[]*awscloudwatch.HorizontalAnnotation
	_jsii_.Get(
		j,
		"redisEngineCpuUsageAnnotations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastiCacheClusterMonitoring) RedisEngineCpuUsageMetric() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"redisEngineCpuUsageMetric",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastiCacheClusterMonitoring) Scope() MonitoringScope {
	var returns MonitoringScope
	_jsii_.Get(
		j,
		"scope",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastiCacheClusterMonitoring) SwapMemoryMetric() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"swapMemoryMetric",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastiCacheClusterMonitoring) Title() *string {
	var returns *string
	_jsii_.Get(
		j,
		"title",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastiCacheClusterMonitoring) UnusedMemoryMetric() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"unusedMemoryMetric",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElastiCacheClusterMonitoring) UsageAlarmFactory() UsageAlarmFactory {
	var returns UsageAlarmFactory
	_jsii_.Get(
		j,
		"usageAlarmFactory",
		&returns,
	)
	return returns
}


// Experimental.
func NewElastiCacheClusterMonitoring(scope MonitoringScope, props *ElastiCacheClusterMonitoringProps) ElastiCacheClusterMonitoring {
	_init_.Initialize()

	if err := validateNewElastiCacheClusterMonitoringParameters(scope, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_ElastiCacheClusterMonitoring{}

	_jsii_.Create(
		"cdk-monitoring-constructs.ElastiCacheClusterMonitoring",
		[]interface{}{scope, props},
		&j,
	)

	return &j
}

// Experimental.
func NewElastiCacheClusterMonitoring_Override(e ElastiCacheClusterMonitoring, scope MonitoringScope, props *ElastiCacheClusterMonitoringProps) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk-monitoring-constructs.ElastiCacheClusterMonitoring",
		[]interface{}{scope, props},
		e,
	)
}

func (e *jsiiProxy_ElastiCacheClusterMonitoring) AddAlarm(alarm *AlarmWithAnnotation) {
	if err := e.validateAddAlarmParameters(alarm); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"addAlarm",
		[]interface{}{alarm},
	)
}

func (e *jsiiProxy_ElastiCacheClusterMonitoring) AlarmWidgets() *[]awscloudwatch.IWidget {
	var returns *[]awscloudwatch.IWidget

	_jsii_.Invoke(
		e,
		"alarmWidgets",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ElastiCacheClusterMonitoring) CreateAlarmFactory(alarmNamePrefix *string) AlarmFactory {
	if err := e.validateCreateAlarmFactoryParameters(alarmNamePrefix); err != nil {
		panic(err)
	}
	var returns AlarmFactory

	_jsii_.Invoke(
		e,
		"createAlarmFactory",
		[]interface{}{alarmNamePrefix},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ElastiCacheClusterMonitoring) CreateConnectionsWidget(width *float64, height *float64) awscloudwatch.GraphWidget {
	if err := e.validateCreateConnectionsWidgetParameters(width, height); err != nil {
		panic(err)
	}
	var returns awscloudwatch.GraphWidget

	_jsii_.Invoke(
		e,
		"createConnectionsWidget",
		[]interface{}{width, height},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ElastiCacheClusterMonitoring) CreateCpuUsageWidget(width *float64, height *float64) awscloudwatch.GraphWidget {
	if err := e.validateCreateCpuUsageWidgetParameters(width, height); err != nil {
		panic(err)
	}
	var returns awscloudwatch.GraphWidget

	_jsii_.Invoke(
		e,
		"createCpuUsageWidget",
		[]interface{}{width, height},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ElastiCacheClusterMonitoring) CreatedAlarms() *[]*AlarmWithAnnotation {
	var returns *[]*AlarmWithAnnotation

	_jsii_.Invoke(
		e,
		"createdAlarms",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ElastiCacheClusterMonitoring) CreateItemCountWidget(width *float64, height *float64) awscloudwatch.GraphWidget {
	if err := e.validateCreateItemCountWidgetParameters(width, height); err != nil {
		panic(err)
	}
	var returns awscloudwatch.GraphWidget

	_jsii_.Invoke(
		e,
		"createItemCountWidget",
		[]interface{}{width, height},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ElastiCacheClusterMonitoring) CreateMemoryUsageWidget(width *float64, height *float64) awscloudwatch.GraphWidget {
	if err := e.validateCreateMemoryUsageWidgetParameters(width, height); err != nil {
		panic(err)
	}
	var returns awscloudwatch.GraphWidget

	_jsii_.Invoke(
		e,
		"createMemoryUsageWidget",
		[]interface{}{width, height},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ElastiCacheClusterMonitoring) CreateMetricFactory() MetricFactory {
	var returns MetricFactory

	_jsii_.Invoke(
		e,
		"createMetricFactory",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ElastiCacheClusterMonitoring) CreateRedisEngineCpuUsageWidget(width *float64, height *float64) awscloudwatch.GraphWidget {
	if err := e.validateCreateRedisEngineCpuUsageWidgetParameters(width, height); err != nil {
		panic(err)
	}
	var returns awscloudwatch.GraphWidget

	_jsii_.Invoke(
		e,
		"createRedisEngineCpuUsageWidget",
		[]interface{}{width, height},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ElastiCacheClusterMonitoring) CreateTitleWidget() MonitoringHeaderWidget {
	var returns MonitoringHeaderWidget

	_jsii_.Invoke(
		e,
		"createTitleWidget",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ElastiCacheClusterMonitoring) CreateWidgetFactory() IWidgetFactory {
	var returns IWidgetFactory

	_jsii_.Invoke(
		e,
		"createWidgetFactory",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ElastiCacheClusterMonitoring) SummaryWidgets() *[]awscloudwatch.IWidget {
	var returns *[]awscloudwatch.IWidget

	_jsii_.Invoke(
		e,
		"summaryWidgets",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ElastiCacheClusterMonitoring) Widgets() *[]awscloudwatch.IWidget {
	var returns *[]awscloudwatch.IWidget

	_jsii_.Invoke(
		e,
		"widgets",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ElastiCacheClusterMonitoring) WidgetsForDashboard(name *string) *[]awscloudwatch.IWidget {
	if err := e.validateWidgetsForDashboardParameters(name); err != nil {
		panic(err)
	}
	var returns *[]awscloudwatch.IWidget

	_jsii_.Invoke(
		e,
		"widgetsForDashboard",
		[]interface{}{name},
		&returns,
	)

	return returns
}


package cdkmonitoringconstructs

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/corymhall/cdk-constructs-go/cdkmonitoringconstructs/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudwatch"
)

// Experimental.
type RedshiftClusterMonitoring interface {
	Monitoring
	// Experimental.
	Alarms() *[]*AlarmWithAnnotation
	// Experimental.
	ConnectionAlarmFactory() ConnectionAlarmFactory
	// Experimental.
	ConnectionAnnotations() *[]*awscloudwatch.HorizontalAnnotation
	// Experimental.
	ConnectionsMetric() interface{}
	// Experimental.
	CpuUsageMetric() interface{}
	// Experimental.
	DiskSpaceUsageMetric() interface{}
	// Experimental.
	LatencyAlarmFactory() LatencyAlarmFactory
	// Experimental.
	LocalAlarmNamePrefixOverride() *string
	// Experimental.
	LongQueryDurationMetric() interface{}
	// Experimental.
	MaintenanceModeMetric() interface{}
	// Experimental.
	MediumQueryDurationMetric() interface{}
	// Experimental.
	QueryDurationAnnotations() *[]*awscloudwatch.HorizontalAnnotation
	// Experimental.
	ReadLatencyMetric() interface{}
	// Experimental.
	Scope() MonitoringScope
	// Experimental.
	ShortQueryDurationMetric() interface{}
	// Experimental.
	Title() *string
	// Experimental.
	Url() *string
	// Experimental.
	UsageAlarmFactory() UsageAlarmFactory
	// Experimental.
	UsageAnnotations() *[]*awscloudwatch.HorizontalAnnotation
	// Experimental.
	WriteLatencyMetric() interface{}
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
	CreateCpuAndDiskUsageWidget(width *float64, height *float64) awscloudwatch.GraphWidget
	// Returns all the alarms created.
	// Experimental.
	CreatedAlarms() *[]*AlarmWithAnnotation
	// Experimental.
	CreateLatencyWidget(width *float64, height *float64) awscloudwatch.GraphWidget
	// Experimental.
	CreateMaintenanceWidget(width *float64, height *float64) awscloudwatch.GraphWidget
	// Creates a new metric factory.
	// Experimental.
	CreateMetricFactory() MetricFactory
	// Experimental.
	CreateQueryDurationWidget(width *float64, height *float64) awscloudwatch.GraphWidget
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

// The jsii proxy struct for RedshiftClusterMonitoring
type jsiiProxy_RedshiftClusterMonitoring struct {
	jsiiProxy_Monitoring
}

func (j *jsiiProxy_RedshiftClusterMonitoring) Alarms() *[]*AlarmWithAnnotation {
	var returns *[]*AlarmWithAnnotation
	_jsii_.Get(
		j,
		"alarms",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedshiftClusterMonitoring) ConnectionAlarmFactory() ConnectionAlarmFactory {
	var returns ConnectionAlarmFactory
	_jsii_.Get(
		j,
		"connectionAlarmFactory",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedshiftClusterMonitoring) ConnectionAnnotations() *[]*awscloudwatch.HorizontalAnnotation {
	var returns *[]*awscloudwatch.HorizontalAnnotation
	_jsii_.Get(
		j,
		"connectionAnnotations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedshiftClusterMonitoring) ConnectionsMetric() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connectionsMetric",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedshiftClusterMonitoring) CpuUsageMetric() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"cpuUsageMetric",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedshiftClusterMonitoring) DiskSpaceUsageMetric() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"diskSpaceUsageMetric",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedshiftClusterMonitoring) LatencyAlarmFactory() LatencyAlarmFactory {
	var returns LatencyAlarmFactory
	_jsii_.Get(
		j,
		"latencyAlarmFactory",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedshiftClusterMonitoring) LocalAlarmNamePrefixOverride() *string {
	var returns *string
	_jsii_.Get(
		j,
		"localAlarmNamePrefixOverride",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedshiftClusterMonitoring) LongQueryDurationMetric() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"longQueryDurationMetric",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedshiftClusterMonitoring) MaintenanceModeMetric() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"maintenanceModeMetric",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedshiftClusterMonitoring) MediumQueryDurationMetric() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"mediumQueryDurationMetric",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedshiftClusterMonitoring) QueryDurationAnnotations() *[]*awscloudwatch.HorizontalAnnotation {
	var returns *[]*awscloudwatch.HorizontalAnnotation
	_jsii_.Get(
		j,
		"queryDurationAnnotations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedshiftClusterMonitoring) ReadLatencyMetric() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"readLatencyMetric",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedshiftClusterMonitoring) Scope() MonitoringScope {
	var returns MonitoringScope
	_jsii_.Get(
		j,
		"scope",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedshiftClusterMonitoring) ShortQueryDurationMetric() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"shortQueryDurationMetric",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedshiftClusterMonitoring) Title() *string {
	var returns *string
	_jsii_.Get(
		j,
		"title",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedshiftClusterMonitoring) Url() *string {
	var returns *string
	_jsii_.Get(
		j,
		"url",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedshiftClusterMonitoring) UsageAlarmFactory() UsageAlarmFactory {
	var returns UsageAlarmFactory
	_jsii_.Get(
		j,
		"usageAlarmFactory",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedshiftClusterMonitoring) UsageAnnotations() *[]*awscloudwatch.HorizontalAnnotation {
	var returns *[]*awscloudwatch.HorizontalAnnotation
	_jsii_.Get(
		j,
		"usageAnnotations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedshiftClusterMonitoring) WriteLatencyMetric() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"writeLatencyMetric",
		&returns,
	)
	return returns
}


// Experimental.
func NewRedshiftClusterMonitoring(scope MonitoringScope, props *RedshiftClusterMonitoringProps) RedshiftClusterMonitoring {
	_init_.Initialize()

	if err := validateNewRedshiftClusterMonitoringParameters(scope, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_RedshiftClusterMonitoring{}

	_jsii_.Create(
		"cdk-monitoring-constructs.RedshiftClusterMonitoring",
		[]interface{}{scope, props},
		&j,
	)

	return &j
}

// Experimental.
func NewRedshiftClusterMonitoring_Override(r RedshiftClusterMonitoring, scope MonitoringScope, props *RedshiftClusterMonitoringProps) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk-monitoring-constructs.RedshiftClusterMonitoring",
		[]interface{}{scope, props},
		r,
	)
}

func (r *jsiiProxy_RedshiftClusterMonitoring) AddAlarm(alarm *AlarmWithAnnotation) {
	if err := r.validateAddAlarmParameters(alarm); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"addAlarm",
		[]interface{}{alarm},
	)
}

func (r *jsiiProxy_RedshiftClusterMonitoring) AlarmWidgets() *[]awscloudwatch.IWidget {
	var returns *[]awscloudwatch.IWidget

	_jsii_.Invoke(
		r,
		"alarmWidgets",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RedshiftClusterMonitoring) CreateAlarmFactory(alarmNamePrefix *string) AlarmFactory {
	if err := r.validateCreateAlarmFactoryParameters(alarmNamePrefix); err != nil {
		panic(err)
	}
	var returns AlarmFactory

	_jsii_.Invoke(
		r,
		"createAlarmFactory",
		[]interface{}{alarmNamePrefix},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RedshiftClusterMonitoring) CreateConnectionsWidget(width *float64, height *float64) awscloudwatch.GraphWidget {
	if err := r.validateCreateConnectionsWidgetParameters(width, height); err != nil {
		panic(err)
	}
	var returns awscloudwatch.GraphWidget

	_jsii_.Invoke(
		r,
		"createConnectionsWidget",
		[]interface{}{width, height},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RedshiftClusterMonitoring) CreateCpuAndDiskUsageWidget(width *float64, height *float64) awscloudwatch.GraphWidget {
	if err := r.validateCreateCpuAndDiskUsageWidgetParameters(width, height); err != nil {
		panic(err)
	}
	var returns awscloudwatch.GraphWidget

	_jsii_.Invoke(
		r,
		"createCpuAndDiskUsageWidget",
		[]interface{}{width, height},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RedshiftClusterMonitoring) CreatedAlarms() *[]*AlarmWithAnnotation {
	var returns *[]*AlarmWithAnnotation

	_jsii_.Invoke(
		r,
		"createdAlarms",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RedshiftClusterMonitoring) CreateLatencyWidget(width *float64, height *float64) awscloudwatch.GraphWidget {
	if err := r.validateCreateLatencyWidgetParameters(width, height); err != nil {
		panic(err)
	}
	var returns awscloudwatch.GraphWidget

	_jsii_.Invoke(
		r,
		"createLatencyWidget",
		[]interface{}{width, height},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RedshiftClusterMonitoring) CreateMaintenanceWidget(width *float64, height *float64) awscloudwatch.GraphWidget {
	if err := r.validateCreateMaintenanceWidgetParameters(width, height); err != nil {
		panic(err)
	}
	var returns awscloudwatch.GraphWidget

	_jsii_.Invoke(
		r,
		"createMaintenanceWidget",
		[]interface{}{width, height},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RedshiftClusterMonitoring) CreateMetricFactory() MetricFactory {
	var returns MetricFactory

	_jsii_.Invoke(
		r,
		"createMetricFactory",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RedshiftClusterMonitoring) CreateQueryDurationWidget(width *float64, height *float64) awscloudwatch.GraphWidget {
	if err := r.validateCreateQueryDurationWidgetParameters(width, height); err != nil {
		panic(err)
	}
	var returns awscloudwatch.GraphWidget

	_jsii_.Invoke(
		r,
		"createQueryDurationWidget",
		[]interface{}{width, height},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RedshiftClusterMonitoring) CreateTitleWidget() MonitoringHeaderWidget {
	var returns MonitoringHeaderWidget

	_jsii_.Invoke(
		r,
		"createTitleWidget",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RedshiftClusterMonitoring) CreateWidgetFactory() IWidgetFactory {
	var returns IWidgetFactory

	_jsii_.Invoke(
		r,
		"createWidgetFactory",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RedshiftClusterMonitoring) SummaryWidgets() *[]awscloudwatch.IWidget {
	var returns *[]awscloudwatch.IWidget

	_jsii_.Invoke(
		r,
		"summaryWidgets",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RedshiftClusterMonitoring) Widgets() *[]awscloudwatch.IWidget {
	var returns *[]awscloudwatch.IWidget

	_jsii_.Invoke(
		r,
		"widgets",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RedshiftClusterMonitoring) WidgetsForDashboard(name *string) *[]awscloudwatch.IWidget {
	if err := r.validateWidgetsForDashboardParameters(name); err != nil {
		panic(err)
	}
	var returns *[]awscloudwatch.IWidget

	_jsii_.Invoke(
		r,
		"widgetsForDashboard",
		[]interface{}{name},
		&returns,
	)

	return returns
}


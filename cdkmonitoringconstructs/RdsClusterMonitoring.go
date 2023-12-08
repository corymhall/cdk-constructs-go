package cdkmonitoringconstructs

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/corymhall/cdk-constructs-go/cdkmonitoringconstructs/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudwatch"
)

// Experimental.
type RdsClusterMonitoring interface {
	Monitoring
	// Experimental.
	Alarms() *[]*AlarmWithAnnotation
	// Experimental.
	CommitLatencyMetric() interface{}
	// Experimental.
	ConnectionAlarmFactory() ConnectionAlarmFactory
	// Experimental.
	ConnectionAnnotations() *[]*awscloudwatch.HorizontalAnnotation
	// Experimental.
	ConnectionsMetric() interface{}
	// Experimental.
	CpuUsageMetric() interface{}
	// Experimental.
	DeleteLatencyMetric() interface{}
	// Experimental.
	DiskSpaceUsageMetric() interface{}
	// Experimental.
	InsertLatencyMetric() interface{}
	// Experimental.
	LocalAlarmNamePrefixOverride() *string
	// Experimental.
	Scope() MonitoringScope
	// Experimental.
	SelectLatencyMetric() interface{}
	// Experimental.
	Title() *string
	// Experimental.
	UpdateLatencyMetric() interface{}
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
	// Experimental.
	CreateConnectionsWidget(width *float64, height *float64) awscloudwatch.GraphWidget
	// Experimental.
	CreateCpuAndDiskUsageWidget(width *float64, height *float64) awscloudwatch.GraphWidget
	// Returns all the alarms created.
	// Experimental.
	CreatedAlarms() *[]*AlarmWithAnnotation
	// Experimental.
	CreateLatencyWidget(width *float64, height *float64) awscloudwatch.GraphWidget
	// Creates a new metric factory.
	// Experimental.
	CreateMetricFactory() MetricFactory
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

// The jsii proxy struct for RdsClusterMonitoring
type jsiiProxy_RdsClusterMonitoring struct {
	jsiiProxy_Monitoring
}

func (j *jsiiProxy_RdsClusterMonitoring) Alarms() *[]*AlarmWithAnnotation {
	var returns *[]*AlarmWithAnnotation
	_jsii_.Get(
		j,
		"alarms",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsClusterMonitoring) CommitLatencyMetric() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"commitLatencyMetric",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsClusterMonitoring) ConnectionAlarmFactory() ConnectionAlarmFactory {
	var returns ConnectionAlarmFactory
	_jsii_.Get(
		j,
		"connectionAlarmFactory",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsClusterMonitoring) ConnectionAnnotations() *[]*awscloudwatch.HorizontalAnnotation {
	var returns *[]*awscloudwatch.HorizontalAnnotation
	_jsii_.Get(
		j,
		"connectionAnnotations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsClusterMonitoring) ConnectionsMetric() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connectionsMetric",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsClusterMonitoring) CpuUsageMetric() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"cpuUsageMetric",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsClusterMonitoring) DeleteLatencyMetric() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"deleteLatencyMetric",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsClusterMonitoring) DiskSpaceUsageMetric() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"diskSpaceUsageMetric",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsClusterMonitoring) InsertLatencyMetric() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"insertLatencyMetric",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsClusterMonitoring) LocalAlarmNamePrefixOverride() *string {
	var returns *string
	_jsii_.Get(
		j,
		"localAlarmNamePrefixOverride",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsClusterMonitoring) Scope() MonitoringScope {
	var returns MonitoringScope
	_jsii_.Get(
		j,
		"scope",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsClusterMonitoring) SelectLatencyMetric() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"selectLatencyMetric",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsClusterMonitoring) Title() *string {
	var returns *string
	_jsii_.Get(
		j,
		"title",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsClusterMonitoring) UpdateLatencyMetric() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"updateLatencyMetric",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsClusterMonitoring) Url() *string {
	var returns *string
	_jsii_.Get(
		j,
		"url",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsClusterMonitoring) UsageAlarmFactory() UsageAlarmFactory {
	var returns UsageAlarmFactory
	_jsii_.Get(
		j,
		"usageAlarmFactory",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsClusterMonitoring) UsageAnnotations() *[]*awscloudwatch.HorizontalAnnotation {
	var returns *[]*awscloudwatch.HorizontalAnnotation
	_jsii_.Get(
		j,
		"usageAnnotations",
		&returns,
	)
	return returns
}


// Experimental.
func NewRdsClusterMonitoring(scope MonitoringScope, props *RdsClusterMonitoringProps) RdsClusterMonitoring {
	_init_.Initialize()

	if err := validateNewRdsClusterMonitoringParameters(scope, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_RdsClusterMonitoring{}

	_jsii_.Create(
		"cdk-monitoring-constructs.RdsClusterMonitoring",
		[]interface{}{scope, props},
		&j,
	)

	return &j
}

// Experimental.
func NewRdsClusterMonitoring_Override(r RdsClusterMonitoring, scope MonitoringScope, props *RdsClusterMonitoringProps) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk-monitoring-constructs.RdsClusterMonitoring",
		[]interface{}{scope, props},
		r,
	)
}

func (r *jsiiProxy_RdsClusterMonitoring) AddAlarm(alarm *AlarmWithAnnotation) {
	if err := r.validateAddAlarmParameters(alarm); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"addAlarm",
		[]interface{}{alarm},
	)
}

func (r *jsiiProxy_RdsClusterMonitoring) AlarmWidgets() *[]awscloudwatch.IWidget {
	var returns *[]awscloudwatch.IWidget

	_jsii_.Invoke(
		r,
		"alarmWidgets",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RdsClusterMonitoring) CreateAlarmFactory(alarmNamePrefix *string) AlarmFactory {
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

func (r *jsiiProxy_RdsClusterMonitoring) CreateConnectionsWidget(width *float64, height *float64) awscloudwatch.GraphWidget {
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

func (r *jsiiProxy_RdsClusterMonitoring) CreateCpuAndDiskUsageWidget(width *float64, height *float64) awscloudwatch.GraphWidget {
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

func (r *jsiiProxy_RdsClusterMonitoring) CreatedAlarms() *[]*AlarmWithAnnotation {
	var returns *[]*AlarmWithAnnotation

	_jsii_.Invoke(
		r,
		"createdAlarms",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RdsClusterMonitoring) CreateLatencyWidget(width *float64, height *float64) awscloudwatch.GraphWidget {
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

func (r *jsiiProxy_RdsClusterMonitoring) CreateMetricFactory() MetricFactory {
	var returns MetricFactory

	_jsii_.Invoke(
		r,
		"createMetricFactory",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RdsClusterMonitoring) CreateTitleWidget() MonitoringHeaderWidget {
	var returns MonitoringHeaderWidget

	_jsii_.Invoke(
		r,
		"createTitleWidget",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RdsClusterMonitoring) CreateWidgetFactory() IWidgetFactory {
	var returns IWidgetFactory

	_jsii_.Invoke(
		r,
		"createWidgetFactory",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RdsClusterMonitoring) SummaryWidgets() *[]awscloudwatch.IWidget {
	var returns *[]awscloudwatch.IWidget

	_jsii_.Invoke(
		r,
		"summaryWidgets",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RdsClusterMonitoring) Widgets() *[]awscloudwatch.IWidget {
	var returns *[]awscloudwatch.IWidget

	_jsii_.Invoke(
		r,
		"widgets",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RdsClusterMonitoring) WidgetsForDashboard(name *string) *[]awscloudwatch.IWidget {
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


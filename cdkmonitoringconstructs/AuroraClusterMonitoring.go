package cdkmonitoringconstructs

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/corymhall/cdk-constructs-go/cdkmonitoringconstructs/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudwatch"
)

// Experimental.
type AuroraClusterMonitoring interface {
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
	InsertLatencyMetric() interface{}
	// Experimental.
	LocalAlarmNamePrefixOverride() *string
	// Experimental.
	Scope() MonitoringScope
	// Experimental.
	SelectLatencyMetric() interface{}
	// Experimental.
	ServerlessCapacityAnnotations() *[]*awscloudwatch.HorizontalAnnotation
	// Experimental.
	ServerlessDatabaseCapacityMetric() interface{}
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
	CreateCpuUsageWidget(width *float64, height *float64) awscloudwatch.GraphWidget
	// Returns all the alarms created.
	// Experimental.
	CreatedAlarms() *[]*AlarmWithAnnotation
	// Experimental.
	CreateLatencyWidget(width *float64, height *float64) awscloudwatch.GraphWidget
	// Creates a new metric factory.
	// Experimental.
	CreateMetricFactory() MetricFactory
	// Experimental.
	CreateServerlessDatabaseCapacityWidget(width *float64, height *float64) awscloudwatch.GraphWidget
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

// The jsii proxy struct for AuroraClusterMonitoring
type jsiiProxy_AuroraClusterMonitoring struct {
	jsiiProxy_Monitoring
}

func (j *jsiiProxy_AuroraClusterMonitoring) Alarms() *[]*AlarmWithAnnotation {
	var returns *[]*AlarmWithAnnotation
	_jsii_.Get(
		j,
		"alarms",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AuroraClusterMonitoring) CommitLatencyMetric() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"commitLatencyMetric",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AuroraClusterMonitoring) ConnectionAlarmFactory() ConnectionAlarmFactory {
	var returns ConnectionAlarmFactory
	_jsii_.Get(
		j,
		"connectionAlarmFactory",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AuroraClusterMonitoring) ConnectionAnnotations() *[]*awscloudwatch.HorizontalAnnotation {
	var returns *[]*awscloudwatch.HorizontalAnnotation
	_jsii_.Get(
		j,
		"connectionAnnotations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AuroraClusterMonitoring) ConnectionsMetric() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connectionsMetric",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AuroraClusterMonitoring) CpuUsageMetric() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"cpuUsageMetric",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AuroraClusterMonitoring) DeleteLatencyMetric() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"deleteLatencyMetric",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AuroraClusterMonitoring) InsertLatencyMetric() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"insertLatencyMetric",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AuroraClusterMonitoring) LocalAlarmNamePrefixOverride() *string {
	var returns *string
	_jsii_.Get(
		j,
		"localAlarmNamePrefixOverride",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AuroraClusterMonitoring) Scope() MonitoringScope {
	var returns MonitoringScope
	_jsii_.Get(
		j,
		"scope",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AuroraClusterMonitoring) SelectLatencyMetric() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"selectLatencyMetric",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AuroraClusterMonitoring) ServerlessCapacityAnnotations() *[]*awscloudwatch.HorizontalAnnotation {
	var returns *[]*awscloudwatch.HorizontalAnnotation
	_jsii_.Get(
		j,
		"serverlessCapacityAnnotations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AuroraClusterMonitoring) ServerlessDatabaseCapacityMetric() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"serverlessDatabaseCapacityMetric",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AuroraClusterMonitoring) Title() *string {
	var returns *string
	_jsii_.Get(
		j,
		"title",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AuroraClusterMonitoring) UpdateLatencyMetric() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"updateLatencyMetric",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AuroraClusterMonitoring) Url() *string {
	var returns *string
	_jsii_.Get(
		j,
		"url",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AuroraClusterMonitoring) UsageAlarmFactory() UsageAlarmFactory {
	var returns UsageAlarmFactory
	_jsii_.Get(
		j,
		"usageAlarmFactory",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AuroraClusterMonitoring) UsageAnnotations() *[]*awscloudwatch.HorizontalAnnotation {
	var returns *[]*awscloudwatch.HorizontalAnnotation
	_jsii_.Get(
		j,
		"usageAnnotations",
		&returns,
	)
	return returns
}


// Experimental.
func NewAuroraClusterMonitoring(scope MonitoringScope, props *AuroraClusterMonitoringProps) AuroraClusterMonitoring {
	_init_.Initialize()

	if err := validateNewAuroraClusterMonitoringParameters(scope, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_AuroraClusterMonitoring{}

	_jsii_.Create(
		"cdk-monitoring-constructs.AuroraClusterMonitoring",
		[]interface{}{scope, props},
		&j,
	)

	return &j
}

// Experimental.
func NewAuroraClusterMonitoring_Override(a AuroraClusterMonitoring, scope MonitoringScope, props *AuroraClusterMonitoringProps) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk-monitoring-constructs.AuroraClusterMonitoring",
		[]interface{}{scope, props},
		a,
	)
}

func (a *jsiiProxy_AuroraClusterMonitoring) AddAlarm(alarm *AlarmWithAnnotation) {
	if err := a.validateAddAlarmParameters(alarm); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addAlarm",
		[]interface{}{alarm},
	)
}

func (a *jsiiProxy_AuroraClusterMonitoring) AlarmWidgets() *[]awscloudwatch.IWidget {
	var returns *[]awscloudwatch.IWidget

	_jsii_.Invoke(
		a,
		"alarmWidgets",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AuroraClusterMonitoring) CreateAlarmFactory(alarmNamePrefix *string) AlarmFactory {
	if err := a.validateCreateAlarmFactoryParameters(alarmNamePrefix); err != nil {
		panic(err)
	}
	var returns AlarmFactory

	_jsii_.Invoke(
		a,
		"createAlarmFactory",
		[]interface{}{alarmNamePrefix},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AuroraClusterMonitoring) CreateConnectionsWidget(width *float64, height *float64) awscloudwatch.GraphWidget {
	if err := a.validateCreateConnectionsWidgetParameters(width, height); err != nil {
		panic(err)
	}
	var returns awscloudwatch.GraphWidget

	_jsii_.Invoke(
		a,
		"createConnectionsWidget",
		[]interface{}{width, height},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AuroraClusterMonitoring) CreateCpuUsageWidget(width *float64, height *float64) awscloudwatch.GraphWidget {
	if err := a.validateCreateCpuUsageWidgetParameters(width, height); err != nil {
		panic(err)
	}
	var returns awscloudwatch.GraphWidget

	_jsii_.Invoke(
		a,
		"createCpuUsageWidget",
		[]interface{}{width, height},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AuroraClusterMonitoring) CreatedAlarms() *[]*AlarmWithAnnotation {
	var returns *[]*AlarmWithAnnotation

	_jsii_.Invoke(
		a,
		"createdAlarms",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AuroraClusterMonitoring) CreateLatencyWidget(width *float64, height *float64) awscloudwatch.GraphWidget {
	if err := a.validateCreateLatencyWidgetParameters(width, height); err != nil {
		panic(err)
	}
	var returns awscloudwatch.GraphWidget

	_jsii_.Invoke(
		a,
		"createLatencyWidget",
		[]interface{}{width, height},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AuroraClusterMonitoring) CreateMetricFactory() MetricFactory {
	var returns MetricFactory

	_jsii_.Invoke(
		a,
		"createMetricFactory",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AuroraClusterMonitoring) CreateServerlessDatabaseCapacityWidget(width *float64, height *float64) awscloudwatch.GraphWidget {
	if err := a.validateCreateServerlessDatabaseCapacityWidgetParameters(width, height); err != nil {
		panic(err)
	}
	var returns awscloudwatch.GraphWidget

	_jsii_.Invoke(
		a,
		"createServerlessDatabaseCapacityWidget",
		[]interface{}{width, height},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AuroraClusterMonitoring) CreateTitleWidget() MonitoringHeaderWidget {
	var returns MonitoringHeaderWidget

	_jsii_.Invoke(
		a,
		"createTitleWidget",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AuroraClusterMonitoring) CreateWidgetFactory() IWidgetFactory {
	var returns IWidgetFactory

	_jsii_.Invoke(
		a,
		"createWidgetFactory",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AuroraClusterMonitoring) SummaryWidgets() *[]awscloudwatch.IWidget {
	var returns *[]awscloudwatch.IWidget

	_jsii_.Invoke(
		a,
		"summaryWidgets",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AuroraClusterMonitoring) Widgets() *[]awscloudwatch.IWidget {
	var returns *[]awscloudwatch.IWidget

	_jsii_.Invoke(
		a,
		"widgets",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AuroraClusterMonitoring) WidgetsForDashboard(name *string) *[]awscloudwatch.IWidget {
	if err := a.validateWidgetsForDashboardParameters(name); err != nil {
		panic(err)
	}
	var returns *[]awscloudwatch.IWidget

	_jsii_.Invoke(
		a,
		"widgetsForDashboard",
		[]interface{}{name},
		&returns,
	)

	return returns
}


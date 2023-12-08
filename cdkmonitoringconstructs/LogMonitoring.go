package cdkmonitoringconstructs

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/corymhall/cdk-constructs-go/cdkmonitoringconstructs/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudwatch"
)

// Monitors a CloudWatch log group for various patterns.
// Experimental.
type LogMonitoring interface {
	Monitoring
	// Experimental.
	AlarmFactory() AlarmFactory
	// Experimental.
	Alarms() *[]*AlarmWithAnnotation
	// Experimental.
	IncomingLogEventsMetric() interface{}
	// Experimental.
	Limit() *float64
	// Experimental.
	LocalAlarmNamePrefixOverride() *string
	// Experimental.
	LogGroupName() *string
	// Experimental.
	LogGroupUrl() *string
	// Experimental.
	Pattern() *string
	// Experimental.
	Scope() MonitoringScope
	// Experimental.
	Title() *string
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
	// Experimental.
	CreateIncomingLogsWidget(width *float64, height *float64) awscloudwatch.GraphWidget
	// Creates a new metric factory.
	// Experimental.
	CreateMetricFactory() MetricFactory
	// Experimental.
	CreateTitleWidget() MonitoringHeaderWidget
	// Creates a new widget factory.
	// Experimental.
	CreateWidgetFactory() IWidgetFactory
	// Experimental.
	ResolveRecommendedHeight(numRows *float64) *float64
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

// The jsii proxy struct for LogMonitoring
type jsiiProxy_LogMonitoring struct {
	jsiiProxy_Monitoring
}

func (j *jsiiProxy_LogMonitoring) AlarmFactory() AlarmFactory {
	var returns AlarmFactory
	_jsii_.Get(
		j,
		"alarmFactory",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogMonitoring) Alarms() *[]*AlarmWithAnnotation {
	var returns *[]*AlarmWithAnnotation
	_jsii_.Get(
		j,
		"alarms",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogMonitoring) IncomingLogEventsMetric() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"incomingLogEventsMetric",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogMonitoring) Limit() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"limit",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogMonitoring) LocalAlarmNamePrefixOverride() *string {
	var returns *string
	_jsii_.Get(
		j,
		"localAlarmNamePrefixOverride",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogMonitoring) LogGroupName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logGroupName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogMonitoring) LogGroupUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logGroupUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogMonitoring) Pattern() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pattern",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogMonitoring) Scope() MonitoringScope {
	var returns MonitoringScope
	_jsii_.Get(
		j,
		"scope",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogMonitoring) Title() *string {
	var returns *string
	_jsii_.Get(
		j,
		"title",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogMonitoring) UsageAlarmFactory() UsageAlarmFactory {
	var returns UsageAlarmFactory
	_jsii_.Get(
		j,
		"usageAlarmFactory",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogMonitoring) UsageAnnotations() *[]*awscloudwatch.HorizontalAnnotation {
	var returns *[]*awscloudwatch.HorizontalAnnotation
	_jsii_.Get(
		j,
		"usageAnnotations",
		&returns,
	)
	return returns
}


// Experimental.
func NewLogMonitoring(scope MonitoringScope, props *LogMonitoringProps) LogMonitoring {
	_init_.Initialize()

	if err := validateNewLogMonitoringParameters(scope, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_LogMonitoring{}

	_jsii_.Create(
		"cdk-monitoring-constructs.LogMonitoring",
		[]interface{}{scope, props},
		&j,
	)

	return &j
}

// Experimental.
func NewLogMonitoring_Override(l LogMonitoring, scope MonitoringScope, props *LogMonitoringProps) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk-monitoring-constructs.LogMonitoring",
		[]interface{}{scope, props},
		l,
	)
}

func (l *jsiiProxy_LogMonitoring) AddAlarm(alarm *AlarmWithAnnotation) {
	if err := l.validateAddAlarmParameters(alarm); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"addAlarm",
		[]interface{}{alarm},
	)
}

func (l *jsiiProxy_LogMonitoring) AlarmWidgets() *[]awscloudwatch.IWidget {
	var returns *[]awscloudwatch.IWidget

	_jsii_.Invoke(
		l,
		"alarmWidgets",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LogMonitoring) CreateAlarmFactory(alarmNamePrefix *string) AlarmFactory {
	if err := l.validateCreateAlarmFactoryParameters(alarmNamePrefix); err != nil {
		panic(err)
	}
	var returns AlarmFactory

	_jsii_.Invoke(
		l,
		"createAlarmFactory",
		[]interface{}{alarmNamePrefix},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LogMonitoring) CreatedAlarms() *[]*AlarmWithAnnotation {
	var returns *[]*AlarmWithAnnotation

	_jsii_.Invoke(
		l,
		"createdAlarms",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LogMonitoring) CreateIncomingLogsWidget(width *float64, height *float64) awscloudwatch.GraphWidget {
	if err := l.validateCreateIncomingLogsWidgetParameters(width, height); err != nil {
		panic(err)
	}
	var returns awscloudwatch.GraphWidget

	_jsii_.Invoke(
		l,
		"createIncomingLogsWidget",
		[]interface{}{width, height},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LogMonitoring) CreateMetricFactory() MetricFactory {
	var returns MetricFactory

	_jsii_.Invoke(
		l,
		"createMetricFactory",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LogMonitoring) CreateTitleWidget() MonitoringHeaderWidget {
	var returns MonitoringHeaderWidget

	_jsii_.Invoke(
		l,
		"createTitleWidget",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LogMonitoring) CreateWidgetFactory() IWidgetFactory {
	var returns IWidgetFactory

	_jsii_.Invoke(
		l,
		"createWidgetFactory",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LogMonitoring) ResolveRecommendedHeight(numRows *float64) *float64 {
	if err := l.validateResolveRecommendedHeightParameters(numRows); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		l,
		"resolveRecommendedHeight",
		[]interface{}{numRows},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LogMonitoring) SummaryWidgets() *[]awscloudwatch.IWidget {
	var returns *[]awscloudwatch.IWidget

	_jsii_.Invoke(
		l,
		"summaryWidgets",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LogMonitoring) Widgets() *[]awscloudwatch.IWidget {
	var returns *[]awscloudwatch.IWidget

	_jsii_.Invoke(
		l,
		"widgets",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LogMonitoring) WidgetsForDashboard(name *string) *[]awscloudwatch.IWidget {
	if err := l.validateWidgetsForDashboardParameters(name); err != nil {
		panic(err)
	}
	var returns *[]awscloudwatch.IWidget

	_jsii_.Invoke(
		l,
		"widgetsForDashboard",
		[]interface{}{name},
		&returns,
	)

	return returns
}


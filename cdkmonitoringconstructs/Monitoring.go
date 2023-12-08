package cdkmonitoringconstructs

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/corymhall/cdk-constructs-go/cdkmonitoringconstructs/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudwatch"
)

// An independent unit of monitoring.
//
// This is the base for all monitoring classes with alarm support.
// Experimental.
type Monitoring interface {
	IDashboardSegment
	IDynamicDashboardSegment
	// Experimental.
	Alarms() *[]*AlarmWithAnnotation
	// Experimental.
	LocalAlarmNamePrefixOverride() *string
	// Experimental.
	Scope() MonitoringScope
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
	// Default: - no widgets.
	//
	// Experimental.
	SummaryWidgets() *[]awscloudwatch.IWidget
	// Returns widgets to be placed on the main dashboard.
	// Experimental.
	Widgets() *[]awscloudwatch.IWidget
	// Returns widgets for the requested dashboard type.
	// Experimental.
	WidgetsForDashboard(name *string) *[]awscloudwatch.IWidget
}

// The jsii proxy struct for Monitoring
type jsiiProxy_Monitoring struct {
	jsiiProxy_IDashboardSegment
	jsiiProxy_IDynamicDashboardSegment
}

func (j *jsiiProxy_Monitoring) Alarms() *[]*AlarmWithAnnotation {
	var returns *[]*AlarmWithAnnotation
	_jsii_.Get(
		j,
		"alarms",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Monitoring) LocalAlarmNamePrefixOverride() *string {
	var returns *string
	_jsii_.Get(
		j,
		"localAlarmNamePrefixOverride",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Monitoring) Scope() MonitoringScope {
	var returns MonitoringScope
	_jsii_.Get(
		j,
		"scope",
		&returns,
	)
	return returns
}


// Experimental.
func NewMonitoring_Override(m Monitoring, scope MonitoringScope, props *BaseMonitoringProps) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk-monitoring-constructs.Monitoring",
		[]interface{}{scope, props},
		m,
	)
}

func (m *jsiiProxy_Monitoring) AddAlarm(alarm *AlarmWithAnnotation) {
	if err := m.validateAddAlarmParameters(alarm); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"addAlarm",
		[]interface{}{alarm},
	)
}

func (m *jsiiProxy_Monitoring) AlarmWidgets() *[]awscloudwatch.IWidget {
	var returns *[]awscloudwatch.IWidget

	_jsii_.Invoke(
		m,
		"alarmWidgets",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_Monitoring) CreateAlarmFactory(alarmNamePrefix *string) AlarmFactory {
	if err := m.validateCreateAlarmFactoryParameters(alarmNamePrefix); err != nil {
		panic(err)
	}
	var returns AlarmFactory

	_jsii_.Invoke(
		m,
		"createAlarmFactory",
		[]interface{}{alarmNamePrefix},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_Monitoring) CreatedAlarms() *[]*AlarmWithAnnotation {
	var returns *[]*AlarmWithAnnotation

	_jsii_.Invoke(
		m,
		"createdAlarms",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_Monitoring) CreateMetricFactory() MetricFactory {
	var returns MetricFactory

	_jsii_.Invoke(
		m,
		"createMetricFactory",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_Monitoring) CreateWidgetFactory() IWidgetFactory {
	var returns IWidgetFactory

	_jsii_.Invoke(
		m,
		"createWidgetFactory",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_Monitoring) SummaryWidgets() *[]awscloudwatch.IWidget {
	var returns *[]awscloudwatch.IWidget

	_jsii_.Invoke(
		m,
		"summaryWidgets",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_Monitoring) Widgets() *[]awscloudwatch.IWidget {
	var returns *[]awscloudwatch.IWidget

	_jsii_.Invoke(
		m,
		"widgets",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_Monitoring) WidgetsForDashboard(name *string) *[]awscloudwatch.IWidget {
	if err := m.validateWidgetsForDashboardParameters(name); err != nil {
		panic(err)
	}
	var returns *[]awscloudwatch.IWidget

	_jsii_.Invoke(
		m,
		"widgetsForDashboard",
		[]interface{}{name},
		&returns,
	)

	return returns
}


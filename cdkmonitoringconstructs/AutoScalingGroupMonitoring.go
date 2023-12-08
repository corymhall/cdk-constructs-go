package cdkmonitoringconstructs

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/corymhall/cdk-constructs-go/cdkmonitoringconstructs/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudwatch"
)

// Experimental.
type AutoScalingGroupMonitoring interface {
	Monitoring
	// Experimental.
	Alarms() *[]*AlarmWithAnnotation
	// Experimental.
	GroupDesiredSizeMetric() interface{}
	// Experimental.
	GroupMaxSizeMetric() interface{}
	// Experimental.
	GroupMinSizeMetric() interface{}
	// Experimental.
	InstancesInServiceMetric() interface{}
	// Experimental.
	InstancesPendingMetric() interface{}
	// Experimental.
	InstancesStandbyMetric() interface{}
	// Experimental.
	InstancesTerminatingMetric() interface{}
	// Experimental.
	InstancesTotalMetric() interface{}
	// Experimental.
	LocalAlarmNamePrefixOverride() *string
	// Experimental.
	Scope() MonitoringScope
	// Experimental.
	Title() *string
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
	CreateGroupSizeWidget(width *float64, height *float64) awscloudwatch.GraphWidget
	// Experimental.
	CreateGroupStatusWidget(width *float64, height *float64) awscloudwatch.GraphWidget
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

// The jsii proxy struct for AutoScalingGroupMonitoring
type jsiiProxy_AutoScalingGroupMonitoring struct {
	jsiiProxy_Monitoring
}

func (j *jsiiProxy_AutoScalingGroupMonitoring) Alarms() *[]*AlarmWithAnnotation {
	var returns *[]*AlarmWithAnnotation
	_jsii_.Get(
		j,
		"alarms",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoScalingGroupMonitoring) GroupDesiredSizeMetric() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"groupDesiredSizeMetric",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoScalingGroupMonitoring) GroupMaxSizeMetric() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"groupMaxSizeMetric",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoScalingGroupMonitoring) GroupMinSizeMetric() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"groupMinSizeMetric",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoScalingGroupMonitoring) InstancesInServiceMetric() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"instancesInServiceMetric",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoScalingGroupMonitoring) InstancesPendingMetric() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"instancesPendingMetric",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoScalingGroupMonitoring) InstancesStandbyMetric() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"instancesStandbyMetric",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoScalingGroupMonitoring) InstancesTerminatingMetric() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"instancesTerminatingMetric",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoScalingGroupMonitoring) InstancesTotalMetric() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"instancesTotalMetric",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoScalingGroupMonitoring) LocalAlarmNamePrefixOverride() *string {
	var returns *string
	_jsii_.Get(
		j,
		"localAlarmNamePrefixOverride",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoScalingGroupMonitoring) Scope() MonitoringScope {
	var returns MonitoringScope
	_jsii_.Get(
		j,
		"scope",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoScalingGroupMonitoring) Title() *string {
	var returns *string
	_jsii_.Get(
		j,
		"title",
		&returns,
	)
	return returns
}


// Experimental.
func NewAutoScalingGroupMonitoring(scope MonitoringScope, props *AutoScalingGroupMonitoringProps) AutoScalingGroupMonitoring {
	_init_.Initialize()

	if err := validateNewAutoScalingGroupMonitoringParameters(scope, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_AutoScalingGroupMonitoring{}

	_jsii_.Create(
		"cdk-monitoring-constructs.AutoScalingGroupMonitoring",
		[]interface{}{scope, props},
		&j,
	)

	return &j
}

// Experimental.
func NewAutoScalingGroupMonitoring_Override(a AutoScalingGroupMonitoring, scope MonitoringScope, props *AutoScalingGroupMonitoringProps) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk-monitoring-constructs.AutoScalingGroupMonitoring",
		[]interface{}{scope, props},
		a,
	)
}

func (a *jsiiProxy_AutoScalingGroupMonitoring) AddAlarm(alarm *AlarmWithAnnotation) {
	if err := a.validateAddAlarmParameters(alarm); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addAlarm",
		[]interface{}{alarm},
	)
}

func (a *jsiiProxy_AutoScalingGroupMonitoring) AlarmWidgets() *[]awscloudwatch.IWidget {
	var returns *[]awscloudwatch.IWidget

	_jsii_.Invoke(
		a,
		"alarmWidgets",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutoScalingGroupMonitoring) CreateAlarmFactory(alarmNamePrefix *string) AlarmFactory {
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

func (a *jsiiProxy_AutoScalingGroupMonitoring) CreatedAlarms() *[]*AlarmWithAnnotation {
	var returns *[]*AlarmWithAnnotation

	_jsii_.Invoke(
		a,
		"createdAlarms",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutoScalingGroupMonitoring) CreateGroupSizeWidget(width *float64, height *float64) awscloudwatch.GraphWidget {
	if err := a.validateCreateGroupSizeWidgetParameters(width, height); err != nil {
		panic(err)
	}
	var returns awscloudwatch.GraphWidget

	_jsii_.Invoke(
		a,
		"createGroupSizeWidget",
		[]interface{}{width, height},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutoScalingGroupMonitoring) CreateGroupStatusWidget(width *float64, height *float64) awscloudwatch.GraphWidget {
	if err := a.validateCreateGroupStatusWidgetParameters(width, height); err != nil {
		panic(err)
	}
	var returns awscloudwatch.GraphWidget

	_jsii_.Invoke(
		a,
		"createGroupStatusWidget",
		[]interface{}{width, height},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutoScalingGroupMonitoring) CreateMetricFactory() MetricFactory {
	var returns MetricFactory

	_jsii_.Invoke(
		a,
		"createMetricFactory",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutoScalingGroupMonitoring) CreateTitleWidget() MonitoringHeaderWidget {
	var returns MonitoringHeaderWidget

	_jsii_.Invoke(
		a,
		"createTitleWidget",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutoScalingGroupMonitoring) CreateWidgetFactory() IWidgetFactory {
	var returns IWidgetFactory

	_jsii_.Invoke(
		a,
		"createWidgetFactory",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutoScalingGroupMonitoring) SummaryWidgets() *[]awscloudwatch.IWidget {
	var returns *[]awscloudwatch.IWidget

	_jsii_.Invoke(
		a,
		"summaryWidgets",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutoScalingGroupMonitoring) Widgets() *[]awscloudwatch.IWidget {
	var returns *[]awscloudwatch.IWidget

	_jsii_.Invoke(
		a,
		"widgets",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutoScalingGroupMonitoring) WidgetsForDashboard(name *string) *[]awscloudwatch.IWidget {
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


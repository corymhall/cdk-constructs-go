package cdkmonitoringconstructs

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/corymhall/cdk-constructs-go/cdkmonitoringconstructs/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudwatch"
)

// Experimental.
type StepFunctionActivityMonitoring interface {
	Monitoring
	// Experimental.
	Alarms() *[]*AlarmWithAnnotation
	// Experimental.
	DurationAlarmFactory() LatencyAlarmFactory
	// Experimental.
	DurationAnnotations() *[]*awscloudwatch.HorizontalAnnotation
	// Experimental.
	ErrorAlarmFactory() ErrorAlarmFactory
	// Experimental.
	ErrorCountAnnotations() *[]*awscloudwatch.HorizontalAnnotation
	// Experimental.
	ErrorRateAnnotations() *[]*awscloudwatch.HorizontalAnnotation
	// Experimental.
	FailedActivitiesMetric() interface{}
	// Experimental.
	FailedActivitiesRateMetric() interface{}
	// Experimental.
	HeartbeatTimedOutActivitiesMetrics() interface{}
	// Experimental.
	LocalAlarmNamePrefixOverride() *string
	// Experimental.
	P50DurationMetric() interface{}
	// Experimental.
	P90DurationMetric() interface{}
	// Experimental.
	P99DurationMetric() interface{}
	// Experimental.
	ScheduledActivitiesMetric() interface{}
	// Experimental.
	Scope() MonitoringScope
	// Experimental.
	StartedActivitiesMetric() interface{}
	// Experimental.
	SucceededActivitiesMetric() interface{}
	// Experimental.
	TimedOutActivitiesMetrics() interface{}
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

// The jsii proxy struct for StepFunctionActivityMonitoring
type jsiiProxy_StepFunctionActivityMonitoring struct {
	jsiiProxy_Monitoring
}

func (j *jsiiProxy_StepFunctionActivityMonitoring) Alarms() *[]*AlarmWithAnnotation {
	var returns *[]*AlarmWithAnnotation
	_jsii_.Get(
		j,
		"alarms",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StepFunctionActivityMonitoring) DurationAlarmFactory() LatencyAlarmFactory {
	var returns LatencyAlarmFactory
	_jsii_.Get(
		j,
		"durationAlarmFactory",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StepFunctionActivityMonitoring) DurationAnnotations() *[]*awscloudwatch.HorizontalAnnotation {
	var returns *[]*awscloudwatch.HorizontalAnnotation
	_jsii_.Get(
		j,
		"durationAnnotations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StepFunctionActivityMonitoring) ErrorAlarmFactory() ErrorAlarmFactory {
	var returns ErrorAlarmFactory
	_jsii_.Get(
		j,
		"errorAlarmFactory",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StepFunctionActivityMonitoring) ErrorCountAnnotations() *[]*awscloudwatch.HorizontalAnnotation {
	var returns *[]*awscloudwatch.HorizontalAnnotation
	_jsii_.Get(
		j,
		"errorCountAnnotations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StepFunctionActivityMonitoring) ErrorRateAnnotations() *[]*awscloudwatch.HorizontalAnnotation {
	var returns *[]*awscloudwatch.HorizontalAnnotation
	_jsii_.Get(
		j,
		"errorRateAnnotations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StepFunctionActivityMonitoring) FailedActivitiesMetric() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"failedActivitiesMetric",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StepFunctionActivityMonitoring) FailedActivitiesRateMetric() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"failedActivitiesRateMetric",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StepFunctionActivityMonitoring) HeartbeatTimedOutActivitiesMetrics() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"heartbeatTimedOutActivitiesMetrics",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StepFunctionActivityMonitoring) LocalAlarmNamePrefixOverride() *string {
	var returns *string
	_jsii_.Get(
		j,
		"localAlarmNamePrefixOverride",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StepFunctionActivityMonitoring) P50DurationMetric() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"p50DurationMetric",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StepFunctionActivityMonitoring) P90DurationMetric() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"p90DurationMetric",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StepFunctionActivityMonitoring) P99DurationMetric() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"p99DurationMetric",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StepFunctionActivityMonitoring) ScheduledActivitiesMetric() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"scheduledActivitiesMetric",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StepFunctionActivityMonitoring) Scope() MonitoringScope {
	var returns MonitoringScope
	_jsii_.Get(
		j,
		"scope",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StepFunctionActivityMonitoring) StartedActivitiesMetric() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"startedActivitiesMetric",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StepFunctionActivityMonitoring) SucceededActivitiesMetric() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"succeededActivitiesMetric",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StepFunctionActivityMonitoring) TimedOutActivitiesMetrics() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timedOutActivitiesMetrics",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StepFunctionActivityMonitoring) Title() *string {
	var returns *string
	_jsii_.Get(
		j,
		"title",
		&returns,
	)
	return returns
}


// Experimental.
func NewStepFunctionActivityMonitoring(scope MonitoringScope, props *StepFunctionActivityMonitoringProps) StepFunctionActivityMonitoring {
	_init_.Initialize()

	if err := validateNewStepFunctionActivityMonitoringParameters(scope, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_StepFunctionActivityMonitoring{}

	_jsii_.Create(
		"cdk-monitoring-constructs.StepFunctionActivityMonitoring",
		[]interface{}{scope, props},
		&j,
	)

	return &j
}

// Experimental.
func NewStepFunctionActivityMonitoring_Override(s StepFunctionActivityMonitoring, scope MonitoringScope, props *StepFunctionActivityMonitoringProps) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk-monitoring-constructs.StepFunctionActivityMonitoring",
		[]interface{}{scope, props},
		s,
	)
}

func (s *jsiiProxy_StepFunctionActivityMonitoring) AddAlarm(alarm *AlarmWithAnnotation) {
	if err := s.validateAddAlarmParameters(alarm); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"addAlarm",
		[]interface{}{alarm},
	)
}

func (s *jsiiProxy_StepFunctionActivityMonitoring) AlarmWidgets() *[]awscloudwatch.IWidget {
	var returns *[]awscloudwatch.IWidget

	_jsii_.Invoke(
		s,
		"alarmWidgets",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StepFunctionActivityMonitoring) CreateAlarmFactory(alarmNamePrefix *string) AlarmFactory {
	if err := s.validateCreateAlarmFactoryParameters(alarmNamePrefix); err != nil {
		panic(err)
	}
	var returns AlarmFactory

	_jsii_.Invoke(
		s,
		"createAlarmFactory",
		[]interface{}{alarmNamePrefix},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StepFunctionActivityMonitoring) CreatedAlarms() *[]*AlarmWithAnnotation {
	var returns *[]*AlarmWithAnnotation

	_jsii_.Invoke(
		s,
		"createdAlarms",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StepFunctionActivityMonitoring) CreateMetricFactory() MetricFactory {
	var returns MetricFactory

	_jsii_.Invoke(
		s,
		"createMetricFactory",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StepFunctionActivityMonitoring) CreateWidgetFactory() IWidgetFactory {
	var returns IWidgetFactory

	_jsii_.Invoke(
		s,
		"createWidgetFactory",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StepFunctionActivityMonitoring) SummaryWidgets() *[]awscloudwatch.IWidget {
	var returns *[]awscloudwatch.IWidget

	_jsii_.Invoke(
		s,
		"summaryWidgets",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StepFunctionActivityMonitoring) Widgets() *[]awscloudwatch.IWidget {
	var returns *[]awscloudwatch.IWidget

	_jsii_.Invoke(
		s,
		"widgets",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StepFunctionActivityMonitoring) WidgetsForDashboard(name *string) *[]awscloudwatch.IWidget {
	if err := s.validateWidgetsForDashboardParameters(name); err != nil {
		panic(err)
	}
	var returns *[]awscloudwatch.IWidget

	_jsii_.Invoke(
		s,
		"widgetsForDashboard",
		[]interface{}{name},
		&returns,
	)

	return returns
}


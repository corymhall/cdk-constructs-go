package cdkmonitoringconstructs

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/corymhall/cdk-constructs-go/cdkmonitoringconstructs/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudwatch"
)

// Experimental.
type StepFunctionMonitoring interface {
	Monitoring
	// Experimental.
	AbortedExecutionsMetric() interface{}
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
	FailedExecutionRateMetric() interface{}
	// Experimental.
	FailedExecutionsMetric() interface{}
	// Experimental.
	LocalAlarmNamePrefixOverride() *string
	// Experimental.
	P50DurationMetric() interface{}
	// Experimental.
	P90DurationMetric() interface{}
	// Experimental.
	P99DurationMetric() interface{}
	// Experimental.
	Scope() MonitoringScope
	// Experimental.
	StartedExecutionsMetric() interface{}
	// Experimental.
	StateMachineUrl() *string
	// Experimental.
	SucceededExecutionsMetric() interface{}
	// Experimental.
	TaskHealthAlarmFactory() TaskHealthAlarmFactory
	// Experimental.
	ThrottledExecutionsMetric() interface{}
	// Experimental.
	TimedOutExecutionsMetrics() interface{}
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

// The jsii proxy struct for StepFunctionMonitoring
type jsiiProxy_StepFunctionMonitoring struct {
	jsiiProxy_Monitoring
}

func (j *jsiiProxy_StepFunctionMonitoring) AbortedExecutionsMetric() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"abortedExecutionsMetric",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StepFunctionMonitoring) Alarms() *[]*AlarmWithAnnotation {
	var returns *[]*AlarmWithAnnotation
	_jsii_.Get(
		j,
		"alarms",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StepFunctionMonitoring) DurationAlarmFactory() LatencyAlarmFactory {
	var returns LatencyAlarmFactory
	_jsii_.Get(
		j,
		"durationAlarmFactory",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StepFunctionMonitoring) DurationAnnotations() *[]*awscloudwatch.HorizontalAnnotation {
	var returns *[]*awscloudwatch.HorizontalAnnotation
	_jsii_.Get(
		j,
		"durationAnnotations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StepFunctionMonitoring) ErrorAlarmFactory() ErrorAlarmFactory {
	var returns ErrorAlarmFactory
	_jsii_.Get(
		j,
		"errorAlarmFactory",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StepFunctionMonitoring) ErrorCountAnnotations() *[]*awscloudwatch.HorizontalAnnotation {
	var returns *[]*awscloudwatch.HorizontalAnnotation
	_jsii_.Get(
		j,
		"errorCountAnnotations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StepFunctionMonitoring) ErrorRateAnnotations() *[]*awscloudwatch.HorizontalAnnotation {
	var returns *[]*awscloudwatch.HorizontalAnnotation
	_jsii_.Get(
		j,
		"errorRateAnnotations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StepFunctionMonitoring) FailedExecutionRateMetric() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"failedExecutionRateMetric",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StepFunctionMonitoring) FailedExecutionsMetric() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"failedExecutionsMetric",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StepFunctionMonitoring) LocalAlarmNamePrefixOverride() *string {
	var returns *string
	_jsii_.Get(
		j,
		"localAlarmNamePrefixOverride",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StepFunctionMonitoring) P50DurationMetric() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"p50DurationMetric",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StepFunctionMonitoring) P90DurationMetric() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"p90DurationMetric",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StepFunctionMonitoring) P99DurationMetric() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"p99DurationMetric",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StepFunctionMonitoring) Scope() MonitoringScope {
	var returns MonitoringScope
	_jsii_.Get(
		j,
		"scope",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StepFunctionMonitoring) StartedExecutionsMetric() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"startedExecutionsMetric",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StepFunctionMonitoring) StateMachineUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stateMachineUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StepFunctionMonitoring) SucceededExecutionsMetric() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"succeededExecutionsMetric",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StepFunctionMonitoring) TaskHealthAlarmFactory() TaskHealthAlarmFactory {
	var returns TaskHealthAlarmFactory
	_jsii_.Get(
		j,
		"taskHealthAlarmFactory",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StepFunctionMonitoring) ThrottledExecutionsMetric() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"throttledExecutionsMetric",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StepFunctionMonitoring) TimedOutExecutionsMetrics() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timedOutExecutionsMetrics",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StepFunctionMonitoring) Title() *string {
	var returns *string
	_jsii_.Get(
		j,
		"title",
		&returns,
	)
	return returns
}


// Experimental.
func NewStepFunctionMonitoring(scope MonitoringScope, props *StepFunctionMonitoringProps) StepFunctionMonitoring {
	_init_.Initialize()

	if err := validateNewStepFunctionMonitoringParameters(scope, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_StepFunctionMonitoring{}

	_jsii_.Create(
		"cdk-monitoring-constructs.StepFunctionMonitoring",
		[]interface{}{scope, props},
		&j,
	)

	return &j
}

// Experimental.
func NewStepFunctionMonitoring_Override(s StepFunctionMonitoring, scope MonitoringScope, props *StepFunctionMonitoringProps) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk-monitoring-constructs.StepFunctionMonitoring",
		[]interface{}{scope, props},
		s,
	)
}

func (s *jsiiProxy_StepFunctionMonitoring) AddAlarm(alarm *AlarmWithAnnotation) {
	if err := s.validateAddAlarmParameters(alarm); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"addAlarm",
		[]interface{}{alarm},
	)
}

func (s *jsiiProxy_StepFunctionMonitoring) AlarmWidgets() *[]awscloudwatch.IWidget {
	var returns *[]awscloudwatch.IWidget

	_jsii_.Invoke(
		s,
		"alarmWidgets",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StepFunctionMonitoring) CreateAlarmFactory(alarmNamePrefix *string) AlarmFactory {
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

func (s *jsiiProxy_StepFunctionMonitoring) CreatedAlarms() *[]*AlarmWithAnnotation {
	var returns *[]*AlarmWithAnnotation

	_jsii_.Invoke(
		s,
		"createdAlarms",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StepFunctionMonitoring) CreateMetricFactory() MetricFactory {
	var returns MetricFactory

	_jsii_.Invoke(
		s,
		"createMetricFactory",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StepFunctionMonitoring) CreateWidgetFactory() IWidgetFactory {
	var returns IWidgetFactory

	_jsii_.Invoke(
		s,
		"createWidgetFactory",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StepFunctionMonitoring) SummaryWidgets() *[]awscloudwatch.IWidget {
	var returns *[]awscloudwatch.IWidget

	_jsii_.Invoke(
		s,
		"summaryWidgets",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StepFunctionMonitoring) Widgets() *[]awscloudwatch.IWidget {
	var returns *[]awscloudwatch.IWidget

	_jsii_.Invoke(
		s,
		"widgets",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StepFunctionMonitoring) WidgetsForDashboard(name *string) *[]awscloudwatch.IWidget {
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


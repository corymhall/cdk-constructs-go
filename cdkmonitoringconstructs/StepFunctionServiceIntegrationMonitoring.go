package cdkmonitoringconstructs

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/corymhall/cdk-constructs-go/cdkmonitoringconstructs/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudwatch"
)

// Experimental.
type StepFunctionServiceIntegrationMonitoring interface {
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
	FailedServiceIntegrationRateMetric() interface{}
	// Experimental.
	FailedServiceIntegrationsMetric() interface{}
	// Experimental.
	LocalAlarmNamePrefixOverride() *string
	// Experimental.
	P50DurationMetric() interface{}
	// Experimental.
	P90DurationMetric() interface{}
	// Experimental.
	P99DurationMetric() interface{}
	// Experimental.
	ScheduledServiceIntegrationsMetric() interface{}
	// Experimental.
	Scope() MonitoringScope
	// Experimental.
	StartedServiceIntegrationsMetric() interface{}
	// Experimental.
	SucceededServiceIntegrationsMetric() interface{}
	// Experimental.
	TimedOutServiceIntegrationsMetrics() interface{}
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

// The jsii proxy struct for StepFunctionServiceIntegrationMonitoring
type jsiiProxy_StepFunctionServiceIntegrationMonitoring struct {
	jsiiProxy_Monitoring
}

func (j *jsiiProxy_StepFunctionServiceIntegrationMonitoring) Alarms() *[]*AlarmWithAnnotation {
	var returns *[]*AlarmWithAnnotation
	_jsii_.Get(
		j,
		"alarms",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StepFunctionServiceIntegrationMonitoring) DurationAlarmFactory() LatencyAlarmFactory {
	var returns LatencyAlarmFactory
	_jsii_.Get(
		j,
		"durationAlarmFactory",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StepFunctionServiceIntegrationMonitoring) DurationAnnotations() *[]*awscloudwatch.HorizontalAnnotation {
	var returns *[]*awscloudwatch.HorizontalAnnotation
	_jsii_.Get(
		j,
		"durationAnnotations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StepFunctionServiceIntegrationMonitoring) ErrorAlarmFactory() ErrorAlarmFactory {
	var returns ErrorAlarmFactory
	_jsii_.Get(
		j,
		"errorAlarmFactory",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StepFunctionServiceIntegrationMonitoring) ErrorCountAnnotations() *[]*awscloudwatch.HorizontalAnnotation {
	var returns *[]*awscloudwatch.HorizontalAnnotation
	_jsii_.Get(
		j,
		"errorCountAnnotations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StepFunctionServiceIntegrationMonitoring) ErrorRateAnnotations() *[]*awscloudwatch.HorizontalAnnotation {
	var returns *[]*awscloudwatch.HorizontalAnnotation
	_jsii_.Get(
		j,
		"errorRateAnnotations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StepFunctionServiceIntegrationMonitoring) FailedServiceIntegrationRateMetric() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"failedServiceIntegrationRateMetric",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StepFunctionServiceIntegrationMonitoring) FailedServiceIntegrationsMetric() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"failedServiceIntegrationsMetric",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StepFunctionServiceIntegrationMonitoring) LocalAlarmNamePrefixOverride() *string {
	var returns *string
	_jsii_.Get(
		j,
		"localAlarmNamePrefixOverride",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StepFunctionServiceIntegrationMonitoring) P50DurationMetric() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"p50DurationMetric",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StepFunctionServiceIntegrationMonitoring) P90DurationMetric() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"p90DurationMetric",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StepFunctionServiceIntegrationMonitoring) P99DurationMetric() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"p99DurationMetric",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StepFunctionServiceIntegrationMonitoring) ScheduledServiceIntegrationsMetric() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"scheduledServiceIntegrationsMetric",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StepFunctionServiceIntegrationMonitoring) Scope() MonitoringScope {
	var returns MonitoringScope
	_jsii_.Get(
		j,
		"scope",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StepFunctionServiceIntegrationMonitoring) StartedServiceIntegrationsMetric() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"startedServiceIntegrationsMetric",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StepFunctionServiceIntegrationMonitoring) SucceededServiceIntegrationsMetric() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"succeededServiceIntegrationsMetric",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StepFunctionServiceIntegrationMonitoring) TimedOutServiceIntegrationsMetrics() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timedOutServiceIntegrationsMetrics",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StepFunctionServiceIntegrationMonitoring) Title() *string {
	var returns *string
	_jsii_.Get(
		j,
		"title",
		&returns,
	)
	return returns
}


// Experimental.
func NewStepFunctionServiceIntegrationMonitoring(scope MonitoringScope, props *StepFunctionServiceIntegrationMonitoringProps) StepFunctionServiceIntegrationMonitoring {
	_init_.Initialize()

	if err := validateNewStepFunctionServiceIntegrationMonitoringParameters(scope, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_StepFunctionServiceIntegrationMonitoring{}

	_jsii_.Create(
		"cdk-monitoring-constructs.StepFunctionServiceIntegrationMonitoring",
		[]interface{}{scope, props},
		&j,
	)

	return &j
}

// Experimental.
func NewStepFunctionServiceIntegrationMonitoring_Override(s StepFunctionServiceIntegrationMonitoring, scope MonitoringScope, props *StepFunctionServiceIntegrationMonitoringProps) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk-monitoring-constructs.StepFunctionServiceIntegrationMonitoring",
		[]interface{}{scope, props},
		s,
	)
}

func (s *jsiiProxy_StepFunctionServiceIntegrationMonitoring) AddAlarm(alarm *AlarmWithAnnotation) {
	if err := s.validateAddAlarmParameters(alarm); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"addAlarm",
		[]interface{}{alarm},
	)
}

func (s *jsiiProxy_StepFunctionServiceIntegrationMonitoring) AlarmWidgets() *[]awscloudwatch.IWidget {
	var returns *[]awscloudwatch.IWidget

	_jsii_.Invoke(
		s,
		"alarmWidgets",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StepFunctionServiceIntegrationMonitoring) CreateAlarmFactory(alarmNamePrefix *string) AlarmFactory {
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

func (s *jsiiProxy_StepFunctionServiceIntegrationMonitoring) CreatedAlarms() *[]*AlarmWithAnnotation {
	var returns *[]*AlarmWithAnnotation

	_jsii_.Invoke(
		s,
		"createdAlarms",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StepFunctionServiceIntegrationMonitoring) CreateMetricFactory() MetricFactory {
	var returns MetricFactory

	_jsii_.Invoke(
		s,
		"createMetricFactory",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StepFunctionServiceIntegrationMonitoring) CreateWidgetFactory() IWidgetFactory {
	var returns IWidgetFactory

	_jsii_.Invoke(
		s,
		"createWidgetFactory",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StepFunctionServiceIntegrationMonitoring) SummaryWidgets() *[]awscloudwatch.IWidget {
	var returns *[]awscloudwatch.IWidget

	_jsii_.Invoke(
		s,
		"summaryWidgets",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StepFunctionServiceIntegrationMonitoring) Widgets() *[]awscloudwatch.IWidget {
	var returns *[]awscloudwatch.IWidget

	_jsii_.Invoke(
		s,
		"widgets",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StepFunctionServiceIntegrationMonitoring) WidgetsForDashboard(name *string) *[]awscloudwatch.IWidget {
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


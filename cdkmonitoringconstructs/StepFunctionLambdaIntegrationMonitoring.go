package cdkmonitoringconstructs

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/corymhall/cdk-constructs-go/cdkmonitoringconstructs/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudwatch"
)

// Experimental.
type StepFunctionLambdaIntegrationMonitoring interface {
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
	FailedFunctionRateMetric() interface{}
	// Experimental.
	FailedFunctionsMetric() interface{}
	// Experimental.
	FunctionUrl() *string
	// Experimental.
	LocalAlarmNamePrefixOverride() *string
	// Experimental.
	P50DurationMetric() interface{}
	// Experimental.
	P90DurationMetric() interface{}
	// Experimental.
	P99DurationMetric() interface{}
	// Experimental.
	ScheduledFunctionsMetric() interface{}
	// Experimental.
	Scope() MonitoringScope
	// Experimental.
	StartedFunctionsMetric() interface{}
	// Experimental.
	SucceededFunctionsMetric() interface{}
	// Experimental.
	TimedOutFunctionsMetrics() interface{}
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

// The jsii proxy struct for StepFunctionLambdaIntegrationMonitoring
type jsiiProxy_StepFunctionLambdaIntegrationMonitoring struct {
	jsiiProxy_Monitoring
}

func (j *jsiiProxy_StepFunctionLambdaIntegrationMonitoring) Alarms() *[]*AlarmWithAnnotation {
	var returns *[]*AlarmWithAnnotation
	_jsii_.Get(
		j,
		"alarms",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StepFunctionLambdaIntegrationMonitoring) DurationAlarmFactory() LatencyAlarmFactory {
	var returns LatencyAlarmFactory
	_jsii_.Get(
		j,
		"durationAlarmFactory",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StepFunctionLambdaIntegrationMonitoring) DurationAnnotations() *[]*awscloudwatch.HorizontalAnnotation {
	var returns *[]*awscloudwatch.HorizontalAnnotation
	_jsii_.Get(
		j,
		"durationAnnotations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StepFunctionLambdaIntegrationMonitoring) ErrorAlarmFactory() ErrorAlarmFactory {
	var returns ErrorAlarmFactory
	_jsii_.Get(
		j,
		"errorAlarmFactory",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StepFunctionLambdaIntegrationMonitoring) ErrorCountAnnotations() *[]*awscloudwatch.HorizontalAnnotation {
	var returns *[]*awscloudwatch.HorizontalAnnotation
	_jsii_.Get(
		j,
		"errorCountAnnotations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StepFunctionLambdaIntegrationMonitoring) ErrorRateAnnotations() *[]*awscloudwatch.HorizontalAnnotation {
	var returns *[]*awscloudwatch.HorizontalAnnotation
	_jsii_.Get(
		j,
		"errorRateAnnotations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StepFunctionLambdaIntegrationMonitoring) FailedFunctionRateMetric() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"failedFunctionRateMetric",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StepFunctionLambdaIntegrationMonitoring) FailedFunctionsMetric() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"failedFunctionsMetric",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StepFunctionLambdaIntegrationMonitoring) FunctionUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"functionUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StepFunctionLambdaIntegrationMonitoring) LocalAlarmNamePrefixOverride() *string {
	var returns *string
	_jsii_.Get(
		j,
		"localAlarmNamePrefixOverride",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StepFunctionLambdaIntegrationMonitoring) P50DurationMetric() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"p50DurationMetric",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StepFunctionLambdaIntegrationMonitoring) P90DurationMetric() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"p90DurationMetric",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StepFunctionLambdaIntegrationMonitoring) P99DurationMetric() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"p99DurationMetric",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StepFunctionLambdaIntegrationMonitoring) ScheduledFunctionsMetric() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"scheduledFunctionsMetric",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StepFunctionLambdaIntegrationMonitoring) Scope() MonitoringScope {
	var returns MonitoringScope
	_jsii_.Get(
		j,
		"scope",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StepFunctionLambdaIntegrationMonitoring) StartedFunctionsMetric() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"startedFunctionsMetric",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StepFunctionLambdaIntegrationMonitoring) SucceededFunctionsMetric() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"succeededFunctionsMetric",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StepFunctionLambdaIntegrationMonitoring) TimedOutFunctionsMetrics() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timedOutFunctionsMetrics",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StepFunctionLambdaIntegrationMonitoring) Title() *string {
	var returns *string
	_jsii_.Get(
		j,
		"title",
		&returns,
	)
	return returns
}


// Experimental.
func NewStepFunctionLambdaIntegrationMonitoring(scope MonitoringScope, props *StepFunctionLambdaIntegrationMonitoringProps) StepFunctionLambdaIntegrationMonitoring {
	_init_.Initialize()

	if err := validateNewStepFunctionLambdaIntegrationMonitoringParameters(scope, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_StepFunctionLambdaIntegrationMonitoring{}

	_jsii_.Create(
		"cdk-monitoring-constructs.StepFunctionLambdaIntegrationMonitoring",
		[]interface{}{scope, props},
		&j,
	)

	return &j
}

// Experimental.
func NewStepFunctionLambdaIntegrationMonitoring_Override(s StepFunctionLambdaIntegrationMonitoring, scope MonitoringScope, props *StepFunctionLambdaIntegrationMonitoringProps) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk-monitoring-constructs.StepFunctionLambdaIntegrationMonitoring",
		[]interface{}{scope, props},
		s,
	)
}

func (s *jsiiProxy_StepFunctionLambdaIntegrationMonitoring) AddAlarm(alarm *AlarmWithAnnotation) {
	if err := s.validateAddAlarmParameters(alarm); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"addAlarm",
		[]interface{}{alarm},
	)
}

func (s *jsiiProxy_StepFunctionLambdaIntegrationMonitoring) AlarmWidgets() *[]awscloudwatch.IWidget {
	var returns *[]awscloudwatch.IWidget

	_jsii_.Invoke(
		s,
		"alarmWidgets",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StepFunctionLambdaIntegrationMonitoring) CreateAlarmFactory(alarmNamePrefix *string) AlarmFactory {
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

func (s *jsiiProxy_StepFunctionLambdaIntegrationMonitoring) CreatedAlarms() *[]*AlarmWithAnnotation {
	var returns *[]*AlarmWithAnnotation

	_jsii_.Invoke(
		s,
		"createdAlarms",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StepFunctionLambdaIntegrationMonitoring) CreateMetricFactory() MetricFactory {
	var returns MetricFactory

	_jsii_.Invoke(
		s,
		"createMetricFactory",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StepFunctionLambdaIntegrationMonitoring) CreateWidgetFactory() IWidgetFactory {
	var returns IWidgetFactory

	_jsii_.Invoke(
		s,
		"createWidgetFactory",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StepFunctionLambdaIntegrationMonitoring) SummaryWidgets() *[]awscloudwatch.IWidget {
	var returns *[]awscloudwatch.IWidget

	_jsii_.Invoke(
		s,
		"summaryWidgets",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StepFunctionLambdaIntegrationMonitoring) Widgets() *[]awscloudwatch.IWidget {
	var returns *[]awscloudwatch.IWidget

	_jsii_.Invoke(
		s,
		"widgets",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StepFunctionLambdaIntegrationMonitoring) WidgetsForDashboard(name *string) *[]awscloudwatch.IWidget {
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


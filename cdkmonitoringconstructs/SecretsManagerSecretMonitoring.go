package cdkmonitoringconstructs

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/corymhall/cdk-constructs-go/cdkmonitoringconstructs/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudwatch"
)

// Experimental.
type SecretsManagerSecretMonitoring interface {
	Monitoring
	// Experimental.
	AlarmFactory() AlarmFactory
	// Experimental.
	Alarms() *[]*AlarmWithAnnotation
	// Experimental.
	DaysSinceLastChangeAnnotations() *[]*awscloudwatch.HorizontalAnnotation
	// Experimental.
	DaysSinceLastChangeMetric() interface{}
	// Experimental.
	DaysSinceLastRotationAnnotations() *[]*awscloudwatch.HorizontalAnnotation
	// Experimental.
	DaysSinceLastRotationMetric() interface{}
	// Experimental.
	LocalAlarmNamePrefixOverride() *string
	// Experimental.
	Scope() MonitoringScope
	// Experimental.
	ShowLastRotationWidget() *bool
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
	CreateDaysSinceLastChangeWidget() awscloudwatch.GraphWidget
	// Experimental.
	CreateDaysSinceLastRotationWidget() awscloudwatch.GraphWidget
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

// The jsii proxy struct for SecretsManagerSecretMonitoring
type jsiiProxy_SecretsManagerSecretMonitoring struct {
	jsiiProxy_Monitoring
}

func (j *jsiiProxy_SecretsManagerSecretMonitoring) AlarmFactory() AlarmFactory {
	var returns AlarmFactory
	_jsii_.Get(
		j,
		"alarmFactory",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecretsManagerSecretMonitoring) Alarms() *[]*AlarmWithAnnotation {
	var returns *[]*AlarmWithAnnotation
	_jsii_.Get(
		j,
		"alarms",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecretsManagerSecretMonitoring) DaysSinceLastChangeAnnotations() *[]*awscloudwatch.HorizontalAnnotation {
	var returns *[]*awscloudwatch.HorizontalAnnotation
	_jsii_.Get(
		j,
		"daysSinceLastChangeAnnotations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecretsManagerSecretMonitoring) DaysSinceLastChangeMetric() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"daysSinceLastChangeMetric",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecretsManagerSecretMonitoring) DaysSinceLastRotationAnnotations() *[]*awscloudwatch.HorizontalAnnotation {
	var returns *[]*awscloudwatch.HorizontalAnnotation
	_jsii_.Get(
		j,
		"daysSinceLastRotationAnnotations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecretsManagerSecretMonitoring) DaysSinceLastRotationMetric() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"daysSinceLastRotationMetric",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecretsManagerSecretMonitoring) LocalAlarmNamePrefixOverride() *string {
	var returns *string
	_jsii_.Get(
		j,
		"localAlarmNamePrefixOverride",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecretsManagerSecretMonitoring) Scope() MonitoringScope {
	var returns MonitoringScope
	_jsii_.Get(
		j,
		"scope",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecretsManagerSecretMonitoring) ShowLastRotationWidget() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"showLastRotationWidget",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecretsManagerSecretMonitoring) Title() *string {
	var returns *string
	_jsii_.Get(
		j,
		"title",
		&returns,
	)
	return returns
}


// Experimental.
func NewSecretsManagerSecretMonitoring(scope MonitoringScope, props *SecretsManagerSecretMonitoringProps) SecretsManagerSecretMonitoring {
	_init_.Initialize()

	if err := validateNewSecretsManagerSecretMonitoringParameters(scope, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_SecretsManagerSecretMonitoring{}

	_jsii_.Create(
		"cdk-monitoring-constructs.SecretsManagerSecretMonitoring",
		[]interface{}{scope, props},
		&j,
	)

	return &j
}

// Experimental.
func NewSecretsManagerSecretMonitoring_Override(s SecretsManagerSecretMonitoring, scope MonitoringScope, props *SecretsManagerSecretMonitoringProps) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk-monitoring-constructs.SecretsManagerSecretMonitoring",
		[]interface{}{scope, props},
		s,
	)
}

func (s *jsiiProxy_SecretsManagerSecretMonitoring) AddAlarm(alarm *AlarmWithAnnotation) {
	if err := s.validateAddAlarmParameters(alarm); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"addAlarm",
		[]interface{}{alarm},
	)
}

func (s *jsiiProxy_SecretsManagerSecretMonitoring) AlarmWidgets() *[]awscloudwatch.IWidget {
	var returns *[]awscloudwatch.IWidget

	_jsii_.Invoke(
		s,
		"alarmWidgets",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SecretsManagerSecretMonitoring) CreateAlarmFactory(alarmNamePrefix *string) AlarmFactory {
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

func (s *jsiiProxy_SecretsManagerSecretMonitoring) CreatedAlarms() *[]*AlarmWithAnnotation {
	var returns *[]*AlarmWithAnnotation

	_jsii_.Invoke(
		s,
		"createdAlarms",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SecretsManagerSecretMonitoring) CreateDaysSinceLastChangeWidget() awscloudwatch.GraphWidget {
	var returns awscloudwatch.GraphWidget

	_jsii_.Invoke(
		s,
		"createDaysSinceLastChangeWidget",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SecretsManagerSecretMonitoring) CreateDaysSinceLastRotationWidget() awscloudwatch.GraphWidget {
	var returns awscloudwatch.GraphWidget

	_jsii_.Invoke(
		s,
		"createDaysSinceLastRotationWidget",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SecretsManagerSecretMonitoring) CreateMetricFactory() MetricFactory {
	var returns MetricFactory

	_jsii_.Invoke(
		s,
		"createMetricFactory",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SecretsManagerSecretMonitoring) CreateTitleWidget() MonitoringHeaderWidget {
	var returns MonitoringHeaderWidget

	_jsii_.Invoke(
		s,
		"createTitleWidget",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SecretsManagerSecretMonitoring) CreateWidgetFactory() IWidgetFactory {
	var returns IWidgetFactory

	_jsii_.Invoke(
		s,
		"createWidgetFactory",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SecretsManagerSecretMonitoring) SummaryWidgets() *[]awscloudwatch.IWidget {
	var returns *[]awscloudwatch.IWidget

	_jsii_.Invoke(
		s,
		"summaryWidgets",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SecretsManagerSecretMonitoring) Widgets() *[]awscloudwatch.IWidget {
	var returns *[]awscloudwatch.IWidget

	_jsii_.Invoke(
		s,
		"widgets",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SecretsManagerSecretMonitoring) WidgetsForDashboard(name *string) *[]awscloudwatch.IWidget {
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


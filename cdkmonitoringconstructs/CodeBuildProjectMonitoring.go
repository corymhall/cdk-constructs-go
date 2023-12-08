package cdkmonitoringconstructs

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/corymhall/cdk-constructs-go/cdkmonitoringconstructs/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudwatch"
	"github.com/aws/aws-cdk-go/awscdk/v2/awscodebuild"
)

// Experimental.
type CodeBuildProjectMonitoring interface {
	Monitoring
	// Experimental.
	Alarms() *[]*AlarmWithAnnotation
	// Experimental.
	BuildCountMetric() interface{}
	// Experimental.
	DurationAlarmFactory() LatencyAlarmFactory
	// Experimental.
	DurationAnnotations() *[]*awscloudwatch.HorizontalAnnotation
	// Experimental.
	DurationP50InSecondsMetric() interface{}
	// Experimental.
	DurationP90InSecondsMetric() interface{}
	// Experimental.
	DurationP99InSecondsMetric() interface{}
	// Experimental.
	ErrorAlarmFactory() ErrorAlarmFactory
	// Experimental.
	ErrorCountAnnotations() *[]*awscloudwatch.HorizontalAnnotation
	// Experimental.
	ErrorRateAnnotations() *[]*awscloudwatch.HorizontalAnnotation
	// Experimental.
	FailedBuildCountMetric() interface{}
	// Experimental.
	FailedBuildRateMetric() interface{}
	// Experimental.
	LocalAlarmNamePrefixOverride() *string
	// Experimental.
	ProjectUrl() *string
	// Experimental.
	Scope() MonitoringScope
	// Experimental.
	SucceededBuildCountMetric() interface{}
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
	// Experimental.
	CreateBuildCountsWidget() awscloudwatch.GraphWidget
	// Returns all the alarms created.
	// Experimental.
	CreatedAlarms() *[]*AlarmWithAnnotation
	// Experimental.
	CreateDurationWidget() awscloudwatch.GraphWidget
	// Experimental.
	CreateFailedBuildRateWidget() awscloudwatch.GraphWidget
	// Creates a new metric factory.
	// Experimental.
	CreateMetricFactory() MetricFactory
	// Experimental.
	CreateTitleWidget() MonitoringHeaderWidget
	// Creates a new widget factory.
	// Experimental.
	CreateWidgetFactory() IWidgetFactory
	// Experimental.
	ResolveProjectName(project awscodebuild.IProject) *string
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

// The jsii proxy struct for CodeBuildProjectMonitoring
type jsiiProxy_CodeBuildProjectMonitoring struct {
	jsiiProxy_Monitoring
}

func (j *jsiiProxy_CodeBuildProjectMonitoring) Alarms() *[]*AlarmWithAnnotation {
	var returns *[]*AlarmWithAnnotation
	_jsii_.Get(
		j,
		"alarms",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodeBuildProjectMonitoring) BuildCountMetric() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"buildCountMetric",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodeBuildProjectMonitoring) DurationAlarmFactory() LatencyAlarmFactory {
	var returns LatencyAlarmFactory
	_jsii_.Get(
		j,
		"durationAlarmFactory",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodeBuildProjectMonitoring) DurationAnnotations() *[]*awscloudwatch.HorizontalAnnotation {
	var returns *[]*awscloudwatch.HorizontalAnnotation
	_jsii_.Get(
		j,
		"durationAnnotations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodeBuildProjectMonitoring) DurationP50InSecondsMetric() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"durationP50InSecondsMetric",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodeBuildProjectMonitoring) DurationP90InSecondsMetric() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"durationP90InSecondsMetric",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodeBuildProjectMonitoring) DurationP99InSecondsMetric() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"durationP99InSecondsMetric",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodeBuildProjectMonitoring) ErrorAlarmFactory() ErrorAlarmFactory {
	var returns ErrorAlarmFactory
	_jsii_.Get(
		j,
		"errorAlarmFactory",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodeBuildProjectMonitoring) ErrorCountAnnotations() *[]*awscloudwatch.HorizontalAnnotation {
	var returns *[]*awscloudwatch.HorizontalAnnotation
	_jsii_.Get(
		j,
		"errorCountAnnotations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodeBuildProjectMonitoring) ErrorRateAnnotations() *[]*awscloudwatch.HorizontalAnnotation {
	var returns *[]*awscloudwatch.HorizontalAnnotation
	_jsii_.Get(
		j,
		"errorRateAnnotations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodeBuildProjectMonitoring) FailedBuildCountMetric() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"failedBuildCountMetric",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodeBuildProjectMonitoring) FailedBuildRateMetric() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"failedBuildRateMetric",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodeBuildProjectMonitoring) LocalAlarmNamePrefixOverride() *string {
	var returns *string
	_jsii_.Get(
		j,
		"localAlarmNamePrefixOverride",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodeBuildProjectMonitoring) ProjectUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"projectUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodeBuildProjectMonitoring) Scope() MonitoringScope {
	var returns MonitoringScope
	_jsii_.Get(
		j,
		"scope",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodeBuildProjectMonitoring) SucceededBuildCountMetric() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"succeededBuildCountMetric",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodeBuildProjectMonitoring) Title() *string {
	var returns *string
	_jsii_.Get(
		j,
		"title",
		&returns,
	)
	return returns
}


// Experimental.
func NewCodeBuildProjectMonitoring(scope MonitoringScope, props *CodeBuildProjectMonitoringProps) CodeBuildProjectMonitoring {
	_init_.Initialize()

	if err := validateNewCodeBuildProjectMonitoringParameters(scope, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_CodeBuildProjectMonitoring{}

	_jsii_.Create(
		"cdk-monitoring-constructs.CodeBuildProjectMonitoring",
		[]interface{}{scope, props},
		&j,
	)

	return &j
}

// Experimental.
func NewCodeBuildProjectMonitoring_Override(c CodeBuildProjectMonitoring, scope MonitoringScope, props *CodeBuildProjectMonitoringProps) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk-monitoring-constructs.CodeBuildProjectMonitoring",
		[]interface{}{scope, props},
		c,
	)
}

func (c *jsiiProxy_CodeBuildProjectMonitoring) AddAlarm(alarm *AlarmWithAnnotation) {
	if err := c.validateAddAlarmParameters(alarm); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addAlarm",
		[]interface{}{alarm},
	)
}

func (c *jsiiProxy_CodeBuildProjectMonitoring) AlarmWidgets() *[]awscloudwatch.IWidget {
	var returns *[]awscloudwatch.IWidget

	_jsii_.Invoke(
		c,
		"alarmWidgets",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CodeBuildProjectMonitoring) CreateAlarmFactory(alarmNamePrefix *string) AlarmFactory {
	if err := c.validateCreateAlarmFactoryParameters(alarmNamePrefix); err != nil {
		panic(err)
	}
	var returns AlarmFactory

	_jsii_.Invoke(
		c,
		"createAlarmFactory",
		[]interface{}{alarmNamePrefix},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CodeBuildProjectMonitoring) CreateBuildCountsWidget() awscloudwatch.GraphWidget {
	var returns awscloudwatch.GraphWidget

	_jsii_.Invoke(
		c,
		"createBuildCountsWidget",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CodeBuildProjectMonitoring) CreatedAlarms() *[]*AlarmWithAnnotation {
	var returns *[]*AlarmWithAnnotation

	_jsii_.Invoke(
		c,
		"createdAlarms",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CodeBuildProjectMonitoring) CreateDurationWidget() awscloudwatch.GraphWidget {
	var returns awscloudwatch.GraphWidget

	_jsii_.Invoke(
		c,
		"createDurationWidget",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CodeBuildProjectMonitoring) CreateFailedBuildRateWidget() awscloudwatch.GraphWidget {
	var returns awscloudwatch.GraphWidget

	_jsii_.Invoke(
		c,
		"createFailedBuildRateWidget",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CodeBuildProjectMonitoring) CreateMetricFactory() MetricFactory {
	var returns MetricFactory

	_jsii_.Invoke(
		c,
		"createMetricFactory",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CodeBuildProjectMonitoring) CreateTitleWidget() MonitoringHeaderWidget {
	var returns MonitoringHeaderWidget

	_jsii_.Invoke(
		c,
		"createTitleWidget",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CodeBuildProjectMonitoring) CreateWidgetFactory() IWidgetFactory {
	var returns IWidgetFactory

	_jsii_.Invoke(
		c,
		"createWidgetFactory",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CodeBuildProjectMonitoring) ResolveProjectName(project awscodebuild.IProject) *string {
	if err := c.validateResolveProjectNameParameters(project); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		c,
		"resolveProjectName",
		[]interface{}{project},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CodeBuildProjectMonitoring) SummaryWidgets() *[]awscloudwatch.IWidget {
	var returns *[]awscloudwatch.IWidget

	_jsii_.Invoke(
		c,
		"summaryWidgets",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CodeBuildProjectMonitoring) Widgets() *[]awscloudwatch.IWidget {
	var returns *[]awscloudwatch.IWidget

	_jsii_.Invoke(
		c,
		"widgets",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CodeBuildProjectMonitoring) WidgetsForDashboard(name *string) *[]awscloudwatch.IWidget {
	if err := c.validateWidgetsForDashboardParameters(name); err != nil {
		panic(err)
	}
	var returns *[]awscloudwatch.IWidget

	_jsii_.Invoke(
		c,
		"widgetsForDashboard",
		[]interface{}{name},
		&returns,
	)

	return returns
}


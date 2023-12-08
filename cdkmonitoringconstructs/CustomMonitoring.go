package cdkmonitoringconstructs

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/corymhall/cdk-constructs-go/cdkmonitoringconstructs/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudwatch"
)

// Custom monitoring is a construct allowing you to monitor your own custom metrics.
//
// The entire construct consists of metric groups.
// Each metric group represents a single graph widget with multiple metrics.
// Each metric inside the metric group represents a single metric inside a graph.
// The widgets will be sized automatically to waste as little space as possible.
// Experimental.
type CustomMonitoring interface {
	Monitoring
	// Experimental.
	AddToSummaryDashboard() *bool
	// Experimental.
	Alarms() *[]*AlarmWithAnnotation
	// Experimental.
	AnomalyDetectingAlarmFactory() AnomalyDetectingAlarmFactory
	// Experimental.
	CustomAlarmFactory() CustomAlarmFactory
	// Experimental.
	Description() *string
	// Experimental.
	DescriptionWidgetHeight() *float64
	// Experimental.
	Height() *float64
	// Experimental.
	LocalAlarmNamePrefixOverride() *string
	// Experimental.
	MetricGroups() *[]*CustomMetricGroupWithAnnotations
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

// The jsii proxy struct for CustomMonitoring
type jsiiProxy_CustomMonitoring struct {
	jsiiProxy_Monitoring
}

func (j *jsiiProxy_CustomMonitoring) AddToSummaryDashboard() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"addToSummaryDashboard",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomMonitoring) Alarms() *[]*AlarmWithAnnotation {
	var returns *[]*AlarmWithAnnotation
	_jsii_.Get(
		j,
		"alarms",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomMonitoring) AnomalyDetectingAlarmFactory() AnomalyDetectingAlarmFactory {
	var returns AnomalyDetectingAlarmFactory
	_jsii_.Get(
		j,
		"anomalyDetectingAlarmFactory",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomMonitoring) CustomAlarmFactory() CustomAlarmFactory {
	var returns CustomAlarmFactory
	_jsii_.Get(
		j,
		"customAlarmFactory",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomMonitoring) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomMonitoring) DescriptionWidgetHeight() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"descriptionWidgetHeight",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomMonitoring) Height() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"height",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomMonitoring) LocalAlarmNamePrefixOverride() *string {
	var returns *string
	_jsii_.Get(
		j,
		"localAlarmNamePrefixOverride",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomMonitoring) MetricGroups() *[]*CustomMetricGroupWithAnnotations {
	var returns *[]*CustomMetricGroupWithAnnotations
	_jsii_.Get(
		j,
		"metricGroups",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomMonitoring) Scope() MonitoringScope {
	var returns MonitoringScope
	_jsii_.Get(
		j,
		"scope",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomMonitoring) Title() *string {
	var returns *string
	_jsii_.Get(
		j,
		"title",
		&returns,
	)
	return returns
}


// Experimental.
func NewCustomMonitoring(scope MonitoringScope, props *CustomMonitoringProps) CustomMonitoring {
	_init_.Initialize()

	if err := validateNewCustomMonitoringParameters(scope, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_CustomMonitoring{}

	_jsii_.Create(
		"cdk-monitoring-constructs.CustomMonitoring",
		[]interface{}{scope, props},
		&j,
	)

	return &j
}

// Experimental.
func NewCustomMonitoring_Override(c CustomMonitoring, scope MonitoringScope, props *CustomMonitoringProps) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk-monitoring-constructs.CustomMonitoring",
		[]interface{}{scope, props},
		c,
	)
}

func (c *jsiiProxy_CustomMonitoring) AddAlarm(alarm *AlarmWithAnnotation) {
	if err := c.validateAddAlarmParameters(alarm); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addAlarm",
		[]interface{}{alarm},
	)
}

func (c *jsiiProxy_CustomMonitoring) AlarmWidgets() *[]awscloudwatch.IWidget {
	var returns *[]awscloudwatch.IWidget

	_jsii_.Invoke(
		c,
		"alarmWidgets",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CustomMonitoring) CreateAlarmFactory(alarmNamePrefix *string) AlarmFactory {
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

func (c *jsiiProxy_CustomMonitoring) CreatedAlarms() *[]*AlarmWithAnnotation {
	var returns *[]*AlarmWithAnnotation

	_jsii_.Invoke(
		c,
		"createdAlarms",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CustomMonitoring) CreateMetricFactory() MetricFactory {
	var returns MetricFactory

	_jsii_.Invoke(
		c,
		"createMetricFactory",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CustomMonitoring) CreateWidgetFactory() IWidgetFactory {
	var returns IWidgetFactory

	_jsii_.Invoke(
		c,
		"createWidgetFactory",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CustomMonitoring) SummaryWidgets() *[]awscloudwatch.IWidget {
	var returns *[]awscloudwatch.IWidget

	_jsii_.Invoke(
		c,
		"summaryWidgets",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CustomMonitoring) Widgets() *[]awscloudwatch.IWidget {
	var returns *[]awscloudwatch.IWidget

	_jsii_.Invoke(
		c,
		"widgets",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CustomMonitoring) WidgetsForDashboard(name *string) *[]awscloudwatch.IWidget {
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


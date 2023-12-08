package cdkmonitoringconstructs

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/corymhall/cdk-constructs-go/cdkmonitoringconstructs/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudwatch"
)

// Experimental.
type CertificateManagerMonitoring interface {
	Monitoring
	// Experimental.
	Alarms() *[]*AlarmWithAnnotation
	// Experimental.
	DaysToExpiryAnnotations() *[]*awscloudwatch.HorizontalAnnotation
	// Experimental.
	DaysToExpiryMetric() interface{}
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
	CreateDaysToExpiryWidget(width *float64, height *float64) awscloudwatch.GraphWidget
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

// The jsii proxy struct for CertificateManagerMonitoring
type jsiiProxy_CertificateManagerMonitoring struct {
	jsiiProxy_Monitoring
}

func (j *jsiiProxy_CertificateManagerMonitoring) Alarms() *[]*AlarmWithAnnotation {
	var returns *[]*AlarmWithAnnotation
	_jsii_.Get(
		j,
		"alarms",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CertificateManagerMonitoring) DaysToExpiryAnnotations() *[]*awscloudwatch.HorizontalAnnotation {
	var returns *[]*awscloudwatch.HorizontalAnnotation
	_jsii_.Get(
		j,
		"daysToExpiryAnnotations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CertificateManagerMonitoring) DaysToExpiryMetric() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"daysToExpiryMetric",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CertificateManagerMonitoring) LocalAlarmNamePrefixOverride() *string {
	var returns *string
	_jsii_.Get(
		j,
		"localAlarmNamePrefixOverride",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CertificateManagerMonitoring) Scope() MonitoringScope {
	var returns MonitoringScope
	_jsii_.Get(
		j,
		"scope",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CertificateManagerMonitoring) Title() *string {
	var returns *string
	_jsii_.Get(
		j,
		"title",
		&returns,
	)
	return returns
}


// Experimental.
func NewCertificateManagerMonitoring(scope MonitoringScope, props *CertificateManagerMonitoringProps) CertificateManagerMonitoring {
	_init_.Initialize()

	if err := validateNewCertificateManagerMonitoringParameters(scope, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_CertificateManagerMonitoring{}

	_jsii_.Create(
		"cdk-monitoring-constructs.CertificateManagerMonitoring",
		[]interface{}{scope, props},
		&j,
	)

	return &j
}

// Experimental.
func NewCertificateManagerMonitoring_Override(c CertificateManagerMonitoring, scope MonitoringScope, props *CertificateManagerMonitoringProps) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk-monitoring-constructs.CertificateManagerMonitoring",
		[]interface{}{scope, props},
		c,
	)
}

func (c *jsiiProxy_CertificateManagerMonitoring) AddAlarm(alarm *AlarmWithAnnotation) {
	if err := c.validateAddAlarmParameters(alarm); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addAlarm",
		[]interface{}{alarm},
	)
}

func (c *jsiiProxy_CertificateManagerMonitoring) AlarmWidgets() *[]awscloudwatch.IWidget {
	var returns *[]awscloudwatch.IWidget

	_jsii_.Invoke(
		c,
		"alarmWidgets",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CertificateManagerMonitoring) CreateAlarmFactory(alarmNamePrefix *string) AlarmFactory {
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

func (c *jsiiProxy_CertificateManagerMonitoring) CreatedAlarms() *[]*AlarmWithAnnotation {
	var returns *[]*AlarmWithAnnotation

	_jsii_.Invoke(
		c,
		"createdAlarms",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CertificateManagerMonitoring) CreateDaysToExpiryWidget(width *float64, height *float64) awscloudwatch.GraphWidget {
	if err := c.validateCreateDaysToExpiryWidgetParameters(width, height); err != nil {
		panic(err)
	}
	var returns awscloudwatch.GraphWidget

	_jsii_.Invoke(
		c,
		"createDaysToExpiryWidget",
		[]interface{}{width, height},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CertificateManagerMonitoring) CreateMetricFactory() MetricFactory {
	var returns MetricFactory

	_jsii_.Invoke(
		c,
		"createMetricFactory",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CertificateManagerMonitoring) CreateTitleWidget() MonitoringHeaderWidget {
	var returns MonitoringHeaderWidget

	_jsii_.Invoke(
		c,
		"createTitleWidget",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CertificateManagerMonitoring) CreateWidgetFactory() IWidgetFactory {
	var returns IWidgetFactory

	_jsii_.Invoke(
		c,
		"createWidgetFactory",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CertificateManagerMonitoring) SummaryWidgets() *[]awscloudwatch.IWidget {
	var returns *[]awscloudwatch.IWidget

	_jsii_.Invoke(
		c,
		"summaryWidgets",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CertificateManagerMonitoring) Widgets() *[]awscloudwatch.IWidget {
	var returns *[]awscloudwatch.IWidget

	_jsii_.Invoke(
		c,
		"widgets",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CertificateManagerMonitoring) WidgetsForDashboard(name *string) *[]awscloudwatch.IWidget {
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


package cdkmonitoringconstructs

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/corymhall/cdk-constructs-go/cdkmonitoringconstructs/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudwatch"
)

// Experimental.
type DocumentDbMonitoring interface {
	Monitoring
	// Experimental.
	Alarms() *[]*AlarmWithAnnotation
	// Experimental.
	ConnectionsMetric() interface{}
	// Experimental.
	CpuUsageMetric() interface{}
	// Experimental.
	CursorsMetric() interface{}
	// Experimental.
	LocalAlarmNamePrefixOverride() *string
	// Experimental.
	ReadLatencyMetric() interface{}
	// Experimental.
	Scope() MonitoringScope
	// Experimental.
	ThrottledMetric() interface{}
	// Experimental.
	Title() *string
	// Experimental.
	TransactionsMetric() interface{}
	// Experimental.
	Url() *string
	// Experimental.
	UsageAlarmFactory() UsageAlarmFactory
	// Experimental.
	UsageAnnotations() *[]*awscloudwatch.HorizontalAnnotation
	// Experimental.
	WriteLatencyMetric() interface{}
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
	CreateConnectionsWidget(width *float64, height *float64) awscloudwatch.GraphWidget
	// Returns all the alarms created.
	// Experimental.
	CreatedAlarms() *[]*AlarmWithAnnotation
	// Experimental.
	CreateLatencyWidget(width *float64, height *float64) awscloudwatch.GraphWidget
	// Creates a new metric factory.
	// Experimental.
	CreateMetricFactory() MetricFactory
	// Experimental.
	CreateResourceUsageWidget(width *float64, height *float64) awscloudwatch.GraphWidget
	// Experimental.
	CreateTitleWidget() MonitoringHeaderWidget
	// Experimental.
	CreateTransactionsWidget(width *float64, height *float64) awscloudwatch.GraphWidget
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

// The jsii proxy struct for DocumentDbMonitoring
type jsiiProxy_DocumentDbMonitoring struct {
	jsiiProxy_Monitoring
}

func (j *jsiiProxy_DocumentDbMonitoring) Alarms() *[]*AlarmWithAnnotation {
	var returns *[]*AlarmWithAnnotation
	_jsii_.Get(
		j,
		"alarms",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DocumentDbMonitoring) ConnectionsMetric() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connectionsMetric",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DocumentDbMonitoring) CpuUsageMetric() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"cpuUsageMetric",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DocumentDbMonitoring) CursorsMetric() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"cursorsMetric",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DocumentDbMonitoring) LocalAlarmNamePrefixOverride() *string {
	var returns *string
	_jsii_.Get(
		j,
		"localAlarmNamePrefixOverride",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DocumentDbMonitoring) ReadLatencyMetric() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"readLatencyMetric",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DocumentDbMonitoring) Scope() MonitoringScope {
	var returns MonitoringScope
	_jsii_.Get(
		j,
		"scope",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DocumentDbMonitoring) ThrottledMetric() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"throttledMetric",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DocumentDbMonitoring) Title() *string {
	var returns *string
	_jsii_.Get(
		j,
		"title",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DocumentDbMonitoring) TransactionsMetric() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"transactionsMetric",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DocumentDbMonitoring) Url() *string {
	var returns *string
	_jsii_.Get(
		j,
		"url",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DocumentDbMonitoring) UsageAlarmFactory() UsageAlarmFactory {
	var returns UsageAlarmFactory
	_jsii_.Get(
		j,
		"usageAlarmFactory",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DocumentDbMonitoring) UsageAnnotations() *[]*awscloudwatch.HorizontalAnnotation {
	var returns *[]*awscloudwatch.HorizontalAnnotation
	_jsii_.Get(
		j,
		"usageAnnotations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DocumentDbMonitoring) WriteLatencyMetric() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"writeLatencyMetric",
		&returns,
	)
	return returns
}


// Experimental.
func NewDocumentDbMonitoring(scope MonitoringScope, props *DocumentDbMonitoringProps) DocumentDbMonitoring {
	_init_.Initialize()

	if err := validateNewDocumentDbMonitoringParameters(scope, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_DocumentDbMonitoring{}

	_jsii_.Create(
		"cdk-monitoring-constructs.DocumentDbMonitoring",
		[]interface{}{scope, props},
		&j,
	)

	return &j
}

// Experimental.
func NewDocumentDbMonitoring_Override(d DocumentDbMonitoring, scope MonitoringScope, props *DocumentDbMonitoringProps) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk-monitoring-constructs.DocumentDbMonitoring",
		[]interface{}{scope, props},
		d,
	)
}

func (d *jsiiProxy_DocumentDbMonitoring) AddAlarm(alarm *AlarmWithAnnotation) {
	if err := d.validateAddAlarmParameters(alarm); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"addAlarm",
		[]interface{}{alarm},
	)
}

func (d *jsiiProxy_DocumentDbMonitoring) AlarmWidgets() *[]awscloudwatch.IWidget {
	var returns *[]awscloudwatch.IWidget

	_jsii_.Invoke(
		d,
		"alarmWidgets",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DocumentDbMonitoring) CreateAlarmFactory(alarmNamePrefix *string) AlarmFactory {
	if err := d.validateCreateAlarmFactoryParameters(alarmNamePrefix); err != nil {
		panic(err)
	}
	var returns AlarmFactory

	_jsii_.Invoke(
		d,
		"createAlarmFactory",
		[]interface{}{alarmNamePrefix},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DocumentDbMonitoring) CreateConnectionsWidget(width *float64, height *float64) awscloudwatch.GraphWidget {
	if err := d.validateCreateConnectionsWidgetParameters(width, height); err != nil {
		panic(err)
	}
	var returns awscloudwatch.GraphWidget

	_jsii_.Invoke(
		d,
		"createConnectionsWidget",
		[]interface{}{width, height},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DocumentDbMonitoring) CreatedAlarms() *[]*AlarmWithAnnotation {
	var returns *[]*AlarmWithAnnotation

	_jsii_.Invoke(
		d,
		"createdAlarms",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DocumentDbMonitoring) CreateLatencyWidget(width *float64, height *float64) awscloudwatch.GraphWidget {
	if err := d.validateCreateLatencyWidgetParameters(width, height); err != nil {
		panic(err)
	}
	var returns awscloudwatch.GraphWidget

	_jsii_.Invoke(
		d,
		"createLatencyWidget",
		[]interface{}{width, height},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DocumentDbMonitoring) CreateMetricFactory() MetricFactory {
	var returns MetricFactory

	_jsii_.Invoke(
		d,
		"createMetricFactory",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DocumentDbMonitoring) CreateResourceUsageWidget(width *float64, height *float64) awscloudwatch.GraphWidget {
	if err := d.validateCreateResourceUsageWidgetParameters(width, height); err != nil {
		panic(err)
	}
	var returns awscloudwatch.GraphWidget

	_jsii_.Invoke(
		d,
		"createResourceUsageWidget",
		[]interface{}{width, height},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DocumentDbMonitoring) CreateTitleWidget() MonitoringHeaderWidget {
	var returns MonitoringHeaderWidget

	_jsii_.Invoke(
		d,
		"createTitleWidget",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DocumentDbMonitoring) CreateTransactionsWidget(width *float64, height *float64) awscloudwatch.GraphWidget {
	if err := d.validateCreateTransactionsWidgetParameters(width, height); err != nil {
		panic(err)
	}
	var returns awscloudwatch.GraphWidget

	_jsii_.Invoke(
		d,
		"createTransactionsWidget",
		[]interface{}{width, height},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DocumentDbMonitoring) CreateWidgetFactory() IWidgetFactory {
	var returns IWidgetFactory

	_jsii_.Invoke(
		d,
		"createWidgetFactory",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DocumentDbMonitoring) SummaryWidgets() *[]awscloudwatch.IWidget {
	var returns *[]awscloudwatch.IWidget

	_jsii_.Invoke(
		d,
		"summaryWidgets",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DocumentDbMonitoring) Widgets() *[]awscloudwatch.IWidget {
	var returns *[]awscloudwatch.IWidget

	_jsii_.Invoke(
		d,
		"widgets",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DocumentDbMonitoring) WidgetsForDashboard(name *string) *[]awscloudwatch.IWidget {
	if err := d.validateWidgetsForDashboardParameters(name); err != nil {
		panic(err)
	}
	var returns *[]awscloudwatch.IWidget

	_jsii_.Invoke(
		d,
		"widgetsForDashboard",
		[]interface{}{name},
		&returns,
	)

	return returns
}


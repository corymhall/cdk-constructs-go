package cdkmonitoringconstructs

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/corymhall/cdk-constructs-go/cdkmonitoringconstructs/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudwatch"
)

// Experimental.
type DynamoTableGlobalSecondaryIndexMonitoring interface {
	Monitoring
	// Experimental.
	Alarms() *[]*AlarmWithAnnotation
	// Experimental.
	ConsumedReadUnitsMetric() interface{}
	// Experimental.
	ConsumedWriteUnitsMetric() interface{}
	// Experimental.
	GsiAlarmFactory() DynamoAlarmFactory
	// Experimental.
	IndexConsumedWriteUnitsMetric() interface{}
	// Experimental.
	IndexThrottleCountMetric() interface{}
	// Experimental.
	LocalAlarmNamePrefixOverride() *string
	// Experimental.
	ProvisionedReadUnitsMetric() interface{}
	// Experimental.
	ProvisionedWriteUnitsMetric() interface{}
	// Experimental.
	ReadThrottleCountMetric() interface{}
	// Experimental.
	Scope() MonitoringScope
	// Experimental.
	TableUrl() *string
	// Experimental.
	ThrottledEventsAnnotations() *[]*awscloudwatch.HorizontalAnnotation
	// Experimental.
	Title() *string
	// Experimental.
	WriteThrottleCountMetric() interface{}
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
	// Experimental.
	CreateReadCapacityWidget(width *float64, height *float64) awscloudwatch.GraphWidget
	// Experimental.
	CreateThrottlesWidget(width *float64, height *float64) awscloudwatch.GraphWidget
	// Experimental.
	CreateTitleWidget() MonitoringHeaderWidget
	// Creates a new widget factory.
	// Experimental.
	CreateWidgetFactory() IWidgetFactory
	// Experimental.
	CreateWriteCapacityWidget(width *float64, height *float64) awscloudwatch.GraphWidget
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

// The jsii proxy struct for DynamoTableGlobalSecondaryIndexMonitoring
type jsiiProxy_DynamoTableGlobalSecondaryIndexMonitoring struct {
	jsiiProxy_Monitoring
}

func (j *jsiiProxy_DynamoTableGlobalSecondaryIndexMonitoring) Alarms() *[]*AlarmWithAnnotation {
	var returns *[]*AlarmWithAnnotation
	_jsii_.Get(
		j,
		"alarms",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamoTableGlobalSecondaryIndexMonitoring) ConsumedReadUnitsMetric() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"consumedReadUnitsMetric",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamoTableGlobalSecondaryIndexMonitoring) ConsumedWriteUnitsMetric() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"consumedWriteUnitsMetric",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamoTableGlobalSecondaryIndexMonitoring) GsiAlarmFactory() DynamoAlarmFactory {
	var returns DynamoAlarmFactory
	_jsii_.Get(
		j,
		"gsiAlarmFactory",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamoTableGlobalSecondaryIndexMonitoring) IndexConsumedWriteUnitsMetric() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"indexConsumedWriteUnitsMetric",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamoTableGlobalSecondaryIndexMonitoring) IndexThrottleCountMetric() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"indexThrottleCountMetric",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamoTableGlobalSecondaryIndexMonitoring) LocalAlarmNamePrefixOverride() *string {
	var returns *string
	_jsii_.Get(
		j,
		"localAlarmNamePrefixOverride",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamoTableGlobalSecondaryIndexMonitoring) ProvisionedReadUnitsMetric() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"provisionedReadUnitsMetric",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamoTableGlobalSecondaryIndexMonitoring) ProvisionedWriteUnitsMetric() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"provisionedWriteUnitsMetric",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamoTableGlobalSecondaryIndexMonitoring) ReadThrottleCountMetric() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"readThrottleCountMetric",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamoTableGlobalSecondaryIndexMonitoring) Scope() MonitoringScope {
	var returns MonitoringScope
	_jsii_.Get(
		j,
		"scope",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamoTableGlobalSecondaryIndexMonitoring) TableUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tableUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamoTableGlobalSecondaryIndexMonitoring) ThrottledEventsAnnotations() *[]*awscloudwatch.HorizontalAnnotation {
	var returns *[]*awscloudwatch.HorizontalAnnotation
	_jsii_.Get(
		j,
		"throttledEventsAnnotations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamoTableGlobalSecondaryIndexMonitoring) Title() *string {
	var returns *string
	_jsii_.Get(
		j,
		"title",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamoTableGlobalSecondaryIndexMonitoring) WriteThrottleCountMetric() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"writeThrottleCountMetric",
		&returns,
	)
	return returns
}


// Experimental.
func NewDynamoTableGlobalSecondaryIndexMonitoring(scope MonitoringScope, props *DynamoTableGlobalSecondaryIndexMonitoringProps) DynamoTableGlobalSecondaryIndexMonitoring {
	_init_.Initialize()

	if err := validateNewDynamoTableGlobalSecondaryIndexMonitoringParameters(scope, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_DynamoTableGlobalSecondaryIndexMonitoring{}

	_jsii_.Create(
		"cdk-monitoring-constructs.DynamoTableGlobalSecondaryIndexMonitoring",
		[]interface{}{scope, props},
		&j,
	)

	return &j
}

// Experimental.
func NewDynamoTableGlobalSecondaryIndexMonitoring_Override(d DynamoTableGlobalSecondaryIndexMonitoring, scope MonitoringScope, props *DynamoTableGlobalSecondaryIndexMonitoringProps) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk-monitoring-constructs.DynamoTableGlobalSecondaryIndexMonitoring",
		[]interface{}{scope, props},
		d,
	)
}

func (d *jsiiProxy_DynamoTableGlobalSecondaryIndexMonitoring) AddAlarm(alarm *AlarmWithAnnotation) {
	if err := d.validateAddAlarmParameters(alarm); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"addAlarm",
		[]interface{}{alarm},
	)
}

func (d *jsiiProxy_DynamoTableGlobalSecondaryIndexMonitoring) AlarmWidgets() *[]awscloudwatch.IWidget {
	var returns *[]awscloudwatch.IWidget

	_jsii_.Invoke(
		d,
		"alarmWidgets",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DynamoTableGlobalSecondaryIndexMonitoring) CreateAlarmFactory(alarmNamePrefix *string) AlarmFactory {
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

func (d *jsiiProxy_DynamoTableGlobalSecondaryIndexMonitoring) CreatedAlarms() *[]*AlarmWithAnnotation {
	var returns *[]*AlarmWithAnnotation

	_jsii_.Invoke(
		d,
		"createdAlarms",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DynamoTableGlobalSecondaryIndexMonitoring) CreateMetricFactory() MetricFactory {
	var returns MetricFactory

	_jsii_.Invoke(
		d,
		"createMetricFactory",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DynamoTableGlobalSecondaryIndexMonitoring) CreateReadCapacityWidget(width *float64, height *float64) awscloudwatch.GraphWidget {
	if err := d.validateCreateReadCapacityWidgetParameters(width, height); err != nil {
		panic(err)
	}
	var returns awscloudwatch.GraphWidget

	_jsii_.Invoke(
		d,
		"createReadCapacityWidget",
		[]interface{}{width, height},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DynamoTableGlobalSecondaryIndexMonitoring) CreateThrottlesWidget(width *float64, height *float64) awscloudwatch.GraphWidget {
	if err := d.validateCreateThrottlesWidgetParameters(width, height); err != nil {
		panic(err)
	}
	var returns awscloudwatch.GraphWidget

	_jsii_.Invoke(
		d,
		"createThrottlesWidget",
		[]interface{}{width, height},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DynamoTableGlobalSecondaryIndexMonitoring) CreateTitleWidget() MonitoringHeaderWidget {
	var returns MonitoringHeaderWidget

	_jsii_.Invoke(
		d,
		"createTitleWidget",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DynamoTableGlobalSecondaryIndexMonitoring) CreateWidgetFactory() IWidgetFactory {
	var returns IWidgetFactory

	_jsii_.Invoke(
		d,
		"createWidgetFactory",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DynamoTableGlobalSecondaryIndexMonitoring) CreateWriteCapacityWidget(width *float64, height *float64) awscloudwatch.GraphWidget {
	if err := d.validateCreateWriteCapacityWidgetParameters(width, height); err != nil {
		panic(err)
	}
	var returns awscloudwatch.GraphWidget

	_jsii_.Invoke(
		d,
		"createWriteCapacityWidget",
		[]interface{}{width, height},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DynamoTableGlobalSecondaryIndexMonitoring) SummaryWidgets() *[]awscloudwatch.IWidget {
	var returns *[]awscloudwatch.IWidget

	_jsii_.Invoke(
		d,
		"summaryWidgets",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DynamoTableGlobalSecondaryIndexMonitoring) Widgets() *[]awscloudwatch.IWidget {
	var returns *[]awscloudwatch.IWidget

	_jsii_.Invoke(
		d,
		"widgets",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DynamoTableGlobalSecondaryIndexMonitoring) WidgetsForDashboard(name *string) *[]awscloudwatch.IWidget {
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


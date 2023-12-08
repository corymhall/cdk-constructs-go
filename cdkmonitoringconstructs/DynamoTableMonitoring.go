package cdkmonitoringconstructs

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/corymhall/cdk-constructs-go/cdkmonitoringconstructs/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudwatch"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsdynamodb"
)

// Experimental.
type DynamoTableMonitoring interface {
	Monitoring
	// Experimental.
	AlarmFactory() AlarmFactory
	// Experimental.
	Alarms() *[]*AlarmWithAnnotation
	// Experimental.
	AveragePerOperationLatencyMetrics() *map[string]interface{}
	// Experimental.
	ConsumedReadUnitsMetric() interface{}
	// Experimental.
	ConsumedWriteUnitsMetric() interface{}
	// Experimental.
	DynamoCapacityAlarmFactory() DynamoAlarmFactory
	// Experimental.
	DynamoReadCapacityAnnotations() *[]*awscloudwatch.HorizontalAnnotation
	// Experimental.
	DynamoWriteCapacityAnnotations() *[]*awscloudwatch.HorizontalAnnotation
	// Experimental.
	ErrorAlarmFactory() ErrorAlarmFactory
	// Experimental.
	ErrorCountAnnotations() *[]*awscloudwatch.HorizontalAnnotation
	// Experimental.
	LatencyAlarmFactory() LatencyAlarmFactory
	// Experimental.
	LatencyAnnotations() *[]*awscloudwatch.HorizontalAnnotation
	// Experimental.
	LatencyAverageSearchMetrics() awscloudwatch.IMetric
	// Experimental.
	LocalAlarmNamePrefixOverride() *string
	// Experimental.
	ProvisionedReadUnitsMetric() interface{}
	// Experimental.
	ProvisionedWriteUnitsMetric() interface{}
	// Experimental.
	ReadCapacityUsageMetric() interface{}
	// Experimental.
	ReadThrottleCountMetric() interface{}
	// Experimental.
	Scope() MonitoringScope
	// Experimental.
	SystemErrorMetric() interface{}
	// Experimental.
	TableBillingMode() awsdynamodb.BillingMode
	// Experimental.
	TableUrl() *string
	// Experimental.
	ThrottledEventsAnnotations() *[]*awscloudwatch.HorizontalAnnotation
	// Experimental.
	Title() *string
	// Experimental.
	WriteCapacityUsageMetric() interface{}
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
	// Experimental.
	CreateErrorsWidget(width *float64, height *float64) awscloudwatch.GraphWidget
	// Experimental.
	CreateLatencyWidget(width *float64, height *float64) awscloudwatch.GraphWidget
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
	// Experimental.
	ForEachOperationLatencyAlarmDefinition(operation awsdynamodb.Operation, alarm *map[string]*LatencyThreshold)
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

// The jsii proxy struct for DynamoTableMonitoring
type jsiiProxy_DynamoTableMonitoring struct {
	jsiiProxy_Monitoring
}

func (j *jsiiProxy_DynamoTableMonitoring) AlarmFactory() AlarmFactory {
	var returns AlarmFactory
	_jsii_.Get(
		j,
		"alarmFactory",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamoTableMonitoring) Alarms() *[]*AlarmWithAnnotation {
	var returns *[]*AlarmWithAnnotation
	_jsii_.Get(
		j,
		"alarms",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamoTableMonitoring) AveragePerOperationLatencyMetrics() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"averagePerOperationLatencyMetrics",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamoTableMonitoring) ConsumedReadUnitsMetric() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"consumedReadUnitsMetric",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamoTableMonitoring) ConsumedWriteUnitsMetric() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"consumedWriteUnitsMetric",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamoTableMonitoring) DynamoCapacityAlarmFactory() DynamoAlarmFactory {
	var returns DynamoAlarmFactory
	_jsii_.Get(
		j,
		"dynamoCapacityAlarmFactory",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamoTableMonitoring) DynamoReadCapacityAnnotations() *[]*awscloudwatch.HorizontalAnnotation {
	var returns *[]*awscloudwatch.HorizontalAnnotation
	_jsii_.Get(
		j,
		"dynamoReadCapacityAnnotations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamoTableMonitoring) DynamoWriteCapacityAnnotations() *[]*awscloudwatch.HorizontalAnnotation {
	var returns *[]*awscloudwatch.HorizontalAnnotation
	_jsii_.Get(
		j,
		"dynamoWriteCapacityAnnotations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamoTableMonitoring) ErrorAlarmFactory() ErrorAlarmFactory {
	var returns ErrorAlarmFactory
	_jsii_.Get(
		j,
		"errorAlarmFactory",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamoTableMonitoring) ErrorCountAnnotations() *[]*awscloudwatch.HorizontalAnnotation {
	var returns *[]*awscloudwatch.HorizontalAnnotation
	_jsii_.Get(
		j,
		"errorCountAnnotations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamoTableMonitoring) LatencyAlarmFactory() LatencyAlarmFactory {
	var returns LatencyAlarmFactory
	_jsii_.Get(
		j,
		"latencyAlarmFactory",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamoTableMonitoring) LatencyAnnotations() *[]*awscloudwatch.HorizontalAnnotation {
	var returns *[]*awscloudwatch.HorizontalAnnotation
	_jsii_.Get(
		j,
		"latencyAnnotations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamoTableMonitoring) LatencyAverageSearchMetrics() awscloudwatch.IMetric {
	var returns awscloudwatch.IMetric
	_jsii_.Get(
		j,
		"latencyAverageSearchMetrics",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamoTableMonitoring) LocalAlarmNamePrefixOverride() *string {
	var returns *string
	_jsii_.Get(
		j,
		"localAlarmNamePrefixOverride",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamoTableMonitoring) ProvisionedReadUnitsMetric() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"provisionedReadUnitsMetric",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamoTableMonitoring) ProvisionedWriteUnitsMetric() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"provisionedWriteUnitsMetric",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamoTableMonitoring) ReadCapacityUsageMetric() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"readCapacityUsageMetric",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamoTableMonitoring) ReadThrottleCountMetric() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"readThrottleCountMetric",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamoTableMonitoring) Scope() MonitoringScope {
	var returns MonitoringScope
	_jsii_.Get(
		j,
		"scope",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamoTableMonitoring) SystemErrorMetric() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"systemErrorMetric",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamoTableMonitoring) TableBillingMode() awsdynamodb.BillingMode {
	var returns awsdynamodb.BillingMode
	_jsii_.Get(
		j,
		"tableBillingMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamoTableMonitoring) TableUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tableUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamoTableMonitoring) ThrottledEventsAnnotations() *[]*awscloudwatch.HorizontalAnnotation {
	var returns *[]*awscloudwatch.HorizontalAnnotation
	_jsii_.Get(
		j,
		"throttledEventsAnnotations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamoTableMonitoring) Title() *string {
	var returns *string
	_jsii_.Get(
		j,
		"title",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamoTableMonitoring) WriteCapacityUsageMetric() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"writeCapacityUsageMetric",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamoTableMonitoring) WriteThrottleCountMetric() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"writeThrottleCountMetric",
		&returns,
	)
	return returns
}


// Experimental.
func NewDynamoTableMonitoring(scope MonitoringScope, props *DynamoTableMonitoringProps) DynamoTableMonitoring {
	_init_.Initialize()

	if err := validateNewDynamoTableMonitoringParameters(scope, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_DynamoTableMonitoring{}

	_jsii_.Create(
		"cdk-monitoring-constructs.DynamoTableMonitoring",
		[]interface{}{scope, props},
		&j,
	)

	return &j
}

// Experimental.
func NewDynamoTableMonitoring_Override(d DynamoTableMonitoring, scope MonitoringScope, props *DynamoTableMonitoringProps) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk-monitoring-constructs.DynamoTableMonitoring",
		[]interface{}{scope, props},
		d,
	)
}

func (d *jsiiProxy_DynamoTableMonitoring) AddAlarm(alarm *AlarmWithAnnotation) {
	if err := d.validateAddAlarmParameters(alarm); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"addAlarm",
		[]interface{}{alarm},
	)
}

func (d *jsiiProxy_DynamoTableMonitoring) AlarmWidgets() *[]awscloudwatch.IWidget {
	var returns *[]awscloudwatch.IWidget

	_jsii_.Invoke(
		d,
		"alarmWidgets",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DynamoTableMonitoring) CreateAlarmFactory(alarmNamePrefix *string) AlarmFactory {
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

func (d *jsiiProxy_DynamoTableMonitoring) CreatedAlarms() *[]*AlarmWithAnnotation {
	var returns *[]*AlarmWithAnnotation

	_jsii_.Invoke(
		d,
		"createdAlarms",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DynamoTableMonitoring) CreateErrorsWidget(width *float64, height *float64) awscloudwatch.GraphWidget {
	if err := d.validateCreateErrorsWidgetParameters(width, height); err != nil {
		panic(err)
	}
	var returns awscloudwatch.GraphWidget

	_jsii_.Invoke(
		d,
		"createErrorsWidget",
		[]interface{}{width, height},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DynamoTableMonitoring) CreateLatencyWidget(width *float64, height *float64) awscloudwatch.GraphWidget {
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

func (d *jsiiProxy_DynamoTableMonitoring) CreateMetricFactory() MetricFactory {
	var returns MetricFactory

	_jsii_.Invoke(
		d,
		"createMetricFactory",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DynamoTableMonitoring) CreateReadCapacityWidget(width *float64, height *float64) awscloudwatch.GraphWidget {
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

func (d *jsiiProxy_DynamoTableMonitoring) CreateThrottlesWidget(width *float64, height *float64) awscloudwatch.GraphWidget {
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

func (d *jsiiProxy_DynamoTableMonitoring) CreateTitleWidget() MonitoringHeaderWidget {
	var returns MonitoringHeaderWidget

	_jsii_.Invoke(
		d,
		"createTitleWidget",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DynamoTableMonitoring) CreateWidgetFactory() IWidgetFactory {
	var returns IWidgetFactory

	_jsii_.Invoke(
		d,
		"createWidgetFactory",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DynamoTableMonitoring) CreateWriteCapacityWidget(width *float64, height *float64) awscloudwatch.GraphWidget {
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

func (d *jsiiProxy_DynamoTableMonitoring) ForEachOperationLatencyAlarmDefinition(operation awsdynamodb.Operation, alarm *map[string]*LatencyThreshold) {
	if err := d.validateForEachOperationLatencyAlarmDefinitionParameters(operation, alarm); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"forEachOperationLatencyAlarmDefinition",
		[]interface{}{operation, alarm},
	)
}

func (d *jsiiProxy_DynamoTableMonitoring) SummaryWidgets() *[]awscloudwatch.IWidget {
	var returns *[]awscloudwatch.IWidget

	_jsii_.Invoke(
		d,
		"summaryWidgets",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DynamoTableMonitoring) Widgets() *[]awscloudwatch.IWidget {
	var returns *[]awscloudwatch.IWidget

	_jsii_.Invoke(
		d,
		"widgets",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DynamoTableMonitoring) WidgetsForDashboard(name *string) *[]awscloudwatch.IWidget {
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

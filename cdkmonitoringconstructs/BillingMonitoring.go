package cdkmonitoringconstructs

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/corymhall/cdk-constructs-go/cdkmonitoringconstructs/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudwatch"
)

// Experimental.
type BillingMonitoring interface {
	Monitoring
	// Experimental.
	AlarmFactory() AlarmFactory
	// Experimental.
	Alarms() *[]*AlarmWithAnnotation
	// Experimental.
	AnomalyDetectingAlarmFactory() AnomalyDetectingAlarmFactory
	// Experimental.
	CostByServiceMetric() awscloudwatch.IMetric
	// Experimental.
	LocalAlarmNamePrefixOverride() *string
	// Experimental.
	Scope() MonitoringScope
	// Experimental.
	Title() *string
	// Experimental.
	TotalCostMetric() interface{}
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
	CreateChargesByServiceWidget(width *float64, height *float64) awscloudwatch.GraphWidget
	// Returns all the alarms created.
	// Experimental.
	CreatedAlarms() *[]*AlarmWithAnnotation
	// Creates a new metric factory.
	// Experimental.
	CreateMetricFactory() MetricFactory
	// Experimental.
	CreateTitleWidget() MonitoringHeaderWidget
	// Experimental.
	CreateTotalChargesWidget(width *float64, height *float64) awscloudwatch.SingleValueWidget
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

// The jsii proxy struct for BillingMonitoring
type jsiiProxy_BillingMonitoring struct {
	jsiiProxy_Monitoring
}

func (j *jsiiProxy_BillingMonitoring) AlarmFactory() AlarmFactory {
	var returns AlarmFactory
	_jsii_.Get(
		j,
		"alarmFactory",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BillingMonitoring) Alarms() *[]*AlarmWithAnnotation {
	var returns *[]*AlarmWithAnnotation
	_jsii_.Get(
		j,
		"alarms",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BillingMonitoring) AnomalyDetectingAlarmFactory() AnomalyDetectingAlarmFactory {
	var returns AnomalyDetectingAlarmFactory
	_jsii_.Get(
		j,
		"anomalyDetectingAlarmFactory",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BillingMonitoring) CostByServiceMetric() awscloudwatch.IMetric {
	var returns awscloudwatch.IMetric
	_jsii_.Get(
		j,
		"costByServiceMetric",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BillingMonitoring) LocalAlarmNamePrefixOverride() *string {
	var returns *string
	_jsii_.Get(
		j,
		"localAlarmNamePrefixOverride",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BillingMonitoring) Scope() MonitoringScope {
	var returns MonitoringScope
	_jsii_.Get(
		j,
		"scope",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BillingMonitoring) Title() *string {
	var returns *string
	_jsii_.Get(
		j,
		"title",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BillingMonitoring) TotalCostMetric() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"totalCostMetric",
		&returns,
	)
	return returns
}


// Experimental.
func NewBillingMonitoring(scope MonitoringScope, props *BillingMonitoringProps) BillingMonitoring {
	_init_.Initialize()

	if err := validateNewBillingMonitoringParameters(scope, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_BillingMonitoring{}

	_jsii_.Create(
		"cdk-monitoring-constructs.BillingMonitoring",
		[]interface{}{scope, props},
		&j,
	)

	return &j
}

// Experimental.
func NewBillingMonitoring_Override(b BillingMonitoring, scope MonitoringScope, props *BillingMonitoringProps) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk-monitoring-constructs.BillingMonitoring",
		[]interface{}{scope, props},
		b,
	)
}

func (b *jsiiProxy_BillingMonitoring) AddAlarm(alarm *AlarmWithAnnotation) {
	if err := b.validateAddAlarmParameters(alarm); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"addAlarm",
		[]interface{}{alarm},
	)
}

func (b *jsiiProxy_BillingMonitoring) AlarmWidgets() *[]awscloudwatch.IWidget {
	var returns *[]awscloudwatch.IWidget

	_jsii_.Invoke(
		b,
		"alarmWidgets",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BillingMonitoring) CreateAlarmFactory(alarmNamePrefix *string) AlarmFactory {
	if err := b.validateCreateAlarmFactoryParameters(alarmNamePrefix); err != nil {
		panic(err)
	}
	var returns AlarmFactory

	_jsii_.Invoke(
		b,
		"createAlarmFactory",
		[]interface{}{alarmNamePrefix},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BillingMonitoring) CreateChargesByServiceWidget(width *float64, height *float64) awscloudwatch.GraphWidget {
	if err := b.validateCreateChargesByServiceWidgetParameters(width, height); err != nil {
		panic(err)
	}
	var returns awscloudwatch.GraphWidget

	_jsii_.Invoke(
		b,
		"createChargesByServiceWidget",
		[]interface{}{width, height},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BillingMonitoring) CreatedAlarms() *[]*AlarmWithAnnotation {
	var returns *[]*AlarmWithAnnotation

	_jsii_.Invoke(
		b,
		"createdAlarms",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BillingMonitoring) CreateMetricFactory() MetricFactory {
	var returns MetricFactory

	_jsii_.Invoke(
		b,
		"createMetricFactory",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BillingMonitoring) CreateTitleWidget() MonitoringHeaderWidget {
	var returns MonitoringHeaderWidget

	_jsii_.Invoke(
		b,
		"createTitleWidget",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BillingMonitoring) CreateTotalChargesWidget(width *float64, height *float64) awscloudwatch.SingleValueWidget {
	if err := b.validateCreateTotalChargesWidgetParameters(width, height); err != nil {
		panic(err)
	}
	var returns awscloudwatch.SingleValueWidget

	_jsii_.Invoke(
		b,
		"createTotalChargesWidget",
		[]interface{}{width, height},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BillingMonitoring) CreateWidgetFactory() IWidgetFactory {
	var returns IWidgetFactory

	_jsii_.Invoke(
		b,
		"createWidgetFactory",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BillingMonitoring) SummaryWidgets() *[]awscloudwatch.IWidget {
	var returns *[]awscloudwatch.IWidget

	_jsii_.Invoke(
		b,
		"summaryWidgets",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BillingMonitoring) Widgets() *[]awscloudwatch.IWidget {
	var returns *[]awscloudwatch.IWidget

	_jsii_.Invoke(
		b,
		"widgets",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BillingMonitoring) WidgetsForDashboard(name *string) *[]awscloudwatch.IWidget {
	if err := b.validateWidgetsForDashboardParameters(name); err != nil {
		panic(err)
	}
	var returns *[]awscloudwatch.IWidget

	_jsii_.Invoke(
		b,
		"widgetsForDashboard",
		[]interface{}{name},
		&returns,
	)

	return returns
}


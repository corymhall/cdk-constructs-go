package cdkmonitoringconstructs

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/corymhall/cdk-constructs-go/cdkmonitoringconstructs/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudwatch"
)

// Monitoring for AWS Web Application Firewall.
// See: https://docs.aws.amazon.com/waf/latest/developerguide/monitoring-cloudwatch.html
//
// Experimental.
type WafV2Monitoring interface {
	Monitoring
	// Experimental.
	AlarmFactory() AlarmFactory
	// Experimental.
	Alarms() *[]*AlarmWithAnnotation
	// Experimental.
	AllowedRequestsMetric() interface{}
	// Experimental.
	BlockedRequestsMetric() interface{}
	// Experimental.
	BlockedRequestsRateMetric() interface{}
	// Experimental.
	ErrorAlarmFactory() ErrorAlarmFactory
	// Experimental.
	ErrorCountAnnotations() *[]*awscloudwatch.HorizontalAnnotation
	// Experimental.
	ErrorRateAnnotations() *[]*awscloudwatch.HorizontalAnnotation
	// Experimental.
	HumanReadableName() *string
	// Experimental.
	LocalAlarmNamePrefixOverride() *string
	// Experimental.
	Scope() MonitoringScope
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
	CreateAllowedRequestsWidget(width *float64, height *float64) awscloudwatch.GraphWidget
	// Experimental.
	CreateBlockedRequestsRateWidget(width *float64, height *float64) awscloudwatch.GraphWidget
	// Experimental.
	CreateBlockedRequestsWidget(width *float64, height *float64) awscloudwatch.GraphWidget
	// Returns all the alarms created.
	// Experimental.
	CreatedAlarms() *[]*AlarmWithAnnotation
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

// The jsii proxy struct for WafV2Monitoring
type jsiiProxy_WafV2Monitoring struct {
	jsiiProxy_Monitoring
}

func (j *jsiiProxy_WafV2Monitoring) AlarmFactory() AlarmFactory {
	var returns AlarmFactory
	_jsii_.Get(
		j,
		"alarmFactory",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WafV2Monitoring) Alarms() *[]*AlarmWithAnnotation {
	var returns *[]*AlarmWithAnnotation
	_jsii_.Get(
		j,
		"alarms",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WafV2Monitoring) AllowedRequestsMetric() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowedRequestsMetric",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WafV2Monitoring) BlockedRequestsMetric() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"blockedRequestsMetric",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WafV2Monitoring) BlockedRequestsRateMetric() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"blockedRequestsRateMetric",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WafV2Monitoring) ErrorAlarmFactory() ErrorAlarmFactory {
	var returns ErrorAlarmFactory
	_jsii_.Get(
		j,
		"errorAlarmFactory",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WafV2Monitoring) ErrorCountAnnotations() *[]*awscloudwatch.HorizontalAnnotation {
	var returns *[]*awscloudwatch.HorizontalAnnotation
	_jsii_.Get(
		j,
		"errorCountAnnotations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WafV2Monitoring) ErrorRateAnnotations() *[]*awscloudwatch.HorizontalAnnotation {
	var returns *[]*awscloudwatch.HorizontalAnnotation
	_jsii_.Get(
		j,
		"errorRateAnnotations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WafV2Monitoring) HumanReadableName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"humanReadableName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WafV2Monitoring) LocalAlarmNamePrefixOverride() *string {
	var returns *string
	_jsii_.Get(
		j,
		"localAlarmNamePrefixOverride",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WafV2Monitoring) Scope() MonitoringScope {
	var returns MonitoringScope
	_jsii_.Get(
		j,
		"scope",
		&returns,
	)
	return returns
}


// Experimental.
func NewWafV2Monitoring(scope MonitoringScope, props *WafV2MonitoringProps) WafV2Monitoring {
	_init_.Initialize()

	if err := validateNewWafV2MonitoringParameters(scope, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_WafV2Monitoring{}

	_jsii_.Create(
		"cdk-monitoring-constructs.WafV2Monitoring",
		[]interface{}{scope, props},
		&j,
	)

	return &j
}

// Experimental.
func NewWafV2Monitoring_Override(w WafV2Monitoring, scope MonitoringScope, props *WafV2MonitoringProps) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk-monitoring-constructs.WafV2Monitoring",
		[]interface{}{scope, props},
		w,
	)
}

func (w *jsiiProxy_WafV2Monitoring) AddAlarm(alarm *AlarmWithAnnotation) {
	if err := w.validateAddAlarmParameters(alarm); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"addAlarm",
		[]interface{}{alarm},
	)
}

func (w *jsiiProxy_WafV2Monitoring) AlarmWidgets() *[]awscloudwatch.IWidget {
	var returns *[]awscloudwatch.IWidget

	_jsii_.Invoke(
		w,
		"alarmWidgets",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WafV2Monitoring) CreateAlarmFactory(alarmNamePrefix *string) AlarmFactory {
	if err := w.validateCreateAlarmFactoryParameters(alarmNamePrefix); err != nil {
		panic(err)
	}
	var returns AlarmFactory

	_jsii_.Invoke(
		w,
		"createAlarmFactory",
		[]interface{}{alarmNamePrefix},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WafV2Monitoring) CreateAllowedRequestsWidget(width *float64, height *float64) awscloudwatch.GraphWidget {
	if err := w.validateCreateAllowedRequestsWidgetParameters(width, height); err != nil {
		panic(err)
	}
	var returns awscloudwatch.GraphWidget

	_jsii_.Invoke(
		w,
		"createAllowedRequestsWidget",
		[]interface{}{width, height},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WafV2Monitoring) CreateBlockedRequestsRateWidget(width *float64, height *float64) awscloudwatch.GraphWidget {
	if err := w.validateCreateBlockedRequestsRateWidgetParameters(width, height); err != nil {
		panic(err)
	}
	var returns awscloudwatch.GraphWidget

	_jsii_.Invoke(
		w,
		"createBlockedRequestsRateWidget",
		[]interface{}{width, height},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WafV2Monitoring) CreateBlockedRequestsWidget(width *float64, height *float64) awscloudwatch.GraphWidget {
	if err := w.validateCreateBlockedRequestsWidgetParameters(width, height); err != nil {
		panic(err)
	}
	var returns awscloudwatch.GraphWidget

	_jsii_.Invoke(
		w,
		"createBlockedRequestsWidget",
		[]interface{}{width, height},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WafV2Monitoring) CreatedAlarms() *[]*AlarmWithAnnotation {
	var returns *[]*AlarmWithAnnotation

	_jsii_.Invoke(
		w,
		"createdAlarms",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WafV2Monitoring) CreateMetricFactory() MetricFactory {
	var returns MetricFactory

	_jsii_.Invoke(
		w,
		"createMetricFactory",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WafV2Monitoring) CreateTitleWidget() MonitoringHeaderWidget {
	var returns MonitoringHeaderWidget

	_jsii_.Invoke(
		w,
		"createTitleWidget",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WafV2Monitoring) CreateWidgetFactory() IWidgetFactory {
	var returns IWidgetFactory

	_jsii_.Invoke(
		w,
		"createWidgetFactory",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WafV2Monitoring) SummaryWidgets() *[]awscloudwatch.IWidget {
	var returns *[]awscloudwatch.IWidget

	_jsii_.Invoke(
		w,
		"summaryWidgets",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WafV2Monitoring) Widgets() *[]awscloudwatch.IWidget {
	var returns *[]awscloudwatch.IWidget

	_jsii_.Invoke(
		w,
		"widgets",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WafV2Monitoring) WidgetsForDashboard(name *string) *[]awscloudwatch.IWidget {
	if err := w.validateWidgetsForDashboardParameters(name); err != nil {
		panic(err)
	}
	var returns *[]awscloudwatch.IWidget

	_jsii_.Invoke(
		w,
		"widgetsForDashboard",
		[]interface{}{name},
		&returns,
	)

	return returns
}


package cdkmonitoringconstructs

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/corymhall/cdk-constructs-go/cdkmonitoringconstructs/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudwatch"
)

// Experimental.
type AppSyncMonitoring interface {
	Monitoring
	// Experimental.
	AlarmFactory() AlarmFactory
	// Experimental.
	Alarms() *[]*AlarmWithAnnotation
	// Experimental.
	Error4xxCountMetric() interface{}
	// Experimental.
	Error4xxRateMetric() interface{}
	// Experimental.
	ErrorAlarmFactory() ErrorAlarmFactory
	// Experimental.
	ErrorCountAnnotations() *[]*awscloudwatch.HorizontalAnnotation
	// Experimental.
	ErrorRateAnnotations() *[]*awscloudwatch.HorizontalAnnotation
	// Experimental.
	Fault5xxCountMetric() interface{}
	// Experimental.
	Fault5xxRateMetric() interface{}
	// Experimental.
	LatencyAlarmFactory() LatencyAlarmFactory
	// Experimental.
	LatencyAnnotations() *[]*awscloudwatch.HorizontalAnnotation
	// Experimental.
	LocalAlarmNamePrefixOverride() *string
	// Experimental.
	MetricFactory() AppSyncMetricFactory
	// Experimental.
	NamingStrategy() MonitoringNamingStrategy
	// Experimental.
	P50LatencyMetric() interface{}
	// Experimental.
	P90LatencyMetric() interface{}
	// Experimental.
	P99LatencyMetric() interface{}
	// Experimental.
	Scope() MonitoringScope
	// Experimental.
	Title() *string
	// Experimental.
	TpsAlarmFactory() TpsAlarmFactory
	// Experimental.
	TpsAnnotations() *[]*awscloudwatch.HorizontalAnnotation
	// Experimental.
	TpsMetric() interface{}
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
	CreateErrorCountWidget(width *float64, height *float64) awscloudwatch.GraphWidget
	// Experimental.
	CreateErrorRateWidget(width *float64, height *float64) awscloudwatch.GraphWidget
	// Experimental.
	CreateLatencyWidget(width *float64, height *float64) awscloudwatch.GraphWidget
	// Creates a new metric factory.
	// Experimental.
	CreateMetricFactory() MetricFactory
	// Experimental.
	CreateTpsWidget(width *float64, height *float64) awscloudwatch.GraphWidget
	// Experimental.
	CreatetTitleWidget() MonitoringHeaderWidget
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

// The jsii proxy struct for AppSyncMonitoring
type jsiiProxy_AppSyncMonitoring struct {
	jsiiProxy_Monitoring
}

func (j *jsiiProxy_AppSyncMonitoring) AlarmFactory() AlarmFactory {
	var returns AlarmFactory
	_jsii_.Get(
		j,
		"alarmFactory",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSyncMonitoring) Alarms() *[]*AlarmWithAnnotation {
	var returns *[]*AlarmWithAnnotation
	_jsii_.Get(
		j,
		"alarms",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSyncMonitoring) Error4xxCountMetric() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"error4xxCountMetric",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSyncMonitoring) Error4xxRateMetric() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"error4xxRateMetric",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSyncMonitoring) ErrorAlarmFactory() ErrorAlarmFactory {
	var returns ErrorAlarmFactory
	_jsii_.Get(
		j,
		"errorAlarmFactory",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSyncMonitoring) ErrorCountAnnotations() *[]*awscloudwatch.HorizontalAnnotation {
	var returns *[]*awscloudwatch.HorizontalAnnotation
	_jsii_.Get(
		j,
		"errorCountAnnotations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSyncMonitoring) ErrorRateAnnotations() *[]*awscloudwatch.HorizontalAnnotation {
	var returns *[]*awscloudwatch.HorizontalAnnotation
	_jsii_.Get(
		j,
		"errorRateAnnotations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSyncMonitoring) Fault5xxCountMetric() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"fault5xxCountMetric",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSyncMonitoring) Fault5xxRateMetric() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"fault5xxRateMetric",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSyncMonitoring) LatencyAlarmFactory() LatencyAlarmFactory {
	var returns LatencyAlarmFactory
	_jsii_.Get(
		j,
		"latencyAlarmFactory",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSyncMonitoring) LatencyAnnotations() *[]*awscloudwatch.HorizontalAnnotation {
	var returns *[]*awscloudwatch.HorizontalAnnotation
	_jsii_.Get(
		j,
		"latencyAnnotations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSyncMonitoring) LocalAlarmNamePrefixOverride() *string {
	var returns *string
	_jsii_.Get(
		j,
		"localAlarmNamePrefixOverride",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSyncMonitoring) MetricFactory() AppSyncMetricFactory {
	var returns AppSyncMetricFactory
	_jsii_.Get(
		j,
		"metricFactory",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSyncMonitoring) NamingStrategy() MonitoringNamingStrategy {
	var returns MonitoringNamingStrategy
	_jsii_.Get(
		j,
		"namingStrategy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSyncMonitoring) P50LatencyMetric() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"p50LatencyMetric",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSyncMonitoring) P90LatencyMetric() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"p90LatencyMetric",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSyncMonitoring) P99LatencyMetric() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"p99LatencyMetric",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSyncMonitoring) Scope() MonitoringScope {
	var returns MonitoringScope
	_jsii_.Get(
		j,
		"scope",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSyncMonitoring) Title() *string {
	var returns *string
	_jsii_.Get(
		j,
		"title",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSyncMonitoring) TpsAlarmFactory() TpsAlarmFactory {
	var returns TpsAlarmFactory
	_jsii_.Get(
		j,
		"tpsAlarmFactory",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSyncMonitoring) TpsAnnotations() *[]*awscloudwatch.HorizontalAnnotation {
	var returns *[]*awscloudwatch.HorizontalAnnotation
	_jsii_.Get(
		j,
		"tpsAnnotations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppSyncMonitoring) TpsMetric() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tpsMetric",
		&returns,
	)
	return returns
}


// Experimental.
func NewAppSyncMonitoring(scope MonitoringScope, props *AppSyncMonitoringProps) AppSyncMonitoring {
	_init_.Initialize()

	if err := validateNewAppSyncMonitoringParameters(scope, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_AppSyncMonitoring{}

	_jsii_.Create(
		"cdk-monitoring-constructs.AppSyncMonitoring",
		[]interface{}{scope, props},
		&j,
	)

	return &j
}

// Experimental.
func NewAppSyncMonitoring_Override(a AppSyncMonitoring, scope MonitoringScope, props *AppSyncMonitoringProps) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk-monitoring-constructs.AppSyncMonitoring",
		[]interface{}{scope, props},
		a,
	)
}

func (a *jsiiProxy_AppSyncMonitoring) AddAlarm(alarm *AlarmWithAnnotation) {
	if err := a.validateAddAlarmParameters(alarm); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addAlarm",
		[]interface{}{alarm},
	)
}

func (a *jsiiProxy_AppSyncMonitoring) AlarmWidgets() *[]awscloudwatch.IWidget {
	var returns *[]awscloudwatch.IWidget

	_jsii_.Invoke(
		a,
		"alarmWidgets",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppSyncMonitoring) CreateAlarmFactory(alarmNamePrefix *string) AlarmFactory {
	if err := a.validateCreateAlarmFactoryParameters(alarmNamePrefix); err != nil {
		panic(err)
	}
	var returns AlarmFactory

	_jsii_.Invoke(
		a,
		"createAlarmFactory",
		[]interface{}{alarmNamePrefix},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppSyncMonitoring) CreatedAlarms() *[]*AlarmWithAnnotation {
	var returns *[]*AlarmWithAnnotation

	_jsii_.Invoke(
		a,
		"createdAlarms",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppSyncMonitoring) CreateErrorCountWidget(width *float64, height *float64) awscloudwatch.GraphWidget {
	if err := a.validateCreateErrorCountWidgetParameters(width, height); err != nil {
		panic(err)
	}
	var returns awscloudwatch.GraphWidget

	_jsii_.Invoke(
		a,
		"createErrorCountWidget",
		[]interface{}{width, height},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppSyncMonitoring) CreateErrorRateWidget(width *float64, height *float64) awscloudwatch.GraphWidget {
	if err := a.validateCreateErrorRateWidgetParameters(width, height); err != nil {
		panic(err)
	}
	var returns awscloudwatch.GraphWidget

	_jsii_.Invoke(
		a,
		"createErrorRateWidget",
		[]interface{}{width, height},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppSyncMonitoring) CreateLatencyWidget(width *float64, height *float64) awscloudwatch.GraphWidget {
	if err := a.validateCreateLatencyWidgetParameters(width, height); err != nil {
		panic(err)
	}
	var returns awscloudwatch.GraphWidget

	_jsii_.Invoke(
		a,
		"createLatencyWidget",
		[]interface{}{width, height},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppSyncMonitoring) CreateMetricFactory() MetricFactory {
	var returns MetricFactory

	_jsii_.Invoke(
		a,
		"createMetricFactory",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppSyncMonitoring) CreateTpsWidget(width *float64, height *float64) awscloudwatch.GraphWidget {
	if err := a.validateCreateTpsWidgetParameters(width, height); err != nil {
		panic(err)
	}
	var returns awscloudwatch.GraphWidget

	_jsii_.Invoke(
		a,
		"createTpsWidget",
		[]interface{}{width, height},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppSyncMonitoring) CreatetTitleWidget() MonitoringHeaderWidget {
	var returns MonitoringHeaderWidget

	_jsii_.Invoke(
		a,
		"createtTitleWidget",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppSyncMonitoring) CreateWidgetFactory() IWidgetFactory {
	var returns IWidgetFactory

	_jsii_.Invoke(
		a,
		"createWidgetFactory",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppSyncMonitoring) SummaryWidgets() *[]awscloudwatch.IWidget {
	var returns *[]awscloudwatch.IWidget

	_jsii_.Invoke(
		a,
		"summaryWidgets",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppSyncMonitoring) Widgets() *[]awscloudwatch.IWidget {
	var returns *[]awscloudwatch.IWidget

	_jsii_.Invoke(
		a,
		"widgets",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppSyncMonitoring) WidgetsForDashboard(name *string) *[]awscloudwatch.IWidget {
	if err := a.validateWidgetsForDashboardParameters(name); err != nil {
		panic(err)
	}
	var returns *[]awscloudwatch.IWidget

	_jsii_.Invoke(
		a,
		"widgetsForDashboard",
		[]interface{}{name},
		&returns,
	)

	return returns
}


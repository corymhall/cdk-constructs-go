package cdkmonitoringconstructs

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/corymhall/cdk-constructs-go/cdkmonitoringconstructs/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudwatch"
)

// Experimental.
type ApiGatewayMonitoring interface {
	Monitoring
	// Experimental.
	AlarmFactory() AlarmFactory
	// Experimental.
	Alarms() *[]*AlarmWithAnnotation
	// Experimental.
	Error4XXCountMetric() interface{}
	// Experimental.
	Error4XXRateMetric() interface{}
	// Experimental.
	ErrorAlarmFactory() ErrorAlarmFactory
	// Experimental.
	ErrorCountAnnotations() *[]*awscloudwatch.HorizontalAnnotation
	// Experimental.
	ErrorRateAnnotations() *[]*awscloudwatch.HorizontalAnnotation
	// Experimental.
	Fault5XXCountMetric() interface{}
	// Experimental.
	Fault5XXRateMetric() interface{}
	// Experimental.
	LatencyAlarmFactory() LatencyAlarmFactory
	// Experimental.
	LatencyAnnotations() *[]*awscloudwatch.HorizontalAnnotation
	// Experimental.
	LatencyMetrics() *map[string]interface{}
	// Experimental.
	LatencyTypesToRender() *[]*string
	// Experimental.
	LocalAlarmNamePrefixOverride() *string
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
	CreateTitleWidget() MonitoringHeaderWidget
	// Experimental.
	CreateTpsWidget(width *float64, height *float64) awscloudwatch.GraphWidget
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

// The jsii proxy struct for ApiGatewayMonitoring
type jsiiProxy_ApiGatewayMonitoring struct {
	jsiiProxy_Monitoring
}

func (j *jsiiProxy_ApiGatewayMonitoring) AlarmFactory() AlarmFactory {
	var returns AlarmFactory
	_jsii_.Get(
		j,
		"alarmFactory",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayMonitoring) Alarms() *[]*AlarmWithAnnotation {
	var returns *[]*AlarmWithAnnotation
	_jsii_.Get(
		j,
		"alarms",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayMonitoring) Error4XXCountMetric() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"error4XXCountMetric",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayMonitoring) Error4XXRateMetric() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"error4XXRateMetric",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayMonitoring) ErrorAlarmFactory() ErrorAlarmFactory {
	var returns ErrorAlarmFactory
	_jsii_.Get(
		j,
		"errorAlarmFactory",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayMonitoring) ErrorCountAnnotations() *[]*awscloudwatch.HorizontalAnnotation {
	var returns *[]*awscloudwatch.HorizontalAnnotation
	_jsii_.Get(
		j,
		"errorCountAnnotations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayMonitoring) ErrorRateAnnotations() *[]*awscloudwatch.HorizontalAnnotation {
	var returns *[]*awscloudwatch.HorizontalAnnotation
	_jsii_.Get(
		j,
		"errorRateAnnotations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayMonitoring) Fault5XXCountMetric() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"fault5XXCountMetric",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayMonitoring) Fault5XXRateMetric() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"fault5XXRateMetric",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayMonitoring) LatencyAlarmFactory() LatencyAlarmFactory {
	var returns LatencyAlarmFactory
	_jsii_.Get(
		j,
		"latencyAlarmFactory",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayMonitoring) LatencyAnnotations() *[]*awscloudwatch.HorizontalAnnotation {
	var returns *[]*awscloudwatch.HorizontalAnnotation
	_jsii_.Get(
		j,
		"latencyAnnotations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayMonitoring) LatencyMetrics() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"latencyMetrics",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayMonitoring) LatencyTypesToRender() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"latencyTypesToRender",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayMonitoring) LocalAlarmNamePrefixOverride() *string {
	var returns *string
	_jsii_.Get(
		j,
		"localAlarmNamePrefixOverride",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayMonitoring) Scope() MonitoringScope {
	var returns MonitoringScope
	_jsii_.Get(
		j,
		"scope",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayMonitoring) Title() *string {
	var returns *string
	_jsii_.Get(
		j,
		"title",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayMonitoring) TpsAlarmFactory() TpsAlarmFactory {
	var returns TpsAlarmFactory
	_jsii_.Get(
		j,
		"tpsAlarmFactory",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayMonitoring) TpsAnnotations() *[]*awscloudwatch.HorizontalAnnotation {
	var returns *[]*awscloudwatch.HorizontalAnnotation
	_jsii_.Get(
		j,
		"tpsAnnotations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayMonitoring) TpsMetric() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tpsMetric",
		&returns,
	)
	return returns
}


// Experimental.
func NewApiGatewayMonitoring(scope MonitoringScope, props *ApiGatewayMonitoringProps) ApiGatewayMonitoring {
	_init_.Initialize()

	if err := validateNewApiGatewayMonitoringParameters(scope, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_ApiGatewayMonitoring{}

	_jsii_.Create(
		"cdk-monitoring-constructs.ApiGatewayMonitoring",
		[]interface{}{scope, props},
		&j,
	)

	return &j
}

// Experimental.
func NewApiGatewayMonitoring_Override(a ApiGatewayMonitoring, scope MonitoringScope, props *ApiGatewayMonitoringProps) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk-monitoring-constructs.ApiGatewayMonitoring",
		[]interface{}{scope, props},
		a,
	)
}

func (a *jsiiProxy_ApiGatewayMonitoring) AddAlarm(alarm *AlarmWithAnnotation) {
	if err := a.validateAddAlarmParameters(alarm); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addAlarm",
		[]interface{}{alarm},
	)
}

func (a *jsiiProxy_ApiGatewayMonitoring) AlarmWidgets() *[]awscloudwatch.IWidget {
	var returns *[]awscloudwatch.IWidget

	_jsii_.Invoke(
		a,
		"alarmWidgets",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiGatewayMonitoring) CreateAlarmFactory(alarmNamePrefix *string) AlarmFactory {
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

func (a *jsiiProxy_ApiGatewayMonitoring) CreatedAlarms() *[]*AlarmWithAnnotation {
	var returns *[]*AlarmWithAnnotation

	_jsii_.Invoke(
		a,
		"createdAlarms",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiGatewayMonitoring) CreateErrorCountWidget(width *float64, height *float64) awscloudwatch.GraphWidget {
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

func (a *jsiiProxy_ApiGatewayMonitoring) CreateErrorRateWidget(width *float64, height *float64) awscloudwatch.GraphWidget {
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

func (a *jsiiProxy_ApiGatewayMonitoring) CreateLatencyWidget(width *float64, height *float64) awscloudwatch.GraphWidget {
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

func (a *jsiiProxy_ApiGatewayMonitoring) CreateMetricFactory() MetricFactory {
	var returns MetricFactory

	_jsii_.Invoke(
		a,
		"createMetricFactory",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiGatewayMonitoring) CreateTitleWidget() MonitoringHeaderWidget {
	var returns MonitoringHeaderWidget

	_jsii_.Invoke(
		a,
		"createTitleWidget",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiGatewayMonitoring) CreateTpsWidget(width *float64, height *float64) awscloudwatch.GraphWidget {
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

func (a *jsiiProxy_ApiGatewayMonitoring) CreateWidgetFactory() IWidgetFactory {
	var returns IWidgetFactory

	_jsii_.Invoke(
		a,
		"createWidgetFactory",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiGatewayMonitoring) SummaryWidgets() *[]awscloudwatch.IWidget {
	var returns *[]awscloudwatch.IWidget

	_jsii_.Invoke(
		a,
		"summaryWidgets",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiGatewayMonitoring) Widgets() *[]awscloudwatch.IWidget {
	var returns *[]awscloudwatch.IWidget

	_jsii_.Invoke(
		a,
		"widgets",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiGatewayMonitoring) WidgetsForDashboard(name *string) *[]awscloudwatch.IWidget {
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


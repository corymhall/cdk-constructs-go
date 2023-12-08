package cdkmonitoringconstructs

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/corymhall/cdk-constructs-go/cdkmonitoringconstructs/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudwatch"
)

// Experimental.
type ApiGatewayV2HttpApiMonitoring interface {
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
	Error5xxCountMetric() interface{}
	// Experimental.
	Error5xxRateMetric() interface{}
	// Experimental.
	ErrorAlarmFactory() ErrorAlarmFactory
	// Experimental.
	ErrorCountAnnotations() *[]*awscloudwatch.HorizontalAnnotation
	// Experimental.
	ErrorRateAnnotations() *[]*awscloudwatch.HorizontalAnnotation
	// Experimental.
	IntegrationLatencyMetrics() *map[string]interface{}
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

// The jsii proxy struct for ApiGatewayV2HttpApiMonitoring
type jsiiProxy_ApiGatewayV2HttpApiMonitoring struct {
	jsiiProxy_Monitoring
}

func (j *jsiiProxy_ApiGatewayV2HttpApiMonitoring) AlarmFactory() AlarmFactory {
	var returns AlarmFactory
	_jsii_.Get(
		j,
		"alarmFactory",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayV2HttpApiMonitoring) Alarms() *[]*AlarmWithAnnotation {
	var returns *[]*AlarmWithAnnotation
	_jsii_.Get(
		j,
		"alarms",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayV2HttpApiMonitoring) Error4xxCountMetric() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"error4xxCountMetric",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayV2HttpApiMonitoring) Error4xxRateMetric() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"error4xxRateMetric",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayV2HttpApiMonitoring) Error5xxCountMetric() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"error5xxCountMetric",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayV2HttpApiMonitoring) Error5xxRateMetric() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"error5xxRateMetric",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayV2HttpApiMonitoring) ErrorAlarmFactory() ErrorAlarmFactory {
	var returns ErrorAlarmFactory
	_jsii_.Get(
		j,
		"errorAlarmFactory",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayV2HttpApiMonitoring) ErrorCountAnnotations() *[]*awscloudwatch.HorizontalAnnotation {
	var returns *[]*awscloudwatch.HorizontalAnnotation
	_jsii_.Get(
		j,
		"errorCountAnnotations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayV2HttpApiMonitoring) ErrorRateAnnotations() *[]*awscloudwatch.HorizontalAnnotation {
	var returns *[]*awscloudwatch.HorizontalAnnotation
	_jsii_.Get(
		j,
		"errorRateAnnotations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayV2HttpApiMonitoring) IntegrationLatencyMetrics() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"integrationLatencyMetrics",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayV2HttpApiMonitoring) LatencyAlarmFactory() LatencyAlarmFactory {
	var returns LatencyAlarmFactory
	_jsii_.Get(
		j,
		"latencyAlarmFactory",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayV2HttpApiMonitoring) LatencyAnnotations() *[]*awscloudwatch.HorizontalAnnotation {
	var returns *[]*awscloudwatch.HorizontalAnnotation
	_jsii_.Get(
		j,
		"latencyAnnotations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayV2HttpApiMonitoring) LatencyMetrics() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"latencyMetrics",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayV2HttpApiMonitoring) LatencyTypesToRender() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"latencyTypesToRender",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayV2HttpApiMonitoring) LocalAlarmNamePrefixOverride() *string {
	var returns *string
	_jsii_.Get(
		j,
		"localAlarmNamePrefixOverride",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayV2HttpApiMonitoring) Scope() MonitoringScope {
	var returns MonitoringScope
	_jsii_.Get(
		j,
		"scope",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayV2HttpApiMonitoring) Title() *string {
	var returns *string
	_jsii_.Get(
		j,
		"title",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayV2HttpApiMonitoring) TpsAlarmFactory() TpsAlarmFactory {
	var returns TpsAlarmFactory
	_jsii_.Get(
		j,
		"tpsAlarmFactory",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayV2HttpApiMonitoring) TpsAnnotations() *[]*awscloudwatch.HorizontalAnnotation {
	var returns *[]*awscloudwatch.HorizontalAnnotation
	_jsii_.Get(
		j,
		"tpsAnnotations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiGatewayV2HttpApiMonitoring) TpsMetric() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tpsMetric",
		&returns,
	)
	return returns
}


// Experimental.
func NewApiGatewayV2HttpApiMonitoring(scope MonitoringScope, props *ApiGatewayV2HttpApiMonitoringProps) ApiGatewayV2HttpApiMonitoring {
	_init_.Initialize()

	if err := validateNewApiGatewayV2HttpApiMonitoringParameters(scope, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_ApiGatewayV2HttpApiMonitoring{}

	_jsii_.Create(
		"cdk-monitoring-constructs.ApiGatewayV2HttpApiMonitoring",
		[]interface{}{scope, props},
		&j,
	)

	return &j
}

// Experimental.
func NewApiGatewayV2HttpApiMonitoring_Override(a ApiGatewayV2HttpApiMonitoring, scope MonitoringScope, props *ApiGatewayV2HttpApiMonitoringProps) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk-monitoring-constructs.ApiGatewayV2HttpApiMonitoring",
		[]interface{}{scope, props},
		a,
	)
}

func (a *jsiiProxy_ApiGatewayV2HttpApiMonitoring) AddAlarm(alarm *AlarmWithAnnotation) {
	if err := a.validateAddAlarmParameters(alarm); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addAlarm",
		[]interface{}{alarm},
	)
}

func (a *jsiiProxy_ApiGatewayV2HttpApiMonitoring) AlarmWidgets() *[]awscloudwatch.IWidget {
	var returns *[]awscloudwatch.IWidget

	_jsii_.Invoke(
		a,
		"alarmWidgets",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiGatewayV2HttpApiMonitoring) CreateAlarmFactory(alarmNamePrefix *string) AlarmFactory {
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

func (a *jsiiProxy_ApiGatewayV2HttpApiMonitoring) CreatedAlarms() *[]*AlarmWithAnnotation {
	var returns *[]*AlarmWithAnnotation

	_jsii_.Invoke(
		a,
		"createdAlarms",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiGatewayV2HttpApiMonitoring) CreateErrorCountWidget(width *float64, height *float64) awscloudwatch.GraphWidget {
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

func (a *jsiiProxy_ApiGatewayV2HttpApiMonitoring) CreateErrorRateWidget(width *float64, height *float64) awscloudwatch.GraphWidget {
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

func (a *jsiiProxy_ApiGatewayV2HttpApiMonitoring) CreateLatencyWidget(width *float64, height *float64) awscloudwatch.GraphWidget {
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

func (a *jsiiProxy_ApiGatewayV2HttpApiMonitoring) CreateMetricFactory() MetricFactory {
	var returns MetricFactory

	_jsii_.Invoke(
		a,
		"createMetricFactory",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiGatewayV2HttpApiMonitoring) CreateTitleWidget() MonitoringHeaderWidget {
	var returns MonitoringHeaderWidget

	_jsii_.Invoke(
		a,
		"createTitleWidget",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiGatewayV2HttpApiMonitoring) CreateTpsWidget(width *float64, height *float64) awscloudwatch.GraphWidget {
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

func (a *jsiiProxy_ApiGatewayV2HttpApiMonitoring) CreateWidgetFactory() IWidgetFactory {
	var returns IWidgetFactory

	_jsii_.Invoke(
		a,
		"createWidgetFactory",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiGatewayV2HttpApiMonitoring) SummaryWidgets() *[]awscloudwatch.IWidget {
	var returns *[]awscloudwatch.IWidget

	_jsii_.Invoke(
		a,
		"summaryWidgets",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiGatewayV2HttpApiMonitoring) Widgets() *[]awscloudwatch.IWidget {
	var returns *[]awscloudwatch.IWidget

	_jsii_.Invoke(
		a,
		"widgets",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiGatewayV2HttpApiMonitoring) WidgetsForDashboard(name *string) *[]awscloudwatch.IWidget {
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


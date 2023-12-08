package cdkmonitoringconstructs

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/corymhall/cdk-constructs-go/cdkmonitoringconstructs/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudwatch"
)

// Experimental.
type LambdaFunctionMonitoring interface {
	Monitoring
	// Experimental.
	AgeAlarmFactory() AgeAlarmFactory
	// Experimental.
	AlarmFactory() AlarmFactory
	// Experimental.
	Alarms() *[]*AlarmWithAnnotation
	// Experimental.
	ConcurrentExecutionsCountMetric() interface{}
	// Experimental.
	CpuTotalTimeAnnotations() *[]*awscloudwatch.HorizontalAnnotation
	// Experimental.
	EnhancedMetricFactory() LambdaFunctionEnhancedMetricFactory
	// Experimental.
	EnhancedMetricFunctionCostMetric() interface{}
	// Experimental.
	EnhancedMonitoringAvgCpuTotalTimeMetric() interface{}
	// Experimental.
	EnhancedMonitoringAvgMemoryUtilizationMetric() interface{}
	// Experimental.
	EnhancedMonitoringMaxCpuTotalTimeMetric() interface{}
	// Experimental.
	EnhancedMonitoringMaxMemoryUtilizationMetric() interface{}
	// Experimental.
	EnhancedMonitoringP90CpuTotalTimeMetric() interface{}
	// Experimental.
	EnhancedMonitoringP90MemoryUtilizationMetric() interface{}
	// Experimental.
	ErrorAlarmFactory() ErrorAlarmFactory
	// Experimental.
	ErrorCountAnnotations() *[]*awscloudwatch.HorizontalAnnotation
	// Experimental.
	ErrorRateAnnotations() *[]*awscloudwatch.HorizontalAnnotation
	// Experimental.
	FaultCountMetric() interface{}
	// Experimental.
	FaultRateMetric() interface{}
	// Experimental.
	FunctionUrl() *string
	// Experimental.
	InvocationCountAnnotations() *[]*awscloudwatch.HorizontalAnnotation
	// Experimental.
	InvocationCountMetric() interface{}
	// Experimental.
	InvocationRateAnnotations() *[]*awscloudwatch.HorizontalAnnotation
	// Experimental.
	LambdaInsightsEnabled() *bool
	// Experimental.
	LatencyAlarmFactory() LatencyAlarmFactory
	// Experimental.
	LatencyAnnotations() *[]*awscloudwatch.HorizontalAnnotation
	// Experimental.
	LocalAlarmNamePrefixOverride() *string
	// Experimental.
	MaxIteratorAgeAnnotations() *[]*awscloudwatch.HorizontalAnnotation
	// Experimental.
	MaxIteratorAgeMetric() interface{}
	// Experimental.
	MemoryUsageAnnotations() *[]*awscloudwatch.HorizontalAnnotation
	// Experimental.
	MetricFactory() LambdaFunctionMetricFactory
	// Experimental.
	NamingStrategy() MonitoringNamingStrategy
	// Experimental.
	P50LatencyMetric() interface{}
	// Experimental.
	P90LatencyMetric() interface{}
	// Experimental.
	P99LatencyMetric() interface{}
	// Experimental.
	ProvisionedConcurrencySpilloverInvocationsCountMetric() interface{}
	// Experimental.
	ProvisionedConcurrencySpilloverInvocationsRateMetric() interface{}
	// Experimental.
	Scope() MonitoringScope
	// Experimental.
	TaskHealthAlarmFactory() TaskHealthAlarmFactory
	// Experimental.
	ThrottlesCountMetric() interface{}
	// Experimental.
	ThrottlesRateMetric() interface{}
	// Experimental.
	Title() *string
	// Experimental.
	TpsAlarmFactory() TpsAlarmFactory
	// Experimental.
	TpsAnnotations() *[]*awscloudwatch.HorizontalAnnotation
	// Experimental.
	TpsMetric() interface{}
	// Experimental.
	UsageAlarmFactory() UsageAlarmFactory
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
	CreateInvocationWidget(width *float64, height *float64) awscloudwatch.GraphWidget
	// Experimental.
	CreateIteratorAgeWidget(width *float64, height *float64) awscloudwatch.GraphWidget
	// Experimental.
	CreateLambdaInsightsCpuWidget(width *float64, height *float64) awscloudwatch.GraphWidget
	// Experimental.
	CreateLambdaInsightsFunctionCostWidget(width *float64, height *float64) awscloudwatch.GraphWidget
	// Experimental.
	CreateLambdaInsightsMemoryWidget(width *float64, height *float64) awscloudwatch.GraphWidget
	// Experimental.
	CreateLatencyWidget(width *float64, height *float64) awscloudwatch.GraphWidget
	// Creates a new metric factory.
	// Experimental.
	CreateMetricFactory() MetricFactory
	// Experimental.
	CreateRateWidget(width *float64, height *float64) awscloudwatch.GraphWidget
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

// The jsii proxy struct for LambdaFunctionMonitoring
type jsiiProxy_LambdaFunctionMonitoring struct {
	jsiiProxy_Monitoring
}

func (j *jsiiProxy_LambdaFunctionMonitoring) AgeAlarmFactory() AgeAlarmFactory {
	var returns AgeAlarmFactory
	_jsii_.Get(
		j,
		"ageAlarmFactory",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunctionMonitoring) AlarmFactory() AlarmFactory {
	var returns AlarmFactory
	_jsii_.Get(
		j,
		"alarmFactory",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunctionMonitoring) Alarms() *[]*AlarmWithAnnotation {
	var returns *[]*AlarmWithAnnotation
	_jsii_.Get(
		j,
		"alarms",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunctionMonitoring) ConcurrentExecutionsCountMetric() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"concurrentExecutionsCountMetric",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunctionMonitoring) CpuTotalTimeAnnotations() *[]*awscloudwatch.HorizontalAnnotation {
	var returns *[]*awscloudwatch.HorizontalAnnotation
	_jsii_.Get(
		j,
		"cpuTotalTimeAnnotations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunctionMonitoring) EnhancedMetricFactory() LambdaFunctionEnhancedMetricFactory {
	var returns LambdaFunctionEnhancedMetricFactory
	_jsii_.Get(
		j,
		"enhancedMetricFactory",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunctionMonitoring) EnhancedMetricFunctionCostMetric() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enhancedMetricFunctionCostMetric",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunctionMonitoring) EnhancedMonitoringAvgCpuTotalTimeMetric() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enhancedMonitoringAvgCpuTotalTimeMetric",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunctionMonitoring) EnhancedMonitoringAvgMemoryUtilizationMetric() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enhancedMonitoringAvgMemoryUtilizationMetric",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunctionMonitoring) EnhancedMonitoringMaxCpuTotalTimeMetric() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enhancedMonitoringMaxCpuTotalTimeMetric",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunctionMonitoring) EnhancedMonitoringMaxMemoryUtilizationMetric() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enhancedMonitoringMaxMemoryUtilizationMetric",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunctionMonitoring) EnhancedMonitoringP90CpuTotalTimeMetric() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enhancedMonitoringP90CpuTotalTimeMetric",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunctionMonitoring) EnhancedMonitoringP90MemoryUtilizationMetric() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enhancedMonitoringP90MemoryUtilizationMetric",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunctionMonitoring) ErrorAlarmFactory() ErrorAlarmFactory {
	var returns ErrorAlarmFactory
	_jsii_.Get(
		j,
		"errorAlarmFactory",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunctionMonitoring) ErrorCountAnnotations() *[]*awscloudwatch.HorizontalAnnotation {
	var returns *[]*awscloudwatch.HorizontalAnnotation
	_jsii_.Get(
		j,
		"errorCountAnnotations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunctionMonitoring) ErrorRateAnnotations() *[]*awscloudwatch.HorizontalAnnotation {
	var returns *[]*awscloudwatch.HorizontalAnnotation
	_jsii_.Get(
		j,
		"errorRateAnnotations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunctionMonitoring) FaultCountMetric() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"faultCountMetric",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunctionMonitoring) FaultRateMetric() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"faultRateMetric",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunctionMonitoring) FunctionUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"functionUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunctionMonitoring) InvocationCountAnnotations() *[]*awscloudwatch.HorizontalAnnotation {
	var returns *[]*awscloudwatch.HorizontalAnnotation
	_jsii_.Get(
		j,
		"invocationCountAnnotations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunctionMonitoring) InvocationCountMetric() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"invocationCountMetric",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunctionMonitoring) InvocationRateAnnotations() *[]*awscloudwatch.HorizontalAnnotation {
	var returns *[]*awscloudwatch.HorizontalAnnotation
	_jsii_.Get(
		j,
		"invocationRateAnnotations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunctionMonitoring) LambdaInsightsEnabled() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"lambdaInsightsEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunctionMonitoring) LatencyAlarmFactory() LatencyAlarmFactory {
	var returns LatencyAlarmFactory
	_jsii_.Get(
		j,
		"latencyAlarmFactory",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunctionMonitoring) LatencyAnnotations() *[]*awscloudwatch.HorizontalAnnotation {
	var returns *[]*awscloudwatch.HorizontalAnnotation
	_jsii_.Get(
		j,
		"latencyAnnotations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunctionMonitoring) LocalAlarmNamePrefixOverride() *string {
	var returns *string
	_jsii_.Get(
		j,
		"localAlarmNamePrefixOverride",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunctionMonitoring) MaxIteratorAgeAnnotations() *[]*awscloudwatch.HorizontalAnnotation {
	var returns *[]*awscloudwatch.HorizontalAnnotation
	_jsii_.Get(
		j,
		"maxIteratorAgeAnnotations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunctionMonitoring) MaxIteratorAgeMetric() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"maxIteratorAgeMetric",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunctionMonitoring) MemoryUsageAnnotations() *[]*awscloudwatch.HorizontalAnnotation {
	var returns *[]*awscloudwatch.HorizontalAnnotation
	_jsii_.Get(
		j,
		"memoryUsageAnnotations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunctionMonitoring) MetricFactory() LambdaFunctionMetricFactory {
	var returns LambdaFunctionMetricFactory
	_jsii_.Get(
		j,
		"metricFactory",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunctionMonitoring) NamingStrategy() MonitoringNamingStrategy {
	var returns MonitoringNamingStrategy
	_jsii_.Get(
		j,
		"namingStrategy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunctionMonitoring) P50LatencyMetric() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"p50LatencyMetric",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunctionMonitoring) P90LatencyMetric() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"p90LatencyMetric",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunctionMonitoring) P99LatencyMetric() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"p99LatencyMetric",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunctionMonitoring) ProvisionedConcurrencySpilloverInvocationsCountMetric() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"provisionedConcurrencySpilloverInvocationsCountMetric",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunctionMonitoring) ProvisionedConcurrencySpilloverInvocationsRateMetric() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"provisionedConcurrencySpilloverInvocationsRateMetric",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunctionMonitoring) Scope() MonitoringScope {
	var returns MonitoringScope
	_jsii_.Get(
		j,
		"scope",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunctionMonitoring) TaskHealthAlarmFactory() TaskHealthAlarmFactory {
	var returns TaskHealthAlarmFactory
	_jsii_.Get(
		j,
		"taskHealthAlarmFactory",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunctionMonitoring) ThrottlesCountMetric() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"throttlesCountMetric",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunctionMonitoring) ThrottlesRateMetric() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"throttlesRateMetric",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunctionMonitoring) Title() *string {
	var returns *string
	_jsii_.Get(
		j,
		"title",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunctionMonitoring) TpsAlarmFactory() TpsAlarmFactory {
	var returns TpsAlarmFactory
	_jsii_.Get(
		j,
		"tpsAlarmFactory",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunctionMonitoring) TpsAnnotations() *[]*awscloudwatch.HorizontalAnnotation {
	var returns *[]*awscloudwatch.HorizontalAnnotation
	_jsii_.Get(
		j,
		"tpsAnnotations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunctionMonitoring) TpsMetric() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tpsMetric",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LambdaFunctionMonitoring) UsageAlarmFactory() UsageAlarmFactory {
	var returns UsageAlarmFactory
	_jsii_.Get(
		j,
		"usageAlarmFactory",
		&returns,
	)
	return returns
}


// Experimental.
func NewLambdaFunctionMonitoring(scope MonitoringScope, props *LambdaFunctionMonitoringProps) LambdaFunctionMonitoring {
	_init_.Initialize()

	if err := validateNewLambdaFunctionMonitoringParameters(scope, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_LambdaFunctionMonitoring{}

	_jsii_.Create(
		"cdk-monitoring-constructs.LambdaFunctionMonitoring",
		[]interface{}{scope, props},
		&j,
	)

	return &j
}

// Experimental.
func NewLambdaFunctionMonitoring_Override(l LambdaFunctionMonitoring, scope MonitoringScope, props *LambdaFunctionMonitoringProps) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk-monitoring-constructs.LambdaFunctionMonitoring",
		[]interface{}{scope, props},
		l,
	)
}

func (l *jsiiProxy_LambdaFunctionMonitoring) AddAlarm(alarm *AlarmWithAnnotation) {
	if err := l.validateAddAlarmParameters(alarm); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"addAlarm",
		[]interface{}{alarm},
	)
}

func (l *jsiiProxy_LambdaFunctionMonitoring) AlarmWidgets() *[]awscloudwatch.IWidget {
	var returns *[]awscloudwatch.IWidget

	_jsii_.Invoke(
		l,
		"alarmWidgets",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LambdaFunctionMonitoring) CreateAlarmFactory(alarmNamePrefix *string) AlarmFactory {
	if err := l.validateCreateAlarmFactoryParameters(alarmNamePrefix); err != nil {
		panic(err)
	}
	var returns AlarmFactory

	_jsii_.Invoke(
		l,
		"createAlarmFactory",
		[]interface{}{alarmNamePrefix},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LambdaFunctionMonitoring) CreatedAlarms() *[]*AlarmWithAnnotation {
	var returns *[]*AlarmWithAnnotation

	_jsii_.Invoke(
		l,
		"createdAlarms",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LambdaFunctionMonitoring) CreateErrorCountWidget(width *float64, height *float64) awscloudwatch.GraphWidget {
	if err := l.validateCreateErrorCountWidgetParameters(width, height); err != nil {
		panic(err)
	}
	var returns awscloudwatch.GraphWidget

	_jsii_.Invoke(
		l,
		"createErrorCountWidget",
		[]interface{}{width, height},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LambdaFunctionMonitoring) CreateErrorRateWidget(width *float64, height *float64) awscloudwatch.GraphWidget {
	if err := l.validateCreateErrorRateWidgetParameters(width, height); err != nil {
		panic(err)
	}
	var returns awscloudwatch.GraphWidget

	_jsii_.Invoke(
		l,
		"createErrorRateWidget",
		[]interface{}{width, height},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LambdaFunctionMonitoring) CreateInvocationWidget(width *float64, height *float64) awscloudwatch.GraphWidget {
	if err := l.validateCreateInvocationWidgetParameters(width, height); err != nil {
		panic(err)
	}
	var returns awscloudwatch.GraphWidget

	_jsii_.Invoke(
		l,
		"createInvocationWidget",
		[]interface{}{width, height},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LambdaFunctionMonitoring) CreateIteratorAgeWidget(width *float64, height *float64) awscloudwatch.GraphWidget {
	if err := l.validateCreateIteratorAgeWidgetParameters(width, height); err != nil {
		panic(err)
	}
	var returns awscloudwatch.GraphWidget

	_jsii_.Invoke(
		l,
		"createIteratorAgeWidget",
		[]interface{}{width, height},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LambdaFunctionMonitoring) CreateLambdaInsightsCpuWidget(width *float64, height *float64) awscloudwatch.GraphWidget {
	if err := l.validateCreateLambdaInsightsCpuWidgetParameters(width, height); err != nil {
		panic(err)
	}
	var returns awscloudwatch.GraphWidget

	_jsii_.Invoke(
		l,
		"createLambdaInsightsCpuWidget",
		[]interface{}{width, height},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LambdaFunctionMonitoring) CreateLambdaInsightsFunctionCostWidget(width *float64, height *float64) awscloudwatch.GraphWidget {
	if err := l.validateCreateLambdaInsightsFunctionCostWidgetParameters(width, height); err != nil {
		panic(err)
	}
	var returns awscloudwatch.GraphWidget

	_jsii_.Invoke(
		l,
		"createLambdaInsightsFunctionCostWidget",
		[]interface{}{width, height},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LambdaFunctionMonitoring) CreateLambdaInsightsMemoryWidget(width *float64, height *float64) awscloudwatch.GraphWidget {
	if err := l.validateCreateLambdaInsightsMemoryWidgetParameters(width, height); err != nil {
		panic(err)
	}
	var returns awscloudwatch.GraphWidget

	_jsii_.Invoke(
		l,
		"createLambdaInsightsMemoryWidget",
		[]interface{}{width, height},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LambdaFunctionMonitoring) CreateLatencyWidget(width *float64, height *float64) awscloudwatch.GraphWidget {
	if err := l.validateCreateLatencyWidgetParameters(width, height); err != nil {
		panic(err)
	}
	var returns awscloudwatch.GraphWidget

	_jsii_.Invoke(
		l,
		"createLatencyWidget",
		[]interface{}{width, height},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LambdaFunctionMonitoring) CreateMetricFactory() MetricFactory {
	var returns MetricFactory

	_jsii_.Invoke(
		l,
		"createMetricFactory",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LambdaFunctionMonitoring) CreateRateWidget(width *float64, height *float64) awscloudwatch.GraphWidget {
	if err := l.validateCreateRateWidgetParameters(width, height); err != nil {
		panic(err)
	}
	var returns awscloudwatch.GraphWidget

	_jsii_.Invoke(
		l,
		"createRateWidget",
		[]interface{}{width, height},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LambdaFunctionMonitoring) CreateTitleWidget() MonitoringHeaderWidget {
	var returns MonitoringHeaderWidget

	_jsii_.Invoke(
		l,
		"createTitleWidget",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LambdaFunctionMonitoring) CreateTpsWidget(width *float64, height *float64) awscloudwatch.GraphWidget {
	if err := l.validateCreateTpsWidgetParameters(width, height); err != nil {
		panic(err)
	}
	var returns awscloudwatch.GraphWidget

	_jsii_.Invoke(
		l,
		"createTpsWidget",
		[]interface{}{width, height},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LambdaFunctionMonitoring) CreateWidgetFactory() IWidgetFactory {
	var returns IWidgetFactory

	_jsii_.Invoke(
		l,
		"createWidgetFactory",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LambdaFunctionMonitoring) SummaryWidgets() *[]awscloudwatch.IWidget {
	var returns *[]awscloudwatch.IWidget

	_jsii_.Invoke(
		l,
		"summaryWidgets",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LambdaFunctionMonitoring) Widgets() *[]awscloudwatch.IWidget {
	var returns *[]awscloudwatch.IWidget

	_jsii_.Invoke(
		l,
		"widgets",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LambdaFunctionMonitoring) WidgetsForDashboard(name *string) *[]awscloudwatch.IWidget {
	if err := l.validateWidgetsForDashboardParameters(name); err != nil {
		panic(err)
	}
	var returns *[]awscloudwatch.IWidget

	_jsii_.Invoke(
		l,
		"widgetsForDashboard",
		[]interface{}{name},
		&returns,
	)

	return returns
}


package cdkmonitoringconstructs

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/corymhall/cdk-constructs-go/cdkmonitoringconstructs/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudwatch"
)

// Experimental.
type FargateServiceMonitoring interface {
	Monitoring
	// Experimental.
	ActiveTcpFlowCountMetric() interface{}
	// Experimental.
	Alarms() *[]*AlarmWithAnnotation
	// Experimental.
	BaseServiceMetricFactory() BaseServiceMetricFactory
	// Experimental.
	CpuUsageAnnotations() *[]*awscloudwatch.HorizontalAnnotation
	// Experimental.
	CpuUtilisationMetric() interface{}
	// Experimental.
	HealthyTaskCountMetric() interface{}
	// Experimental.
	HealthyTaskPercentMetric() interface{}
	// Experimental.
	LoadBalancerMetricFactory() ILoadBalancerMetricFactory
	// Experimental.
	LocalAlarmNamePrefixOverride() *string
	// Experimental.
	MemoryUsageAnnotations() *[]*awscloudwatch.HorizontalAnnotation
	// Experimental.
	MemoryUtilisationMetric() interface{}
	// Experimental.
	MetricFactory() MetricFactory
	// Experimental.
	NewTcpFlowCountMetric() interface{}
	// Experimental.
	ProcessedBytesAnnotations() *[]*awscloudwatch.HorizontalAnnotation
	// Experimental.
	ProcessedBytesMetric() interface{}
	// Experimental.
	RunningTaskCountMetric() interface{}
	// Experimental.
	Scope() MonitoringScope
	// Experimental.
	TaskHealthAlarmFactory() TaskHealthAlarmFactory
	// Experimental.
	TaskHealthAnnotations() *[]*awscloudwatch.HorizontalAnnotation
	// Experimental.
	ThroughputAlarmFactory() ThroughputAlarmFactory
	// Experimental.
	Title() *string
	// Experimental.
	UnhealthyTaskCountMetric() interface{}
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
	// Experimental.
	CreateCpuWidget(width *float64, height *float64) awscloudwatch.GraphWidget
	// Returns all the alarms created.
	// Experimental.
	CreatedAlarms() *[]*AlarmWithAnnotation
	// Experimental.
	CreateMemoryWidget(width *float64, height *float64) awscloudwatch.GraphWidget
	// Creates a new metric factory.
	// Experimental.
	CreateMetricFactory() MetricFactory
	// Experimental.
	CreateTaskHealthWidget(width *float64, height *float64) awscloudwatch.GraphWidget
	// Experimental.
	CreateTitleWidget() MonitoringHeaderWidget
	// Experimental.
	CreateTpcFlowsWidget(width *float64, height *float64) awscloudwatch.GraphWidget
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

// The jsii proxy struct for FargateServiceMonitoring
type jsiiProxy_FargateServiceMonitoring struct {
	jsiiProxy_Monitoring
}

func (j *jsiiProxy_FargateServiceMonitoring) ActiveTcpFlowCountMetric() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"activeTcpFlowCountMetric",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FargateServiceMonitoring) Alarms() *[]*AlarmWithAnnotation {
	var returns *[]*AlarmWithAnnotation
	_jsii_.Get(
		j,
		"alarms",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FargateServiceMonitoring) BaseServiceMetricFactory() BaseServiceMetricFactory {
	var returns BaseServiceMetricFactory
	_jsii_.Get(
		j,
		"baseServiceMetricFactory",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FargateServiceMonitoring) CpuUsageAnnotations() *[]*awscloudwatch.HorizontalAnnotation {
	var returns *[]*awscloudwatch.HorizontalAnnotation
	_jsii_.Get(
		j,
		"cpuUsageAnnotations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FargateServiceMonitoring) CpuUtilisationMetric() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"cpuUtilisationMetric",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FargateServiceMonitoring) HealthyTaskCountMetric() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"healthyTaskCountMetric",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FargateServiceMonitoring) HealthyTaskPercentMetric() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"healthyTaskPercentMetric",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FargateServiceMonitoring) LoadBalancerMetricFactory() ILoadBalancerMetricFactory {
	var returns ILoadBalancerMetricFactory
	_jsii_.Get(
		j,
		"loadBalancerMetricFactory",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FargateServiceMonitoring) LocalAlarmNamePrefixOverride() *string {
	var returns *string
	_jsii_.Get(
		j,
		"localAlarmNamePrefixOverride",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FargateServiceMonitoring) MemoryUsageAnnotations() *[]*awscloudwatch.HorizontalAnnotation {
	var returns *[]*awscloudwatch.HorizontalAnnotation
	_jsii_.Get(
		j,
		"memoryUsageAnnotations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FargateServiceMonitoring) MemoryUtilisationMetric() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"memoryUtilisationMetric",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FargateServiceMonitoring) MetricFactory() MetricFactory {
	var returns MetricFactory
	_jsii_.Get(
		j,
		"metricFactory",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FargateServiceMonitoring) NewTcpFlowCountMetric() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"newTcpFlowCountMetric",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FargateServiceMonitoring) ProcessedBytesAnnotations() *[]*awscloudwatch.HorizontalAnnotation {
	var returns *[]*awscloudwatch.HorizontalAnnotation
	_jsii_.Get(
		j,
		"processedBytesAnnotations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FargateServiceMonitoring) ProcessedBytesMetric() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"processedBytesMetric",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FargateServiceMonitoring) RunningTaskCountMetric() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"runningTaskCountMetric",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FargateServiceMonitoring) Scope() MonitoringScope {
	var returns MonitoringScope
	_jsii_.Get(
		j,
		"scope",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FargateServiceMonitoring) TaskHealthAlarmFactory() TaskHealthAlarmFactory {
	var returns TaskHealthAlarmFactory
	_jsii_.Get(
		j,
		"taskHealthAlarmFactory",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FargateServiceMonitoring) TaskHealthAnnotations() *[]*awscloudwatch.HorizontalAnnotation {
	var returns *[]*awscloudwatch.HorizontalAnnotation
	_jsii_.Get(
		j,
		"taskHealthAnnotations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FargateServiceMonitoring) ThroughputAlarmFactory() ThroughputAlarmFactory {
	var returns ThroughputAlarmFactory
	_jsii_.Get(
		j,
		"throughputAlarmFactory",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FargateServiceMonitoring) Title() *string {
	var returns *string
	_jsii_.Get(
		j,
		"title",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FargateServiceMonitoring) UnhealthyTaskCountMetric() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"unhealthyTaskCountMetric",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FargateServiceMonitoring) UsageAlarmFactory() UsageAlarmFactory {
	var returns UsageAlarmFactory
	_jsii_.Get(
		j,
		"usageAlarmFactory",
		&returns,
	)
	return returns
}


// Experimental.
func NewFargateServiceMonitoring(scope MonitoringScope, props *CustomFargateServiceMonitoringProps) FargateServiceMonitoring {
	_init_.Initialize()

	if err := validateNewFargateServiceMonitoringParameters(scope, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_FargateServiceMonitoring{}

	_jsii_.Create(
		"cdk-monitoring-constructs.FargateServiceMonitoring",
		[]interface{}{scope, props},
		&j,
	)

	return &j
}

// Experimental.
func NewFargateServiceMonitoring_Override(f FargateServiceMonitoring, scope MonitoringScope, props *CustomFargateServiceMonitoringProps) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk-monitoring-constructs.FargateServiceMonitoring",
		[]interface{}{scope, props},
		f,
	)
}

func (f *jsiiProxy_FargateServiceMonitoring) AddAlarm(alarm *AlarmWithAnnotation) {
	if err := f.validateAddAlarmParameters(alarm); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		f,
		"addAlarm",
		[]interface{}{alarm},
	)
}

func (f *jsiiProxy_FargateServiceMonitoring) AlarmWidgets() *[]awscloudwatch.IWidget {
	var returns *[]awscloudwatch.IWidget

	_jsii_.Invoke(
		f,
		"alarmWidgets",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FargateServiceMonitoring) CreateAlarmFactory(alarmNamePrefix *string) AlarmFactory {
	if err := f.validateCreateAlarmFactoryParameters(alarmNamePrefix); err != nil {
		panic(err)
	}
	var returns AlarmFactory

	_jsii_.Invoke(
		f,
		"createAlarmFactory",
		[]interface{}{alarmNamePrefix},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FargateServiceMonitoring) CreateCpuWidget(width *float64, height *float64) awscloudwatch.GraphWidget {
	if err := f.validateCreateCpuWidgetParameters(width, height); err != nil {
		panic(err)
	}
	var returns awscloudwatch.GraphWidget

	_jsii_.Invoke(
		f,
		"createCpuWidget",
		[]interface{}{width, height},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FargateServiceMonitoring) CreatedAlarms() *[]*AlarmWithAnnotation {
	var returns *[]*AlarmWithAnnotation

	_jsii_.Invoke(
		f,
		"createdAlarms",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FargateServiceMonitoring) CreateMemoryWidget(width *float64, height *float64) awscloudwatch.GraphWidget {
	if err := f.validateCreateMemoryWidgetParameters(width, height); err != nil {
		panic(err)
	}
	var returns awscloudwatch.GraphWidget

	_jsii_.Invoke(
		f,
		"createMemoryWidget",
		[]interface{}{width, height},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FargateServiceMonitoring) CreateMetricFactory() MetricFactory {
	var returns MetricFactory

	_jsii_.Invoke(
		f,
		"createMetricFactory",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FargateServiceMonitoring) CreateTaskHealthWidget(width *float64, height *float64) awscloudwatch.GraphWidget {
	if err := f.validateCreateTaskHealthWidgetParameters(width, height); err != nil {
		panic(err)
	}
	var returns awscloudwatch.GraphWidget

	_jsii_.Invoke(
		f,
		"createTaskHealthWidget",
		[]interface{}{width, height},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FargateServiceMonitoring) CreateTitleWidget() MonitoringHeaderWidget {
	var returns MonitoringHeaderWidget

	_jsii_.Invoke(
		f,
		"createTitleWidget",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FargateServiceMonitoring) CreateTpcFlowsWidget(width *float64, height *float64) awscloudwatch.GraphWidget {
	if err := f.validateCreateTpcFlowsWidgetParameters(width, height); err != nil {
		panic(err)
	}
	var returns awscloudwatch.GraphWidget

	_jsii_.Invoke(
		f,
		"createTpcFlowsWidget",
		[]interface{}{width, height},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FargateServiceMonitoring) CreateWidgetFactory() IWidgetFactory {
	var returns IWidgetFactory

	_jsii_.Invoke(
		f,
		"createWidgetFactory",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FargateServiceMonitoring) SummaryWidgets() *[]awscloudwatch.IWidget {
	var returns *[]awscloudwatch.IWidget

	_jsii_.Invoke(
		f,
		"summaryWidgets",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FargateServiceMonitoring) Widgets() *[]awscloudwatch.IWidget {
	var returns *[]awscloudwatch.IWidget

	_jsii_.Invoke(
		f,
		"widgets",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FargateServiceMonitoring) WidgetsForDashboard(name *string) *[]awscloudwatch.IWidget {
	if err := f.validateWidgetsForDashboardParameters(name); err != nil {
		panic(err)
	}
	var returns *[]awscloudwatch.IWidget

	_jsii_.Invoke(
		f,
		"widgetsForDashboard",
		[]interface{}{name},
		&returns,
	)

	return returns
}


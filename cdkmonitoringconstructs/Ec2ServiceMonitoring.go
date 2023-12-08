package cdkmonitoringconstructs

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/corymhall/cdk-constructs-go/cdkmonitoringconstructs/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudwatch"
)

// Experimental.
type Ec2ServiceMonitoring interface {
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

// The jsii proxy struct for Ec2ServiceMonitoring
type jsiiProxy_Ec2ServiceMonitoring struct {
	jsiiProxy_Monitoring
}

func (j *jsiiProxy_Ec2ServiceMonitoring) ActiveTcpFlowCountMetric() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"activeTcpFlowCountMetric",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2ServiceMonitoring) Alarms() *[]*AlarmWithAnnotation {
	var returns *[]*AlarmWithAnnotation
	_jsii_.Get(
		j,
		"alarms",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2ServiceMonitoring) BaseServiceMetricFactory() BaseServiceMetricFactory {
	var returns BaseServiceMetricFactory
	_jsii_.Get(
		j,
		"baseServiceMetricFactory",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2ServiceMonitoring) CpuUsageAnnotations() *[]*awscloudwatch.HorizontalAnnotation {
	var returns *[]*awscloudwatch.HorizontalAnnotation
	_jsii_.Get(
		j,
		"cpuUsageAnnotations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2ServiceMonitoring) CpuUtilisationMetric() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"cpuUtilisationMetric",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2ServiceMonitoring) HealthyTaskCountMetric() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"healthyTaskCountMetric",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2ServiceMonitoring) HealthyTaskPercentMetric() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"healthyTaskPercentMetric",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2ServiceMonitoring) LoadBalancerMetricFactory() ILoadBalancerMetricFactory {
	var returns ILoadBalancerMetricFactory
	_jsii_.Get(
		j,
		"loadBalancerMetricFactory",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2ServiceMonitoring) LocalAlarmNamePrefixOverride() *string {
	var returns *string
	_jsii_.Get(
		j,
		"localAlarmNamePrefixOverride",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2ServiceMonitoring) MemoryUsageAnnotations() *[]*awscloudwatch.HorizontalAnnotation {
	var returns *[]*awscloudwatch.HorizontalAnnotation
	_jsii_.Get(
		j,
		"memoryUsageAnnotations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2ServiceMonitoring) MemoryUtilisationMetric() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"memoryUtilisationMetric",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2ServiceMonitoring) MetricFactory() MetricFactory {
	var returns MetricFactory
	_jsii_.Get(
		j,
		"metricFactory",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2ServiceMonitoring) NewTcpFlowCountMetric() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"newTcpFlowCountMetric",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2ServiceMonitoring) ProcessedBytesAnnotations() *[]*awscloudwatch.HorizontalAnnotation {
	var returns *[]*awscloudwatch.HorizontalAnnotation
	_jsii_.Get(
		j,
		"processedBytesAnnotations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2ServiceMonitoring) ProcessedBytesMetric() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"processedBytesMetric",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2ServiceMonitoring) RunningTaskCountMetric() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"runningTaskCountMetric",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2ServiceMonitoring) Scope() MonitoringScope {
	var returns MonitoringScope
	_jsii_.Get(
		j,
		"scope",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2ServiceMonitoring) TaskHealthAlarmFactory() TaskHealthAlarmFactory {
	var returns TaskHealthAlarmFactory
	_jsii_.Get(
		j,
		"taskHealthAlarmFactory",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2ServiceMonitoring) TaskHealthAnnotations() *[]*awscloudwatch.HorizontalAnnotation {
	var returns *[]*awscloudwatch.HorizontalAnnotation
	_jsii_.Get(
		j,
		"taskHealthAnnotations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2ServiceMonitoring) ThroughputAlarmFactory() ThroughputAlarmFactory {
	var returns ThroughputAlarmFactory
	_jsii_.Get(
		j,
		"throughputAlarmFactory",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2ServiceMonitoring) Title() *string {
	var returns *string
	_jsii_.Get(
		j,
		"title",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2ServiceMonitoring) UnhealthyTaskCountMetric() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"unhealthyTaskCountMetric",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2ServiceMonitoring) UsageAlarmFactory() UsageAlarmFactory {
	var returns UsageAlarmFactory
	_jsii_.Get(
		j,
		"usageAlarmFactory",
		&returns,
	)
	return returns
}


// Experimental.
func NewEc2ServiceMonitoring(scope MonitoringScope, props *CustomEc2ServiceMonitoringProps) Ec2ServiceMonitoring {
	_init_.Initialize()

	if err := validateNewEc2ServiceMonitoringParameters(scope, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_Ec2ServiceMonitoring{}

	_jsii_.Create(
		"cdk-monitoring-constructs.Ec2ServiceMonitoring",
		[]interface{}{scope, props},
		&j,
	)

	return &j
}

// Experimental.
func NewEc2ServiceMonitoring_Override(e Ec2ServiceMonitoring, scope MonitoringScope, props *CustomEc2ServiceMonitoringProps) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk-monitoring-constructs.Ec2ServiceMonitoring",
		[]interface{}{scope, props},
		e,
	)
}

func (e *jsiiProxy_Ec2ServiceMonitoring) AddAlarm(alarm *AlarmWithAnnotation) {
	if err := e.validateAddAlarmParameters(alarm); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"addAlarm",
		[]interface{}{alarm},
	)
}

func (e *jsiiProxy_Ec2ServiceMonitoring) AlarmWidgets() *[]awscloudwatch.IWidget {
	var returns *[]awscloudwatch.IWidget

	_jsii_.Invoke(
		e,
		"alarmWidgets",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Ec2ServiceMonitoring) CreateAlarmFactory(alarmNamePrefix *string) AlarmFactory {
	if err := e.validateCreateAlarmFactoryParameters(alarmNamePrefix); err != nil {
		panic(err)
	}
	var returns AlarmFactory

	_jsii_.Invoke(
		e,
		"createAlarmFactory",
		[]interface{}{alarmNamePrefix},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Ec2ServiceMonitoring) CreateCpuWidget(width *float64, height *float64) awscloudwatch.GraphWidget {
	if err := e.validateCreateCpuWidgetParameters(width, height); err != nil {
		panic(err)
	}
	var returns awscloudwatch.GraphWidget

	_jsii_.Invoke(
		e,
		"createCpuWidget",
		[]interface{}{width, height},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Ec2ServiceMonitoring) CreatedAlarms() *[]*AlarmWithAnnotation {
	var returns *[]*AlarmWithAnnotation

	_jsii_.Invoke(
		e,
		"createdAlarms",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Ec2ServiceMonitoring) CreateMemoryWidget(width *float64, height *float64) awscloudwatch.GraphWidget {
	if err := e.validateCreateMemoryWidgetParameters(width, height); err != nil {
		panic(err)
	}
	var returns awscloudwatch.GraphWidget

	_jsii_.Invoke(
		e,
		"createMemoryWidget",
		[]interface{}{width, height},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Ec2ServiceMonitoring) CreateMetricFactory() MetricFactory {
	var returns MetricFactory

	_jsii_.Invoke(
		e,
		"createMetricFactory",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Ec2ServiceMonitoring) CreateTaskHealthWidget(width *float64, height *float64) awscloudwatch.GraphWidget {
	if err := e.validateCreateTaskHealthWidgetParameters(width, height); err != nil {
		panic(err)
	}
	var returns awscloudwatch.GraphWidget

	_jsii_.Invoke(
		e,
		"createTaskHealthWidget",
		[]interface{}{width, height},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Ec2ServiceMonitoring) CreateTitleWidget() MonitoringHeaderWidget {
	var returns MonitoringHeaderWidget

	_jsii_.Invoke(
		e,
		"createTitleWidget",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Ec2ServiceMonitoring) CreateTpcFlowsWidget(width *float64, height *float64) awscloudwatch.GraphWidget {
	if err := e.validateCreateTpcFlowsWidgetParameters(width, height); err != nil {
		panic(err)
	}
	var returns awscloudwatch.GraphWidget

	_jsii_.Invoke(
		e,
		"createTpcFlowsWidget",
		[]interface{}{width, height},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Ec2ServiceMonitoring) CreateWidgetFactory() IWidgetFactory {
	var returns IWidgetFactory

	_jsii_.Invoke(
		e,
		"createWidgetFactory",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Ec2ServiceMonitoring) SummaryWidgets() *[]awscloudwatch.IWidget {
	var returns *[]awscloudwatch.IWidget

	_jsii_.Invoke(
		e,
		"summaryWidgets",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Ec2ServiceMonitoring) Widgets() *[]awscloudwatch.IWidget {
	var returns *[]awscloudwatch.IWidget

	_jsii_.Invoke(
		e,
		"widgets",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Ec2ServiceMonitoring) WidgetsForDashboard(name *string) *[]awscloudwatch.IWidget {
	if err := e.validateWidgetsForDashboardParameters(name); err != nil {
		panic(err)
	}
	var returns *[]awscloudwatch.IWidget

	_jsii_.Invoke(
		e,
		"widgetsForDashboard",
		[]interface{}{name},
		&returns,
	)

	return returns
}


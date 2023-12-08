package cdkmonitoringconstructs

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/corymhall/cdk-constructs-go/cdkmonitoringconstructs/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudwatch"
)

// Experimental.
type NetworkLoadBalancerMonitoring interface {
	Monitoring
	// Experimental.
	ActiveTcpFlowCountMetric() interface{}
	// Experimental.
	Alarms() *[]*AlarmWithAnnotation
	// Experimental.
	HealthyTaskCountMetric() interface{}
	// Experimental.
	HealthyTaskPercentMetric() interface{}
	// Experimental.
	HumanReadableName() *string
	// Experimental.
	LocalAlarmNamePrefixOverride() *string
	// Experimental.
	MetricFactory() NetworkLoadBalancerMetricFactory
	// Experimental.
	NewTcpFlowCountMetric() interface{}
	// Experimental.
	ProcessedBytesAnnotations() *[]*awscloudwatch.HorizontalAnnotation
	// Experimental.
	ProcessedBytesMetric() interface{}
	// Experimental.
	Scope() MonitoringScope
	// Experimental.
	TaskHealthAlarmFactory() TaskHealthAlarmFactory
	// Experimental.
	TaskHealthAnnotations() *[]*awscloudwatch.HorizontalAnnotation
	// Experimental.
	ThroughputAlarmFactory() ThroughputAlarmFactory
	// Experimental.
	UnhealthyTaskCountMetric() interface{}
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
	CreateTaskHealthWidget(width *float64, height *float64) awscloudwatch.GraphWidget
	// Experimental.
	CreateTcpFlowsWidget(width *float64, height *float64) awscloudwatch.GraphWidget
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

// The jsii proxy struct for NetworkLoadBalancerMonitoring
type jsiiProxy_NetworkLoadBalancerMonitoring struct {
	jsiiProxy_Monitoring
}

func (j *jsiiProxy_NetworkLoadBalancerMonitoring) ActiveTcpFlowCountMetric() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"activeTcpFlowCountMetric",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkLoadBalancerMonitoring) Alarms() *[]*AlarmWithAnnotation {
	var returns *[]*AlarmWithAnnotation
	_jsii_.Get(
		j,
		"alarms",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkLoadBalancerMonitoring) HealthyTaskCountMetric() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"healthyTaskCountMetric",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkLoadBalancerMonitoring) HealthyTaskPercentMetric() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"healthyTaskPercentMetric",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkLoadBalancerMonitoring) HumanReadableName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"humanReadableName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkLoadBalancerMonitoring) LocalAlarmNamePrefixOverride() *string {
	var returns *string
	_jsii_.Get(
		j,
		"localAlarmNamePrefixOverride",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkLoadBalancerMonitoring) MetricFactory() NetworkLoadBalancerMetricFactory {
	var returns NetworkLoadBalancerMetricFactory
	_jsii_.Get(
		j,
		"metricFactory",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkLoadBalancerMonitoring) NewTcpFlowCountMetric() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"newTcpFlowCountMetric",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkLoadBalancerMonitoring) ProcessedBytesAnnotations() *[]*awscloudwatch.HorizontalAnnotation {
	var returns *[]*awscloudwatch.HorizontalAnnotation
	_jsii_.Get(
		j,
		"processedBytesAnnotations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkLoadBalancerMonitoring) ProcessedBytesMetric() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"processedBytesMetric",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkLoadBalancerMonitoring) Scope() MonitoringScope {
	var returns MonitoringScope
	_jsii_.Get(
		j,
		"scope",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkLoadBalancerMonitoring) TaskHealthAlarmFactory() TaskHealthAlarmFactory {
	var returns TaskHealthAlarmFactory
	_jsii_.Get(
		j,
		"taskHealthAlarmFactory",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkLoadBalancerMonitoring) TaskHealthAnnotations() *[]*awscloudwatch.HorizontalAnnotation {
	var returns *[]*awscloudwatch.HorizontalAnnotation
	_jsii_.Get(
		j,
		"taskHealthAnnotations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkLoadBalancerMonitoring) ThroughputAlarmFactory() ThroughputAlarmFactory {
	var returns ThroughputAlarmFactory
	_jsii_.Get(
		j,
		"throughputAlarmFactory",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkLoadBalancerMonitoring) UnhealthyTaskCountMetric() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"unhealthyTaskCountMetric",
		&returns,
	)
	return returns
}


// Experimental.
func NewNetworkLoadBalancerMonitoring(scope MonitoringScope, props *NetworkLoadBalancerMonitoringProps) NetworkLoadBalancerMonitoring {
	_init_.Initialize()

	if err := validateNewNetworkLoadBalancerMonitoringParameters(scope, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_NetworkLoadBalancerMonitoring{}

	_jsii_.Create(
		"cdk-monitoring-constructs.NetworkLoadBalancerMonitoring",
		[]interface{}{scope, props},
		&j,
	)

	return &j
}

// Experimental.
func NewNetworkLoadBalancerMonitoring_Override(n NetworkLoadBalancerMonitoring, scope MonitoringScope, props *NetworkLoadBalancerMonitoringProps) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk-monitoring-constructs.NetworkLoadBalancerMonitoring",
		[]interface{}{scope, props},
		n,
	)
}

func (n *jsiiProxy_NetworkLoadBalancerMonitoring) AddAlarm(alarm *AlarmWithAnnotation) {
	if err := n.validateAddAlarmParameters(alarm); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"addAlarm",
		[]interface{}{alarm},
	)
}

func (n *jsiiProxy_NetworkLoadBalancerMonitoring) AlarmWidgets() *[]awscloudwatch.IWidget {
	var returns *[]awscloudwatch.IWidget

	_jsii_.Invoke(
		n,
		"alarmWidgets",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkLoadBalancerMonitoring) CreateAlarmFactory(alarmNamePrefix *string) AlarmFactory {
	if err := n.validateCreateAlarmFactoryParameters(alarmNamePrefix); err != nil {
		panic(err)
	}
	var returns AlarmFactory

	_jsii_.Invoke(
		n,
		"createAlarmFactory",
		[]interface{}{alarmNamePrefix},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkLoadBalancerMonitoring) CreatedAlarms() *[]*AlarmWithAnnotation {
	var returns *[]*AlarmWithAnnotation

	_jsii_.Invoke(
		n,
		"createdAlarms",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkLoadBalancerMonitoring) CreateMetricFactory() MetricFactory {
	var returns MetricFactory

	_jsii_.Invoke(
		n,
		"createMetricFactory",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkLoadBalancerMonitoring) CreateTaskHealthWidget(width *float64, height *float64) awscloudwatch.GraphWidget {
	if err := n.validateCreateTaskHealthWidgetParameters(width, height); err != nil {
		panic(err)
	}
	var returns awscloudwatch.GraphWidget

	_jsii_.Invoke(
		n,
		"createTaskHealthWidget",
		[]interface{}{width, height},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkLoadBalancerMonitoring) CreateTcpFlowsWidget(width *float64, height *float64) awscloudwatch.GraphWidget {
	if err := n.validateCreateTcpFlowsWidgetParameters(width, height); err != nil {
		panic(err)
	}
	var returns awscloudwatch.GraphWidget

	_jsii_.Invoke(
		n,
		"createTcpFlowsWidget",
		[]interface{}{width, height},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkLoadBalancerMonitoring) CreateTitleWidget() MonitoringHeaderWidget {
	var returns MonitoringHeaderWidget

	_jsii_.Invoke(
		n,
		"createTitleWidget",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkLoadBalancerMonitoring) CreateWidgetFactory() IWidgetFactory {
	var returns IWidgetFactory

	_jsii_.Invoke(
		n,
		"createWidgetFactory",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkLoadBalancerMonitoring) SummaryWidgets() *[]awscloudwatch.IWidget {
	var returns *[]awscloudwatch.IWidget

	_jsii_.Invoke(
		n,
		"summaryWidgets",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkLoadBalancerMonitoring) Widgets() *[]awscloudwatch.IWidget {
	var returns *[]awscloudwatch.IWidget

	_jsii_.Invoke(
		n,
		"widgets",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkLoadBalancerMonitoring) WidgetsForDashboard(name *string) *[]awscloudwatch.IWidget {
	if err := n.validateWidgetsForDashboardParameters(name); err != nil {
		panic(err)
	}
	var returns *[]awscloudwatch.IWidget

	_jsii_.Invoke(
		n,
		"widgetsForDashboard",
		[]interface{}{name},
		&returns,
	)

	return returns
}


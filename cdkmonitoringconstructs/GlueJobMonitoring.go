package cdkmonitoringconstructs

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/corymhall/cdk-constructs-go/cdkmonitoringconstructs/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudwatch"
)

// Experimental.
type GlueJobMonitoring interface {
	Monitoring
	// Experimental.
	ActiveExecutorsMetric() interface{}
	// Experimental.
	AlarmFactory() AlarmFactory
	// Experimental.
	Alarms() *[]*AlarmWithAnnotation
	// Experimental.
	BytesReadFromS3Metric() interface{}
	// Experimental.
	BytesWrittenToS3Metric() interface{}
	// Experimental.
	CompletedStagesMetric() interface{}
	// Experimental.
	CpuUsageMetric() interface{}
	// Experimental.
	ErrorAlarmFactory() ErrorAlarmFactory
	// Experimental.
	ErrorCountAnnotations() *[]*awscloudwatch.HorizontalAnnotation
	// Experimental.
	ErrorRateAnnotations() *[]*awscloudwatch.HorizontalAnnotation
	// Experimental.
	FailedTaskCountMetric() interface{}
	// Experimental.
	FailedTaskRateMetric() interface{}
	// Experimental.
	HeapMemoryUsageMetric() interface{}
	// Experimental.
	KilledTaskCountMetric() interface{}
	// Experimental.
	KilledTaskRateMetric() interface{}
	// Experimental.
	LocalAlarmNamePrefixOverride() *string
	// Experimental.
	NeededExecutorsMetric() interface{}
	// Experimental.
	Scope() MonitoringScope
	// Experimental.
	Title() *string
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
	CreateDataMovementWidget(width *float64, height *float64) awscloudwatch.GraphWidget
	// Experimental.
	CreateErrorCountWidget(width *float64, height *float64) awscloudwatch.GraphWidget
	// Experimental.
	CreateErrorRateWidget(width *float64, height *float64) awscloudwatch.GraphWidget
	// Experimental.
	CreateJobExecutionWidget(width *float64, height *float64) awscloudwatch.GraphWidget
	// Creates a new metric factory.
	// Experimental.
	CreateMetricFactory() MetricFactory
	// Experimental.
	CreateTitleWidget() MonitoringHeaderWidget
	// Experimental.
	CreateUtilizationWidget(width *float64, height *float64) awscloudwatch.GraphWidget
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

// The jsii proxy struct for GlueJobMonitoring
type jsiiProxy_GlueJobMonitoring struct {
	jsiiProxy_Monitoring
}

func (j *jsiiProxy_GlueJobMonitoring) ActiveExecutorsMetric() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"activeExecutorsMetric",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueJobMonitoring) AlarmFactory() AlarmFactory {
	var returns AlarmFactory
	_jsii_.Get(
		j,
		"alarmFactory",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueJobMonitoring) Alarms() *[]*AlarmWithAnnotation {
	var returns *[]*AlarmWithAnnotation
	_jsii_.Get(
		j,
		"alarms",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueJobMonitoring) BytesReadFromS3Metric() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"bytesReadFromS3Metric",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueJobMonitoring) BytesWrittenToS3Metric() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"bytesWrittenToS3Metric",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueJobMonitoring) CompletedStagesMetric() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"completedStagesMetric",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueJobMonitoring) CpuUsageMetric() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"cpuUsageMetric",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueJobMonitoring) ErrorAlarmFactory() ErrorAlarmFactory {
	var returns ErrorAlarmFactory
	_jsii_.Get(
		j,
		"errorAlarmFactory",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueJobMonitoring) ErrorCountAnnotations() *[]*awscloudwatch.HorizontalAnnotation {
	var returns *[]*awscloudwatch.HorizontalAnnotation
	_jsii_.Get(
		j,
		"errorCountAnnotations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueJobMonitoring) ErrorRateAnnotations() *[]*awscloudwatch.HorizontalAnnotation {
	var returns *[]*awscloudwatch.HorizontalAnnotation
	_jsii_.Get(
		j,
		"errorRateAnnotations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueJobMonitoring) FailedTaskCountMetric() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"failedTaskCountMetric",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueJobMonitoring) FailedTaskRateMetric() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"failedTaskRateMetric",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueJobMonitoring) HeapMemoryUsageMetric() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"heapMemoryUsageMetric",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueJobMonitoring) KilledTaskCountMetric() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"killedTaskCountMetric",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueJobMonitoring) KilledTaskRateMetric() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"killedTaskRateMetric",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueJobMonitoring) LocalAlarmNamePrefixOverride() *string {
	var returns *string
	_jsii_.Get(
		j,
		"localAlarmNamePrefixOverride",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueJobMonitoring) NeededExecutorsMetric() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"neededExecutorsMetric",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueJobMonitoring) Scope() MonitoringScope {
	var returns MonitoringScope
	_jsii_.Get(
		j,
		"scope",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GlueJobMonitoring) Title() *string {
	var returns *string
	_jsii_.Get(
		j,
		"title",
		&returns,
	)
	return returns
}


// Experimental.
func NewGlueJobMonitoring(scope MonitoringScope, props *GlueJobMonitoringProps) GlueJobMonitoring {
	_init_.Initialize()

	if err := validateNewGlueJobMonitoringParameters(scope, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_GlueJobMonitoring{}

	_jsii_.Create(
		"cdk-monitoring-constructs.GlueJobMonitoring",
		[]interface{}{scope, props},
		&j,
	)

	return &j
}

// Experimental.
func NewGlueJobMonitoring_Override(g GlueJobMonitoring, scope MonitoringScope, props *GlueJobMonitoringProps) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk-monitoring-constructs.GlueJobMonitoring",
		[]interface{}{scope, props},
		g,
	)
}

func (g *jsiiProxy_GlueJobMonitoring) AddAlarm(alarm *AlarmWithAnnotation) {
	if err := g.validateAddAlarmParameters(alarm); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"addAlarm",
		[]interface{}{alarm},
	)
}

func (g *jsiiProxy_GlueJobMonitoring) AlarmWidgets() *[]awscloudwatch.IWidget {
	var returns *[]awscloudwatch.IWidget

	_jsii_.Invoke(
		g,
		"alarmWidgets",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GlueJobMonitoring) CreateAlarmFactory(alarmNamePrefix *string) AlarmFactory {
	if err := g.validateCreateAlarmFactoryParameters(alarmNamePrefix); err != nil {
		panic(err)
	}
	var returns AlarmFactory

	_jsii_.Invoke(
		g,
		"createAlarmFactory",
		[]interface{}{alarmNamePrefix},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GlueJobMonitoring) CreatedAlarms() *[]*AlarmWithAnnotation {
	var returns *[]*AlarmWithAnnotation

	_jsii_.Invoke(
		g,
		"createdAlarms",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GlueJobMonitoring) CreateDataMovementWidget(width *float64, height *float64) awscloudwatch.GraphWidget {
	if err := g.validateCreateDataMovementWidgetParameters(width, height); err != nil {
		panic(err)
	}
	var returns awscloudwatch.GraphWidget

	_jsii_.Invoke(
		g,
		"createDataMovementWidget",
		[]interface{}{width, height},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GlueJobMonitoring) CreateErrorCountWidget(width *float64, height *float64) awscloudwatch.GraphWidget {
	if err := g.validateCreateErrorCountWidgetParameters(width, height); err != nil {
		panic(err)
	}
	var returns awscloudwatch.GraphWidget

	_jsii_.Invoke(
		g,
		"createErrorCountWidget",
		[]interface{}{width, height},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GlueJobMonitoring) CreateErrorRateWidget(width *float64, height *float64) awscloudwatch.GraphWidget {
	if err := g.validateCreateErrorRateWidgetParameters(width, height); err != nil {
		panic(err)
	}
	var returns awscloudwatch.GraphWidget

	_jsii_.Invoke(
		g,
		"createErrorRateWidget",
		[]interface{}{width, height},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GlueJobMonitoring) CreateJobExecutionWidget(width *float64, height *float64) awscloudwatch.GraphWidget {
	if err := g.validateCreateJobExecutionWidgetParameters(width, height); err != nil {
		panic(err)
	}
	var returns awscloudwatch.GraphWidget

	_jsii_.Invoke(
		g,
		"createJobExecutionWidget",
		[]interface{}{width, height},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GlueJobMonitoring) CreateMetricFactory() MetricFactory {
	var returns MetricFactory

	_jsii_.Invoke(
		g,
		"createMetricFactory",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GlueJobMonitoring) CreateTitleWidget() MonitoringHeaderWidget {
	var returns MonitoringHeaderWidget

	_jsii_.Invoke(
		g,
		"createTitleWidget",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GlueJobMonitoring) CreateUtilizationWidget(width *float64, height *float64) awscloudwatch.GraphWidget {
	if err := g.validateCreateUtilizationWidgetParameters(width, height); err != nil {
		panic(err)
	}
	var returns awscloudwatch.GraphWidget

	_jsii_.Invoke(
		g,
		"createUtilizationWidget",
		[]interface{}{width, height},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GlueJobMonitoring) CreateWidgetFactory() IWidgetFactory {
	var returns IWidgetFactory

	_jsii_.Invoke(
		g,
		"createWidgetFactory",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GlueJobMonitoring) SummaryWidgets() *[]awscloudwatch.IWidget {
	var returns *[]awscloudwatch.IWidget

	_jsii_.Invoke(
		g,
		"summaryWidgets",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GlueJobMonitoring) Widgets() *[]awscloudwatch.IWidget {
	var returns *[]awscloudwatch.IWidget

	_jsii_.Invoke(
		g,
		"widgets",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GlueJobMonitoring) WidgetsForDashboard(name *string) *[]awscloudwatch.IWidget {
	if err := g.validateWidgetsForDashboardParameters(name); err != nil {
		panic(err)
	}
	var returns *[]awscloudwatch.IWidget

	_jsii_.Invoke(
		g,
		"widgetsForDashboard",
		[]interface{}{name},
		&returns,
	)

	return returns
}


package cdkmonitoringconstructs

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/corymhall/cdk-constructs-go/cdkmonitoringconstructs/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudwatch"
)

// Experimental.
type EC2Monitoring interface {
	Monitoring
	// Experimental.
	Alarms() *[]*AlarmWithAnnotation
	// Experimental.
	CpuUtilisationMetrics() *[]awscloudwatch.IMetric
	// Experimental.
	DiskReadBytesMetrics() *[]awscloudwatch.IMetric
	// Experimental.
	DiskReadOpsMetrics() *[]awscloudwatch.IMetric
	// Experimental.
	DiskWriteBytesMetrics() *[]awscloudwatch.IMetric
	// Experimental.
	DiskWriteOpsMetrics() *[]awscloudwatch.IMetric
	// Experimental.
	Family() *string
	// Experimental.
	LocalAlarmNamePrefixOverride() *string
	// Experimental.
	NetworkInMetrics() *[]awscloudwatch.IMetric
	// Experimental.
	NetworkOutMetrics() *[]awscloudwatch.IMetric
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
	// Experimental.
	CreateCpuWidget(width *float64, height *float64) awscloudwatch.GraphWidget
	// Returns all the alarms created.
	// Experimental.
	CreatedAlarms() *[]*AlarmWithAnnotation
	// Experimental.
	CreateDiskOpsWidget(width *float64, height *float64) awscloudwatch.GraphWidget
	// Experimental.
	CreateDiskWidget(width *float64, height *float64) awscloudwatch.GraphWidget
	// Creates a new metric factory.
	// Experimental.
	CreateMetricFactory() MetricFactory
	// Experimental.
	CreateNetworkWidget(width *float64, height *float64) awscloudwatch.GraphWidget
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

// The jsii proxy struct for EC2Monitoring
type jsiiProxy_EC2Monitoring struct {
	jsiiProxy_Monitoring
}

func (j *jsiiProxy_EC2Monitoring) Alarms() *[]*AlarmWithAnnotation {
	var returns *[]*AlarmWithAnnotation
	_jsii_.Get(
		j,
		"alarms",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EC2Monitoring) CpuUtilisationMetrics() *[]awscloudwatch.IMetric {
	var returns *[]awscloudwatch.IMetric
	_jsii_.Get(
		j,
		"cpuUtilisationMetrics",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EC2Monitoring) DiskReadBytesMetrics() *[]awscloudwatch.IMetric {
	var returns *[]awscloudwatch.IMetric
	_jsii_.Get(
		j,
		"diskReadBytesMetrics",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EC2Monitoring) DiskReadOpsMetrics() *[]awscloudwatch.IMetric {
	var returns *[]awscloudwatch.IMetric
	_jsii_.Get(
		j,
		"diskReadOpsMetrics",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EC2Monitoring) DiskWriteBytesMetrics() *[]awscloudwatch.IMetric {
	var returns *[]awscloudwatch.IMetric
	_jsii_.Get(
		j,
		"diskWriteBytesMetrics",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EC2Monitoring) DiskWriteOpsMetrics() *[]awscloudwatch.IMetric {
	var returns *[]awscloudwatch.IMetric
	_jsii_.Get(
		j,
		"diskWriteOpsMetrics",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EC2Monitoring) Family() *string {
	var returns *string
	_jsii_.Get(
		j,
		"family",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EC2Monitoring) LocalAlarmNamePrefixOverride() *string {
	var returns *string
	_jsii_.Get(
		j,
		"localAlarmNamePrefixOverride",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EC2Monitoring) NetworkInMetrics() *[]awscloudwatch.IMetric {
	var returns *[]awscloudwatch.IMetric
	_jsii_.Get(
		j,
		"networkInMetrics",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EC2Monitoring) NetworkOutMetrics() *[]awscloudwatch.IMetric {
	var returns *[]awscloudwatch.IMetric
	_jsii_.Get(
		j,
		"networkOutMetrics",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EC2Monitoring) Scope() MonitoringScope {
	var returns MonitoringScope
	_jsii_.Get(
		j,
		"scope",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EC2Monitoring) Title() *string {
	var returns *string
	_jsii_.Get(
		j,
		"title",
		&returns,
	)
	return returns
}


// Experimental.
func NewEC2Monitoring(scope MonitoringScope, props *EC2MonitoringProps) EC2Monitoring {
	_init_.Initialize()

	if err := validateNewEC2MonitoringParameters(scope, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_EC2Monitoring{}

	_jsii_.Create(
		"cdk-monitoring-constructs.EC2Monitoring",
		[]interface{}{scope, props},
		&j,
	)

	return &j
}

// Experimental.
func NewEC2Monitoring_Override(e EC2Monitoring, scope MonitoringScope, props *EC2MonitoringProps) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk-monitoring-constructs.EC2Monitoring",
		[]interface{}{scope, props},
		e,
	)
}

func (e *jsiiProxy_EC2Monitoring) AddAlarm(alarm *AlarmWithAnnotation) {
	if err := e.validateAddAlarmParameters(alarm); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"addAlarm",
		[]interface{}{alarm},
	)
}

func (e *jsiiProxy_EC2Monitoring) AlarmWidgets() *[]awscloudwatch.IWidget {
	var returns *[]awscloudwatch.IWidget

	_jsii_.Invoke(
		e,
		"alarmWidgets",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EC2Monitoring) CreateAlarmFactory(alarmNamePrefix *string) AlarmFactory {
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

func (e *jsiiProxy_EC2Monitoring) CreateCpuWidget(width *float64, height *float64) awscloudwatch.GraphWidget {
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

func (e *jsiiProxy_EC2Monitoring) CreatedAlarms() *[]*AlarmWithAnnotation {
	var returns *[]*AlarmWithAnnotation

	_jsii_.Invoke(
		e,
		"createdAlarms",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EC2Monitoring) CreateDiskOpsWidget(width *float64, height *float64) awscloudwatch.GraphWidget {
	if err := e.validateCreateDiskOpsWidgetParameters(width, height); err != nil {
		panic(err)
	}
	var returns awscloudwatch.GraphWidget

	_jsii_.Invoke(
		e,
		"createDiskOpsWidget",
		[]interface{}{width, height},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EC2Monitoring) CreateDiskWidget(width *float64, height *float64) awscloudwatch.GraphWidget {
	if err := e.validateCreateDiskWidgetParameters(width, height); err != nil {
		panic(err)
	}
	var returns awscloudwatch.GraphWidget

	_jsii_.Invoke(
		e,
		"createDiskWidget",
		[]interface{}{width, height},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EC2Monitoring) CreateMetricFactory() MetricFactory {
	var returns MetricFactory

	_jsii_.Invoke(
		e,
		"createMetricFactory",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EC2Monitoring) CreateNetworkWidget(width *float64, height *float64) awscloudwatch.GraphWidget {
	if err := e.validateCreateNetworkWidgetParameters(width, height); err != nil {
		panic(err)
	}
	var returns awscloudwatch.GraphWidget

	_jsii_.Invoke(
		e,
		"createNetworkWidget",
		[]interface{}{width, height},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EC2Monitoring) CreateTitleWidget() MonitoringHeaderWidget {
	var returns MonitoringHeaderWidget

	_jsii_.Invoke(
		e,
		"createTitleWidget",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EC2Monitoring) CreateWidgetFactory() IWidgetFactory {
	var returns IWidgetFactory

	_jsii_.Invoke(
		e,
		"createWidgetFactory",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EC2Monitoring) SummaryWidgets() *[]awscloudwatch.IWidget {
	var returns *[]awscloudwatch.IWidget

	_jsii_.Invoke(
		e,
		"summaryWidgets",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EC2Monitoring) Widgets() *[]awscloudwatch.IWidget {
	var returns *[]awscloudwatch.IWidget

	_jsii_.Invoke(
		e,
		"widgets",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EC2Monitoring) WidgetsForDashboard(name *string) *[]awscloudwatch.IWidget {
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


package cdkmonitoringconstructs

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/corymhall/cdk-constructs-go/cdkmonitoringconstructs/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudwatch"
)

// Experimental.
type KinesisDataAnalyticsMonitoring interface {
	Monitoring
	// Experimental.
	Alarms() *[]*AlarmWithAnnotation
	// Experimental.
	CpuUtilizationPercentMetric() interface{}
	// Experimental.
	DowntimeAnnotations() *[]*awscloudwatch.HorizontalAnnotation
	// Experimental.
	DowntimeMsMetric() interface{}
	// Experimental.
	FullRestartAnnotations() *[]*awscloudwatch.HorizontalAnnotation
	// Experimental.
	FullRestartsCountMetric() interface{}
	// Experimental.
	HeapMemoryUtilizationPercentMetric() interface{}
	// Experimental.
	KdaAlarmFactory() KinesisDataAnalyticsAlarmFactory
	// Experimental.
	KinesisDataAnalyticsUrl() *string
	// Experimental.
	KpusCountMetric() interface{}
	// Experimental.
	LastCheckpointDurationMsMetric() interface{}
	// Experimental.
	LastCheckpointSizeBytesMetric() interface{}
	// Experimental.
	LocalAlarmNamePrefixOverride() *string
	// Experimental.
	NumberOfFailedCheckpointsCountMetric() interface{}
	// Experimental.
	OldGenerationGCCountMetric() interface{}
	// Experimental.
	OldGenerationGCTimeMsMetric() interface{}
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
	CreateDownTimeWidget(width *float64, height *float64) awscloudwatch.GraphWidget
	// Experimental.
	CreateFullRestartsWidget(width *float64, height *float64) awscloudwatch.GraphWidget
	// Experimental.
	CreateGarbageCollectionWidget(width *float64, height *float64) awscloudwatch.GraphWidget
	// Experimental.
	CreateKPUWidget(width *float64, height *float64) awscloudwatch.GraphWidget
	// Experimental.
	CreateLastCheckpointDurationWidget(width *float64, height *float64) awscloudwatch.GraphWidget
	// Experimental.
	CreateLastCheckpointSizeWidget(width *float64, height *float64) awscloudwatch.GraphWidget
	// Creates a new metric factory.
	// Experimental.
	CreateMetricFactory() MetricFactory
	// Experimental.
	CreateNumberOfFailedCheckpointsWidget(width *float64, height *float64) awscloudwatch.GraphWidget
	// Experimental.
	CreateResourceUtilizationWidget(width *float64, height *float64) awscloudwatch.GraphWidget
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

// The jsii proxy struct for KinesisDataAnalyticsMonitoring
type jsiiProxy_KinesisDataAnalyticsMonitoring struct {
	jsiiProxy_Monitoring
}

func (j *jsiiProxy_KinesisDataAnalyticsMonitoring) Alarms() *[]*AlarmWithAnnotation {
	var returns *[]*AlarmWithAnnotation
	_jsii_.Get(
		j,
		"alarms",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisDataAnalyticsMonitoring) CpuUtilizationPercentMetric() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"cpuUtilizationPercentMetric",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisDataAnalyticsMonitoring) DowntimeAnnotations() *[]*awscloudwatch.HorizontalAnnotation {
	var returns *[]*awscloudwatch.HorizontalAnnotation
	_jsii_.Get(
		j,
		"downtimeAnnotations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisDataAnalyticsMonitoring) DowntimeMsMetric() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"downtimeMsMetric",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisDataAnalyticsMonitoring) FullRestartAnnotations() *[]*awscloudwatch.HorizontalAnnotation {
	var returns *[]*awscloudwatch.HorizontalAnnotation
	_jsii_.Get(
		j,
		"fullRestartAnnotations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisDataAnalyticsMonitoring) FullRestartsCountMetric() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"fullRestartsCountMetric",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisDataAnalyticsMonitoring) HeapMemoryUtilizationPercentMetric() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"heapMemoryUtilizationPercentMetric",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisDataAnalyticsMonitoring) KdaAlarmFactory() KinesisDataAnalyticsAlarmFactory {
	var returns KinesisDataAnalyticsAlarmFactory
	_jsii_.Get(
		j,
		"kdaAlarmFactory",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisDataAnalyticsMonitoring) KinesisDataAnalyticsUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kinesisDataAnalyticsUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisDataAnalyticsMonitoring) KpusCountMetric() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"kpusCountMetric",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisDataAnalyticsMonitoring) LastCheckpointDurationMsMetric() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"lastCheckpointDurationMsMetric",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisDataAnalyticsMonitoring) LastCheckpointSizeBytesMetric() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"lastCheckpointSizeBytesMetric",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisDataAnalyticsMonitoring) LocalAlarmNamePrefixOverride() *string {
	var returns *string
	_jsii_.Get(
		j,
		"localAlarmNamePrefixOverride",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisDataAnalyticsMonitoring) NumberOfFailedCheckpointsCountMetric() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"numberOfFailedCheckpointsCountMetric",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisDataAnalyticsMonitoring) OldGenerationGCCountMetric() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"oldGenerationGCCountMetric",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisDataAnalyticsMonitoring) OldGenerationGCTimeMsMetric() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"oldGenerationGCTimeMsMetric",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisDataAnalyticsMonitoring) Scope() MonitoringScope {
	var returns MonitoringScope
	_jsii_.Get(
		j,
		"scope",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisDataAnalyticsMonitoring) Title() *string {
	var returns *string
	_jsii_.Get(
		j,
		"title",
		&returns,
	)
	return returns
}


// Experimental.
func NewKinesisDataAnalyticsMonitoring(scope MonitoringScope, props *KinesisDataAnalyticsMonitoringProps) KinesisDataAnalyticsMonitoring {
	_init_.Initialize()

	if err := validateNewKinesisDataAnalyticsMonitoringParameters(scope, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_KinesisDataAnalyticsMonitoring{}

	_jsii_.Create(
		"cdk-monitoring-constructs.KinesisDataAnalyticsMonitoring",
		[]interface{}{scope, props},
		&j,
	)

	return &j
}

// Experimental.
func NewKinesisDataAnalyticsMonitoring_Override(k KinesisDataAnalyticsMonitoring, scope MonitoringScope, props *KinesisDataAnalyticsMonitoringProps) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk-monitoring-constructs.KinesisDataAnalyticsMonitoring",
		[]interface{}{scope, props},
		k,
	)
}

func (k *jsiiProxy_KinesisDataAnalyticsMonitoring) AddAlarm(alarm *AlarmWithAnnotation) {
	if err := k.validateAddAlarmParameters(alarm); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"addAlarm",
		[]interface{}{alarm},
	)
}

func (k *jsiiProxy_KinesisDataAnalyticsMonitoring) AlarmWidgets() *[]awscloudwatch.IWidget {
	var returns *[]awscloudwatch.IWidget

	_jsii_.Invoke(
		k,
		"alarmWidgets",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KinesisDataAnalyticsMonitoring) CreateAlarmFactory(alarmNamePrefix *string) AlarmFactory {
	if err := k.validateCreateAlarmFactoryParameters(alarmNamePrefix); err != nil {
		panic(err)
	}
	var returns AlarmFactory

	_jsii_.Invoke(
		k,
		"createAlarmFactory",
		[]interface{}{alarmNamePrefix},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KinesisDataAnalyticsMonitoring) CreatedAlarms() *[]*AlarmWithAnnotation {
	var returns *[]*AlarmWithAnnotation

	_jsii_.Invoke(
		k,
		"createdAlarms",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KinesisDataAnalyticsMonitoring) CreateDownTimeWidget(width *float64, height *float64) awscloudwatch.GraphWidget {
	if err := k.validateCreateDownTimeWidgetParameters(width, height); err != nil {
		panic(err)
	}
	var returns awscloudwatch.GraphWidget

	_jsii_.Invoke(
		k,
		"createDownTimeWidget",
		[]interface{}{width, height},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KinesisDataAnalyticsMonitoring) CreateFullRestartsWidget(width *float64, height *float64) awscloudwatch.GraphWidget {
	if err := k.validateCreateFullRestartsWidgetParameters(width, height); err != nil {
		panic(err)
	}
	var returns awscloudwatch.GraphWidget

	_jsii_.Invoke(
		k,
		"createFullRestartsWidget",
		[]interface{}{width, height},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KinesisDataAnalyticsMonitoring) CreateGarbageCollectionWidget(width *float64, height *float64) awscloudwatch.GraphWidget {
	if err := k.validateCreateGarbageCollectionWidgetParameters(width, height); err != nil {
		panic(err)
	}
	var returns awscloudwatch.GraphWidget

	_jsii_.Invoke(
		k,
		"createGarbageCollectionWidget",
		[]interface{}{width, height},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KinesisDataAnalyticsMonitoring) CreateKPUWidget(width *float64, height *float64) awscloudwatch.GraphWidget {
	if err := k.validateCreateKPUWidgetParameters(width, height); err != nil {
		panic(err)
	}
	var returns awscloudwatch.GraphWidget

	_jsii_.Invoke(
		k,
		"createKPUWidget",
		[]interface{}{width, height},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KinesisDataAnalyticsMonitoring) CreateLastCheckpointDurationWidget(width *float64, height *float64) awscloudwatch.GraphWidget {
	if err := k.validateCreateLastCheckpointDurationWidgetParameters(width, height); err != nil {
		panic(err)
	}
	var returns awscloudwatch.GraphWidget

	_jsii_.Invoke(
		k,
		"createLastCheckpointDurationWidget",
		[]interface{}{width, height},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KinesisDataAnalyticsMonitoring) CreateLastCheckpointSizeWidget(width *float64, height *float64) awscloudwatch.GraphWidget {
	if err := k.validateCreateLastCheckpointSizeWidgetParameters(width, height); err != nil {
		panic(err)
	}
	var returns awscloudwatch.GraphWidget

	_jsii_.Invoke(
		k,
		"createLastCheckpointSizeWidget",
		[]interface{}{width, height},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KinesisDataAnalyticsMonitoring) CreateMetricFactory() MetricFactory {
	var returns MetricFactory

	_jsii_.Invoke(
		k,
		"createMetricFactory",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KinesisDataAnalyticsMonitoring) CreateNumberOfFailedCheckpointsWidget(width *float64, height *float64) awscloudwatch.GraphWidget {
	if err := k.validateCreateNumberOfFailedCheckpointsWidgetParameters(width, height); err != nil {
		panic(err)
	}
	var returns awscloudwatch.GraphWidget

	_jsii_.Invoke(
		k,
		"createNumberOfFailedCheckpointsWidget",
		[]interface{}{width, height},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KinesisDataAnalyticsMonitoring) CreateResourceUtilizationWidget(width *float64, height *float64) awscloudwatch.GraphWidget {
	if err := k.validateCreateResourceUtilizationWidgetParameters(width, height); err != nil {
		panic(err)
	}
	var returns awscloudwatch.GraphWidget

	_jsii_.Invoke(
		k,
		"createResourceUtilizationWidget",
		[]interface{}{width, height},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KinesisDataAnalyticsMonitoring) CreateTitleWidget() MonitoringHeaderWidget {
	var returns MonitoringHeaderWidget

	_jsii_.Invoke(
		k,
		"createTitleWidget",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KinesisDataAnalyticsMonitoring) CreateWidgetFactory() IWidgetFactory {
	var returns IWidgetFactory

	_jsii_.Invoke(
		k,
		"createWidgetFactory",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KinesisDataAnalyticsMonitoring) SummaryWidgets() *[]awscloudwatch.IWidget {
	var returns *[]awscloudwatch.IWidget

	_jsii_.Invoke(
		k,
		"summaryWidgets",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KinesisDataAnalyticsMonitoring) Widgets() *[]awscloudwatch.IWidget {
	var returns *[]awscloudwatch.IWidget

	_jsii_.Invoke(
		k,
		"widgets",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KinesisDataAnalyticsMonitoring) WidgetsForDashboard(name *string) *[]awscloudwatch.IWidget {
	if err := k.validateWidgetsForDashboardParameters(name); err != nil {
		panic(err)
	}
	var returns *[]awscloudwatch.IWidget

	_jsii_.Invoke(
		k,
		"widgetsForDashboard",
		[]interface{}{name},
		&returns,
	)

	return returns
}


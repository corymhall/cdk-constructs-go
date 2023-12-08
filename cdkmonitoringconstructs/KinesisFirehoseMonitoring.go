package cdkmonitoringconstructs

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/corymhall/cdk-constructs-go/cdkmonitoringconstructs/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudwatch"
)

// Experimental.
type KinesisFirehoseMonitoring interface {
	Monitoring
	// Experimental.
	Alarms() *[]*AlarmWithAnnotation
	// Experimental.
	FailedConversionMetric() interface{}
	// Experimental.
	IncomingBytesMetric() interface{}
	// Experimental.
	IncomingBytesToLimitRate() interface{}
	// Experimental.
	IncomingLimitAnnotations() *[]*awscloudwatch.HorizontalAnnotation
	// Experimental.
	IncomingPutRequestsToLimitRate() interface{}
	// Experimental.
	IncomingRecordsMetric() interface{}
	// Experimental.
	IncomingRecordsToLimitRate() interface{}
	// Experimental.
	KinesisAlarmFactory() KinesisAlarmFactory
	// Experimental.
	LocalAlarmNamePrefixOverride() *string
	// Experimental.
	PutRecordBatchLatency() interface{}
	// Experimental.
	PutRecordLatency() interface{}
	// Experimental.
	RecordCountAnnotations() *[]*awscloudwatch.HorizontalAnnotation
	// Experimental.
	Scope() MonitoringScope
	// Experimental.
	StreamUrl() *string
	// Experimental.
	SuccessfulConversionMetric() interface{}
	// Experimental.
	ThrottledRecordsMetric() interface{}
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
	CreateConversionWidget(width *float64, height *float64) awscloudwatch.GraphWidget
	// Returns all the alarms created.
	// Experimental.
	CreatedAlarms() *[]*AlarmWithAnnotation
	// Experimental.
	CreateIncomingRecordWidget(width *float64, height *float64) awscloudwatch.GraphWidget
	// Experimental.
	CreateLatencyWidget(width *float64, height *float64) awscloudwatch.GraphWidget
	// Experimental.
	CreateLimitWidget(width *float64, height *float64) awscloudwatch.GraphWidget
	// Creates a new metric factory.
	// Experimental.
	CreateMetricFactory() MetricFactory
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

// The jsii proxy struct for KinesisFirehoseMonitoring
type jsiiProxy_KinesisFirehoseMonitoring struct {
	jsiiProxy_Monitoring
}

func (j *jsiiProxy_KinesisFirehoseMonitoring) Alarms() *[]*AlarmWithAnnotation {
	var returns *[]*AlarmWithAnnotation
	_jsii_.Get(
		j,
		"alarms",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisFirehoseMonitoring) FailedConversionMetric() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"failedConversionMetric",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisFirehoseMonitoring) IncomingBytesMetric() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"incomingBytesMetric",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisFirehoseMonitoring) IncomingBytesToLimitRate() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"incomingBytesToLimitRate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisFirehoseMonitoring) IncomingLimitAnnotations() *[]*awscloudwatch.HorizontalAnnotation {
	var returns *[]*awscloudwatch.HorizontalAnnotation
	_jsii_.Get(
		j,
		"incomingLimitAnnotations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisFirehoseMonitoring) IncomingPutRequestsToLimitRate() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"incomingPutRequestsToLimitRate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisFirehoseMonitoring) IncomingRecordsMetric() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"incomingRecordsMetric",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisFirehoseMonitoring) IncomingRecordsToLimitRate() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"incomingRecordsToLimitRate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisFirehoseMonitoring) KinesisAlarmFactory() KinesisAlarmFactory {
	var returns KinesisAlarmFactory
	_jsii_.Get(
		j,
		"kinesisAlarmFactory",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisFirehoseMonitoring) LocalAlarmNamePrefixOverride() *string {
	var returns *string
	_jsii_.Get(
		j,
		"localAlarmNamePrefixOverride",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisFirehoseMonitoring) PutRecordBatchLatency() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"putRecordBatchLatency",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisFirehoseMonitoring) PutRecordLatency() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"putRecordLatency",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisFirehoseMonitoring) RecordCountAnnotations() *[]*awscloudwatch.HorizontalAnnotation {
	var returns *[]*awscloudwatch.HorizontalAnnotation
	_jsii_.Get(
		j,
		"recordCountAnnotations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisFirehoseMonitoring) Scope() MonitoringScope {
	var returns MonitoringScope
	_jsii_.Get(
		j,
		"scope",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisFirehoseMonitoring) StreamUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"streamUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisFirehoseMonitoring) SuccessfulConversionMetric() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"successfulConversionMetric",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisFirehoseMonitoring) ThrottledRecordsMetric() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"throttledRecordsMetric",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisFirehoseMonitoring) Title() *string {
	var returns *string
	_jsii_.Get(
		j,
		"title",
		&returns,
	)
	return returns
}


// Experimental.
func NewKinesisFirehoseMonitoring(scope MonitoringScope, props *KinesisFirehoseMonitoringProps) KinesisFirehoseMonitoring {
	_init_.Initialize()

	if err := validateNewKinesisFirehoseMonitoringParameters(scope, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_KinesisFirehoseMonitoring{}

	_jsii_.Create(
		"cdk-monitoring-constructs.KinesisFirehoseMonitoring",
		[]interface{}{scope, props},
		&j,
	)

	return &j
}

// Experimental.
func NewKinesisFirehoseMonitoring_Override(k KinesisFirehoseMonitoring, scope MonitoringScope, props *KinesisFirehoseMonitoringProps) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk-monitoring-constructs.KinesisFirehoseMonitoring",
		[]interface{}{scope, props},
		k,
	)
}

func (k *jsiiProxy_KinesisFirehoseMonitoring) AddAlarm(alarm *AlarmWithAnnotation) {
	if err := k.validateAddAlarmParameters(alarm); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"addAlarm",
		[]interface{}{alarm},
	)
}

func (k *jsiiProxy_KinesisFirehoseMonitoring) AlarmWidgets() *[]awscloudwatch.IWidget {
	var returns *[]awscloudwatch.IWidget

	_jsii_.Invoke(
		k,
		"alarmWidgets",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KinesisFirehoseMonitoring) CreateAlarmFactory(alarmNamePrefix *string) AlarmFactory {
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

func (k *jsiiProxy_KinesisFirehoseMonitoring) CreateConversionWidget(width *float64, height *float64) awscloudwatch.GraphWidget {
	if err := k.validateCreateConversionWidgetParameters(width, height); err != nil {
		panic(err)
	}
	var returns awscloudwatch.GraphWidget

	_jsii_.Invoke(
		k,
		"createConversionWidget",
		[]interface{}{width, height},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KinesisFirehoseMonitoring) CreatedAlarms() *[]*AlarmWithAnnotation {
	var returns *[]*AlarmWithAnnotation

	_jsii_.Invoke(
		k,
		"createdAlarms",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KinesisFirehoseMonitoring) CreateIncomingRecordWidget(width *float64, height *float64) awscloudwatch.GraphWidget {
	if err := k.validateCreateIncomingRecordWidgetParameters(width, height); err != nil {
		panic(err)
	}
	var returns awscloudwatch.GraphWidget

	_jsii_.Invoke(
		k,
		"createIncomingRecordWidget",
		[]interface{}{width, height},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KinesisFirehoseMonitoring) CreateLatencyWidget(width *float64, height *float64) awscloudwatch.GraphWidget {
	if err := k.validateCreateLatencyWidgetParameters(width, height); err != nil {
		panic(err)
	}
	var returns awscloudwatch.GraphWidget

	_jsii_.Invoke(
		k,
		"createLatencyWidget",
		[]interface{}{width, height},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KinesisFirehoseMonitoring) CreateLimitWidget(width *float64, height *float64) awscloudwatch.GraphWidget {
	if err := k.validateCreateLimitWidgetParameters(width, height); err != nil {
		panic(err)
	}
	var returns awscloudwatch.GraphWidget

	_jsii_.Invoke(
		k,
		"createLimitWidget",
		[]interface{}{width, height},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KinesisFirehoseMonitoring) CreateMetricFactory() MetricFactory {
	var returns MetricFactory

	_jsii_.Invoke(
		k,
		"createMetricFactory",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KinesisFirehoseMonitoring) CreateTitleWidget() MonitoringHeaderWidget {
	var returns MonitoringHeaderWidget

	_jsii_.Invoke(
		k,
		"createTitleWidget",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KinesisFirehoseMonitoring) CreateWidgetFactory() IWidgetFactory {
	var returns IWidgetFactory

	_jsii_.Invoke(
		k,
		"createWidgetFactory",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KinesisFirehoseMonitoring) SummaryWidgets() *[]awscloudwatch.IWidget {
	var returns *[]awscloudwatch.IWidget

	_jsii_.Invoke(
		k,
		"summaryWidgets",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KinesisFirehoseMonitoring) Widgets() *[]awscloudwatch.IWidget {
	var returns *[]awscloudwatch.IWidget

	_jsii_.Invoke(
		k,
		"widgets",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KinesisFirehoseMonitoring) WidgetsForDashboard(name *string) *[]awscloudwatch.IWidget {
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


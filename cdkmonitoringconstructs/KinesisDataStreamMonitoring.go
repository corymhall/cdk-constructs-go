package cdkmonitoringconstructs

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/corymhall/cdk-constructs-go/cdkmonitoringconstructs/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudwatch"
)

// Experimental.
type KinesisDataStreamMonitoring interface {
	Monitoring
	// Experimental.
	AgeAnnotations() *[]*awscloudwatch.HorizontalAnnotation
	// Experimental.
	Alarms() *[]*AlarmWithAnnotation
	// Experimental.
	IncomingDataSumBytesMetric() interface{}
	// Experimental.
	IncomingDataSumCountMetric() interface{}
	// Experimental.
	KinesisAlarmFactory() KinesisAlarmFactory
	// Experimental.
	LocalAlarmNamePrefixOverride() *string
	// Experimental.
	MetricGetRecordsIteratorAge() interface{}
	// Experimental.
	MetricGetRecordsLatencyAverage() interface{}
	// Experimental.
	MetricGetRecordsSuccessCount() interface{}
	// Experimental.
	MetricGetRecordsSumCount() interface{}
	// Experimental.
	MetricGetRecordSumBytes() interface{}
	// Experimental.
	ProvisionedCapacityAnnotations() *[]*awscloudwatch.HorizontalAnnotation
	// Experimental.
	PutRecordLatencyAverageMetric() interface{}
	// Experimental.
	PutRecordsFailedRecordsCountMetric() interface{}
	// Experimental.
	PutRecordsLatencyAverageMetric() interface{}
	// Experimental.
	PutRecordsSuccessCountMetric() interface{}
	// Experimental.
	PutRecordsSuccessfulRecordsCountMetric() interface{}
	// Experimental.
	PutRecordsSumBytesMetric() interface{}
	// Experimental.
	PutRecordsThrottledRecordsCountMetric() interface{}
	// Experimental.
	PutRecordsTotalRecordsCountMetric() interface{}
	// Experimental.
	PutRecordSuccessCountMetric() interface{}
	// Experimental.
	PutRecordSumBytesMetric() interface{}
	// Experimental.
	ReadProvisionedThroughputExceededMetric() interface{}
	// Experimental.
	RecordCountAnnotations() *[]*awscloudwatch.HorizontalAnnotation
	// Experimental.
	Scope() MonitoringScope
	// Experimental.
	StreamUrl() *string
	// Experimental.
	Title() *string
	// Experimental.
	WriteProvisionedThroughputExceededMetric() interface{}
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
	CreateCapacityWidget(width *float64, height *float64) awscloudwatch.GraphWidget
	// Returns all the alarms created.
	// Experimental.
	CreatedAlarms() *[]*AlarmWithAnnotation
	// Experimental.
	CreateIncomingDataWidget(width *float64, height *float64) awscloudwatch.GraphWidget
	// Experimental.
	CreateIteratorAgeWidget(width *float64, height *float64) awscloudwatch.GraphWidget
	// Experimental.
	CreateLatencyWidget(width *float64, height *float64) awscloudwatch.GraphWidget
	// Creates a new metric factory.
	// Experimental.
	CreateMetricFactory() MetricFactory
	// Experimental.
	CreateOperationWidget(width *float64, height *float64) awscloudwatch.GraphWidget
	// Experimental.
	CreateRecordNumberWidget(width *float64, height *float64) awscloudwatch.GraphWidget
	// Experimental.
	CreateRecordSizeWidget(width *float64, height *float64) awscloudwatch.GraphWidget
	// Experimental.
	CreateSecondAdditionalRow() awscloudwatch.Row
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

// The jsii proxy struct for KinesisDataStreamMonitoring
type jsiiProxy_KinesisDataStreamMonitoring struct {
	jsiiProxy_Monitoring
}

func (j *jsiiProxy_KinesisDataStreamMonitoring) AgeAnnotations() *[]*awscloudwatch.HorizontalAnnotation {
	var returns *[]*awscloudwatch.HorizontalAnnotation
	_jsii_.Get(
		j,
		"ageAnnotations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisDataStreamMonitoring) Alarms() *[]*AlarmWithAnnotation {
	var returns *[]*AlarmWithAnnotation
	_jsii_.Get(
		j,
		"alarms",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisDataStreamMonitoring) IncomingDataSumBytesMetric() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"incomingDataSumBytesMetric",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisDataStreamMonitoring) IncomingDataSumCountMetric() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"incomingDataSumCountMetric",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisDataStreamMonitoring) KinesisAlarmFactory() KinesisAlarmFactory {
	var returns KinesisAlarmFactory
	_jsii_.Get(
		j,
		"kinesisAlarmFactory",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisDataStreamMonitoring) LocalAlarmNamePrefixOverride() *string {
	var returns *string
	_jsii_.Get(
		j,
		"localAlarmNamePrefixOverride",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisDataStreamMonitoring) MetricGetRecordsIteratorAge() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"metricGetRecordsIteratorAge",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisDataStreamMonitoring) MetricGetRecordsLatencyAverage() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"metricGetRecordsLatencyAverage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisDataStreamMonitoring) MetricGetRecordsSuccessCount() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"metricGetRecordsSuccessCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisDataStreamMonitoring) MetricGetRecordsSumCount() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"metricGetRecordsSumCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisDataStreamMonitoring) MetricGetRecordSumBytes() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"metricGetRecordSumBytes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisDataStreamMonitoring) ProvisionedCapacityAnnotations() *[]*awscloudwatch.HorizontalAnnotation {
	var returns *[]*awscloudwatch.HorizontalAnnotation
	_jsii_.Get(
		j,
		"provisionedCapacityAnnotations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisDataStreamMonitoring) PutRecordLatencyAverageMetric() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"putRecordLatencyAverageMetric",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisDataStreamMonitoring) PutRecordsFailedRecordsCountMetric() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"putRecordsFailedRecordsCountMetric",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisDataStreamMonitoring) PutRecordsLatencyAverageMetric() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"putRecordsLatencyAverageMetric",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisDataStreamMonitoring) PutRecordsSuccessCountMetric() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"putRecordsSuccessCountMetric",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisDataStreamMonitoring) PutRecordsSuccessfulRecordsCountMetric() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"putRecordsSuccessfulRecordsCountMetric",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisDataStreamMonitoring) PutRecordsSumBytesMetric() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"putRecordsSumBytesMetric",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisDataStreamMonitoring) PutRecordsThrottledRecordsCountMetric() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"putRecordsThrottledRecordsCountMetric",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisDataStreamMonitoring) PutRecordsTotalRecordsCountMetric() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"putRecordsTotalRecordsCountMetric",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisDataStreamMonitoring) PutRecordSuccessCountMetric() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"putRecordSuccessCountMetric",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisDataStreamMonitoring) PutRecordSumBytesMetric() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"putRecordSumBytesMetric",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisDataStreamMonitoring) ReadProvisionedThroughputExceededMetric() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"readProvisionedThroughputExceededMetric",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisDataStreamMonitoring) RecordCountAnnotations() *[]*awscloudwatch.HorizontalAnnotation {
	var returns *[]*awscloudwatch.HorizontalAnnotation
	_jsii_.Get(
		j,
		"recordCountAnnotations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisDataStreamMonitoring) Scope() MonitoringScope {
	var returns MonitoringScope
	_jsii_.Get(
		j,
		"scope",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisDataStreamMonitoring) StreamUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"streamUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisDataStreamMonitoring) Title() *string {
	var returns *string
	_jsii_.Get(
		j,
		"title",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisDataStreamMonitoring) WriteProvisionedThroughputExceededMetric() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"writeProvisionedThroughputExceededMetric",
		&returns,
	)
	return returns
}


// Experimental.
func NewKinesisDataStreamMonitoring(scope MonitoringScope, props *KinesisDataStreamMonitoringProps) KinesisDataStreamMonitoring {
	_init_.Initialize()

	if err := validateNewKinesisDataStreamMonitoringParameters(scope, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_KinesisDataStreamMonitoring{}

	_jsii_.Create(
		"cdk-monitoring-constructs.KinesisDataStreamMonitoring",
		[]interface{}{scope, props},
		&j,
	)

	return &j
}

// Experimental.
func NewKinesisDataStreamMonitoring_Override(k KinesisDataStreamMonitoring, scope MonitoringScope, props *KinesisDataStreamMonitoringProps) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk-monitoring-constructs.KinesisDataStreamMonitoring",
		[]interface{}{scope, props},
		k,
	)
}

func (k *jsiiProxy_KinesisDataStreamMonitoring) AddAlarm(alarm *AlarmWithAnnotation) {
	if err := k.validateAddAlarmParameters(alarm); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"addAlarm",
		[]interface{}{alarm},
	)
}

func (k *jsiiProxy_KinesisDataStreamMonitoring) AlarmWidgets() *[]awscloudwatch.IWidget {
	var returns *[]awscloudwatch.IWidget

	_jsii_.Invoke(
		k,
		"alarmWidgets",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KinesisDataStreamMonitoring) CreateAlarmFactory(alarmNamePrefix *string) AlarmFactory {
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

func (k *jsiiProxy_KinesisDataStreamMonitoring) CreateCapacityWidget(width *float64, height *float64) awscloudwatch.GraphWidget {
	if err := k.validateCreateCapacityWidgetParameters(width, height); err != nil {
		panic(err)
	}
	var returns awscloudwatch.GraphWidget

	_jsii_.Invoke(
		k,
		"createCapacityWidget",
		[]interface{}{width, height},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KinesisDataStreamMonitoring) CreatedAlarms() *[]*AlarmWithAnnotation {
	var returns *[]*AlarmWithAnnotation

	_jsii_.Invoke(
		k,
		"createdAlarms",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KinesisDataStreamMonitoring) CreateIncomingDataWidget(width *float64, height *float64) awscloudwatch.GraphWidget {
	if err := k.validateCreateIncomingDataWidgetParameters(width, height); err != nil {
		panic(err)
	}
	var returns awscloudwatch.GraphWidget

	_jsii_.Invoke(
		k,
		"createIncomingDataWidget",
		[]interface{}{width, height},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KinesisDataStreamMonitoring) CreateIteratorAgeWidget(width *float64, height *float64) awscloudwatch.GraphWidget {
	if err := k.validateCreateIteratorAgeWidgetParameters(width, height); err != nil {
		panic(err)
	}
	var returns awscloudwatch.GraphWidget

	_jsii_.Invoke(
		k,
		"createIteratorAgeWidget",
		[]interface{}{width, height},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KinesisDataStreamMonitoring) CreateLatencyWidget(width *float64, height *float64) awscloudwatch.GraphWidget {
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

func (k *jsiiProxy_KinesisDataStreamMonitoring) CreateMetricFactory() MetricFactory {
	var returns MetricFactory

	_jsii_.Invoke(
		k,
		"createMetricFactory",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KinesisDataStreamMonitoring) CreateOperationWidget(width *float64, height *float64) awscloudwatch.GraphWidget {
	if err := k.validateCreateOperationWidgetParameters(width, height); err != nil {
		panic(err)
	}
	var returns awscloudwatch.GraphWidget

	_jsii_.Invoke(
		k,
		"createOperationWidget",
		[]interface{}{width, height},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KinesisDataStreamMonitoring) CreateRecordNumberWidget(width *float64, height *float64) awscloudwatch.GraphWidget {
	if err := k.validateCreateRecordNumberWidgetParameters(width, height); err != nil {
		panic(err)
	}
	var returns awscloudwatch.GraphWidget

	_jsii_.Invoke(
		k,
		"createRecordNumberWidget",
		[]interface{}{width, height},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KinesisDataStreamMonitoring) CreateRecordSizeWidget(width *float64, height *float64) awscloudwatch.GraphWidget {
	if err := k.validateCreateRecordSizeWidgetParameters(width, height); err != nil {
		panic(err)
	}
	var returns awscloudwatch.GraphWidget

	_jsii_.Invoke(
		k,
		"createRecordSizeWidget",
		[]interface{}{width, height},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KinesisDataStreamMonitoring) CreateSecondAdditionalRow() awscloudwatch.Row {
	var returns awscloudwatch.Row

	_jsii_.Invoke(
		k,
		"createSecondAdditionalRow",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KinesisDataStreamMonitoring) CreateTitleWidget() MonitoringHeaderWidget {
	var returns MonitoringHeaderWidget

	_jsii_.Invoke(
		k,
		"createTitleWidget",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KinesisDataStreamMonitoring) CreateWidgetFactory() IWidgetFactory {
	var returns IWidgetFactory

	_jsii_.Invoke(
		k,
		"createWidgetFactory",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KinesisDataStreamMonitoring) SummaryWidgets() *[]awscloudwatch.IWidget {
	var returns *[]awscloudwatch.IWidget

	_jsii_.Invoke(
		k,
		"summaryWidgets",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KinesisDataStreamMonitoring) Widgets() *[]awscloudwatch.IWidget {
	var returns *[]awscloudwatch.IWidget

	_jsii_.Invoke(
		k,
		"widgets",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KinesisDataStreamMonitoring) WidgetsForDashboard(name *string) *[]awscloudwatch.IWidget {
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


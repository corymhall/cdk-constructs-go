package cdkmonitoringconstructs

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/corymhall/cdk-constructs-go/cdkmonitoringconstructs/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudwatch"
	"github.com/aws/aws-cdk-go/awscdk/v2/awssqs"
)

// Experimental.
type SqsQueueMonitoringWithDlq interface {
	SqsQueueMonitoring
	// Experimental.
	AddDeadLetterQueueToSummaryDashboard() *bool
	// Experimental.
	AgeAnnotations() *[]*awscloudwatch.HorizontalAnnotation
	// Experimental.
	Alarms() *[]*AlarmWithAnnotation
	// Experimental.
	ConsumptionRateMetric() interface{}
	// Experimental.
	CountAnnotations() *[]*awscloudwatch.HorizontalAnnotation
	// Experimental.
	DeadLetterAgeAnnotations() *[]*awscloudwatch.HorizontalAnnotation
	// Experimental.
	DeadLetterCountAnnotations() *[]*awscloudwatch.HorizontalAnnotation
	// Experimental.
	DeadLetterQueueAlarmFactory() QueueAlarmFactory
	// Experimental.
	DeadLetterQueueIncomingMessagesMetric() interface{}
	// Experimental.
	DeadLetterQueueOldestMessageAgeMetric() interface{}
	// Experimental.
	DeadLetterQueueVisibleMessagesMetric() interface{}
	// Experimental.
	DeadLetterTitle() *string
	// Experimental.
	DeadLetterUrl() *string
	// Experimental.
	DeletedMessagesMetric() interface{}
	// Experimental.
	IncomingMessagesMetric() interface{}
	// Experimental.
	LocalAlarmNamePrefixOverride() *string
	// Experimental.
	MessageSizeMetric() interface{}
	// Experimental.
	OldestMessageAgeMetric() interface{}
	// Experimental.
	ProductionRateMetric() interface{}
	// Experimental.
	QueueAlarmFactory() QueueAlarmFactory
	// Experimental.
	QueueUrl() *string
	// Experimental.
	Scope() MonitoringScope
	// Experimental.
	TimeToDrainAnnotations() *[]*awscloudwatch.HorizontalAnnotation
	// Experimental.
	TimeToDrainMetric() interface{}
	// Experimental.
	Title() *string
	// Experimental.
	VisibleMessagesMetric() interface{}
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
	CreateDeadLetterMessageAgeWidget(width *float64, height *float64) awscloudwatch.GraphWidget
	// Experimental.
	CreateDeadLetterMessageCountWidget(width *float64, height *float64) awscloudwatch.GraphWidget
	// Experimental.
	CreateDeadLetterTitleWidget() MonitoringHeaderWidget
	// Experimental.
	CreateMessageAgeWidget(width *float64, height *float64) awscloudwatch.GraphWidget
	// Experimental.
	CreateMessageCountWidget(width *float64, height *float64) awscloudwatch.GraphWidget
	// Experimental.
	CreateMessageSizeWidget(width *float64, height *float64) awscloudwatch.GraphWidget
	// Creates a new metric factory.
	// Experimental.
	CreateMetricFactory() MetricFactory
	// Experimental.
	CreateProducerAndConsumerRateWidget(width *float64, height *float64) awscloudwatch.GraphWidget
	// Experimental.
	CreateTimeToDrainWidget(width *float64, height *float64) awscloudwatch.GraphWidget
	// Experimental.
	CreateTitleWidget() MonitoringHeaderWidget
	// Creates a new widget factory.
	// Experimental.
	CreateWidgetFactory() IWidgetFactory
	// Experimental.
	ResolveQueueName(queue awssqs.IQueue) *string
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

// The jsii proxy struct for SqsQueueMonitoringWithDlq
type jsiiProxy_SqsQueueMonitoringWithDlq struct {
	jsiiProxy_SqsQueueMonitoring
}

func (j *jsiiProxy_SqsQueueMonitoringWithDlq) AddDeadLetterQueueToSummaryDashboard() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"addDeadLetterQueueToSummaryDashboard",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqsQueueMonitoringWithDlq) AgeAnnotations() *[]*awscloudwatch.HorizontalAnnotation {
	var returns *[]*awscloudwatch.HorizontalAnnotation
	_jsii_.Get(
		j,
		"ageAnnotations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqsQueueMonitoringWithDlq) Alarms() *[]*AlarmWithAnnotation {
	var returns *[]*AlarmWithAnnotation
	_jsii_.Get(
		j,
		"alarms",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqsQueueMonitoringWithDlq) ConsumptionRateMetric() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"consumptionRateMetric",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqsQueueMonitoringWithDlq) CountAnnotations() *[]*awscloudwatch.HorizontalAnnotation {
	var returns *[]*awscloudwatch.HorizontalAnnotation
	_jsii_.Get(
		j,
		"countAnnotations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqsQueueMonitoringWithDlq) DeadLetterAgeAnnotations() *[]*awscloudwatch.HorizontalAnnotation {
	var returns *[]*awscloudwatch.HorizontalAnnotation
	_jsii_.Get(
		j,
		"deadLetterAgeAnnotations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqsQueueMonitoringWithDlq) DeadLetterCountAnnotations() *[]*awscloudwatch.HorizontalAnnotation {
	var returns *[]*awscloudwatch.HorizontalAnnotation
	_jsii_.Get(
		j,
		"deadLetterCountAnnotations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqsQueueMonitoringWithDlq) DeadLetterQueueAlarmFactory() QueueAlarmFactory {
	var returns QueueAlarmFactory
	_jsii_.Get(
		j,
		"deadLetterQueueAlarmFactory",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqsQueueMonitoringWithDlq) DeadLetterQueueIncomingMessagesMetric() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"deadLetterQueueIncomingMessagesMetric",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqsQueueMonitoringWithDlq) DeadLetterQueueOldestMessageAgeMetric() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"deadLetterQueueOldestMessageAgeMetric",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqsQueueMonitoringWithDlq) DeadLetterQueueVisibleMessagesMetric() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"deadLetterQueueVisibleMessagesMetric",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqsQueueMonitoringWithDlq) DeadLetterTitle() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deadLetterTitle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqsQueueMonitoringWithDlq) DeadLetterUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deadLetterUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqsQueueMonitoringWithDlq) DeletedMessagesMetric() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"deletedMessagesMetric",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqsQueueMonitoringWithDlq) IncomingMessagesMetric() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"incomingMessagesMetric",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqsQueueMonitoringWithDlq) LocalAlarmNamePrefixOverride() *string {
	var returns *string
	_jsii_.Get(
		j,
		"localAlarmNamePrefixOverride",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqsQueueMonitoringWithDlq) MessageSizeMetric() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"messageSizeMetric",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqsQueueMonitoringWithDlq) OldestMessageAgeMetric() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"oldestMessageAgeMetric",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqsQueueMonitoringWithDlq) ProductionRateMetric() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"productionRateMetric",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqsQueueMonitoringWithDlq) QueueAlarmFactory() QueueAlarmFactory {
	var returns QueueAlarmFactory
	_jsii_.Get(
		j,
		"queueAlarmFactory",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqsQueueMonitoringWithDlq) QueueUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"queueUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqsQueueMonitoringWithDlq) Scope() MonitoringScope {
	var returns MonitoringScope
	_jsii_.Get(
		j,
		"scope",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqsQueueMonitoringWithDlq) TimeToDrainAnnotations() *[]*awscloudwatch.HorizontalAnnotation {
	var returns *[]*awscloudwatch.HorizontalAnnotation
	_jsii_.Get(
		j,
		"timeToDrainAnnotations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqsQueueMonitoringWithDlq) TimeToDrainMetric() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeToDrainMetric",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqsQueueMonitoringWithDlq) Title() *string {
	var returns *string
	_jsii_.Get(
		j,
		"title",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqsQueueMonitoringWithDlq) VisibleMessagesMetric() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"visibleMessagesMetric",
		&returns,
	)
	return returns
}


// Experimental.
func NewSqsQueueMonitoringWithDlq(scope MonitoringScope, props *SqsQueueMonitoringWithDlqProps) SqsQueueMonitoringWithDlq {
	_init_.Initialize()

	if err := validateNewSqsQueueMonitoringWithDlqParameters(scope, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_SqsQueueMonitoringWithDlq{}

	_jsii_.Create(
		"cdk-monitoring-constructs.SqsQueueMonitoringWithDlq",
		[]interface{}{scope, props},
		&j,
	)

	return &j
}

// Experimental.
func NewSqsQueueMonitoringWithDlq_Override(s SqsQueueMonitoringWithDlq, scope MonitoringScope, props *SqsQueueMonitoringWithDlqProps) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk-monitoring-constructs.SqsQueueMonitoringWithDlq",
		[]interface{}{scope, props},
		s,
	)
}

func (s *jsiiProxy_SqsQueueMonitoringWithDlq) AddAlarm(alarm *AlarmWithAnnotation) {
	if err := s.validateAddAlarmParameters(alarm); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"addAlarm",
		[]interface{}{alarm},
	)
}

func (s *jsiiProxy_SqsQueueMonitoringWithDlq) AlarmWidgets() *[]awscloudwatch.IWidget {
	var returns *[]awscloudwatch.IWidget

	_jsii_.Invoke(
		s,
		"alarmWidgets",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SqsQueueMonitoringWithDlq) CreateAlarmFactory(alarmNamePrefix *string) AlarmFactory {
	if err := s.validateCreateAlarmFactoryParameters(alarmNamePrefix); err != nil {
		panic(err)
	}
	var returns AlarmFactory

	_jsii_.Invoke(
		s,
		"createAlarmFactory",
		[]interface{}{alarmNamePrefix},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SqsQueueMonitoringWithDlq) CreatedAlarms() *[]*AlarmWithAnnotation {
	var returns *[]*AlarmWithAnnotation

	_jsii_.Invoke(
		s,
		"createdAlarms",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SqsQueueMonitoringWithDlq) CreateDeadLetterMessageAgeWidget(width *float64, height *float64) awscloudwatch.GraphWidget {
	if err := s.validateCreateDeadLetterMessageAgeWidgetParameters(width, height); err != nil {
		panic(err)
	}
	var returns awscloudwatch.GraphWidget

	_jsii_.Invoke(
		s,
		"createDeadLetterMessageAgeWidget",
		[]interface{}{width, height},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SqsQueueMonitoringWithDlq) CreateDeadLetterMessageCountWidget(width *float64, height *float64) awscloudwatch.GraphWidget {
	if err := s.validateCreateDeadLetterMessageCountWidgetParameters(width, height); err != nil {
		panic(err)
	}
	var returns awscloudwatch.GraphWidget

	_jsii_.Invoke(
		s,
		"createDeadLetterMessageCountWidget",
		[]interface{}{width, height},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SqsQueueMonitoringWithDlq) CreateDeadLetterTitleWidget() MonitoringHeaderWidget {
	var returns MonitoringHeaderWidget

	_jsii_.Invoke(
		s,
		"createDeadLetterTitleWidget",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SqsQueueMonitoringWithDlq) CreateMessageAgeWidget(width *float64, height *float64) awscloudwatch.GraphWidget {
	if err := s.validateCreateMessageAgeWidgetParameters(width, height); err != nil {
		panic(err)
	}
	var returns awscloudwatch.GraphWidget

	_jsii_.Invoke(
		s,
		"createMessageAgeWidget",
		[]interface{}{width, height},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SqsQueueMonitoringWithDlq) CreateMessageCountWidget(width *float64, height *float64) awscloudwatch.GraphWidget {
	if err := s.validateCreateMessageCountWidgetParameters(width, height); err != nil {
		panic(err)
	}
	var returns awscloudwatch.GraphWidget

	_jsii_.Invoke(
		s,
		"createMessageCountWidget",
		[]interface{}{width, height},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SqsQueueMonitoringWithDlq) CreateMessageSizeWidget(width *float64, height *float64) awscloudwatch.GraphWidget {
	if err := s.validateCreateMessageSizeWidgetParameters(width, height); err != nil {
		panic(err)
	}
	var returns awscloudwatch.GraphWidget

	_jsii_.Invoke(
		s,
		"createMessageSizeWidget",
		[]interface{}{width, height},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SqsQueueMonitoringWithDlq) CreateMetricFactory() MetricFactory {
	var returns MetricFactory

	_jsii_.Invoke(
		s,
		"createMetricFactory",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SqsQueueMonitoringWithDlq) CreateProducerAndConsumerRateWidget(width *float64, height *float64) awscloudwatch.GraphWidget {
	if err := s.validateCreateProducerAndConsumerRateWidgetParameters(width, height); err != nil {
		panic(err)
	}
	var returns awscloudwatch.GraphWidget

	_jsii_.Invoke(
		s,
		"createProducerAndConsumerRateWidget",
		[]interface{}{width, height},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SqsQueueMonitoringWithDlq) CreateTimeToDrainWidget(width *float64, height *float64) awscloudwatch.GraphWidget {
	if err := s.validateCreateTimeToDrainWidgetParameters(width, height); err != nil {
		panic(err)
	}
	var returns awscloudwatch.GraphWidget

	_jsii_.Invoke(
		s,
		"createTimeToDrainWidget",
		[]interface{}{width, height},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SqsQueueMonitoringWithDlq) CreateTitleWidget() MonitoringHeaderWidget {
	var returns MonitoringHeaderWidget

	_jsii_.Invoke(
		s,
		"createTitleWidget",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SqsQueueMonitoringWithDlq) CreateWidgetFactory() IWidgetFactory {
	var returns IWidgetFactory

	_jsii_.Invoke(
		s,
		"createWidgetFactory",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SqsQueueMonitoringWithDlq) ResolveQueueName(queue awssqs.IQueue) *string {
	if err := s.validateResolveQueueNameParameters(queue); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		s,
		"resolveQueueName",
		[]interface{}{queue},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SqsQueueMonitoringWithDlq) SummaryWidgets() *[]awscloudwatch.IWidget {
	var returns *[]awscloudwatch.IWidget

	_jsii_.Invoke(
		s,
		"summaryWidgets",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SqsQueueMonitoringWithDlq) Widgets() *[]awscloudwatch.IWidget {
	var returns *[]awscloudwatch.IWidget

	_jsii_.Invoke(
		s,
		"widgets",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SqsQueueMonitoringWithDlq) WidgetsForDashboard(name *string) *[]awscloudwatch.IWidget {
	if err := s.validateWidgetsForDashboardParameters(name); err != nil {
		panic(err)
	}
	var returns *[]awscloudwatch.IWidget

	_jsii_.Invoke(
		s,
		"widgetsForDashboard",
		[]interface{}{name},
		&returns,
	)

	return returns
}


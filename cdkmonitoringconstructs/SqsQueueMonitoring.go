package cdkmonitoringconstructs

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/corymhall/cdk-constructs-go/cdkmonitoringconstructs/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudwatch"
	"github.com/aws/aws-cdk-go/awscdk/v2/awssqs"
)

// Experimental.
type SqsQueueMonitoring interface {
	Monitoring
	// Experimental.
	AgeAnnotations() *[]*awscloudwatch.HorizontalAnnotation
	// Experimental.
	Alarms() *[]*AlarmWithAnnotation
	// Experimental.
	ConsumptionRateMetric() interface{}
	// Experimental.
	CountAnnotations() *[]*awscloudwatch.HorizontalAnnotation
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

// The jsii proxy struct for SqsQueueMonitoring
type jsiiProxy_SqsQueueMonitoring struct {
	jsiiProxy_Monitoring
}

func (j *jsiiProxy_SqsQueueMonitoring) AgeAnnotations() *[]*awscloudwatch.HorizontalAnnotation {
	var returns *[]*awscloudwatch.HorizontalAnnotation
	_jsii_.Get(
		j,
		"ageAnnotations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqsQueueMonitoring) Alarms() *[]*AlarmWithAnnotation {
	var returns *[]*AlarmWithAnnotation
	_jsii_.Get(
		j,
		"alarms",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqsQueueMonitoring) ConsumptionRateMetric() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"consumptionRateMetric",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqsQueueMonitoring) CountAnnotations() *[]*awscloudwatch.HorizontalAnnotation {
	var returns *[]*awscloudwatch.HorizontalAnnotation
	_jsii_.Get(
		j,
		"countAnnotations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqsQueueMonitoring) DeletedMessagesMetric() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"deletedMessagesMetric",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqsQueueMonitoring) IncomingMessagesMetric() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"incomingMessagesMetric",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqsQueueMonitoring) LocalAlarmNamePrefixOverride() *string {
	var returns *string
	_jsii_.Get(
		j,
		"localAlarmNamePrefixOverride",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqsQueueMonitoring) MessageSizeMetric() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"messageSizeMetric",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqsQueueMonitoring) OldestMessageAgeMetric() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"oldestMessageAgeMetric",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqsQueueMonitoring) ProductionRateMetric() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"productionRateMetric",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqsQueueMonitoring) QueueAlarmFactory() QueueAlarmFactory {
	var returns QueueAlarmFactory
	_jsii_.Get(
		j,
		"queueAlarmFactory",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqsQueueMonitoring) QueueUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"queueUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqsQueueMonitoring) Scope() MonitoringScope {
	var returns MonitoringScope
	_jsii_.Get(
		j,
		"scope",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqsQueueMonitoring) TimeToDrainAnnotations() *[]*awscloudwatch.HorizontalAnnotation {
	var returns *[]*awscloudwatch.HorizontalAnnotation
	_jsii_.Get(
		j,
		"timeToDrainAnnotations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqsQueueMonitoring) TimeToDrainMetric() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeToDrainMetric",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqsQueueMonitoring) Title() *string {
	var returns *string
	_jsii_.Get(
		j,
		"title",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqsQueueMonitoring) VisibleMessagesMetric() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"visibleMessagesMetric",
		&returns,
	)
	return returns
}


// Experimental.
func NewSqsQueueMonitoring(scope MonitoringScope, props *SqsQueueMonitoringProps, invokedFromSuper *bool) SqsQueueMonitoring {
	_init_.Initialize()

	if err := validateNewSqsQueueMonitoringParameters(scope, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_SqsQueueMonitoring{}

	_jsii_.Create(
		"cdk-monitoring-constructs.SqsQueueMonitoring",
		[]interface{}{scope, props, invokedFromSuper},
		&j,
	)

	return &j
}

// Experimental.
func NewSqsQueueMonitoring_Override(s SqsQueueMonitoring, scope MonitoringScope, props *SqsQueueMonitoringProps, invokedFromSuper *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk-monitoring-constructs.SqsQueueMonitoring",
		[]interface{}{scope, props, invokedFromSuper},
		s,
	)
}

func (s *jsiiProxy_SqsQueueMonitoring) AddAlarm(alarm *AlarmWithAnnotation) {
	if err := s.validateAddAlarmParameters(alarm); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"addAlarm",
		[]interface{}{alarm},
	)
}

func (s *jsiiProxy_SqsQueueMonitoring) AlarmWidgets() *[]awscloudwatch.IWidget {
	var returns *[]awscloudwatch.IWidget

	_jsii_.Invoke(
		s,
		"alarmWidgets",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SqsQueueMonitoring) CreateAlarmFactory(alarmNamePrefix *string) AlarmFactory {
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

func (s *jsiiProxy_SqsQueueMonitoring) CreatedAlarms() *[]*AlarmWithAnnotation {
	var returns *[]*AlarmWithAnnotation

	_jsii_.Invoke(
		s,
		"createdAlarms",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SqsQueueMonitoring) CreateMessageAgeWidget(width *float64, height *float64) awscloudwatch.GraphWidget {
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

func (s *jsiiProxy_SqsQueueMonitoring) CreateMessageCountWidget(width *float64, height *float64) awscloudwatch.GraphWidget {
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

func (s *jsiiProxy_SqsQueueMonitoring) CreateMessageSizeWidget(width *float64, height *float64) awscloudwatch.GraphWidget {
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

func (s *jsiiProxy_SqsQueueMonitoring) CreateMetricFactory() MetricFactory {
	var returns MetricFactory

	_jsii_.Invoke(
		s,
		"createMetricFactory",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SqsQueueMonitoring) CreateProducerAndConsumerRateWidget(width *float64, height *float64) awscloudwatch.GraphWidget {
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

func (s *jsiiProxy_SqsQueueMonitoring) CreateTimeToDrainWidget(width *float64, height *float64) awscloudwatch.GraphWidget {
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

func (s *jsiiProxy_SqsQueueMonitoring) CreateTitleWidget() MonitoringHeaderWidget {
	var returns MonitoringHeaderWidget

	_jsii_.Invoke(
		s,
		"createTitleWidget",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SqsQueueMonitoring) CreateWidgetFactory() IWidgetFactory {
	var returns IWidgetFactory

	_jsii_.Invoke(
		s,
		"createWidgetFactory",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SqsQueueMonitoring) ResolveQueueName(queue awssqs.IQueue) *string {
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

func (s *jsiiProxy_SqsQueueMonitoring) SummaryWidgets() *[]awscloudwatch.IWidget {
	var returns *[]awscloudwatch.IWidget

	_jsii_.Invoke(
		s,
		"summaryWidgets",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SqsQueueMonitoring) Widgets() *[]awscloudwatch.IWidget {
	var returns *[]awscloudwatch.IWidget

	_jsii_.Invoke(
		s,
		"widgets",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SqsQueueMonitoring) WidgetsForDashboard(name *string) *[]awscloudwatch.IWidget {
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


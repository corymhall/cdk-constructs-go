package cdkmonitoringconstructs

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/corymhall/cdk-constructs-go/cdkmonitoringconstructs/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudwatch"
)

// Experimental.
type SnsTopicMonitoring interface {
	Monitoring
	// Experimental.
	Alarms() *[]*AlarmWithAnnotation
	// Experimental.
	FailedDeliveryAnnotations() *[]*awscloudwatch.HorizontalAnnotation
	// Experimental.
	IncomingMessagesAnnotations() *[]*awscloudwatch.HorizontalAnnotation
	// Experimental.
	IncomingMessagesMetric() interface{}
	// Experimental.
	LocalAlarmNamePrefixOverride() *string
	// Experimental.
	MessagesFailedMetric() interface{}
	// Experimental.
	MessageSizeMetric() interface{}
	// Experimental.
	OutgoingMessagesMetric() interface{}
	// Experimental.
	Scope() MonitoringScope
	// Experimental.
	Title() *string
	// Experimental.
	TopicAlarmFactory() TopicAlarmFactory
	// Experimental.
	TopicUrl() *string
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
	CreateMessageCountWidget(width *float64, height *float64) awscloudwatch.GraphWidget
	// Experimental.
	CreateMessageFailedWidget(width *float64, height *float64) awscloudwatch.GraphWidget
	// Experimental.
	CreateMessageSizeWidget(width *float64, height *float64) awscloudwatch.GraphWidget
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

// The jsii proxy struct for SnsTopicMonitoring
type jsiiProxy_SnsTopicMonitoring struct {
	jsiiProxy_Monitoring
}

func (j *jsiiProxy_SnsTopicMonitoring) Alarms() *[]*AlarmWithAnnotation {
	var returns *[]*AlarmWithAnnotation
	_jsii_.Get(
		j,
		"alarms",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnsTopicMonitoring) FailedDeliveryAnnotations() *[]*awscloudwatch.HorizontalAnnotation {
	var returns *[]*awscloudwatch.HorizontalAnnotation
	_jsii_.Get(
		j,
		"failedDeliveryAnnotations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnsTopicMonitoring) IncomingMessagesAnnotations() *[]*awscloudwatch.HorizontalAnnotation {
	var returns *[]*awscloudwatch.HorizontalAnnotation
	_jsii_.Get(
		j,
		"incomingMessagesAnnotations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnsTopicMonitoring) IncomingMessagesMetric() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"incomingMessagesMetric",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnsTopicMonitoring) LocalAlarmNamePrefixOverride() *string {
	var returns *string
	_jsii_.Get(
		j,
		"localAlarmNamePrefixOverride",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnsTopicMonitoring) MessagesFailedMetric() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"messagesFailedMetric",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnsTopicMonitoring) MessageSizeMetric() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"messageSizeMetric",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnsTopicMonitoring) OutgoingMessagesMetric() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"outgoingMessagesMetric",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnsTopicMonitoring) Scope() MonitoringScope {
	var returns MonitoringScope
	_jsii_.Get(
		j,
		"scope",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnsTopicMonitoring) Title() *string {
	var returns *string
	_jsii_.Get(
		j,
		"title",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnsTopicMonitoring) TopicAlarmFactory() TopicAlarmFactory {
	var returns TopicAlarmFactory
	_jsii_.Get(
		j,
		"topicAlarmFactory",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnsTopicMonitoring) TopicUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"topicUrl",
		&returns,
	)
	return returns
}


// Experimental.
func NewSnsTopicMonitoring(scope MonitoringScope, props *SnsTopicMonitoringProps) SnsTopicMonitoring {
	_init_.Initialize()

	if err := validateNewSnsTopicMonitoringParameters(scope, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_SnsTopicMonitoring{}

	_jsii_.Create(
		"cdk-monitoring-constructs.SnsTopicMonitoring",
		[]interface{}{scope, props},
		&j,
	)

	return &j
}

// Experimental.
func NewSnsTopicMonitoring_Override(s SnsTopicMonitoring, scope MonitoringScope, props *SnsTopicMonitoringProps) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk-monitoring-constructs.SnsTopicMonitoring",
		[]interface{}{scope, props},
		s,
	)
}

func (s *jsiiProxy_SnsTopicMonitoring) AddAlarm(alarm *AlarmWithAnnotation) {
	if err := s.validateAddAlarmParameters(alarm); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"addAlarm",
		[]interface{}{alarm},
	)
}

func (s *jsiiProxy_SnsTopicMonitoring) AlarmWidgets() *[]awscloudwatch.IWidget {
	var returns *[]awscloudwatch.IWidget

	_jsii_.Invoke(
		s,
		"alarmWidgets",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SnsTopicMonitoring) CreateAlarmFactory(alarmNamePrefix *string) AlarmFactory {
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

func (s *jsiiProxy_SnsTopicMonitoring) CreatedAlarms() *[]*AlarmWithAnnotation {
	var returns *[]*AlarmWithAnnotation

	_jsii_.Invoke(
		s,
		"createdAlarms",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SnsTopicMonitoring) CreateMessageCountWidget(width *float64, height *float64) awscloudwatch.GraphWidget {
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

func (s *jsiiProxy_SnsTopicMonitoring) CreateMessageFailedWidget(width *float64, height *float64) awscloudwatch.GraphWidget {
	if err := s.validateCreateMessageFailedWidgetParameters(width, height); err != nil {
		panic(err)
	}
	var returns awscloudwatch.GraphWidget

	_jsii_.Invoke(
		s,
		"createMessageFailedWidget",
		[]interface{}{width, height},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SnsTopicMonitoring) CreateMessageSizeWidget(width *float64, height *float64) awscloudwatch.GraphWidget {
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

func (s *jsiiProxy_SnsTopicMonitoring) CreateMetricFactory() MetricFactory {
	var returns MetricFactory

	_jsii_.Invoke(
		s,
		"createMetricFactory",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SnsTopicMonitoring) CreateTitleWidget() MonitoringHeaderWidget {
	var returns MonitoringHeaderWidget

	_jsii_.Invoke(
		s,
		"createTitleWidget",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SnsTopicMonitoring) CreateWidgetFactory() IWidgetFactory {
	var returns IWidgetFactory

	_jsii_.Invoke(
		s,
		"createWidgetFactory",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SnsTopicMonitoring) SummaryWidgets() *[]awscloudwatch.IWidget {
	var returns *[]awscloudwatch.IWidget

	_jsii_.Invoke(
		s,
		"summaryWidgets",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SnsTopicMonitoring) Widgets() *[]awscloudwatch.IWidget {
	var returns *[]awscloudwatch.IWidget

	_jsii_.Invoke(
		s,
		"widgets",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SnsTopicMonitoring) WidgetsForDashboard(name *string) *[]awscloudwatch.IWidget {
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


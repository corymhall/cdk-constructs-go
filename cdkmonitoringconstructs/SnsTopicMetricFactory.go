package cdkmonitoringconstructs

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/corymhall/cdk-constructs-go/cdkmonitoringconstructs/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2/awssns"
)

// Experimental.
type SnsTopicMetricFactory interface {
	// Experimental.
	MetricFactory() MetricFactory
	// Experimental.
	Topic() awssns.ITopic
	// Experimental.
	MetricAverageMessageSizeInBytes() interface{}
	// Experimental.
	MetricIncomingMessageCount() interface{}
	// Experimental.
	MetricNumberOfNotificationsFailed() interface{}
	// Experimental.
	MetricOutgoingMessageCount() interface{}
}

// The jsii proxy struct for SnsTopicMetricFactory
type jsiiProxy_SnsTopicMetricFactory struct {
	_ byte // padding
}

func (j *jsiiProxy_SnsTopicMetricFactory) MetricFactory() MetricFactory {
	var returns MetricFactory
	_jsii_.Get(
		j,
		"metricFactory",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnsTopicMetricFactory) Topic() awssns.ITopic {
	var returns awssns.ITopic
	_jsii_.Get(
		j,
		"topic",
		&returns,
	)
	return returns
}


// Experimental.
func NewSnsTopicMetricFactory(metricFactory MetricFactory, props *SnsTopicMetricFactoryProps) SnsTopicMetricFactory {
	_init_.Initialize()

	if err := validateNewSnsTopicMetricFactoryParameters(metricFactory, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_SnsTopicMetricFactory{}

	_jsii_.Create(
		"cdk-monitoring-constructs.SnsTopicMetricFactory",
		[]interface{}{metricFactory, props},
		&j,
	)

	return &j
}

// Experimental.
func NewSnsTopicMetricFactory_Override(s SnsTopicMetricFactory, metricFactory MetricFactory, props *SnsTopicMetricFactoryProps) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk-monitoring-constructs.SnsTopicMetricFactory",
		[]interface{}{metricFactory, props},
		s,
	)
}

func (s *jsiiProxy_SnsTopicMetricFactory) MetricAverageMessageSizeInBytes() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"metricAverageMessageSizeInBytes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SnsTopicMetricFactory) MetricIncomingMessageCount() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"metricIncomingMessageCount",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SnsTopicMetricFactory) MetricNumberOfNotificationsFailed() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"metricNumberOfNotificationsFailed",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SnsTopicMetricFactory) MetricOutgoingMessageCount() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"metricOutgoingMessageCount",
		nil, // no parameters
		&returns,
	)

	return returns
}


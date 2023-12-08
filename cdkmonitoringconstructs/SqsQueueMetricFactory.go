package cdkmonitoringconstructs

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/corymhall/cdk-constructs-go/cdkmonitoringconstructs/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2/awssqs"
)

// Experimental.
type SqsQueueMetricFactory interface {
	// Experimental.
	MetricFactory() MetricFactory
	// Experimental.
	Queue() awssqs.IQueue
	// Experimental.
	MetricApproximateAgeOfOldestMessageInSeconds() interface{}
	// Experimental.
	MetricApproximateVisibleMessageCount() interface{}
	// Experimental.
	MetricAverageMessageSizeInBytes() interface{}
	// Experimental.
	MetricConsumptionRate() interface{}
	// Experimental.
	MetricDeletedMessageCount() interface{}
	// Experimental.
	MetricIncomingMessageCount() interface{}
	// Experimental.
	MetricProductionRate() interface{}
	// Experimental.
	MetricTimeToDrain() interface{}
}

// The jsii proxy struct for SqsQueueMetricFactory
type jsiiProxy_SqsQueueMetricFactory struct {
	_ byte // padding
}

func (j *jsiiProxy_SqsQueueMetricFactory) MetricFactory() MetricFactory {
	var returns MetricFactory
	_jsii_.Get(
		j,
		"metricFactory",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SqsQueueMetricFactory) Queue() awssqs.IQueue {
	var returns awssqs.IQueue
	_jsii_.Get(
		j,
		"queue",
		&returns,
	)
	return returns
}


// Experimental.
func NewSqsQueueMetricFactory(metricFactory MetricFactory, props *SqsQueueMetricFactoryProps) SqsQueueMetricFactory {
	_init_.Initialize()

	if err := validateNewSqsQueueMetricFactoryParameters(metricFactory, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_SqsQueueMetricFactory{}

	_jsii_.Create(
		"cdk-monitoring-constructs.SqsQueueMetricFactory",
		[]interface{}{metricFactory, props},
		&j,
	)

	return &j
}

// Experimental.
func NewSqsQueueMetricFactory_Override(s SqsQueueMetricFactory, metricFactory MetricFactory, props *SqsQueueMetricFactoryProps) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk-monitoring-constructs.SqsQueueMetricFactory",
		[]interface{}{metricFactory, props},
		s,
	)
}

func (s *jsiiProxy_SqsQueueMetricFactory) MetricApproximateAgeOfOldestMessageInSeconds() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"metricApproximateAgeOfOldestMessageInSeconds",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SqsQueueMetricFactory) MetricApproximateVisibleMessageCount() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"metricApproximateVisibleMessageCount",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SqsQueueMetricFactory) MetricAverageMessageSizeInBytes() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"metricAverageMessageSizeInBytes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SqsQueueMetricFactory) MetricConsumptionRate() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"metricConsumptionRate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SqsQueueMetricFactory) MetricDeletedMessageCount() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"metricDeletedMessageCount",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SqsQueueMetricFactory) MetricIncomingMessageCount() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"metricIncomingMessageCount",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SqsQueueMetricFactory) MetricProductionRate() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"metricProductionRate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SqsQueueMetricFactory) MetricTimeToDrain() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"metricTimeToDrain",
		nil, // no parameters
		&returns,
	)

	return returns
}


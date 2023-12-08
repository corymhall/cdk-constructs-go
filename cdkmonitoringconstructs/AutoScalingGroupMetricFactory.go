package cdkmonitoringconstructs

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/corymhall/cdk-constructs-go/cdkmonitoringconstructs/jsii"
)

// Experimental.
type AutoScalingGroupMetricFactory interface {
	// Experimental.
	DimensionsMap() *map[string]*string
	// Experimental.
	MetricFactory() MetricFactory
	// The number of instances that the Auto Scaling group attempts to maintain.
	// Experimental.
	MetricGroupDesiredCapacity() interface{}
	// The number of instances that are running as part of the Auto Scaling group.
	//
	// This metric does not include instances that are pending or terminating.
	// Experimental.
	MetricGroupInServiceInstances() interface{}
	// The maximum size of the Auto Scaling group.
	// Experimental.
	MetricGroupMaxSize() interface{}
	// The minimum size of the Auto Scaling group.
	// Experimental.
	MetricGroupMinSize() interface{}
	// The number of instances that are pending.
	//
	// A pending instance is not yet in service.
	// This metric does not include instances that are in service or terminating.
	// Experimental.
	MetricGroupPendingInstances() interface{}
	// The number of instances that are in a Standby state.
	//
	// Instances in this state are still running but are not actively in service.
	// Experimental.
	MetricGroupStandbyInstances() interface{}
	// The number of instances that are in the process of terminating.
	//
	// This metric does not include instances that are in service or pending.
	// Experimental.
	MetricGroupTerminatingInstances() interface{}
	// The total number of instances in the Auto Scaling group.
	//
	// This metric identifies the number of instances that are in service, pending, and terminating.
	// Experimental.
	MetricGroupTotalInstances() interface{}
}

// The jsii proxy struct for AutoScalingGroupMetricFactory
type jsiiProxy_AutoScalingGroupMetricFactory struct {
	_ byte // padding
}

func (j *jsiiProxy_AutoScalingGroupMetricFactory) DimensionsMap() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"dimensionsMap",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutoScalingGroupMetricFactory) MetricFactory() MetricFactory {
	var returns MetricFactory
	_jsii_.Get(
		j,
		"metricFactory",
		&returns,
	)
	return returns
}


// Experimental.
func NewAutoScalingGroupMetricFactory(metricFactory MetricFactory, props *AutoScalingGroupMetricFactoryProps) AutoScalingGroupMetricFactory {
	_init_.Initialize()

	if err := validateNewAutoScalingGroupMetricFactoryParameters(metricFactory, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_AutoScalingGroupMetricFactory{}

	_jsii_.Create(
		"cdk-monitoring-constructs.AutoScalingGroupMetricFactory",
		[]interface{}{metricFactory, props},
		&j,
	)

	return &j
}

// Experimental.
func NewAutoScalingGroupMetricFactory_Override(a AutoScalingGroupMetricFactory, metricFactory MetricFactory, props *AutoScalingGroupMetricFactoryProps) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk-monitoring-constructs.AutoScalingGroupMetricFactory",
		[]interface{}{metricFactory, props},
		a,
	)
}

func (a *jsiiProxy_AutoScalingGroupMetricFactory) MetricGroupDesiredCapacity() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"metricGroupDesiredCapacity",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutoScalingGroupMetricFactory) MetricGroupInServiceInstances() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"metricGroupInServiceInstances",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutoScalingGroupMetricFactory) MetricGroupMaxSize() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"metricGroupMaxSize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutoScalingGroupMetricFactory) MetricGroupMinSize() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"metricGroupMinSize",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutoScalingGroupMetricFactory) MetricGroupPendingInstances() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"metricGroupPendingInstances",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutoScalingGroupMetricFactory) MetricGroupStandbyInstances() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"metricGroupStandbyInstances",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutoScalingGroupMetricFactory) MetricGroupTerminatingInstances() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"metricGroupTerminatingInstances",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutoScalingGroupMetricFactory) MetricGroupTotalInstances() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"metricGroupTotalInstances",
		nil, // no parameters
		&returns,
	)

	return returns
}


package cdkmonitoringconstructs

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/corymhall/cdk-constructs-go/cdkmonitoringconstructs/jsii"

	"github.com/aws/constructs-go/constructs/v10"
)

// Allows to apply a collection of {@link IMetricAdjuster} to a metric.
// Experimental.
type CompositeMetricAdjuster interface {
	IMetricAdjuster
	// Adjusts a metric.
	// Experimental.
	AdjustMetric(metric interface{}, alarmScope constructs.Construct, props *AddAlarmProps) interface{}
}

// The jsii proxy struct for CompositeMetricAdjuster
type jsiiProxy_CompositeMetricAdjuster struct {
	jsiiProxy_IMetricAdjuster
}

// Experimental.
func NewCompositeMetricAdjuster(adjusters *[]IMetricAdjuster) CompositeMetricAdjuster {
	_init_.Initialize()

	if err := validateNewCompositeMetricAdjusterParameters(adjusters); err != nil {
		panic(err)
	}
	j := jsiiProxy_CompositeMetricAdjuster{}

	_jsii_.Create(
		"cdk-monitoring-constructs.CompositeMetricAdjuster",
		[]interface{}{adjusters},
		&j,
	)

	return &j
}

// Experimental.
func NewCompositeMetricAdjuster_Override(c CompositeMetricAdjuster, adjusters *[]IMetricAdjuster) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk-monitoring-constructs.CompositeMetricAdjuster",
		[]interface{}{adjusters},
		c,
	)
}

// Experimental.
func CompositeMetricAdjuster_Of(adjusters ...IMetricAdjuster) CompositeMetricAdjuster {
	_init_.Initialize()

	args := []interface{}{}
	for _, a := range adjusters {
		args = append(args, a)
	}

	var returns CompositeMetricAdjuster

	_jsii_.StaticInvoke(
		"cdk-monitoring-constructs.CompositeMetricAdjuster",
		"of",
		args,
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CompositeMetricAdjuster) AdjustMetric(metric interface{}, alarmScope constructs.Construct, props *AddAlarmProps) interface{} {
	if err := c.validateAdjustMetricParameters(metric, alarmScope, props); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		c,
		"adjustMetric",
		[]interface{}{metric, alarmScope, props},
		&returns,
	)

	return returns
}


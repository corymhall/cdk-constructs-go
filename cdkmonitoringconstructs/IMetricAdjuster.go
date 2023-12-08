package cdkmonitoringconstructs

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/constructs-go/constructs/v10"
)

// Adjusts a metric before creating adding an alarm to it.
// Experimental.
type IMetricAdjuster interface {
	// Adjusts a metric.
	//
	// Returns: The adjusted metric.
	// Experimental.
	AdjustMetric(metric interface{}, alarmScope constructs.Construct, props *AddAlarmProps) interface{}
}

// The jsii proxy for IMetricAdjuster
type jsiiProxy_IMetricAdjuster struct {
	_ byte // padding
}

func (i *jsiiProxy_IMetricAdjuster) AdjustMetric(metric interface{}, alarmScope constructs.Construct, props *AddAlarmProps) interface{} {
	if err := i.validateAdjustMetricParameters(metric, alarmScope, props); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		i,
		"adjustMetric",
		[]interface{}{metric, alarmScope, props},
		&returns,
	)

	return returns
}


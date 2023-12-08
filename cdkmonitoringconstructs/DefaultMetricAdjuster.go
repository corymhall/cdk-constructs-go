package cdkmonitoringconstructs

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/corymhall/cdk-constructs-go/cdkmonitoringconstructs/jsii"

	"github.com/aws/constructs-go/constructs/v10"
)

// Applies the default metric adjustments.
//
// These adjustments are always applied last, regardless the value configured in {@link AddAlarmProps.metricAdjuster}.
// Experimental.
type DefaultMetricAdjuster interface {
	IMetricAdjuster
	// Adjusts a metric.
	// Experimental.
	AdjustMetric(metric interface{}, _arg constructs.Construct, props *AddAlarmProps) interface{}
}

// The jsii proxy struct for DefaultMetricAdjuster
type jsiiProxy_DefaultMetricAdjuster struct {
	jsiiProxy_IMetricAdjuster
}

// Experimental.
func NewDefaultMetricAdjuster() DefaultMetricAdjuster {
	_init_.Initialize()

	j := jsiiProxy_DefaultMetricAdjuster{}

	_jsii_.Create(
		"cdk-monitoring-constructs.DefaultMetricAdjuster",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewDefaultMetricAdjuster_Override(d DefaultMetricAdjuster) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk-monitoring-constructs.DefaultMetricAdjuster",
		nil, // no parameters
		d,
	)
}

func DefaultMetricAdjuster_INSTANCE() DefaultMetricAdjuster {
	_init_.Initialize()
	var returns DefaultMetricAdjuster
	_jsii_.StaticGet(
		"cdk-monitoring-constructs.DefaultMetricAdjuster",
		"INSTANCE",
		&returns,
	)
	return returns
}

func (d *jsiiProxy_DefaultMetricAdjuster) AdjustMetric(metric interface{}, _arg constructs.Construct, props *AddAlarmProps) interface{} {
	if err := d.validateAdjustMetricParameters(metric, _arg, props); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		d,
		"adjustMetric",
		[]interface{}{metric, _arg, props},
		&returns,
	)

	return returns
}


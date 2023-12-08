package cdkmonitoringconstructs

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/corymhall/cdk-constructs-go/cdkmonitoringconstructs/jsii"

	"github.com/aws/constructs-go/constructs/v10"
)

// Adjusts a metric so that alarms created from it can be used in Route53 Health Checks.
//
// The metric will be validated to ensure it satisfies Route53 Health Check alarm requirements, otherwise it will throw an {@link Error}.
// See: https://docs.aws.amazon.com/Route53/latest/DeveloperGuide/health-checks-types.html
//
// Experimental.
type Route53HealthCheckMetricAdjuster interface {
	IMetricAdjuster
	// Adjusts a metric.
	// Experimental.
	AdjustMetric(metric interface{}, alarmScope constructs.Construct, props *AddAlarmProps) interface{}
}

// The jsii proxy struct for Route53HealthCheckMetricAdjuster
type jsiiProxy_Route53HealthCheckMetricAdjuster struct {
	jsiiProxy_IMetricAdjuster
}

// Experimental.
func NewRoute53HealthCheckMetricAdjuster() Route53HealthCheckMetricAdjuster {
	_init_.Initialize()

	j := jsiiProxy_Route53HealthCheckMetricAdjuster{}

	_jsii_.Create(
		"cdk-monitoring-constructs.Route53HealthCheckMetricAdjuster",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewRoute53HealthCheckMetricAdjuster_Override(r Route53HealthCheckMetricAdjuster) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk-monitoring-constructs.Route53HealthCheckMetricAdjuster",
		nil, // no parameters
		r,
	)
}

func Route53HealthCheckMetricAdjuster_INSTANCE() Route53HealthCheckMetricAdjuster {
	_init_.Initialize()
	var returns Route53HealthCheckMetricAdjuster
	_jsii_.StaticGet(
		"cdk-monitoring-constructs.Route53HealthCheckMetricAdjuster",
		"INSTANCE",
		&returns,
	)
	return returns
}

func (r *jsiiProxy_Route53HealthCheckMetricAdjuster) AdjustMetric(metric interface{}, alarmScope constructs.Construct, props *AddAlarmProps) interface{} {
	if err := r.validateAdjustMetricParameters(metric, alarmScope, props); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		r,
		"adjustMetric",
		[]interface{}{metric, alarmScope, props},
		&returns,
	)

	return returns
}


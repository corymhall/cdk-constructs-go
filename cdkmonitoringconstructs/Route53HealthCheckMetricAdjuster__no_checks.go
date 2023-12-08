//go:build no_runtime_type_checking

package cdkmonitoringconstructs

// Building without runtime type checking enabled, so all the below just return nil

func (r *jsiiProxy_Route53HealthCheckMetricAdjuster) validateAdjustMetricParameters(metric interface{}, alarmScope constructs.Construct, props *AddAlarmProps) error {
	return nil
}


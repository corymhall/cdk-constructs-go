//go:build no_runtime_type_checking

package cdkmonitoringconstructs

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CompositeMetricAdjuster) validateAdjustMetricParameters(metric interface{}, alarmScope constructs.Construct, props *AddAlarmProps) error {
	return nil
}

func validateNewCompositeMetricAdjusterParameters(adjusters *[]IMetricAdjuster) error {
	return nil
}


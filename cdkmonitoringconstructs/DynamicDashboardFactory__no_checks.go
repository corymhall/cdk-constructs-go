//go:build no_runtime_type_checking

package cdkmonitoringconstructs

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DynamicDashboardFactory) validateAddDynamicSegmentParameters(segment IDynamicDashboardSegment) error {
	return nil
}

func (d *jsiiProxy_DynamicDashboardFactory) validateCreateDashboardParameters(renderingPreference DashboardRenderingPreference, id *string, props *awscloudwatch.DashboardProps) error {
	return nil
}

func (d *jsiiProxy_DynamicDashboardFactory) validateGetDashboardParameters(type_ *string) error {
	return nil
}

func validateDynamicDashboardFactory_IsConstructParameters(x interface{}) error {
	return nil
}

func validateNewDynamicDashboardFactoryParameters(scope constructs.Construct, id *string, props *MonitoringDynamicDashboardsProps) error {
	return nil
}


//go:build no_runtime_type_checking

package cdkmonitoringconstructs

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DefaultDashboardFactory) validateAddDynamicSegmentParameters(segment IDynamicDashboardSegment) error {
	return nil
}

func (d *jsiiProxy_DefaultDashboardFactory) validateAddSegmentParameters(props IDashboardFactoryProps) error {
	return nil
}

func (d *jsiiProxy_DefaultDashboardFactory) validateCreateDashboardParameters(renderingPreference DashboardRenderingPreference, id *string, props *awscloudwatch.DashboardProps) error {
	return nil
}

func (d *jsiiProxy_DefaultDashboardFactory) validateGetDashboardParameters(name *string) error {
	return nil
}

func validateDefaultDashboardFactory_IsConstructParameters(x interface{}) error {
	return nil
}

func validateNewDefaultDashboardFactoryParameters(scope constructs.Construct, id *string, props *MonitoringDashboardsProps) error {
	return nil
}


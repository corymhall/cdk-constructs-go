//go:build no_runtime_type_checking

package cdkmonitoringconstructs

// Building without runtime type checking enabled, so all the below just return nil

func (n *jsiiProxy_NetworkLoadBalancerMonitoring) validateAddAlarmParameters(alarm *AlarmWithAnnotation) error {
	return nil
}

func (n *jsiiProxy_NetworkLoadBalancerMonitoring) validateCreateAlarmFactoryParameters(alarmNamePrefix *string) error {
	return nil
}

func (n *jsiiProxy_NetworkLoadBalancerMonitoring) validateCreateTaskHealthWidgetParameters(width *float64, height *float64) error {
	return nil
}

func (n *jsiiProxy_NetworkLoadBalancerMonitoring) validateCreateTcpFlowsWidgetParameters(width *float64, height *float64) error {
	return nil
}

func (n *jsiiProxy_NetworkLoadBalancerMonitoring) validateWidgetsForDashboardParameters(name *string) error {
	return nil
}

func validateNewNetworkLoadBalancerMonitoringParameters(scope MonitoringScope, props *NetworkLoadBalancerMonitoringProps) error {
	return nil
}


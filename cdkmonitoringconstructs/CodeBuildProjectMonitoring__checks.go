//go:build !no_runtime_type_checking

package cdkmonitoringconstructs

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awscodebuild"
)

func (c *jsiiProxy_CodeBuildProjectMonitoring) validateAddAlarmParameters(alarm *AlarmWithAnnotation) error {
	if alarm == nil {
		return fmt.Errorf("parameter alarm is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(alarm, func() string { return "parameter alarm" }); err != nil {
		return err
	}

	return nil
}

func (c *jsiiProxy_CodeBuildProjectMonitoring) validateCreateAlarmFactoryParameters(alarmNamePrefix *string) error {
	if alarmNamePrefix == nil {
		return fmt.Errorf("parameter alarmNamePrefix is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_CodeBuildProjectMonitoring) validateResolveProjectNameParameters(project awscodebuild.IProject) error {
	if project == nil {
		return fmt.Errorf("parameter project is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_CodeBuildProjectMonitoring) validateWidgetsForDashboardParameters(name *string) error {
	if name == nil {
		return fmt.Errorf("parameter name is required, but nil was provided")
	}

	return nil
}

func validateNewCodeBuildProjectMonitoringParameters(scope MonitoringScope, props *CodeBuildProjectMonitoringProps) error {
	if scope == nil {
		return fmt.Errorf("parameter scope is required, but nil was provided")
	}

	if props == nil {
		return fmt.Errorf("parameter props is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}


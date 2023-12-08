//go:build !no_runtime_type_checking

package cdkmonitoringconstructs

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudwatch"
	"github.com/aws/constructs-go/constructs/v10"
)

func (d *jsiiProxy_DynamicDashboardFactory) validateAddDynamicSegmentParameters(segment IDynamicDashboardSegment) error {
	if segment == nil {
		return fmt.Errorf("parameter segment is required, but nil was provided")
	}

	return nil
}

func (d *jsiiProxy_DynamicDashboardFactory) validateCreateDashboardParameters(renderingPreference DashboardRenderingPreference, id *string, props *awscloudwatch.DashboardProps) error {
	if renderingPreference == "" {
		return fmt.Errorf("parameter renderingPreference is required, but nil was provided")
	}

	if id == nil {
		return fmt.Errorf("parameter id is required, but nil was provided")
	}

	if props == nil {
		return fmt.Errorf("parameter props is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}

func (d *jsiiProxy_DynamicDashboardFactory) validateGetDashboardParameters(type_ *string) error {
	if type_ == nil {
		return fmt.Errorf("parameter type_ is required, but nil was provided")
	}

	return nil
}

func validateDynamicDashboardFactory_IsConstructParameters(x interface{}) error {
	if x == nil {
		return fmt.Errorf("parameter x is required, but nil was provided")
	}

	return nil
}

func validateNewDynamicDashboardFactoryParameters(scope constructs.Construct, id *string, props *MonitoringDynamicDashboardsProps) error {
	if scope == nil {
		return fmt.Errorf("parameter scope is required, but nil was provided")
	}

	if id == nil {
		return fmt.Errorf("parameter id is required, but nil was provided")
	}

	if props == nil {
		return fmt.Errorf("parameter props is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}


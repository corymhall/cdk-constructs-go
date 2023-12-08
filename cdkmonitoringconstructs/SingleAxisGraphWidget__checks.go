//go:build !no_runtime_type_checking

package cdkmonitoringconstructs

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudwatch"
)

func (s *jsiiProxy_SingleAxisGraphWidget) validateAddLeftMetricParameters(metric awscloudwatch.IMetric) error {
	if metric == nil {
		return fmt.Errorf("parameter metric is required, but nil was provided")
	}

	return nil
}

func (s *jsiiProxy_SingleAxisGraphWidget) validateAddRightMetricParameters(metric awscloudwatch.IMetric) error {
	if metric == nil {
		return fmt.Errorf("parameter metric is required, but nil was provided")
	}

	return nil
}

func (s *jsiiProxy_SingleAxisGraphWidget) validatePositionParameters(x *float64, y *float64) error {
	if x == nil {
		return fmt.Errorf("parameter x is required, but nil was provided")
	}

	if y == nil {
		return fmt.Errorf("parameter y is required, but nil was provided")
	}

	return nil
}

func validateNewSingleAxisGraphWidgetParameters(props *SingleAxisGraphWidgetProps) error {
	if props == nil {
		return fmt.Errorf("parameter props is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}


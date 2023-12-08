//go:build !no_runtime_type_checking

package cdkmonitoringconstructs

import (
	"fmt"

	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudwatch"
)

func (s *jsiiProxy_SingleWidgetDashboardSegment) validateWidgetsForDashboardParameters(name *string) error {
	if name == nil {
		return fmt.Errorf("parameter name is required, but nil was provided")
	}

	return nil
}

func validateNewSingleWidgetDashboardSegmentParameters(widget awscloudwatch.IWidget) error {
	if widget == nil {
		return fmt.Errorf("parameter widget is required, but nil was provided")
	}

	return nil
}


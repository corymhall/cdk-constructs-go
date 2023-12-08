//go:build !no_runtime_type_checking

package cdkmonitoringconstructs

import (
	"fmt"

	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudwatch"
	"github.com/aws/constructs-go/constructs/v10"
)

func (b *jsiiProxy_BitmapWidgetRenderingSupport) validateAsBitmapParameters(widget awscloudwatch.IWidget) error {
	if widget == nil {
		return fmt.Errorf("parameter widget is required, but nil was provided")
	}

	return nil
}

func (b *jsiiProxy_BitmapWidgetRenderingSupport) validateGetWidgetPropertiesParameters(widget awscloudwatch.IWidget) error {
	if widget == nil {
		return fmt.Errorf("parameter widget is required, but nil was provided")
	}

	return nil
}

func validateBitmapWidgetRenderingSupport_IsConstructParameters(x interface{}) error {
	if x == nil {
		return fmt.Errorf("parameter x is required, but nil was provided")
	}

	return nil
}

func validateNewBitmapWidgetRenderingSupportParameters(scope constructs.Construct, id *string) error {
	if scope == nil {
		return fmt.Errorf("parameter scope is required, but nil was provided")
	}

	if id == nil {
		return fmt.Errorf("parameter id is required, but nil was provided")
	}

	return nil
}


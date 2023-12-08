package cdkmonitoringconstructs

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudwatch"
)

// Experimental.
type IDynamicDashboardSegment interface {
	// Returns widgets for the requested dashboard type.
	// Experimental.
	WidgetsForDashboard(name *string) *[]awscloudwatch.IWidget
}

// The jsii proxy for IDynamicDashboardSegment
type jsiiProxy_IDynamicDashboardSegment struct {
	_ byte // padding
}

func (i *jsiiProxy_IDynamicDashboardSegment) WidgetsForDashboard(name *string) *[]awscloudwatch.IWidget {
	if err := i.validateWidgetsForDashboardParameters(name); err != nil {
		panic(err)
	}
	var returns *[]awscloudwatch.IWidget

	_jsii_.Invoke(
		i,
		"widgetsForDashboard",
		[]interface{}{name},
		&returns,
	)

	return returns
}


package cdkmonitoringconstructs

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudwatch"
)

// This dashboard factory interface provides for dynamic dashboard generation through IDynamicDashboard segments which will return different content depending on the dashboard type.
// Experimental.
type IDynamicDashboardFactory interface {
	// Adds a dynamic dashboard segment.
	// Experimental.
	AddDynamicSegment(segment IDynamicDashboardSegment)
	// Gets the dashboard for the requested dashboard type.
	// Experimental.
	GetDashboard(type_ *string) awscloudwatch.Dashboard
}

// The jsii proxy for IDynamicDashboardFactory
type jsiiProxy_IDynamicDashboardFactory struct {
	_ byte // padding
}

func (i *jsiiProxy_IDynamicDashboardFactory) AddDynamicSegment(segment IDynamicDashboardSegment) {
	if err := i.validateAddDynamicSegmentParameters(segment); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"addDynamicSegment",
		[]interface{}{segment},
	)
}

func (i *jsiiProxy_IDynamicDashboardFactory) GetDashboard(type_ *string) awscloudwatch.Dashboard {
	if err := i.validateGetDashboardParameters(type_); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Dashboard

	_jsii_.Invoke(
		i,
		"getDashboard",
		[]interface{}{type_},
		&returns,
	)

	return returns
}


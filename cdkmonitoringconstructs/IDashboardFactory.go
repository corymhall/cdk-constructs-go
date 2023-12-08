package cdkmonitoringconstructs

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudwatch"
)

// Experimental.
type IDashboardFactory interface {
	// Experimental.
	AddSegment(props IDashboardFactoryProps)
	// Experimental.
	CreatedAlarmDashboard() awscloudwatch.Dashboard
	// Experimental.
	CreatedDashboard() awscloudwatch.Dashboard
	// Experimental.
	CreatedSummaryDashboard() awscloudwatch.Dashboard
}

// The jsii proxy for IDashboardFactory
type jsiiProxy_IDashboardFactory struct {
	_ byte // padding
}

func (i *jsiiProxy_IDashboardFactory) AddSegment(props IDashboardFactoryProps) {
	if err := i.validateAddSegmentParameters(props); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"addSegment",
		[]interface{}{props},
	)
}

func (i *jsiiProxy_IDashboardFactory) CreatedAlarmDashboard() awscloudwatch.Dashboard {
	var returns awscloudwatch.Dashboard

	_jsii_.Invoke(
		i,
		"createdAlarmDashboard",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IDashboardFactory) CreatedDashboard() awscloudwatch.Dashboard {
	var returns awscloudwatch.Dashboard

	_jsii_.Invoke(
		i,
		"createdDashboard",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IDashboardFactory) CreatedSummaryDashboard() awscloudwatch.Dashboard {
	var returns awscloudwatch.Dashboard

	_jsii_.Invoke(
		i,
		"createdSummaryDashboard",
		nil, // no parameters
		&returns,
	)

	return returns
}


package cdkmonitoringconstructs

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudwatch"
)

// Experimental.
type IDashboardSegment interface {
	// Returns widgets for all alarms.
	//
	// These should go to the runbook or service dashboard.
	// Experimental.
	AlarmWidgets() *[]awscloudwatch.IWidget
	// Returns widgets for the summary.
	//
	// These should go to the team OPS dashboard.
	// Experimental.
	SummaryWidgets() *[]awscloudwatch.IWidget
	// Returns all widgets.
	//
	// These should go to the detailed service dashboard.
	// Experimental.
	Widgets() *[]awscloudwatch.IWidget
}

// The jsii proxy for IDashboardSegment
type jsiiProxy_IDashboardSegment struct {
	_ byte // padding
}

func (i *jsiiProxy_IDashboardSegment) AlarmWidgets() *[]awscloudwatch.IWidget {
	var returns *[]awscloudwatch.IWidget

	_jsii_.Invoke(
		i,
		"alarmWidgets",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IDashboardSegment) SummaryWidgets() *[]awscloudwatch.IWidget {
	var returns *[]awscloudwatch.IWidget

	_jsii_.Invoke(
		i,
		"summaryWidgets",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IDashboardSegment) Widgets() *[]awscloudwatch.IWidget {
	var returns *[]awscloudwatch.IWidget

	_jsii_.Invoke(
		i,
		"widgets",
		nil, // no parameters
		&returns,
	)

	return returns
}


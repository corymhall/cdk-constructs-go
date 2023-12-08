package cdkmonitoringconstructs

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/corymhall/cdk-constructs-go/cdkmonitoringconstructs/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudwatch"
)

// Experimental.
type SingleWidgetDashboardSegment interface {
	IDashboardSegment
	IDynamicDashboardSegment
	// Experimental.
	DashboardsToInclude() *[]*string
	// Experimental.
	Widget() awscloudwatch.IWidget
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
	// Returns widgets for the requested dashboard type.
	// Experimental.
	WidgetsForDashboard(name *string) *[]awscloudwatch.IWidget
}

// The jsii proxy struct for SingleWidgetDashboardSegment
type jsiiProxy_SingleWidgetDashboardSegment struct {
	jsiiProxy_IDashboardSegment
	jsiiProxy_IDynamicDashboardSegment
}

func (j *jsiiProxy_SingleWidgetDashboardSegment) DashboardsToInclude() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dashboardsToInclude",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SingleWidgetDashboardSegment) Widget() awscloudwatch.IWidget {
	var returns awscloudwatch.IWidget
	_jsii_.Get(
		j,
		"widget",
		&returns,
	)
	return returns
}


// Create a dashboard segment representing a single widget.
// Experimental.
func NewSingleWidgetDashboardSegment(widget awscloudwatch.IWidget, dashboardsToInclude *[]*string) SingleWidgetDashboardSegment {
	_init_.Initialize()

	if err := validateNewSingleWidgetDashboardSegmentParameters(widget); err != nil {
		panic(err)
	}
	j := jsiiProxy_SingleWidgetDashboardSegment{}

	_jsii_.Create(
		"cdk-monitoring-constructs.SingleWidgetDashboardSegment",
		[]interface{}{widget, dashboardsToInclude},
		&j,
	)

	return &j
}

// Create a dashboard segment representing a single widget.
// Experimental.
func NewSingleWidgetDashboardSegment_Override(s SingleWidgetDashboardSegment, widget awscloudwatch.IWidget, dashboardsToInclude *[]*string) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk-monitoring-constructs.SingleWidgetDashboardSegment",
		[]interface{}{widget, dashboardsToInclude},
		s,
	)
}

func (s *jsiiProxy_SingleWidgetDashboardSegment) AlarmWidgets() *[]awscloudwatch.IWidget {
	var returns *[]awscloudwatch.IWidget

	_jsii_.Invoke(
		s,
		"alarmWidgets",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SingleWidgetDashboardSegment) SummaryWidgets() *[]awscloudwatch.IWidget {
	var returns *[]awscloudwatch.IWidget

	_jsii_.Invoke(
		s,
		"summaryWidgets",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SingleWidgetDashboardSegment) Widgets() *[]awscloudwatch.IWidget {
	var returns *[]awscloudwatch.IWidget

	_jsii_.Invoke(
		s,
		"widgets",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SingleWidgetDashboardSegment) WidgetsForDashboard(name *string) *[]awscloudwatch.IWidget {
	if err := s.validateWidgetsForDashboardParameters(name); err != nil {
		panic(err)
	}
	var returns *[]awscloudwatch.IWidget

	_jsii_.Invoke(
		s,
		"widgetsForDashboard",
		[]interface{}{name},
		&returns,
	)

	return returns
}


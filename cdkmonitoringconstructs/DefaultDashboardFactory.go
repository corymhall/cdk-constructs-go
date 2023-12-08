package cdkmonitoringconstructs

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/corymhall/cdk-constructs-go/cdkmonitoringconstructs/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudwatch"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/corymhall/cdk-constructs-go/cdkmonitoringconstructs/internal"
)

// Experimental.
type DefaultDashboardFactory interface {
	constructs.Construct
	IDashboardFactory
	IDynamicDashboardFactory
	// Experimental.
	AlarmDashboard() awscloudwatch.Dashboard
	// Experimental.
	AnyDashboardCreated() *bool
	// Experimental.
	Dashboard() awscloudwatch.Dashboard
	// Experimental.
	Dashboards() *map[string]awscloudwatch.Dashboard
	// The tree node.
	// Experimental.
	Node() constructs.Node
	// Experimental.
	SummaryDashboard() awscloudwatch.Dashboard
	// Adds a dynamic dashboard segment.
	// Experimental.
	AddDynamicSegment(segment IDynamicDashboardSegment)
	// Experimental.
	AddSegment(props IDashboardFactoryProps)
	// Experimental.
	CreatedAlarmDashboard() awscloudwatch.Dashboard
	// Experimental.
	CreateDashboard(renderingPreference DashboardRenderingPreference, id *string, props *awscloudwatch.DashboardProps) awscloudwatch.Dashboard
	// Experimental.
	CreatedDashboard() awscloudwatch.Dashboard
	// Experimental.
	CreatedSummaryDashboard() awscloudwatch.Dashboard
	// Gets the dashboard for the requested dashboard type.
	// Experimental.
	GetDashboard(name *string) awscloudwatch.Dashboard
	// Returns a string representation of this construct.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DefaultDashboardFactory
type jsiiProxy_DefaultDashboardFactory struct {
	internal.Type__constructsConstruct
	jsiiProxy_IDashboardFactory
	jsiiProxy_IDynamicDashboardFactory
}

func (j *jsiiProxy_DefaultDashboardFactory) AlarmDashboard() awscloudwatch.Dashboard {
	var returns awscloudwatch.Dashboard
	_jsii_.Get(
		j,
		"alarmDashboard",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DefaultDashboardFactory) AnyDashboardCreated() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"anyDashboardCreated",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DefaultDashboardFactory) Dashboard() awscloudwatch.Dashboard {
	var returns awscloudwatch.Dashboard
	_jsii_.Get(
		j,
		"dashboard",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DefaultDashboardFactory) Dashboards() *map[string]awscloudwatch.Dashboard {
	var returns *map[string]awscloudwatch.Dashboard
	_jsii_.Get(
		j,
		"dashboards",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DefaultDashboardFactory) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DefaultDashboardFactory) SummaryDashboard() awscloudwatch.Dashboard {
	var returns awscloudwatch.Dashboard
	_jsii_.Get(
		j,
		"summaryDashboard",
		&returns,
	)
	return returns
}


// Experimental.
func NewDefaultDashboardFactory(scope constructs.Construct, id *string, props *MonitoringDashboardsProps) DefaultDashboardFactory {
	_init_.Initialize()

	if err := validateNewDefaultDashboardFactoryParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_DefaultDashboardFactory{}

	_jsii_.Create(
		"cdk-monitoring-constructs.DefaultDashboardFactory",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Experimental.
func NewDefaultDashboardFactory_Override(d DefaultDashboardFactory, scope constructs.Construct, id *string, props *MonitoringDashboardsProps) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk-monitoring-constructs.DefaultDashboardFactory",
		[]interface{}{scope, id, props},
		d,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead.
func DefaultDashboardFactory_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDefaultDashboardFactory_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"cdk-monitoring-constructs.DefaultDashboardFactory",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DefaultDashboardFactory) AddDynamicSegment(segment IDynamicDashboardSegment) {
	if err := d.validateAddDynamicSegmentParameters(segment); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"addDynamicSegment",
		[]interface{}{segment},
	)
}

func (d *jsiiProxy_DefaultDashboardFactory) AddSegment(props IDashboardFactoryProps) {
	if err := d.validateAddSegmentParameters(props); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"addSegment",
		[]interface{}{props},
	)
}

func (d *jsiiProxy_DefaultDashboardFactory) CreatedAlarmDashboard() awscloudwatch.Dashboard {
	var returns awscloudwatch.Dashboard

	_jsii_.Invoke(
		d,
		"createdAlarmDashboard",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DefaultDashboardFactory) CreateDashboard(renderingPreference DashboardRenderingPreference, id *string, props *awscloudwatch.DashboardProps) awscloudwatch.Dashboard {
	if err := d.validateCreateDashboardParameters(renderingPreference, id, props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Dashboard

	_jsii_.Invoke(
		d,
		"createDashboard",
		[]interface{}{renderingPreference, id, props},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DefaultDashboardFactory) CreatedDashboard() awscloudwatch.Dashboard {
	var returns awscloudwatch.Dashboard

	_jsii_.Invoke(
		d,
		"createdDashboard",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DefaultDashboardFactory) CreatedSummaryDashboard() awscloudwatch.Dashboard {
	var returns awscloudwatch.Dashboard

	_jsii_.Invoke(
		d,
		"createdSummaryDashboard",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DefaultDashboardFactory) GetDashboard(name *string) awscloudwatch.Dashboard {
	if err := d.validateGetDashboardParameters(name); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Dashboard

	_jsii_.Invoke(
		d,
		"getDashboard",
		[]interface{}{name},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DefaultDashboardFactory) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}


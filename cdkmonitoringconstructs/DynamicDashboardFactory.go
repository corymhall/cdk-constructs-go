package cdkmonitoringconstructs

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/corymhall/cdk-constructs-go/cdkmonitoringconstructs/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudwatch"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/corymhall/cdk-constructs-go/cdkmonitoringconstructs/internal"
)

// Experimental.
type DynamicDashboardFactory interface {
	constructs.Construct
	IDynamicDashboardFactory
	// Experimental.
	Dashboards() *map[string]awscloudwatch.Dashboard
	// The tree node.
	// Experimental.
	Node() constructs.Node
	// Adds a dynamic dashboard segment.
	// Experimental.
	AddDynamicSegment(segment IDynamicDashboardSegment)
	// Experimental.
	CreateDashboard(renderingPreference DashboardRenderingPreference, id *string, props *awscloudwatch.DashboardProps) awscloudwatch.Dashboard
	// Gets the dashboard for the requested dashboard type.
	// Experimental.
	GetDashboard(type_ *string) awscloudwatch.Dashboard
	// Returns a string representation of this construct.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DynamicDashboardFactory
type jsiiProxy_DynamicDashboardFactory struct {
	internal.Type__constructsConstruct
	jsiiProxy_IDynamicDashboardFactory
}

func (j *jsiiProxy_DynamicDashboardFactory) Dashboards() *map[string]awscloudwatch.Dashboard {
	var returns *map[string]awscloudwatch.Dashboard
	_jsii_.Get(
		j,
		"dashboards",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamicDashboardFactory) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}


// Experimental.
func NewDynamicDashboardFactory(scope constructs.Construct, id *string, props *MonitoringDynamicDashboardsProps) DynamicDashboardFactory {
	_init_.Initialize()

	if err := validateNewDynamicDashboardFactoryParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_DynamicDashboardFactory{}

	_jsii_.Create(
		"cdk-monitoring-constructs.DynamicDashboardFactory",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Experimental.
func NewDynamicDashboardFactory_Override(d DynamicDashboardFactory, scope constructs.Construct, id *string, props *MonitoringDynamicDashboardsProps) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk-monitoring-constructs.DynamicDashboardFactory",
		[]interface{}{scope, id, props},
		d,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead.
func DynamicDashboardFactory_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDynamicDashboardFactory_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"cdk-monitoring-constructs.DynamicDashboardFactory",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DynamicDashboardFactory) AddDynamicSegment(segment IDynamicDashboardSegment) {
	if err := d.validateAddDynamicSegmentParameters(segment); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"addDynamicSegment",
		[]interface{}{segment},
	)
}

func (d *jsiiProxy_DynamicDashboardFactory) CreateDashboard(renderingPreference DashboardRenderingPreference, id *string, props *awscloudwatch.DashboardProps) awscloudwatch.Dashboard {
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

func (d *jsiiProxy_DynamicDashboardFactory) GetDashboard(type_ *string) awscloudwatch.Dashboard {
	if err := d.validateGetDashboardParameters(type_); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Dashboard

	_jsii_.Invoke(
		d,
		"getDashboard",
		[]interface{}{type_},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DynamicDashboardFactory) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}


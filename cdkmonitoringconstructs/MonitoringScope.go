package cdkmonitoringconstructs

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/corymhall/cdk-constructs-go/cdkmonitoringconstructs/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/corymhall/cdk-constructs-go/cdkmonitoringconstructs/internal"
)

// A scope where all monitored constructs are managed from (i.e., alarms, dashboards, etc.).
//
// Standard usages will use {@link MonitoringFacade}.
// Experimental.
type MonitoringScope interface {
	constructs.Construct
	// The tree node.
	// Experimental.
	Node() constructs.Node
	// Creates a new alarm factory.
	//
	// Alarms created will be named with the given prefix, unless a local name override is present.
	// Experimental.
	CreateAlarmFactory(alarmNamePrefix *string) AlarmFactory
	// Creates a new factory that creates AWS Console URLs.
	// Experimental.
	CreateAwsConsoleUrlFactory() AwsConsoleUrlFactory
	// Creates a new metric factory.
	// Experimental.
	CreateMetricFactory() MetricFactory
	// Creates a new widget factory.
	// Experimental.
	CreateWidgetFactory() IWidgetFactory
	// Returns a string representation of this construct.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for MonitoringScope
type jsiiProxy_MonitoringScope struct {
	internal.Type__constructsConstruct
}

func (j *jsiiProxy_MonitoringScope) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}


// Creates a new construct node.
// Experimental.
func NewMonitoringScope_Override(m MonitoringScope, scope constructs.Construct, id *string) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk-monitoring-constructs.MonitoringScope",
		[]interface{}{scope, id},
		m,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead.
func MonitoringScope_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateMonitoringScope_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"cdk-monitoring-constructs.MonitoringScope",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitoringScope) CreateAlarmFactory(alarmNamePrefix *string) AlarmFactory {
	if err := m.validateCreateAlarmFactoryParameters(alarmNamePrefix); err != nil {
		panic(err)
	}
	var returns AlarmFactory

	_jsii_.Invoke(
		m,
		"createAlarmFactory",
		[]interface{}{alarmNamePrefix},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitoringScope) CreateAwsConsoleUrlFactory() AwsConsoleUrlFactory {
	var returns AwsConsoleUrlFactory

	_jsii_.Invoke(
		m,
		"createAwsConsoleUrlFactory",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitoringScope) CreateMetricFactory() MetricFactory {
	var returns MetricFactory

	_jsii_.Invoke(
		m,
		"createMetricFactory",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitoringScope) CreateWidgetFactory() IWidgetFactory {
	var returns IWidgetFactory

	_jsii_.Invoke(
		m,
		"createWidgetFactory",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitoringScope) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}


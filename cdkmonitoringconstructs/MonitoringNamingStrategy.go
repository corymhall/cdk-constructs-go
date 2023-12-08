package cdkmonitoringconstructs

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/corymhall/cdk-constructs-go/cdkmonitoringconstructs/jsii"
)

// Utility class to unify approach to naming monitoring sections.
// See: https://docs.aws.amazon.com/cdk/latest/guide/tokens.html#tokens_lazy
//
// Experimental.
type MonitoringNamingStrategy interface {
	// Experimental.
	Input() *NameResolutionInput
	// Experimental.
	ResolveAlarmFriendlyName() *string
	// Experimental.
	ResolveHumanReadableName() *string
}

// The jsii proxy struct for MonitoringNamingStrategy
type jsiiProxy_MonitoringNamingStrategy struct {
	_ byte // padding
}

func (j *jsiiProxy_MonitoringNamingStrategy) Input() *NameResolutionInput {
	var returns *NameResolutionInput
	_jsii_.Get(
		j,
		"input",
		&returns,
	)
	return returns
}


// Experimental.
func NewMonitoringNamingStrategy(input *NameResolutionInput) MonitoringNamingStrategy {
	_init_.Initialize()

	if err := validateNewMonitoringNamingStrategyParameters(input); err != nil {
		panic(err)
	}
	j := jsiiProxy_MonitoringNamingStrategy{}

	_jsii_.Create(
		"cdk-monitoring-constructs.MonitoringNamingStrategy",
		[]interface{}{input},
		&j,
	)

	return &j
}

// Experimental.
func NewMonitoringNamingStrategy_Override(m MonitoringNamingStrategy, input *NameResolutionInput) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk-monitoring-constructs.MonitoringNamingStrategy",
		[]interface{}{input},
		m,
	)
}

// Experimental.
func MonitoringNamingStrategy_IsAlarmFriendly(str *string) interface{} {
	_init_.Initialize()

	if err := validateMonitoringNamingStrategy_IsAlarmFriendlyParameters(str); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.StaticInvoke(
		"cdk-monitoring-constructs.MonitoringNamingStrategy",
		"isAlarmFriendly",
		[]interface{}{str},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitoringNamingStrategy) ResolveAlarmFriendlyName() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"resolveAlarmFriendlyName",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitoringNamingStrategy) ResolveHumanReadableName() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"resolveHumanReadableName",
		nil, // no parameters
		&returns,
	)

	return returns
}


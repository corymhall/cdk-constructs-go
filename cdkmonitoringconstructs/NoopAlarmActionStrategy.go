package cdkmonitoringconstructs

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/corymhall/cdk-constructs-go/cdkmonitoringconstructs/jsii"
)

// Alarm action strategy that does not add any actions.
// Experimental.
type NoopAlarmActionStrategy interface {
	IAlarmActionStrategy
	// Experimental.
	AddAlarmActions(_props *AlarmActionStrategyProps)
}

// The jsii proxy struct for NoopAlarmActionStrategy
type jsiiProxy_NoopAlarmActionStrategy struct {
	jsiiProxy_IAlarmActionStrategy
}

// Experimental.
func NewNoopAlarmActionStrategy() NoopAlarmActionStrategy {
	_init_.Initialize()

	j := jsiiProxy_NoopAlarmActionStrategy{}

	_jsii_.Create(
		"cdk-monitoring-constructs.NoopAlarmActionStrategy",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewNoopAlarmActionStrategy_Override(n NoopAlarmActionStrategy) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk-monitoring-constructs.NoopAlarmActionStrategy",
		nil, // no parameters
		n,
	)
}

func (n *jsiiProxy_NoopAlarmActionStrategy) AddAlarmActions(_props *AlarmActionStrategyProps) {
	if err := n.validateAddAlarmActionsParameters(_props); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		n,
		"addAlarmActions",
		[]interface{}{_props},
	)
}


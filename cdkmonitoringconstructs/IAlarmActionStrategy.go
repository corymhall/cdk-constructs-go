package cdkmonitoringconstructs

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// An object that appends actions to alarms.
// Experimental.
type IAlarmActionStrategy interface {
	// Experimental.
	AddAlarmActions(props *AlarmActionStrategyProps)
}

// The jsii proxy for IAlarmActionStrategy
type jsiiProxy_IAlarmActionStrategy struct {
	_ byte // padding
}

func (i *jsiiProxy_IAlarmActionStrategy) AddAlarmActions(props *AlarmActionStrategyProps) {
	if err := i.validateAddAlarmActionsParameters(props); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"addAlarmActions",
		[]interface{}{props},
	)
}


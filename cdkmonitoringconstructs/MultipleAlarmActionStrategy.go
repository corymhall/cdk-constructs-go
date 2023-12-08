package cdkmonitoringconstructs

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/corymhall/cdk-constructs-go/cdkmonitoringconstructs/jsii"
)

// Alarm action strategy that combines multiple actions in the same order as they were given.
// Experimental.
type MultipleAlarmActionStrategy interface {
	IAlarmActionStrategy
	// Experimental.
	Actions() *[]IAlarmActionStrategy
	// Experimental.
	AddAlarmActions(props *AlarmActionStrategyProps)
	// Returns list of alarm actions where any nested instances of MultipleAlarmActionStrategy are flattened.
	//
	// Returns: flattened list of alarm actions.
	// Experimental.
	FlattenedAlarmActions() *[]IAlarmActionStrategy
}

// The jsii proxy struct for MultipleAlarmActionStrategy
type jsiiProxy_MultipleAlarmActionStrategy struct {
	jsiiProxy_IAlarmActionStrategy
}

func (j *jsiiProxy_MultipleAlarmActionStrategy) Actions() *[]IAlarmActionStrategy {
	var returns *[]IAlarmActionStrategy
	_jsii_.Get(
		j,
		"actions",
		&returns,
	)
	return returns
}


// Experimental.
func NewMultipleAlarmActionStrategy(actions *[]IAlarmActionStrategy) MultipleAlarmActionStrategy {
	_init_.Initialize()

	if err := validateNewMultipleAlarmActionStrategyParameters(actions); err != nil {
		panic(err)
	}
	j := jsiiProxy_MultipleAlarmActionStrategy{}

	_jsii_.Create(
		"cdk-monitoring-constructs.MultipleAlarmActionStrategy",
		[]interface{}{actions},
		&j,
	)

	return &j
}

// Experimental.
func NewMultipleAlarmActionStrategy_Override(m MultipleAlarmActionStrategy, actions *[]IAlarmActionStrategy) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk-monitoring-constructs.MultipleAlarmActionStrategy",
		[]interface{}{actions},
		m,
	)
}

func (m *jsiiProxy_MultipleAlarmActionStrategy) AddAlarmActions(props *AlarmActionStrategyProps) {
	if err := m.validateAddAlarmActionsParameters(props); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"addAlarmActions",
		[]interface{}{props},
	)
}

func (m *jsiiProxy_MultipleAlarmActionStrategy) FlattenedAlarmActions() *[]IAlarmActionStrategy {
	var returns *[]IAlarmActionStrategy

	_jsii_.Invoke(
		m,
		"flattenedAlarmActions",
		nil, // no parameters
		&returns,
	)

	return returns
}


package cdkmonitoringconstructs

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/corymhall/cdk-constructs-go/cdkmonitoringconstructs/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2/awssns"
)

// Alarm action strategy that sends a notification to the specified SNS topic.
// Experimental.
type SnsAlarmActionStrategy interface {
	IAlarmActionStrategy
	// Experimental.
	OnAlarmTopic() awssns.ITopic
	// Experimental.
	OnInsufficientDataTopic() awssns.ITopic
	// Experimental.
	OnOkTopic() awssns.ITopic
	// Experimental.
	AddAlarmActions(props *AlarmActionStrategyProps)
}

// The jsii proxy struct for SnsAlarmActionStrategy
type jsiiProxy_SnsAlarmActionStrategy struct {
	jsiiProxy_IAlarmActionStrategy
}

func (j *jsiiProxy_SnsAlarmActionStrategy) OnAlarmTopic() awssns.ITopic {
	var returns awssns.ITopic
	_jsii_.Get(
		j,
		"onAlarmTopic",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnsAlarmActionStrategy) OnInsufficientDataTopic() awssns.ITopic {
	var returns awssns.ITopic
	_jsii_.Get(
		j,
		"onInsufficientDataTopic",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SnsAlarmActionStrategy) OnOkTopic() awssns.ITopic {
	var returns awssns.ITopic
	_jsii_.Get(
		j,
		"onOkTopic",
		&returns,
	)
	return returns
}


// Experimental.
func NewSnsAlarmActionStrategy(props *SnsAlarmActionStrategyProps) SnsAlarmActionStrategy {
	_init_.Initialize()

	if err := validateNewSnsAlarmActionStrategyParameters(props); err != nil {
		panic(err)
	}
	j := jsiiProxy_SnsAlarmActionStrategy{}

	_jsii_.Create(
		"cdk-monitoring-constructs.SnsAlarmActionStrategy",
		[]interface{}{props},
		&j,
	)

	return &j
}

// Experimental.
func NewSnsAlarmActionStrategy_Override(s SnsAlarmActionStrategy, props *SnsAlarmActionStrategyProps) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk-monitoring-constructs.SnsAlarmActionStrategy",
		[]interface{}{props},
		s,
	)
}

func (s *jsiiProxy_SnsAlarmActionStrategy) AddAlarmActions(props *AlarmActionStrategyProps) {
	if err := s.validateAddAlarmActionsParameters(props); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"addAlarmActions",
		[]interface{}{props},
	)
}


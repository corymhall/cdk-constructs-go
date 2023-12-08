package cdkmonitoringconstructs

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/corymhall/cdk-constructs-go/cdkmonitoringconstructs/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudwatchactions"
)

// Alarm action strategy that creates an AWS OpsCenter OpsItem.
// Experimental.
type OpsItemAlarmActionStrategy interface {
	IAlarmActionStrategy
	// OPS Item Category.
	// Experimental.
	Category() awscloudwatchactions.OpsItemCategory
	// OPS Item Severity.
	// Experimental.
	Severity() awscloudwatchactions.OpsItemSeverity
	// Experimental.
	AddAlarmActions(props *AlarmActionStrategyProps)
}

// The jsii proxy struct for OpsItemAlarmActionStrategy
type jsiiProxy_OpsItemAlarmActionStrategy struct {
	jsiiProxy_IAlarmActionStrategy
}

func (j *jsiiProxy_OpsItemAlarmActionStrategy) Category() awscloudwatchactions.OpsItemCategory {
	var returns awscloudwatchactions.OpsItemCategory
	_jsii_.Get(
		j,
		"category",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpsItemAlarmActionStrategy) Severity() awscloudwatchactions.OpsItemSeverity {
	var returns awscloudwatchactions.OpsItemSeverity
	_jsii_.Get(
		j,
		"severity",
		&returns,
	)
	return returns
}


// Experimental.
func NewOpsItemAlarmActionStrategy(severity awscloudwatchactions.OpsItemSeverity, category awscloudwatchactions.OpsItemCategory) OpsItemAlarmActionStrategy {
	_init_.Initialize()

	if err := validateNewOpsItemAlarmActionStrategyParameters(severity); err != nil {
		panic(err)
	}
	j := jsiiProxy_OpsItemAlarmActionStrategy{}

	_jsii_.Create(
		"cdk-monitoring-constructs.OpsItemAlarmActionStrategy",
		[]interface{}{severity, category},
		&j,
	)

	return &j
}

// Experimental.
func NewOpsItemAlarmActionStrategy_Override(o OpsItemAlarmActionStrategy, severity awscloudwatchactions.OpsItemSeverity, category awscloudwatchactions.OpsItemCategory) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk-monitoring-constructs.OpsItemAlarmActionStrategy",
		[]interface{}{severity, category},
		o,
	)
}

func (o *jsiiProxy_OpsItemAlarmActionStrategy) AddAlarmActions(props *AlarmActionStrategyProps) {
	if err := o.validateAddAlarmActionsParameters(props); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"addAlarmActions",
		[]interface{}{props},
	)
}


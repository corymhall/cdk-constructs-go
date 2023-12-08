package cdkmonitoringconstructs

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Strategy used to name alarms, their widgets, and their dedupe strings.
// Experimental.
type IAlarmNamingStrategy interface {
	// How to generate the deduplication string for an alarm.
	// Experimental.
	GetDedupeString(props *AlarmNamingInput) *string
	// How to generate the name of an alarm.
	// Experimental.
	GetName(props *AlarmNamingInput) *string
	// How to generate the label for the alarm displayed on a widget.
	// Experimental.
	GetWidgetLabel(props *AlarmNamingInput) *string
}

// The jsii proxy for IAlarmNamingStrategy
type jsiiProxy_IAlarmNamingStrategy struct {
	_ byte // padding
}

func (i *jsiiProxy_IAlarmNamingStrategy) GetDedupeString(props *AlarmNamingInput) *string {
	if err := i.validateGetDedupeStringParameters(props); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		i,
		"getDedupeString",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IAlarmNamingStrategy) GetName(props *AlarmNamingInput) *string {
	if err := i.validateGetNameParameters(props); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		i,
		"getName",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IAlarmNamingStrategy) GetWidgetLabel(props *AlarmNamingInput) *string {
	if err := i.validateGetWidgetLabelParameters(props); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		i,
		"getWidgetLabel",
		[]interface{}{props},
		&returns,
	)

	return returns
}


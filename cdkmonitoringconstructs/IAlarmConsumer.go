package cdkmonitoringconstructs

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Experimental.
type IAlarmConsumer interface {
	// Experimental.
	Consume(alarms *[]*AlarmWithAnnotation)
}

// The jsii proxy for IAlarmConsumer
type jsiiProxy_IAlarmConsumer struct {
	_ byte // padding
}

func (i *jsiiProxy_IAlarmConsumer) Consume(alarms *[]*AlarmWithAnnotation) {
	if err := i.validateConsumeParameters(alarms); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"consume",
		[]interface{}{alarms},
	)
}


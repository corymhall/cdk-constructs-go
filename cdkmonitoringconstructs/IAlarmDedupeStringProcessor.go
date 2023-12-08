package cdkmonitoringconstructs

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Strategy used to finalize dedupe string.
// Experimental.
type IAlarmDedupeStringProcessor interface {
	// Process the dedupe string which was auto-generated.
	//
	// Returns: final dedupe string.
	// Experimental.
	ProcessDedupeString(dedupeString *string) *string
	// Process the dedupe string which was specified by the user as an override.
	//
	// Returns: final dedupe string.
	// Experimental.
	ProcessDedupeStringOverride(dedupeString *string) *string
}

// The jsii proxy for IAlarmDedupeStringProcessor
type jsiiProxy_IAlarmDedupeStringProcessor struct {
	_ byte // padding
}

func (i *jsiiProxy_IAlarmDedupeStringProcessor) ProcessDedupeString(dedupeString *string) *string {
	if err := i.validateProcessDedupeStringParameters(dedupeString); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		i,
		"processDedupeString",
		[]interface{}{dedupeString},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IAlarmDedupeStringProcessor) ProcessDedupeStringOverride(dedupeString *string) *string {
	if err := i.validateProcessDedupeStringOverrideParameters(dedupeString); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		i,
		"processDedupeStringOverride",
		[]interface{}{dedupeString},
		&returns,
	)

	return returns
}


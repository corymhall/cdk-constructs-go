package cdkmonitoringconstructs

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/corymhall/cdk-constructs-go/cdkmonitoringconstructs/jsii"
)

// Dedupe string processor that adds prefix and/or suffix to the dedupe string.
// Experimental.
type ExtendDedupeString interface {
	IAlarmDedupeStringProcessor
	// Process the dedupe string which was auto-generated.
	// Experimental.
	ProcessDedupeString(dedupeString *string) *string
	// Process the dedupe string which was specified by the user as an override.
	// Experimental.
	ProcessDedupeStringOverride(dedupeString *string) *string
}

// The jsii proxy struct for ExtendDedupeString
type jsiiProxy_ExtendDedupeString struct {
	jsiiProxy_IAlarmDedupeStringProcessor
}

// Experimental.
func NewExtendDedupeString(prefix *string, suffix *string) ExtendDedupeString {
	_init_.Initialize()

	j := jsiiProxy_ExtendDedupeString{}

	_jsii_.Create(
		"cdk-monitoring-constructs.ExtendDedupeString",
		[]interface{}{prefix, suffix},
		&j,
	)

	return &j
}

// Experimental.
func NewExtendDedupeString_Override(e ExtendDedupeString, prefix *string, suffix *string) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk-monitoring-constructs.ExtendDedupeString",
		[]interface{}{prefix, suffix},
		e,
	)
}

func (e *jsiiProxy_ExtendDedupeString) ProcessDedupeString(dedupeString *string) *string {
	if err := e.validateProcessDedupeStringParameters(dedupeString); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		e,
		"processDedupeString",
		[]interface{}{dedupeString},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ExtendDedupeString) ProcessDedupeStringOverride(dedupeString *string) *string {
	if err := e.validateProcessDedupeStringOverrideParameters(dedupeString); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		e,
		"processDedupeStringOverride",
		[]interface{}{dedupeString},
		&returns,
	)

	return returns
}


package cdkmonitoringconstructs

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/corymhall/cdk-constructs-go/cdkmonitoringconstructs/jsii"
)

// Default dedupe strategy - does not add any prefix nor suffix.
// Experimental.
type DoNotModifyDedupeString interface {
	ExtendDedupeString
	// Process the dedupe string which was auto-generated.
	// Experimental.
	ProcessDedupeString(dedupeString *string) *string
	// Process the dedupe string which was specified by the user as an override.
	// Experimental.
	ProcessDedupeStringOverride(dedupeString *string) *string
}

// The jsii proxy struct for DoNotModifyDedupeString
type jsiiProxy_DoNotModifyDedupeString struct {
	jsiiProxy_ExtendDedupeString
}

// Experimental.
func NewDoNotModifyDedupeString(prefix *string, suffix *string) DoNotModifyDedupeString {
	_init_.Initialize()

	j := jsiiProxy_DoNotModifyDedupeString{}

	_jsii_.Create(
		"cdk-monitoring-constructs.DoNotModifyDedupeString",
		[]interface{}{prefix, suffix},
		&j,
	)

	return &j
}

// Experimental.
func NewDoNotModifyDedupeString_Override(d DoNotModifyDedupeString, prefix *string, suffix *string) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk-monitoring-constructs.DoNotModifyDedupeString",
		[]interface{}{prefix, suffix},
		d,
	)
}

func (d *jsiiProxy_DoNotModifyDedupeString) ProcessDedupeString(dedupeString *string) *string {
	if err := d.validateProcessDedupeStringParameters(dedupeString); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		d,
		"processDedupeString",
		[]interface{}{dedupeString},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DoNotModifyDedupeString) ProcessDedupeStringOverride(dedupeString *string) *string {
	if err := d.validateProcessDedupeStringOverrideParameters(dedupeString); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		d,
		"processDedupeStringOverride",
		[]interface{}{dedupeString},
		&returns,
	)

	return returns
}


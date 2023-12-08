package cdkmonitoringconstructs

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/corymhall/cdk-constructs-go/cdkmonitoringconstructs/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudwatch"
	"github.com/corymhall/cdk-constructs-go/cdkmonitoringconstructs/internal"
)

// Experimental.
type KeyValueTableWidget interface {
	awscloudwatch.TextWidget
	// The amount of vertical grid units the widget will take up.
	// Experimental.
	Height() *float64
	// Any warnings that are produced as a result of putting together this widget.
	// Experimental.
	Warnings() *[]*string
	// Any warnings that are produced as a result of putting together this widget.
	// Experimental.
	WarningsV2() *map[string]*string
	// The amount of horizontal grid units the widget will take up.
	// Experimental.
	Width() *float64
	// Experimental.
	X() *float64
	// Experimental.
	SetX(val *float64)
	// Experimental.
	Y() *float64
	// Experimental.
	SetY(val *float64)
	// Copy the warnings from the given metric.
	// Experimental.
	CopyMetricWarnings(ms ...awscloudwatch.IMetric)
	// Place the widget at a given position.
	// Experimental.
	Position(x *float64, y *float64)
	// Return the widget JSON for use in the dashboard.
	// Experimental.
	ToJson() *[]interface{}
}

// The jsii proxy struct for KeyValueTableWidget
type jsiiProxy_KeyValueTableWidget struct {
	internal.Type__awscloudwatchTextWidget
}

func (j *jsiiProxy_KeyValueTableWidget) Height() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"height",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KeyValueTableWidget) Warnings() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"warnings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KeyValueTableWidget) WarningsV2() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"warningsV2",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KeyValueTableWidget) Width() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"width",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KeyValueTableWidget) X() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"x",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KeyValueTableWidget) Y() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"y",
		&returns,
	)
	return returns
}


// Experimental.
func NewKeyValueTableWidget(data *[]*map[string]interface{}) KeyValueTableWidget {
	_init_.Initialize()

	if err := validateNewKeyValueTableWidgetParameters(data); err != nil {
		panic(err)
	}
	j := jsiiProxy_KeyValueTableWidget{}

	_jsii_.Create(
		"cdk-monitoring-constructs.KeyValueTableWidget",
		[]interface{}{data},
		&j,
	)

	return &j
}

// Experimental.
func NewKeyValueTableWidget_Override(k KeyValueTableWidget, data *[]*map[string]interface{}) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk-monitoring-constructs.KeyValueTableWidget",
		[]interface{}{data},
		k,
	)
}

func (j *jsiiProxy_KeyValueTableWidget)SetX(val *float64) {
	_jsii_.Set(
		j,
		"x",
		val,
	)
}

func (j *jsiiProxy_KeyValueTableWidget)SetY(val *float64) {
	_jsii_.Set(
		j,
		"y",
		val,
	)
}

func (k *jsiiProxy_KeyValueTableWidget) CopyMetricWarnings(ms ...awscloudwatch.IMetric) {
	args := []interface{}{}
	for _, a := range ms {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		k,
		"copyMetricWarnings",
		args,
	)
}

func (k *jsiiProxy_KeyValueTableWidget) Position(x *float64, y *float64) {
	if err := k.validatePositionParameters(x, y); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"position",
		[]interface{}{x, y},
	)
}

func (k *jsiiProxy_KeyValueTableWidget) ToJson() *[]interface{} {
	var returns *[]interface{}

	_jsii_.Invoke(
		k,
		"toJson",
		nil, // no parameters
		&returns,
	)

	return returns
}


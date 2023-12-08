package cdkmonitoringconstructs

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/corymhall/cdk-constructs-go/cdkmonitoringconstructs/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudwatch"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/corymhall/cdk-constructs-go/cdkmonitoringconstructs/internal"
)

// Support for rendering bitmap widgets on the server side.
//
// It is a custom widget lambda with some additional roles and helper methods.
// Experimental.
type BitmapWidgetRenderingSupport interface {
	constructs.Construct
	// Experimental.
	Handler() awslambda.IFunction
	// The tree node.
	// Experimental.
	Node() constructs.Node
	// Experimental.
	AsBitmap(widget awscloudwatch.IWidget) CustomWidget
	// Experimental.
	GetWidgetProperties(widget awscloudwatch.IWidget) interface{}
	// Returns a string representation of this construct.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for BitmapWidgetRenderingSupport
type jsiiProxy_BitmapWidgetRenderingSupport struct {
	internal.Type__constructsConstruct
}

func (j *jsiiProxy_BitmapWidgetRenderingSupport) Handler() awslambda.IFunction {
	var returns awslambda.IFunction
	_jsii_.Get(
		j,
		"handler",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BitmapWidgetRenderingSupport) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}


// Experimental.
func NewBitmapWidgetRenderingSupport(scope constructs.Construct, id *string) BitmapWidgetRenderingSupport {
	_init_.Initialize()

	if err := validateNewBitmapWidgetRenderingSupportParameters(scope, id); err != nil {
		panic(err)
	}
	j := jsiiProxy_BitmapWidgetRenderingSupport{}

	_jsii_.Create(
		"cdk-monitoring-constructs.BitmapWidgetRenderingSupport",
		[]interface{}{scope, id},
		&j,
	)

	return &j
}

// Experimental.
func NewBitmapWidgetRenderingSupport_Override(b BitmapWidgetRenderingSupport, scope constructs.Construct, id *string) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk-monitoring-constructs.BitmapWidgetRenderingSupport",
		[]interface{}{scope, id},
		b,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead.
func BitmapWidgetRenderingSupport_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateBitmapWidgetRenderingSupport_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"cdk-monitoring-constructs.BitmapWidgetRenderingSupport",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BitmapWidgetRenderingSupport) AsBitmap(widget awscloudwatch.IWidget) CustomWidget {
	if err := b.validateAsBitmapParameters(widget); err != nil {
		panic(err)
	}
	var returns CustomWidget

	_jsii_.Invoke(
		b,
		"asBitmap",
		[]interface{}{widget},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BitmapWidgetRenderingSupport) GetWidgetProperties(widget awscloudwatch.IWidget) interface{} {
	if err := b.validateGetWidgetPropertiesParameters(widget); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		b,
		"getWidgetProperties",
		[]interface{}{widget},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BitmapWidgetRenderingSupport) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}


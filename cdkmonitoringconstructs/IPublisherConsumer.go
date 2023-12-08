package cdkmonitoringconstructs

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
)

// Experimental.
type IPublisherConsumer interface {
	// Experimental.
	Consume(lambdaFunction awslambda.IFunction)
}

// The jsii proxy for IPublisherConsumer
type jsiiProxy_IPublisherConsumer struct {
	_ byte // padding
}

func (i *jsiiProxy_IPublisherConsumer) Consume(lambdaFunction awslambda.IFunction) {
	if err := i.validateConsumeParameters(lambdaFunction); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"consume",
		[]interface{}{lambdaFunction},
	)
}


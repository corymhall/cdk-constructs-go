package cdkmonitoringconstructs

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/corymhall/cdk-constructs-go/cdkmonitoringconstructs/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
	"github.com/aws/aws-cdk-go/awscdk/v2/awssecretsmanager"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/corymhall/cdk-constructs-go/cdkmonitoringconstructs/internal"
)

// Experimental.
type SecretsManagerMetricsPublisher interface {
	constructs.Construct
	// Experimental.
	Lambda() awslambda.IFunction
	// The tree node.
	// Experimental.
	Node() constructs.Node
	// Experimental.
	AddSecret(secret awssecretsmanager.ISecret)
	// Returns a string representation of this construct.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for SecretsManagerMetricsPublisher
type jsiiProxy_SecretsManagerMetricsPublisher struct {
	internal.Type__constructsConstruct
}

func (j *jsiiProxy_SecretsManagerMetricsPublisher) Lambda() awslambda.IFunction {
	var returns awslambda.IFunction
	_jsii_.Get(
		j,
		"lambda",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecretsManagerMetricsPublisher) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}


// Experimental.
func SecretsManagerMetricsPublisher_GetInstance(scope MonitoringScope) SecretsManagerMetricsPublisher {
	_init_.Initialize()

	if err := validateSecretsManagerMetricsPublisher_GetInstanceParameters(scope); err != nil {
		panic(err)
	}
	var returns SecretsManagerMetricsPublisher

	_jsii_.StaticInvoke(
		"cdk-monitoring-constructs.SecretsManagerMetricsPublisher",
		"getInstance",
		[]interface{}{scope},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead.
func SecretsManagerMetricsPublisher_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateSecretsManagerMetricsPublisher_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"cdk-monitoring-constructs.SecretsManagerMetricsPublisher",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SecretsManagerMetricsPublisher) AddSecret(secret awssecretsmanager.ISecret) {
	if err := s.validateAddSecretParameters(secret); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"addSecret",
		[]interface{}{secret},
	)
}

func (s *jsiiProxy_SecretsManagerMetricsPublisher) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}


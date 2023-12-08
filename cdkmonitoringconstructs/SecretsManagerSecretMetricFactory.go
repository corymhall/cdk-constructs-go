package cdkmonitoringconstructs

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/corymhall/cdk-constructs-go/cdkmonitoringconstructs/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2/awssecretsmanager"
)

// Experimental.
type SecretsManagerSecretMetricFactory interface {
	// Experimental.
	DimensionsMap() *map[string]*string
	// Experimental.
	MetricFactory() MetricFactory
	// Experimental.
	Secret() awssecretsmanager.ISecret
	// Experimental.
	MetricDaysSinceLastChange() interface{}
	// Experimental.
	MetricDaysSinceLastRotation() interface{}
}

// The jsii proxy struct for SecretsManagerSecretMetricFactory
type jsiiProxy_SecretsManagerSecretMetricFactory struct {
	_ byte // padding
}

func (j *jsiiProxy_SecretsManagerSecretMetricFactory) DimensionsMap() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"dimensionsMap",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecretsManagerSecretMetricFactory) MetricFactory() MetricFactory {
	var returns MetricFactory
	_jsii_.Get(
		j,
		"metricFactory",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecretsManagerSecretMetricFactory) Secret() awssecretsmanager.ISecret {
	var returns awssecretsmanager.ISecret
	_jsii_.Get(
		j,
		"secret",
		&returns,
	)
	return returns
}


// Experimental.
func NewSecretsManagerSecretMetricFactory(metricFactory MetricFactory, props *SecretsManagerSecretMetricFactoryProps) SecretsManagerSecretMetricFactory {
	_init_.Initialize()

	if err := validateNewSecretsManagerSecretMetricFactoryParameters(metricFactory, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_SecretsManagerSecretMetricFactory{}

	_jsii_.Create(
		"cdk-monitoring-constructs.SecretsManagerSecretMetricFactory",
		[]interface{}{metricFactory, props},
		&j,
	)

	return &j
}

// Experimental.
func NewSecretsManagerSecretMetricFactory_Override(s SecretsManagerSecretMetricFactory, metricFactory MetricFactory, props *SecretsManagerSecretMetricFactoryProps) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk-monitoring-constructs.SecretsManagerSecretMetricFactory",
		[]interface{}{metricFactory, props},
		s,
	)
}

func SecretsManagerSecretMetricFactory_MetricNameDaysSinceLastChange() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"cdk-monitoring-constructs.SecretsManagerSecretMetricFactory",
		"MetricNameDaysSinceLastChange",
		&returns,
	)
	return returns
}

func SecretsManagerSecretMetricFactory_MetricNameDaysSinceLastRotation() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"cdk-monitoring-constructs.SecretsManagerSecretMetricFactory",
		"MetricNameDaysSinceLastRotation",
		&returns,
	)
	return returns
}

func SecretsManagerSecretMetricFactory_Namespace() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"cdk-monitoring-constructs.SecretsManagerSecretMetricFactory",
		"Namespace",
		&returns,
	)
	return returns
}

func (s *jsiiProxy_SecretsManagerSecretMetricFactory) MetricDaysSinceLastChange() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"metricDaysSinceLastChange",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SecretsManagerSecretMetricFactory) MetricDaysSinceLastRotation() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"metricDaysSinceLastRotation",
		nil, // no parameters
		&returns,
	)

	return returns
}


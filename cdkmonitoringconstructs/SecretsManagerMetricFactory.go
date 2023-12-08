package cdkmonitoringconstructs

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/corymhall/cdk-constructs-go/cdkmonitoringconstructs/jsii"
)

// Experimental.
type SecretsManagerMetricFactory interface {
	// Experimental.
	MetricFactory() MetricFactory
	// Experimental.
	MetricSecretCount() interface{}
}

// The jsii proxy struct for SecretsManagerMetricFactory
type jsiiProxy_SecretsManagerMetricFactory struct {
	_ byte // padding
}

func (j *jsiiProxy_SecretsManagerMetricFactory) MetricFactory() MetricFactory {
	var returns MetricFactory
	_jsii_.Get(
		j,
		"metricFactory",
		&returns,
	)
	return returns
}


// Experimental.
func NewSecretsManagerMetricFactory(metricFactory MetricFactory) SecretsManagerMetricFactory {
	_init_.Initialize()

	if err := validateNewSecretsManagerMetricFactoryParameters(metricFactory); err != nil {
		panic(err)
	}
	j := jsiiProxy_SecretsManagerMetricFactory{}

	_jsii_.Create(
		"cdk-monitoring-constructs.SecretsManagerMetricFactory",
		[]interface{}{metricFactory},
		&j,
	)

	return &j
}

// Experimental.
func NewSecretsManagerMetricFactory_Override(s SecretsManagerMetricFactory, metricFactory MetricFactory) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk-monitoring-constructs.SecretsManagerMetricFactory",
		[]interface{}{metricFactory},
		s,
	)
}

func (s *jsiiProxy_SecretsManagerMetricFactory) MetricSecretCount() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"metricSecretCount",
		nil, // no parameters
		&returns,
	)

	return returns
}


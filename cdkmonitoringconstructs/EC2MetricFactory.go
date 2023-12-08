package cdkmonitoringconstructs

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/corymhall/cdk-constructs-go/cdkmonitoringconstructs/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudwatch"
)

// Experimental.
type EC2MetricFactory interface {
	// Experimental.
	MetricFactory() MetricFactory
	// Experimental.
	Strategy() IEC2MetricFactoryStrategy
	// The percentage of allocated EC2 compute units that are currently in use on the instance.
	//
	// This metric identifies the processing power required to run an application on a selected instance.
	// Depending on the instance type, tools in your operating system can show a lower percentage than
	// CloudWatch when the instance is not allocated a full processor core.
	// Experimental.
	MetricAverageCpuUtilisationPercent() *[]awscloudwatch.IMetric
	// Bytes read from all instance store volumes available to the instance.
	//
	// This metric is used to determine the volume of the data the application reads from the hard disk of the instance.
	// This can be used to determine the speed of the application.
	// Experimental.
	MetricAverageDiskReadBytes() *[]interface{}
	// Completed read operations from all instance store volumes available to the instance in a specified period of time.
	// Experimental.
	MetricAverageDiskReadOps() *[]interface{}
	// Bytes written to all instance store volumes available to the instance.
	//
	// This metric is used to determine the volume of the data the application writes onto the hard disk of the instance.
	// This can be used to determine the speed of the application.
	// Experimental.
	MetricAverageDiskWriteBytes() *[]interface{}
	// Completed write operations to all instance store volumes available to the instance in a specified period of time.
	// Experimental.
	MetricAverageDiskWriteOps() *[]interface{}
	// The number of bytes received on all network interfaces by the instance.
	//
	// This metric identifies the volume of incoming network traffic to a single instance.
	// Experimental.
	MetricAverageNetworkInRateBytes() *[]awscloudwatch.IMetric
	// The number of bytes sent out on all network interfaces by the instance.
	//
	// This metric identifies the volume of outgoing network traffic from a single instance.
	// Experimental.
	MetricAverageNetworkOutRateBytes() *[]awscloudwatch.IMetric
}

// The jsii proxy struct for EC2MetricFactory
type jsiiProxy_EC2MetricFactory struct {
	_ byte // padding
}

func (j *jsiiProxy_EC2MetricFactory) MetricFactory() MetricFactory {
	var returns MetricFactory
	_jsii_.Get(
		j,
		"metricFactory",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EC2MetricFactory) Strategy() IEC2MetricFactoryStrategy {
	var returns IEC2MetricFactoryStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Experimental.
func NewEC2MetricFactory(metricFactory MetricFactory, props *EC2MetricFactoryProps) EC2MetricFactory {
	_init_.Initialize()

	if err := validateNewEC2MetricFactoryParameters(metricFactory, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_EC2MetricFactory{}

	_jsii_.Create(
		"cdk-monitoring-constructs.EC2MetricFactory",
		[]interface{}{metricFactory, props},
		&j,
	)

	return &j
}

// Experimental.
func NewEC2MetricFactory_Override(e EC2MetricFactory, metricFactory MetricFactory, props *EC2MetricFactoryProps) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk-monitoring-constructs.EC2MetricFactory",
		[]interface{}{metricFactory, props},
		e,
	)
}

func (e *jsiiProxy_EC2MetricFactory) MetricAverageCpuUtilisationPercent() *[]awscloudwatch.IMetric {
	var returns *[]awscloudwatch.IMetric

	_jsii_.Invoke(
		e,
		"metricAverageCpuUtilisationPercent",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EC2MetricFactory) MetricAverageDiskReadBytes() *[]interface{} {
	var returns *[]interface{}

	_jsii_.Invoke(
		e,
		"metricAverageDiskReadBytes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EC2MetricFactory) MetricAverageDiskReadOps() *[]interface{} {
	var returns *[]interface{}

	_jsii_.Invoke(
		e,
		"metricAverageDiskReadOps",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EC2MetricFactory) MetricAverageDiskWriteBytes() *[]interface{} {
	var returns *[]interface{}

	_jsii_.Invoke(
		e,
		"metricAverageDiskWriteBytes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EC2MetricFactory) MetricAverageDiskWriteOps() *[]interface{} {
	var returns *[]interface{}

	_jsii_.Invoke(
		e,
		"metricAverageDiskWriteOps",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EC2MetricFactory) MetricAverageNetworkInRateBytes() *[]awscloudwatch.IMetric {
	var returns *[]awscloudwatch.IMetric

	_jsii_.Invoke(
		e,
		"metricAverageNetworkInRateBytes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EC2MetricFactory) MetricAverageNetworkOutRateBytes() *[]awscloudwatch.IMetric {
	var returns *[]awscloudwatch.IMetric

	_jsii_.Invoke(
		e,
		"metricAverageNetworkOutRateBytes",
		nil, // no parameters
		&returns,
	)

	return returns
}


package cdkmonitoringconstructs

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudwatch"
)

// Experimental.
type IEC2MetricFactoryStrategy interface {
	// Experimental.
	CreateMetrics(metricFactory MetricFactory, metricName *string, statistic MetricStatistic) *[]awscloudwatch.IMetric
}

// The jsii proxy for IEC2MetricFactoryStrategy
type jsiiProxy_IEC2MetricFactoryStrategy struct {
	_ byte // padding
}

func (i *jsiiProxy_IEC2MetricFactoryStrategy) CreateMetrics(metricFactory MetricFactory, metricName *string, statistic MetricStatistic) *[]awscloudwatch.IMetric {
	if err := i.validateCreateMetricsParameters(metricFactory, metricName, statistic); err != nil {
		panic(err)
	}
	var returns *[]awscloudwatch.IMetric

	_jsii_.Invoke(
		i,
		"createMetrics",
		[]interface{}{metricFactory, metricName, statistic},
		&returns,
	)

	return returns
}


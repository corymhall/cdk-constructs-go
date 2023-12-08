//go:build !no_runtime_type_checking

package cdkmonitoringconstructs

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudwatch"
	"github.com/aws/aws-cdk-go/awscdk/v2/awselasticsearch"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsopensearchservice"
)

func (o *jsiiProxy_OpenSearchBackportedMetrics) validateMetricParameters(metricName *string, props *awscloudwatch.MetricOptions) error {
	if metricName == nil {
		return fmt.Errorf("parameter metricName is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}

func (o *jsiiProxy_OpenSearchBackportedMetrics) validateMetricAutomatedSnapshotFailureParameters(props *awscloudwatch.MetricOptions) error {
	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}

func (o *jsiiProxy_OpenSearchBackportedMetrics) validateMetricClusterIndexWriteBlockedParameters(props *awscloudwatch.MetricOptions) error {
	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}

func (o *jsiiProxy_OpenSearchBackportedMetrics) validateMetricClusterIndexWritesBlockedParameters(props *awscloudwatch.MetricOptions) error {
	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}

func (o *jsiiProxy_OpenSearchBackportedMetrics) validateMetricClusterStatusRedParameters(props *awscloudwatch.MetricOptions) error {
	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}

func (o *jsiiProxy_OpenSearchBackportedMetrics) validateMetricClusterStatusYellowParameters(props *awscloudwatch.MetricOptions) error {
	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}

func (o *jsiiProxy_OpenSearchBackportedMetrics) validateMetricCPUUtilizationParameters(props *awscloudwatch.MetricOptions) error {
	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}

func (o *jsiiProxy_OpenSearchBackportedMetrics) validateMetricFreeStorageSpaceParameters(props *awscloudwatch.MetricOptions) error {
	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}

func (o *jsiiProxy_OpenSearchBackportedMetrics) validateMetricIndexingLatencyParameters(props *awscloudwatch.MetricOptions) error {
	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}

func (o *jsiiProxy_OpenSearchBackportedMetrics) validateMetricJVMMemoryPressureParameters(props *awscloudwatch.MetricOptions) error {
	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}

func (o *jsiiProxy_OpenSearchBackportedMetrics) validateMetricKMSKeyErrorParameters(props *awscloudwatch.MetricOptions) error {
	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}

func (o *jsiiProxy_OpenSearchBackportedMetrics) validateMetricKMSKeyInaccessibleParameters(props *awscloudwatch.MetricOptions) error {
	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}

func (o *jsiiProxy_OpenSearchBackportedMetrics) validateMetricMasterCPUUtilizationParameters(props *awscloudwatch.MetricOptions) error {
	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}

func (o *jsiiProxy_OpenSearchBackportedMetrics) validateMetricMasterJVMMemoryPressureParameters(props *awscloudwatch.MetricOptions) error {
	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}

func (o *jsiiProxy_OpenSearchBackportedMetrics) validateMetricNodesParameters(props *awscloudwatch.MetricOptions) error {
	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}

func (o *jsiiProxy_OpenSearchBackportedMetrics) validateMetricSearchableDocumentsParameters(props *awscloudwatch.MetricOptions) error {
	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}

func (o *jsiiProxy_OpenSearchBackportedMetrics) validateMetricSearchLatencyParameters(props *awscloudwatch.MetricOptions) error {
	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}

func validateNewOpenSearchBackportedMetricsParameters(domain interface{}) error {
	if domain == nil {
		return fmt.Errorf("parameter domain is required, but nil was provided")
	}
	switch domain.(type) {
	case awselasticsearch.IDomain:
		// ok
	case awselasticsearch.CfnDomain:
		// ok
	case awsopensearchservice.IDomain:
		// ok
	case awsopensearchservice.CfnDomain:
		// ok
	default:
		if !_jsii_.IsAnonymousProxy(domain) {
			return fmt.Errorf("parameter domain must be one of the allowed types: awselasticsearch.IDomain, awselasticsearch.CfnDomain, awsopensearchservice.IDomain, awsopensearchservice.CfnDomain; received %#v (a %T)", domain, domain)
		}
	}

	return nil
}


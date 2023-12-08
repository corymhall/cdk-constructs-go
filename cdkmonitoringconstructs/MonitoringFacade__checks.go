//go:build !no_runtime_type_checking

package cdkmonitoringconstructs

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudwatch"
	"github.com/aws/constructs-go/constructs/v10"
)

func (m *jsiiProxy_MonitoringFacade) validateAddDynamicSegmentParameters(segment IDynamicDashboardSegment) error {
	if segment == nil {
		return fmt.Errorf("parameter segment is required, but nil was provided")
	}

	return nil
}

func (m *jsiiProxy_MonitoringFacade) validateAddLargeHeaderParameters(text *string) error {
	if text == nil {
		return fmt.Errorf("parameter text is required, but nil was provided")
	}

	return nil
}

func (m *jsiiProxy_MonitoringFacade) validateAddMediumHeaderParameters(text *string) error {
	if text == nil {
		return fmt.Errorf("parameter text is required, but nil was provided")
	}

	return nil
}

func (m *jsiiProxy_MonitoringFacade) validateAddSegmentParameters(segment IDashboardSegment, overrideProps *MonitoringDashboardsOverrideProps) error {
	if segment == nil {
		return fmt.Errorf("parameter segment is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(overrideProps, func() string { return "parameter overrideProps" }); err != nil {
		return err
	}

	return nil
}

func (m *jsiiProxy_MonitoringFacade) validateAddSmallHeaderParameters(text *string) error {
	if text == nil {
		return fmt.Errorf("parameter text is required, but nil was provided")
	}

	return nil
}

func (m *jsiiProxy_MonitoringFacade) validateAddWidgetParameters(widget awscloudwatch.IWidget) error {
	if widget == nil {
		return fmt.Errorf("parameter widget is required, but nil was provided")
	}

	return nil
}

func (m *jsiiProxy_MonitoringFacade) validateCreateAlarmFactoryParameters(alarmNamePrefix *string) error {
	if alarmNamePrefix == nil {
		return fmt.Errorf("parameter alarmNamePrefix is required, but nil was provided")
	}

	return nil
}

func (m *jsiiProxy_MonitoringFacade) validateCreateCompositeAlarmUsingDisambiguatorParameters(alarmDisambiguator *string, props *AddCompositeAlarmProps) error {
	if alarmDisambiguator == nil {
		return fmt.Errorf("parameter alarmDisambiguator is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}

func (m *jsiiProxy_MonitoringFacade) validateCreateCompositeAlarmUsingTagParameters(customTag *string, props *AddCompositeAlarmProps) error {
	if customTag == nil {
		return fmt.Errorf("parameter customTag is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}

func (m *jsiiProxy_MonitoringFacade) validateCreatedAlarmsWithDisambiguatorParameters(disambiguator *string) error {
	if disambiguator == nil {
		return fmt.Errorf("parameter disambiguator is required, but nil was provided")
	}

	return nil
}

func (m *jsiiProxy_MonitoringFacade) validateCreatedAlarmsWithTagParameters(customTag *string) error {
	if customTag == nil {
		return fmt.Errorf("parameter customTag is required, but nil was provided")
	}

	return nil
}

func (m *jsiiProxy_MonitoringFacade) validateMonitorApiGatewayParameters(props *ApiGatewayMonitoringProps) error {
	if props == nil {
		return fmt.Errorf("parameter props is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}

func (m *jsiiProxy_MonitoringFacade) validateMonitorApiGatewayV2HttpApiParameters(props *ApiGatewayV2HttpApiMonitoringProps) error {
	if props == nil {
		return fmt.Errorf("parameter props is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}

func (m *jsiiProxy_MonitoringFacade) validateMonitorAppSyncApiParameters(props *AppSyncMonitoringProps) error {
	if props == nil {
		return fmt.Errorf("parameter props is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}

func (m *jsiiProxy_MonitoringFacade) validateMonitorAuroraClusterParameters(props *AuroraClusterMonitoringProps) error {
	if props == nil {
		return fmt.Errorf("parameter props is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}

func (m *jsiiProxy_MonitoringFacade) validateMonitorAutoScalingGroupParameters(props *AutoScalingGroupMonitoringProps) error {
	if props == nil {
		return fmt.Errorf("parameter props is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}

func (m *jsiiProxy_MonitoringFacade) validateMonitorBillingParameters(props *BillingMonitoringProps) error {
	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}

func (m *jsiiProxy_MonitoringFacade) validateMonitorCertificateParameters(props *CertificateManagerMonitoringProps) error {
	if props == nil {
		return fmt.Errorf("parameter props is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}

func (m *jsiiProxy_MonitoringFacade) validateMonitorCloudFrontDistributionParameters(props *CloudFrontDistributionMonitoringProps) error {
	if props == nil {
		return fmt.Errorf("parameter props is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}

func (m *jsiiProxy_MonitoringFacade) validateMonitorCodeBuildProjectParameters(props *CodeBuildProjectMonitoringProps) error {
	if props == nil {
		return fmt.Errorf("parameter props is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}

func (m *jsiiProxy_MonitoringFacade) validateMonitorCustomParameters(props *CustomMonitoringProps) error {
	if props == nil {
		return fmt.Errorf("parameter props is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}

func (m *jsiiProxy_MonitoringFacade) validateMonitorDocumentDbClusterParameters(props *DocumentDbMonitoringProps) error {
	if props == nil {
		return fmt.Errorf("parameter props is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}

func (m *jsiiProxy_MonitoringFacade) validateMonitorDynamoTableParameters(props *DynamoTableMonitoringProps) error {
	if props == nil {
		return fmt.Errorf("parameter props is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}

func (m *jsiiProxy_MonitoringFacade) validateMonitorDynamoTableGlobalSecondaryIndexParameters(props *DynamoTableGlobalSecondaryIndexMonitoringProps) error {
	if props == nil {
		return fmt.Errorf("parameter props is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}

func (m *jsiiProxy_MonitoringFacade) validateMonitorEc2ApplicationLoadBalancerParameters(props *Ec2ApplicationLoadBalancerMonitoringProps) error {
	if props == nil {
		return fmt.Errorf("parameter props is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}

func (m *jsiiProxy_MonitoringFacade) validateMonitorEC2InstancesParameters(props *EC2MonitoringProps) error {
	if props == nil {
		return fmt.Errorf("parameter props is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}

func (m *jsiiProxy_MonitoringFacade) validateMonitorEc2NetworkLoadBalancerParameters(props *Ec2NetworkLoadBalancerMonitoringProps) error {
	if props == nil {
		return fmt.Errorf("parameter props is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}

func (m *jsiiProxy_MonitoringFacade) validateMonitorEc2ServiceParameters(props *Ec2ServiceMonitoringProps) error {
	if props == nil {
		return fmt.Errorf("parameter props is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}

func (m *jsiiProxy_MonitoringFacade) validateMonitorElastiCacheClusterParameters(props *ElastiCacheClusterMonitoringProps) error {
	if props == nil {
		return fmt.Errorf("parameter props is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}

func (m *jsiiProxy_MonitoringFacade) validateMonitorElasticsearchClusterParameters(props *OpenSearchClusterMonitoringProps) error {
	if props == nil {
		return fmt.Errorf("parameter props is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}

func (m *jsiiProxy_MonitoringFacade) validateMonitorFargateApplicationLoadBalancerParameters(props *FargateApplicationLoadBalancerMonitoringProps) error {
	if props == nil {
		return fmt.Errorf("parameter props is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}

func (m *jsiiProxy_MonitoringFacade) validateMonitorFargateNetworkLoadBalancerParameters(props *FargateNetworkLoadBalancerMonitoringProps) error {
	if props == nil {
		return fmt.Errorf("parameter props is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}

func (m *jsiiProxy_MonitoringFacade) validateMonitorFargateServiceParameters(props *FargateServiceMonitoringProps) error {
	if props == nil {
		return fmt.Errorf("parameter props is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}

func (m *jsiiProxy_MonitoringFacade) validateMonitorGlueJobParameters(props *GlueJobMonitoringProps) error {
	if props == nil {
		return fmt.Errorf("parameter props is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}

func (m *jsiiProxy_MonitoringFacade) validateMonitorKinesisDataAnalyticsParameters(props *KinesisDataAnalyticsMonitoringProps) error {
	if props == nil {
		return fmt.Errorf("parameter props is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}

func (m *jsiiProxy_MonitoringFacade) validateMonitorKinesisDataStreamParameters(props *KinesisDataStreamMonitoringProps) error {
	if props == nil {
		return fmt.Errorf("parameter props is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}

func (m *jsiiProxy_MonitoringFacade) validateMonitorKinesisFirehoseParameters(props *KinesisFirehoseMonitoringProps) error {
	if props == nil {
		return fmt.Errorf("parameter props is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}

func (m *jsiiProxy_MonitoringFacade) validateMonitorLambdaFunctionParameters(props *LambdaFunctionMonitoringProps) error {
	if props == nil {
		return fmt.Errorf("parameter props is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}

func (m *jsiiProxy_MonitoringFacade) validateMonitorLogParameters(props *LogMonitoringProps) error {
	if props == nil {
		return fmt.Errorf("parameter props is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}

func (m *jsiiProxy_MonitoringFacade) validateMonitorNetworkLoadBalancerParameters(props *NetworkLoadBalancerMonitoringProps) error {
	if props == nil {
		return fmt.Errorf("parameter props is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}

func (m *jsiiProxy_MonitoringFacade) validateMonitorOpenSearchClusterParameters(props *OpenSearchClusterMonitoringProps) error {
	if props == nil {
		return fmt.Errorf("parameter props is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}

func (m *jsiiProxy_MonitoringFacade) validateMonitorQueueProcessingEc2ServiceParameters(props *QueueProcessingEc2ServiceMonitoringProps) error {
	if props == nil {
		return fmt.Errorf("parameter props is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}

func (m *jsiiProxy_MonitoringFacade) validateMonitorQueueProcessingFargateServiceParameters(props *QueueProcessingFargateServiceMonitoringProps) error {
	if props == nil {
		return fmt.Errorf("parameter props is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}

func (m *jsiiProxy_MonitoringFacade) validateMonitorRdsClusterParameters(props *RdsClusterMonitoringProps) error {
	if props == nil {
		return fmt.Errorf("parameter props is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}

func (m *jsiiProxy_MonitoringFacade) validateMonitorRedshiftClusterParameters(props *RedshiftClusterMonitoringProps) error {
	if props == nil {
		return fmt.Errorf("parameter props is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}

func (m *jsiiProxy_MonitoringFacade) validateMonitorS3BucketParameters(props *S3BucketMonitoringProps) error {
	if props == nil {
		return fmt.Errorf("parameter props is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}

func (m *jsiiProxy_MonitoringFacade) validateMonitorScopeParameters(scope constructs.Construct, aspectProps *MonitoringAspectProps) error {
	if scope == nil {
		return fmt.Errorf("parameter scope is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(aspectProps, func() string { return "parameter aspectProps" }); err != nil {
		return err
	}

	return nil
}

func (m *jsiiProxy_MonitoringFacade) validateMonitorSecretsManagerParameters(props *SecretsManagerMonitoringProps) error {
	if props == nil {
		return fmt.Errorf("parameter props is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}

func (m *jsiiProxy_MonitoringFacade) validateMonitorSecretsManagerSecretParameters(props *SecretsManagerSecretMonitoringProps) error {
	if props == nil {
		return fmt.Errorf("parameter props is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}

func (m *jsiiProxy_MonitoringFacade) validateMonitorSimpleEc2ServiceParameters(props *SimpleEc2ServiceMonitoringProps) error {
	if props == nil {
		return fmt.Errorf("parameter props is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}

func (m *jsiiProxy_MonitoringFacade) validateMonitorSimpleFargateServiceParameters(props *SimpleFargateServiceMonitoringProps) error {
	if props == nil {
		return fmt.Errorf("parameter props is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}

func (m *jsiiProxy_MonitoringFacade) validateMonitorSnsTopicParameters(props *SnsTopicMonitoringProps) error {
	if props == nil {
		return fmt.Errorf("parameter props is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}

func (m *jsiiProxy_MonitoringFacade) validateMonitorSqsQueueParameters(props *SqsQueueMonitoringProps) error {
	if props == nil {
		return fmt.Errorf("parameter props is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}

func (m *jsiiProxy_MonitoringFacade) validateMonitorSqsQueueWithDlqParameters(props *SqsQueueMonitoringWithDlqProps) error {
	if props == nil {
		return fmt.Errorf("parameter props is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}

func (m *jsiiProxy_MonitoringFacade) validateMonitorStepFunctionParameters(props *StepFunctionMonitoringProps) error {
	if props == nil {
		return fmt.Errorf("parameter props is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}

func (m *jsiiProxy_MonitoringFacade) validateMonitorStepFunctionActivityParameters(props *StepFunctionActivityMonitoringProps) error {
	if props == nil {
		return fmt.Errorf("parameter props is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}

func (m *jsiiProxy_MonitoringFacade) validateMonitorStepFunctionLambdaIntegrationParameters(props *StepFunctionLambdaIntegrationMonitoringProps) error {
	if props == nil {
		return fmt.Errorf("parameter props is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}

func (m *jsiiProxy_MonitoringFacade) validateMonitorStepFunctionServiceIntegrationParameters(props *StepFunctionServiceIntegrationMonitoringProps) error {
	if props == nil {
		return fmt.Errorf("parameter props is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}

func (m *jsiiProxy_MonitoringFacade) validateMonitorSyntheticsCanaryParameters(props *SyntheticsCanaryMonitoringProps) error {
	if props == nil {
		return fmt.Errorf("parameter props is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}

func (m *jsiiProxy_MonitoringFacade) validateMonitorWebApplicationFirewallAclV2Parameters(props *WafV2MonitoringProps) error {
	if props == nil {
		return fmt.Errorf("parameter props is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}

func validateMonitoringFacade_IsConstructParameters(x interface{}) error {
	if x == nil {
		return fmt.Errorf("parameter x is required, but nil was provided")
	}

	return nil
}

func validateNewMonitoringFacadeParameters(scope constructs.Construct, id *string, props *MonitoringFacadeProps) error {
	if scope == nil {
		return fmt.Errorf("parameter scope is required, but nil was provided")
	}

	if id == nil {
		return fmt.Errorf("parameter id is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}


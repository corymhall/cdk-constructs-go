//go:build !no_runtime_type_checking

package cdkmonitoringconstructs

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
)

func (a *jsiiProxy_AwsConsoleUrlFactory) validateGetApiGatewayUrlParameters(restApiId *string) error {
	if restApiId == nil {
		return fmt.Errorf("parameter restApiId is required, but nil was provided")
	}

	return nil
}

func (a *jsiiProxy_AwsConsoleUrlFactory) validateGetCloudFrontDistributionUrlParameters(distributionId *string) error {
	if distributionId == nil {
		return fmt.Errorf("parameter distributionId is required, but nil was provided")
	}

	return nil
}

func (a *jsiiProxy_AwsConsoleUrlFactory) validateGetCloudWatchLogGroupUrlParameters(logGroupName *string) error {
	if logGroupName == nil {
		return fmt.Errorf("parameter logGroupName is required, but nil was provided")
	}

	return nil
}

func (a *jsiiProxy_AwsConsoleUrlFactory) validateGetCodeBuildProjectUrlParameters(projectName *string) error {
	if projectName == nil {
		return fmt.Errorf("parameter projectName is required, but nil was provided")
	}

	return nil
}

func (a *jsiiProxy_AwsConsoleUrlFactory) validateGetDocumentDbClusterUrlParameters(clusterId *string) error {
	if clusterId == nil {
		return fmt.Errorf("parameter clusterId is required, but nil was provided")
	}

	return nil
}

func (a *jsiiProxy_AwsConsoleUrlFactory) validateGetDynamoTableUrlParameters(tableName *string) error {
	if tableName == nil {
		return fmt.Errorf("parameter tableName is required, but nil was provided")
	}

	return nil
}

func (a *jsiiProxy_AwsConsoleUrlFactory) validateGetElastiCacheClusterUrlParameters(clusterId *string, clusterType ElastiCacheClusterType) error {
	if clusterId == nil {
		return fmt.Errorf("parameter clusterId is required, but nil was provided")
	}

	if clusterType == "" {
		return fmt.Errorf("parameter clusterType is required, but nil was provided")
	}

	return nil
}

func (a *jsiiProxy_AwsConsoleUrlFactory) validateGetKinesisAnalyticsUrlParameters(application *string) error {
	if application == nil {
		return fmt.Errorf("parameter application is required, but nil was provided")
	}

	return nil
}

func (a *jsiiProxy_AwsConsoleUrlFactory) validateGetKinesisDataStreamUrlParameters(streamName *string) error {
	if streamName == nil {
		return fmt.Errorf("parameter streamName is required, but nil was provided")
	}

	return nil
}

func (a *jsiiProxy_AwsConsoleUrlFactory) validateGetKinesisFirehoseDeliveryStreamUrlParameters(streamName *string) error {
	if streamName == nil {
		return fmt.Errorf("parameter streamName is required, but nil was provided")
	}

	return nil
}

func (a *jsiiProxy_AwsConsoleUrlFactory) validateGetLambdaFunctionUrlParameters(functionName *string) error {
	if functionName == nil {
		return fmt.Errorf("parameter functionName is required, but nil was provided")
	}

	return nil
}

func (a *jsiiProxy_AwsConsoleUrlFactory) validateGetOpenSearchClusterUrlParameters(domainName *string) error {
	if domainName == nil {
		return fmt.Errorf("parameter domainName is required, but nil was provided")
	}

	return nil
}

func (a *jsiiProxy_AwsConsoleUrlFactory) validateGetRdsClusterUrlParameters(clusterId *string) error {
	if clusterId == nil {
		return fmt.Errorf("parameter clusterId is required, but nil was provided")
	}

	return nil
}

func (a *jsiiProxy_AwsConsoleUrlFactory) validateGetRedshiftClusterUrlParameters(clusterId *string) error {
	if clusterId == nil {
		return fmt.Errorf("parameter clusterId is required, but nil was provided")
	}

	return nil
}

func (a *jsiiProxy_AwsConsoleUrlFactory) validateGetResolvedDestinationUrlParameters(context awscdk.IResolveContext, destinationUrl *string) error {
	if context == nil {
		return fmt.Errorf("parameter context is required, but nil was provided")
	}

	if destinationUrl == nil {
		return fmt.Errorf("parameter destinationUrl is required, but nil was provided")
	}

	return nil
}

func (a *jsiiProxy_AwsConsoleUrlFactory) validateGetS3BucketUrlParameters(bucketName *string) error {
	if bucketName == nil {
		return fmt.Errorf("parameter bucketName is required, but nil was provided")
	}

	return nil
}

func (a *jsiiProxy_AwsConsoleUrlFactory) validateGetSnsTopicUrlParameters(topicArn *string) error {
	if topicArn == nil {
		return fmt.Errorf("parameter topicArn is required, but nil was provided")
	}

	return nil
}

func (a *jsiiProxy_AwsConsoleUrlFactory) validateGetSqsQueueUrlParameters(queueUrl *string) error {
	if queueUrl == nil {
		return fmt.Errorf("parameter queueUrl is required, but nil was provided")
	}

	return nil
}

func (a *jsiiProxy_AwsConsoleUrlFactory) validateGetStateMachineUrlParameters(stateMachineArn *string) error {
	if stateMachineArn == nil {
		return fmt.Errorf("parameter stateMachineArn is required, but nil was provided")
	}

	return nil
}

func validateNewAwsConsoleUrlFactoryParameters(props *AwsConsoleUrlFactoryProps) error {
	if props == nil {
		return fmt.Errorf("parameter props is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}


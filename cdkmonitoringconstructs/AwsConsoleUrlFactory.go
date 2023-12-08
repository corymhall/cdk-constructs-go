package cdkmonitoringconstructs

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/corymhall/cdk-constructs-go/cdkmonitoringconstructs/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Experimental.
type AwsConsoleUrlFactory interface {
	// Experimental.
	AwsAccountId() *string
	// Experimental.
	AwsAccountRegion() *string
	// Experimental.
	GetApiGatewayUrl(restApiId *string) *string
	// Experimental.
	GetAwsConsoleUrl(destinationUrl *string) *string
	// Experimental.
	GetCloudFrontDistributionUrl(distributionId *string) *string
	// Experimental.
	GetCloudWatchLogGroupUrl(logGroupName *string) *string
	// Experimental.
	GetCodeBuildProjectUrl(projectName *string) *string
	// Experimental.
	GetDocumentDbClusterUrl(clusterId *string) *string
	// Experimental.
	GetDynamoTableUrl(tableName *string) *string
	// Experimental.
	GetElastiCacheClusterUrl(clusterId *string, clusterType ElastiCacheClusterType) *string
	// Experimental.
	GetKinesisAnalyticsUrl(application *string) *string
	// Experimental.
	GetKinesisDataStreamUrl(streamName *string) *string
	// Experimental.
	GetKinesisFirehoseDeliveryStreamUrl(streamName *string) *string
	// Experimental.
	GetLambdaFunctionUrl(functionName *string) *string
	// Experimental.
	GetOpenSearchClusterUrl(domainName *string) *string
	// Experimental.
	GetRdsClusterUrl(clusterId *string) *string
	// Experimental.
	GetRedshiftClusterUrl(clusterId *string) *string
	// Resolves a destination URL within a resolution context.
	// See: https://docs.aws.amazon.com/cdk/latest/guide/tokens.html
	//
	// Experimental.
	GetResolvedDestinationUrl(context awscdk.IResolveContext, destinationUrl *string) *string
	// Experimental.
	GetS3BucketUrl(bucketName *string) *string
	// Experimental.
	GetSnsTopicUrl(topicArn *string) *string
	// Experimental.
	GetSqsQueueUrl(queueUrl *string) *string
	// Experimental.
	GetStateMachineUrl(stateMachineArn *string) *string
}

// The jsii proxy struct for AwsConsoleUrlFactory
type jsiiProxy_AwsConsoleUrlFactory struct {
	_ byte // padding
}

func (j *jsiiProxy_AwsConsoleUrlFactory) AwsAccountId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"awsAccountId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConsoleUrlFactory) AwsAccountRegion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"awsAccountRegion",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsConsoleUrlFactory(props *AwsConsoleUrlFactoryProps) AwsConsoleUrlFactory {
	_init_.Initialize()

	if err := validateNewAwsConsoleUrlFactoryParameters(props); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsConsoleUrlFactory{}

	_jsii_.Create(
		"cdk-monitoring-constructs.AwsConsoleUrlFactory",
		[]interface{}{props},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsConsoleUrlFactory_Override(a AwsConsoleUrlFactory, props *AwsConsoleUrlFactoryProps) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk-monitoring-constructs.AwsConsoleUrlFactory",
		[]interface{}{props},
		a,
	)
}

func (a *jsiiProxy_AwsConsoleUrlFactory) GetApiGatewayUrl(restApiId *string) *string {
	if err := a.validateGetApiGatewayUrlParameters(restApiId); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		a,
		"getApiGatewayUrl",
		[]interface{}{restApiId},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsConsoleUrlFactory) GetAwsConsoleUrl(destinationUrl *string) *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"getAwsConsoleUrl",
		[]interface{}{destinationUrl},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsConsoleUrlFactory) GetCloudFrontDistributionUrl(distributionId *string) *string {
	if err := a.validateGetCloudFrontDistributionUrlParameters(distributionId); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		a,
		"getCloudFrontDistributionUrl",
		[]interface{}{distributionId},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsConsoleUrlFactory) GetCloudWatchLogGroupUrl(logGroupName *string) *string {
	if err := a.validateGetCloudWatchLogGroupUrlParameters(logGroupName); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		a,
		"getCloudWatchLogGroupUrl",
		[]interface{}{logGroupName},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsConsoleUrlFactory) GetCodeBuildProjectUrl(projectName *string) *string {
	if err := a.validateGetCodeBuildProjectUrlParameters(projectName); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		a,
		"getCodeBuildProjectUrl",
		[]interface{}{projectName},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsConsoleUrlFactory) GetDocumentDbClusterUrl(clusterId *string) *string {
	if err := a.validateGetDocumentDbClusterUrlParameters(clusterId); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		a,
		"getDocumentDbClusterUrl",
		[]interface{}{clusterId},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsConsoleUrlFactory) GetDynamoTableUrl(tableName *string) *string {
	if err := a.validateGetDynamoTableUrlParameters(tableName); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		a,
		"getDynamoTableUrl",
		[]interface{}{tableName},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsConsoleUrlFactory) GetElastiCacheClusterUrl(clusterId *string, clusterType ElastiCacheClusterType) *string {
	if err := a.validateGetElastiCacheClusterUrlParameters(clusterId, clusterType); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		a,
		"getElastiCacheClusterUrl",
		[]interface{}{clusterId, clusterType},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsConsoleUrlFactory) GetKinesisAnalyticsUrl(application *string) *string {
	if err := a.validateGetKinesisAnalyticsUrlParameters(application); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		a,
		"getKinesisAnalyticsUrl",
		[]interface{}{application},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsConsoleUrlFactory) GetKinesisDataStreamUrl(streamName *string) *string {
	if err := a.validateGetKinesisDataStreamUrlParameters(streamName); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		a,
		"getKinesisDataStreamUrl",
		[]interface{}{streamName},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsConsoleUrlFactory) GetKinesisFirehoseDeliveryStreamUrl(streamName *string) *string {
	if err := a.validateGetKinesisFirehoseDeliveryStreamUrlParameters(streamName); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		a,
		"getKinesisFirehoseDeliveryStreamUrl",
		[]interface{}{streamName},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsConsoleUrlFactory) GetLambdaFunctionUrl(functionName *string) *string {
	if err := a.validateGetLambdaFunctionUrlParameters(functionName); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		a,
		"getLambdaFunctionUrl",
		[]interface{}{functionName},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsConsoleUrlFactory) GetOpenSearchClusterUrl(domainName *string) *string {
	if err := a.validateGetOpenSearchClusterUrlParameters(domainName); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		a,
		"getOpenSearchClusterUrl",
		[]interface{}{domainName},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsConsoleUrlFactory) GetRdsClusterUrl(clusterId *string) *string {
	if err := a.validateGetRdsClusterUrlParameters(clusterId); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		a,
		"getRdsClusterUrl",
		[]interface{}{clusterId},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsConsoleUrlFactory) GetRedshiftClusterUrl(clusterId *string) *string {
	if err := a.validateGetRedshiftClusterUrlParameters(clusterId); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		a,
		"getRedshiftClusterUrl",
		[]interface{}{clusterId},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsConsoleUrlFactory) GetResolvedDestinationUrl(context awscdk.IResolveContext, destinationUrl *string) *string {
	if err := a.validateGetResolvedDestinationUrlParameters(context, destinationUrl); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		a,
		"getResolvedDestinationUrl",
		[]interface{}{context, destinationUrl},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsConsoleUrlFactory) GetS3BucketUrl(bucketName *string) *string {
	if err := a.validateGetS3BucketUrlParameters(bucketName); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		a,
		"getS3BucketUrl",
		[]interface{}{bucketName},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsConsoleUrlFactory) GetSnsTopicUrl(topicArn *string) *string {
	if err := a.validateGetSnsTopicUrlParameters(topicArn); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		a,
		"getSnsTopicUrl",
		[]interface{}{topicArn},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsConsoleUrlFactory) GetSqsQueueUrl(queueUrl *string) *string {
	if err := a.validateGetSqsQueueUrlParameters(queueUrl); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		a,
		"getSqsQueueUrl",
		[]interface{}{queueUrl},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsConsoleUrlFactory) GetStateMachineUrl(stateMachineArn *string) *string {
	if err := a.validateGetStateMachineUrlParameters(stateMachineArn); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		a,
		"getStateMachineUrl",
		[]interface{}{stateMachineArn},
		&returns,
	)

	return returns
}


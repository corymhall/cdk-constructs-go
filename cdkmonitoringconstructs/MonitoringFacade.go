package cdkmonitoringconstructs

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/corymhall/cdk-constructs-go/cdkmonitoringconstructs/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudwatch"
	"github.com/aws/constructs-go/constructs/v10"
)

// An implementation of a {@link MonitoringScope}.
//
// This is a convenient main entrypoint to monitor resources.
//
// Provides methods for retrieving and creating alarms based on added segments that are subclasses of
// {@link Monitoring}.
// Experimental.
type MonitoringFacade interface {
	MonitoringScope
	// Experimental.
	AlarmFactoryDefaults() *AlarmFactoryDefaults
	// Experimental.
	CreatedSegments() *[]interface{}
	// Experimental.
	DashboardFactory() IDynamicDashboardFactory
	// Experimental.
	MetricFactoryDefaults() *MetricFactoryDefaults
	// The tree node.
	// Experimental.
	Node() constructs.Node
	// Adds a dashboard segment which returns dynamic content depending on dashboard type.
	// Experimental.
	AddDynamicSegment(segment IDynamicDashboardSegment) MonitoringFacade
	// Experimental.
	AddLargeHeader(text *string, addToSummary *bool, addToAlarm *bool) MonitoringFacade
	// Experimental.
	AddMediumHeader(text *string, addToSummary *bool, addToAlarm *bool) MonitoringFacade
	// Adds a dashboard segment to go on one of the {@link DefaultDashboards}.
	// Experimental.
	AddSegment(segment IDashboardSegment, overrideProps *MonitoringDashboardsOverrideProps) MonitoringFacade
	// Experimental.
	AddSmallHeader(text *string, addToSummary *bool, addToAlarm *bool) MonitoringFacade
	// Experimental.
	AddWidget(widget awscloudwatch.IWidget, addToSummary *bool, addToAlarm *bool) MonitoringFacade
	// Creates a new alarm factory.
	//
	// Alarms created will be named with the given prefix, unless a local name override is present.
	// Experimental.
	CreateAlarmFactory(alarmNamePrefix *string) AlarmFactory
	// Creates a new factory that creates AWS Console URLs.
	// Experimental.
	CreateAwsConsoleUrlFactory() AwsConsoleUrlFactory
	// Finds a subset of created alarms that are marked by a specific disambiguator and creates a composite alarm.
	//
	// This composite alarm is created with an 'OR' condition, so it triggers with any child alarm.
	// NOTE: This composite alarm is not added among other alarms, so it is not returned by createdAlarms() calls.
	// Experimental.
	CreateCompositeAlarmUsingDisambiguator(alarmDisambiguator *string, props *AddCompositeAlarmProps) awscloudwatch.CompositeAlarm
	// Finds a subset of created alarms that are marked by a specific custom tag and creates a composite alarm.
	//
	// This composite alarm is created with an 'OR' condition, so it triggers with any child alarm.
	// NOTE: This composite alarm is not added among other alarms, so it is not returned by createdAlarms() calls.
	// Experimental.
	CreateCompositeAlarmUsingTag(customTag *string, props *AddCompositeAlarmProps) awscloudwatch.CompositeAlarm
	// Returns: default alarms dashboard.
	// Deprecated: - prefer calling dashboardFactory.getDashboard directly.
	CreatedAlarmDashboard() awscloudwatch.Dashboard
	// Returns the created alarms across all added segments that subclass {@link Monitoring} added up until now.
	// Experimental.
	CreatedAlarms() *[]*AlarmWithAnnotation
	// Returns a subset of created alarms that are marked by a specific disambiguator.
	// Experimental.
	CreatedAlarmsWithDisambiguator(disambiguator *string) *[]*AlarmWithAnnotation
	// Returns a subset of created alarms that are marked by a specific custom tag.
	// Experimental.
	CreatedAlarmsWithTag(customTag *string) *[]*AlarmWithAnnotation
	// Returns: default detail dashboard.
	// Deprecated: - prefer calling dashboardFactory.getDashboard directly.
	CreatedDashboard() awscloudwatch.Dashboard
	// Returns the added segments that subclass {@link Monitoring}.
	// Experimental.
	CreatedMonitorings() *[]Monitoring
	// Returns: default summary dashboard.
	// Deprecated: - prefer calling dashboardFactory.getDashboard directly.
	CreatedSummaryDashboard() awscloudwatch.Dashboard
	// Creates a new metric factory.
	// Experimental.
	CreateMetricFactory() MetricFactory
	// Creates a new widget factory.
	// Experimental.
	CreateWidgetFactory() IWidgetFactory
	// Experimental.
	MonitorApiGateway(props *ApiGatewayMonitoringProps) MonitoringFacade
	// Experimental.
	MonitorApiGatewayV2HttpApi(props *ApiGatewayV2HttpApiMonitoringProps) MonitoringFacade
	// Experimental.
	MonitorAppSyncApi(props *AppSyncMonitoringProps) MonitoringFacade
	// Experimental.
	MonitorAuroraCluster(props *AuroraClusterMonitoringProps) MonitoringFacade
	// Experimental.
	MonitorAutoScalingGroup(props *AutoScalingGroupMonitoringProps) MonitoringFacade
	// Experimental.
	MonitorBilling(props *BillingMonitoringProps) MonitoringFacade
	// Experimental.
	MonitorCertificate(props *CertificateManagerMonitoringProps) MonitoringFacade
	// Experimental.
	MonitorCloudFrontDistribution(props *CloudFrontDistributionMonitoringProps) MonitoringFacade
	// Experimental.
	MonitorCodeBuildProject(props *CodeBuildProjectMonitoringProps) MonitoringFacade
	// Experimental.
	MonitorCustom(props *CustomMonitoringProps) MonitoringFacade
	// Experimental.
	MonitorDocumentDbCluster(props *DocumentDbMonitoringProps) MonitoringFacade
	// Experimental.
	MonitorDynamoTable(props *DynamoTableMonitoringProps) MonitoringFacade
	// Experimental.
	MonitorDynamoTableGlobalSecondaryIndex(props *DynamoTableGlobalSecondaryIndexMonitoringProps) MonitoringFacade
	// Experimental.
	MonitorEc2ApplicationLoadBalancer(props *Ec2ApplicationLoadBalancerMonitoringProps) MonitoringFacade
	// Experimental.
	MonitorEC2Instances(props *EC2MonitoringProps) MonitoringFacade
	// Experimental.
	MonitorEc2NetworkLoadBalancer(props *Ec2NetworkLoadBalancerMonitoringProps) MonitoringFacade
	// Experimental.
	MonitorEc2Service(props *Ec2ServiceMonitoringProps) MonitoringFacade
	// Experimental.
	MonitorElastiCacheCluster(props *ElastiCacheClusterMonitoringProps) MonitoringFacade
	// Experimental.
	MonitorElasticsearchCluster(props *OpenSearchClusterMonitoringProps) MonitoringFacade
	// Experimental.
	MonitorFargateApplicationLoadBalancer(props *FargateApplicationLoadBalancerMonitoringProps) MonitoringFacade
	// Experimental.
	MonitorFargateNetworkLoadBalancer(props *FargateNetworkLoadBalancerMonitoringProps) MonitoringFacade
	// Experimental.
	MonitorFargateService(props *FargateServiceMonitoringProps) MonitoringFacade
	// Experimental.
	MonitorGlueJob(props *GlueJobMonitoringProps) MonitoringFacade
	// Experimental.
	MonitorKinesisDataAnalytics(props *KinesisDataAnalyticsMonitoringProps) MonitoringFacade
	// Experimental.
	MonitorKinesisDataStream(props *KinesisDataStreamMonitoringProps) MonitoringFacade
	// Experimental.
	MonitorKinesisFirehose(props *KinesisFirehoseMonitoringProps) MonitoringFacade
	// Experimental.
	MonitorLambdaFunction(props *LambdaFunctionMonitoringProps) MonitoringFacade
	// Experimental.
	MonitorLog(props *LogMonitoringProps) MonitoringFacade
	// Experimental.
	MonitorNetworkLoadBalancer(props *NetworkLoadBalancerMonitoringProps) MonitoringFacade
	// Experimental.
	MonitorOpenSearchCluster(props *OpenSearchClusterMonitoringProps) MonitoringFacade
	// Experimental.
	MonitorQueueProcessingEc2Service(props *QueueProcessingEc2ServiceMonitoringProps) MonitoringFacade
	// Experimental.
	MonitorQueueProcessingFargateService(props *QueueProcessingFargateServiceMonitoringProps) MonitoringFacade
	// Experimental.
	MonitorRdsCluster(props *RdsClusterMonitoringProps) MonitoringFacade
	// Experimental.
	MonitorRedshiftCluster(props *RedshiftClusterMonitoringProps) MonitoringFacade
	// Experimental.
	MonitorS3Bucket(props *S3BucketMonitoringProps) MonitoringFacade
	// Uses an aspect to automatically monitor all resources in the given scope.
	// Experimental.
	MonitorScope(scope constructs.Construct, aspectProps *MonitoringAspectProps)
	// Experimental.
	MonitorSecretsManager(props *SecretsManagerMonitoringProps) MonitoringFacade
	// Experimental.
	MonitorSecretsManagerSecret(props *SecretsManagerSecretMonitoringProps) MonitoringFacade
	// Experimental.
	MonitorSimpleEc2Service(props *SimpleEc2ServiceMonitoringProps) MonitoringFacade
	// Experimental.
	MonitorSimpleFargateService(props *SimpleFargateServiceMonitoringProps) MonitoringFacade
	// Experimental.
	MonitorSnsTopic(props *SnsTopicMonitoringProps) MonitoringFacade
	// Experimental.
	MonitorSqsQueue(props *SqsQueueMonitoringProps) MonitoringFacade
	// Experimental.
	MonitorSqsQueueWithDlq(props *SqsQueueMonitoringWithDlqProps) MonitoringFacade
	// Experimental.
	MonitorStepFunction(props *StepFunctionMonitoringProps) MonitoringFacade
	// Experimental.
	MonitorStepFunctionActivity(props *StepFunctionActivityMonitoringProps) MonitoringFacade
	// Experimental.
	MonitorStepFunctionLambdaIntegration(props *StepFunctionLambdaIntegrationMonitoringProps) MonitoringFacade
	// Experimental.
	MonitorStepFunctionServiceIntegration(props *StepFunctionServiceIntegrationMonitoringProps) MonitoringFacade
	// Experimental.
	MonitorSyntheticsCanary(props *SyntheticsCanaryMonitoringProps) MonitoringFacade
	// Experimental.
	MonitorWebApplicationFirewallAclV2(props *WafV2MonitoringProps) MonitoringFacade
	// Returns a string representation of this construct.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for MonitoringFacade
type jsiiProxy_MonitoringFacade struct {
	jsiiProxy_MonitoringScope
}

func (j *jsiiProxy_MonitoringFacade) AlarmFactoryDefaults() *AlarmFactoryDefaults {
	var returns *AlarmFactoryDefaults
	_jsii_.Get(
		j,
		"alarmFactoryDefaults",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitoringFacade) CreatedSegments() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"createdSegments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitoringFacade) DashboardFactory() IDynamicDashboardFactory {
	var returns IDynamicDashboardFactory
	_jsii_.Get(
		j,
		"dashboardFactory",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitoringFacade) MetricFactoryDefaults() *MetricFactoryDefaults {
	var returns *MetricFactoryDefaults
	_jsii_.Get(
		j,
		"metricFactoryDefaults",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitoringFacade) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}


// Experimental.
func NewMonitoringFacade(scope constructs.Construct, id *string, props *MonitoringFacadeProps) MonitoringFacade {
	_init_.Initialize()

	if err := validateNewMonitoringFacadeParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_MonitoringFacade{}

	_jsii_.Create(
		"cdk-monitoring-constructs.MonitoringFacade",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

// Experimental.
func NewMonitoringFacade_Override(m MonitoringFacade, scope constructs.Construct, id *string, props *MonitoringFacadeProps) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk-monitoring-constructs.MonitoringFacade",
		[]interface{}{scope, id, props},
		m,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead.
func MonitoringFacade_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateMonitoringFacade_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"cdk-monitoring-constructs.MonitoringFacade",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitoringFacade) AddDynamicSegment(segment IDynamicDashboardSegment) MonitoringFacade {
	if err := m.validateAddDynamicSegmentParameters(segment); err != nil {
		panic(err)
	}
	var returns MonitoringFacade

	_jsii_.Invoke(
		m,
		"addDynamicSegment",
		[]interface{}{segment},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitoringFacade) AddLargeHeader(text *string, addToSummary *bool, addToAlarm *bool) MonitoringFacade {
	if err := m.validateAddLargeHeaderParameters(text); err != nil {
		panic(err)
	}
	var returns MonitoringFacade

	_jsii_.Invoke(
		m,
		"addLargeHeader",
		[]interface{}{text, addToSummary, addToAlarm},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitoringFacade) AddMediumHeader(text *string, addToSummary *bool, addToAlarm *bool) MonitoringFacade {
	if err := m.validateAddMediumHeaderParameters(text); err != nil {
		panic(err)
	}
	var returns MonitoringFacade

	_jsii_.Invoke(
		m,
		"addMediumHeader",
		[]interface{}{text, addToSummary, addToAlarm},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitoringFacade) AddSegment(segment IDashboardSegment, overrideProps *MonitoringDashboardsOverrideProps) MonitoringFacade {
	if err := m.validateAddSegmentParameters(segment, overrideProps); err != nil {
		panic(err)
	}
	var returns MonitoringFacade

	_jsii_.Invoke(
		m,
		"addSegment",
		[]interface{}{segment, overrideProps},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitoringFacade) AddSmallHeader(text *string, addToSummary *bool, addToAlarm *bool) MonitoringFacade {
	if err := m.validateAddSmallHeaderParameters(text); err != nil {
		panic(err)
	}
	var returns MonitoringFacade

	_jsii_.Invoke(
		m,
		"addSmallHeader",
		[]interface{}{text, addToSummary, addToAlarm},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitoringFacade) AddWidget(widget awscloudwatch.IWidget, addToSummary *bool, addToAlarm *bool) MonitoringFacade {
	if err := m.validateAddWidgetParameters(widget); err != nil {
		panic(err)
	}
	var returns MonitoringFacade

	_jsii_.Invoke(
		m,
		"addWidget",
		[]interface{}{widget, addToSummary, addToAlarm},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitoringFacade) CreateAlarmFactory(alarmNamePrefix *string) AlarmFactory {
	if err := m.validateCreateAlarmFactoryParameters(alarmNamePrefix); err != nil {
		panic(err)
	}
	var returns AlarmFactory

	_jsii_.Invoke(
		m,
		"createAlarmFactory",
		[]interface{}{alarmNamePrefix},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitoringFacade) CreateAwsConsoleUrlFactory() AwsConsoleUrlFactory {
	var returns AwsConsoleUrlFactory

	_jsii_.Invoke(
		m,
		"createAwsConsoleUrlFactory",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitoringFacade) CreateCompositeAlarmUsingDisambiguator(alarmDisambiguator *string, props *AddCompositeAlarmProps) awscloudwatch.CompositeAlarm {
	if err := m.validateCreateCompositeAlarmUsingDisambiguatorParameters(alarmDisambiguator, props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.CompositeAlarm

	_jsii_.Invoke(
		m,
		"createCompositeAlarmUsingDisambiguator",
		[]interface{}{alarmDisambiguator, props},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitoringFacade) CreateCompositeAlarmUsingTag(customTag *string, props *AddCompositeAlarmProps) awscloudwatch.CompositeAlarm {
	if err := m.validateCreateCompositeAlarmUsingTagParameters(customTag, props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.CompositeAlarm

	_jsii_.Invoke(
		m,
		"createCompositeAlarmUsingTag",
		[]interface{}{customTag, props},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitoringFacade) CreatedAlarmDashboard() awscloudwatch.Dashboard {
	var returns awscloudwatch.Dashboard

	_jsii_.Invoke(
		m,
		"createdAlarmDashboard",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitoringFacade) CreatedAlarms() *[]*AlarmWithAnnotation {
	var returns *[]*AlarmWithAnnotation

	_jsii_.Invoke(
		m,
		"createdAlarms",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitoringFacade) CreatedAlarmsWithDisambiguator(disambiguator *string) *[]*AlarmWithAnnotation {
	if err := m.validateCreatedAlarmsWithDisambiguatorParameters(disambiguator); err != nil {
		panic(err)
	}
	var returns *[]*AlarmWithAnnotation

	_jsii_.Invoke(
		m,
		"createdAlarmsWithDisambiguator",
		[]interface{}{disambiguator},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitoringFacade) CreatedAlarmsWithTag(customTag *string) *[]*AlarmWithAnnotation {
	if err := m.validateCreatedAlarmsWithTagParameters(customTag); err != nil {
		panic(err)
	}
	var returns *[]*AlarmWithAnnotation

	_jsii_.Invoke(
		m,
		"createdAlarmsWithTag",
		[]interface{}{customTag},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitoringFacade) CreatedDashboard() awscloudwatch.Dashboard {
	var returns awscloudwatch.Dashboard

	_jsii_.Invoke(
		m,
		"createdDashboard",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitoringFacade) CreatedMonitorings() *[]Monitoring {
	var returns *[]Monitoring

	_jsii_.Invoke(
		m,
		"createdMonitorings",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitoringFacade) CreatedSummaryDashboard() awscloudwatch.Dashboard {
	var returns awscloudwatch.Dashboard

	_jsii_.Invoke(
		m,
		"createdSummaryDashboard",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitoringFacade) CreateMetricFactory() MetricFactory {
	var returns MetricFactory

	_jsii_.Invoke(
		m,
		"createMetricFactory",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitoringFacade) CreateWidgetFactory() IWidgetFactory {
	var returns IWidgetFactory

	_jsii_.Invoke(
		m,
		"createWidgetFactory",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitoringFacade) MonitorApiGateway(props *ApiGatewayMonitoringProps) MonitoringFacade {
	if err := m.validateMonitorApiGatewayParameters(props); err != nil {
		panic(err)
	}
	var returns MonitoringFacade

	_jsii_.Invoke(
		m,
		"monitorApiGateway",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitoringFacade) MonitorApiGatewayV2HttpApi(props *ApiGatewayV2HttpApiMonitoringProps) MonitoringFacade {
	if err := m.validateMonitorApiGatewayV2HttpApiParameters(props); err != nil {
		panic(err)
	}
	var returns MonitoringFacade

	_jsii_.Invoke(
		m,
		"monitorApiGatewayV2HttpApi",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitoringFacade) MonitorAppSyncApi(props *AppSyncMonitoringProps) MonitoringFacade {
	if err := m.validateMonitorAppSyncApiParameters(props); err != nil {
		panic(err)
	}
	var returns MonitoringFacade

	_jsii_.Invoke(
		m,
		"monitorAppSyncApi",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitoringFacade) MonitorAuroraCluster(props *AuroraClusterMonitoringProps) MonitoringFacade {
	if err := m.validateMonitorAuroraClusterParameters(props); err != nil {
		panic(err)
	}
	var returns MonitoringFacade

	_jsii_.Invoke(
		m,
		"monitorAuroraCluster",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitoringFacade) MonitorAutoScalingGroup(props *AutoScalingGroupMonitoringProps) MonitoringFacade {
	if err := m.validateMonitorAutoScalingGroupParameters(props); err != nil {
		panic(err)
	}
	var returns MonitoringFacade

	_jsii_.Invoke(
		m,
		"monitorAutoScalingGroup",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitoringFacade) MonitorBilling(props *BillingMonitoringProps) MonitoringFacade {
	if err := m.validateMonitorBillingParameters(props); err != nil {
		panic(err)
	}
	var returns MonitoringFacade

	_jsii_.Invoke(
		m,
		"monitorBilling",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitoringFacade) MonitorCertificate(props *CertificateManagerMonitoringProps) MonitoringFacade {
	if err := m.validateMonitorCertificateParameters(props); err != nil {
		panic(err)
	}
	var returns MonitoringFacade

	_jsii_.Invoke(
		m,
		"monitorCertificate",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitoringFacade) MonitorCloudFrontDistribution(props *CloudFrontDistributionMonitoringProps) MonitoringFacade {
	if err := m.validateMonitorCloudFrontDistributionParameters(props); err != nil {
		panic(err)
	}
	var returns MonitoringFacade

	_jsii_.Invoke(
		m,
		"monitorCloudFrontDistribution",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitoringFacade) MonitorCodeBuildProject(props *CodeBuildProjectMonitoringProps) MonitoringFacade {
	if err := m.validateMonitorCodeBuildProjectParameters(props); err != nil {
		panic(err)
	}
	var returns MonitoringFacade

	_jsii_.Invoke(
		m,
		"monitorCodeBuildProject",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitoringFacade) MonitorCustom(props *CustomMonitoringProps) MonitoringFacade {
	if err := m.validateMonitorCustomParameters(props); err != nil {
		panic(err)
	}
	var returns MonitoringFacade

	_jsii_.Invoke(
		m,
		"monitorCustom",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitoringFacade) MonitorDocumentDbCluster(props *DocumentDbMonitoringProps) MonitoringFacade {
	if err := m.validateMonitorDocumentDbClusterParameters(props); err != nil {
		panic(err)
	}
	var returns MonitoringFacade

	_jsii_.Invoke(
		m,
		"monitorDocumentDbCluster",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitoringFacade) MonitorDynamoTable(props *DynamoTableMonitoringProps) MonitoringFacade {
	if err := m.validateMonitorDynamoTableParameters(props); err != nil {
		panic(err)
	}
	var returns MonitoringFacade

	_jsii_.Invoke(
		m,
		"monitorDynamoTable",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitoringFacade) MonitorDynamoTableGlobalSecondaryIndex(props *DynamoTableGlobalSecondaryIndexMonitoringProps) MonitoringFacade {
	if err := m.validateMonitorDynamoTableGlobalSecondaryIndexParameters(props); err != nil {
		panic(err)
	}
	var returns MonitoringFacade

	_jsii_.Invoke(
		m,
		"monitorDynamoTableGlobalSecondaryIndex",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitoringFacade) MonitorEc2ApplicationLoadBalancer(props *Ec2ApplicationLoadBalancerMonitoringProps) MonitoringFacade {
	if err := m.validateMonitorEc2ApplicationLoadBalancerParameters(props); err != nil {
		panic(err)
	}
	var returns MonitoringFacade

	_jsii_.Invoke(
		m,
		"monitorEc2ApplicationLoadBalancer",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitoringFacade) MonitorEC2Instances(props *EC2MonitoringProps) MonitoringFacade {
	if err := m.validateMonitorEC2InstancesParameters(props); err != nil {
		panic(err)
	}
	var returns MonitoringFacade

	_jsii_.Invoke(
		m,
		"monitorEC2Instances",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitoringFacade) MonitorEc2NetworkLoadBalancer(props *Ec2NetworkLoadBalancerMonitoringProps) MonitoringFacade {
	if err := m.validateMonitorEc2NetworkLoadBalancerParameters(props); err != nil {
		panic(err)
	}
	var returns MonitoringFacade

	_jsii_.Invoke(
		m,
		"monitorEc2NetworkLoadBalancer",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitoringFacade) MonitorEc2Service(props *Ec2ServiceMonitoringProps) MonitoringFacade {
	if err := m.validateMonitorEc2ServiceParameters(props); err != nil {
		panic(err)
	}
	var returns MonitoringFacade

	_jsii_.Invoke(
		m,
		"monitorEc2Service",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitoringFacade) MonitorElastiCacheCluster(props *ElastiCacheClusterMonitoringProps) MonitoringFacade {
	if err := m.validateMonitorElastiCacheClusterParameters(props); err != nil {
		panic(err)
	}
	var returns MonitoringFacade

	_jsii_.Invoke(
		m,
		"monitorElastiCacheCluster",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitoringFacade) MonitorElasticsearchCluster(props *OpenSearchClusterMonitoringProps) MonitoringFacade {
	if err := m.validateMonitorElasticsearchClusterParameters(props); err != nil {
		panic(err)
	}
	var returns MonitoringFacade

	_jsii_.Invoke(
		m,
		"monitorElasticsearchCluster",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitoringFacade) MonitorFargateApplicationLoadBalancer(props *FargateApplicationLoadBalancerMonitoringProps) MonitoringFacade {
	if err := m.validateMonitorFargateApplicationLoadBalancerParameters(props); err != nil {
		panic(err)
	}
	var returns MonitoringFacade

	_jsii_.Invoke(
		m,
		"monitorFargateApplicationLoadBalancer",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitoringFacade) MonitorFargateNetworkLoadBalancer(props *FargateNetworkLoadBalancerMonitoringProps) MonitoringFacade {
	if err := m.validateMonitorFargateNetworkLoadBalancerParameters(props); err != nil {
		panic(err)
	}
	var returns MonitoringFacade

	_jsii_.Invoke(
		m,
		"monitorFargateNetworkLoadBalancer",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitoringFacade) MonitorFargateService(props *FargateServiceMonitoringProps) MonitoringFacade {
	if err := m.validateMonitorFargateServiceParameters(props); err != nil {
		panic(err)
	}
	var returns MonitoringFacade

	_jsii_.Invoke(
		m,
		"monitorFargateService",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitoringFacade) MonitorGlueJob(props *GlueJobMonitoringProps) MonitoringFacade {
	if err := m.validateMonitorGlueJobParameters(props); err != nil {
		panic(err)
	}
	var returns MonitoringFacade

	_jsii_.Invoke(
		m,
		"monitorGlueJob",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitoringFacade) MonitorKinesisDataAnalytics(props *KinesisDataAnalyticsMonitoringProps) MonitoringFacade {
	if err := m.validateMonitorKinesisDataAnalyticsParameters(props); err != nil {
		panic(err)
	}
	var returns MonitoringFacade

	_jsii_.Invoke(
		m,
		"monitorKinesisDataAnalytics",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitoringFacade) MonitorKinesisDataStream(props *KinesisDataStreamMonitoringProps) MonitoringFacade {
	if err := m.validateMonitorKinesisDataStreamParameters(props); err != nil {
		panic(err)
	}
	var returns MonitoringFacade

	_jsii_.Invoke(
		m,
		"monitorKinesisDataStream",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitoringFacade) MonitorKinesisFirehose(props *KinesisFirehoseMonitoringProps) MonitoringFacade {
	if err := m.validateMonitorKinesisFirehoseParameters(props); err != nil {
		panic(err)
	}
	var returns MonitoringFacade

	_jsii_.Invoke(
		m,
		"monitorKinesisFirehose",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitoringFacade) MonitorLambdaFunction(props *LambdaFunctionMonitoringProps) MonitoringFacade {
	if err := m.validateMonitorLambdaFunctionParameters(props); err != nil {
		panic(err)
	}
	var returns MonitoringFacade

	_jsii_.Invoke(
		m,
		"monitorLambdaFunction",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitoringFacade) MonitorLog(props *LogMonitoringProps) MonitoringFacade {
	if err := m.validateMonitorLogParameters(props); err != nil {
		panic(err)
	}
	var returns MonitoringFacade

	_jsii_.Invoke(
		m,
		"monitorLog",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitoringFacade) MonitorNetworkLoadBalancer(props *NetworkLoadBalancerMonitoringProps) MonitoringFacade {
	if err := m.validateMonitorNetworkLoadBalancerParameters(props); err != nil {
		panic(err)
	}
	var returns MonitoringFacade

	_jsii_.Invoke(
		m,
		"monitorNetworkLoadBalancer",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitoringFacade) MonitorOpenSearchCluster(props *OpenSearchClusterMonitoringProps) MonitoringFacade {
	if err := m.validateMonitorOpenSearchClusterParameters(props); err != nil {
		panic(err)
	}
	var returns MonitoringFacade

	_jsii_.Invoke(
		m,
		"monitorOpenSearchCluster",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitoringFacade) MonitorQueueProcessingEc2Service(props *QueueProcessingEc2ServiceMonitoringProps) MonitoringFacade {
	if err := m.validateMonitorQueueProcessingEc2ServiceParameters(props); err != nil {
		panic(err)
	}
	var returns MonitoringFacade

	_jsii_.Invoke(
		m,
		"monitorQueueProcessingEc2Service",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitoringFacade) MonitorQueueProcessingFargateService(props *QueueProcessingFargateServiceMonitoringProps) MonitoringFacade {
	if err := m.validateMonitorQueueProcessingFargateServiceParameters(props); err != nil {
		panic(err)
	}
	var returns MonitoringFacade

	_jsii_.Invoke(
		m,
		"monitorQueueProcessingFargateService",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitoringFacade) MonitorRdsCluster(props *RdsClusterMonitoringProps) MonitoringFacade {
	if err := m.validateMonitorRdsClusterParameters(props); err != nil {
		panic(err)
	}
	var returns MonitoringFacade

	_jsii_.Invoke(
		m,
		"monitorRdsCluster",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitoringFacade) MonitorRedshiftCluster(props *RedshiftClusterMonitoringProps) MonitoringFacade {
	if err := m.validateMonitorRedshiftClusterParameters(props); err != nil {
		panic(err)
	}
	var returns MonitoringFacade

	_jsii_.Invoke(
		m,
		"monitorRedshiftCluster",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitoringFacade) MonitorS3Bucket(props *S3BucketMonitoringProps) MonitoringFacade {
	if err := m.validateMonitorS3BucketParameters(props); err != nil {
		panic(err)
	}
	var returns MonitoringFacade

	_jsii_.Invoke(
		m,
		"monitorS3Bucket",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitoringFacade) MonitorScope(scope constructs.Construct, aspectProps *MonitoringAspectProps) {
	if err := m.validateMonitorScopeParameters(scope, aspectProps); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"monitorScope",
		[]interface{}{scope, aspectProps},
	)
}

func (m *jsiiProxy_MonitoringFacade) MonitorSecretsManager(props *SecretsManagerMonitoringProps) MonitoringFacade {
	if err := m.validateMonitorSecretsManagerParameters(props); err != nil {
		panic(err)
	}
	var returns MonitoringFacade

	_jsii_.Invoke(
		m,
		"monitorSecretsManager",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitoringFacade) MonitorSecretsManagerSecret(props *SecretsManagerSecretMonitoringProps) MonitoringFacade {
	if err := m.validateMonitorSecretsManagerSecretParameters(props); err != nil {
		panic(err)
	}
	var returns MonitoringFacade

	_jsii_.Invoke(
		m,
		"monitorSecretsManagerSecret",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitoringFacade) MonitorSimpleEc2Service(props *SimpleEc2ServiceMonitoringProps) MonitoringFacade {
	if err := m.validateMonitorSimpleEc2ServiceParameters(props); err != nil {
		panic(err)
	}
	var returns MonitoringFacade

	_jsii_.Invoke(
		m,
		"monitorSimpleEc2Service",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitoringFacade) MonitorSimpleFargateService(props *SimpleFargateServiceMonitoringProps) MonitoringFacade {
	if err := m.validateMonitorSimpleFargateServiceParameters(props); err != nil {
		panic(err)
	}
	var returns MonitoringFacade

	_jsii_.Invoke(
		m,
		"monitorSimpleFargateService",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitoringFacade) MonitorSnsTopic(props *SnsTopicMonitoringProps) MonitoringFacade {
	if err := m.validateMonitorSnsTopicParameters(props); err != nil {
		panic(err)
	}
	var returns MonitoringFacade

	_jsii_.Invoke(
		m,
		"monitorSnsTopic",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitoringFacade) MonitorSqsQueue(props *SqsQueueMonitoringProps) MonitoringFacade {
	if err := m.validateMonitorSqsQueueParameters(props); err != nil {
		panic(err)
	}
	var returns MonitoringFacade

	_jsii_.Invoke(
		m,
		"monitorSqsQueue",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitoringFacade) MonitorSqsQueueWithDlq(props *SqsQueueMonitoringWithDlqProps) MonitoringFacade {
	if err := m.validateMonitorSqsQueueWithDlqParameters(props); err != nil {
		panic(err)
	}
	var returns MonitoringFacade

	_jsii_.Invoke(
		m,
		"monitorSqsQueueWithDlq",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitoringFacade) MonitorStepFunction(props *StepFunctionMonitoringProps) MonitoringFacade {
	if err := m.validateMonitorStepFunctionParameters(props); err != nil {
		panic(err)
	}
	var returns MonitoringFacade

	_jsii_.Invoke(
		m,
		"monitorStepFunction",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitoringFacade) MonitorStepFunctionActivity(props *StepFunctionActivityMonitoringProps) MonitoringFacade {
	if err := m.validateMonitorStepFunctionActivityParameters(props); err != nil {
		panic(err)
	}
	var returns MonitoringFacade

	_jsii_.Invoke(
		m,
		"monitorStepFunctionActivity",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitoringFacade) MonitorStepFunctionLambdaIntegration(props *StepFunctionLambdaIntegrationMonitoringProps) MonitoringFacade {
	if err := m.validateMonitorStepFunctionLambdaIntegrationParameters(props); err != nil {
		panic(err)
	}
	var returns MonitoringFacade

	_jsii_.Invoke(
		m,
		"monitorStepFunctionLambdaIntegration",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitoringFacade) MonitorStepFunctionServiceIntegration(props *StepFunctionServiceIntegrationMonitoringProps) MonitoringFacade {
	if err := m.validateMonitorStepFunctionServiceIntegrationParameters(props); err != nil {
		panic(err)
	}
	var returns MonitoringFacade

	_jsii_.Invoke(
		m,
		"monitorStepFunctionServiceIntegration",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitoringFacade) MonitorSyntheticsCanary(props *SyntheticsCanaryMonitoringProps) MonitoringFacade {
	if err := m.validateMonitorSyntheticsCanaryParameters(props); err != nil {
		panic(err)
	}
	var returns MonitoringFacade

	_jsii_.Invoke(
		m,
		"monitorSyntheticsCanary",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitoringFacade) MonitorWebApplicationFirewallAclV2(props *WafV2MonitoringProps) MonitoringFacade {
	if err := m.validateMonitorWebApplicationFirewallAclV2Parameters(props); err != nil {
		panic(err)
	}
	var returns MonitoringFacade

	_jsii_.Invoke(
		m,
		"monitorWebApplicationFirewallAclV2",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitoringFacade) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}


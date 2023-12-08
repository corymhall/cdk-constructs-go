package cdkmonitoringconstructs

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/corymhall/cdk-constructs-go/cdkmonitoringconstructs/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudwatch"
)

// Experimental.
type CloudFrontDistributionMonitoring interface {
	Monitoring
	// Experimental.
	AdditionalMetricsEnabled() *bool
	// Experimental.
	AlarmFactory() AlarmFactory
	// Experimental.
	Alarms() *[]*AlarmWithAnnotation
	// Experimental.
	CacheHitRate() interface{}
	// Experimental.
	DistributionUrl() *string
	// Experimental.
	DownloadedBytesMetric() interface{}
	// Experimental.
	Error4xxRate() interface{}
	// Experimental.
	Error5xxRate() interface{}
	// Experimental.
	ErrorAlarmFactory() ErrorAlarmFactory
	// Experimental.
	ErrorRateAnnotations() *[]*awscloudwatch.HorizontalAnnotation
	// Experimental.
	LocalAlarmNamePrefixOverride() *string
	// Experimental.
	NamingStrategy() MonitoringNamingStrategy
	// Experimental.
	Scope() MonitoringScope
	// Experimental.
	Title() *string
	// Experimental.
	TpsAlarmFactory() TpsAlarmFactory
	// Experimental.
	TpsAnnotations() *[]*awscloudwatch.HorizontalAnnotation
	// Experimental.
	TpsMetric() interface{}
	// Experimental.
	UploadedBytesMetric() interface{}
	// Adds an alarm.
	// Experimental.
	AddAlarm(alarm *AlarmWithAnnotation)
	// Returns widgets for all alarms.
	//
	// These can go to runbook or to service dashboard.
	// Experimental.
	AlarmWidgets() *[]awscloudwatch.IWidget
	// Creates a new alarm factory.
	//
	// Alarms created will be named with the given prefix, unless a local name override is present.
	// Experimental.
	CreateAlarmFactory(alarmNamePrefix *string) AlarmFactory
	// Experimental.
	CreateCacheWidget(width *float64, height *float64) awscloudwatch.GraphWidget
	// Returns all the alarms created.
	// Experimental.
	CreatedAlarms() *[]*AlarmWithAnnotation
	// Experimental.
	CreateErrorRateWidget(width *float64, height *float64) awscloudwatch.GraphWidget
	// Creates a new metric factory.
	// Experimental.
	CreateMetricFactory() MetricFactory
	// Experimental.
	CreateTitleWidget() MonitoringHeaderWidget
	// Experimental.
	CreateTpsWidget(width *float64, height *float64) awscloudwatch.GraphWidget
	// Experimental.
	CreateTrafficWidget(width *float64, height *float64) awscloudwatch.GraphWidget
	// Creates a new widget factory.
	// Experimental.
	CreateWidgetFactory() IWidgetFactory
	// Returns widgets to be placed on the summary dashboard.
	// Experimental.
	SummaryWidgets() *[]awscloudwatch.IWidget
	// Returns widgets to be placed on the main dashboard.
	// Experimental.
	Widgets() *[]awscloudwatch.IWidget
	// Returns widgets for the requested dashboard type.
	// Experimental.
	WidgetsForDashboard(name *string) *[]awscloudwatch.IWidget
}

// The jsii proxy struct for CloudFrontDistributionMonitoring
type jsiiProxy_CloudFrontDistributionMonitoring struct {
	jsiiProxy_Monitoring
}

func (j *jsiiProxy_CloudFrontDistributionMonitoring) AdditionalMetricsEnabled() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"additionalMetricsEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudFrontDistributionMonitoring) AlarmFactory() AlarmFactory {
	var returns AlarmFactory
	_jsii_.Get(
		j,
		"alarmFactory",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudFrontDistributionMonitoring) Alarms() *[]*AlarmWithAnnotation {
	var returns *[]*AlarmWithAnnotation
	_jsii_.Get(
		j,
		"alarms",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudFrontDistributionMonitoring) CacheHitRate() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"cacheHitRate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudFrontDistributionMonitoring) DistributionUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"distributionUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudFrontDistributionMonitoring) DownloadedBytesMetric() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"downloadedBytesMetric",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudFrontDistributionMonitoring) Error4xxRate() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"error4xxRate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudFrontDistributionMonitoring) Error5xxRate() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"error5xxRate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudFrontDistributionMonitoring) ErrorAlarmFactory() ErrorAlarmFactory {
	var returns ErrorAlarmFactory
	_jsii_.Get(
		j,
		"errorAlarmFactory",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudFrontDistributionMonitoring) ErrorRateAnnotations() *[]*awscloudwatch.HorizontalAnnotation {
	var returns *[]*awscloudwatch.HorizontalAnnotation
	_jsii_.Get(
		j,
		"errorRateAnnotations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudFrontDistributionMonitoring) LocalAlarmNamePrefixOverride() *string {
	var returns *string
	_jsii_.Get(
		j,
		"localAlarmNamePrefixOverride",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudFrontDistributionMonitoring) NamingStrategy() MonitoringNamingStrategy {
	var returns MonitoringNamingStrategy
	_jsii_.Get(
		j,
		"namingStrategy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudFrontDistributionMonitoring) Scope() MonitoringScope {
	var returns MonitoringScope
	_jsii_.Get(
		j,
		"scope",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudFrontDistributionMonitoring) Title() *string {
	var returns *string
	_jsii_.Get(
		j,
		"title",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudFrontDistributionMonitoring) TpsAlarmFactory() TpsAlarmFactory {
	var returns TpsAlarmFactory
	_jsii_.Get(
		j,
		"tpsAlarmFactory",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudFrontDistributionMonitoring) TpsAnnotations() *[]*awscloudwatch.HorizontalAnnotation {
	var returns *[]*awscloudwatch.HorizontalAnnotation
	_jsii_.Get(
		j,
		"tpsAnnotations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudFrontDistributionMonitoring) TpsMetric() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tpsMetric",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudFrontDistributionMonitoring) UploadedBytesMetric() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"uploadedBytesMetric",
		&returns,
	)
	return returns
}


// Experimental.
func NewCloudFrontDistributionMonitoring(scope MonitoringScope, props *CloudFrontDistributionMonitoringProps) CloudFrontDistributionMonitoring {
	_init_.Initialize()

	if err := validateNewCloudFrontDistributionMonitoringParameters(scope, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_CloudFrontDistributionMonitoring{}

	_jsii_.Create(
		"cdk-monitoring-constructs.CloudFrontDistributionMonitoring",
		[]interface{}{scope, props},
		&j,
	)

	return &j
}

// Experimental.
func NewCloudFrontDistributionMonitoring_Override(c CloudFrontDistributionMonitoring, scope MonitoringScope, props *CloudFrontDistributionMonitoringProps) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk-monitoring-constructs.CloudFrontDistributionMonitoring",
		[]interface{}{scope, props},
		c,
	)
}

func (c *jsiiProxy_CloudFrontDistributionMonitoring) AddAlarm(alarm *AlarmWithAnnotation) {
	if err := c.validateAddAlarmParameters(alarm); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"addAlarm",
		[]interface{}{alarm},
	)
}

func (c *jsiiProxy_CloudFrontDistributionMonitoring) AlarmWidgets() *[]awscloudwatch.IWidget {
	var returns *[]awscloudwatch.IWidget

	_jsii_.Invoke(
		c,
		"alarmWidgets",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudFrontDistributionMonitoring) CreateAlarmFactory(alarmNamePrefix *string) AlarmFactory {
	if err := c.validateCreateAlarmFactoryParameters(alarmNamePrefix); err != nil {
		panic(err)
	}
	var returns AlarmFactory

	_jsii_.Invoke(
		c,
		"createAlarmFactory",
		[]interface{}{alarmNamePrefix},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudFrontDistributionMonitoring) CreateCacheWidget(width *float64, height *float64) awscloudwatch.GraphWidget {
	if err := c.validateCreateCacheWidgetParameters(width, height); err != nil {
		panic(err)
	}
	var returns awscloudwatch.GraphWidget

	_jsii_.Invoke(
		c,
		"createCacheWidget",
		[]interface{}{width, height},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudFrontDistributionMonitoring) CreatedAlarms() *[]*AlarmWithAnnotation {
	var returns *[]*AlarmWithAnnotation

	_jsii_.Invoke(
		c,
		"createdAlarms",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudFrontDistributionMonitoring) CreateErrorRateWidget(width *float64, height *float64) awscloudwatch.GraphWidget {
	if err := c.validateCreateErrorRateWidgetParameters(width, height); err != nil {
		panic(err)
	}
	var returns awscloudwatch.GraphWidget

	_jsii_.Invoke(
		c,
		"createErrorRateWidget",
		[]interface{}{width, height},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudFrontDistributionMonitoring) CreateMetricFactory() MetricFactory {
	var returns MetricFactory

	_jsii_.Invoke(
		c,
		"createMetricFactory",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudFrontDistributionMonitoring) CreateTitleWidget() MonitoringHeaderWidget {
	var returns MonitoringHeaderWidget

	_jsii_.Invoke(
		c,
		"createTitleWidget",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudFrontDistributionMonitoring) CreateTpsWidget(width *float64, height *float64) awscloudwatch.GraphWidget {
	if err := c.validateCreateTpsWidgetParameters(width, height); err != nil {
		panic(err)
	}
	var returns awscloudwatch.GraphWidget

	_jsii_.Invoke(
		c,
		"createTpsWidget",
		[]interface{}{width, height},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudFrontDistributionMonitoring) CreateTrafficWidget(width *float64, height *float64) awscloudwatch.GraphWidget {
	if err := c.validateCreateTrafficWidgetParameters(width, height); err != nil {
		panic(err)
	}
	var returns awscloudwatch.GraphWidget

	_jsii_.Invoke(
		c,
		"createTrafficWidget",
		[]interface{}{width, height},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudFrontDistributionMonitoring) CreateWidgetFactory() IWidgetFactory {
	var returns IWidgetFactory

	_jsii_.Invoke(
		c,
		"createWidgetFactory",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudFrontDistributionMonitoring) SummaryWidgets() *[]awscloudwatch.IWidget {
	var returns *[]awscloudwatch.IWidget

	_jsii_.Invoke(
		c,
		"summaryWidgets",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudFrontDistributionMonitoring) Widgets() *[]awscloudwatch.IWidget {
	var returns *[]awscloudwatch.IWidget

	_jsii_.Invoke(
		c,
		"widgets",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudFrontDistributionMonitoring) WidgetsForDashboard(name *string) *[]awscloudwatch.IWidget {
	if err := c.validateWidgetsForDashboardParameters(name); err != nil {
		panic(err)
	}
	var returns *[]awscloudwatch.IWidget

	_jsii_.Invoke(
		c,
		"widgetsForDashboard",
		[]interface{}{name},
		&returns,
	)

	return returns
}


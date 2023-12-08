package cdkmonitoringconstructs

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/corymhall/cdk-constructs-go/cdkmonitoringconstructs/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudwatch"
	"github.com/aws/constructs-go/constructs/v10"
)

// Experimental.
type AlarmFactory interface {
	// Experimental.
	AlarmNamingStrategy() IAlarmNamingStrategy
	// Experimental.
	AlarmScope() constructs.Construct
	// Experimental.
	GlobalAlarmDefaults() *AlarmFactoryDefaults
	// Experimental.
	GlobalMetricDefaults() *MetricFactoryDefaults
	// Experimental.
	ShouldUseDefaultDedupeForError() *bool
	// Experimental.
	ShouldUseDefaultDedupeForLatency() *bool
	// Experimental.
	AddAlarm(metric interface{}, props *AddAlarmProps) *AlarmWithAnnotation
	// Experimental.
	AddCompositeAlarm(alarms *[]*AlarmWithAnnotation, props *AddCompositeAlarmProps) awscloudwatch.CompositeAlarm
	// Experimental.
	CreateAnnotation(props *AlarmAnnotationStrategyProps) *awscloudwatch.HorizontalAnnotation
	// Experimental.
	DetermineAction(disambiguator *string, actionOverride IAlarmActionStrategy) IAlarmActionStrategy
	// Experimental.
	DetermineActionsEnabled(actionsEnabled *bool, disambiguator *string) *bool
	// Experimental.
	DetermineCompositeAlarmRule(alarms *[]*AlarmWithAnnotation, props *AddCompositeAlarmProps) awscloudwatch.IAlarmRule
	// Experimental.
	GenerateDescription(alarmDescription *string, alarmDescriptionOverride *string, runbookLinkOverride *string, documentationLinkOverride *string) *string
	// Experimental.
	JoinDescriptionParts(parts ...*string) *string
}

// The jsii proxy struct for AlarmFactory
type jsiiProxy_AlarmFactory struct {
	_ byte // padding
}

func (j *jsiiProxy_AlarmFactory) AlarmNamingStrategy() IAlarmNamingStrategy {
	var returns IAlarmNamingStrategy
	_jsii_.Get(
		j,
		"alarmNamingStrategy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlarmFactory) AlarmScope() constructs.Construct {
	var returns constructs.Construct
	_jsii_.Get(
		j,
		"alarmScope",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlarmFactory) GlobalAlarmDefaults() *AlarmFactoryDefaults {
	var returns *AlarmFactoryDefaults
	_jsii_.Get(
		j,
		"globalAlarmDefaults",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlarmFactory) GlobalMetricDefaults() *MetricFactoryDefaults {
	var returns *MetricFactoryDefaults
	_jsii_.Get(
		j,
		"globalMetricDefaults",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlarmFactory) ShouldUseDefaultDedupeForError() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"shouldUseDefaultDedupeForError",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlarmFactory) ShouldUseDefaultDedupeForLatency() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"shouldUseDefaultDedupeForLatency",
		&returns,
	)
	return returns
}


// Experimental.
func NewAlarmFactory(alarmScope constructs.Construct, props *AlarmFactoryProps) AlarmFactory {
	_init_.Initialize()

	if err := validateNewAlarmFactoryParameters(alarmScope, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_AlarmFactory{}

	_jsii_.Create(
		"cdk-monitoring-constructs.AlarmFactory",
		[]interface{}{alarmScope, props},
		&j,
	)

	return &j
}

// Experimental.
func NewAlarmFactory_Override(a AlarmFactory, alarmScope constructs.Construct, props *AlarmFactoryProps) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk-monitoring-constructs.AlarmFactory",
		[]interface{}{alarmScope, props},
		a,
	)
}

func (a *jsiiProxy_AlarmFactory) AddAlarm(metric interface{}, props *AddAlarmProps) *AlarmWithAnnotation {
	if err := a.validateAddAlarmParameters(metric, props); err != nil {
		panic(err)
	}
	var returns *AlarmWithAnnotation

	_jsii_.Invoke(
		a,
		"addAlarm",
		[]interface{}{metric, props},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlarmFactory) AddCompositeAlarm(alarms *[]*AlarmWithAnnotation, props *AddCompositeAlarmProps) awscloudwatch.CompositeAlarm {
	if err := a.validateAddCompositeAlarmParameters(alarms, props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.CompositeAlarm

	_jsii_.Invoke(
		a,
		"addCompositeAlarm",
		[]interface{}{alarms, props},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlarmFactory) CreateAnnotation(props *AlarmAnnotationStrategyProps) *awscloudwatch.HorizontalAnnotation {
	if err := a.validateCreateAnnotationParameters(props); err != nil {
		panic(err)
	}
	var returns *awscloudwatch.HorizontalAnnotation

	_jsii_.Invoke(
		a,
		"createAnnotation",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlarmFactory) DetermineAction(disambiguator *string, actionOverride IAlarmActionStrategy) IAlarmActionStrategy {
	var returns IAlarmActionStrategy

	_jsii_.Invoke(
		a,
		"determineAction",
		[]interface{}{disambiguator, actionOverride},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlarmFactory) DetermineActionsEnabled(actionsEnabled *bool, disambiguator *string) *bool {
	var returns *bool

	_jsii_.Invoke(
		a,
		"determineActionsEnabled",
		[]interface{}{actionsEnabled, disambiguator},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlarmFactory) DetermineCompositeAlarmRule(alarms *[]*AlarmWithAnnotation, props *AddCompositeAlarmProps) awscloudwatch.IAlarmRule {
	if err := a.validateDetermineCompositeAlarmRuleParameters(alarms, props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.IAlarmRule

	_jsii_.Invoke(
		a,
		"determineCompositeAlarmRule",
		[]interface{}{alarms, props},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlarmFactory) GenerateDescription(alarmDescription *string, alarmDescriptionOverride *string, runbookLinkOverride *string, documentationLinkOverride *string) *string {
	if err := a.validateGenerateDescriptionParameters(alarmDescription); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		a,
		"generateDescription",
		[]interface{}{alarmDescription, alarmDescriptionOverride, runbookLinkOverride, documentationLinkOverride},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AlarmFactory) JoinDescriptionParts(parts ...*string) *string {
	args := []interface{}{}
	for _, a := range parts {
		args = append(args, a)
	}

	var returns *string

	_jsii_.Invoke(
		a,
		"joinDescriptionParts",
		args,
		&returns,
	)

	return returns
}


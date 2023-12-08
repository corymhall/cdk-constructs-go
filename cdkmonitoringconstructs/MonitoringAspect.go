package cdkmonitoringconstructs

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/corymhall/cdk-constructs-go/cdkmonitoringconstructs/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/corymhall/cdk-constructs-go/cdkmonitoringconstructs/internal"
)

// A CDK aspect that adds support for monitoring all resources within scope.
// Experimental.
type MonitoringAspect interface {
	awscdk.IAspect
	// All aspects can visit an IConstruct.
	// Experimental.
	Visit(node constructs.IConstruct)
}

// The jsii proxy struct for MonitoringAspect
type jsiiProxy_MonitoringAspect struct {
	internal.Type__awscdkIAspect
}

// Experimental.
func NewMonitoringAspect(monitoringFacade MonitoringFacade, props *MonitoringAspectProps) MonitoringAspect {
	_init_.Initialize()

	if err := validateNewMonitoringAspectParameters(monitoringFacade, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_MonitoringAspect{}

	_jsii_.Create(
		"cdk-monitoring-constructs.MonitoringAspect",
		[]interface{}{monitoringFacade, props},
		&j,
	)

	return &j
}

// Experimental.
func NewMonitoringAspect_Override(m MonitoringAspect, monitoringFacade MonitoringFacade, props *MonitoringAspectProps) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk-monitoring-constructs.MonitoringAspect",
		[]interface{}{monitoringFacade, props},
		m,
	)
}

func (m *jsiiProxy_MonitoringAspect) Visit(node constructs.IConstruct) {
	if err := m.validateVisitParameters(node); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"visit",
		[]interface{}{node},
	)
}


package cdkmonitoringconstructs

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/corymhall/cdk-constructs-go/cdkmonitoringconstructs/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudwatch"
)

// Experimental.
type StaticSegmentDynamicAdapter interface {
	IDynamicDashboardSegment
	// Experimental.
	Props() IDashboardFactoryProps
	// Adapts an IDashboardSegment to the IDynamicDashboardSegment interface by using overrideProps to determine if a segment should be shown on a specific dashboard.
	//
	// The default values are true, so consumers must set these to false if they would
	// like to hide these items from dashboards.
	// Experimental.
	WidgetsForDashboard(name *string) *[]awscloudwatch.IWidget
}

// The jsii proxy struct for StaticSegmentDynamicAdapter
type jsiiProxy_StaticSegmentDynamicAdapter struct {
	jsiiProxy_IDynamicDashboardSegment
}

func (j *jsiiProxy_StaticSegmentDynamicAdapter) Props() IDashboardFactoryProps {
	var returns IDashboardFactoryProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}


// Experimental.
func NewStaticSegmentDynamicAdapter(props IDashboardFactoryProps) StaticSegmentDynamicAdapter {
	_init_.Initialize()

	if err := validateNewStaticSegmentDynamicAdapterParameters(props); err != nil {
		panic(err)
	}
	j := jsiiProxy_StaticSegmentDynamicAdapter{}

	_jsii_.Create(
		"cdk-monitoring-constructs.StaticSegmentDynamicAdapter",
		[]interface{}{props},
		&j,
	)

	return &j
}

// Experimental.
func NewStaticSegmentDynamicAdapter_Override(s StaticSegmentDynamicAdapter, props IDashboardFactoryProps) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk-monitoring-constructs.StaticSegmentDynamicAdapter",
		[]interface{}{props},
		s,
	)
}

func (s *jsiiProxy_StaticSegmentDynamicAdapter) WidgetsForDashboard(name *string) *[]awscloudwatch.IWidget {
	if err := s.validateWidgetsForDashboardParameters(name); err != nil {
		panic(err)
	}
	var returns *[]awscloudwatch.IWidget

	_jsii_.Invoke(
		s,
		"widgetsForDashboard",
		[]interface{}{name},
		&returns,
	)

	return returns
}


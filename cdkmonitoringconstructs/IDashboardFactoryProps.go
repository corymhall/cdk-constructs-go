package cdkmonitoringconstructs

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Experimental.
type IDashboardFactoryProps interface {
	// Dashboard placement override props.
	// Default: - all default.
	//
	// Experimental.
	OverrideProps() *MonitoringDashboardsOverrideProps
	// Experimental.
	SetOverrideProps(o *MonitoringDashboardsOverrideProps)
	// Segment to be placed on the dashboard.
	// Experimental.
	Segment() IDashboardSegment
	// Experimental.
	SetSegment(s IDashboardSegment)
}

// The jsii proxy for IDashboardFactoryProps
type jsiiProxy_IDashboardFactoryProps struct {
	_ byte // padding
}

func (j *jsiiProxy_IDashboardFactoryProps) OverrideProps() *MonitoringDashboardsOverrideProps {
	var returns *MonitoringDashboardsOverrideProps
	_jsii_.Get(
		j,
		"overrideProps",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IDashboardFactoryProps)SetOverrideProps(val *MonitoringDashboardsOverrideProps) {
	if err := j.validateSetOverridePropsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"overrideProps",
		val,
	)
}

func (j *jsiiProxy_IDashboardFactoryProps) Segment() IDashboardSegment {
	var returns IDashboardSegment
	_jsii_.Get(
		j,
		"segment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IDashboardFactoryProps)SetSegment(val IDashboardSegment) {
	if err := j.validateSetSegmentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"segment",
		val,
	)
}


package cdkmonitoringconstructs

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudwatch"
)

// Experimental.
type DynamicDashboardConfiguration struct {
	// Name of the dashboard. Full dashboard name will take the form of: `{@link MonitoringDynamicDashboardsProps.dashboardNamePrefix}-{@link name}`.
	//
	// NOTE: The dashboard names in {@link DefaultDashboardFactory.DefaultDashboards}
	// are reserved and cannot be used as dashboard names.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Period override for the dashboard.
	// Default: - respect individual graphs (PeriodOverride.INHERIT)
	//
	// Experimental.
	PeriodOverride awscloudwatch.PeriodOverride `field:"optional" json:"periodOverride" yaml:"periodOverride"`
	// Range of the dashboard.
	// Default: - 8 hours.
	//
	// Experimental.
	Range awscdk.Duration `field:"optional" json:"range" yaml:"range"`
	// Dashboard rendering preference.
	// Default: - DashboardRenderingPreference.INTERACTIVE_ONLY
	//
	// Experimental.
	RenderingPreference DashboardRenderingPreference `field:"optional" json:"renderingPreference" yaml:"renderingPreference"`
}


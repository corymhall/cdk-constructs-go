package cdkmonitoringconstructs

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudwatch"
)

// Experimental.
type MonitoringDashboardsProps struct {
	// Prefix added to each dashboard name.
	//
	// This allows to have all dashboards sorted close to each other and also separate multiple monitoring facades.
	// Experimental.
	DashboardNamePrefix *string `field:"required" json:"dashboardNamePrefix" yaml:"dashboardNamePrefix"`
	// Flag indicating whether the alarm dashboard should be created.
	//
	// This is independent on other create dashboard flags.
	// Default: - false.
	//
	// Experimental.
	CreateAlarmDashboard *bool `field:"optional" json:"createAlarmDashboard" yaml:"createAlarmDashboard"`
	// Flag indicating whether the default dashboard should be created.
	//
	// This is independent on other create dashboard flags.
	// Default: - true.
	//
	// Experimental.
	CreateDashboard *bool `field:"optional" json:"createDashboard" yaml:"createDashboard"`
	// Flag indicating whether the summary dashboard should be created.
	//
	// This is independent on other create dashboard flags.
	// Default: - false.
	//
	// Experimental.
	CreateSummaryDashboard *bool `field:"optional" json:"createSummaryDashboard" yaml:"createSummaryDashboard"`
	// Period override for the detail dashboard (and other auxiliary dashboards).
	// Default: - respect individual graphs (PeriodOverride.INHERIT)
	//
	// Experimental.
	DetailDashboardPeriodOverride awscloudwatch.PeriodOverride `field:"optional" json:"detailDashboardPeriodOverride" yaml:"detailDashboardPeriodOverride"`
	// Range of the detail dashboard (and other auxiliary dashboards).
	// See: DefaultDetailDashboardRange.
	//
	// Default: - 8 hours.
	//
	// Experimental.
	DetailDashboardRange awscdk.Duration `field:"optional" json:"detailDashboardRange" yaml:"detailDashboardRange"`
	// Dashboard rendering preference.
	// Default: - DashboardRenderingPreference.INTERACTIVE_ONLY
	//
	// Experimental.
	RenderingPreference DashboardRenderingPreference `field:"optional" json:"renderingPreference" yaml:"renderingPreference"`
	// Period override for the summary dashboard.
	// Default: - respect individual graphs (PeriodOverride.INHERIT)
	//
	// Experimental.
	SummaryDashboardPeriodOverride awscloudwatch.PeriodOverride `field:"optional" json:"summaryDashboardPeriodOverride" yaml:"summaryDashboardPeriodOverride"`
	// Range of the summary dashboard.
	// Default: - 14 days.
	//
	// Experimental.
	SummaryDashboardRange awscdk.Duration `field:"optional" json:"summaryDashboardRange" yaml:"summaryDashboardRange"`
}


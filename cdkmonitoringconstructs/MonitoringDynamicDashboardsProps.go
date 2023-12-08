package cdkmonitoringconstructs


// Experimental.
type MonitoringDynamicDashboardsProps struct {
	// List of dashboard types to generate.
	// Experimental.
	DashboardConfigs *[]*DynamicDashboardConfiguration `field:"required" json:"dashboardConfigs" yaml:"dashboardConfigs"`
	// Prefix added to each dashboard's name.
	//
	// This allows to have all dashboards sorted close to each other and also separate multiple monitoring facades.
	// Experimental.
	DashboardNamePrefix *string `field:"required" json:"dashboardNamePrefix" yaml:"dashboardNamePrefix"`
}


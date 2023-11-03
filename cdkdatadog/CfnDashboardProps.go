package cdkdatadog


// Datadog Dashboard 2.1.0.
type CfnDashboardProps struct {
	// JSON string of the dashboard definition.
	DashboardDefinition interface{} `field:"required" json:"dashboardDefinition" yaml:"dashboardDefinition"`
}


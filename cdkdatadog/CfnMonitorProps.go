package cdkdatadog


// Datadog Monitor 4.7.0.
type CfnMonitorProps struct {
	// The monitor query.
	Query *string `field:"required" json:"query" yaml:"query"`
	// The type of the monitor.
	Type CfnMonitorPropsType `field:"required" json:"type" yaml:"type"`
	Creator *Creator `field:"optional" json:"creator" yaml:"creator"`
	// A message to include with notifications for the monitor.
	Message *string `field:"optional" json:"message" yaml:"message"`
	// Whether or not the monitor is multi alert.
	Multi *bool `field:"optional" json:"multi" yaml:"multi"`
	// Name of the monitor.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The monitor options.
	Options *MonitorOptions `field:"optional" json:"options" yaml:"options"`
	// Integer from 1 (high) to 5 (low) indicating alert severity.
	Priority *float64 `field:"optional" json:"priority" yaml:"priority"`
	// A list of unique role identifiers to define which roles are allowed to edit the monitor.
	//
	// The unique identifiers for all roles can be pulled from the [Roles API](https://docs.datadoghq.com/api/latest/roles/#list-roles) and are located in the `data.id` field. Editing a monitor includes any updates to the monitor configuration, monitor deletion, and muting of the monitor for any amount of time. `restricted_roles` is the successor of `locked`. For more information about `locked` and `restricted_roles`, see the [monitor options docs](https://docs.datadoghq.com/monitors/guide/monitor_api_options/#permissions-options).
	RestrictedRoles *[]*string `field:"optional" json:"restrictedRoles" yaml:"restrictedRoles"`
	// Tags associated with the monitor.
	Tags *[]*string `field:"optional" json:"tags" yaml:"tags"`
}


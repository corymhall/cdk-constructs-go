package cdkmonitoringconstructs


// Preferred way of rendering dashboard widgets.
// Experimental.
type DashboardRenderingPreference string

const (
	// Create standard set of dashboards with interactive widgets only.
	// Experimental.
	DashboardRenderingPreference_INTERACTIVE_ONLY DashboardRenderingPreference = "INTERACTIVE_ONLY"
	// Create standard set of dashboards with bitmap widgets only.
	// Experimental.
	DashboardRenderingPreference_BITMAP_ONLY DashboardRenderingPreference = "BITMAP_ONLY"
	// Create a two sets of dashboards: standard set (interactive) and a copy (bitmap).
	// Experimental.
	DashboardRenderingPreference_INTERACTIVE_AND_BITMAP DashboardRenderingPreference = "INTERACTIVE_AND_BITMAP"
)


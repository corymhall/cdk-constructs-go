// cdk-constructs-datadog-go
package cdkdatadog

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"cdk-constructs-datadog-go.CfnDashboard",
		reflect.TypeOf((*CfnDashboard)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDeletionOverride", GoMethod: "AddDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addDependsOn", GoMethod: "AddDependsOn"},
			_jsii_.MemberMethod{JsiiMethod: "addMetadata", GoMethod: "AddMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyDeletionOverride", GoMethod: "AddPropertyDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyOverride", GoMethod: "AddPropertyOverride"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "attrId", GoGetter: "AttrId"},
			_jsii_.MemberProperty{JsiiProperty: "attrUrl", GoGetter: "AttrUrl"},
			_jsii_.MemberProperty{JsiiProperty: "cfnOptions", GoGetter: "CfnOptions"},
			_jsii_.MemberProperty{JsiiProperty: "cfnProperties", GoGetter: "CfnProperties"},
			_jsii_.MemberProperty{JsiiProperty: "cfnResourceType", GoGetter: "CfnResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberMethod{JsiiMethod: "getAtt", GoMethod: "GetAtt"},
			_jsii_.MemberMethod{JsiiMethod: "getMetadata", GoMethod: "GetMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "logicalId", GoGetter: "LogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "obtainDependencies", GoMethod: "ObtainDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "obtainResourceDependencies", GoMethod: "ObtainResourceDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "ref", GoGetter: "Ref"},
			_jsii_.MemberMethod{JsiiMethod: "removeDependency", GoMethod: "RemoveDependency"},
			_jsii_.MemberMethod{JsiiMethod: "renderProperties", GoMethod: "RenderProperties"},
			_jsii_.MemberMethod{JsiiMethod: "replaceDependency", GoMethod: "ReplaceDependency"},
			_jsii_.MemberMethod{JsiiMethod: "shouldSynthesize", GoMethod: "ShouldSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperites", GoGetter: "UpdatedProperites"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperties", GoGetter: "UpdatedProperties"},
			_jsii_.MemberMethod{JsiiMethod: "validateProperties", GoMethod: "ValidateProperties"},
		},
		func() interface{} {
			j := jsiiProxy_CfnDashboard{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkCfnResource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"cdk-constructs-datadog-go.CfnDashboardProps",
		reflect.TypeOf((*CfnDashboardProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"cdk-constructs-datadog-go.CfnMonitor",
		reflect.TypeOf((*CfnMonitor)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDeletionOverride", GoMethod: "AddDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addDependsOn", GoMethod: "AddDependsOn"},
			_jsii_.MemberMethod{JsiiMethod: "addMetadata", GoMethod: "AddMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyDeletionOverride", GoMethod: "AddPropertyDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyOverride", GoMethod: "AddPropertyOverride"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "attrCreated", GoGetter: "AttrCreated"},
			_jsii_.MemberProperty{JsiiProperty: "attrDeleted", GoGetter: "AttrDeleted"},
			_jsii_.MemberProperty{JsiiProperty: "attrId", GoGetter: "AttrId"},
			_jsii_.MemberProperty{JsiiProperty: "attrModified", GoGetter: "AttrModified"},
			_jsii_.MemberProperty{JsiiProperty: "cfnOptions", GoGetter: "CfnOptions"},
			_jsii_.MemberProperty{JsiiProperty: "cfnProperties", GoGetter: "CfnProperties"},
			_jsii_.MemberProperty{JsiiProperty: "cfnResourceType", GoGetter: "CfnResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberMethod{JsiiMethod: "getAtt", GoMethod: "GetAtt"},
			_jsii_.MemberMethod{JsiiMethod: "getMetadata", GoMethod: "GetMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "logicalId", GoGetter: "LogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "obtainDependencies", GoMethod: "ObtainDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "obtainResourceDependencies", GoMethod: "ObtainResourceDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "ref", GoGetter: "Ref"},
			_jsii_.MemberMethod{JsiiMethod: "removeDependency", GoMethod: "RemoveDependency"},
			_jsii_.MemberMethod{JsiiMethod: "renderProperties", GoMethod: "RenderProperties"},
			_jsii_.MemberMethod{JsiiMethod: "replaceDependency", GoMethod: "ReplaceDependency"},
			_jsii_.MemberMethod{JsiiMethod: "shouldSynthesize", GoMethod: "ShouldSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperites", GoGetter: "UpdatedProperites"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperties", GoGetter: "UpdatedProperties"},
			_jsii_.MemberMethod{JsiiMethod: "validateProperties", GoMethod: "ValidateProperties"},
		},
		func() interface{} {
			j := jsiiProxy_CfnMonitor{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkCfnResource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"cdk-constructs-datadog-go.CfnMonitorProps",
		reflect.TypeOf((*CfnMonitorProps)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"cdk-constructs-datadog-go.CfnMonitorPropsType",
		reflect.TypeOf((*CfnMonitorPropsType)(nil)).Elem(),
		map[string]interface{}{
			"AUDIT_ALERT": CfnMonitorPropsType_AUDIT_ALERT,
			"COMPOSITE": CfnMonitorPropsType_COMPOSITE,
			"EVENT_ALERT": CfnMonitorPropsType_EVENT_ALERT,
			"EVENT_V2_ALERT": CfnMonitorPropsType_EVENT_V2_ALERT,
			"LOG_ALERT": CfnMonitorPropsType_LOG_ALERT,
			"METRIC_ALERT": CfnMonitorPropsType_METRIC_ALERT,
			"PROCESS_ALERT": CfnMonitorPropsType_PROCESS_ALERT,
			"QUERY_ALERT": CfnMonitorPropsType_QUERY_ALERT,
			"SERVICE_CHECK": CfnMonitorPropsType_SERVICE_CHECK,
			"SYNTHETICS_ALERT": CfnMonitorPropsType_SYNTHETICS_ALERT,
			"TRACE_ANALYTICS_ALERT": CfnMonitorPropsType_TRACE_ANALYTICS_ALERT,
			"SLO_ALERT": CfnMonitorPropsType_SLO_ALERT,
			"RUM_ALERT": CfnMonitorPropsType_RUM_ALERT,
			"CI_PIPELINES_ALERT": CfnMonitorPropsType_CI_PIPELINES_ALERT,
			"ERROR_TRACKING_ALERT": CfnMonitorPropsType_ERROR_TRACKING_ALERT,
			"CI_TESTS_ALERT": CfnMonitorPropsType_CI_TESTS_ALERT,
		},
	)
	_jsii_.RegisterStruct(
		"cdk-constructs-datadog-go.Creator",
		reflect.TypeOf((*Creator)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"cdk-constructs-datadog-go.MonitorNotificationPresetName",
		reflect.TypeOf((*MonitorNotificationPresetName)(nil)).Elem(),
		map[string]interface{}{
			"SHOW_ALL": MonitorNotificationPresetName_SHOW_ALL,
			"HIDE_QUERY": MonitorNotificationPresetName_HIDE_QUERY,
			"HIDE_HANDLES": MonitorNotificationPresetName_HIDE_HANDLES,
			"HIDE_ALL": MonitorNotificationPresetName_HIDE_ALL,
		},
	)
	_jsii_.RegisterEnum(
		"cdk-constructs-datadog-go.MonitorOnMissingData",
		reflect.TypeOf((*MonitorOnMissingData)(nil)).Elem(),
		map[string]interface{}{
			"DEFAULT": MonitorOnMissingData_DEFAULT,
			"SHOW_NO_DATA": MonitorOnMissingData_SHOW_NO_DATA,
			"SHOW_AND_NOTIFY_NO_DATA": MonitorOnMissingData_SHOW_AND_NOTIFY_NO_DATA,
			"RESOLVE": MonitorOnMissingData_RESOLVE,
		},
	)
	_jsii_.RegisterStruct(
		"cdk-constructs-datadog-go.MonitorOptions",
		reflect.TypeOf((*MonitorOptions)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"cdk-constructs-datadog-go.MonitorOptionsRenotifyStatuses",
		reflect.TypeOf((*MonitorOptionsRenotifyStatuses)(nil)).Elem(),
		map[string]interface{}{
			"ALERT": MonitorOptionsRenotifyStatuses_ALERT,
			"NO_DATA": MonitorOptionsRenotifyStatuses_NO_DATA,
			"WARN": MonitorOptionsRenotifyStatuses_WARN,
		},
	)
	_jsii_.RegisterStruct(
		"cdk-constructs-datadog-go.MonitorSchedulingOptions",
		reflect.TypeOf((*MonitorSchedulingOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk-constructs-datadog-go.MonitorSchedulingOptionsEvaluationWindow",
		reflect.TypeOf((*MonitorSchedulingOptionsEvaluationWindow)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk-constructs-datadog-go.MonitorThresholdWindows",
		reflect.TypeOf((*MonitorThresholdWindows)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"cdk-constructs-datadog-go.MonitorThresholds",
		reflect.TypeOf((*MonitorThresholds)(nil)).Elem(),
	)
}

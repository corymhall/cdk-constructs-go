// Generated by cdk-import
import * as cdk from 'aws-cdk-lib';
import * as constructs from 'constructs';

/**
 * Datadog Monitor 4.7.0
 *
 * @schema CfnMonitorProps
 */
export interface CfnMonitorProps {
  /**
   * @schema CfnMonitorProps#Creator
   */
  readonly creator?: Creator;

  /**
   * A message to include with notifications for the monitor
   *
   * @schema CfnMonitorProps#Message
   */
  readonly message?: string;

  /**
   * Name of the monitor
   *
   * @schema CfnMonitorProps#Name
   */
  readonly name?: string;

  /**
   * Tags associated with the monitor
   *
   * @schema CfnMonitorProps#Tags
   */
  readonly tags?: string[];

  /**
   * Integer from 1 (high) to 5 (low) indicating alert severity.
   *
   * @schema CfnMonitorProps#Priority
   */
  readonly priority?: number;

  /**
   * The monitor options
   *
   * @schema CfnMonitorProps#Options
   */
  readonly options?: MonitorOptions;

  /**
   * The monitor query
   *
   * @schema CfnMonitorProps#Query
   */
  readonly query: string;

  /**
   * The type of the monitor
   *
   * @schema CfnMonitorProps#Type
   */
  readonly type: CfnMonitorPropsType;

  /**
   * Whether or not the monitor is multi alert
   *
   * @schema CfnMonitorProps#Multi
   */
  readonly multi?: boolean;

  /**
   * A list of unique role identifiers to define which roles are allowed to edit the monitor. The unique identifiers for all roles can be pulled from the [Roles API](https://docs.datadoghq.com/api/latest/roles/#list-roles) and are located in the `data.id` field. Editing a monitor includes any updates to the monitor configuration, monitor deletion, and muting of the monitor for any amount of time. `restricted_roles` is the successor of `locked`. For more information about `locked` and `restricted_roles`, see the [monitor options docs](https://docs.datadoghq.com/monitors/guide/monitor_api_options/#permissions-options).
   *
   * @schema CfnMonitorProps#RestrictedRoles
   */
  readonly restrictedRoles?: string[];

}

/**
 * Converts an object of type 'CfnMonitorProps' to JSON representation.
 */
/* eslint-disable max-len, quote-props */
export function toJson_CfnMonitorProps(obj: CfnMonitorProps | undefined): Record<string, any> | undefined {
  if (obj === undefined) { return undefined; }
  const result = {
    'Creator': toJson_Creator(obj.creator),
    'Message': obj.message,
    'Name': obj.name,
    'Tags': obj.tags?.map(y => y),
    'Priority': obj.priority,
    'Options': toJson_MonitorOptions(obj.options),
    'Query': obj.query,
    'Type': obj.type,
    'Multi': obj.multi,
    'RestrictedRoles': obj.restrictedRoles?.map(y => y),
  };
  // filter undefined values
  return Object.entries(result).reduce((r, i) => (i[1] === undefined) ? r : ({ ...r, [i[0]]: i[1] }), {});
}
/* eslint-enable max-len, quote-props */

/**
 * @schema Creator
 */
export interface Creator {
  /**
   * Name of the creator of the monitor
   *
   * @schema Creator#Name
   */
  readonly name?: string;

  /**
   * Handle of the creator of the monitor
   *
   * @schema Creator#Handle
   */
  readonly handle?: string;

  /**
   * Email of the creator of the monitor
   *
   * @schema Creator#Email
   */
  readonly email?: string;

}

/**
 * Converts an object of type 'Creator' to JSON representation.
 */
/* eslint-disable max-len, quote-props */
export function toJson_Creator(obj: Creator | undefined): Record<string, any> | undefined {
  if (obj === undefined) { return undefined; }
  const result = {
    'Name': obj.name,
    'Handle': obj.handle,
    'Email': obj.email,
  };
  // filter undefined values
  return Object.entries(result).reduce((r, i) => (i[1] === undefined) ? r : ({ ...r, [i[0]]: i[1] }), {});
}
/* eslint-enable max-len, quote-props */

/**
 * @schema MonitorOptions
 */
export interface MonitorOptions {
  /**
   * Whether or not to include a sample of the logs
   *
   * @schema MonitorOptions#EnableLogsSample
   */
  readonly enableLogsSample?: boolean;

  /**
   * Whether or not to send a list of samples when the monitor triggers. This is only used by CI Test and Pipeline monitors.
   *
   * @schema MonitorOptions#EnableSamples
   */
  readonly enableSamples?: boolean;

  /**
   * Message to include with a re-notification when renotify_interval is set
   *
   * @schema MonitorOptions#EscalationMessage
   */
  readonly escalationMessage?: string;

  /**
   * Time in seconds to delay evaluation
   *
   * @schema MonitorOptions#EvaluationDelay
   */
  readonly evaluationDelay?: number;

  /**
   * The time span after which groups with missing data are dropped from the monitor state.
   * The minimum value is one hour, and the maximum value is 72 hours.
   * Example values are: "60m", "1h", and "2d".
   * This option is only available for APM Trace Analytics, Audit Trail, CI, Error Tracking, Event, Logs, and RUM monitors.
   *
   * @schema MonitorOptions#GroupRetentionDuration
   */
  readonly groupRetentionDuration?: string;

  /**
   * Whether or not to include triggering tags into notification title
   *
   * @schema MonitorOptions#IncludeTags
   */
  readonly includeTags?: boolean;

  /**
   * Whether or not changes to this monitor should be restricted to the creator or admins
   *
   * @schema MonitorOptions#Locked
   */
  readonly locked?: boolean;

  /**
   * Number of locations allowed to fail before triggering alert
   *
   * @schema MonitorOptions#MinLocationFailed
   */
  readonly minLocationFailed?: number;

  /**
   * Time in seconds to allow a host to start reporting data before starting the evaluation of monitor results
   *
   * @schema MonitorOptions#NewHostDelay
   */
  readonly newHostDelay?: number;

  /**
   * Number of minutes data stopped reporting before notifying
   *
   * @schema MonitorOptions#NoDataTimeframe
   */
  readonly noDataTimeframe?: number;

  /**
   * Whether or not to notify tagged users when changes are made to the monitor
   *
   * @schema MonitorOptions#NotifyAudit
   */
  readonly notifyAudit?: boolean;

  /**
   * Controls what granularity a monitor alerts on. Only available for monitors with groupings.
   * For instance, a monitor grouped by `cluster`, `namespace`, and `pod` can be configured to only notify on each new `cluster` violating the alert conditions by setting `notify_by` to `["cluster"]`.
   * Tags mentioned in `notify_by` must be a subset of the grouping tags in the query.
   * For example, a query grouped by `cluster` and `namespace` cannot notify on `region`.
   * Setting `notify_by` to `[*]` configures the monitor to notify as a simple-alert.
   *
   * @schema MonitorOptions#NotifyBy
   */
  readonly notifyBy?: string[];

  /**
   * Whether or not to notify when data stops reporting
   *
   * @schema MonitorOptions#NotifyNoData
   */
  readonly notifyNoData?: boolean;

  /**
   * @schema MonitorOptions#NotificationPresetName
   */
  readonly notificationPresetName?: MonitorNotificationPresetName;

  /**
   * @schema MonitorOptions#OnMissingData
   */
  readonly onMissingData?: MonitorOnMissingData;

  /**
   * Number of minutes after the last notification before the monitor re-notifies on the current status
   *
   * @schema MonitorOptions#RenotifyInterval
   */
  readonly renotifyInterval?: number;

  /**
   * Whether or not the monitor requires a full window of data before it is evaluated
   *
   * @schema MonitorOptions#RequireFullWindow
   */
  readonly requireFullWindow?: boolean;

  /**
   * @schema MonitorOptions#SchedulingOptions
   */
  readonly schedulingOptions?: MonitorSchedulingOptions;

  /**
   * ID of the corresponding synthetics check
   *
   * @schema MonitorOptions#SyntheticsCheckID
   */
  readonly syntheticsCheckId?: number;

  /**
   * The threshold definitions
   *
   * @schema MonitorOptions#Thresholds
   */
  readonly thresholds?: MonitorThresholds;

  /**
   * The threshold window definitions
   *
   * @schema MonitorOptions#ThresholdWindows
   */
  readonly thresholdWindows?: MonitorThresholdWindows;

  /**
   * Number of hours of the monitor not reporting data before it automatically resolves
   *
   * @schema MonitorOptions#TimeoutH
   */
  readonly timeoutH?: number;

  /**
   * The number of times re-notification messages should be sent on the current status at the provided re-notification interval.
   *
   * @schema MonitorOptions#RenotifyOccurrences
   */
  readonly renotifyOccurrences?: number;

  /**
   * The types of monitor statuses for which re-notification messages are sent.
   *
   * @schema MonitorOptions#RenotifyStatuses
   */
  readonly renotifyStatuses?: MonitorOptionsRenotifyStatuses[];

  /**
   * How long the test should be in failure before alerting (integer, number of seconds, max 7200).
   *
   * @schema MonitorOptions#MinFailureDuration
   */
  readonly minFailureDuration?: number;

  /**
   * Time (in seconds) to skip evaluations for new groups. For example, this option can be used to skip evaluations for new hosts while they initialize. Must be a non negative integer.
   *
   * @schema MonitorOptions#NewGroupDelay
   */
  readonly newGroupDelay?: number;

  /**
   * List of requests that can be used in the monitor query.
   *
   * @schema MonitorOptions#Variables
   */
  readonly variables?: any[];

}

/**
 * Converts an object of type 'MonitorOptions' to JSON representation.
 */
/* eslint-disable max-len, quote-props */
export function toJson_MonitorOptions(obj: MonitorOptions | undefined): Record<string, any> | undefined {
  if (obj === undefined) { return undefined; }
  const result = {
    'EnableLogsSample': obj.enableLogsSample,
    'EnableSamples': obj.enableSamples,
    'EscalationMessage': obj.escalationMessage,
    'EvaluationDelay': obj.evaluationDelay,
    'GroupRetentionDuration': obj.groupRetentionDuration,
    'IncludeTags': obj.includeTags,
    'Locked': obj.locked,
    'MinLocationFailed': obj.minLocationFailed,
    'NewHostDelay': obj.newHostDelay,
    'NoDataTimeframe': obj.noDataTimeframe,
    'NotifyAudit': obj.notifyAudit,
    'NotifyBy': obj.notifyBy?.map(y => y),
    'NotifyNoData': obj.notifyNoData,
    'NotificationPresetName': obj.notificationPresetName,
    'OnMissingData': obj.onMissingData,
    'RenotifyInterval': obj.renotifyInterval,
    'RequireFullWindow': obj.requireFullWindow,
    'SchedulingOptions': toJson_MonitorSchedulingOptions(obj.schedulingOptions),
    'SyntheticsCheckID': obj.syntheticsCheckId,
    'Thresholds': toJson_MonitorThresholds(obj.thresholds),
    'ThresholdWindows': toJson_MonitorThresholdWindows(obj.thresholdWindows),
    'TimeoutH': obj.timeoutH,
    'RenotifyOccurrences': obj.renotifyOccurrences,
    'RenotifyStatuses': obj.renotifyStatuses?.map(y => y),
    'MinFailureDuration': obj.minFailureDuration,
    'NewGroupDelay': obj.newGroupDelay,
    'Variables': obj.variables?.map(y => y),
  };
  // filter undefined values
  return Object.entries(result).reduce((r, i) => (i[1] === undefined) ? r : ({ ...r, [i[0]]: i[1] }), {});
}
/* eslint-enable max-len, quote-props */

/**
 * The type of the monitor
 *
 * @schema CfnMonitorPropsType
 */
export enum CfnMonitorPropsType {
  /** audit alert */
  AUDIT_ALERT = 'audit alert',
  /** composite */
  COMPOSITE = 'composite',
  /** event alert */
  EVENT_ALERT = 'event alert',
  /** event-v2 alert */
  EVENT_V2_ALERT = 'event-v2 alert',
  /** log alert */
  LOG_ALERT = 'log alert',
  /** metric alert */
  METRIC_ALERT = 'metric alert',
  /** process alert */
  PROCESS_ALERT = 'process alert',
  /** query alert */
  QUERY_ALERT = 'query alert',
  /** service check */
  SERVICE_CHECK = 'service check',
  /** synthetics alert */
  SYNTHETICS_ALERT = 'synthetics alert',
  /** trace-analytics alert */
  TRACE_ANALYTICS_ALERT = 'trace-analytics alert',
  /** slo alert */
  SLO_ALERT = 'slo alert',
  /** rum alert */
  RUM_ALERT = 'rum alert',
  /** ci-pipelines alert */
  CI_PIPELINES_ALERT = 'ci-pipelines alert',
  /** error-tracking alert */
  ERROR_TRACKING_ALERT = 'error-tracking alert',
  /** ci-tests alert */
  CI_TESTS_ALERT = 'ci-tests alert',
}

/**
 * Toggles the display of additional content sent in the monitor notification.
 *
 * @schema MonitorNotificationPresetName
 */
export enum MonitorNotificationPresetName {
  /** show_all */
  SHOW_ALL = 'show_all',
  /** hide_query */
  HIDE_QUERY = 'hide_query',
  /** hide_handles */
  HIDE_HANDLES = 'hide_handles',
  /** hide_all */
  HIDE_ALL = 'hide_all',
}

/**
 * Controls how groups or monitors are treated if an evaluation does not return any data points.
 * The default option results in different behavior depending on the monitor query type.
 * For monitors using Count queries, an empty monitor evaluation is treated as 0 and is compared to the threshold conditions.
 * For monitors using any query type other than Count, for example Gauge, Measure, or Rate, the monitor shows the last known status.
 * This option is only available for APM Trace Analytics, Audit Trail, CI, Error Tracking, Event, Logs, and RUM monitors.
 *
 * @schema MonitorOnMissingData
 */
export enum MonitorOnMissingData {
  /** default */
  DEFAULT = 'default',
  /** show_no_data */
  SHOW_NO_DATA = 'show_no_data',
  /** show_and_notify_no_data */
  SHOW_AND_NOTIFY_NO_DATA = 'show_and_notify_no_data',
  /** resolve */
  RESOLVE = 'resolve',
}

/**
 * Configuration options for scheduling.
 *
 * @schema MonitorSchedulingOptions
 */
export interface MonitorSchedulingOptions {
  /**
   * @schema MonitorSchedulingOptions#EvaluationWindow
   */
  readonly evaluationWindow?: MonitorSchedulingOptionsEvaluationWindow;

}

/**
 * Converts an object of type 'MonitorSchedulingOptions' to JSON representation.
 */
/* eslint-disable max-len, quote-props */
export function toJson_MonitorSchedulingOptions(obj: MonitorSchedulingOptions | undefined): Record<string, any> | undefined {
  if (obj === undefined) { return undefined; }
  const result = {
    'EvaluationWindow': toJson_MonitorSchedulingOptionsEvaluationWindow(obj.evaluationWindow),
  };
  // filter undefined values
  return Object.entries(result).reduce((r, i) => (i[1] === undefined) ? r : ({ ...r, [i[0]]: i[1] }), {});
}
/* eslint-enable max-len, quote-props */

/**
 * @schema MonitorThresholds
 */
export interface MonitorThresholds {
  /**
   * Threshold value for triggering an alert
   *
   * @schema MonitorThresholds#Critical
   */
  readonly critical?: number;

  /**
   * Threshold value for recovering from an alert state
   *
   * @schema MonitorThresholds#CriticalRecovery
   */
  readonly criticalRecovery?: number;

  /**
   * Threshold value for recovering from an alert state
   *
   * @schema MonitorThresholds#OK
   */
  readonly ok?: number;

  /**
   * Threshold value for triggering a warning
   *
   * @schema MonitorThresholds#Warning
   */
  readonly warning?: number;

  /**
   * Threshold value for recovering from a warning state
   *
   * @schema MonitorThresholds#WarningRecovery
   */
  readonly warningRecovery?: number;

}

/**
 * Converts an object of type 'MonitorThresholds' to JSON representation.
 */
/* eslint-disable max-len, quote-props */
export function toJson_MonitorThresholds(obj: MonitorThresholds | undefined): Record<string, any> | undefined {
  if (obj === undefined) { return undefined; }
  const result = {
    'Critical': obj.critical,
    'CriticalRecovery': obj.criticalRecovery,
    'OK': obj.ok,
    'Warning': obj.warning,
    'WarningRecovery': obj.warningRecovery,
  };
  // filter undefined values
  return Object.entries(result).reduce((r, i) => (i[1] === undefined) ? r : ({ ...r, [i[0]]: i[1] }), {});
}
/* eslint-enable max-len, quote-props */

/**
 * @schema MonitorThresholdWindows
 */
export interface MonitorThresholdWindows {
  /**
   * How long a metric must be anomalous before triggering an alert
   *
   * @schema MonitorThresholdWindows#TriggerWindow
   */
  readonly triggerWindow?: string;

  /**
   * How long an anomalous metric must be normal before recovering from an alert state
   *
   * @schema MonitorThresholdWindows#RecoveryWindow
   */
  readonly recoveryWindow?: string;

}

/**
 * Converts an object of type 'MonitorThresholdWindows' to JSON representation.
 */
/* eslint-disable max-len, quote-props */
export function toJson_MonitorThresholdWindows(obj: MonitorThresholdWindows | undefined): Record<string, any> | undefined {
  if (obj === undefined) { return undefined; }
  const result = {
    'TriggerWindow': obj.triggerWindow,
    'RecoveryWindow': obj.recoveryWindow,
  };
  // filter undefined values
  return Object.entries(result).reduce((r, i) => (i[1] === undefined) ? r : ({ ...r, [i[0]]: i[1] }), {});
}
/* eslint-enable max-len, quote-props */

/**
 * @schema MonitorOptionsRenotifyStatuses
 */
export enum MonitorOptionsRenotifyStatuses {
  /** alert */
  ALERT = 'alert',
  /** no data */
  NO_DATA = 'no data',
  /** warn */
  WARN = 'warn',
}

/**
 * Configuration options for the evaluation window. If `hour_starts` is set, no other fields may be set. Otherwise, `day_starts` and `month_starts` must be set together.
 *
 * @schema MonitorSchedulingOptionsEvaluationWindow
 */
export interface MonitorSchedulingOptionsEvaluationWindow {
  /**
   * The time of the day at which a one day cumulative evaluation window starts. Must be defined in UTC time in `HH:mm` format.
   *
   * @schema MonitorSchedulingOptionsEvaluationWindow#DayStarts
   */
  readonly dayStarts?: string;

  /**
   * The day of the month at which a one month cumulative evaluation window starts.
   *
   * @schema MonitorSchedulingOptionsEvaluationWindow#MonthStarts
   */
  readonly monthStarts?: number;

  /**
   * The minute of the hour at which a one hour cumulative evaluation window starts.
   *
   * @schema MonitorSchedulingOptionsEvaluationWindow#HourStarts
   */
  readonly hourStarts?: number;

}

/**
 * Converts an object of type 'MonitorSchedulingOptionsEvaluationWindow' to JSON representation.
 */
/* eslint-disable max-len, quote-props */
export function toJson_MonitorSchedulingOptionsEvaluationWindow(obj: MonitorSchedulingOptionsEvaluationWindow | undefined): Record<string, any> | undefined {
  if (obj === undefined) { return undefined; }
  const result = {
    'DayStarts': obj.dayStarts,
    'MonthStarts': obj.monthStarts,
    'HourStarts': obj.hourStarts,
  };
  // filter undefined values
  return Object.entries(result).reduce((r, i) => (i[1] === undefined) ? r : ({ ...r, [i[0]]: i[1] }), {});
}
/* eslint-enable max-len, quote-props */


/**
 * A CloudFormation `Datadog::Monitors::Monitor`
 *
 * @cloudformationResource Datadog::Monitors::Monitor
 * @stability external
 * @link http://unknown-url
 */
export class CfnMonitor extends cdk.CfnResource {
  /**
  * The CloudFormation resource type name for this resource class.
  */
  public static readonly CFN_RESOURCE_TYPE_NAME = 'Datadog::Monitors::Monitor';

  /**
   * Resource props.
   */
  public readonly props: CfnMonitorProps;

  /**
   * Attribute `Datadog::Monitors::Monitor.Modified`
   * @link http://unknown-url
   */
  public readonly attrModified: string;
  /**
   * Attribute `Datadog::Monitors::Monitor.Id`
   * @link http://unknown-url
   */
  public readonly attrId: number;
  /**
   * Attribute `Datadog::Monitors::Monitor.Deleted`
   * @link http://unknown-url
   */
  public readonly attrDeleted: string;
  /**
   * Attribute `Datadog::Monitors::Monitor.Created`
   * @link http://unknown-url
   */
  public readonly attrCreated: string;

  /**
   * Create a new `Datadog::Monitors::Monitor`.
   *
   * @param scope - scope in which this resource is defined
   * @param id    - scoped id of the resource
   * @param props - resource properties
   */
  constructor(scope: constructs.Construct, id: string, props: CfnMonitorProps) {
    super(scope, id, { type: CfnMonitor.CFN_RESOURCE_TYPE_NAME, properties: toJson_CfnMonitorProps(props)! });

    this.props = props;

    this.attrModified = cdk.Token.asString(this.getAtt('Modified'));
    this.attrId = cdk.Token.asNumber(this.getAtt('Id'));
    this.attrDeleted = cdk.Token.asString(this.getAtt('Deleted'));
    this.attrCreated = cdk.Token.asString(this.getAtt('Created'));
  }
}
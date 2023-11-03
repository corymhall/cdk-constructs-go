# API Reference <a name="API Reference" id="api-reference"></a>

## Constructs <a name="Constructs" id="Constructs"></a>

### CfnDashboard <a name="CfnDashboard" id="cdk-constructs-datadog-go.CfnDashboard"></a>

A CloudFormation `Datadog::Dashboards::Dashboard`.

> [arn:aws:cloudformation:us-east-1::type/resource/7171b96e5d207b947eb72ca9ce05247c246de623/Datadog-Dashboards-Dashboard](arn:aws:cloudformation:us-east-1::type/resource/7171b96e5d207b947eb72ca9ce05247c246de623/Datadog-Dashboards-Dashboard)

#### Initializers <a name="Initializers" id="cdk-constructs-datadog-go.CfnDashboard.Initializer"></a>

```typescript
import { CfnDashboard } from 'cdk-constructs-datadog-go'

new CfnDashboard(scope: Construct, id: string, props: CfnDashboardProps)
```

| **Name** | **Type** | **Description** |
| --- | --- | --- |
| <code><a href="#cdk-constructs-datadog-go.CfnDashboard.Initializer.parameter.scope">scope</a></code> | <code>constructs.Construct</code> | - scope in which this resource is defined. |
| <code><a href="#cdk-constructs-datadog-go.CfnDashboard.Initializer.parameter.id">id</a></code> | <code>string</code> | - scoped id of the resource. |
| <code><a href="#cdk-constructs-datadog-go.CfnDashboard.Initializer.parameter.props">props</a></code> | <code><a href="#cdk-constructs-datadog-go.CfnDashboardProps">CfnDashboardProps</a></code> | - resource properties. |

---

##### `scope`<sup>Required</sup> <a name="scope" id="cdk-constructs-datadog-go.CfnDashboard.Initializer.parameter.scope"></a>

- *Type:* constructs.Construct

scope in which this resource is defined.

---

##### `id`<sup>Required</sup> <a name="id" id="cdk-constructs-datadog-go.CfnDashboard.Initializer.parameter.id"></a>

- *Type:* string

scoped id of the resource.

---

##### `props`<sup>Required</sup> <a name="props" id="cdk-constructs-datadog-go.CfnDashboard.Initializer.parameter.props"></a>

- *Type:* <a href="#cdk-constructs-datadog-go.CfnDashboardProps">CfnDashboardProps</a>

resource properties.

---

#### Methods <a name="Methods" id="Methods"></a>

| **Name** | **Description** |
| --- | --- |
| <code><a href="#cdk-constructs-datadog-go.CfnDashboard.toString">toString</a></code> | Returns a string representation of this construct. |
| <code><a href="#cdk-constructs-datadog-go.CfnDashboard.overrideLogicalId">overrideLogicalId</a></code> | Overrides the auto-generated logical ID with a specific ID. |
| <code><a href="#cdk-constructs-datadog-go.CfnDashboard.addDeletionOverride">addDeletionOverride</a></code> | Syntactic sugar for `addOverride(path, undefined)`. |
| <code><a href="#cdk-constructs-datadog-go.CfnDashboard.addDependency">addDependency</a></code> | Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned. |
| <code><a href="#cdk-constructs-datadog-go.CfnDashboard.addDependsOn">addDependsOn</a></code> | Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned. |
| <code><a href="#cdk-constructs-datadog-go.CfnDashboard.addMetadata">addMetadata</a></code> | Add a value to the CloudFormation Resource Metadata. |
| <code><a href="#cdk-constructs-datadog-go.CfnDashboard.addOverride">addOverride</a></code> | Adds an override to the synthesized CloudFormation resource. |
| <code><a href="#cdk-constructs-datadog-go.CfnDashboard.addPropertyDeletionOverride">addPropertyDeletionOverride</a></code> | Adds an override that deletes the value of a property from the resource definition. |
| <code><a href="#cdk-constructs-datadog-go.CfnDashboard.addPropertyOverride">addPropertyOverride</a></code> | Adds an override to a resource property. |
| <code><a href="#cdk-constructs-datadog-go.CfnDashboard.applyRemovalPolicy">applyRemovalPolicy</a></code> | Sets the deletion policy of the resource based on the removal policy specified. |
| <code><a href="#cdk-constructs-datadog-go.CfnDashboard.getAtt">getAtt</a></code> | Returns a token for an runtime attribute of this resource. |
| <code><a href="#cdk-constructs-datadog-go.CfnDashboard.getMetadata">getMetadata</a></code> | Retrieve a value value from the CloudFormation Resource Metadata. |
| <code><a href="#cdk-constructs-datadog-go.CfnDashboard.obtainDependencies">obtainDependencies</a></code> | Retrieves an array of resources this resource depends on. |
| <code><a href="#cdk-constructs-datadog-go.CfnDashboard.obtainResourceDependencies">obtainResourceDependencies</a></code> | Get a shallow copy of dependencies between this resource and other resources in the same stack. |
| <code><a href="#cdk-constructs-datadog-go.CfnDashboard.removeDependency">removeDependency</a></code> | Indicates that this resource no longer depends on another resource. |
| <code><a href="#cdk-constructs-datadog-go.CfnDashboard.replaceDependency">replaceDependency</a></code> | Replaces one dependency with another. |

---

##### `toString` <a name="toString" id="cdk-constructs-datadog-go.CfnDashboard.toString"></a>

```typescript
public toString(): string
```

Returns a string representation of this construct.

##### `overrideLogicalId` <a name="overrideLogicalId" id="cdk-constructs-datadog-go.CfnDashboard.overrideLogicalId"></a>

```typescript
public overrideLogicalId(newLogicalId: string): void
```

Overrides the auto-generated logical ID with a specific ID.

###### `newLogicalId`<sup>Required</sup> <a name="newLogicalId" id="cdk-constructs-datadog-go.CfnDashboard.overrideLogicalId.parameter.newLogicalId"></a>

- *Type:* string

The new logical ID to use for this stack element.

---

##### `addDeletionOverride` <a name="addDeletionOverride" id="cdk-constructs-datadog-go.CfnDashboard.addDeletionOverride"></a>

```typescript
public addDeletionOverride(path: string): void
```

Syntactic sugar for `addOverride(path, undefined)`.

###### `path`<sup>Required</sup> <a name="path" id="cdk-constructs-datadog-go.CfnDashboard.addDeletionOverride.parameter.path"></a>

- *Type:* string

The path of the value to delete.

---

##### `addDependency` <a name="addDependency" id="cdk-constructs-datadog-go.CfnDashboard.addDependency"></a>

```typescript
public addDependency(target: CfnResource): void
```

Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.

This can be used for resources across stacks (or nested stack) boundaries
and the dependency will automatically be transferred to the relevant scope.

###### `target`<sup>Required</sup> <a name="target" id="cdk-constructs-datadog-go.CfnDashboard.addDependency.parameter.target"></a>

- *Type:* aws-cdk-lib.CfnResource

---

##### ~~`addDependsOn`~~ <a name="addDependsOn" id="cdk-constructs-datadog-go.CfnDashboard.addDependsOn"></a>

```typescript
public addDependsOn(target: CfnResource): void
```

Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.

###### `target`<sup>Required</sup> <a name="target" id="cdk-constructs-datadog-go.CfnDashboard.addDependsOn.parameter.target"></a>

- *Type:* aws-cdk-lib.CfnResource

---

##### `addMetadata` <a name="addMetadata" id="cdk-constructs-datadog-go.CfnDashboard.addMetadata"></a>

```typescript
public addMetadata(key: string, value: any): void
```

Add a value to the CloudFormation Resource Metadata.

> [https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html

Note that this is a different set of metadata from CDK node metadata; this
metadata ends up in the stack template under the resource, whereas CDK
node metadata ends up in the Cloud Assembly.](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html

Note that this is a different set of metadata from CDK node metadata; this
metadata ends up in the stack template under the resource, whereas CDK
node metadata ends up in the Cloud Assembly.)

###### `key`<sup>Required</sup> <a name="key" id="cdk-constructs-datadog-go.CfnDashboard.addMetadata.parameter.key"></a>

- *Type:* string

---

###### `value`<sup>Required</sup> <a name="value" id="cdk-constructs-datadog-go.CfnDashboard.addMetadata.parameter.value"></a>

- *Type:* any

---

##### `addOverride` <a name="addOverride" id="cdk-constructs-datadog-go.CfnDashboard.addOverride"></a>

```typescript
public addOverride(path: string, value: any): void
```

Adds an override to the synthesized CloudFormation resource.

To add a
property override, either use `addPropertyOverride` or prefix `path` with
"Properties." (i.e. `Properties.TopicName`).

If the override is nested, separate each nested level using a dot (.) in the path parameter.
If there is an array as part of the nesting, specify the index in the path.

To include a literal `.` in the property name, prefix with a `\`. In most
programming languages you will need to write this as `"\\."` because the
`\` itself will need to be escaped.

For example,
```typescript
cfnResource.addOverride('Properties.GlobalSecondaryIndexes.0.Projection.NonKeyAttributes', ['myattribute']);
cfnResource.addOverride('Properties.GlobalSecondaryIndexes.1.ProjectionType', 'INCLUDE');
```
would add the overrides
```json
"Properties": {
  "GlobalSecondaryIndexes": [
    {
      "Projection": {
        "NonKeyAttributes": [ "myattribute" ]
        ...
      }
      ...
    },
    {
      "ProjectionType": "INCLUDE"
      ...
    },
  ]
  ...
}
```

The `value` argument to `addOverride` will not be processed or translated
in any way. Pass raw JSON values in here with the correct capitalization
for CloudFormation. If you pass CDK classes or structs, they will be
rendered with lowercased key names, and CloudFormation will reject the
template.

###### `path`<sup>Required</sup> <a name="path" id="cdk-constructs-datadog-go.CfnDashboard.addOverride.parameter.path"></a>

- *Type:* string

The path of the property, you can use dot notation to override values in complex types.

Any intermediate keys
will be created as needed.

---

###### `value`<sup>Required</sup> <a name="value" id="cdk-constructs-datadog-go.CfnDashboard.addOverride.parameter.value"></a>

- *Type:* any

The value.

Could be primitive or complex.

---

##### `addPropertyDeletionOverride` <a name="addPropertyDeletionOverride" id="cdk-constructs-datadog-go.CfnDashboard.addPropertyDeletionOverride"></a>

```typescript
public addPropertyDeletionOverride(propertyPath: string): void
```

Adds an override that deletes the value of a property from the resource definition.

###### `propertyPath`<sup>Required</sup> <a name="propertyPath" id="cdk-constructs-datadog-go.CfnDashboard.addPropertyDeletionOverride.parameter.propertyPath"></a>

- *Type:* string

The path to the property.

---

##### `addPropertyOverride` <a name="addPropertyOverride" id="cdk-constructs-datadog-go.CfnDashboard.addPropertyOverride"></a>

```typescript
public addPropertyOverride(propertyPath: string, value: any): void
```

Adds an override to a resource property.

Syntactic sugar for `addOverride("Properties.<...>", value)`.

###### `propertyPath`<sup>Required</sup> <a name="propertyPath" id="cdk-constructs-datadog-go.CfnDashboard.addPropertyOverride.parameter.propertyPath"></a>

- *Type:* string

The path of the property.

---

###### `value`<sup>Required</sup> <a name="value" id="cdk-constructs-datadog-go.CfnDashboard.addPropertyOverride.parameter.value"></a>

- *Type:* any

The value.

---

##### `applyRemovalPolicy` <a name="applyRemovalPolicy" id="cdk-constructs-datadog-go.CfnDashboard.applyRemovalPolicy"></a>

```typescript
public applyRemovalPolicy(policy?: RemovalPolicy, options?: RemovalPolicyOptions): void
```

Sets the deletion policy of the resource based on the removal policy specified.

The Removal Policy controls what happens to this resource when it stops
being managed by CloudFormation, either because you've removed it from the
CDK application or because you've made a change that requires the resource
to be replaced.

The resource can be deleted (`RemovalPolicy.DESTROY`), or left in your AWS
account for data recovery and cleanup later (`RemovalPolicy.RETAIN`). In some
cases, a snapshot can be taken of the resource prior to deletion
(`RemovalPolicy.SNAPSHOT`). A list of resources that support this policy
can be found in the following link:

> [https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-deletionpolicy.html#aws-attribute-deletionpolicy-options](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-deletionpolicy.html#aws-attribute-deletionpolicy-options)

###### `policy`<sup>Optional</sup> <a name="policy" id="cdk-constructs-datadog-go.CfnDashboard.applyRemovalPolicy.parameter.policy"></a>

- *Type:* aws-cdk-lib.RemovalPolicy

---

###### `options`<sup>Optional</sup> <a name="options" id="cdk-constructs-datadog-go.CfnDashboard.applyRemovalPolicy.parameter.options"></a>

- *Type:* aws-cdk-lib.RemovalPolicyOptions

---

##### `getAtt` <a name="getAtt" id="cdk-constructs-datadog-go.CfnDashboard.getAtt"></a>

```typescript
public getAtt(attributeName: string, typeHint?: ResolutionTypeHint): Reference
```

Returns a token for an runtime attribute of this resource.

Ideally, use generated attribute accessors (e.g. `resource.arn`), but this can be used for future compatibility
in case there is no generated attribute.

###### `attributeName`<sup>Required</sup> <a name="attributeName" id="cdk-constructs-datadog-go.CfnDashboard.getAtt.parameter.attributeName"></a>

- *Type:* string

The name of the attribute.

---

###### `typeHint`<sup>Optional</sup> <a name="typeHint" id="cdk-constructs-datadog-go.CfnDashboard.getAtt.parameter.typeHint"></a>

- *Type:* aws-cdk-lib.ResolutionTypeHint

---

##### `getMetadata` <a name="getMetadata" id="cdk-constructs-datadog-go.CfnDashboard.getMetadata"></a>

```typescript
public getMetadata(key: string): any
```

Retrieve a value value from the CloudFormation Resource Metadata.

> [https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html

Note that this is a different set of metadata from CDK node metadata; this
metadata ends up in the stack template under the resource, whereas CDK
node metadata ends up in the Cloud Assembly.](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html

Note that this is a different set of metadata from CDK node metadata; this
metadata ends up in the stack template under the resource, whereas CDK
node metadata ends up in the Cloud Assembly.)

###### `key`<sup>Required</sup> <a name="key" id="cdk-constructs-datadog-go.CfnDashboard.getMetadata.parameter.key"></a>

- *Type:* string

---

##### `obtainDependencies` <a name="obtainDependencies" id="cdk-constructs-datadog-go.CfnDashboard.obtainDependencies"></a>

```typescript
public obtainDependencies(): Stack | CfnResource[]
```

Retrieves an array of resources this resource depends on.

This assembles dependencies on resources across stacks (including nested stacks)
automatically.

##### `obtainResourceDependencies` <a name="obtainResourceDependencies" id="cdk-constructs-datadog-go.CfnDashboard.obtainResourceDependencies"></a>

```typescript
public obtainResourceDependencies(): CfnResource[]
```

Get a shallow copy of dependencies between this resource and other resources in the same stack.

##### `removeDependency` <a name="removeDependency" id="cdk-constructs-datadog-go.CfnDashboard.removeDependency"></a>

```typescript
public removeDependency(target: CfnResource): void
```

Indicates that this resource no longer depends on another resource.

This can be used for resources across stacks (including nested stacks)
and the dependency will automatically be removed from the relevant scope.

###### `target`<sup>Required</sup> <a name="target" id="cdk-constructs-datadog-go.CfnDashboard.removeDependency.parameter.target"></a>

- *Type:* aws-cdk-lib.CfnResource

---

##### `replaceDependency` <a name="replaceDependency" id="cdk-constructs-datadog-go.CfnDashboard.replaceDependency"></a>

```typescript
public replaceDependency(target: CfnResource, newTarget: CfnResource): void
```

Replaces one dependency with another.

###### `target`<sup>Required</sup> <a name="target" id="cdk-constructs-datadog-go.CfnDashboard.replaceDependency.parameter.target"></a>

- *Type:* aws-cdk-lib.CfnResource

The dependency to replace.

---

###### `newTarget`<sup>Required</sup> <a name="newTarget" id="cdk-constructs-datadog-go.CfnDashboard.replaceDependency.parameter.newTarget"></a>

- *Type:* aws-cdk-lib.CfnResource

The new dependency to add.

---

#### Static Functions <a name="Static Functions" id="Static Functions"></a>

| **Name** | **Description** |
| --- | --- |
| <code><a href="#cdk-constructs-datadog-go.CfnDashboard.isConstruct">isConstruct</a></code> | Checks if `x` is a construct. |
| <code><a href="#cdk-constructs-datadog-go.CfnDashboard.isCfnElement">isCfnElement</a></code> | Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template). |
| <code><a href="#cdk-constructs-datadog-go.CfnDashboard.isCfnResource">isCfnResource</a></code> | Check whether the given construct is a CfnResource. |

---

##### ~~`isConstruct`~~ <a name="isConstruct" id="cdk-constructs-datadog-go.CfnDashboard.isConstruct"></a>

```typescript
import { CfnDashboard } from 'cdk-constructs-datadog-go'

CfnDashboard.isConstruct(x: any)
```

Checks if `x` is a construct.

###### `x`<sup>Required</sup> <a name="x" id="cdk-constructs-datadog-go.CfnDashboard.isConstruct.parameter.x"></a>

- *Type:* any

Any object.

---

##### `isCfnElement` <a name="isCfnElement" id="cdk-constructs-datadog-go.CfnDashboard.isCfnElement"></a>

```typescript
import { CfnDashboard } from 'cdk-constructs-datadog-go'

CfnDashboard.isCfnElement(x: any)
```

Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).

Uses duck-typing instead of `instanceof` to allow stack elements from different
versions of this library to be included in the same stack.

###### `x`<sup>Required</sup> <a name="x" id="cdk-constructs-datadog-go.CfnDashboard.isCfnElement.parameter.x"></a>

- *Type:* any

---

##### `isCfnResource` <a name="isCfnResource" id="cdk-constructs-datadog-go.CfnDashboard.isCfnResource"></a>

```typescript
import { CfnDashboard } from 'cdk-constructs-datadog-go'

CfnDashboard.isCfnResource(construct: IConstruct)
```

Check whether the given construct is a CfnResource.

###### `construct`<sup>Required</sup> <a name="construct" id="cdk-constructs-datadog-go.CfnDashboard.isCfnResource.parameter.construct"></a>

- *Type:* constructs.IConstruct

---

#### Properties <a name="Properties" id="Properties"></a>

| **Name** | **Type** | **Description** |
| --- | --- | --- |
| <code><a href="#cdk-constructs-datadog-go.CfnDashboard.property.node">node</a></code> | <code>constructs.Node</code> | The tree node. |
| <code><a href="#cdk-constructs-datadog-go.CfnDashboard.property.creationStack">creationStack</a></code> | <code>string[]</code> | *No description.* |
| <code><a href="#cdk-constructs-datadog-go.CfnDashboard.property.logicalId">logicalId</a></code> | <code>string</code> | The logical ID for this CloudFormation stack element. |
| <code><a href="#cdk-constructs-datadog-go.CfnDashboard.property.stack">stack</a></code> | <code>aws-cdk-lib.Stack</code> | The stack in which this element is defined. |
| <code><a href="#cdk-constructs-datadog-go.CfnDashboard.property.ref">ref</a></code> | <code>string</code> | Return a string that will be resolved to a CloudFormation `{ Ref }` for this element. |
| <code><a href="#cdk-constructs-datadog-go.CfnDashboard.property.cfnOptions">cfnOptions</a></code> | <code>aws-cdk-lib.ICfnResourceOptions</code> | Options for this resource, such as condition, update policy etc. |
| <code><a href="#cdk-constructs-datadog-go.CfnDashboard.property.cfnResourceType">cfnResourceType</a></code> | <code>string</code> | AWS resource type. |
| <code><a href="#cdk-constructs-datadog-go.CfnDashboard.property.attrId">attrId</a></code> | <code>string</code> | Attribute `Datadog::Dashboards::Dashboard.Id`. |
| <code><a href="#cdk-constructs-datadog-go.CfnDashboard.property.attrUrl">attrUrl</a></code> | <code>string</code> | Attribute `Datadog::Dashboards::Dashboard.Url`. |
| <code><a href="#cdk-constructs-datadog-go.CfnDashboard.property.props">props</a></code> | <code><a href="#cdk-constructs-datadog-go.CfnDashboardProps">CfnDashboardProps</a></code> | Resource props. |

---

##### `node`<sup>Required</sup> <a name="node" id="cdk-constructs-datadog-go.CfnDashboard.property.node"></a>

```typescript
public readonly node: Node;
```

- *Type:* constructs.Node

The tree node.

---

##### `creationStack`<sup>Required</sup> <a name="creationStack" id="cdk-constructs-datadog-go.CfnDashboard.property.creationStack"></a>

```typescript
public readonly creationStack: string[];
```

- *Type:* string[]

---

##### `logicalId`<sup>Required</sup> <a name="logicalId" id="cdk-constructs-datadog-go.CfnDashboard.property.logicalId"></a>

```typescript
public readonly logicalId: string;
```

- *Type:* string

The logical ID for this CloudFormation stack element.

The logical ID of the element
is calculated from the path of the resource node in the construct tree.

To override this value, use `overrideLogicalId(newLogicalId)`.

---

##### `stack`<sup>Required</sup> <a name="stack" id="cdk-constructs-datadog-go.CfnDashboard.property.stack"></a>

```typescript
public readonly stack: Stack;
```

- *Type:* aws-cdk-lib.Stack

The stack in which this element is defined.

CfnElements must be defined within a stack scope (directly or indirectly).

---

##### `ref`<sup>Required</sup> <a name="ref" id="cdk-constructs-datadog-go.CfnDashboard.property.ref"></a>

```typescript
public readonly ref: string;
```

- *Type:* string

Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.

If, by any chance, the intrinsic reference of a resource is not a string, you could
coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.

---

##### `cfnOptions`<sup>Required</sup> <a name="cfnOptions" id="cdk-constructs-datadog-go.CfnDashboard.property.cfnOptions"></a>

```typescript
public readonly cfnOptions: ICfnResourceOptions;
```

- *Type:* aws-cdk-lib.ICfnResourceOptions

Options for this resource, such as condition, update policy etc.

---

##### `cfnResourceType`<sup>Required</sup> <a name="cfnResourceType" id="cdk-constructs-datadog-go.CfnDashboard.property.cfnResourceType"></a>

```typescript
public readonly cfnResourceType: string;
```

- *Type:* string

AWS resource type.

---

##### `attrId`<sup>Required</sup> <a name="attrId" id="cdk-constructs-datadog-go.CfnDashboard.property.attrId"></a>

```typescript
public readonly attrId: string;
```

- *Type:* string

Attribute `Datadog::Dashboards::Dashboard.Id`.

> [arn:aws:cloudformation:us-east-1::type/resource/7171b96e5d207b947eb72ca9ce05247c246de623/Datadog-Dashboards-Dashboard](arn:aws:cloudformation:us-east-1::type/resource/7171b96e5d207b947eb72ca9ce05247c246de623/Datadog-Dashboards-Dashboard)

---

##### `attrUrl`<sup>Required</sup> <a name="attrUrl" id="cdk-constructs-datadog-go.CfnDashboard.property.attrUrl"></a>

```typescript
public readonly attrUrl: string;
```

- *Type:* string

Attribute `Datadog::Dashboards::Dashboard.Url`.

> [arn:aws:cloudformation:us-east-1::type/resource/7171b96e5d207b947eb72ca9ce05247c246de623/Datadog-Dashboards-Dashboard](arn:aws:cloudformation:us-east-1::type/resource/7171b96e5d207b947eb72ca9ce05247c246de623/Datadog-Dashboards-Dashboard)

---

##### `props`<sup>Required</sup> <a name="props" id="cdk-constructs-datadog-go.CfnDashboard.property.props"></a>

```typescript
public readonly props: CfnDashboardProps;
```

- *Type:* <a href="#cdk-constructs-datadog-go.CfnDashboardProps">CfnDashboardProps</a>

Resource props.

---

#### Constants <a name="Constants" id="Constants"></a>

| **Name** | **Type** | **Description** |
| --- | --- | --- |
| <code><a href="#cdk-constructs-datadog-go.CfnDashboard.property.CFN_RESOURCE_TYPE_NAME">CFN_RESOURCE_TYPE_NAME</a></code> | <code>string</code> | The CloudFormation resource type name for this resource class. |

---

##### `CFN_RESOURCE_TYPE_NAME`<sup>Required</sup> <a name="CFN_RESOURCE_TYPE_NAME" id="cdk-constructs-datadog-go.CfnDashboard.property.CFN_RESOURCE_TYPE_NAME"></a>

```typescript
public readonly CFN_RESOURCE_TYPE_NAME: string;
```

- *Type:* string

The CloudFormation resource type name for this resource class.

---

### CfnMonitor <a name="CfnMonitor" id="cdk-constructs-datadog-go.CfnMonitor"></a>

A CloudFormation `Datadog::Monitors::Monitor`.

> [http://unknown-url](http://unknown-url)

#### Initializers <a name="Initializers" id="cdk-constructs-datadog-go.CfnMonitor.Initializer"></a>

```typescript
import { CfnMonitor } from 'cdk-constructs-datadog-go'

new CfnMonitor(scope: Construct, id: string, props: CfnMonitorProps)
```

| **Name** | **Type** | **Description** |
| --- | --- | --- |
| <code><a href="#cdk-constructs-datadog-go.CfnMonitor.Initializer.parameter.scope">scope</a></code> | <code>constructs.Construct</code> | - scope in which this resource is defined. |
| <code><a href="#cdk-constructs-datadog-go.CfnMonitor.Initializer.parameter.id">id</a></code> | <code>string</code> | - scoped id of the resource. |
| <code><a href="#cdk-constructs-datadog-go.CfnMonitor.Initializer.parameter.props">props</a></code> | <code><a href="#cdk-constructs-datadog-go.CfnMonitorProps">CfnMonitorProps</a></code> | - resource properties. |

---

##### `scope`<sup>Required</sup> <a name="scope" id="cdk-constructs-datadog-go.CfnMonitor.Initializer.parameter.scope"></a>

- *Type:* constructs.Construct

scope in which this resource is defined.

---

##### `id`<sup>Required</sup> <a name="id" id="cdk-constructs-datadog-go.CfnMonitor.Initializer.parameter.id"></a>

- *Type:* string

scoped id of the resource.

---

##### `props`<sup>Required</sup> <a name="props" id="cdk-constructs-datadog-go.CfnMonitor.Initializer.parameter.props"></a>

- *Type:* <a href="#cdk-constructs-datadog-go.CfnMonitorProps">CfnMonitorProps</a>

resource properties.

---

#### Methods <a name="Methods" id="Methods"></a>

| **Name** | **Description** |
| --- | --- |
| <code><a href="#cdk-constructs-datadog-go.CfnMonitor.toString">toString</a></code> | Returns a string representation of this construct. |
| <code><a href="#cdk-constructs-datadog-go.CfnMonitor.overrideLogicalId">overrideLogicalId</a></code> | Overrides the auto-generated logical ID with a specific ID. |
| <code><a href="#cdk-constructs-datadog-go.CfnMonitor.addDeletionOverride">addDeletionOverride</a></code> | Syntactic sugar for `addOverride(path, undefined)`. |
| <code><a href="#cdk-constructs-datadog-go.CfnMonitor.addDependency">addDependency</a></code> | Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned. |
| <code><a href="#cdk-constructs-datadog-go.CfnMonitor.addDependsOn">addDependsOn</a></code> | Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned. |
| <code><a href="#cdk-constructs-datadog-go.CfnMonitor.addMetadata">addMetadata</a></code> | Add a value to the CloudFormation Resource Metadata. |
| <code><a href="#cdk-constructs-datadog-go.CfnMonitor.addOverride">addOverride</a></code> | Adds an override to the synthesized CloudFormation resource. |
| <code><a href="#cdk-constructs-datadog-go.CfnMonitor.addPropertyDeletionOverride">addPropertyDeletionOverride</a></code> | Adds an override that deletes the value of a property from the resource definition. |
| <code><a href="#cdk-constructs-datadog-go.CfnMonitor.addPropertyOverride">addPropertyOverride</a></code> | Adds an override to a resource property. |
| <code><a href="#cdk-constructs-datadog-go.CfnMonitor.applyRemovalPolicy">applyRemovalPolicy</a></code> | Sets the deletion policy of the resource based on the removal policy specified. |
| <code><a href="#cdk-constructs-datadog-go.CfnMonitor.getAtt">getAtt</a></code> | Returns a token for an runtime attribute of this resource. |
| <code><a href="#cdk-constructs-datadog-go.CfnMonitor.getMetadata">getMetadata</a></code> | Retrieve a value value from the CloudFormation Resource Metadata. |
| <code><a href="#cdk-constructs-datadog-go.CfnMonitor.obtainDependencies">obtainDependencies</a></code> | Retrieves an array of resources this resource depends on. |
| <code><a href="#cdk-constructs-datadog-go.CfnMonitor.obtainResourceDependencies">obtainResourceDependencies</a></code> | Get a shallow copy of dependencies between this resource and other resources in the same stack. |
| <code><a href="#cdk-constructs-datadog-go.CfnMonitor.removeDependency">removeDependency</a></code> | Indicates that this resource no longer depends on another resource. |
| <code><a href="#cdk-constructs-datadog-go.CfnMonitor.replaceDependency">replaceDependency</a></code> | Replaces one dependency with another. |

---

##### `toString` <a name="toString" id="cdk-constructs-datadog-go.CfnMonitor.toString"></a>

```typescript
public toString(): string
```

Returns a string representation of this construct.

##### `overrideLogicalId` <a name="overrideLogicalId" id="cdk-constructs-datadog-go.CfnMonitor.overrideLogicalId"></a>

```typescript
public overrideLogicalId(newLogicalId: string): void
```

Overrides the auto-generated logical ID with a specific ID.

###### `newLogicalId`<sup>Required</sup> <a name="newLogicalId" id="cdk-constructs-datadog-go.CfnMonitor.overrideLogicalId.parameter.newLogicalId"></a>

- *Type:* string

The new logical ID to use for this stack element.

---

##### `addDeletionOverride` <a name="addDeletionOverride" id="cdk-constructs-datadog-go.CfnMonitor.addDeletionOverride"></a>

```typescript
public addDeletionOverride(path: string): void
```

Syntactic sugar for `addOverride(path, undefined)`.

###### `path`<sup>Required</sup> <a name="path" id="cdk-constructs-datadog-go.CfnMonitor.addDeletionOverride.parameter.path"></a>

- *Type:* string

The path of the value to delete.

---

##### `addDependency` <a name="addDependency" id="cdk-constructs-datadog-go.CfnMonitor.addDependency"></a>

```typescript
public addDependency(target: CfnResource): void
```

Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.

This can be used for resources across stacks (or nested stack) boundaries
and the dependency will automatically be transferred to the relevant scope.

###### `target`<sup>Required</sup> <a name="target" id="cdk-constructs-datadog-go.CfnMonitor.addDependency.parameter.target"></a>

- *Type:* aws-cdk-lib.CfnResource

---

##### ~~`addDependsOn`~~ <a name="addDependsOn" id="cdk-constructs-datadog-go.CfnMonitor.addDependsOn"></a>

```typescript
public addDependsOn(target: CfnResource): void
```

Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.

###### `target`<sup>Required</sup> <a name="target" id="cdk-constructs-datadog-go.CfnMonitor.addDependsOn.parameter.target"></a>

- *Type:* aws-cdk-lib.CfnResource

---

##### `addMetadata` <a name="addMetadata" id="cdk-constructs-datadog-go.CfnMonitor.addMetadata"></a>

```typescript
public addMetadata(key: string, value: any): void
```

Add a value to the CloudFormation Resource Metadata.

> [https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html

Note that this is a different set of metadata from CDK node metadata; this
metadata ends up in the stack template under the resource, whereas CDK
node metadata ends up in the Cloud Assembly.](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html

Note that this is a different set of metadata from CDK node metadata; this
metadata ends up in the stack template under the resource, whereas CDK
node metadata ends up in the Cloud Assembly.)

###### `key`<sup>Required</sup> <a name="key" id="cdk-constructs-datadog-go.CfnMonitor.addMetadata.parameter.key"></a>

- *Type:* string

---

###### `value`<sup>Required</sup> <a name="value" id="cdk-constructs-datadog-go.CfnMonitor.addMetadata.parameter.value"></a>

- *Type:* any

---

##### `addOverride` <a name="addOverride" id="cdk-constructs-datadog-go.CfnMonitor.addOverride"></a>

```typescript
public addOverride(path: string, value: any): void
```

Adds an override to the synthesized CloudFormation resource.

To add a
property override, either use `addPropertyOverride` or prefix `path` with
"Properties." (i.e. `Properties.TopicName`).

If the override is nested, separate each nested level using a dot (.) in the path parameter.
If there is an array as part of the nesting, specify the index in the path.

To include a literal `.` in the property name, prefix with a `\`. In most
programming languages you will need to write this as `"\\."` because the
`\` itself will need to be escaped.

For example,
```typescript
cfnResource.addOverride('Properties.GlobalSecondaryIndexes.0.Projection.NonKeyAttributes', ['myattribute']);
cfnResource.addOverride('Properties.GlobalSecondaryIndexes.1.ProjectionType', 'INCLUDE');
```
would add the overrides
```json
"Properties": {
  "GlobalSecondaryIndexes": [
    {
      "Projection": {
        "NonKeyAttributes": [ "myattribute" ]
        ...
      }
      ...
    },
    {
      "ProjectionType": "INCLUDE"
      ...
    },
  ]
  ...
}
```

The `value` argument to `addOverride` will not be processed or translated
in any way. Pass raw JSON values in here with the correct capitalization
for CloudFormation. If you pass CDK classes or structs, they will be
rendered with lowercased key names, and CloudFormation will reject the
template.

###### `path`<sup>Required</sup> <a name="path" id="cdk-constructs-datadog-go.CfnMonitor.addOverride.parameter.path"></a>

- *Type:* string

The path of the property, you can use dot notation to override values in complex types.

Any intermediate keys
will be created as needed.

---

###### `value`<sup>Required</sup> <a name="value" id="cdk-constructs-datadog-go.CfnMonitor.addOverride.parameter.value"></a>

- *Type:* any

The value.

Could be primitive or complex.

---

##### `addPropertyDeletionOverride` <a name="addPropertyDeletionOverride" id="cdk-constructs-datadog-go.CfnMonitor.addPropertyDeletionOverride"></a>

```typescript
public addPropertyDeletionOverride(propertyPath: string): void
```

Adds an override that deletes the value of a property from the resource definition.

###### `propertyPath`<sup>Required</sup> <a name="propertyPath" id="cdk-constructs-datadog-go.CfnMonitor.addPropertyDeletionOverride.parameter.propertyPath"></a>

- *Type:* string

The path to the property.

---

##### `addPropertyOverride` <a name="addPropertyOverride" id="cdk-constructs-datadog-go.CfnMonitor.addPropertyOverride"></a>

```typescript
public addPropertyOverride(propertyPath: string, value: any): void
```

Adds an override to a resource property.

Syntactic sugar for `addOverride("Properties.<...>", value)`.

###### `propertyPath`<sup>Required</sup> <a name="propertyPath" id="cdk-constructs-datadog-go.CfnMonitor.addPropertyOverride.parameter.propertyPath"></a>

- *Type:* string

The path of the property.

---

###### `value`<sup>Required</sup> <a name="value" id="cdk-constructs-datadog-go.CfnMonitor.addPropertyOverride.parameter.value"></a>

- *Type:* any

The value.

---

##### `applyRemovalPolicy` <a name="applyRemovalPolicy" id="cdk-constructs-datadog-go.CfnMonitor.applyRemovalPolicy"></a>

```typescript
public applyRemovalPolicy(policy?: RemovalPolicy, options?: RemovalPolicyOptions): void
```

Sets the deletion policy of the resource based on the removal policy specified.

The Removal Policy controls what happens to this resource when it stops
being managed by CloudFormation, either because you've removed it from the
CDK application or because you've made a change that requires the resource
to be replaced.

The resource can be deleted (`RemovalPolicy.DESTROY`), or left in your AWS
account for data recovery and cleanup later (`RemovalPolicy.RETAIN`). In some
cases, a snapshot can be taken of the resource prior to deletion
(`RemovalPolicy.SNAPSHOT`). A list of resources that support this policy
can be found in the following link:

> [https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-deletionpolicy.html#aws-attribute-deletionpolicy-options](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-deletionpolicy.html#aws-attribute-deletionpolicy-options)

###### `policy`<sup>Optional</sup> <a name="policy" id="cdk-constructs-datadog-go.CfnMonitor.applyRemovalPolicy.parameter.policy"></a>

- *Type:* aws-cdk-lib.RemovalPolicy

---

###### `options`<sup>Optional</sup> <a name="options" id="cdk-constructs-datadog-go.CfnMonitor.applyRemovalPolicy.parameter.options"></a>

- *Type:* aws-cdk-lib.RemovalPolicyOptions

---

##### `getAtt` <a name="getAtt" id="cdk-constructs-datadog-go.CfnMonitor.getAtt"></a>

```typescript
public getAtt(attributeName: string, typeHint?: ResolutionTypeHint): Reference
```

Returns a token for an runtime attribute of this resource.

Ideally, use generated attribute accessors (e.g. `resource.arn`), but this can be used for future compatibility
in case there is no generated attribute.

###### `attributeName`<sup>Required</sup> <a name="attributeName" id="cdk-constructs-datadog-go.CfnMonitor.getAtt.parameter.attributeName"></a>

- *Type:* string

The name of the attribute.

---

###### `typeHint`<sup>Optional</sup> <a name="typeHint" id="cdk-constructs-datadog-go.CfnMonitor.getAtt.parameter.typeHint"></a>

- *Type:* aws-cdk-lib.ResolutionTypeHint

---

##### `getMetadata` <a name="getMetadata" id="cdk-constructs-datadog-go.CfnMonitor.getMetadata"></a>

```typescript
public getMetadata(key: string): any
```

Retrieve a value value from the CloudFormation Resource Metadata.

> [https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html

Note that this is a different set of metadata from CDK node metadata; this
metadata ends up in the stack template under the resource, whereas CDK
node metadata ends up in the Cloud Assembly.](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html

Note that this is a different set of metadata from CDK node metadata; this
metadata ends up in the stack template under the resource, whereas CDK
node metadata ends up in the Cloud Assembly.)

###### `key`<sup>Required</sup> <a name="key" id="cdk-constructs-datadog-go.CfnMonitor.getMetadata.parameter.key"></a>

- *Type:* string

---

##### `obtainDependencies` <a name="obtainDependencies" id="cdk-constructs-datadog-go.CfnMonitor.obtainDependencies"></a>

```typescript
public obtainDependencies(): Stack | CfnResource[]
```

Retrieves an array of resources this resource depends on.

This assembles dependencies on resources across stacks (including nested stacks)
automatically.

##### `obtainResourceDependencies` <a name="obtainResourceDependencies" id="cdk-constructs-datadog-go.CfnMonitor.obtainResourceDependencies"></a>

```typescript
public obtainResourceDependencies(): CfnResource[]
```

Get a shallow copy of dependencies between this resource and other resources in the same stack.

##### `removeDependency` <a name="removeDependency" id="cdk-constructs-datadog-go.CfnMonitor.removeDependency"></a>

```typescript
public removeDependency(target: CfnResource): void
```

Indicates that this resource no longer depends on another resource.

This can be used for resources across stacks (including nested stacks)
and the dependency will automatically be removed from the relevant scope.

###### `target`<sup>Required</sup> <a name="target" id="cdk-constructs-datadog-go.CfnMonitor.removeDependency.parameter.target"></a>

- *Type:* aws-cdk-lib.CfnResource

---

##### `replaceDependency` <a name="replaceDependency" id="cdk-constructs-datadog-go.CfnMonitor.replaceDependency"></a>

```typescript
public replaceDependency(target: CfnResource, newTarget: CfnResource): void
```

Replaces one dependency with another.

###### `target`<sup>Required</sup> <a name="target" id="cdk-constructs-datadog-go.CfnMonitor.replaceDependency.parameter.target"></a>

- *Type:* aws-cdk-lib.CfnResource

The dependency to replace.

---

###### `newTarget`<sup>Required</sup> <a name="newTarget" id="cdk-constructs-datadog-go.CfnMonitor.replaceDependency.parameter.newTarget"></a>

- *Type:* aws-cdk-lib.CfnResource

The new dependency to add.

---

#### Static Functions <a name="Static Functions" id="Static Functions"></a>

| **Name** | **Description** |
| --- | --- |
| <code><a href="#cdk-constructs-datadog-go.CfnMonitor.isConstruct">isConstruct</a></code> | Checks if `x` is a construct. |
| <code><a href="#cdk-constructs-datadog-go.CfnMonitor.isCfnElement">isCfnElement</a></code> | Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template). |
| <code><a href="#cdk-constructs-datadog-go.CfnMonitor.isCfnResource">isCfnResource</a></code> | Check whether the given construct is a CfnResource. |

---

##### ~~`isConstruct`~~ <a name="isConstruct" id="cdk-constructs-datadog-go.CfnMonitor.isConstruct"></a>

```typescript
import { CfnMonitor } from 'cdk-constructs-datadog-go'

CfnMonitor.isConstruct(x: any)
```

Checks if `x` is a construct.

###### `x`<sup>Required</sup> <a name="x" id="cdk-constructs-datadog-go.CfnMonitor.isConstruct.parameter.x"></a>

- *Type:* any

Any object.

---

##### `isCfnElement` <a name="isCfnElement" id="cdk-constructs-datadog-go.CfnMonitor.isCfnElement"></a>

```typescript
import { CfnMonitor } from 'cdk-constructs-datadog-go'

CfnMonitor.isCfnElement(x: any)
```

Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).

Uses duck-typing instead of `instanceof` to allow stack elements from different
versions of this library to be included in the same stack.

###### `x`<sup>Required</sup> <a name="x" id="cdk-constructs-datadog-go.CfnMonitor.isCfnElement.parameter.x"></a>

- *Type:* any

---

##### `isCfnResource` <a name="isCfnResource" id="cdk-constructs-datadog-go.CfnMonitor.isCfnResource"></a>

```typescript
import { CfnMonitor } from 'cdk-constructs-datadog-go'

CfnMonitor.isCfnResource(construct: IConstruct)
```

Check whether the given construct is a CfnResource.

###### `construct`<sup>Required</sup> <a name="construct" id="cdk-constructs-datadog-go.CfnMonitor.isCfnResource.parameter.construct"></a>

- *Type:* constructs.IConstruct

---

#### Properties <a name="Properties" id="Properties"></a>

| **Name** | **Type** | **Description** |
| --- | --- | --- |
| <code><a href="#cdk-constructs-datadog-go.CfnMonitor.property.node">node</a></code> | <code>constructs.Node</code> | The tree node. |
| <code><a href="#cdk-constructs-datadog-go.CfnMonitor.property.creationStack">creationStack</a></code> | <code>string[]</code> | *No description.* |
| <code><a href="#cdk-constructs-datadog-go.CfnMonitor.property.logicalId">logicalId</a></code> | <code>string</code> | The logical ID for this CloudFormation stack element. |
| <code><a href="#cdk-constructs-datadog-go.CfnMonitor.property.stack">stack</a></code> | <code>aws-cdk-lib.Stack</code> | The stack in which this element is defined. |
| <code><a href="#cdk-constructs-datadog-go.CfnMonitor.property.ref">ref</a></code> | <code>string</code> | Return a string that will be resolved to a CloudFormation `{ Ref }` for this element. |
| <code><a href="#cdk-constructs-datadog-go.CfnMonitor.property.cfnOptions">cfnOptions</a></code> | <code>aws-cdk-lib.ICfnResourceOptions</code> | Options for this resource, such as condition, update policy etc. |
| <code><a href="#cdk-constructs-datadog-go.CfnMonitor.property.cfnResourceType">cfnResourceType</a></code> | <code>string</code> | AWS resource type. |
| <code><a href="#cdk-constructs-datadog-go.CfnMonitor.property.attrCreated">attrCreated</a></code> | <code>string</code> | Attribute `Datadog::Monitors::Monitor.Created`. |
| <code><a href="#cdk-constructs-datadog-go.CfnMonitor.property.attrDeleted">attrDeleted</a></code> | <code>string</code> | Attribute `Datadog::Monitors::Monitor.Deleted`. |
| <code><a href="#cdk-constructs-datadog-go.CfnMonitor.property.attrId">attrId</a></code> | <code>number</code> | Attribute `Datadog::Monitors::Monitor.Id`. |
| <code><a href="#cdk-constructs-datadog-go.CfnMonitor.property.attrModified">attrModified</a></code> | <code>string</code> | Attribute `Datadog::Monitors::Monitor.Modified`. |
| <code><a href="#cdk-constructs-datadog-go.CfnMonitor.property.props">props</a></code> | <code><a href="#cdk-constructs-datadog-go.CfnMonitorProps">CfnMonitorProps</a></code> | Resource props. |

---

##### `node`<sup>Required</sup> <a name="node" id="cdk-constructs-datadog-go.CfnMonitor.property.node"></a>

```typescript
public readonly node: Node;
```

- *Type:* constructs.Node

The tree node.

---

##### `creationStack`<sup>Required</sup> <a name="creationStack" id="cdk-constructs-datadog-go.CfnMonitor.property.creationStack"></a>

```typescript
public readonly creationStack: string[];
```

- *Type:* string[]

---

##### `logicalId`<sup>Required</sup> <a name="logicalId" id="cdk-constructs-datadog-go.CfnMonitor.property.logicalId"></a>

```typescript
public readonly logicalId: string;
```

- *Type:* string

The logical ID for this CloudFormation stack element.

The logical ID of the element
is calculated from the path of the resource node in the construct tree.

To override this value, use `overrideLogicalId(newLogicalId)`.

---

##### `stack`<sup>Required</sup> <a name="stack" id="cdk-constructs-datadog-go.CfnMonitor.property.stack"></a>

```typescript
public readonly stack: Stack;
```

- *Type:* aws-cdk-lib.Stack

The stack in which this element is defined.

CfnElements must be defined within a stack scope (directly or indirectly).

---

##### `ref`<sup>Required</sup> <a name="ref" id="cdk-constructs-datadog-go.CfnMonitor.property.ref"></a>

```typescript
public readonly ref: string;
```

- *Type:* string

Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.

If, by any chance, the intrinsic reference of a resource is not a string, you could
coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.

---

##### `cfnOptions`<sup>Required</sup> <a name="cfnOptions" id="cdk-constructs-datadog-go.CfnMonitor.property.cfnOptions"></a>

```typescript
public readonly cfnOptions: ICfnResourceOptions;
```

- *Type:* aws-cdk-lib.ICfnResourceOptions

Options for this resource, such as condition, update policy etc.

---

##### `cfnResourceType`<sup>Required</sup> <a name="cfnResourceType" id="cdk-constructs-datadog-go.CfnMonitor.property.cfnResourceType"></a>

```typescript
public readonly cfnResourceType: string;
```

- *Type:* string

AWS resource type.

---

##### `attrCreated`<sup>Required</sup> <a name="attrCreated" id="cdk-constructs-datadog-go.CfnMonitor.property.attrCreated"></a>

```typescript
public readonly attrCreated: string;
```

- *Type:* string

Attribute `Datadog::Monitors::Monitor.Created`.

> [http://unknown-url](http://unknown-url)

---

##### `attrDeleted`<sup>Required</sup> <a name="attrDeleted" id="cdk-constructs-datadog-go.CfnMonitor.property.attrDeleted"></a>

```typescript
public readonly attrDeleted: string;
```

- *Type:* string

Attribute `Datadog::Monitors::Monitor.Deleted`.

> [http://unknown-url](http://unknown-url)

---

##### `attrId`<sup>Required</sup> <a name="attrId" id="cdk-constructs-datadog-go.CfnMonitor.property.attrId"></a>

```typescript
public readonly attrId: number;
```

- *Type:* number

Attribute `Datadog::Monitors::Monitor.Id`.

> [http://unknown-url](http://unknown-url)

---

##### `attrModified`<sup>Required</sup> <a name="attrModified" id="cdk-constructs-datadog-go.CfnMonitor.property.attrModified"></a>

```typescript
public readonly attrModified: string;
```

- *Type:* string

Attribute `Datadog::Monitors::Monitor.Modified`.

> [http://unknown-url](http://unknown-url)

---

##### `props`<sup>Required</sup> <a name="props" id="cdk-constructs-datadog-go.CfnMonitor.property.props"></a>

```typescript
public readonly props: CfnMonitorProps;
```

- *Type:* <a href="#cdk-constructs-datadog-go.CfnMonitorProps">CfnMonitorProps</a>

Resource props.

---

#### Constants <a name="Constants" id="Constants"></a>

| **Name** | **Type** | **Description** |
| --- | --- | --- |
| <code><a href="#cdk-constructs-datadog-go.CfnMonitor.property.CFN_RESOURCE_TYPE_NAME">CFN_RESOURCE_TYPE_NAME</a></code> | <code>string</code> | The CloudFormation resource type name for this resource class. |

---

##### `CFN_RESOURCE_TYPE_NAME`<sup>Required</sup> <a name="CFN_RESOURCE_TYPE_NAME" id="cdk-constructs-datadog-go.CfnMonitor.property.CFN_RESOURCE_TYPE_NAME"></a>

```typescript
public readonly CFN_RESOURCE_TYPE_NAME: string;
```

- *Type:* string

The CloudFormation resource type name for this resource class.

---

## Structs <a name="Structs" id="Structs"></a>

### CfnDashboardProps <a name="CfnDashboardProps" id="cdk-constructs-datadog-go.CfnDashboardProps"></a>

Datadog Dashboard 2.1.0.

#### Initializer <a name="Initializer" id="cdk-constructs-datadog-go.CfnDashboardProps.Initializer"></a>

```typescript
import { CfnDashboardProps } from 'cdk-constructs-datadog-go'

const cfnDashboardProps: CfnDashboardProps = { ... }
```

#### Properties <a name="Properties" id="Properties"></a>

| **Name** | **Type** | **Description** |
| --- | --- | --- |
| <code><a href="#cdk-constructs-datadog-go.CfnDashboardProps.property.dashboardDefinition">dashboardDefinition</a></code> | <code>string \| aws-cdk-lib.IResolvable</code> | JSON string of the dashboard definition. |

---

##### `dashboardDefinition`<sup>Required</sup> <a name="dashboardDefinition" id="cdk-constructs-datadog-go.CfnDashboardProps.property.dashboardDefinition"></a>

```typescript
public readonly dashboardDefinition: string | IResolvable;
```

- *Type:* string | aws-cdk-lib.IResolvable

JSON string of the dashboard definition.

---

### CfnMonitorProps <a name="CfnMonitorProps" id="cdk-constructs-datadog-go.CfnMonitorProps"></a>

Datadog Monitor 4.7.0.

#### Initializer <a name="Initializer" id="cdk-constructs-datadog-go.CfnMonitorProps.Initializer"></a>

```typescript
import { CfnMonitorProps } from 'cdk-constructs-datadog-go'

const cfnMonitorProps: CfnMonitorProps = { ... }
```

#### Properties <a name="Properties" id="Properties"></a>

| **Name** | **Type** | **Description** |
| --- | --- | --- |
| <code><a href="#cdk-constructs-datadog-go.CfnMonitorProps.property.query">query</a></code> | <code>string</code> | The monitor query. |
| <code><a href="#cdk-constructs-datadog-go.CfnMonitorProps.property.type">type</a></code> | <code><a href="#cdk-constructs-datadog-go.CfnMonitorPropsType">CfnMonitorPropsType</a></code> | The type of the monitor. |
| <code><a href="#cdk-constructs-datadog-go.CfnMonitorProps.property.creator">creator</a></code> | <code><a href="#cdk-constructs-datadog-go.Creator">Creator</a></code> | *No description.* |
| <code><a href="#cdk-constructs-datadog-go.CfnMonitorProps.property.message">message</a></code> | <code>string</code> | A message to include with notifications for the monitor. |
| <code><a href="#cdk-constructs-datadog-go.CfnMonitorProps.property.multi">multi</a></code> | <code>boolean</code> | Whether or not the monitor is multi alert. |
| <code><a href="#cdk-constructs-datadog-go.CfnMonitorProps.property.name">name</a></code> | <code>string</code> | Name of the monitor. |
| <code><a href="#cdk-constructs-datadog-go.CfnMonitorProps.property.options">options</a></code> | <code><a href="#cdk-constructs-datadog-go.MonitorOptions">MonitorOptions</a></code> | The monitor options. |
| <code><a href="#cdk-constructs-datadog-go.CfnMonitorProps.property.priority">priority</a></code> | <code>number</code> | Integer from 1 (high) to 5 (low) indicating alert severity. |
| <code><a href="#cdk-constructs-datadog-go.CfnMonitorProps.property.restrictedRoles">restrictedRoles</a></code> | <code>string[]</code> | A list of unique role identifiers to define which roles are allowed to edit the monitor. |
| <code><a href="#cdk-constructs-datadog-go.CfnMonitorProps.property.tags">tags</a></code> | <code>string[]</code> | Tags associated with the monitor. |

---

##### `query`<sup>Required</sup> <a name="query" id="cdk-constructs-datadog-go.CfnMonitorProps.property.query"></a>

```typescript
public readonly query: string;
```

- *Type:* string

The monitor query.

---

##### `type`<sup>Required</sup> <a name="type" id="cdk-constructs-datadog-go.CfnMonitorProps.property.type"></a>

```typescript
public readonly type: CfnMonitorPropsType;
```

- *Type:* <a href="#cdk-constructs-datadog-go.CfnMonitorPropsType">CfnMonitorPropsType</a>

The type of the monitor.

---

##### `creator`<sup>Optional</sup> <a name="creator" id="cdk-constructs-datadog-go.CfnMonitorProps.property.creator"></a>

```typescript
public readonly creator: Creator;
```

- *Type:* <a href="#cdk-constructs-datadog-go.Creator">Creator</a>

---

##### `message`<sup>Optional</sup> <a name="message" id="cdk-constructs-datadog-go.CfnMonitorProps.property.message"></a>

```typescript
public readonly message: string;
```

- *Type:* string

A message to include with notifications for the monitor.

---

##### `multi`<sup>Optional</sup> <a name="multi" id="cdk-constructs-datadog-go.CfnMonitorProps.property.multi"></a>

```typescript
public readonly multi: boolean;
```

- *Type:* boolean

Whether or not the monitor is multi alert.

---

##### `name`<sup>Optional</sup> <a name="name" id="cdk-constructs-datadog-go.CfnMonitorProps.property.name"></a>

```typescript
public readonly name: string;
```

- *Type:* string

Name of the monitor.

---

##### `options`<sup>Optional</sup> <a name="options" id="cdk-constructs-datadog-go.CfnMonitorProps.property.options"></a>

```typescript
public readonly options: MonitorOptions;
```

- *Type:* <a href="#cdk-constructs-datadog-go.MonitorOptions">MonitorOptions</a>

The monitor options.

---

##### `priority`<sup>Optional</sup> <a name="priority" id="cdk-constructs-datadog-go.CfnMonitorProps.property.priority"></a>

```typescript
public readonly priority: number;
```

- *Type:* number

Integer from 1 (high) to 5 (low) indicating alert severity.

---

##### `restrictedRoles`<sup>Optional</sup> <a name="restrictedRoles" id="cdk-constructs-datadog-go.CfnMonitorProps.property.restrictedRoles"></a>

```typescript
public readonly restrictedRoles: string[];
```

- *Type:* string[]

A list of unique role identifiers to define which roles are allowed to edit the monitor.

The unique identifiers for all roles can be pulled from the [Roles API](https://docs.datadoghq.com/api/latest/roles/#list-roles) and are located in the `data.id` field. Editing a monitor includes any updates to the monitor configuration, monitor deletion, and muting of the monitor for any amount of time. `restricted_roles` is the successor of `locked`. For more information about `locked` and `restricted_roles`, see the [monitor options docs](https://docs.datadoghq.com/monitors/guide/monitor_api_options/#permissions-options).

---

##### `tags`<sup>Optional</sup> <a name="tags" id="cdk-constructs-datadog-go.CfnMonitorProps.property.tags"></a>

```typescript
public readonly tags: string[];
```

- *Type:* string[]

Tags associated with the monitor.

---

### Creator <a name="Creator" id="cdk-constructs-datadog-go.Creator"></a>

#### Initializer <a name="Initializer" id="cdk-constructs-datadog-go.Creator.Initializer"></a>

```typescript
import { Creator } from 'cdk-constructs-datadog-go'

const creator: Creator = { ... }
```

#### Properties <a name="Properties" id="Properties"></a>

| **Name** | **Type** | **Description** |
| --- | --- | --- |
| <code><a href="#cdk-constructs-datadog-go.Creator.property.email">email</a></code> | <code>string</code> | Email of the creator of the monitor. |
| <code><a href="#cdk-constructs-datadog-go.Creator.property.handle">handle</a></code> | <code>string</code> | Handle of the creator of the monitor. |
| <code><a href="#cdk-constructs-datadog-go.Creator.property.name">name</a></code> | <code>string</code> | Name of the creator of the monitor. |

---

##### `email`<sup>Optional</sup> <a name="email" id="cdk-constructs-datadog-go.Creator.property.email"></a>

```typescript
public readonly email: string;
```

- *Type:* string

Email of the creator of the monitor.

---

##### `handle`<sup>Optional</sup> <a name="handle" id="cdk-constructs-datadog-go.Creator.property.handle"></a>

```typescript
public readonly handle: string;
```

- *Type:* string

Handle of the creator of the monitor.

---

##### `name`<sup>Optional</sup> <a name="name" id="cdk-constructs-datadog-go.Creator.property.name"></a>

```typescript
public readonly name: string;
```

- *Type:* string

Name of the creator of the monitor.

---

### MonitorOptions <a name="MonitorOptions" id="cdk-constructs-datadog-go.MonitorOptions"></a>

#### Initializer <a name="Initializer" id="cdk-constructs-datadog-go.MonitorOptions.Initializer"></a>

```typescript
import { MonitorOptions } from 'cdk-constructs-datadog-go'

const monitorOptions: MonitorOptions = { ... }
```

#### Properties <a name="Properties" id="Properties"></a>

| **Name** | **Type** | **Description** |
| --- | --- | --- |
| <code><a href="#cdk-constructs-datadog-go.MonitorOptions.property.enableLogsSample">enableLogsSample</a></code> | <code>boolean</code> | Whether or not to include a sample of the logs. |
| <code><a href="#cdk-constructs-datadog-go.MonitorOptions.property.enableSamples">enableSamples</a></code> | <code>boolean</code> | Whether or not to send a list of samples when the monitor triggers. |
| <code><a href="#cdk-constructs-datadog-go.MonitorOptions.property.escalationMessage">escalationMessage</a></code> | <code>string</code> | Message to include with a re-notification when renotify_interval is set. |
| <code><a href="#cdk-constructs-datadog-go.MonitorOptions.property.evaluationDelay">evaluationDelay</a></code> | <code>number</code> | Time in seconds to delay evaluation. |
| <code><a href="#cdk-constructs-datadog-go.MonitorOptions.property.groupRetentionDuration">groupRetentionDuration</a></code> | <code>string</code> | The time span after which groups with missing data are dropped from the monitor state. |
| <code><a href="#cdk-constructs-datadog-go.MonitorOptions.property.includeTags">includeTags</a></code> | <code>boolean</code> | Whether or not to include triggering tags into notification title. |
| <code><a href="#cdk-constructs-datadog-go.MonitorOptions.property.locked">locked</a></code> | <code>boolean</code> | Whether or not changes to this monitor should be restricted to the creator or admins. |
| <code><a href="#cdk-constructs-datadog-go.MonitorOptions.property.minFailureDuration">minFailureDuration</a></code> | <code>number</code> | How long the test should be in failure before alerting (integer, number of seconds, max 7200). |
| <code><a href="#cdk-constructs-datadog-go.MonitorOptions.property.minLocationFailed">minLocationFailed</a></code> | <code>number</code> | Number of locations allowed to fail before triggering alert. |
| <code><a href="#cdk-constructs-datadog-go.MonitorOptions.property.newGroupDelay">newGroupDelay</a></code> | <code>number</code> | Time (in seconds) to skip evaluations for new groups. |
| <code><a href="#cdk-constructs-datadog-go.MonitorOptions.property.newHostDelay">newHostDelay</a></code> | <code>number</code> | Time in seconds to allow a host to start reporting data before starting the evaluation of monitor results. |
| <code><a href="#cdk-constructs-datadog-go.MonitorOptions.property.noDataTimeframe">noDataTimeframe</a></code> | <code>number</code> | Number of minutes data stopped reporting before notifying. |
| <code><a href="#cdk-constructs-datadog-go.MonitorOptions.property.notificationPresetName">notificationPresetName</a></code> | <code><a href="#cdk-constructs-datadog-go.MonitorNotificationPresetName">MonitorNotificationPresetName</a></code> | *No description.* |
| <code><a href="#cdk-constructs-datadog-go.MonitorOptions.property.notifyAudit">notifyAudit</a></code> | <code>boolean</code> | Whether or not to notify tagged users when changes are made to the monitor. |
| <code><a href="#cdk-constructs-datadog-go.MonitorOptions.property.notifyBy">notifyBy</a></code> | <code>string[]</code> | Controls what granularity a monitor alerts on. |
| <code><a href="#cdk-constructs-datadog-go.MonitorOptions.property.notifyNoData">notifyNoData</a></code> | <code>boolean</code> | Whether or not to notify when data stops reporting. |
| <code><a href="#cdk-constructs-datadog-go.MonitorOptions.property.onMissingData">onMissingData</a></code> | <code><a href="#cdk-constructs-datadog-go.MonitorOnMissingData">MonitorOnMissingData</a></code> | *No description.* |
| <code><a href="#cdk-constructs-datadog-go.MonitorOptions.property.renotifyInterval">renotifyInterval</a></code> | <code>number</code> | Number of minutes after the last notification before the monitor re-notifies on the current status. |
| <code><a href="#cdk-constructs-datadog-go.MonitorOptions.property.renotifyOccurrences">renotifyOccurrences</a></code> | <code>number</code> | The number of times re-notification messages should be sent on the current status at the provided re-notification interval. |
| <code><a href="#cdk-constructs-datadog-go.MonitorOptions.property.renotifyStatuses">renotifyStatuses</a></code> | <code><a href="#cdk-constructs-datadog-go.MonitorOptionsRenotifyStatuses">MonitorOptionsRenotifyStatuses</a>[]</code> | The types of monitor statuses for which re-notification messages are sent. |
| <code><a href="#cdk-constructs-datadog-go.MonitorOptions.property.requireFullWindow">requireFullWindow</a></code> | <code>boolean</code> | Whether or not the monitor requires a full window of data before it is evaluated. |
| <code><a href="#cdk-constructs-datadog-go.MonitorOptions.property.schedulingOptions">schedulingOptions</a></code> | <code><a href="#cdk-constructs-datadog-go.MonitorSchedulingOptions">MonitorSchedulingOptions</a></code> | *No description.* |
| <code><a href="#cdk-constructs-datadog-go.MonitorOptions.property.syntheticsCheckId">syntheticsCheckId</a></code> | <code>number</code> | ID of the corresponding synthetics check. |
| <code><a href="#cdk-constructs-datadog-go.MonitorOptions.property.thresholds">thresholds</a></code> | <code><a href="#cdk-constructs-datadog-go.MonitorThresholds">MonitorThresholds</a></code> | The threshold definitions. |
| <code><a href="#cdk-constructs-datadog-go.MonitorOptions.property.thresholdWindows">thresholdWindows</a></code> | <code><a href="#cdk-constructs-datadog-go.MonitorThresholdWindows">MonitorThresholdWindows</a></code> | The threshold window definitions. |
| <code><a href="#cdk-constructs-datadog-go.MonitorOptions.property.timeoutH">timeoutH</a></code> | <code>number</code> | Number of hours of the monitor not reporting data before it automatically resolves. |
| <code><a href="#cdk-constructs-datadog-go.MonitorOptions.property.variables">variables</a></code> | <code>any[]</code> | List of requests that can be used in the monitor query. |

---

##### `enableLogsSample`<sup>Optional</sup> <a name="enableLogsSample" id="cdk-constructs-datadog-go.MonitorOptions.property.enableLogsSample"></a>

```typescript
public readonly enableLogsSample: boolean;
```

- *Type:* boolean

Whether or not to include a sample of the logs.

---

##### `enableSamples`<sup>Optional</sup> <a name="enableSamples" id="cdk-constructs-datadog-go.MonitorOptions.property.enableSamples"></a>

```typescript
public readonly enableSamples: boolean;
```

- *Type:* boolean

Whether or not to send a list of samples when the monitor triggers.

This is only used by CI Test and Pipeline monitors.

---

##### `escalationMessage`<sup>Optional</sup> <a name="escalationMessage" id="cdk-constructs-datadog-go.MonitorOptions.property.escalationMessage"></a>

```typescript
public readonly escalationMessage: string;
```

- *Type:* string

Message to include with a re-notification when renotify_interval is set.

---

##### `evaluationDelay`<sup>Optional</sup> <a name="evaluationDelay" id="cdk-constructs-datadog-go.MonitorOptions.property.evaluationDelay"></a>

```typescript
public readonly evaluationDelay: number;
```

- *Type:* number

Time in seconds to delay evaluation.

---

##### `groupRetentionDuration`<sup>Optional</sup> <a name="groupRetentionDuration" id="cdk-constructs-datadog-go.MonitorOptions.property.groupRetentionDuration"></a>

```typescript
public readonly groupRetentionDuration: string;
```

- *Type:* string

The time span after which groups with missing data are dropped from the monitor state.

The minimum value is one hour, and the maximum value is 72 hours.
Example values are: "60m", "1h", and "2d".
This option is only available for APM Trace Analytics, Audit Trail, CI, Error Tracking, Event, Logs, and RUM monitors.

---

##### `includeTags`<sup>Optional</sup> <a name="includeTags" id="cdk-constructs-datadog-go.MonitorOptions.property.includeTags"></a>

```typescript
public readonly includeTags: boolean;
```

- *Type:* boolean

Whether or not to include triggering tags into notification title.

---

##### `locked`<sup>Optional</sup> <a name="locked" id="cdk-constructs-datadog-go.MonitorOptions.property.locked"></a>

```typescript
public readonly locked: boolean;
```

- *Type:* boolean

Whether or not changes to this monitor should be restricted to the creator or admins.

---

##### `minFailureDuration`<sup>Optional</sup> <a name="minFailureDuration" id="cdk-constructs-datadog-go.MonitorOptions.property.minFailureDuration"></a>

```typescript
public readonly minFailureDuration: number;
```

- *Type:* number

How long the test should be in failure before alerting (integer, number of seconds, max 7200).

---

##### `minLocationFailed`<sup>Optional</sup> <a name="minLocationFailed" id="cdk-constructs-datadog-go.MonitorOptions.property.minLocationFailed"></a>

```typescript
public readonly minLocationFailed: number;
```

- *Type:* number

Number of locations allowed to fail before triggering alert.

---

##### `newGroupDelay`<sup>Optional</sup> <a name="newGroupDelay" id="cdk-constructs-datadog-go.MonitorOptions.property.newGroupDelay"></a>

```typescript
public readonly newGroupDelay: number;
```

- *Type:* number

Time (in seconds) to skip evaluations for new groups.

For example, this option can be used to skip evaluations for new hosts while they initialize. Must be a non negative integer.

---

##### `newHostDelay`<sup>Optional</sup> <a name="newHostDelay" id="cdk-constructs-datadog-go.MonitorOptions.property.newHostDelay"></a>

```typescript
public readonly newHostDelay: number;
```

- *Type:* number

Time in seconds to allow a host to start reporting data before starting the evaluation of monitor results.

---

##### `noDataTimeframe`<sup>Optional</sup> <a name="noDataTimeframe" id="cdk-constructs-datadog-go.MonitorOptions.property.noDataTimeframe"></a>

```typescript
public readonly noDataTimeframe: number;
```

- *Type:* number

Number of minutes data stopped reporting before notifying.

---

##### `notificationPresetName`<sup>Optional</sup> <a name="notificationPresetName" id="cdk-constructs-datadog-go.MonitorOptions.property.notificationPresetName"></a>

```typescript
public readonly notificationPresetName: MonitorNotificationPresetName;
```

- *Type:* <a href="#cdk-constructs-datadog-go.MonitorNotificationPresetName">MonitorNotificationPresetName</a>

---

##### `notifyAudit`<sup>Optional</sup> <a name="notifyAudit" id="cdk-constructs-datadog-go.MonitorOptions.property.notifyAudit"></a>

```typescript
public readonly notifyAudit: boolean;
```

- *Type:* boolean

Whether or not to notify tagged users when changes are made to the monitor.

---

##### `notifyBy`<sup>Optional</sup> <a name="notifyBy" id="cdk-constructs-datadog-go.MonitorOptions.property.notifyBy"></a>

```typescript
public readonly notifyBy: string[];
```

- *Type:* string[]

Controls what granularity a monitor alerts on.

Only available for monitors with groupings.
For instance, a monitor grouped by `cluster`, `namespace`, and `pod` can be configured to only notify on each new `cluster` violating the alert conditions by setting `notify_by` to `["cluster"]`.
Tags mentioned in `notify_by` must be a subset of the grouping tags in the query.
For example, a query grouped by `cluster` and `namespace` cannot notify on `region`.
Setting `notify_by` to `[*]` configures the monitor to notify as a simple-alert.

---

##### `notifyNoData`<sup>Optional</sup> <a name="notifyNoData" id="cdk-constructs-datadog-go.MonitorOptions.property.notifyNoData"></a>

```typescript
public readonly notifyNoData: boolean;
```

- *Type:* boolean

Whether or not to notify when data stops reporting.

---

##### `onMissingData`<sup>Optional</sup> <a name="onMissingData" id="cdk-constructs-datadog-go.MonitorOptions.property.onMissingData"></a>

```typescript
public readonly onMissingData: MonitorOnMissingData;
```

- *Type:* <a href="#cdk-constructs-datadog-go.MonitorOnMissingData">MonitorOnMissingData</a>

---

##### `renotifyInterval`<sup>Optional</sup> <a name="renotifyInterval" id="cdk-constructs-datadog-go.MonitorOptions.property.renotifyInterval"></a>

```typescript
public readonly renotifyInterval: number;
```

- *Type:* number

Number of minutes after the last notification before the monitor re-notifies on the current status.

---

##### `renotifyOccurrences`<sup>Optional</sup> <a name="renotifyOccurrences" id="cdk-constructs-datadog-go.MonitorOptions.property.renotifyOccurrences"></a>

```typescript
public readonly renotifyOccurrences: number;
```

- *Type:* number

The number of times re-notification messages should be sent on the current status at the provided re-notification interval.

---

##### `renotifyStatuses`<sup>Optional</sup> <a name="renotifyStatuses" id="cdk-constructs-datadog-go.MonitorOptions.property.renotifyStatuses"></a>

```typescript
public readonly renotifyStatuses: MonitorOptionsRenotifyStatuses[];
```

- *Type:* <a href="#cdk-constructs-datadog-go.MonitorOptionsRenotifyStatuses">MonitorOptionsRenotifyStatuses</a>[]

The types of monitor statuses for which re-notification messages are sent.

---

##### `requireFullWindow`<sup>Optional</sup> <a name="requireFullWindow" id="cdk-constructs-datadog-go.MonitorOptions.property.requireFullWindow"></a>

```typescript
public readonly requireFullWindow: boolean;
```

- *Type:* boolean

Whether or not the monitor requires a full window of data before it is evaluated.

---

##### `schedulingOptions`<sup>Optional</sup> <a name="schedulingOptions" id="cdk-constructs-datadog-go.MonitorOptions.property.schedulingOptions"></a>

```typescript
public readonly schedulingOptions: MonitorSchedulingOptions;
```

- *Type:* <a href="#cdk-constructs-datadog-go.MonitorSchedulingOptions">MonitorSchedulingOptions</a>

---

##### `syntheticsCheckId`<sup>Optional</sup> <a name="syntheticsCheckId" id="cdk-constructs-datadog-go.MonitorOptions.property.syntheticsCheckId"></a>

```typescript
public readonly syntheticsCheckId: number;
```

- *Type:* number

ID of the corresponding synthetics check.

---

##### `thresholds`<sup>Optional</sup> <a name="thresholds" id="cdk-constructs-datadog-go.MonitorOptions.property.thresholds"></a>

```typescript
public readonly thresholds: MonitorThresholds;
```

- *Type:* <a href="#cdk-constructs-datadog-go.MonitorThresholds">MonitorThresholds</a>

The threshold definitions.

---

##### `thresholdWindows`<sup>Optional</sup> <a name="thresholdWindows" id="cdk-constructs-datadog-go.MonitorOptions.property.thresholdWindows"></a>

```typescript
public readonly thresholdWindows: MonitorThresholdWindows;
```

- *Type:* <a href="#cdk-constructs-datadog-go.MonitorThresholdWindows">MonitorThresholdWindows</a>

The threshold window definitions.

---

##### `timeoutH`<sup>Optional</sup> <a name="timeoutH" id="cdk-constructs-datadog-go.MonitorOptions.property.timeoutH"></a>

```typescript
public readonly timeoutH: number;
```

- *Type:* number

Number of hours of the monitor not reporting data before it automatically resolves.

---

##### `variables`<sup>Optional</sup> <a name="variables" id="cdk-constructs-datadog-go.MonitorOptions.property.variables"></a>

```typescript
public readonly variables: any[];
```

- *Type:* any[]

List of requests that can be used in the monitor query.

---

### MonitorSchedulingOptions <a name="MonitorSchedulingOptions" id="cdk-constructs-datadog-go.MonitorSchedulingOptions"></a>

Configuration options for scheduling.

#### Initializer <a name="Initializer" id="cdk-constructs-datadog-go.MonitorSchedulingOptions.Initializer"></a>

```typescript
import { MonitorSchedulingOptions } from 'cdk-constructs-datadog-go'

const monitorSchedulingOptions: MonitorSchedulingOptions = { ... }
```

#### Properties <a name="Properties" id="Properties"></a>

| **Name** | **Type** | **Description** |
| --- | --- | --- |
| <code><a href="#cdk-constructs-datadog-go.MonitorSchedulingOptions.property.evaluationWindow">evaluationWindow</a></code> | <code><a href="#cdk-constructs-datadog-go.MonitorSchedulingOptionsEvaluationWindow">MonitorSchedulingOptionsEvaluationWindow</a></code> | *No description.* |

---

##### `evaluationWindow`<sup>Optional</sup> <a name="evaluationWindow" id="cdk-constructs-datadog-go.MonitorSchedulingOptions.property.evaluationWindow"></a>

```typescript
public readonly evaluationWindow: MonitorSchedulingOptionsEvaluationWindow;
```

- *Type:* <a href="#cdk-constructs-datadog-go.MonitorSchedulingOptionsEvaluationWindow">MonitorSchedulingOptionsEvaluationWindow</a>

---

### MonitorSchedulingOptionsEvaluationWindow <a name="MonitorSchedulingOptionsEvaluationWindow" id="cdk-constructs-datadog-go.MonitorSchedulingOptionsEvaluationWindow"></a>

Configuration options for the evaluation window.

If `hour_starts` is set, no other fields may be set. Otherwise, `day_starts` and `month_starts` must be set together.

#### Initializer <a name="Initializer" id="cdk-constructs-datadog-go.MonitorSchedulingOptionsEvaluationWindow.Initializer"></a>

```typescript
import { MonitorSchedulingOptionsEvaluationWindow } from 'cdk-constructs-datadog-go'

const monitorSchedulingOptionsEvaluationWindow: MonitorSchedulingOptionsEvaluationWindow = { ... }
```

#### Properties <a name="Properties" id="Properties"></a>

| **Name** | **Type** | **Description** |
| --- | --- | --- |
| <code><a href="#cdk-constructs-datadog-go.MonitorSchedulingOptionsEvaluationWindow.property.dayStarts">dayStarts</a></code> | <code>string</code> | The time of the day at which a one day cumulative evaluation window starts. |
| <code><a href="#cdk-constructs-datadog-go.MonitorSchedulingOptionsEvaluationWindow.property.hourStarts">hourStarts</a></code> | <code>number</code> | The minute of the hour at which a one hour cumulative evaluation window starts. |
| <code><a href="#cdk-constructs-datadog-go.MonitorSchedulingOptionsEvaluationWindow.property.monthStarts">monthStarts</a></code> | <code>number</code> | The day of the month at which a one month cumulative evaluation window starts. |

---

##### `dayStarts`<sup>Optional</sup> <a name="dayStarts" id="cdk-constructs-datadog-go.MonitorSchedulingOptionsEvaluationWindow.property.dayStarts"></a>

```typescript
public readonly dayStarts: string;
```

- *Type:* string

The time of the day at which a one day cumulative evaluation window starts.

Must be defined in UTC time in `HH:mm` format.

---

##### `hourStarts`<sup>Optional</sup> <a name="hourStarts" id="cdk-constructs-datadog-go.MonitorSchedulingOptionsEvaluationWindow.property.hourStarts"></a>

```typescript
public readonly hourStarts: number;
```

- *Type:* number

The minute of the hour at which a one hour cumulative evaluation window starts.

---

##### `monthStarts`<sup>Optional</sup> <a name="monthStarts" id="cdk-constructs-datadog-go.MonitorSchedulingOptionsEvaluationWindow.property.monthStarts"></a>

```typescript
public readonly monthStarts: number;
```

- *Type:* number

The day of the month at which a one month cumulative evaluation window starts.

---

### MonitorThresholds <a name="MonitorThresholds" id="cdk-constructs-datadog-go.MonitorThresholds"></a>

#### Initializer <a name="Initializer" id="cdk-constructs-datadog-go.MonitorThresholds.Initializer"></a>

```typescript
import { MonitorThresholds } from 'cdk-constructs-datadog-go'

const monitorThresholds: MonitorThresholds = { ... }
```

#### Properties <a name="Properties" id="Properties"></a>

| **Name** | **Type** | **Description** |
| --- | --- | --- |
| <code><a href="#cdk-constructs-datadog-go.MonitorThresholds.property.critical">critical</a></code> | <code>number</code> | Threshold value for triggering an alert. |
| <code><a href="#cdk-constructs-datadog-go.MonitorThresholds.property.criticalRecovery">criticalRecovery</a></code> | <code>number</code> | Threshold value for recovering from an alert state. |
| <code><a href="#cdk-constructs-datadog-go.MonitorThresholds.property.ok">ok</a></code> | <code>number</code> | Threshold value for recovering from an alert state. |
| <code><a href="#cdk-constructs-datadog-go.MonitorThresholds.property.warning">warning</a></code> | <code>number</code> | Threshold value for triggering a warning. |
| <code><a href="#cdk-constructs-datadog-go.MonitorThresholds.property.warningRecovery">warningRecovery</a></code> | <code>number</code> | Threshold value for recovering from a warning state. |

---

##### `critical`<sup>Optional</sup> <a name="critical" id="cdk-constructs-datadog-go.MonitorThresholds.property.critical"></a>

```typescript
public readonly critical: number;
```

- *Type:* number

Threshold value for triggering an alert.

---

##### `criticalRecovery`<sup>Optional</sup> <a name="criticalRecovery" id="cdk-constructs-datadog-go.MonitorThresholds.property.criticalRecovery"></a>

```typescript
public readonly criticalRecovery: number;
```

- *Type:* number

Threshold value for recovering from an alert state.

---

##### `ok`<sup>Optional</sup> <a name="ok" id="cdk-constructs-datadog-go.MonitorThresholds.property.ok"></a>

```typescript
public readonly ok: number;
```

- *Type:* number

Threshold value for recovering from an alert state.

---

##### `warning`<sup>Optional</sup> <a name="warning" id="cdk-constructs-datadog-go.MonitorThresholds.property.warning"></a>

```typescript
public readonly warning: number;
```

- *Type:* number

Threshold value for triggering a warning.

---

##### `warningRecovery`<sup>Optional</sup> <a name="warningRecovery" id="cdk-constructs-datadog-go.MonitorThresholds.property.warningRecovery"></a>

```typescript
public readonly warningRecovery: number;
```

- *Type:* number

Threshold value for recovering from a warning state.

---

### MonitorThresholdWindows <a name="MonitorThresholdWindows" id="cdk-constructs-datadog-go.MonitorThresholdWindows"></a>

#### Initializer <a name="Initializer" id="cdk-constructs-datadog-go.MonitorThresholdWindows.Initializer"></a>

```typescript
import { MonitorThresholdWindows } from 'cdk-constructs-datadog-go'

const monitorThresholdWindows: MonitorThresholdWindows = { ... }
```

#### Properties <a name="Properties" id="Properties"></a>

| **Name** | **Type** | **Description** |
| --- | --- | --- |
| <code><a href="#cdk-constructs-datadog-go.MonitorThresholdWindows.property.recoveryWindow">recoveryWindow</a></code> | <code>string</code> | How long an anomalous metric must be normal before recovering from an alert state. |
| <code><a href="#cdk-constructs-datadog-go.MonitorThresholdWindows.property.triggerWindow">triggerWindow</a></code> | <code>string</code> | How long a metric must be anomalous before triggering an alert. |

---

##### `recoveryWindow`<sup>Optional</sup> <a name="recoveryWindow" id="cdk-constructs-datadog-go.MonitorThresholdWindows.property.recoveryWindow"></a>

```typescript
public readonly recoveryWindow: string;
```

- *Type:* string

How long an anomalous metric must be normal before recovering from an alert state.

---

##### `triggerWindow`<sup>Optional</sup> <a name="triggerWindow" id="cdk-constructs-datadog-go.MonitorThresholdWindows.property.triggerWindow"></a>

```typescript
public readonly triggerWindow: string;
```

- *Type:* string

How long a metric must be anomalous before triggering an alert.

---



## Enums <a name="Enums" id="Enums"></a>

### CfnMonitorPropsType <a name="CfnMonitorPropsType" id="cdk-constructs-datadog-go.CfnMonitorPropsType"></a>

The type of the monitor.

#### Members <a name="Members" id="Members"></a>

| **Name** | **Description** |
| --- | --- |
| <code><a href="#cdk-constructs-datadog-go.CfnMonitorPropsType.AUDIT_ALERT">AUDIT_ALERT</a></code> | audit alert. |
| <code><a href="#cdk-constructs-datadog-go.CfnMonitorPropsType.COMPOSITE">COMPOSITE</a></code> | composite. |
| <code><a href="#cdk-constructs-datadog-go.CfnMonitorPropsType.EVENT_ALERT">EVENT_ALERT</a></code> | event alert. |
| <code><a href="#cdk-constructs-datadog-go.CfnMonitorPropsType.EVENT_V2_ALERT">EVENT_V2_ALERT</a></code> | event-v2 alert. |
| <code><a href="#cdk-constructs-datadog-go.CfnMonitorPropsType.LOG_ALERT">LOG_ALERT</a></code> | log alert. |
| <code><a href="#cdk-constructs-datadog-go.CfnMonitorPropsType.METRIC_ALERT">METRIC_ALERT</a></code> | metric alert. |
| <code><a href="#cdk-constructs-datadog-go.CfnMonitorPropsType.PROCESS_ALERT">PROCESS_ALERT</a></code> | process alert. |
| <code><a href="#cdk-constructs-datadog-go.CfnMonitorPropsType.QUERY_ALERT">QUERY_ALERT</a></code> | query alert. |
| <code><a href="#cdk-constructs-datadog-go.CfnMonitorPropsType.SERVICE_CHECK">SERVICE_CHECK</a></code> | service check. |
| <code><a href="#cdk-constructs-datadog-go.CfnMonitorPropsType.SYNTHETICS_ALERT">SYNTHETICS_ALERT</a></code> | synthetics alert. |
| <code><a href="#cdk-constructs-datadog-go.CfnMonitorPropsType.TRACE_ANALYTICS_ALERT">TRACE_ANALYTICS_ALERT</a></code> | trace-analytics alert. |
| <code><a href="#cdk-constructs-datadog-go.CfnMonitorPropsType.SLO_ALERT">SLO_ALERT</a></code> | slo alert. |
| <code><a href="#cdk-constructs-datadog-go.CfnMonitorPropsType.RUM_ALERT">RUM_ALERT</a></code> | rum alert. |
| <code><a href="#cdk-constructs-datadog-go.CfnMonitorPropsType.CI_PIPELINES_ALERT">CI_PIPELINES_ALERT</a></code> | ci-pipelines alert. |
| <code><a href="#cdk-constructs-datadog-go.CfnMonitorPropsType.ERROR_TRACKING_ALERT">ERROR_TRACKING_ALERT</a></code> | error-tracking alert. |
| <code><a href="#cdk-constructs-datadog-go.CfnMonitorPropsType.CI_TESTS_ALERT">CI_TESTS_ALERT</a></code> | ci-tests alert. |

---

##### `AUDIT_ALERT` <a name="AUDIT_ALERT" id="cdk-constructs-datadog-go.CfnMonitorPropsType.AUDIT_ALERT"></a>

audit alert.

---


##### `COMPOSITE` <a name="COMPOSITE" id="cdk-constructs-datadog-go.CfnMonitorPropsType.COMPOSITE"></a>

composite.

---


##### `EVENT_ALERT` <a name="EVENT_ALERT" id="cdk-constructs-datadog-go.CfnMonitorPropsType.EVENT_ALERT"></a>

event alert.

---


##### `EVENT_V2_ALERT` <a name="EVENT_V2_ALERT" id="cdk-constructs-datadog-go.CfnMonitorPropsType.EVENT_V2_ALERT"></a>

event-v2 alert.

---


##### `LOG_ALERT` <a name="LOG_ALERT" id="cdk-constructs-datadog-go.CfnMonitorPropsType.LOG_ALERT"></a>

log alert.

---


##### `METRIC_ALERT` <a name="METRIC_ALERT" id="cdk-constructs-datadog-go.CfnMonitorPropsType.METRIC_ALERT"></a>

metric alert.

---


##### `PROCESS_ALERT` <a name="PROCESS_ALERT" id="cdk-constructs-datadog-go.CfnMonitorPropsType.PROCESS_ALERT"></a>

process alert.

---


##### `QUERY_ALERT` <a name="QUERY_ALERT" id="cdk-constructs-datadog-go.CfnMonitorPropsType.QUERY_ALERT"></a>

query alert.

---


##### `SERVICE_CHECK` <a name="SERVICE_CHECK" id="cdk-constructs-datadog-go.CfnMonitorPropsType.SERVICE_CHECK"></a>

service check.

---


##### `SYNTHETICS_ALERT` <a name="SYNTHETICS_ALERT" id="cdk-constructs-datadog-go.CfnMonitorPropsType.SYNTHETICS_ALERT"></a>

synthetics alert.

---


##### `TRACE_ANALYTICS_ALERT` <a name="TRACE_ANALYTICS_ALERT" id="cdk-constructs-datadog-go.CfnMonitorPropsType.TRACE_ANALYTICS_ALERT"></a>

trace-analytics alert.

---


##### `SLO_ALERT` <a name="SLO_ALERT" id="cdk-constructs-datadog-go.CfnMonitorPropsType.SLO_ALERT"></a>

slo alert.

---


##### `RUM_ALERT` <a name="RUM_ALERT" id="cdk-constructs-datadog-go.CfnMonitorPropsType.RUM_ALERT"></a>

rum alert.

---


##### `CI_PIPELINES_ALERT` <a name="CI_PIPELINES_ALERT" id="cdk-constructs-datadog-go.CfnMonitorPropsType.CI_PIPELINES_ALERT"></a>

ci-pipelines alert.

---


##### `ERROR_TRACKING_ALERT` <a name="ERROR_TRACKING_ALERT" id="cdk-constructs-datadog-go.CfnMonitorPropsType.ERROR_TRACKING_ALERT"></a>

error-tracking alert.

---


##### `CI_TESTS_ALERT` <a name="CI_TESTS_ALERT" id="cdk-constructs-datadog-go.CfnMonitorPropsType.CI_TESTS_ALERT"></a>

ci-tests alert.

---


### MonitorNotificationPresetName <a name="MonitorNotificationPresetName" id="cdk-constructs-datadog-go.MonitorNotificationPresetName"></a>

Toggles the display of additional content sent in the monitor notification.

#### Members <a name="Members" id="Members"></a>

| **Name** | **Description** |
| --- | --- |
| <code><a href="#cdk-constructs-datadog-go.MonitorNotificationPresetName.SHOW_ALL">SHOW_ALL</a></code> | show_all. |
| <code><a href="#cdk-constructs-datadog-go.MonitorNotificationPresetName.HIDE_QUERY">HIDE_QUERY</a></code> | hide_query. |
| <code><a href="#cdk-constructs-datadog-go.MonitorNotificationPresetName.HIDE_HANDLES">HIDE_HANDLES</a></code> | hide_handles. |
| <code><a href="#cdk-constructs-datadog-go.MonitorNotificationPresetName.HIDE_ALL">HIDE_ALL</a></code> | hide_all. |

---

##### `SHOW_ALL` <a name="SHOW_ALL" id="cdk-constructs-datadog-go.MonitorNotificationPresetName.SHOW_ALL"></a>

show_all.

---


##### `HIDE_QUERY` <a name="HIDE_QUERY" id="cdk-constructs-datadog-go.MonitorNotificationPresetName.HIDE_QUERY"></a>

hide_query.

---


##### `HIDE_HANDLES` <a name="HIDE_HANDLES" id="cdk-constructs-datadog-go.MonitorNotificationPresetName.HIDE_HANDLES"></a>

hide_handles.

---


##### `HIDE_ALL` <a name="HIDE_ALL" id="cdk-constructs-datadog-go.MonitorNotificationPresetName.HIDE_ALL"></a>

hide_all.

---


### MonitorOnMissingData <a name="MonitorOnMissingData" id="cdk-constructs-datadog-go.MonitorOnMissingData"></a>

Controls how groups or monitors are treated if an evaluation does not return any data points.

The default option results in different behavior depending on the monitor query type.
For monitors using Count queries, an empty monitor evaluation is treated as 0 and is compared to the threshold conditions.
For monitors using any query type other than Count, for example Gauge, Measure, or Rate, the monitor shows the last known status.
This option is only available for APM Trace Analytics, Audit Trail, CI, Error Tracking, Event, Logs, and RUM monitors.

#### Members <a name="Members" id="Members"></a>

| **Name** | **Description** |
| --- | --- |
| <code><a href="#cdk-constructs-datadog-go.MonitorOnMissingData.DEFAULT">DEFAULT</a></code> | default. |
| <code><a href="#cdk-constructs-datadog-go.MonitorOnMissingData.SHOW_NO_DATA">SHOW_NO_DATA</a></code> | show_no_data. |
| <code><a href="#cdk-constructs-datadog-go.MonitorOnMissingData.SHOW_AND_NOTIFY_NO_DATA">SHOW_AND_NOTIFY_NO_DATA</a></code> | show_and_notify_no_data. |
| <code><a href="#cdk-constructs-datadog-go.MonitorOnMissingData.RESOLVE">RESOLVE</a></code> | resolve. |

---

##### `DEFAULT` <a name="DEFAULT" id="cdk-constructs-datadog-go.MonitorOnMissingData.DEFAULT"></a>

default.

---


##### `SHOW_NO_DATA` <a name="SHOW_NO_DATA" id="cdk-constructs-datadog-go.MonitorOnMissingData.SHOW_NO_DATA"></a>

show_no_data.

---


##### `SHOW_AND_NOTIFY_NO_DATA` <a name="SHOW_AND_NOTIFY_NO_DATA" id="cdk-constructs-datadog-go.MonitorOnMissingData.SHOW_AND_NOTIFY_NO_DATA"></a>

show_and_notify_no_data.

---


##### `RESOLVE` <a name="RESOLVE" id="cdk-constructs-datadog-go.MonitorOnMissingData.RESOLVE"></a>

resolve.

---


### MonitorOptionsRenotifyStatuses <a name="MonitorOptionsRenotifyStatuses" id="cdk-constructs-datadog-go.MonitorOptionsRenotifyStatuses"></a>

#### Members <a name="Members" id="Members"></a>

| **Name** | **Description** |
| --- | --- |
| <code><a href="#cdk-constructs-datadog-go.MonitorOptionsRenotifyStatuses.ALERT">ALERT</a></code> | alert. |
| <code><a href="#cdk-constructs-datadog-go.MonitorOptionsRenotifyStatuses.NO_DATA">NO_DATA</a></code> | no data. |
| <code><a href="#cdk-constructs-datadog-go.MonitorOptionsRenotifyStatuses.WARN">WARN</a></code> | warn. |

---

##### `ALERT` <a name="ALERT" id="cdk-constructs-datadog-go.MonitorOptionsRenotifyStatuses.ALERT"></a>

alert.

---


##### `NO_DATA` <a name="NO_DATA" id="cdk-constructs-datadog-go.MonitorOptionsRenotifyStatuses.NO_DATA"></a>

no data.

---


##### `WARN` <a name="WARN" id="cdk-constructs-datadog-go.MonitorOptionsRenotifyStatuses.WARN"></a>

warn.

---


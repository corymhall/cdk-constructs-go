package cdkmonitoringconstructs

import (
	"github.com/aws/constructs-go/constructs/v10"
)

// Experimental.
type NameResolutionInput struct {
	// Plain name, used in naming alarms.
	//
	// This unique among other resources, and respect the AWS CDK restriction posed on alarm names.
	// The length must be 1 - 255 characters and although the validation rules are undocumented, we recommend using ASCII and hyphens.
	// Default: - derives name from the construct itself.
	//
	// Experimental.
	AlarmFriendlyName *string `field:"optional" json:"alarmFriendlyName" yaml:"alarmFriendlyName"`
	// Human-readable name is a freeform string, used as a caption or description.
	//
	// There are no limitations on what it can be.
	// Default: - use alarmFriendlyName.
	//
	// Experimental.
	HumanReadableName *string `field:"optional" json:"humanReadableName" yaml:"humanReadableName"`
	// If this is defined, the local alarm name prefix used in naming alarms for the construct will be set to this value.
	//
	// The length must be 1 - 255 characters and although the validation rules are undocumented, we recommend using ASCII and hyphens.
	// See: AlarmNamingStrategy for more details on alarm name prefixes.
	//
	// Experimental.
	LocalAlarmNamePrefixOverride *string `field:"optional" json:"localAlarmNamePrefixOverride" yaml:"localAlarmNamePrefixOverride"`
	// Fallback name before we fallback to extracting name from the construct itself.
	//
	// This might be some construct reference, such is cluster ID, stream name, and so on.
	// Default: - use namedConstruct to extract the name.
	//
	// Experimental.
	FallbackConstructName *string `field:"optional" json:"fallbackConstructName" yaml:"fallbackConstructName"`
	// Construct that this naming strategy is naming.
	//
	// It is used as a last resort for naming.
	// Experimental.
	NamedConstruct constructs.IConstruct `field:"optional" json:"namedConstruct" yaml:"namedConstruct"`
}


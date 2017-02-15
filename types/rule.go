package types

// Versions and Prefixes used in API and KV URLs
const (

	// RuleAPIPrefix is a partial path to the HTTP endpoint.
	RuleAPIPrefix string = "rules"

	// RuleActionAdd specifies to add labels to the object(s).
	RuleActionAdd RuleAction = "add"

	// RuleActionRemove specifies to remove labels from the object(s).
	RuleActionRemove RuleAction = "remove"
)

// RuleAction - rule action type
type RuleAction string

// Rule is used to define a rule
type Rule struct {

	// Volume unique ID.
	// Read Only: true
	ID string `json:"ID,omitempty"`

	// Rule name.
	// Required: true
	Name string `json:"name"`

	// Rule description.
	Description string `json:"description"`

	// Flag describing whether rule is active.
	// Required: true
	// Default: false
	Active bool `json:"active"`

	// Weight is used to determine order during rule processing.  Rules with
	// heavier weights are processed later.
	// default: 0
	Weight int `json:"weight"`

	// Operator is used to compare objects or labels.
	Operator Operator `json:"operator"`

	// RuleAction controls whether the action is to add or remove a label from the
	// matching object(s).
	// default: add
	RuleAction RuleAction `json:"rule_action"`

	// Selectors defines the list of labels that should trigger a rule.
	Selectors map[string]string `json:"selectors"`

	// Labels defines the list of labels that will be added or removed from the
	// matching object(s).
	Labels map[string]string `json:"labels"`
}

// Rules is a collection of Rules.
type Rules []*Rule

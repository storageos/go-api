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

	// Rule unique ID.
	// Read Only: true
	ID string `json:"ID,omitempty"`

	// Rule name.
	// Required: true
	Name string `json:"Name"`

	// Rule description.
	Description string `json:"Description"`

	// Flag describing whether the rule is active.
	// Default: false
	Active bool `json:"Active"`

	// Weight is used to determine order during rule processing.  Rules with
	// heavier weights are processed later.
	// default: 0
	Weight int `json:"Weight"`

	// Operator is used to compare objects or labels.
	Operator Operator `json:"Operator"`

	// RuleAction controls whether the action is to add or remove a label from the
	// matching object(s).
	RuleAction RuleAction `json:"RuleAction"`

	// Selectors defines the list of labels that should trigger a rule.
	Selectors map[string]string `json:"Selectors"`

	// Labels defines the list of labels that will be added or removed from the
	// matching object(s).
	Labels map[string]string `json:"Labels"`
}

// Rules is a collection of Rules.
type Rules []*Rule

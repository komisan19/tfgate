package iamrules

// Represents Terraform change operations.
type Operation string

const (
	OpCreate Operation = "create"
	OpDelete Operation = "delete"
	OpUpdate Operation = "update"
)

// Rule is an IAM action requirement for one resource type × one operation.
type Rule struct {
	BaseActions        []string
	UpdateActions      []string
	ConditionalActions map[string][]string
}

// registry is inner dictionary.
// for example, resourceType: "aws_s3_bucket"
var registry = map[string]map[Operation]Rule{}

func register(resourceType string, op Operation, rule Rule) {
	if registry[resourceType] == nil {
		registry[resourceType] = map[Operation]Rule{}
	}

	registry[resourceType][op] = rule
}

// Lookup retrieves a Rule based on the resource type and operation.
func Lookup(resourceType string, op Operation) (Rule, bool) {
	inner, ok := registry[resourceType]
	if !ok {
		return Rule{}, false
	}

	rule, ok := inner[op]
	if !ok {
		return Rule{}, false
	}
	return rule, true
}

// Resolve function returns a list of necessary actions from the Rule.
func Resolve(rule Rule, op Operation, changedKeys []string) []string {
	seen := map[string]bool{}
	var result []string

	add := func(actions []string) {
		for _, a := range actions {
			if !seen[a] {
				seen[a] = true
				result = append(result, a)
			}
		}
	}

	switch op {
	case OpUpdate:
		add(rule.UpdateActions)
	default:
		add(rule.BaseActions)
	}

	for _, key := range changedKeys {
		if actions, ok := rule.ConditionalActions[key]; ok {
			add(actions)
		}
	}

	return result
}

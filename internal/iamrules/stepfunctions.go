package iamrules

func init() {
	// aws_sfn_state_machine
	register("aws_sfn_state_machine", OpCreate, Rule{
		BaseActions: []string{
			"states:CreateStateMachine",
			"states:DescribeStateMachine",
			"states:TagResource",
		},
	})
	register("aws_sfn_state_machine", OpDelete, Rule{
		BaseActions: []string{
			"states:DeleteStateMachine",
			"states:DescribeStateMachine",
		},
	})
	register("aws_sfn_state_machine", OpUpdate, Rule{
		UpdateActions: []string{"states:DescribeStateMachine"},
		ConditionalActions: map[string][]string{
			"definition":             {"states:UpdateStateMachine"},
			"role_arn":               {"states:UpdateStateMachine"},
			"logging_configuration":  {"states:UpdateStateMachine"},
			"tracing_configuration":  {"states:UpdateStateMachine"},
			"type":                   {"states:UpdateStateMachine"},
			"tags":                   {"states:TagResource", "states:UntagResource"},
		},
	})

	// aws_sfn_activity
	register("aws_sfn_activity", OpCreate, Rule{
		BaseActions: []string{
			"states:CreateActivity",
			"states:DescribeActivity",
			"states:TagResource",
		},
	})
	register("aws_sfn_activity", OpDelete, Rule{
		BaseActions: []string{
			"states:DeleteActivity",
			"states:DescribeActivity",
		},
	})
	register("aws_sfn_activity", OpUpdate, Rule{
		UpdateActions: []string{"states:DescribeActivity"},
		ConditionalActions: map[string][]string{
			"tags": {"states:TagResource", "states:UntagResource"},
		},
	})
}

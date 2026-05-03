package iamrules

func init() {
	// aws_cloudwatch_event_rule (EventBridge)
	register("aws_cloudwatch_event_rule", OpCreate, Rule{
		BaseActions: []string{
			"events:PutRule",
			"events:DescribeRule",
			"events:TagResource",
		},
	})
	register("aws_cloudwatch_event_rule", OpDelete, Rule{
		BaseActions: []string{
			"events:DeleteRule",
			"events:DescribeRule",
			"events:RemoveTargets",
		},
	})
	register("aws_cloudwatch_event_rule", OpUpdate, Rule{
		UpdateActions: []string{"events:DescribeRule"},
		ConditionalActions: map[string][]string{
			"schedule_expression": {"events:PutRule"},
			"event_pattern":       {"events:PutRule"},
			"event_bus_name":      {"events:PutRule"},
			"role_arn":            {"events:PutRule"},
			"description":         {"events:PutRule"},
			"is_enabled":          {"events:EnableRule", "events:DisableRule"},
			"tags":                {"events:TagResource", "events:UntagResource"},
		},
	})

	// aws_cloudwatch_event_target
	register("aws_cloudwatch_event_target", OpCreate, Rule{
		BaseActions: []string{
			"events:PutTargets",
			"events:ListTargetsByRule",
		},
	})
	register("aws_cloudwatch_event_target", OpDelete, Rule{
		BaseActions: []string{
			"events:RemoveTargets",
			"events:ListTargetsByRule",
		},
	})
	register("aws_cloudwatch_event_target", OpUpdate, Rule{
		UpdateActions: []string{"events:ListTargetsByRule"},
		ConditionalActions: map[string][]string{
			"arn":            {"events:PutTargets"},
			"input":          {"events:PutTargets"},
			"input_path":     {"events:PutTargets"},
			"input_transformer": {"events:PutTargets"},
			"run_command_targets": {"events:PutTargets"},
			"ecs_target":     {"events:PutTargets"},
		},
	})

	// aws_cloudwatch_event_bus
	register("aws_cloudwatch_event_bus", OpCreate, Rule{
		BaseActions: []string{
			"events:CreateEventBus",
			"events:DescribeEventBus",
			"events:TagResource",
		},
	})
	register("aws_cloudwatch_event_bus", OpDelete, Rule{
		BaseActions: []string{
			"events:DeleteEventBus",
			"events:DescribeEventBus",
		},
	})
	register("aws_cloudwatch_event_bus", OpUpdate, Rule{
		UpdateActions: []string{"events:DescribeEventBus"},
		ConditionalActions: map[string][]string{
			"tags": {"events:TagResource", "events:UntagResource"},
		},
	})

	// aws_cloudwatch_event_permission (resource-based policy)
	register("aws_cloudwatch_event_permission", OpCreate, Rule{
		BaseActions: []string{
			"events:PutPermission",
			"events:DescribeEventBus",
		},
	})
	register("aws_cloudwatch_event_permission", OpDelete, Rule{
		BaseActions: []string{
			"events:RemovePermission",
			"events:DescribeEventBus",
		},
	})
	register("aws_cloudwatch_event_permission", OpUpdate, Rule{
		UpdateActions: []string{"events:DescribeEventBus"},
		ConditionalActions: map[string][]string{
			"action":    {"events:PutPermission"},
			"principal": {"events:PutPermission"},
			"condition": {"events:PutPermission"},
		},
	})
}

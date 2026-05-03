package iamrules

func init() {
	// aws_ecs_cluster
	register("aws_ecs_cluster", OpCreate, Rule{
		BaseActions: []string{
			"ecs:CreateCluster",
			"ecs:DescribeClusters",
			"ecs:TagResource",
		},
	})
	register("aws_ecs_cluster", OpDelete, Rule{
		BaseActions: []string{
			"ecs:DeleteCluster",
			"ecs:DescribeClusters",
		},
	})
	register("aws_ecs_cluster", OpUpdate, Rule{
		UpdateActions: []string{"ecs:DescribeClusters"},
		ConditionalActions: map[string][]string{
			"setting":             {"ecs:UpdateCluster"},
			"configuration":       {"ecs:UpdateCluster"},
			"capacity_providers":  {"ecs:PutClusterCapacityProviders"},
			"default_capacity_provider_strategy": {"ecs:PutClusterCapacityProviders"},
			"tags":                {"ecs:TagResource", "ecs:UntagResource"},
		},
	})

	// aws_ecs_service
	register("aws_ecs_service", OpCreate, Rule{
		BaseActions: []string{
			"ecs:CreateService",
			"ecs:DescribeServices",
			"ecs:TagResource",
		},
	})
	register("aws_ecs_service", OpDelete, Rule{
		BaseActions: []string{
			"ecs:DeleteService",
			"ecs:DescribeServices",
			"ecs:UpdateService",
		},
	})
	register("aws_ecs_service", OpUpdate, Rule{
		UpdateActions: []string{"ecs:DescribeServices"},
		ConditionalActions: map[string][]string{
			"desired_count":              {"ecs:UpdateService"},
			"task_definition":            {"ecs:UpdateService"},
			"deployment_configuration":   {"ecs:UpdateService"},
			"network_configuration":      {"ecs:UpdateService"},
			"load_balancer":              {"ecs:UpdateService"},
			"capacity_provider_strategy": {"ecs:UpdateService"},
			"health_check_grace_period_seconds": {"ecs:UpdateService"},
			"enable_execute_command":     {"ecs:UpdateService"},
			"force_new_deployment":       {"ecs:UpdateService"},
			"tags":                       {"ecs:TagResource", "ecs:UntagResource"},
		},
	})

	// aws_ecs_task_definition (immutable; update creates a new revision)
	register("aws_ecs_task_definition", OpCreate, Rule{
		BaseActions: []string{
			"ecs:RegisterTaskDefinition",
			"ecs:DescribeTaskDefinition",
			"ecs:TagResource",
		},
	})
	register("aws_ecs_task_definition", OpDelete, Rule{
		BaseActions: []string{
			"ecs:DeregisterTaskDefinition",
			"ecs:DescribeTaskDefinition",
		},
	})

	// aws_ecs_capacity_provider
	register("aws_ecs_capacity_provider", OpCreate, Rule{
		BaseActions: []string{
			"ecs:CreateCapacityProvider",
			"ecs:DescribeCapacityProviders",
			"ecs:TagResource",
		},
	})
	register("aws_ecs_capacity_provider", OpDelete, Rule{
		BaseActions: []string{
			"ecs:DeleteCapacityProvider",
			"ecs:DescribeCapacityProviders",
		},
	})
	register("aws_ecs_capacity_provider", OpUpdate, Rule{
		UpdateActions: []string{"ecs:DescribeCapacityProviders"},
		ConditionalActions: map[string][]string{
			"auto_scaling_group_provider": {"ecs:UpdateCapacityProvider"},
			"tags":                        {"ecs:TagResource", "ecs:UntagResource"},
		},
	})
}

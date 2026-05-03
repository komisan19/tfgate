package iamrules

func init() {
	// aws_db_instance
	register("aws_db_instance", OpCreate, Rule{
		BaseActions: []string{
			"rds:CreateDBInstance",
			"rds:DescribeDBInstances",
			"rds:AddTagsToResource",
		},
	})
	register("aws_db_instance", OpDelete, Rule{
		BaseActions: []string{
			"rds:DeleteDBInstance",
			"rds:DescribeDBInstances",
		},
	})
	register("aws_db_instance", OpUpdate, Rule{
		UpdateActions: []string{"rds:DescribeDBInstances"},
		ConditionalActions: map[string][]string{
			"instance_class":                  {"rds:ModifyDBInstance"},
			"allocated_storage":               {"rds:ModifyDBInstance"},
			"engine_version":                  {"rds:ModifyDBInstance"},
			"password":                        {"rds:ModifyDBInstance"},
			"backup_retention_period":         {"rds:ModifyDBInstance"},
			"backup_window":                   {"rds:ModifyDBInstance"},
			"maintenance_window":              {"rds:ModifyDBInstance"},
			"multi_az":                        {"rds:ModifyDBInstance"},
			"publicly_accessible":             {"rds:ModifyDBInstance"},
			"storage_type":                    {"rds:ModifyDBInstance"},
			"iops":                            {"rds:ModifyDBInstance"},
			"auto_minor_version_upgrade":      {"rds:ModifyDBInstance"},
			"allow_major_version_upgrade":     {"rds:ModifyDBInstance"},
			"parameter_group_name":            {"rds:ModifyDBInstance"},
			"option_group_name":               {"rds:ModifyDBInstance"},
			"db_subnet_group_name":            {"rds:ModifyDBInstance"},
			"vpc_security_group_ids":          {"rds:ModifyDBInstance"},
			"performance_insights_enabled":    {"rds:ModifyDBInstance"},
			"deletion_protection":             {"rds:ModifyDBInstance"},
			"enabled_cloudwatch_logs_exports": {"rds:ModifyDBInstance"},
			"tags":                            {"rds:AddTagsToResource", "rds:RemoveTagsFromResource"},
		},
	})

	// aws_db_subnet_group
	register("aws_db_subnet_group", OpCreate, Rule{
		BaseActions: []string{
			"rds:CreateDBSubnetGroup",
			"rds:DescribeDBSubnetGroups",
			"rds:AddTagsToResource",
		},
	})
	register("aws_db_subnet_group", OpDelete, Rule{
		BaseActions: []string{
			"rds:DeleteDBSubnetGroup",
			"rds:DescribeDBSubnetGroups",
		},
	})
	register("aws_db_subnet_group", OpUpdate, Rule{
		UpdateActions: []string{"rds:DescribeDBSubnetGroups"},
		ConditionalActions: map[string][]string{
			"subnet_ids":   {"rds:ModifyDBSubnetGroup"},
			"description":  {"rds:ModifyDBSubnetGroup"},
			"tags":         {"rds:AddTagsToResource", "rds:RemoveTagsFromResource"},
		},
	})

	// aws_db_parameter_group
	register("aws_db_parameter_group", OpCreate, Rule{
		BaseActions: []string{
			"rds:CreateDBParameterGroup",
			"rds:DescribeDBParameterGroups",
			"rds:ModifyDBParameterGroup",
			"rds:AddTagsToResource",
		},
	})
	register("aws_db_parameter_group", OpDelete, Rule{
		BaseActions: []string{
			"rds:DeleteDBParameterGroup",
			"rds:DescribeDBParameterGroups",
		},
	})
	register("aws_db_parameter_group", OpUpdate, Rule{
		UpdateActions: []string{"rds:DescribeDBParameterGroups"},
		ConditionalActions: map[string][]string{
			"parameter": {"rds:ModifyDBParameterGroup"},
			"tags":      {"rds:AddTagsToResource", "rds:RemoveTagsFromResource"},
		},
	})

	// aws_db_option_group
	register("aws_db_option_group", OpCreate, Rule{
		BaseActions: []string{
			"rds:CreateOptionGroup",
			"rds:DescribeOptionGroups",
			"rds:AddTagsToResource",
		},
	})
	register("aws_db_option_group", OpDelete, Rule{
		BaseActions: []string{
			"rds:DeleteOptionGroup",
			"rds:DescribeOptionGroups",
		},
	})
	register("aws_db_option_group", OpUpdate, Rule{
		UpdateActions: []string{"rds:DescribeOptionGroups"},
		ConditionalActions: map[string][]string{
			"option": {"rds:ModifyOptionGroup"},
			"tags":   {"rds:AddTagsToResource", "rds:RemoveTagsFromResource"},
		},
	})

	// aws_rds_cluster (Aurora)
	register("aws_rds_cluster", OpCreate, Rule{
		BaseActions: []string{
			"rds:CreateDBCluster",
			"rds:DescribeDBClusters",
			"rds:AddTagsToResource",
		},
	})
	register("aws_rds_cluster", OpDelete, Rule{
		BaseActions: []string{
			"rds:DeleteDBCluster",
			"rds:DescribeDBClusters",
		},
	})
	register("aws_rds_cluster", OpUpdate, Rule{
		UpdateActions: []string{"rds:DescribeDBClusters"},
		ConditionalActions: map[string][]string{
			"master_password":             {"rds:ModifyDBCluster"},
			"backup_retention_period":     {"rds:ModifyDBCluster"},
			"preferred_backup_window":     {"rds:ModifyDBCluster"},
			"preferred_maintenance_window": {"rds:ModifyDBCluster"},
			"vpc_security_group_ids":      {"rds:ModifyDBCluster"},
			"db_cluster_parameter_group_name": {"rds:ModifyDBCluster"},
			"engine_version":              {"rds:ModifyDBCluster"},
			"deletion_protection":         {"rds:ModifyDBCluster"},
			"enabled_cloudwatch_logs_exports": {"rds:ModifyDBCluster"},
			"scaling_configuration":       {"rds:ModifyDBCluster"},
			"tags":                        {"rds:AddTagsToResource", "rds:RemoveTagsFromResource"},
		},
	})

	// aws_rds_cluster_instance
	register("aws_rds_cluster_instance", OpCreate, Rule{
		BaseActions: []string{
			"rds:CreateDBInstance",
			"rds:DescribeDBInstances",
			"rds:AddTagsToResource",
		},
	})
	register("aws_rds_cluster_instance", OpDelete, Rule{
		BaseActions: []string{
			"rds:DeleteDBInstance",
			"rds:DescribeDBInstances",
		},
	})
	register("aws_rds_cluster_instance", OpUpdate, Rule{
		UpdateActions: []string{"rds:DescribeDBInstances"},
		ConditionalActions: map[string][]string{
			"instance_class":             {"rds:ModifyDBInstance"},
			"auto_minor_version_upgrade": {"rds:ModifyDBInstance"},
			"promotion_tier":             {"rds:ModifyDBInstance"},
			"tags":                       {"rds:AddTagsToResource", "rds:RemoveTagsFromResource"},
		},
	})
}

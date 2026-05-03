package iamrules

func init() {
	// aws_elasticache_cluster
	register("aws_elasticache_cluster", OpCreate, Rule{
		BaseActions: []string{
			"elasticache:CreateCacheCluster",
			"elasticache:DescribeCacheClusters",
			"elasticache:AddTagsToResource",
		},
	})
	register("aws_elasticache_cluster", OpDelete, Rule{
		BaseActions: []string{
			"elasticache:DeleteCacheCluster",
			"elasticache:DescribeCacheClusters",
		},
	})
	register("aws_elasticache_cluster", OpUpdate, Rule{
		UpdateActions: []string{"elasticache:DescribeCacheClusters"},
		ConditionalActions: map[string][]string{
			"num_cache_nodes":               {"elasticache:ModifyCacheCluster"},
			"node_type":                     {"elasticache:ModifyCacheCluster"},
			"engine_version":                {"elasticache:ModifyCacheCluster"},
			"maintenance_window":            {"elasticache:ModifyCacheCluster"},
			"snapshot_window":               {"elasticache:ModifyCacheCluster"},
			"snapshot_retention_limit":      {"elasticache:ModifyCacheCluster"},
			"parameter_group_name":          {"elasticache:ModifyCacheCluster"},
			"security_group_ids":            {"elasticache:ModifyCacheCluster"},
			"notification_topic_arn":        {"elasticache:ModifyCacheCluster"},
			"auto_minor_version_upgrade":    {"elasticache:ModifyCacheCluster"},
			"az_mode":                       {"elasticache:ModifyCacheCluster"},
			"preferred_availability_zones":  {"elasticache:ModifyCacheCluster"},
			"tags":                          {"elasticache:AddTagsToResource", "elasticache:RemoveTagsFromResource"},
		},
	})

	// aws_elasticache_replication_group
	register("aws_elasticache_replication_group", OpCreate, Rule{
		BaseActions: []string{
			"elasticache:CreateReplicationGroup",
			"elasticache:DescribeReplicationGroups",
			"elasticache:AddTagsToResource",
		},
	})
	register("aws_elasticache_replication_group", OpDelete, Rule{
		BaseActions: []string{
			"elasticache:DeleteReplicationGroup",
			"elasticache:DescribeReplicationGroups",
		},
	})
	register("aws_elasticache_replication_group", OpUpdate, Rule{
		UpdateActions: []string{"elasticache:DescribeReplicationGroups"},
		ConditionalActions: map[string][]string{
			"num_cache_clusters":            {"elasticache:ModifyReplicationGroup"},
			"num_node_groups":               {"elasticache:ModifyReplicationGroupShardConfiguration"},
			"replicas_per_node_group":       {"elasticache:ModifyReplicationGroupShardConfiguration"},
			"node_type":                     {"elasticache:ModifyReplicationGroup"},
			"description":                   {"elasticache:ModifyReplicationGroup"},
			"automatic_failover_enabled":    {"elasticache:ModifyReplicationGroup"},
			"multi_az_enabled":              {"elasticache:ModifyReplicationGroup"},
			"engine_version":                {"elasticache:ModifyReplicationGroup"},
			"parameter_group_name":          {"elasticache:ModifyReplicationGroup"},
			"security_group_ids":            {"elasticache:ModifyReplicationGroup"},
			"maintenance_window":            {"elasticache:ModifyReplicationGroup"},
			"snapshot_window":               {"elasticache:ModifyReplicationGroup"},
			"snapshot_retention_limit":      {"elasticache:ModifyReplicationGroup"},
			"auth_token":                    {"elasticache:ModifyReplicationGroup"},
			"at_rest_encryption_enabled":    {"elasticache:ModifyReplicationGroup"},
			"auto_minor_version_upgrade":    {"elasticache:ModifyReplicationGroup"},
			"tags":                          {"elasticache:AddTagsToResource", "elasticache:RemoveTagsFromResource"},
		},
	})

	// aws_elasticache_parameter_group
	register("aws_elasticache_parameter_group", OpCreate, Rule{
		BaseActions: []string{
			"elasticache:CreateCacheParameterGroup",
			"elasticache:DescribeCacheParameterGroups",
			"elasticache:ModifyCacheParameterGroup",
			"elasticache:AddTagsToResource",
		},
	})
	register("aws_elasticache_parameter_group", OpDelete, Rule{
		BaseActions: []string{
			"elasticache:DeleteCacheParameterGroup",
			"elasticache:DescribeCacheParameterGroups",
			"elasticache:ResetCacheParameterGroup",
		},
	})
	register("aws_elasticache_parameter_group", OpUpdate, Rule{
		UpdateActions: []string{"elasticache:DescribeCacheParameterGroups"},
		ConditionalActions: map[string][]string{
			"parameter": {"elasticache:ModifyCacheParameterGroup"},
			"tags":      {"elasticache:AddTagsToResource", "elasticache:RemoveTagsFromResource"},
		},
	})

	// aws_elasticache_subnet_group
	register("aws_elasticache_subnet_group", OpCreate, Rule{
		BaseActions: []string{
			"elasticache:CreateCacheSubnetGroup",
			"elasticache:DescribeCacheSubnetGroups",
			"elasticache:AddTagsToResource",
		},
	})
	register("aws_elasticache_subnet_group", OpDelete, Rule{
		BaseActions: []string{
			"elasticache:DeleteCacheSubnetGroup",
			"elasticache:DescribeCacheSubnetGroups",
		},
	})
	register("aws_elasticache_subnet_group", OpUpdate, Rule{
		UpdateActions: []string{"elasticache:DescribeCacheSubnetGroups"},
		ConditionalActions: map[string][]string{
			"subnet_ids":  {"elasticache:ModifyCacheSubnetGroup"},
			"description": {"elasticache:ModifyCacheSubnetGroup"},
			"tags":        {"elasticache:AddTagsToResource", "elasticache:RemoveTagsFromResource"},
		},
	})
}

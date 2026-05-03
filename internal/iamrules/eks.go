package iamrules

func init() {
	// aws_eks_cluster
	register("aws_eks_cluster", OpCreate, Rule{
		BaseActions: []string{
			"eks:CreateCluster",
			"eks:DescribeCluster",
			"eks:TagResource",
		},
	})
	register("aws_eks_cluster", OpDelete, Rule{
		BaseActions: []string{
			"eks:DeleteCluster",
			"eks:DescribeCluster",
		},
	})
	register("aws_eks_cluster", OpUpdate, Rule{
		UpdateActions: []string{"eks:DescribeCluster"},
		ConditionalActions: map[string][]string{
			"version":                  {"eks:UpdateClusterVersion"},
			"kubernetes_network_config": {"eks:UpdateClusterConfig"},
			"resources_vpc_config":     {"eks:UpdateClusterConfig"},
			"enabled_cluster_log_types": {"eks:UpdateClusterConfig"},
			"logging":                  {"eks:UpdateClusterConfig"},
			"tags":                     {"eks:TagResource", "eks:UntagResource"},
		},
	})

	// aws_eks_node_group
	register("aws_eks_node_group", OpCreate, Rule{
		BaseActions: []string{
			"eks:CreateNodegroup",
			"eks:DescribeNodegroup",
			"eks:TagResource",
		},
	})
	register("aws_eks_node_group", OpDelete, Rule{
		BaseActions: []string{
			"eks:DeleteNodegroup",
			"eks:DescribeNodegroup",
		},
	})
	register("aws_eks_node_group", OpUpdate, Rule{
		UpdateActions: []string{"eks:DescribeNodegroup"},
		ConditionalActions: map[string][]string{
			"scaling_config":  {"eks:UpdateNodegroupConfig"},
			"update_config":   {"eks:UpdateNodegroupConfig"},
			"labels":          {"eks:UpdateNodegroupConfig"},
			"taints":          {"eks:UpdateNodegroupConfig"},
			"version":         {"eks:UpdateNodegroupVersion"},
			"release_version": {"eks:UpdateNodegroupVersion"},
			"ami_type":        {"eks:UpdateNodegroupVersion"},
			"tags":            {"eks:TagResource", "eks:UntagResource"},
		},
	})

	// aws_eks_fargate_profile
	register("aws_eks_fargate_profile", OpCreate, Rule{
		BaseActions: []string{
			"eks:CreateFargateProfile",
			"eks:DescribeFargateProfile",
			"eks:TagResource",
		},
	})
	register("aws_eks_fargate_profile", OpDelete, Rule{
		BaseActions: []string{
			"eks:DeleteFargateProfile",
			"eks:DescribeFargateProfile",
		},
	})
	register("aws_eks_fargate_profile", OpUpdate, Rule{
		UpdateActions: []string{"eks:DescribeFargateProfile"},
		ConditionalActions: map[string][]string{
			"tags": {"eks:TagResource", "eks:UntagResource"},
		},
	})

	// aws_eks_addon
	register("aws_eks_addon", OpCreate, Rule{
		BaseActions: []string{
			"eks:CreateAddon",
			"eks:DescribeAddon",
			"eks:TagResource",
		},
	})
	register("aws_eks_addon", OpDelete, Rule{
		BaseActions: []string{
			"eks:DeleteAddon",
			"eks:DescribeAddon",
		},
	})
	register("aws_eks_addon", OpUpdate, Rule{
		UpdateActions: []string{"eks:DescribeAddon"},
		ConditionalActions: map[string][]string{
			"addon_version":               {"eks:UpdateAddon"},
			"service_account_role_arn":    {"eks:UpdateAddon"},
			"resolve_conflicts":           {"eks:UpdateAddon"},
			"configuration_values":        {"eks:UpdateAddon"},
			"tags":                        {"eks:TagResource", "eks:UntagResource"},
		},
	})
}

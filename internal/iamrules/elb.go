package iamrules

func init() {
	// aws_lb (ALB / NLB)
	register("aws_lb", OpCreate, Rule{
		BaseActions: []string{
			"elasticloadbalancing:CreateLoadBalancer",
			"elasticloadbalancing:DescribeLoadBalancers",
			"elasticloadbalancing:ModifyLoadBalancerAttributes",
			"elasticloadbalancing:AddTags",
		},
	})
	register("aws_lb", OpDelete, Rule{
		BaseActions: []string{
			"elasticloadbalancing:DeleteLoadBalancer",
			"elasticloadbalancing:DescribeLoadBalancers",
		},
	})
	register("aws_lb", OpUpdate, Rule{
		UpdateActions: []string{"elasticloadbalancing:DescribeLoadBalancers"},
		ConditionalActions: map[string][]string{
			"access_logs":                      {"elasticloadbalancing:ModifyLoadBalancerAttributes"},
			"idle_timeout":                     {"elasticloadbalancing:ModifyLoadBalancerAttributes"},
			"enable_deletion_protection":       {"elasticloadbalancing:ModifyLoadBalancerAttributes"},
			"enable_http2":                     {"elasticloadbalancing:ModifyLoadBalancerAttributes"},
			"enable_cross_zone_load_balancing": {"elasticloadbalancing:ModifyLoadBalancerAttributes"},
			"drop_invalid_header_fields":       {"elasticloadbalancing:ModifyLoadBalancerAttributes"},
			"security_groups":                  {"elasticloadbalancing:SetSecurityGroups"},
			"subnets":                          {"elasticloadbalancing:SetSubnets"},
			"subnet_mapping":                   {"elasticloadbalancing:SetSubnets"},
			"tags":                             {"elasticloadbalancing:AddTags", "elasticloadbalancing:RemoveTags"},
		},
	})

	// aws_lb_listener
	register("aws_lb_listener", OpCreate, Rule{
		BaseActions: []string{
			"elasticloadbalancing:CreateListener",
			"elasticloadbalancing:DescribeListeners",
			"elasticloadbalancing:AddTags",
		},
	})
	register("aws_lb_listener", OpDelete, Rule{
		BaseActions: []string{
			"elasticloadbalancing:DeleteListener",
			"elasticloadbalancing:DescribeListeners",
		},
	})
	register("aws_lb_listener", OpUpdate, Rule{
		UpdateActions: []string{"elasticloadbalancing:DescribeListeners"},
		ConditionalActions: map[string][]string{
			"default_action":   {"elasticloadbalancing:ModifyListener"},
			"port":             {"elasticloadbalancing:ModifyListener"},
			"protocol":         {"elasticloadbalancing:ModifyListener"},
			"ssl_policy":       {"elasticloadbalancing:ModifyListener"},
			"certificate_arn":  {"elasticloadbalancing:ModifyListener"},
			"alpn_policy":      {"elasticloadbalancing:ModifyListener"},
			"tags":             {"elasticloadbalancing:AddTags", "elasticloadbalancing:RemoveTags"},
		},
	})

	// aws_lb_listener_certificate
	register("aws_lb_listener_certificate", OpCreate, Rule{
		BaseActions: []string{"elasticloadbalancing:AddListenerCertificates"},
	})
	register("aws_lb_listener_certificate", OpDelete, Rule{
		BaseActions: []string{"elasticloadbalancing:RemoveListenerCertificates"},
	})

	// aws_lb_target_group
	register("aws_lb_target_group", OpCreate, Rule{
		BaseActions: []string{
			"elasticloadbalancing:CreateTargetGroup",
			"elasticloadbalancing:DescribeTargetGroups",
			"elasticloadbalancing:AddTags",
		},
	})
	register("aws_lb_target_group", OpDelete, Rule{
		BaseActions: []string{
			"elasticloadbalancing:DeleteTargetGroup",
			"elasticloadbalancing:DescribeTargetGroups",
		},
	})
	register("aws_lb_target_group", OpUpdate, Rule{
		UpdateActions: []string{"elasticloadbalancing:DescribeTargetGroups"},
		ConditionalActions: map[string][]string{
			"health_check":          {"elasticloadbalancing:ModifyTargetGroup"},
			"stickiness":            {"elasticloadbalancing:ModifyTargetGroupAttributes"},
			"deregistration_delay":  {"elasticloadbalancing:ModifyTargetGroupAttributes"},
			"load_balancing_algorithm_type": {"elasticloadbalancing:ModifyTargetGroupAttributes"},
			"tags":                  {"elasticloadbalancing:AddTags", "elasticloadbalancing:RemoveTags"},
		},
	})

	// aws_lb_target_group_attachment
	register("aws_lb_target_group_attachment", OpCreate, Rule{
		BaseActions: []string{
			"elasticloadbalancing:RegisterTargets",
			"elasticloadbalancing:DescribeTargetHealth",
		},
	})
	register("aws_lb_target_group_attachment", OpDelete, Rule{
		BaseActions: []string{"elasticloadbalancing:DeregisterTargets"},
	})

	// aws_lb_listener_rule
	register("aws_lb_listener_rule", OpCreate, Rule{
		BaseActions: []string{
			"elasticloadbalancing:CreateRule",
			"elasticloadbalancing:DescribeRules",
			"elasticloadbalancing:AddTags",
		},
	})
	register("aws_lb_listener_rule", OpDelete, Rule{
		BaseActions: []string{
			"elasticloadbalancing:DeleteRule",
			"elasticloadbalancing:DescribeRules",
		},
	})
	register("aws_lb_listener_rule", OpUpdate, Rule{
		UpdateActions: []string{"elasticloadbalancing:DescribeRules"},
		ConditionalActions: map[string][]string{
			"action":    {"elasticloadbalancing:ModifyRule"},
			"condition": {"elasticloadbalancing:ModifyRule"},
			"priority":  {"elasticloadbalancing:SetRulePriorities"},
			"tags":      {"elasticloadbalancing:AddTags", "elasticloadbalancing:RemoveTags"},
		},
	})
}

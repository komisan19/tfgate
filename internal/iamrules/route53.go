package iamrules

func init() {
	// aws_route53_zone
	register("aws_route53_zone", OpCreate, Rule{
		BaseActions: []string{
			"route53:CreateHostedZone",
			"route53:GetHostedZone",
			"route53:ChangeTagsForResource",
		},
	})
	register("aws_route53_zone", OpDelete, Rule{
		BaseActions: []string{
			"route53:DeleteHostedZone",
			"route53:GetHostedZone",
			"route53:ListResourceRecordSets",
		},
	})
	register("aws_route53_zone", OpUpdate, Rule{
		UpdateActions: []string{"route53:GetHostedZone"},
		ConditionalActions: map[string][]string{
			"comment": {"route53:UpdateHostedZoneComment"},
			"tags":    {"route53:ChangeTagsForResource"},
			"vpc":     {"route53:AssociateVPCWithHostedZone", "route53:DisassociateVPCFromHostedZone"},
		},
	})

	// aws_route53_record
	register("aws_route53_record", OpCreate, Rule{
		BaseActions: []string{
			"route53:ChangeResourceRecordSets",
			"route53:GetChange",
			"route53:ListResourceRecordSets",
		},
	})
	register("aws_route53_record", OpDelete, Rule{
		BaseActions: []string{
			"route53:ChangeResourceRecordSets",
			"route53:GetChange",
			"route53:ListResourceRecordSets",
		},
	})
	register("aws_route53_record", OpUpdate, Rule{
		UpdateActions: []string{"route53:ListResourceRecordSets"},
		ConditionalActions: map[string][]string{
			"records":                  {"route53:ChangeResourceRecordSets", "route53:GetChange"},
			"ttl":                      {"route53:ChangeResourceRecordSets", "route53:GetChange"},
			"alias":                    {"route53:ChangeResourceRecordSets", "route53:GetChange"},
			"weighted_routing_policy":  {"route53:ChangeResourceRecordSets", "route53:GetChange"},
			"latency_routing_policy":   {"route53:ChangeResourceRecordSets", "route53:GetChange"},
			"failover_routing_policy":  {"route53:ChangeResourceRecordSets", "route53:GetChange"},
			"geolocation_routing_policy": {"route53:ChangeResourceRecordSets", "route53:GetChange"},
			"health_check_id":          {"route53:ChangeResourceRecordSets", "route53:GetChange"},
		},
	})

	// aws_route53_health_check
	register("aws_route53_health_check", OpCreate, Rule{
		BaseActions: []string{
			"route53:CreateHealthCheck",
			"route53:GetHealthCheck",
			"route53:ChangeTagsForResource",
		},
	})
	register("aws_route53_health_check", OpDelete, Rule{
		BaseActions: []string{
			"route53:DeleteHealthCheck",
			"route53:GetHealthCheck",
		},
	})
	register("aws_route53_health_check", OpUpdate, Rule{
		UpdateActions: []string{"route53:GetHealthCheck"},
		ConditionalActions: map[string][]string{
			"fqdn":              {"route53:UpdateHealthCheck"},
			"ip_address":        {"route53:UpdateHealthCheck"},
			"port":              {"route53:UpdateHealthCheck"},
			"failure_threshold": {"route53:UpdateHealthCheck"},
			"tags":              {"route53:ChangeTagsForResource"},
		},
	})
}

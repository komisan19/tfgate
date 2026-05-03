package iamrules

func init() {
	// aws_cloudfront_distribution
	register("aws_cloudfront_distribution", OpCreate, Rule{
		BaseActions: []string{
			"cloudfront:CreateDistribution",
			"cloudfront:GetDistribution",
			"cloudfront:TagResource",
		},
	})
	register("aws_cloudfront_distribution", OpDelete, Rule{
		BaseActions: []string{
			"cloudfront:DeleteDistribution",
			"cloudfront:GetDistribution",
			"cloudfront:UpdateDistribution",
		},
	})
	register("aws_cloudfront_distribution", OpUpdate, Rule{
		UpdateActions: []string{"cloudfront:GetDistribution"},
		ConditionalActions: map[string][]string{
			"origin":                    {"cloudfront:UpdateDistribution"},
			"default_cache_behavior":    {"cloudfront:UpdateDistribution"},
			"ordered_cache_behavior":    {"cloudfront:UpdateDistribution"},
			"price_class":               {"cloudfront:UpdateDistribution"},
			"aliases":                   {"cloudfront:UpdateDistribution"},
			"viewer_certificate":        {"cloudfront:UpdateDistribution"},
			"restrictions":              {"cloudfront:UpdateDistribution"},
			"enabled":                   {"cloudfront:UpdateDistribution"},
			"http_version":              {"cloudfront:UpdateDistribution"},
			"is_ipv6_enabled":           {"cloudfront:UpdateDistribution"},
			"comment":                   {"cloudfront:UpdateDistribution"},
			"web_acl_id":                {"cloudfront:UpdateDistribution"},
			"logging_config":            {"cloudfront:UpdateDistribution"},
			"custom_error_response":     {"cloudfront:UpdateDistribution"},
			"default_root_object":       {"cloudfront:UpdateDistribution"},
			"tags":                      {"cloudfront:TagResource", "cloudfront:UntagResource"},
		},
	})

	// aws_cloudfront_origin_access_identity
	register("aws_cloudfront_origin_access_identity", OpCreate, Rule{
		BaseActions: []string{
			"cloudfront:CreateCloudFrontOriginAccessIdentity",
			"cloudfront:GetCloudFrontOriginAccessIdentity",
		},
	})
	register("aws_cloudfront_origin_access_identity", OpDelete, Rule{
		BaseActions: []string{
			"cloudfront:DeleteCloudFrontOriginAccessIdentity",
			"cloudfront:GetCloudFrontOriginAccessIdentity",
		},
	})
	register("aws_cloudfront_origin_access_identity", OpUpdate, Rule{
		UpdateActions: []string{"cloudfront:GetCloudFrontOriginAccessIdentity"},
		ConditionalActions: map[string][]string{
			"comment": {"cloudfront:UpdateCloudFrontOriginAccessIdentity"},
		},
	})

	// aws_cloudfront_origin_access_control
	register("aws_cloudfront_origin_access_control", OpCreate, Rule{
		BaseActions: []string{
			"cloudfront:CreateOriginAccessControl",
			"cloudfront:GetOriginAccessControl",
		},
	})
	register("aws_cloudfront_origin_access_control", OpDelete, Rule{
		BaseActions: []string{
			"cloudfront:DeleteOriginAccessControl",
			"cloudfront:GetOriginAccessControl",
		},
	})
	register("aws_cloudfront_origin_access_control", OpUpdate, Rule{
		UpdateActions: []string{"cloudfront:GetOriginAccessControl"},
		ConditionalActions: map[string][]string{
			"name":                           {"cloudfront:UpdateOriginAccessControl"},
			"description":                    {"cloudfront:UpdateOriginAccessControl"},
			"origin_access_control_origin_type": {"cloudfront:UpdateOriginAccessControl"},
			"signing_behavior":               {"cloudfront:UpdateOriginAccessControl"},
			"signing_protocol":               {"cloudfront:UpdateOriginAccessControl"},
		},
	})

	// aws_cloudfront_cache_policy
	register("aws_cloudfront_cache_policy", OpCreate, Rule{
		BaseActions: []string{
			"cloudfront:CreateCachePolicy",
			"cloudfront:GetCachePolicy",
		},
	})
	register("aws_cloudfront_cache_policy", OpDelete, Rule{
		BaseActions: []string{
			"cloudfront:DeleteCachePolicy",
			"cloudfront:GetCachePolicy",
		},
	})
	register("aws_cloudfront_cache_policy", OpUpdate, Rule{
		UpdateActions: []string{"cloudfront:GetCachePolicy"},
		ConditionalActions: map[string][]string{
			"parameters_in_cache_key_and_forwarded_to_origin": {"cloudfront:UpdateCachePolicy"},
			"default_ttl":    {"cloudfront:UpdateCachePolicy"},
			"max_ttl":        {"cloudfront:UpdateCachePolicy"},
			"min_ttl":        {"cloudfront:UpdateCachePolicy"},
			"name":           {"cloudfront:UpdateCachePolicy"},
			"comment":        {"cloudfront:UpdateCachePolicy"},
		},
	})
}

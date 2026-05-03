package iamrules

func init() {
	// aws_cloudtrail
	register("aws_cloudtrail", OpCreate, Rule{
		BaseActions: []string{
			"cloudtrail:CreateTrail",
			"cloudtrail:StartLogging",
			"cloudtrail:GetTrail",
			"cloudtrail:AddTags",
		},
	})
	register("aws_cloudtrail", OpDelete, Rule{
		BaseActions: []string{
			"cloudtrail:DeleteTrail",
			"cloudtrail:StopLogging",
			"cloudtrail:GetTrail",
		},
	})
	register("aws_cloudtrail", OpUpdate, Rule{
		UpdateActions: []string{"cloudtrail:GetTrail"},
		ConditionalActions: map[string][]string{
			"enable_logging":                  {"cloudtrail:StartLogging", "cloudtrail:StopLogging"},
			"s3_bucket_name":                  {"cloudtrail:UpdateTrail"},
			"s3_key_prefix":                   {"cloudtrail:UpdateTrail"},
			"include_global_service_events":   {"cloudtrail:UpdateTrail"},
			"is_multi_region_trail":           {"cloudtrail:UpdateTrail"},
			"is_organization_trail":           {"cloudtrail:UpdateTrail"},
			"cloud_watch_logs_log_group_arn":  {"cloudtrail:UpdateTrail"},
			"cloud_watch_logs_role_arn":        {"cloudtrail:UpdateTrail"},
			"sns_topic_name":                  {"cloudtrail:UpdateTrail"},
			"enable_log_file_validation":      {"cloudtrail:UpdateTrail"},
			"kms_key_id":                      {"cloudtrail:UpdateTrail"},
			"event_selector":                  {"cloudtrail:PutEventSelectors"},
			"insight_selector":                {"cloudtrail:PutInsightSelectors"},
			"tags":                            {"cloudtrail:AddTags", "cloudtrail:RemoveTags"},
		},
	})
}

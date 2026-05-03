package iamrules

func init() {
	register("aws_sns_topic", OpCreate, Rule{
		BaseActions: []string{
			"sns:CreateTopic",
			"sns:GetTopicAttributes",
		},
	})
	register("aws_sns_topic", OpDelete, Rule{
		BaseActions: []string{
			"sns:DeleteTopic",
			"sns:GetTopicAttributes",
		},
	})
	register("aws_sns_topic", OpUpdate, Rule{
		UpdateActions: []string{"sns:GetTopicAttributes"},
		ConditionalActions: map[string][]string{
			"display_name":               {"sns:SetTopicAttributes"},
			"policy":                     {"sns:SetTopicAttributes"},
			"delivery_policy":            {"sns:SetTopicAttributes"},
			"kms_master_key_id":          {"sns:SetTopicAttributes"},
			"tracing_config":             {"sns:SetTopicAttributes"},
			"signature_version":          {"sns:SetTopicAttributes"},
			"fifo_topic":                 {"sns:SetTopicAttributes"},
			"content_based_deduplication": {"sns:SetTopicAttributes"},
			"tags":                       {"sns:TagResource", "sns:UntagResource"},
		},
	})

	register("aws_sns_topic_subscription", OpCreate, Rule{
		BaseActions: []string{
			"sns:Subscribe",
			"sns:GetSubscriptionAttributes",
		},
	})
	register("aws_sns_topic_subscription", OpDelete, Rule{
		BaseActions: []string{"sns:Unsubscribe"},
	})
	register("aws_sns_topic_subscription", OpUpdate, Rule{
		UpdateActions: []string{"sns:GetSubscriptionAttributes"},
		ConditionalActions: map[string][]string{
			"filter_policy":       {"sns:SetSubscriptionAttributes"},
			"filter_policy_scope": {"sns:SetSubscriptionAttributes"},
			"raw_message_delivery": {"sns:SetSubscriptionAttributes"},
			"redrive_policy":      {"sns:SetSubscriptionAttributes"},
			"delivery_policy":     {"sns:SetSubscriptionAttributes"},
		},
	})
}

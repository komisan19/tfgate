package iamrules

func init() {
	register("aws_sqs_queue", OpCreate, Rule{
		BaseActions: []string{
			"sqs:CreateQueue",
			"sqs:GetQueueUrl",
			"sqs:GetQueueAttributes",
		},
	})
	register("aws_sqs_queue", OpDelete, Rule{
		BaseActions: []string{
			"sqs:DeleteQueue",
			"sqs:GetQueueUrl",
		},
	})
	register("aws_sqs_queue", OpUpdate, Rule{
		UpdateActions: []string{"sqs:GetQueueAttributes"},
		ConditionalActions: map[string][]string{
			"visibility_timeout_seconds":        {"sqs:SetQueueAttributes"},
			"message_retention_seconds":         {"sqs:SetQueueAttributes"},
			"max_message_size":                  {"sqs:SetQueueAttributes"},
			"delay_seconds":                     {"sqs:SetQueueAttributes"},
			"receive_wait_time_seconds":         {"sqs:SetQueueAttributes"},
			"policy":                            {"sqs:SetQueueAttributes"},
			"redrive_policy":                    {"sqs:SetQueueAttributes"},
			"redrive_allow_policy":              {"sqs:SetQueueAttributes"},
			"kms_master_key_id":                 {"sqs:SetQueueAttributes"},
			"kms_data_key_reuse_period_seconds": {"sqs:SetQueueAttributes"},
			"sqs_managed_sse_enabled":           {"sqs:SetQueueAttributes"},
			"tags":                              {"sqs:TagQueue", "sqs:UntagQueue"},
		},
	})

	register("aws_sqs_queue_policy", OpCreate, Rule{
		BaseActions: []string{
			"sqs:SetQueueAttributes",
			"sqs:GetQueueAttributes",
		},
	})
	register("aws_sqs_queue_policy", OpDelete, Rule{
		BaseActions: []string{"sqs:SetQueueAttributes"},
	})
	register("aws_sqs_queue_policy", OpUpdate, Rule{
		UpdateActions: []string{"sqs:GetQueueAttributes"},
		ConditionalActions: map[string][]string{
			"policy": {"sqs:SetQueueAttributes"},
		},
	})
}

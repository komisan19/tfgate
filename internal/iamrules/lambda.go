package iamrules

func init() {
	register("aws_lambda_function", OpCreate, Rule{
		BaseActions: []string{
			"lambda:CreateFunction",
			"lambda:GetFunction",
			"lambda:GetFunctionConfiguration",
		},
	})
	register("aws_lambda_function", OpDelete, Rule{
		BaseActions: []string{
			"lambda:DeleteFunction",
			"lambda:GetFunction",
		},
	})
	register("aws_lambda_function", OpUpdate, Rule{
		UpdateActions: []string{
			"lambda:GetFunction",
			"lambda:GetFunctionConfiguration",
		},
		ConditionalActions: map[string][]string{
			"filename":          {"lambda:UpdateFunctionCode"},
			"s3_bucket":         {"lambda:UpdateFunctionCode"},
			"s3_key":            {"lambda:UpdateFunctionCode"},
			"s3_object_version": {"lambda:UpdateFunctionCode"},
			"image_uri":         {"lambda:UpdateFunctionCode"},
			"role":              {"lambda:UpdateFunctionConfiguration"},
			"handler":           {"lambda:UpdateFunctionConfiguration"},
			"runtime":           {"lambda:UpdateFunctionConfiguration"},
			"timeout":           {"lambda:UpdateFunctionConfiguration"},
			"memory_size":       {"lambda:UpdateFunctionConfiguration"},
			"environment":       {"lambda:UpdateFunctionConfiguration"},
			"description":       {"lambda:UpdateFunctionConfiguration"},
			"layers":            {"lambda:UpdateFunctionConfiguration"},
			"vpc_config":        {"lambda:UpdateFunctionConfiguration"},
			"dead_letter_config": {"lambda:UpdateFunctionConfiguration"},
			"tracing_config":    {"lambda:UpdateFunctionConfiguration"},
			"ephemeral_storage": {"lambda:UpdateFunctionConfiguration"},
			"snap_start":        {"lambda:UpdateFunctionConfiguration"},
			"reserved_concurrent_executions": {
				"lambda:PutFunctionConcurrency",
				"lambda:DeleteFunctionConcurrency",
			},
			"tags": {"lambda:TagResource", "lambda:UntagResource"},
		},
	})

	register("aws_lambda_event_source_mapping", OpCreate, Rule{
		BaseActions: []string{
			"lambda:CreateEventSourceMapping",
			"lambda:GetEventSourceMapping",
		},
	})
	register("aws_lambda_event_source_mapping", OpDelete, Rule{
		BaseActions: []string{
			"lambda:DeleteEventSourceMapping",
			"lambda:GetEventSourceMapping",
		},
	})
	register("aws_lambda_event_source_mapping", OpUpdate, Rule{
		UpdateActions: []string{"lambda:GetEventSourceMapping"},
		ConditionalActions: map[string][]string{
			"enabled":                                {"lambda:UpdateEventSourceMapping"},
			"batch_size":                             {"lambda:UpdateEventSourceMapping"},
			"function_name":                          {"lambda:UpdateEventSourceMapping"},
			"maximum_batching_window_in_seconds":     {"lambda:UpdateEventSourceMapping"},
			"bisect_batch_on_function_error":         {"lambda:UpdateEventSourceMapping"},
			"maximum_retry_attempts":                 {"lambda:UpdateEventSourceMapping"},
			"destination_config":                     {"lambda:UpdateEventSourceMapping"},
			"filter_criteria":                        {"lambda:UpdateEventSourceMapping"},
			"scaling_config":                         {"lambda:UpdateEventSourceMapping"},
			"function_response_types":                {"lambda:UpdateEventSourceMapping"},
		},
	})

	register("aws_lambda_permission", OpCreate, Rule{
		BaseActions: []string{
			"lambda:AddPermission",
			"lambda:GetPolicy",
		},
	})
	register("aws_lambda_permission", OpDelete, Rule{
		BaseActions: []string{
			"lambda:RemovePermission",
			"lambda:GetPolicy",
		},
	})
}

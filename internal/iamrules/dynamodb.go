package iamrules

func init() {
	// aws_dynamodb_table
	register("aws_dynamodb_table", OpCreate, Rule{
		BaseActions: []string{
			"dynamodb:CreateTable",
			"dynamodb:DescribeTable",
			"dynamodb:TagResource",
		},
	})
	register("aws_dynamodb_table", OpDelete, Rule{
		BaseActions: []string{
			"dynamodb:DeleteTable",
			"dynamodb:DescribeTable",
		},
	})
	register("aws_dynamodb_table", OpUpdate, Rule{
		UpdateActions: []string{"dynamodb:DescribeTable"},
		ConditionalActions: map[string][]string{
			"billing_mode":              {"dynamodb:UpdateTable"},
			"read_capacity":             {"dynamodb:UpdateTable"},
			"write_capacity":            {"dynamodb:UpdateTable"},
			"global_secondary_index":    {"dynamodb:UpdateTable"},
			"local_secondary_index":     {"dynamodb:UpdateTable"},
			"stream_enabled":            {"dynamodb:UpdateTable"},
			"stream_view_type":          {"dynamodb:UpdateTable"},
			"server_side_encryption":    {"dynamodb:UpdateTable"},
			"ttl":                       {"dynamodb:UpdateTimeToLive"},
			"point_in_time_recovery":    {"dynamodb:UpdateContinuousBackups"},
			"replica":                   {"dynamodb:UpdateTable"},
			"table_class":               {"dynamodb:UpdateTable"},
			"tags":                      {"dynamodb:TagResource", "dynamodb:UntagResource"},
		},
	})

	// aws_dynamodb_global_table
	register("aws_dynamodb_global_table", OpCreate, Rule{
		BaseActions: []string{
			"dynamodb:CreateGlobalTable",
			"dynamodb:DescribeGlobalTable",
		},
	})
	register("aws_dynamodb_global_table", OpDelete, Rule{
		BaseActions: []string{"dynamodb:DeleteTable"},
	})
	register("aws_dynamodb_global_table", OpUpdate, Rule{
		UpdateActions: []string{"dynamodb:DescribeGlobalTable"},
		ConditionalActions: map[string][]string{
			"replica": {"dynamodb:UpdateGlobalTable"},
		},
	})
}

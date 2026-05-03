package iamrules

func init() {
	// aws_cloudwatch_metric_alarm
	register("aws_cloudwatch_metric_alarm", OpCreate, Rule{
		BaseActions: []string{
			"cloudwatch:PutMetricAlarm",
			"cloudwatch:DescribeAlarms",
			"cloudwatch:TagResource",
		},
	})
	register("aws_cloudwatch_metric_alarm", OpDelete, Rule{
		BaseActions: []string{
			"cloudwatch:DeleteAlarms",
			"cloudwatch:DescribeAlarms",
		},
	})
	register("aws_cloudwatch_metric_alarm", OpUpdate, Rule{
		UpdateActions: []string{"cloudwatch:DescribeAlarms"},
		ConditionalActions: map[string][]string{
			"alarm_actions":            {"cloudwatch:PutMetricAlarm"},
			"ok_actions":               {"cloudwatch:PutMetricAlarm"},
			"insufficient_data_actions": {"cloudwatch:PutMetricAlarm"},
			"comparison_operator":      {"cloudwatch:PutMetricAlarm"},
			"evaluation_periods":       {"cloudwatch:PutMetricAlarm"},
			"metric_name":              {"cloudwatch:PutMetricAlarm"},
			"namespace":                {"cloudwatch:PutMetricAlarm"},
			"statistic":                {"cloudwatch:PutMetricAlarm"},
			"threshold":                {"cloudwatch:PutMetricAlarm"},
			"period":                   {"cloudwatch:PutMetricAlarm"},
			"metric_query":             {"cloudwatch:PutMetricAlarm"},
			"tags":                     {"cloudwatch:TagResource", "cloudwatch:UntagResource"},
		},
	})

	// aws_cloudwatch_log_group
	register("aws_cloudwatch_log_group", OpCreate, Rule{
		BaseActions: []string{
			"logs:CreateLogGroup",
			"logs:DescribeLogGroups",
			"logs:TagLogGroup",
		},
	})
	register("aws_cloudwatch_log_group", OpDelete, Rule{
		BaseActions: []string{
			"logs:DeleteLogGroup",
			"logs:DescribeLogGroups",
		},
	})
	register("aws_cloudwatch_log_group", OpUpdate, Rule{
		UpdateActions: []string{"logs:DescribeLogGroups"},
		ConditionalActions: map[string][]string{
			"retention_in_days": {"logs:PutRetentionPolicy"},
			"kms_key_id":        {"logs:AssociateKmsKey", "logs:DisassociateKmsKey"},
			"tags":              {"logs:TagLogGroup", "logs:UntagLogGroup"},
		},
	})

	// aws_cloudwatch_log_metric_filter
	register("aws_cloudwatch_log_metric_filter", OpCreate, Rule{
		BaseActions: []string{
			"logs:PutMetricFilter",
			"logs:DescribeMetricFilters",
		},
	})
	register("aws_cloudwatch_log_metric_filter", OpDelete, Rule{
		BaseActions: []string{
			"logs:DeleteMetricFilter",
			"logs:DescribeMetricFilters",
		},
	})
	register("aws_cloudwatch_log_metric_filter", OpUpdate, Rule{
		UpdateActions: []string{"logs:DescribeMetricFilters"},
		ConditionalActions: map[string][]string{
			"pattern":         {"logs:PutMetricFilter"},
			"metric_transformation": {"logs:PutMetricFilter"},
		},
	})

	// aws_cloudwatch_log_subscription_filter
	register("aws_cloudwatch_log_subscription_filter", OpCreate, Rule{
		BaseActions: []string{
			"logs:PutSubscriptionFilter",
			"logs:DescribeSubscriptionFilters",
		},
	})
	register("aws_cloudwatch_log_subscription_filter", OpDelete, Rule{
		BaseActions: []string{
			"logs:DeleteSubscriptionFilter",
			"logs:DescribeSubscriptionFilters",
		},
	})
	register("aws_cloudwatch_log_subscription_filter", OpUpdate, Rule{
		UpdateActions: []string{"logs:DescribeSubscriptionFilters"},
		ConditionalActions: map[string][]string{
			"filter_pattern":  {"logs:PutSubscriptionFilter"},
			"destination_arn": {"logs:PutSubscriptionFilter"},
		},
	})

	// aws_cloudwatch_dashboard
	register("aws_cloudwatch_dashboard", OpCreate, Rule{
		BaseActions: []string{
			"cloudwatch:PutDashboard",
			"cloudwatch:GetDashboard",
		},
	})
	register("aws_cloudwatch_dashboard", OpDelete, Rule{
		BaseActions: []string{
			"cloudwatch:DeleteDashboards",
			"cloudwatch:GetDashboard",
		},
	})
	register("aws_cloudwatch_dashboard", OpUpdate, Rule{
		UpdateActions: []string{"cloudwatch:GetDashboard"},
		ConditionalActions: map[string][]string{
			"dashboard_body": {"cloudwatch:PutDashboard"},
		},
	})

	// aws_cloudwatch_composite_alarm
	register("aws_cloudwatch_composite_alarm", OpCreate, Rule{
		BaseActions: []string{
			"cloudwatch:PutCompositeAlarm",
			"cloudwatch:DescribeAlarms",
			"cloudwatch:TagResource",
		},
	})
	register("aws_cloudwatch_composite_alarm", OpDelete, Rule{
		BaseActions: []string{
			"cloudwatch:DeleteAlarms",
			"cloudwatch:DescribeAlarms",
		},
	})
	register("aws_cloudwatch_composite_alarm", OpUpdate, Rule{
		UpdateActions: []string{"cloudwatch:DescribeAlarms"},
		ConditionalActions: map[string][]string{
			"alarm_rule":                {"cloudwatch:PutCompositeAlarm"},
			"alarm_actions":             {"cloudwatch:PutCompositeAlarm"},
			"ok_actions":                {"cloudwatch:PutCompositeAlarm"},
			"insufficient_data_actions": {"cloudwatch:PutCompositeAlarm"},
			"tags":                      {"cloudwatch:TagResource", "cloudwatch:UntagResource"},
		},
	})
}

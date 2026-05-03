package iamrules

func init() {
	// aws_codebuild_project
	register("aws_codebuild_project", OpCreate, Rule{
		BaseActions: []string{
			"codebuild:CreateProject",
			"codebuild:BatchGetProjects",
			"codebuild:TagResource",
		},
	})
	register("aws_codebuild_project", OpDelete, Rule{
		BaseActions: []string{
			"codebuild:DeleteProject",
			"codebuild:BatchGetProjects",
		},
	})
	register("aws_codebuild_project", OpUpdate, Rule{
		UpdateActions: []string{"codebuild:BatchGetProjects"},
		ConditionalActions: map[string][]string{
			"source":          {"codebuild:UpdateProject"},
			"environment":     {"codebuild:UpdateProject"},
			"artifacts":       {"codebuild:UpdateProject"},
			"secondary_artifacts": {"codebuild:UpdateProject"},
			"secondary_sources":   {"codebuild:UpdateProject"},
			"build_timeout":   {"codebuild:UpdateProject"},
			"queued_timeout":  {"codebuild:UpdateProject"},
			"cache":           {"codebuild:UpdateProject"},
			"vpc_config":      {"codebuild:UpdateProject"},
			"logs_config":     {"codebuild:UpdateProject"},
			"description":     {"codebuild:UpdateProject"},
			"service_role":    {"codebuild:UpdateProject"},
			"tags":            {"codebuild:TagResource", "codebuild:UntagResource"},
		},
	})

	// aws_codebuild_report_group
	register("aws_codebuild_report_group", OpCreate, Rule{
		BaseActions: []string{
			"codebuild:CreateReportGroup",
			"codebuild:BatchGetReportGroups",
			"codebuild:TagResource",
		},
	})
	register("aws_codebuild_report_group", OpDelete, Rule{
		BaseActions: []string{
			"codebuild:DeleteReportGroup",
			"codebuild:BatchGetReportGroups",
		},
	})
	register("aws_codebuild_report_group", OpUpdate, Rule{
		UpdateActions: []string{"codebuild:BatchGetReportGroups"},
		ConditionalActions: map[string][]string{
			"export_config": {"codebuild:UpdateReportGroup"},
			"tags":          {"codebuild:TagResource", "codebuild:UntagResource"},
		},
	})

	// aws_codebuild_webhook
	register("aws_codebuild_webhook", OpCreate, Rule{
		BaseActions: []string{
			"codebuild:CreateWebhook",
			"codebuild:BatchGetProjects",
		},
	})
	register("aws_codebuild_webhook", OpDelete, Rule{
		BaseActions: []string{
			"codebuild:DeleteWebhook",
			"codebuild:BatchGetProjects",
		},
	})
	register("aws_codebuild_webhook", OpUpdate, Rule{
		UpdateActions: []string{"codebuild:BatchGetProjects"},
		ConditionalActions: map[string][]string{
			"filter_group":         {"codebuild:UpdateWebhook"},
			"build_type":           {"codebuild:UpdateWebhook"},
			"branch_filter":        {"codebuild:UpdateWebhook"},
		},
	})
}

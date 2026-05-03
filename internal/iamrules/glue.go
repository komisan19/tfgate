package iamrules

func init() {
	// aws_glue_job
	register("aws_glue_job", OpCreate, Rule{
		BaseActions: []string{
			"glue:CreateJob",
			"glue:GetJob",
			"glue:TagResource",
		},
	})
	register("aws_glue_job", OpDelete, Rule{
		BaseActions: []string{
			"glue:DeleteJob",
			"glue:GetJob",
		},
	})
	register("aws_glue_job", OpUpdate, Rule{
		UpdateActions: []string{"glue:GetJob"},
		ConditionalActions: map[string][]string{
			"command":                 {"glue:UpdateJob"},
			"connections":             {"glue:UpdateJob"},
			"default_arguments":       {"glue:UpdateJob"},
			"non_overridable_arguments": {"glue:UpdateJob"},
			"max_capacity":            {"glue:UpdateJob"},
			"number_of_workers":       {"glue:UpdateJob"},
			"worker_type":             {"glue:UpdateJob"},
			"timeout":                 {"glue:UpdateJob"},
			"max_retries":             {"glue:UpdateJob"},
			"security_configuration":  {"glue:UpdateJob"},
			"glue_version":            {"glue:UpdateJob"},
			"description":             {"glue:UpdateJob"},
			"tags":                    {"glue:TagResource", "glue:UntagResource"},
		},
	})

	// aws_glue_database
	register("aws_glue_database", OpCreate, Rule{
		BaseActions: []string{
			"glue:CreateDatabase",
			"glue:GetDatabase",
		},
	})
	register("aws_glue_database", OpDelete, Rule{
		BaseActions: []string{
			"glue:DeleteDatabase",
			"glue:GetDatabase",
		},
	})
	register("aws_glue_database", OpUpdate, Rule{
		UpdateActions: []string{"glue:GetDatabase"},
		ConditionalActions: map[string][]string{
			"description":   {"glue:UpdateDatabase"},
			"location_uri":  {"glue:UpdateDatabase"},
			"parameters":    {"glue:UpdateDatabase"},
		},
	})

	// aws_glue_crawler
	register("aws_glue_crawler", OpCreate, Rule{
		BaseActions: []string{
			"glue:CreateCrawler",
			"glue:GetCrawler",
			"glue:TagResource",
		},
	})
	register("aws_glue_crawler", OpDelete, Rule{
		BaseActions: []string{
			"glue:DeleteCrawler",
			"glue:GetCrawler",
		},
	})
	register("aws_glue_crawler", OpUpdate, Rule{
		UpdateActions: []string{"glue:GetCrawler"},
		ConditionalActions: map[string][]string{
			"s3_target":               {"glue:UpdateCrawler"},
			"jdbc_target":             {"glue:UpdateCrawler"},
			"dynamodb_target":         {"glue:UpdateCrawler"},
			"catalog_target":          {"glue:UpdateCrawler"},
			"role":                    {"glue:UpdateCrawler"},
			"configuration":           {"glue:UpdateCrawler"},
			"classifiers":             {"glue:UpdateCrawler"},
			"schema_change_policy":    {"glue:UpdateCrawler"},
			"security_configuration":  {"glue:UpdateCrawler"},
			"description":             {"glue:UpdateCrawler"},
			"schedule":                {"glue:UpdateCrawlerSchedule"},
			"tags":                    {"glue:TagResource", "glue:UntagResource"},
		},
	})

	// aws_glue_catalog_table
	register("aws_glue_catalog_table", OpCreate, Rule{
		BaseActions: []string{
			"glue:CreateTable",
			"glue:GetTable",
		},
	})
	register("aws_glue_catalog_table", OpDelete, Rule{
		BaseActions: []string{
			"glue:DeleteTable",
			"glue:GetTable",
		},
	})
	register("aws_glue_catalog_table", OpUpdate, Rule{
		UpdateActions: []string{"glue:GetTable"},
		ConditionalActions: map[string][]string{
			"storage_descriptor": {"glue:UpdateTable"},
			"partition_keys":     {"glue:UpdateTable"},
			"parameters":         {"glue:UpdateTable"},
			"description":        {"glue:UpdateTable"},
		},
	})

	// aws_glue_trigger
	register("aws_glue_trigger", OpCreate, Rule{
		BaseActions: []string{
			"glue:CreateTrigger",
			"glue:GetTrigger",
			"glue:TagResource",
		},
	})
	register("aws_glue_trigger", OpDelete, Rule{
		BaseActions: []string{
			"glue:DeleteTrigger",
			"glue:GetTrigger",
		},
	})
	register("aws_glue_trigger", OpUpdate, Rule{
		UpdateActions: []string{"glue:GetTrigger"},
		ConditionalActions: map[string][]string{
			"actions":      {"glue:UpdateTrigger"},
			"predicate":    {"glue:UpdateTrigger"},
			"schedule":     {"glue:UpdateTrigger"},
			"description":  {"glue:UpdateTrigger"},
			"enabled":      {"glue:StartTrigger", "glue:StopTrigger"},
			"tags":         {"glue:TagResource", "glue:UntagResource"},
		},
	})
}

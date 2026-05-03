package iamrules

func init() {
	// aws_athena_workgroup
	register("aws_athena_workgroup", OpCreate, Rule{
		BaseActions: []string{
			"athena:CreateWorkGroup",
			"athena:GetWorkGroup",
			"athena:TagResource",
		},
	})
	register("aws_athena_workgroup", OpDelete, Rule{
		BaseActions: []string{
			"athena:DeleteWorkGroup",
			"athena:GetWorkGroup",
		},
	})
	register("aws_athena_workgroup", OpUpdate, Rule{
		UpdateActions: []string{"athena:GetWorkGroup"},
		ConditionalActions: map[string][]string{
			"configuration": {"athena:UpdateWorkGroup"},
			"description":   {"athena:UpdateWorkGroup"},
			"state":         {"athena:UpdateWorkGroup"},
			"tags":          {"athena:TagResource", "athena:UntagResource"},
		},
	})

	// aws_athena_database
	register("aws_athena_database", OpCreate, Rule{
		BaseActions: []string{
			"athena:StartQueryExecution",
			"athena:GetQueryExecution",
		},
	})
	register("aws_athena_database", OpDelete, Rule{
		BaseActions: []string{
			"athena:StartQueryExecution",
			"athena:GetQueryExecution",
		},
	})

	// aws_athena_named_query
	register("aws_athena_named_query", OpCreate, Rule{
		BaseActions: []string{
			"athena:CreateNamedQuery",
			"athena:GetNamedQuery",
		},
	})
	register("aws_athena_named_query", OpDelete, Rule{
		BaseActions: []string{
			"athena:DeleteNamedQuery",
			"athena:GetNamedQuery",
		},
	})

	// aws_athena_data_catalog
	register("aws_athena_data_catalog", OpCreate, Rule{
		BaseActions: []string{
			"athena:CreateDataCatalog",
			"athena:GetDataCatalog",
			"athena:TagResource",
		},
	})
	register("aws_athena_data_catalog", OpDelete, Rule{
		BaseActions: []string{
			"athena:DeleteDataCatalog",
			"athena:GetDataCatalog",
		},
	})
	register("aws_athena_data_catalog", OpUpdate, Rule{
		UpdateActions: []string{"athena:GetDataCatalog"},
		ConditionalActions: map[string][]string{
			"parameters":  {"athena:UpdateDataCatalog"},
			"description": {"athena:UpdateDataCatalog"},
			"tags":        {"athena:TagResource", "athena:UntagResource"},
		},
	})
}

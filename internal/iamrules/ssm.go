package iamrules

func init() {
	// aws_ssm_parameter
	register("aws_ssm_parameter", OpCreate, Rule{
		BaseActions: []string{
			"ssm:PutParameter",
			"ssm:GetParameters",
			"ssm:AddTagsToResource",
		},
	})
	register("aws_ssm_parameter", OpDelete, Rule{
		BaseActions: []string{
			"ssm:DeleteParameter",
			"ssm:GetParameters",
		},
	})
	register("aws_ssm_parameter", OpUpdate, Rule{
		UpdateActions: []string{"ssm:GetParameters"},
		ConditionalActions: map[string][]string{
			"value":       {"ssm:PutParameter"},
			"type":        {"ssm:PutParameter"},
			"description": {"ssm:PutParameter"},
			"tier":        {"ssm:PutParameter"},
			"key_id":      {"ssm:PutParameter"},
			"tags":        {"ssm:AddTagsToResource", "ssm:RemoveTagsFromResource"},
		},
	})

	// aws_ssm_document
	register("aws_ssm_document", OpCreate, Rule{
		BaseActions: []string{
			"ssm:CreateDocument",
			"ssm:DescribeDocument",
			"ssm:AddTagsToResource",
		},
	})
	register("aws_ssm_document", OpDelete, Rule{
		BaseActions: []string{
			"ssm:DeleteDocument",
			"ssm:DescribeDocument",
		},
	})
	register("aws_ssm_document", OpUpdate, Rule{
		UpdateActions: []string{"ssm:DescribeDocument"},
		ConditionalActions: map[string][]string{
			"content":     {"ssm:UpdateDocument"},
			"permissions": {"ssm:ModifyDocumentPermission"},
			"tags":        {"ssm:AddTagsToResource", "ssm:RemoveTagsFromResource"},
		},
	})

	// aws_ssm_association
	register("aws_ssm_association", OpCreate, Rule{
		BaseActions: []string{
			"ssm:CreateAssociation",
			"ssm:DescribeAssociation",
		},
	})
	register("aws_ssm_association", OpDelete, Rule{
		BaseActions: []string{
			"ssm:DeleteAssociation",
			"ssm:DescribeAssociation",
		},
	})
	register("aws_ssm_association", OpUpdate, Rule{
		UpdateActions: []string{"ssm:DescribeAssociation"},
		ConditionalActions: map[string][]string{
			"parameters":          {"ssm:UpdateAssociation"},
			"schedule_expression": {"ssm:UpdateAssociation"},
			"targets":             {"ssm:UpdateAssociation"},
			"output_location":     {"ssm:UpdateAssociation"},
		},
	})

	// aws_ssm_patch_baseline
	register("aws_ssm_patch_baseline", OpCreate, Rule{
		BaseActions: []string{
			"ssm:CreatePatchBaseline",
			"ssm:GetPatchBaseline",
			"ssm:AddTagsToResource",
		},
	})
	register("aws_ssm_patch_baseline", OpDelete, Rule{
		BaseActions: []string{
			"ssm:DeletePatchBaseline",
			"ssm:GetPatchBaseline",
		},
	})
	register("aws_ssm_patch_baseline", OpUpdate, Rule{
		UpdateActions: []string{"ssm:GetPatchBaseline"},
		ConditionalActions: map[string][]string{
			"approval_rule":         {"ssm:UpdatePatchBaseline"},
			"approved_patches":      {"ssm:UpdatePatchBaseline"},
			"rejected_patches":      {"ssm:UpdatePatchBaseline"},
			"global_filter":         {"ssm:UpdatePatchBaseline"},
			"description":           {"ssm:UpdatePatchBaseline"},
			"tags":                  {"ssm:AddTagsToResource", "ssm:RemoveTagsFromResource"},
		},
	})
}

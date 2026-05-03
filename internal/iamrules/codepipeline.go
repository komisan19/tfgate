package iamrules

func init() {
	// aws_codepipeline
	register("aws_codepipeline", OpCreate, Rule{
		BaseActions: []string{
			"codepipeline:CreatePipeline",
			"codepipeline:GetPipeline",
			"codepipeline:TagResource",
		},
	})
	register("aws_codepipeline", OpDelete, Rule{
		BaseActions: []string{
			"codepipeline:DeletePipeline",
			"codepipeline:GetPipeline",
		},
	})
	register("aws_codepipeline", OpUpdate, Rule{
		UpdateActions: []string{"codepipeline:GetPipeline"},
		ConditionalActions: map[string][]string{
			"stage":     {"codepipeline:UpdatePipeline"},
			"artifact_store": {"codepipeline:UpdatePipeline"},
			"role_arn":  {"codepipeline:UpdatePipeline"},
			"tags":      {"codepipeline:TagResource", "codepipeline:UntagResource"},
		},
	})

	// aws_codepipeline_webhook
	register("aws_codepipeline_webhook", OpCreate, Rule{
		BaseActions: []string{
			"codepipeline:PutWebhook",
			"codepipeline:GetWebhook",
			"codepipeline:RegisterWebhookWithThirdParty",
			"codepipeline:TagResource",
		},
	})
	register("aws_codepipeline_webhook", OpDelete, Rule{
		BaseActions: []string{
			"codepipeline:DeregisterWebhookWithThirdParty",
			"codepipeline:DeleteWebhook",
			"codepipeline:GetWebhook",
		},
	})
	register("aws_codepipeline_webhook", OpUpdate, Rule{
		UpdateActions: []string{"codepipeline:GetWebhook"},
		ConditionalActions: map[string][]string{
			"filter":         {"codepipeline:PutWebhook"},
			"authentication": {"codepipeline:PutWebhook"},
			"target_action":  {"codepipeline:PutWebhook"},
			"tags":           {"codepipeline:TagResource", "codepipeline:UntagResource"},
		},
	})

	// aws_codecommit_repository
	register("aws_codecommit_repository", OpCreate, Rule{
		BaseActions: []string{
			"codecommit:CreateRepository",
			"codecommit:GetRepository",
			"codecommit:TagResource",
		},
	})
	register("aws_codecommit_repository", OpDelete, Rule{
		BaseActions: []string{
			"codecommit:DeleteRepository",
			"codecommit:GetRepository",
		},
	})
	register("aws_codecommit_repository", OpUpdate, Rule{
		UpdateActions: []string{"codecommit:GetRepository"},
		ConditionalActions: map[string][]string{
			"description": {"codecommit:UpdateRepositoryDescription"},
			"default_branch": {"codecommit:UpdateDefaultBranch"},
			"tags":        {"codecommit:TagResource", "codecommit:UntagResource"},
		},
	})
}

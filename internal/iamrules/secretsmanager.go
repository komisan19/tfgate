package iamrules

func init() {
	// aws_secretsmanager_secret
	register("aws_secretsmanager_secret", OpCreate, Rule{
		BaseActions: []string{
			"secretsmanager:CreateSecret",
			"secretsmanager:DescribeSecret",
			"secretsmanager:TagResource",
		},
	})
	register("aws_secretsmanager_secret", OpDelete, Rule{
		BaseActions: []string{
			"secretsmanager:DeleteSecret",
			"secretsmanager:DescribeSecret",
		},
	})
	register("aws_secretsmanager_secret", OpUpdate, Rule{
		UpdateActions: []string{"secretsmanager:DescribeSecret"},
		ConditionalActions: map[string][]string{
			"name":                    {"secretsmanager:UpdateSecret"},
			"description":             {"secretsmanager:UpdateSecret"},
			"kms_key_id":              {"secretsmanager:UpdateSecret"},
			"rotation_lambda_arn":     {"secretsmanager:RotateSecret"},
			"rotation_rules":          {"secretsmanager:RotateSecret"},
			"policy":                  {"secretsmanager:PutResourcePolicy"},
			"recovery_window_in_days": {"secretsmanager:DeleteSecret"},
			"tags":                    {"secretsmanager:TagResource", "secretsmanager:UntagResource"},
		},
	})

	// aws_secretsmanager_secret_version
	register("aws_secretsmanager_secret_version", OpCreate, Rule{
		BaseActions: []string{
			"secretsmanager:PutSecretValue",
			"secretsmanager:GetSecretValue",
		},
	})
	register("aws_secretsmanager_secret_version", OpDelete, Rule{
		BaseActions: []string{
			"secretsmanager:UpdateSecret",
			"secretsmanager:GetSecretValue",
		},
	})
	register("aws_secretsmanager_secret_version", OpUpdate, Rule{
		UpdateActions: []string{"secretsmanager:GetSecretValue"},
		ConditionalActions: map[string][]string{
			"secret_string": {"secretsmanager:PutSecretValue"},
			"secret_binary": {"secretsmanager:PutSecretValue"},
		},
	})

	// aws_secretsmanager_secret_rotation
	register("aws_secretsmanager_secret_rotation", OpCreate, Rule{
		BaseActions: []string{
			"secretsmanager:RotateSecret",
			"secretsmanager:DescribeSecret",
		},
	})
	register("aws_secretsmanager_secret_rotation", OpDelete, Rule{
		BaseActions: []string{
			"secretsmanager:CancelRotateSecret",
			"secretsmanager:DescribeSecret",
		},
	})
	register("aws_secretsmanager_secret_rotation", OpUpdate, Rule{
		UpdateActions: []string{"secretsmanager:DescribeSecret"},
		ConditionalActions: map[string][]string{
			"rotation_lambda_arn": {"secretsmanager:RotateSecret"},
			"rotation_rules":      {"secretsmanager:RotateSecret"},
		},
	})

	// aws_secretsmanager_secret_policy
	register("aws_secretsmanager_secret_policy", OpCreate, Rule{
		BaseActions: []string{
			"secretsmanager:PutResourcePolicy",
			"secretsmanager:GetResourcePolicy",
		},
	})
	register("aws_secretsmanager_secret_policy", OpDelete, Rule{
		BaseActions: []string{
			"secretsmanager:DeleteResourcePolicy",
			"secretsmanager:GetResourcePolicy",
		},
	})
	register("aws_secretsmanager_secret_policy", OpUpdate, Rule{
		UpdateActions: []string{"secretsmanager:GetResourcePolicy"},
		ConditionalActions: map[string][]string{
			"policy": {"secretsmanager:PutResourcePolicy"},
		},
	})
}

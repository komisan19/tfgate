package iamrules

func init() {
	register("aws_iam_role", OpCreate, Rule{
		BaseActions: []string{
			"iam:CreateRole",
			"iam:GetRole",
			"iam:TagRole",
		},
	})

	register("aws_iam_role", OpDelete, Rule{
		BaseActions: []string{
			"iam:DeleteRole",
			"iam:GetRole",
			"iam:ListAttachedRolePolicies",
			"iam:ListRolePolicies",
			"iam:ListInstanceProfilesForRole",
		},
	})

	register("aws_iam_role", OpUpdate, Rule{
		BaseActions: []string{
			"iam:GetRole",
		},
		ConditionalActions: map[string][]string{
			// キー名 = plan.json の After/Before に出てくる属性名
			// 値     = その属性が変わったとき追加で必要な action
			"name":                 {"iam:UpdateRole"},
			"description":          {"iam:UpdateRole"},
			"max_session_duration": {"iam:UpdateRole"},
			"assume_role_policy":   {"iam:UpdateAssumeRolePolicy"},
			"tags":                 {"iam:TagRole", "iam:UntagRole"},
			"inline_policy":        {"iam:PutRolePolicy", "iam:DeleteRolePolicy"},
			"managed_policy_arns":  {"iam:AttachRolePolicy", "iam:DetachRolePolicy"},
		},
	})
}

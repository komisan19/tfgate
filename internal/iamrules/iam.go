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
}

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
		UpdateActions: []string{
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

	register("aws_iam_policy", OpCreate, Rule{
		BaseActions: []string{"iam:CreatePolicy", "iam:GetPolicy"},
	})
	register("aws_iam_policy", OpDelete, Rule{
		BaseActions: []string{"iam:DeletePolicy", "iam:GetPolicy", "iam:ListPolicyVersions"},
	})
	register("aws_iam_policy", OpUpdate, Rule{
		UpdateActions: []string{"iam:GetPolicy"},
		ConditionalActions: map[string][]string{
			"policy": {"iam:CreatePolicyVersion", "iam:DeletePolicyVersion"},
			"tags":   {"iam:TagPolicy", "iam:UntagPolicy"},
		},
	})

	register("aws_iam_user", OpCreate, Rule{
		BaseActions: []string{"iam:CreateUser", "iam:GetUser"},
	})
	register("aws_iam_user", OpDelete, Rule{
		BaseActions: []string{
			"iam:DeleteUser", "iam:GetUser",
			"iam:ListAttachedUserPolicies", "iam:ListUserPolicies", "iam:ListGroupsForUser",
		},
	})
	register("aws_iam_user", OpUpdate, Rule{
		UpdateActions: []string{"iam:GetUser"},
		ConditionalActions: map[string][]string{
			"name": {"iam:UpdateUser"},
			"path": {"iam:UpdateUser"},
			"tags": {"iam:TagUser", "iam:UntagUser"},
		},
	})

	register("aws_iam_group", OpCreate, Rule{
		BaseActions: []string{"iam:CreateGroup", "iam:GetGroup"},
	})
	register("aws_iam_group", OpDelete, Rule{
		BaseActions: []string{
			"iam:DeleteGroup", "iam:GetGroup",
			"iam:ListAttachedGroupPolicies", "iam:ListGroupPolicies",
		},
	})
	register("aws_iam_group", OpUpdate, Rule{
		UpdateActions: []string{"iam:GetGroup"},
		ConditionalActions: map[string][]string{
			"name": {"iam:UpdateGroup"},
			"path": {"iam:UpdateGroup"},
		},
	})

	register("aws_iam_instance_profile", OpCreate, Rule{
		BaseActions: []string{
			"iam:CreateInstanceProfile", "iam:GetInstanceProfile", "iam:AddRoleToInstanceProfile",
		},
	})
	register("aws_iam_instance_profile", OpDelete, Rule{
		BaseActions: []string{
			"iam:DeleteInstanceProfile", "iam:GetInstanceProfile", "iam:RemoveRoleFromInstanceProfile",
		},
	})
	register("aws_iam_instance_profile", OpUpdate, Rule{
		UpdateActions: []string{"iam:GetInstanceProfile"},
		ConditionalActions: map[string][]string{
			"role": {"iam:AddRoleToInstanceProfile", "iam:RemoveRoleFromInstanceProfile"},
			"tags": {"iam:TagInstanceProfile", "iam:UntagInstanceProfile"},
		},
	})

	register("aws_iam_role_policy", OpCreate, Rule{
		BaseActions: []string{"iam:PutRolePolicy", "iam:GetRolePolicy"},
	})
	register("aws_iam_role_policy", OpDelete, Rule{
		BaseActions: []string{"iam:DeleteRolePolicy", "iam:GetRolePolicy"},
	})
	register("aws_iam_role_policy", OpUpdate, Rule{
		UpdateActions: []string{"iam:GetRolePolicy"},
		ConditionalActions: map[string][]string{
			"policy": {"iam:PutRolePolicy"},
		},
	})

	register("aws_iam_role_policy_attachment", OpCreate, Rule{
		BaseActions: []string{"iam:AttachRolePolicy", "iam:GetPolicy"},
	})
	register("aws_iam_role_policy_attachment", OpDelete, Rule{
		BaseActions: []string{"iam:DetachRolePolicy"},
	})

	register("aws_iam_user_policy", OpCreate, Rule{
		BaseActions: []string{"iam:PutUserPolicy", "iam:GetUserPolicy"},
	})
	register("aws_iam_user_policy", OpDelete, Rule{
		BaseActions: []string{"iam:DeleteUserPolicy", "iam:GetUserPolicy"},
	})
	register("aws_iam_user_policy", OpUpdate, Rule{
		UpdateActions: []string{"iam:GetUserPolicy"},
		ConditionalActions: map[string][]string{
			"policy": {"iam:PutUserPolicy"},
		},
	})

	register("aws_iam_user_policy_attachment", OpCreate, Rule{
		BaseActions: []string{"iam:AttachUserPolicy", "iam:GetPolicy"},
	})
	register("aws_iam_user_policy_attachment", OpDelete, Rule{
		BaseActions: []string{"iam:DetachUserPolicy"},
	})

	register("aws_iam_access_key", OpCreate, Rule{
		BaseActions: []string{"iam:CreateAccessKey"},
	})
	register("aws_iam_access_key", OpDelete, Rule{
		BaseActions: []string{"iam:DeleteAccessKey"},
	})
	register("aws_iam_access_key", OpUpdate, Rule{
		UpdateActions: []string{},
		ConditionalActions: map[string][]string{
			"status": {"iam:UpdateAccessKey"},
		},
	})

	register("aws_iam_openid_connect_provider", OpCreate, Rule{
		BaseActions: []string{"iam:CreateOpenIDConnectProvider", "iam:GetOpenIDConnectProvider"},
	})
	register("aws_iam_openid_connect_provider", OpDelete, Rule{
		BaseActions: []string{"iam:DeleteOpenIDConnectProvider", "iam:GetOpenIDConnectProvider"},
	})
	register("aws_iam_openid_connect_provider", OpUpdate, Rule{
		UpdateActions: []string{"iam:GetOpenIDConnectProvider"},
		ConditionalActions: map[string][]string{
			"thumbprint_list": {"iam:UpdateOpenIDConnectProviderThumbprint"},
			"client_id_list":  {"iam:AddClientIDToOpenIDConnectProvider", "iam:RemoveClientIDFromOpenIDConnectProvider"},
			"tags":            {"iam:TagOpenIDConnectProvider", "iam:UntagOpenIDConnectProvider"},
		},
	})
}

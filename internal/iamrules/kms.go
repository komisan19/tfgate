package iamrules

func init() {
	// aws_kms_key
	register("aws_kms_key", OpCreate, Rule{
		BaseActions: []string{
			"kms:CreateKey",
			"kms:DescribeKey",
			"kms:GetKeyPolicy",
			"kms:TagResource",
		},
	})
	register("aws_kms_key", OpDelete, Rule{
		BaseActions: []string{
			"kms:ScheduleKeyDeletion",
			"kms:DescribeKey",
		},
	})
	register("aws_kms_key", OpUpdate, Rule{
		UpdateActions: []string{"kms:DescribeKey"},
		ConditionalActions: map[string][]string{
			"description":             {"kms:UpdateKeyDescription"},
			"policy":                  {"kms:PutKeyPolicy"},
			"enable_key_rotation":     {"kms:EnableKeyRotation", "kms:DisableKeyRotation"},
			"is_enabled":              {"kms:EnableKey", "kms:DisableKey"},
			"deletion_window_in_days": {"kms:ScheduleKeyDeletion"},
			"tags":                    {"kms:TagResource", "kms:UntagResource"},
		},
	})

	// aws_kms_alias
	register("aws_kms_alias", OpCreate, Rule{
		BaseActions: []string{
			"kms:CreateAlias",
			"kms:ListAliases",
		},
	})
	register("aws_kms_alias", OpDelete, Rule{
		BaseActions: []string{
			"kms:DeleteAlias",
			"kms:ListAliases",
		},
	})
	register("aws_kms_alias", OpUpdate, Rule{
		UpdateActions: []string{"kms:ListAliases"},
		ConditionalActions: map[string][]string{
			"target_key_id": {"kms:UpdateAlias"},
		},
	})

	// aws_kms_grant
	register("aws_kms_grant", OpCreate, Rule{
		BaseActions: []string{
			"kms:CreateGrant",
			"kms:ListGrants",
		},
	})
	register("aws_kms_grant", OpDelete, Rule{
		BaseActions: []string{
			"kms:RevokeGrant",
			"kms:ListGrants",
		},
	})

	// aws_kms_key_policy
	register("aws_kms_key_policy", OpCreate, Rule{
		BaseActions: []string{
			"kms:PutKeyPolicy",
			"kms:GetKeyPolicy",
		},
	})
	register("aws_kms_key_policy", OpDelete, Rule{
		BaseActions: []string{
			"kms:PutKeyPolicy",
			"kms:GetKeyPolicy",
		},
	})
	register("aws_kms_key_policy", OpUpdate, Rule{
		UpdateActions: []string{"kms:GetKeyPolicy"},
		ConditionalActions: map[string][]string{
			"policy": {"kms:PutKeyPolicy"},
		},
	})
}

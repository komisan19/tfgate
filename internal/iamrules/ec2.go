package iamrules

func init() {
	register("aws_instance", OpCreate, Rule{
		BaseActions: []string{
			"ec2:RunInstances",
			"ec2:DescribeInstances",
			"ec2:CreateTags",
			"ec2:DescribeInstanceTypes",
			"ec2:DescribeImages",
			"ec2:DescribeVpcs",
			"ec2:DescribeSubnets",
			"ec2:DescribeSecurityGroups",
		},
	})

	register("aws_instance", OpDelete, Rule{
		BaseActions: []string{
			"ec2:TerminateInstances",
			"ec2:DescribeInstances",
		},
	})

	register("aws_instance", OpUpdate, Rule{
		UpdateActions: []string{"ec2:DescribeInstances"},
		ConditionalActions: map[string][]string{
			"instance_type":        {"ec2:ModifyInstanceAttribute", "ec2:StopInstances", "ec2:StartInstances"},
			"tags":                 {"ec2:CreateTags", "ec2:DeleteTags"},
			"security_groups":      {"ec2:ModifyInstanceAttribute"},
			"ebs_optimized":        {"ec2:ModifyInstanceAttribute"},
			"user_data":            {"ec2:ModifyInstanceAttribute", "ec2:StopInstances", "ec2:StartInstances"},
			"iam_instance_profile": {"ec2:AssociateIamInstanceProfile", "ec2:DisassociateIamInstanceProfile"},
		},
	})
}

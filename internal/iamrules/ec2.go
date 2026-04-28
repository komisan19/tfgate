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
}

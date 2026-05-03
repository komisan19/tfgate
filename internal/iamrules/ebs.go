package iamrules

func init() {
	// aws_ebs_volume
	register("aws_ebs_volume", OpCreate, Rule{
		BaseActions: []string{
			"ec2:CreateVolume",
			"ec2:DescribeVolumes",
			"ec2:CreateTags",
		},
	})
	register("aws_ebs_volume", OpDelete, Rule{
		BaseActions: []string{
			"ec2:DeleteVolume",
			"ec2:DescribeVolumes",
		},
	})
	register("aws_ebs_volume", OpUpdate, Rule{
		UpdateActions: []string{"ec2:DescribeVolumes"},
		ConditionalActions: map[string][]string{
			"size":       {"ec2:ModifyVolume"},
			"type":       {"ec2:ModifyVolume"},
			"iops":       {"ec2:ModifyVolume"},
			"throughput": {"ec2:ModifyVolume"},
			"tags":       {"ec2:CreateTags", "ec2:DeleteTags"},
		},
	})

	// aws_volume_attachment
	register("aws_volume_attachment", OpCreate, Rule{
		BaseActions: []string{
			"ec2:AttachVolume",
			"ec2:DescribeVolumes",
			"ec2:DescribeInstances",
		},
	})
	register("aws_volume_attachment", OpDelete, Rule{
		BaseActions: []string{
			"ec2:DetachVolume",
			"ec2:DescribeVolumes",
		},
	})

	// aws_ebs_snapshot
	register("aws_ebs_snapshot", OpCreate, Rule{
		BaseActions: []string{
			"ec2:CreateSnapshot",
			"ec2:DescribeSnapshots",
			"ec2:CreateTags",
		},
	})
	register("aws_ebs_snapshot", OpDelete, Rule{
		BaseActions: []string{
			"ec2:DeleteSnapshot",
			"ec2:DescribeSnapshots",
		},
	})
	register("aws_ebs_snapshot", OpUpdate, Rule{
		UpdateActions: []string{"ec2:DescribeSnapshots"},
		ConditionalActions: map[string][]string{
			"tags": {"ec2:CreateTags", "ec2:DeleteTags"},
		},
	})
}

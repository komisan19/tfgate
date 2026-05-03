package iamrules

func init() {
	// aws_efs_file_system
	register("aws_efs_file_system", OpCreate, Rule{
		BaseActions: []string{
			"elasticfilesystem:CreateFileSystem",
			"elasticfilesystem:DescribeFileSystems",
			"elasticfilesystem:TagResource",
		},
	})
	register("aws_efs_file_system", OpDelete, Rule{
		BaseActions: []string{
			"elasticfilesystem:DeleteFileSystem",
			"elasticfilesystem:DescribeFileSystems",
		},
	})
	register("aws_efs_file_system", OpUpdate, Rule{
		UpdateActions: []string{"elasticfilesystem:DescribeFileSystems"},
		ConditionalActions: map[string][]string{
			"throughput_mode":                {"elasticfilesystem:UpdateFileSystem"},
			"provisioned_throughput_in_mibps": {"elasticfilesystem:UpdateFileSystem"},
			"lifecycle_policy":               {"elasticfilesystem:PutLifecycleConfiguration"},
			"tags":                           {"elasticfilesystem:TagResource", "elasticfilesystem:UntagResource"},
		},
	})

	// aws_efs_mount_target
	register("aws_efs_mount_target", OpCreate, Rule{
		BaseActions: []string{
			"elasticfilesystem:CreateMountTarget",
			"elasticfilesystem:DescribeMountTargets",
		},
	})
	register("aws_efs_mount_target", OpDelete, Rule{
		BaseActions: []string{
			"elasticfilesystem:DeleteMountTarget",
			"elasticfilesystem:DescribeMountTargets",
		},
	})
	register("aws_efs_mount_target", OpUpdate, Rule{
		UpdateActions: []string{"elasticfilesystem:DescribeMountTargets"},
		ConditionalActions: map[string][]string{
			"security_groups": {"elasticfilesystem:ModifyMountTargetSecurityGroups"},
		},
	})

	// aws_efs_access_point
	register("aws_efs_access_point", OpCreate, Rule{
		BaseActions: []string{
			"elasticfilesystem:CreateAccessPoint",
			"elasticfilesystem:DescribeAccessPoints",
			"elasticfilesystem:TagResource",
		},
	})
	register("aws_efs_access_point", OpDelete, Rule{
		BaseActions: []string{
			"elasticfilesystem:DeleteAccessPoint",
			"elasticfilesystem:DescribeAccessPoints",
		},
	})
	register("aws_efs_access_point", OpUpdate, Rule{
		UpdateActions: []string{"elasticfilesystem:DescribeAccessPoints"},
		ConditionalActions: map[string][]string{
			"tags": {"elasticfilesystem:TagResource", "elasticfilesystem:UntagResource"},
		},
	})

	// aws_efs_file_system_policy
	register("aws_efs_file_system_policy", OpCreate, Rule{
		BaseActions: []string{
			"elasticfilesystem:PutFileSystemPolicy",
			"elasticfilesystem:DescribeFileSystemPolicy",
		},
	})
	register("aws_efs_file_system_policy", OpDelete, Rule{
		BaseActions: []string{
			"elasticfilesystem:DeleteFileSystemPolicy",
			"elasticfilesystem:DescribeFileSystemPolicy",
		},
	})
	register("aws_efs_file_system_policy", OpUpdate, Rule{
		UpdateActions: []string{"elasticfilesystem:DescribeFileSystemPolicy"},
		ConditionalActions: map[string][]string{
			"policy": {"elasticfilesystem:PutFileSystemPolicy"},
		},
	})
}

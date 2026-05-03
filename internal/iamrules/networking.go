package iamrules

func init() {
	// aws_vpc
	register("aws_vpc", OpCreate, Rule{
		BaseActions: []string{
			"ec2:CreateVpc",
			"ec2:DescribeVpcs",
			"ec2:ModifyVpcAttribute",
		},
	})
	register("aws_vpc", OpDelete, Rule{
		BaseActions: []string{
			"ec2:DeleteVpc",
			"ec2:DescribeVpcs",
		},
	})
	register("aws_vpc", OpUpdate, Rule{
		UpdateActions: []string{"ec2:DescribeVpcs"},
		ConditionalActions: map[string][]string{
			"enable_dns_hostnames": {"ec2:ModifyVpcAttribute"},
			"enable_dns_support":   {"ec2:ModifyVpcAttribute"},
			"instance_tenancy":     {"ec2:ModifyVpcTenancy"},
			"tags":                 {"ec2:CreateTags", "ec2:DeleteTags"},
		},
	})

	// aws_subnet
	register("aws_subnet", OpCreate, Rule{
		BaseActions: []string{
			"ec2:CreateSubnet",
			"ec2:DescribeSubnets",
			"ec2:ModifySubnetAttribute",
		},
	})
	register("aws_subnet", OpDelete, Rule{
		BaseActions: []string{
			"ec2:DeleteSubnet",
			"ec2:DescribeSubnets",
		},
	})
	register("aws_subnet", OpUpdate, Rule{
		UpdateActions: []string{"ec2:DescribeSubnets"},
		ConditionalActions: map[string][]string{
			"map_public_ip_on_launch":         {"ec2:ModifySubnetAttribute"},
			"map_customer_owned_ip_on_launch": {"ec2:ModifySubnetAttribute"},
			"customer_owned_ipv4_pool":        {"ec2:ModifySubnetAttribute"},
			"enable_dns64":                    {"ec2:ModifySubnetAttribute"},
			"tags":                            {"ec2:CreateTags", "ec2:DeleteTags"},
		},
	})

	// aws_security_group
	register("aws_security_group", OpCreate, Rule{
		BaseActions: []string{
			"ec2:CreateSecurityGroup",
			"ec2:DescribeSecurityGroups",
			"ec2:AuthorizeSecurityGroupIngress",
			"ec2:AuthorizeSecurityGroupEgress",
		},
	})
	register("aws_security_group", OpDelete, Rule{
		BaseActions: []string{
			"ec2:DeleteSecurityGroup",
			"ec2:DescribeSecurityGroups",
			"ec2:RevokeSecurityGroupIngress",
			"ec2:RevokeSecurityGroupEgress",
		},
	})
	register("aws_security_group", OpUpdate, Rule{
		UpdateActions: []string{"ec2:DescribeSecurityGroups"},
		ConditionalActions: map[string][]string{
			"ingress": {"ec2:AuthorizeSecurityGroupIngress", "ec2:RevokeSecurityGroupIngress"},
			"egress":  {"ec2:AuthorizeSecurityGroupEgress", "ec2:RevokeSecurityGroupEgress"},
			"tags":    {"ec2:CreateTags", "ec2:DeleteTags"},
		},
	})

	// aws_security_group_rule
	register("aws_security_group_rule", OpCreate, Rule{
		BaseActions: []string{
			"ec2:AuthorizeSecurityGroupIngress",
			"ec2:AuthorizeSecurityGroupEgress",
			"ec2:DescribeSecurityGroups",
		},
	})
	register("aws_security_group_rule", OpDelete, Rule{
		BaseActions: []string{
			"ec2:RevokeSecurityGroupIngress",
			"ec2:RevokeSecurityGroupEgress",
			"ec2:DescribeSecurityGroups",
		},
	})

	// aws_internet_gateway
	register("aws_internet_gateway", OpCreate, Rule{
		BaseActions: []string{
			"ec2:CreateInternetGateway",
			"ec2:AttachInternetGateway",
			"ec2:DescribeInternetGateways",
		},
	})
	register("aws_internet_gateway", OpDelete, Rule{
		BaseActions: []string{
			"ec2:DeleteInternetGateway",
			"ec2:DetachInternetGateway",
			"ec2:DescribeInternetGateways",
		},
	})
	register("aws_internet_gateway", OpUpdate, Rule{
		UpdateActions: []string{"ec2:DescribeInternetGateways"},
		ConditionalActions: map[string][]string{
			"vpc_id": {"ec2:AttachInternetGateway", "ec2:DetachInternetGateway"},
			"tags":   {"ec2:CreateTags", "ec2:DeleteTags"},
		},
	})

	// aws_nat_gateway
	register("aws_nat_gateway", OpCreate, Rule{
		BaseActions: []string{
			"ec2:CreateNatGateway",
			"ec2:DescribeNatGateways",
			"ec2:AllocateAddress",
		},
	})
	register("aws_nat_gateway", OpDelete, Rule{
		BaseActions: []string{
			"ec2:DeleteNatGateway",
			"ec2:DescribeNatGateways",
		},
	})
	register("aws_nat_gateway", OpUpdate, Rule{
		UpdateActions: []string{"ec2:DescribeNatGateways"},
		ConditionalActions: map[string][]string{
			"tags": {"ec2:CreateTags", "ec2:DeleteTags"},
		},
	})

	// aws_route_table
	register("aws_route_table", OpCreate, Rule{
		BaseActions: []string{
			"ec2:CreateRouteTable",
			"ec2:DescribeRouteTables",
			"ec2:CreateRoute",
			"ec2:AssociateRouteTable",
		},
	})
	register("aws_route_table", OpDelete, Rule{
		BaseActions: []string{
			"ec2:DeleteRouteTable",
			"ec2:DescribeRouteTables",
			"ec2:DisassociateRouteTable",
			"ec2:DeleteRoute",
		},
	})
	register("aws_route_table", OpUpdate, Rule{
		UpdateActions: []string{"ec2:DescribeRouteTables"},
		ConditionalActions: map[string][]string{
			"route": {"ec2:CreateRoute", "ec2:DeleteRoute", "ec2:ReplaceRoute"},
			"tags":  {"ec2:CreateTags", "ec2:DeleteTags"},
		},
	})

	// aws_route_table_association
	register("aws_route_table_association", OpCreate, Rule{
		BaseActions: []string{
			"ec2:AssociateRouteTable",
			"ec2:DescribeRouteTables",
		},
	})
	register("aws_route_table_association", OpDelete, Rule{
		BaseActions: []string{
			"ec2:DisassociateRouteTable",
			"ec2:DescribeRouteTables",
		},
	})
	register("aws_route_table_association", OpUpdate, Rule{
		UpdateActions: []string{"ec2:DescribeRouteTables"},
		ConditionalActions: map[string][]string{
			"route_table_id": {"ec2:ReplaceRouteTableAssociation"},
		},
	})

	// aws_eip
	register("aws_eip", OpCreate, Rule{
		BaseActions: []string{
			"ec2:AllocateAddress",
			"ec2:DescribeAddresses",
			"ec2:AssociateAddress",
		},
	})
	register("aws_eip", OpDelete, Rule{
		BaseActions: []string{
			"ec2:ReleaseAddress",
			"ec2:DescribeAddresses",
			"ec2:DisassociateAddress",
		},
	})
	register("aws_eip", OpUpdate, Rule{
		UpdateActions: []string{"ec2:DescribeAddresses"},
		ConditionalActions: map[string][]string{
			"instance":          {"ec2:AssociateAddress", "ec2:DisassociateAddress"},
			"network_interface": {"ec2:AssociateAddress", "ec2:DisassociateAddress"},
			"tags":              {"ec2:CreateTags", "ec2:DeleteTags"},
		},
	})

	// aws_vpc_endpoint
	register("aws_vpc_endpoint", OpCreate, Rule{
		BaseActions: []string{
			"ec2:CreateVpcEndpoint",
			"ec2:DescribeVpcEndpoints",
		},
	})
	register("aws_vpc_endpoint", OpDelete, Rule{
		BaseActions: []string{
			"ec2:DeleteVpcEndpoints",
			"ec2:DescribeVpcEndpoints",
		},
	})
	register("aws_vpc_endpoint", OpUpdate, Rule{
		UpdateActions: []string{"ec2:DescribeVpcEndpoints"},
		ConditionalActions: map[string][]string{
			"policy":             {"ec2:ModifyVpcEndpoint"},
			"route_table_ids":    {"ec2:ModifyVpcEndpoint"},
			"subnet_ids":         {"ec2:ModifyVpcEndpoint"},
			"security_group_ids": {"ec2:ModifyVpcEndpoint"},
			"tags":               {"ec2:CreateTags", "ec2:DeleteTags"},
		},
	})
}

package iamrules

func init() {
	// aws_codedeploy_app
	register("aws_codedeploy_app", OpCreate, Rule{
		BaseActions: []string{
			"codedeploy:CreateApplication",
			"codedeploy:GetApplication",
			"codedeploy:TagResource",
		},
	})
	register("aws_codedeploy_app", OpDelete, Rule{
		BaseActions: []string{
			"codedeploy:DeleteApplication",
			"codedeploy:GetApplication",
		},
	})
	register("aws_codedeploy_app", OpUpdate, Rule{
		UpdateActions: []string{"codedeploy:GetApplication"},
		ConditionalActions: map[string][]string{
			"name": {"codedeploy:UpdateApplication"},
			"tags": {"codedeploy:TagResource", "codedeploy:UntagResource"},
		},
	})

	// aws_codedeploy_deployment_group
	register("aws_codedeploy_deployment_group", OpCreate, Rule{
		BaseActions: []string{
			"codedeploy:CreateDeploymentGroup",
			"codedeploy:GetDeploymentGroup",
			"codedeploy:TagResource",
		},
	})
	register("aws_codedeploy_deployment_group", OpDelete, Rule{
		BaseActions: []string{
			"codedeploy:DeleteDeploymentGroup",
			"codedeploy:GetDeploymentGroup",
		},
	})
	register("aws_codedeploy_deployment_group", OpUpdate, Rule{
		UpdateActions: []string{"codedeploy:GetDeploymentGroup"},
		ConditionalActions: map[string][]string{
			"service_role_arn":               {"codedeploy:UpdateDeploymentGroup"},
			"deployment_style":               {"codedeploy:UpdateDeploymentGroup"},
			"deployment_config_name":         {"codedeploy:UpdateDeploymentGroup"},
			"alarm_configuration":            {"codedeploy:UpdateDeploymentGroup"},
			"auto_rollback_configuration":    {"codedeploy:UpdateDeploymentGroup"},
			"ec2_tag_set":                    {"codedeploy:UpdateDeploymentGroup"},
			"ecs_service":                    {"codedeploy:UpdateDeploymentGroup"},
			"load_balancer_info":             {"codedeploy:UpdateDeploymentGroup"},
			"blue_green_deployment_config":   {"codedeploy:UpdateDeploymentGroup"},
			"autoscaling_groups":             {"codedeploy:UpdateDeploymentGroup"},
			"tags":                           {"codedeploy:TagResource", "codedeploy:UntagResource"},
		},
	})

	// aws_codedeploy_deployment_config
	register("aws_codedeploy_deployment_config", OpCreate, Rule{
		BaseActions: []string{
			"codedeploy:CreateDeploymentConfig",
			"codedeploy:GetDeploymentConfig",
		},
	})
	register("aws_codedeploy_deployment_config", OpDelete, Rule{
		BaseActions: []string{
			"codedeploy:DeleteDeploymentConfig",
			"codedeploy:GetDeploymentConfig",
		},
	})
}

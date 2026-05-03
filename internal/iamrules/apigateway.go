package iamrules

func init() {
	// aws_api_gateway_rest_api
	register("aws_api_gateway_rest_api", OpCreate, Rule{
		BaseActions: []string{
			"apigateway:POST",
			"apigateway:GET",
		},
	})
	register("aws_api_gateway_rest_api", OpDelete, Rule{
		BaseActions: []string{
			"apigateway:DELETE",
			"apigateway:GET",
		},
	})
	register("aws_api_gateway_rest_api", OpUpdate, Rule{
		UpdateActions: []string{"apigateway:GET"},
		ConditionalActions: map[string][]string{
			"name":              {"apigateway:PATCH"},
			"description":       {"apigateway:PATCH"},
			"endpoint_configuration": {"apigateway:PATCH"},
			"policy":            {"apigateway:PUT"},
			"body":              {"apigateway:PUT"},
			"tags":              {"apigateway:PUT"},
		},
	})

	// aws_api_gateway_deployment
	register("aws_api_gateway_deployment", OpCreate, Rule{
		BaseActions: []string{
			"apigateway:POST",
			"apigateway:GET",
		},
	})
	register("aws_api_gateway_deployment", OpDelete, Rule{
		BaseActions: []string{
			"apigateway:DELETE",
			"apigateway:GET",
		},
	})

	// aws_api_gateway_stage
	register("aws_api_gateway_stage", OpCreate, Rule{
		BaseActions: []string{
			"apigateway:POST",
			"apigateway:GET",
			"apigateway:PUT",
		},
	})
	register("aws_api_gateway_stage", OpDelete, Rule{
		BaseActions: []string{
			"apigateway:DELETE",
			"apigateway:GET",
		},
	})
	register("aws_api_gateway_stage", OpUpdate, Rule{
		UpdateActions: []string{"apigateway:GET"},
		ConditionalActions: map[string][]string{
			"access_log_settings":   {"apigateway:PATCH"},
			"cache_cluster_enabled": {"apigateway:PATCH"},
			"cache_cluster_size":    {"apigateway:PATCH"},
			"xray_tracing_enabled":  {"apigateway:PATCH"},
			"default_route_settings": {"apigateway:PATCH"},
			"tags":                  {"apigateway:PUT"},
		},
	})

	// aws_api_gateway_resource
	register("aws_api_gateway_resource", OpCreate, Rule{
		BaseActions: []string{
			"apigateway:POST",
			"apigateway:GET",
		},
	})
	register("aws_api_gateway_resource", OpDelete, Rule{
		BaseActions: []string{
			"apigateway:DELETE",
			"apigateway:GET",
		},
	})
	register("aws_api_gateway_resource", OpUpdate, Rule{
		UpdateActions: []string{"apigateway:GET"},
		ConditionalActions: map[string][]string{
			"path_part": {"apigateway:PATCH"},
		},
	})

	// aws_api_gateway_method
	register("aws_api_gateway_method", OpCreate, Rule{
		BaseActions: []string{
			"apigateway:PUT",
			"apigateway:GET",
		},
	})
	register("aws_api_gateway_method", OpDelete, Rule{
		BaseActions: []string{
			"apigateway:DELETE",
			"apigateway:GET",
		},
	})
	register("aws_api_gateway_method", OpUpdate, Rule{
		UpdateActions: []string{"apigateway:GET"},
		ConditionalActions: map[string][]string{
			"authorization":     {"apigateway:PATCH"},
			"authorizer_id":     {"apigateway:PATCH"},
			"request_parameters": {"apigateway:PATCH"},
		},
	})

	// aws_apigatewayv2_api (HTTP API / WebSocket API)
	register("aws_apigatewayv2_api", OpCreate, Rule{
		BaseActions: []string{
			"apigateway:POST",
			"apigateway:GET",
		},
	})
	register("aws_apigatewayv2_api", OpDelete, Rule{
		BaseActions: []string{
			"apigateway:DELETE",
			"apigateway:GET",
		},
	})
	register("aws_apigatewayv2_api", OpUpdate, Rule{
		UpdateActions: []string{"apigateway:GET"},
		ConditionalActions: map[string][]string{
			"name":                 {"apigateway:PATCH"},
			"description":          {"apigateway:PATCH"},
			"cors_configuration":   {"apigateway:PATCH"},
			"disable_execute_api_endpoint": {"apigateway:PATCH"},
			"route_key":            {"apigateway:PATCH"},
			"target":               {"apigateway:PATCH"},
			"body":                 {"apigateway:PUT"},
			"tags":                 {"apigateway:POST"},
		},
	})

	// aws_apigatewayv2_stage
	register("aws_apigatewayv2_stage", OpCreate, Rule{
		BaseActions: []string{
			"apigateway:POST",
			"apigateway:GET",
		},
	})
	register("aws_apigatewayv2_stage", OpDelete, Rule{
		BaseActions: []string{
			"apigateway:DELETE",
			"apigateway:GET",
		},
	})
	register("aws_apigatewayv2_stage", OpUpdate, Rule{
		UpdateActions: []string{"apigateway:GET"},
		ConditionalActions: map[string][]string{
			"auto_deploy":           {"apigateway:PATCH"},
			"access_log_settings":   {"apigateway:PATCH"},
			"default_route_settings": {"apigateway:PATCH"},
			"stage_variables":       {"apigateway:PATCH"},
			"tags":                  {"apigateway:POST"},
		},
	})

	// aws_apigatewayv2_integration
	register("aws_apigatewayv2_integration", OpCreate, Rule{
		BaseActions: []string{
			"apigateway:POST",
			"apigateway:GET",
		},
	})
	register("aws_apigatewayv2_integration", OpDelete, Rule{
		BaseActions: []string{
			"apigateway:DELETE",
			"apigateway:GET",
		},
	})
	register("aws_apigatewayv2_integration", OpUpdate, Rule{
		UpdateActions: []string{"apigateway:GET"},
		ConditionalActions: map[string][]string{
			"integration_uri":    {"apigateway:PATCH"},
			"integration_type":   {"apigateway:PATCH"},
			"integration_method": {"apigateway:PATCH"},
			"payload_format_version": {"apigateway:PATCH"},
			"timeout_milliseconds":   {"apigateway:PATCH"},
		},
	})

	// aws_apigatewayv2_route
	register("aws_apigatewayv2_route", OpCreate, Rule{
		BaseActions: []string{
			"apigateway:POST",
			"apigateway:GET",
		},
	})
	register("aws_apigatewayv2_route", OpDelete, Rule{
		BaseActions: []string{
			"apigateway:DELETE",
			"apigateway:GET",
		},
	})
	register("aws_apigatewayv2_route", OpUpdate, Rule{
		UpdateActions: []string{"apigateway:GET"},
		ConditionalActions: map[string][]string{
			"route_key":      {"apigateway:PATCH"},
			"target":         {"apigateway:PATCH"},
			"authorizer_id":  {"apigateway:PATCH"},
			"authorization_type": {"apigateway:PATCH"},
		},
	})
}

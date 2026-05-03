package iamrules

func init() {
	// aws_kinesis_stream
	register("aws_kinesis_stream", OpCreate, Rule{
		BaseActions: []string{
			"kinesis:CreateStream",
			"kinesis:DescribeStream",
			"kinesis:AddTagsToStream",
		},
	})
	register("aws_kinesis_stream", OpDelete, Rule{
		BaseActions: []string{
			"kinesis:DeleteStream",
			"kinesis:DescribeStream",
		},
	})
	register("aws_kinesis_stream", OpUpdate, Rule{
		UpdateActions: []string{"kinesis:DescribeStream"},
		ConditionalActions: map[string][]string{
			"shard_count":          {"kinesis:UpdateShardCount"},
			"retention_period":     {"kinesis:IncreaseStreamRetentionPeriod", "kinesis:DecreaseStreamRetentionPeriod"},
			"encryption_type":      {"kinesis:StartStreamEncryption", "kinesis:StopStreamEncryption"},
			"kms_key_id":           {"kinesis:StartStreamEncryption", "kinesis:StopStreamEncryption"},
			"stream_mode_details":  {"kinesis:UpdateStreamMode"},
			"tags":                 {"kinesis:AddTagsToStream", "kinesis:RemoveTagsFromStream"},
		},
	})

	// aws_kinesis_firehose_delivery_stream
	register("aws_kinesis_firehose_delivery_stream", OpCreate, Rule{
		BaseActions: []string{
			"firehose:CreateDeliveryStream",
			"firehose:DescribeDeliveryStream",
			"firehose:TagDeliveryStream",
		},
	})
	register("aws_kinesis_firehose_delivery_stream", OpDelete, Rule{
		BaseActions: []string{
			"firehose:DeleteDeliveryStream",
			"firehose:DescribeDeliveryStream",
		},
	})
	register("aws_kinesis_firehose_delivery_stream", OpUpdate, Rule{
		UpdateActions: []string{"firehose:DescribeDeliveryStream"},
		ConditionalActions: map[string][]string{
			"s3_configuration":              {"firehose:UpdateDestination"},
			"extended_s3_configuration":     {"firehose:UpdateDestination"},
			"redshift_configuration":        {"firehose:UpdateDestination"},
			"elasticsearch_configuration":   {"firehose:UpdateDestination"},
			"opensearch_configuration":      {"firehose:UpdateDestination"},
			"splunk_configuration":          {"firehose:UpdateDestination"},
			"http_endpoint_configuration":   {"firehose:UpdateDestination"},
			"server_side_encryption":        {"firehose:StartDeliveryStreamEncryption", "firehose:StopDeliveryStreamEncryption"},
			"tags":                          {"firehose:TagDeliveryStream", "firehose:UntagDeliveryStream"},
		},
	})

	// aws_kinesis_analytics_application
	register("aws_kinesis_analytics_application", OpCreate, Rule{
		BaseActions: []string{
			"kinesisanalytics:CreateApplication",
			"kinesisanalytics:DescribeApplication",
			"kinesisanalytics:AddApplicationTag",
		},
	})
	register("aws_kinesis_analytics_application", OpDelete, Rule{
		BaseActions: []string{
			"kinesisanalytics:DeleteApplication",
			"kinesisanalytics:DescribeApplication",
		},
	})
	register("aws_kinesis_analytics_application", OpUpdate, Rule{
		UpdateActions: []string{"kinesisanalytics:DescribeApplication"},
		ConditionalActions: map[string][]string{
			"inputs":  {"kinesisanalytics:UpdateApplication"},
			"outputs": {"kinesisanalytics:UpdateApplication"},
			"code":    {"kinesisanalytics:UpdateApplication"},
			"tags":    {"kinesisanalytics:AddApplicationTag", "kinesisanalytics:RemoveApplicationTag"},
		},
	})
}

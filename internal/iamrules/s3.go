package iamrules

func init() {
	register("aws_s3_bucket", OpCreate, Rule{BaseActions: []string{"s3:CreateBucket"}})
	register("aws_s3_bucket", OpDelete, Rule{BaseActions: []string{"s3:DeleteBucket"}})

	register("aws_s3_bucket", OpUpdate, Rule{
		UpdateActions: []string{}, // update 単体の base はなし
		ConditionalActions: map[string][]string{
			"tags":                                 {"s3:PutBucketTagging"},
			"versioning":                           {"s3:PutBucketVersioning"},
			"server_side_encryption_configuration": {"s3:PutEncryptionConfiguration"},
			"logging":                              {"s3:PutBucketLogging"},
			"website":                              {"s3:PutBucketWebsite"},
			"cors_rule":                            {"s3:PutBucketCORS"},
			"lifecycle_rule":                       {"s3:PutLifecycleConfiguration"},
			"replication_configuration":            {"s3:PutReplicationConfiguration"},
			"acl":                                  {"s3:PutBucketAcl"},
			"policy":                               {"s3:PutBucketPolicy"},
		},
	})
}

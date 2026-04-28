package iamrules

func init() {
	register("aws_s3_bucket", OpCreate, Rule{BaseActions: []string{"s3:CreateBucket"}})
	register("aws_s3_bucket", OpDelete, Rule{BaseActions: []string{"s3:DeleteBucket"}})
}

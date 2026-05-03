# tfgate

* `tfgate` is a CLI tool to check IAM permissions before `terraform apply` on AWS.
* It parses your Terraform plan, computes the required IAM actions, and calls `iam:SimulatePrincipalPolicy` to verify your credentials — before apply.
* `terraform plan` only requires read permissions, so permission errors are invisible until `apply` fails halfway through. `tfgate` catches them early.

## Features

* [x] Parse `terraform show -json` output
* [x] Automatically derive required IAM actions from plan
* [x] Simulate against current credentials via `iam:SimulatePrincipalPolicy`
* [x] Support create / delete / update operations
* [x] Text and JSON output (`--format`)
* [x] AWS profile and region flags
* [x] Exit code 1 on denied — CI-friendly

### Supported services

EC2, S3, RDS, Aurora, DynamoDB, ElastiCache, ECS, EKS, Lambda, SNS, SQS, EventBridge, Kinesis, Step Functions, VPC, ALB/NLB, Route 53, CloudFront, API Gateway, IAM, KMS, Secrets Manager, ACM, CloudWatch, CloudTrail, SSM, EBS, EFS, CodeBuild, CodeDeploy, CodePipeline, CodeCommit, Glue, Athena

## Installation

```
go install github.com/komisan19/tfgate@latest
```

## Usage

```bash
# 1. Generate a plan and export it as JSON
terraform plan -out=tfplan
terraform show -json tfplan > plan.json

# 2. Check IAM permissions
tfgate check plan.json
```

### Flags

```
tfgate check [flags] <plan.json>

  --format string   Output format: text or json (default "text")
  --profile string  AWS shared config profile (overrides AWS_PROFILE)
  --region string   AWS region
```

## CI integration

```yaml
- name: Generate plan
  run: |
    terraform plan -out=tfplan
    terraform show -json tfplan > plan.json

- name: Check IAM permissions
  run: tfgate check plan.json
```

## License

MIT — see [LICENSE](./LICENSE).

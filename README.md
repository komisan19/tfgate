# tfgate

Check IAM permissions before `terraform apply` on AWS.

## Why

`terraform plan` succeeds but `terraform apply` fails with permission errors?

`terraform plan` only requires read permissions, so it cannot detect missing write permissions until `apply` runs and fails halfway through.

`tfgate` parses your plan, computes the IAM actions each resource change requires, and calls `iam:SimulatePrincipalPolicy` to verify your current credentials can perform them — before you apply.

## Install

### From source

```bash
go install github.com/komisan19/tfgate@latest
```

### Pre-built binaries

Download from [Releases](https://github.com/komisan19/tfgate/releases).

Supported platforms: `linux-amd64`, `linux-arm64`, `darwin-amd64`, `darwin-arm64`, `windows-amd64`.

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
tfgate [--version] <command> [flags] [args]

Global Flags:
  --version         Print version and exit

tfgate check [flags] <plan.json>

  --format string   Output format: text or json (default "text")
  --profile string  AWS shared config profile (overrides AWS_PROFILE)
  --region string   AWS region
```

### Examples

```bash
# JSON output (useful for programmatic processing)
tfgate check --format json plan.json | jq .

# Use a specific AWS profile and region
tfgate check --profile staging --region us-east-1 plan.json

# Print version
tfgate --version
```

## CI integration

Example GitHub Actions step:

```yaml
- name: Generate plan
  run: |
    terraform plan -out=tfplan
    terraform show -json tfplan > plan.json

- name: Check IAM permissions
  run: tfgate check plan.json
```

The job fails on exit code 1, blocking apply when permissions are insufficient.

## Comparison with `aws_iam_principal_policy_simulation`

Both tools call the same AWS API (`iam:SimulatePrincipalPolicy`), but they fit different workflows.

|                  | `aws_iam_principal_policy_simulation` | `tfgate`                                          |
| ---------------- | ------------------------------------- | ------------------------------------------------ |
| Form             | Terraform data source (HCL)           | External CLI                                     |
| Action discovery | Manual — you list actions in HCL      | Automatic — derived from `plan.json`             |
| Plan integration | Inside Terraform run                  | Runs separately on `terraform show -json` output |
| CI integration   | Via `terraform plan` exit code        | Independent step, language-agnostic              |
| Maintenance      | Update HCL when resources change      | Update once in tfgate's rule registry             |
| Coverage         | Whatever you write in HCL             | Resources covered by tfgate's registry            |
| Best for         | Specific permission assertions        | Pre-apply sanity check across the entire plan    |

## Supported resources

- Resources
  - [ ] `aws_s3_bucket` (create, delete)
  - [ ] `aws_iam_role` (create, delete)
  - [ ] `aws_instance` (create, delete)

## License

MIT — see [LICENSE](./LICENSE).

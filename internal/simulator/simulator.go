package simulator

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/iam"
	"github.com/aws/aws-sdk-go-v2/service/iam/types"
	"github.com/aws/aws-sdk-go-v2/service/sts"
)

type Client struct {
	sts       *sts.Client
	iam       *iam.Client
	callerARN string
}

func New(ctx context.Context) (*Client, error) {
	cfg, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		return nil, fmt.Errorf("load AWS config: %w", err)
	}

	c := &Client{
		sts: sts.NewFromConfig(cfg),
		iam: iam.NewFromConfig(cfg),
	}

	arn, err := c.fetchCallerARN(ctx)
	if err != nil {
		return nil, fmt.Errorf("fetch caller ARN: %w", err)
	}

	c.callerARN = arn

	return c, nil

}

func (c *Client) CallerARN() string {
	return c.callerARN
}

func (c *Client) fetchCallerARN(ctx context.Context) (string, error) {
	out, err := c.sts.GetCallerIdentity(ctx, &sts.GetCallerIdentityInput{})
	if err != nil {
		return "", fmt.Errorf("get caller identity: %w", err)
	}
	if out.Arn == nil {
		return "", fmt.Errorf("got nil ARN from STS")
	}
	return *out.Arn, nil
}

func (c *Client) Simulate(ctx context.Context, actions []string) (allowed []string, denied []string, err error) {
	if len(actions) == 0 {
		return nil, nil, fmt.Errorf("actions is empty")
	}

	input := &iam.SimulatePrincipalPolicyInput{
		PolicySourceArn: aws.String(c.callerARN),
		ActionNames:     actions,
		ResourceArns:    []string{"*"},
	}

	out, err := c.iam.SimulatePrincipalPolicy(ctx, input)
	if err != nil {
		return nil, nil, fmt.Errorf("call iam:SimulatePrincipalPolicy: %w", err)
	}

	for _, result := range out.EvaluationResults {
		if result.EvalActionName == nil {
			continue
		}
		actionName := *result.EvalActionName

		if result.EvalDecision == types.PolicyEvaluationDecisionTypeAllowed {
			allowed = append(allowed, actionName)
		} else {
			denied = append(denied, actionName)
		}
	}

	return allowed, denied, nil
}

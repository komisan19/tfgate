package main

import (
	"context"
	"fmt"
	"os"

	"github.com/komisan19/tfgate/internal/iamrules"
	"github.com/komisan19/tfgate/internal/plan"
	"github.com/komisan19/tfgate/internal/report"
	"github.com/komisan19/tfgate/internal/simulator"
)

func main() {
	if len(os.Args) < 3 || os.Args[1] != "check" {
		fmt.Fprintln(os.Stderr, "usage: tfgate check <plan.json>")
		os.Exit(2)
	}

	ctx := context.Background()

	client, err := simulator.New(ctx)
	if err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(2)
	}

	p, err := plan.Load(os.Args[2])
	if err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(2)
	}

	var results []report.Result

	for _, rc := range p.ResourceChanges {
		ops, reason := determineOps(rc.Change.Actions)
		if len(ops) == 0 {
			if reason != "" {
				results = append(results, report.Result{
					Address:      rc.Address,
					ResourceType: rc.Type,
					Operation:    "",
					Status:       report.StatusWarning,
					Message:      reason,
				})
			}
			continue
		}

		for _, op := range ops {
			rule, ok := iamrules.Lookup(rc.Type, op)
			if !ok {
				results = append(results, report.Result{
					Address:      rc.Address,
					ResourceType: rc.Type,
					Operation:    string(op),
					Status:       report.StatusWarning,
					Message:      "no rule registered for " + rc.Type,
				})
				continue
			}

			actions := iamrules.Resolve(rule)
			_, denied, err := client.Simulate(ctx, actions)
			if err != nil {
				results = append(results, report.Result{
					Address:      rc.Address,
					ResourceType: rc.Type,
					Operation:    string(op),
					Status:       report.StatusWarning,
					Message:      "simulation failed: " + err.Error(),
				})
				continue
			}

			if len(denied) == 0 {
				results = append(results, report.Result{
					Address:      rc.Address,
					ResourceType: rc.Type,
					Operation:    string(op),
					Status:       report.StatusOK,
					Required:     actions,
				})
			} else {
				results = append(results, report.Result{
					Address:      rc.Address,
					ResourceType: rc.Type,
					Operation:    string(op),
					Status:       report.StatusDenied,
					Required:     actions,
					Missing:      denied,
				})
			}
		}
	}

	if err := report.WriteText(os.Stdout, client.CallerARN(), len(p.ResourceChanges), results); err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(2)
	}

	if report.HasFailure(results) {
		os.Exit(1)
	}

}

func determineOps(actions []string) (ops []iamrules.Operation, reason string) {
	switch {
	case len(actions) == 1:
		switch actions[0] {
		case "create":
			return []iamrules.Operation{iamrules.OpCreate}, ""
		case "delete":
			return []iamrules.Operation{iamrules.OpDelete}, ""
		case "update":
			return nil, "update not supported in v0.1"
		case "no-op", "read":
			return nil, ""
		default:
			return nil, fmt.Sprintf("unknown action: %s", actions[0])
		}
	case len(actions) == 2:
		set := map[string]bool{actions[0]: true, actions[1]: true}
		if set["create"] && set["delete"] {
			return []iamrules.Operation{iamrules.OpCreate, iamrules.OpDelete}, ""
		}
		return nil, fmt.Sprintf("unknown action combo: %v", actions)
	default:
		return nil, fmt.Sprintf("unexpected actions length: %v", actions)
	}
}

package main

import (
	"context"
	"flag"
	"fmt"
	"os"

	"github.com/komisan19/tfgate/internal/iamrules"
	"github.com/komisan19/tfgate/internal/plan"
	"github.com/komisan19/tfgate/internal/report"
	"github.com/komisan19/tfgate/internal/simulator"
)

var version = "dev"

func main() {
	globalFlags := flag.NewFlagSet("tfgate", flag.ContinueOnError)
	versionFlag := globalFlags.Bool("version", false, "print version and exit")
	globalFlags.Usage = func() {
		fmt.Fprintln(os.Stderr, "Usage: tfgate [--version] <command> [flags] [args]")
		fmt.Fprintln(os.Stderr, "")
		fmt.Fprintln(os.Stderr, "Commands:")
		fmt.Fprintln(os.Stderr, "  check   Simulate IAM permissions against a Terraform plan")
		fmt.Fprintln(os.Stderr, "")
		fmt.Fprintln(os.Stderr, "Global Flags:")
		globalFlags.PrintDefaults()
	}

	if err := globalFlags.Parse(os.Args[1:]); err != nil {
		if err == flag.ErrHelp {
			os.Exit(0)
		}
		os.Exit(2)
	}

	if *versionFlag {
		fmt.Println("tfgate", version)
		os.Exit(0)
	}

	args := globalFlags.Args()
	if len(args) == 0 {
		globalFlags.Usage()
		os.Exit(2)
	}

	switch args[0] {
	case "check":
		runCheck(args[1:])
	default:
		fmt.Fprintf(os.Stderr, "unknown command: %s\n", args[0])
		globalFlags.Usage()
		os.Exit(2)
	}
}

func runCheck(args []string) {
	checkFlags := flag.NewFlagSet("check", flag.ContinueOnError)
	format := checkFlags.String("format", "text", "output format: text or json")
	profile := checkFlags.String("profile", "", "AWS shared config profile (overrides AWS_PROFILE)")
	region := checkFlags.String("region", "", "AWS region")
	checkFlags.Usage = func() {
		fmt.Fprintln(os.Stderr, "Usage: tfgate check [flags] <plan.json>")
		fmt.Fprintln(os.Stderr, "")
		fmt.Fprintln(os.Stderr, "Flags:")
		checkFlags.PrintDefaults()
	}

	if err := checkFlags.Parse(args); err != nil {
		if err == flag.ErrHelp {
			os.Exit(0)
		}
		os.Exit(2)
	}

	if checkFlags.NArg() < 1 {
		checkFlags.Usage()
		os.Exit(2)
	}

	if *format != "text" && *format != "json" {
		fmt.Fprintf(os.Stderr, "invalid --format %q: must be text or json\n", *format)
		os.Exit(2)
	}

	planPath := checkFlags.Arg(0)
	ctx := context.Background()

	client, err := simulator.New(ctx, simulator.Options{
		Profile: *profile,
		Region:  *region,
	})
	if err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(2)
	}

	p, err := plan.Load(planPath)
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

	var writeErr error
	switch *format {
	case "json":
		writeErr = report.WriteJSON(os.Stdout, client.CallerARN(), len(p.ResourceChanges), results)
	default:
		writeErr = report.WriteText(os.Stdout, client.CallerARN(), len(p.ResourceChanges), results)
	}
	if writeErr != nil {
		fmt.Fprintln(os.Stderr, "error:", writeErr)
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

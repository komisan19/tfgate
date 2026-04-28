package report

import (
	"fmt"
	"io"
	"strings"
)

type ResultStatus string

const (
	StatusOK      ResultStatus = "ok"
	StatusDenied  ResultStatus = "denied"
	StatusWarning ResultStatus = "warning"
)

type Result struct {
	Address      string
	ResourceType string
	Operation    string
	Status       ResultStatus
	Required     []string
	Missing      []string
	Message      string
}

func WriteText(w io.Writer, principalARN string, totalChanges int, results []Result) error {
	if _, err := fmt.Fprintf(w, "Current principal: %s\n", principalARN); err != nil {
		return err
	}
	if _, err := fmt.Fprintf(w, "Loaded %d resource changes\n", totalChanges); err != nil {
		return err
	}

	for _, r := range results {
		switch r.Status {
		case StatusOK:
			if _, err := fmt.Fprintf(w, "✓ %s (%s) — OK\n", r.Address, r.Operation); err != nil {
				return err
			}
		case StatusDenied:
			if _, err := fmt.Fprintf(w, "✗ %s (%s)\n    Missing: %s\n", r.Address, r.Operation, strings.Join(r.Missing, ", ")); err != nil {
				return err
			}
		case StatusWarning:
			if r.Operation != "" {
				if _, err := fmt.Fprintf(w, "⚠ %s (%s): %s\n", r.Address, r.Operation, r.Message); err != nil {
					return err
				}
			} else {
				if _, err := fmt.Fprintf(w, "⚠ %s: %s\n", r.Address, r.Message); err != nil {
					return err
				}
			}
		}
	}

	return nil
}

func HasFailure(results []Result) bool {
	for _, r := range results {
		if r.Status == StatusDenied {
			return true
		}
	}
	return false
}

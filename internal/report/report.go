package report

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strings"
)

type ResultStatus string

const (
	StatusOK      ResultStatus = "ok"
	StatusDenied  ResultStatus = "denied"
	StatusWarning ResultStatus = "warning"
	colorReset                 = "\033[0m"
	colorGreen                 = "\033[32m"
	colorRed                   = "\033[31m"
	colorYellow                = "\033[33m"
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

type jsonReport struct {
	PrincipalARN string       `json:"principal_arn"`
	TotalChanges int          `json:"total_changes"`
	Results      []jsonResult `json:"results"`
}

type jsonResult struct {
	Address      string   `json:"address"`
	ResourceType string   `json:"resource_type"`
	Operation    string   `json:"operation,omitempty"`
	Status       string   `json:"status"`
	Required     []string `json:"required,omitempty"`
	Missing      []string `json:"missing,omitempty"`
	Message      string   `json:"message,omitempty"`
}

func WriteText(w io.Writer, principalARN string, totalChanges int, results []Result, color bool) error {
	useColor := color && isTerminal(w)

	colorize := func(code, text string) string {
		if !useColor {
			return text
		}
		return code + text + colorReset
	}

	if _, err := fmt.Fprintf(w, "Current principal: %s\n", principalARN); err != nil {
		return err
	}
	if _, err := fmt.Fprintf(w, "Loaded %d resource changes\n", totalChanges); err != nil {
		return err
	}

	for _, r := range results {
		var line string
		switch r.Status {
		case StatusOK:
			line = colorize(colorGreen, fmt.Sprintf("✓ %s (%s) — OK", r.Address, r.Operation))
		case StatusDenied:
			line = colorize(colorRed, fmt.Sprintf("✗ %s (%s)", r.Address, r.Operation))
		case StatusWarning:
			if r.Operation != "" {
				line = colorize(colorYellow, fmt.Sprintf("⚠ %s (%s): %s", r.Address, r.Operation, r.Message))
			} else {
				line = colorize(colorYellow, fmt.Sprintf("⚠ %s: %s", r.Address, r.Message))
			}
		}

		if _, err := fmt.Fprintln(w, line); err != nil {
			return err
		}

		if r.Status == StatusDenied && len(r.Missing) > 0 {
			missing := "    Missing: " + strings.Join(r.Missing, ", ")
			if _, err := fmt.Fprintln(w, missing); err != nil {
				return err
			}
		}
	}

	return nil
}

func WriteJSON(w io.Writer, principalARN string, totalChanges int, results []Result) error {
	jr := jsonReport{
		PrincipalARN: principalARN,
		TotalChanges: totalChanges,
		Results:      make([]jsonResult, 0, len(results)),
	}
	for _, r := range results {
		jr.Results = append(jr.Results, jsonResult{
			Address:      r.Address,
			ResourceType: r.ResourceType,
			Operation:    r.Operation,
			Status:       string(r.Status),
			Required:     r.Required,
			Missing:      r.Missing,
			Message:      r.Message,
		})
	}
	enc := json.NewEncoder(w)
	enc.SetIndent("", "  ")
	return enc.Encode(jr)
}

func HasFailure(results []Result) bool {
	for _, r := range results {
		if r.Status == StatusDenied {
			return true
		}
	}
	return false
}

func isTerminal(w io.Writer) bool {
	if f, ok := w.(*os.File); ok {
		fi, _ := f.Stat()
		return (fi.Mode() & os.ModeCharDevice) != 0
	}
	return false
}

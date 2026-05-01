package plan

import (
	"encoding/json"
	"fmt"
	"os"
)

type Plan struct {
	ResourceChanges []ResourceChange `json:"resource_changes"`
}

type ResourceChange struct {
	Address string `json:"address"`
	Type    string `json:"type"`
	Name    string `json:"name"`
	Change  Change `json:"change"`
}

type Change struct {
	Actions []string        `json:"actions"`
	After   json.RawMessage `json:"after"`
	Before  json.RawMessage `json:"before"`
}

func Load(path string) (*Plan, error) {
	b, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("read plan: %w", err)
	}
	var p Plan
	if err := json.Unmarshal(b, &p); err != nil {
		return nil, fmt.Errorf("parse plan: %w", err)
	}

	return &p, nil
}

func rawToMap(raw json.RawMessage) (map[string]any, error) {
	if len(raw) == 0 || string(raw) == "null" {
		return nil, nil
	}
	var m map[string]any
	if err := json.Unmarshal(raw, &m); err != nil {
		return nil, fmt.Errorf("unmarshal: %w", err)
	}
	return m, nil
}

func (c *Change) AfterMap() (map[string]any, error) {
	return rawToMap(c.After)
}

func (c *Change) BeforeMap() (map[string]any, error) {
	return rawToMap(c.Before)
}

func (c *Change) ChangedKeys() ([]string, error) {
	before, err := rawToMap(c.Before)
	if err != nil {
		return nil, err
	}
	after, err := rawToMap(c.After)
	if err != nil {
		return nil, err
	}

	if before == nil {
		keys := make([]string, 0, len(after))
		for k := range after {
			keys = append(keys, k)
		}
		return keys, nil
	}

	if after == nil {
		keys := make([]string, 0, len(before))
		for k := range before {
			keys = append(keys, k)
		}
		return keys, nil
	}

	changed := map[string]bool{}
	for k, v := range after {
		bv, ok := before[k]
		if !ok {
			changed[k] = true
			continue
		}
		av, _ := json.Marshal(v)
		bvv, _ := json.Marshal(bv)
		if string(av) != string(bvv) {
			changed[k] = true
		}
	}
	for k := range before {
		if _, ok := after[k]; !ok {
			changed[k] = true
		}
	}

	result := make([]string, 0, len(changed))
	for k := range changed {
		result = append(result, k)
	}
	return result, nil
}

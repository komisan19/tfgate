package main

import (
	"encoding/json"
	"reflect"
	"sort"
	"testing"

	"github.com/komisan19/tfgate/internal/plan"
)

func TestChangedKeys_BeforeNil(t *testing.T) {
	c := plan.Change{
		Before: json.RawMessage("null"),
		After:  json.RawMessage(`{"a":1,"b":2}`),
	}
	keys, err := c.ChangedKeys()
	if err != nil {
		t.Fatal(err)
	}
	sort.Strings(keys)
	if !reflect.DeepEqual(keys, []string{"a", "b"}) {
		t.Errorf("got %v", keys)
	}
}

func TestChangedKeys_AfterNil(t *testing.T) {
	c := plan.Change{
		Before: json.RawMessage(`{"a":1,"b":2}`),
		After:  json.RawMessage("null"),
	}
	keys, err := c.ChangedKeys()
	if err != nil {
		t.Fatal(err)
	}
	sort.Strings(keys)
	if !reflect.DeepEqual(keys, []string{"a", "b"}) {
		t.Errorf("got %v", keys)
	}
}

func TestChangedKeys_BothNil(t *testing.T) {
	c := plan.Change{
		Before: json.RawMessage("null"),
		After:  json.RawMessage("null"),
	}
	keys, err := c.ChangedKeys()
	if err != nil {
		t.Fatal(err)
	}
	if len(keys) != 0 {
		t.Errorf("expected empty, got %v", keys)
	}
}

func TestChangedKeys_NoChange(t *testing.T) {
	c := plan.Change{
		Before: json.RawMessage(`{"a":1}`),
		After:  json.RawMessage(`{"a":1}`),
	}
	keys, err := c.ChangedKeys()
	if err != nil {
		t.Fatal(err)
	}
	if len(keys) != 0 {
		t.Errorf("expected empty, got %v", keys)
	}
}

func TestChangedKeys_ValueChanged(t *testing.T) {
	c := plan.Change{
		Before: json.RawMessage(`{"a":1,"b":2}`),
		After:  json.RawMessage(`{"a":9,"b":2}`),
	}
	keys, err := c.ChangedKeys()
	if err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(keys, []string{"a"}) {
		t.Errorf("got %v", keys)
	}
}

func TestChangedKeys_KeyAdded(t *testing.T) {
	c := plan.Change{
		Before: json.RawMessage(`{"a":1}`),
		After:  json.RawMessage(`{"a":1,"b":2}`),
	}
	keys, err := c.ChangedKeys()
	if err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(keys, []string{"b"}) {
		t.Errorf("got %v", keys)
	}
}

func TestChangedKeys_KeyRemoved(t *testing.T) {
	c := plan.Change{
		Before: json.RawMessage(`{"a":1,"b":2}`),
		After:  json.RawMessage(`{"a":1}`),
	}
	keys, err := c.ChangedKeys()
	if err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(keys, []string{"b"}) {
		t.Errorf("got %v", keys)
	}
}

func TestChangedKeys_NestedChange(t *testing.T) {
	c := plan.Change{
		Before: json.RawMessage(`{"tags":{"env":"dev"}}`),
		After:  json.RawMessage(`{"tags":{"env":"prod"}}`),
	}
	keys, err := c.ChangedKeys()
	if err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(keys, []string{"tags"}) {
		t.Errorf("got %v", keys)
	}
}

func TestChangedKeys_MultipleChanges(t *testing.T) {
	c := plan.Change{
		Before: json.RawMessage(`{"a":1,"b":2,"c":3}`),
		After:  json.RawMessage(`{"a":9,"b":2,"d":4}`),
	}
	keys, err := c.ChangedKeys()
	if err != nil {
		t.Fatal(err)
	}
	sort.Strings(keys)
	if !reflect.DeepEqual(keys, []string{"a", "c", "d"}) {
		t.Errorf("got %v", keys)
	}
}

func TestAfterMap_ValidJSON(t *testing.T) {
	c := plan.Change{After: json.RawMessage(`{"key":"val"}`)}
	m, err := c.AfterMap()
	if err != nil {
		t.Fatal(err)
	}
	if m["key"] != "val" {
		t.Errorf("got %v", m)
	}
}

func TestAfterMap_Null(t *testing.T) {
	c := plan.Change{After: json.RawMessage("null")}
	m, err := c.AfterMap()
	if err != nil {
		t.Fatal(err)
	}
	if m != nil {
		t.Errorf("expected nil, got %v", m)
	}
}

func TestAfterMap_InvalidJSON(t *testing.T) {
	c := plan.Change{After: json.RawMessage(`{broken}`)}
	_, err := c.AfterMap()
	if err == nil {
		t.Error("expected error")
	}
}

func TestBeforeMap_ValidJSON(t *testing.T) {
	c := plan.Change{Before: json.RawMessage(`{"a":1}`)}
	m, err := c.BeforeMap()
	if err != nil {
		t.Fatal(err)
	}
	if _, ok := m["a"]; !ok {
		t.Errorf("expected key 'a', got %v", m)
	}
}

func TestLoad_ValidFile(t *testing.T) {
	p, err := plan.Load("testdata/valid.json")
	if err != nil {
		t.Fatal(err)
	}
	if len(p.ResourceChanges) != 2 {
		t.Errorf("expected 2 resource changes, got %d", len(p.ResourceChanges))
	}
	if p.ResourceChanges[0].Address != "aws_iam_role.test" {
		t.Errorf("unexpected address: %s", p.ResourceChanges[0].Address)
	}
	if len(p.ResourceChanges[0].Change.Actions) != 1 || p.ResourceChanges[0].Change.Actions[0] != "create" {
		t.Errorf("unexpected actions: %v", p.ResourceChanges[0].Change.Actions)
	}
}

func TestLoad_NotFound(t *testing.T) {
	_, err := plan.Load("testdata/nonexistent.json")
	if err == nil {
		t.Error("expected error")
	}
}

func TestLoad_InvalidJSON(t *testing.T) {
	_, err := plan.Load("testdata/invalid.json")
	if err == nil {
		t.Error("expected error")
	}
}

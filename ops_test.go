package main

import (
	"testing"

	"github.com/komisan19/tfgate/internal/iamrules"
)

func TestDetermineOps_Create(t *testing.T) {
	ops, reason := determineOps([]string{"create"})
	if len(ops) != 1 || ops[0] != iamrules.OpCreate || reason != "" {
		t.Errorf("got ops=%v reason=%q", ops, reason)
	}
}

func TestDetermineOps_Delete(t *testing.T) {
	ops, reason := determineOps([]string{"delete"})
	if len(ops) != 1 || ops[0] != iamrules.OpDelete || reason != "" {
		t.Errorf("got ops=%v reason=%q", ops, reason)
	}
}

func TestDetermineOps_Update(t *testing.T) {
	ops, reason := determineOps([]string{"update"})
	if len(ops) != 1 || ops[0] != iamrules.OpUpdate || reason != "" {
		t.Errorf("got ops=%v reason=%q", ops, reason)
	}
}

func TestDetermineOps_NoOp(t *testing.T) {
	ops, reason := determineOps([]string{"no-op"})
	if len(ops) != 0 || reason != "" {
		t.Errorf("got ops=%v reason=%q", ops, reason)
	}
}

func TestDetermineOps_Read(t *testing.T) {
	ops, reason := determineOps([]string{"read"})
	if len(ops) != 0 || reason != "" {
		t.Errorf("got ops=%v reason=%q", ops, reason)
	}
}

func TestDetermineOps_CreateDelete(t *testing.T) {
	ops, reason := determineOps([]string{"create", "delete"})
	if len(ops) != 2 || reason != "" {
		t.Fatalf("got ops=%v reason=%q", ops, reason)
	}
	opSet := map[iamrules.Operation]bool{ops[0]: true, ops[1]: true}
	if !opSet[iamrules.OpCreate] || !opSet[iamrules.OpDelete] {
		t.Errorf("expected OpCreate+OpDelete, got %v", ops)
	}
}

func TestDetermineOps_DeleteCreate(t *testing.T) {
	ops, reason := determineOps([]string{"delete", "create"})
	if len(ops) != 2 || reason != "" {
		t.Fatalf("got ops=%v reason=%q", ops, reason)
	}
	opSet := map[iamrules.Operation]bool{ops[0]: true, ops[1]: true}
	if !opSet[iamrules.OpCreate] || !opSet[iamrules.OpDelete] {
		t.Errorf("expected OpCreate+OpDelete, got %v", ops)
	}
}

func TestDetermineOps_UnknownSingle(t *testing.T) {
	ops, reason := determineOps([]string{"replace"})
	if len(ops) != 0 || reason == "" {
		t.Errorf("got ops=%v reason=%q", ops, reason)
	}
}

func TestDetermineOps_InvalidCombo(t *testing.T) {
	ops, reason := determineOps([]string{"create", "update"})
	if len(ops) != 0 || reason == "" {
		t.Errorf("got ops=%v reason=%q", ops, reason)
	}
}

func TestDetermineOps_Empty(t *testing.T) {
	ops, reason := determineOps([]string{})
	if len(ops) != 0 || reason == "" {
		t.Errorf("got ops=%v reason=%q", ops, reason)
	}
}

func TestDetermineOps_ThreeActions(t *testing.T) {
	ops, reason := determineOps([]string{"create", "delete", "update"})
	if len(ops) != 0 || reason == "" {
		t.Errorf("got ops=%v reason=%q", ops, reason)
	}
}

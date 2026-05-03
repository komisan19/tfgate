package main

import (
	"reflect"
	"sort"
	"testing"

	"github.com/komisan19/tfgate/internal/iamrules"
)

func TestLookup_KnownResourceCreate(t *testing.T) {
	rule, ok := iamrules.Lookup("aws_iam_role", iamrules.OpCreate)
	if !ok {
		t.Fatal("expected ok=true")
	}
	if len(rule.BaseActions) == 0 {
		t.Error("expected BaseActions to be non-empty")
	}
}

func TestLookup_KnownResourceUpdate(t *testing.T) {
	_, ok := iamrules.Lookup("aws_iam_role", iamrules.OpUpdate)
	if !ok {
		t.Fatal("expected ok=true")
	}
}

func TestLookup_KnownResourceDelete(t *testing.T) {
	_, ok := iamrules.Lookup("aws_iam_role", iamrules.OpDelete)
	if !ok {
		t.Fatal("expected ok=true")
	}
}

func TestLookup_UnknownResource(t *testing.T) {
	_, ok := iamrules.Lookup("aws_nonexistent_resource", iamrules.OpCreate)
	if ok {
		t.Error("expected ok=false")
	}
}

func TestLookup_KnownResourceUnregisteredOp(t *testing.T) {
	// aws_iam_role_policy_attachment は Create/Delete のみ登録
	_, ok := iamrules.Lookup("aws_iam_role_policy_attachment", iamrules.OpUpdate)
	if ok {
		t.Error("expected ok=false")
	}
}

func TestResolve_Create_UsesBaseActions(t *testing.T) {
	rule := iamrules.Rule{
		BaseActions: []string{"iam:CreateRole", "iam:GetRole"},
	}
	actions := iamrules.Resolve(rule, iamrules.OpCreate, nil)
	sort.Strings(actions)
	if !reflect.DeepEqual(actions, []string{"iam:CreateRole", "iam:GetRole"}) {
		t.Errorf("got %v", actions)
	}
}

func TestResolve_Delete_UsesBaseActions(t *testing.T) {
	rule := iamrules.Rule{
		BaseActions: []string{"iam:DeleteRole"},
	}
	actions := iamrules.Resolve(rule, iamrules.OpDelete, nil)
	if !reflect.DeepEqual(actions, []string{"iam:DeleteRole"}) {
		t.Errorf("got %v", actions)
	}
}

func TestResolve_Update_UsesUpdateActions(t *testing.T) {
	rule := iamrules.Rule{
		BaseActions:   []string{"ignored"},
		UpdateActions: []string{"iam:GetPolicy"},
	}
	actions := iamrules.Resolve(rule, iamrules.OpUpdate, nil)
	if !reflect.DeepEqual(actions, []string{"iam:GetPolicy"}) {
		t.Errorf("got %v, BaseActions must be ignored for update", actions)
	}
}

func TestResolve_Update_EmptyUpdateActions(t *testing.T) {
	rule := iamrules.Rule{
		BaseActions:   []string{"iam:GetRole"},
		UpdateActions: []string{},
	}
	actions := iamrules.Resolve(rule, iamrules.OpUpdate, nil)
	if len(actions) != 0 {
		t.Errorf("expected empty, got %v", actions)
	}
}

func TestResolve_ConditionalAction_Matched(t *testing.T) {
	rule := iamrules.Rule{
		BaseActions: []string{"s3:CreateBucket"},
		ConditionalActions: map[string][]string{
			"tags": {"s3:PutBucketTagging"},
		},
	}
	actions := iamrules.Resolve(rule, iamrules.OpCreate, []string{"tags"})
	sort.Strings(actions)
	if !reflect.DeepEqual(actions, []string{"s3:CreateBucket", "s3:PutBucketTagging"}) {
		t.Errorf("got %v", actions)
	}
}

func TestResolve_ConditionalAction_NotMatched(t *testing.T) {
	rule := iamrules.Rule{
		BaseActions: []string{"s3:CreateBucket"},
		ConditionalActions: map[string][]string{
			"tags": {"s3:PutBucketTagging"},
		},
	}
	actions := iamrules.Resolve(rule, iamrules.OpCreate, []string{"bucket"})
	if !reflect.DeepEqual(actions, []string{"s3:CreateBucket"}) {
		t.Errorf("got %v", actions)
	}
}

func TestResolve_Deduplication(t *testing.T) {
	rule := iamrules.Rule{
		BaseActions: []string{"iam:TagRole"},
		ConditionalActions: map[string][]string{
			"tags": {"iam:TagRole"},
		},
	}
	actions := iamrules.Resolve(rule, iamrules.OpCreate, []string{"tags"})
	if len(actions) != 1 || actions[0] != "iam:TagRole" {
		t.Errorf("expected deduplication, got %v", actions)
	}
}

func TestResolve_MultipleConditional(t *testing.T) {
	rule := iamrules.Rule{
		UpdateActions: []string{"iam:GetRole"},
		ConditionalActions: map[string][]string{
			"tags":        {"iam:TagRole", "iam:UntagRole"},
			"description": {"iam:UpdateRole"},
		},
	}
	actions := iamrules.Resolve(rule, iamrules.OpUpdate, []string{"tags", "description"})
	sort.Strings(actions)
	want := []string{"iam:GetRole", "iam:TagRole", "iam:UntagRole", "iam:UpdateRole"}
	if !reflect.DeepEqual(actions, want) {
		t.Errorf("got %v, want %v", actions, want)
	}
}

func TestResolve_IamRole_UpdateActions_HasGetRole(t *testing.T) {
	// aws_iam_role OpUpdate のバグ修正を確認：UpdateActionsに iam:GetRole が含まれること
	rule, ok := iamrules.Lookup("aws_iam_role", iamrules.OpUpdate)
	if !ok {
		t.Fatal("rule not found")
	}
	actions := iamrules.Resolve(rule, iamrules.OpUpdate, nil)
	for _, a := range actions {
		if a == "iam:GetRole" {
			return
		}
	}
	t.Errorf("iam:GetRole not found in update actions; got %v", actions)
}

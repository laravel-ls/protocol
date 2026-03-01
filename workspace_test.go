package protocol_test

import (
	"encoding/json"
	"testing"

	"github.com/laravel-ls/protocol"
)

func Test_Workspace_StructsUnmarshalValidJSON(t *testing.T) {
	var ws protocol.WorkspaceEdit
	if err := json.Unmarshal([]byte(`{
		"changes":{"file:///tmp/main.go":[{"range":{"start":{"line":0,"character":0},"end":{"line":0,"character":0}},"newText":"package main\n"}]},
		"changeAnnotations":{"a1":{"label":"Add package declaration","description":"Needed for file header","needsConfirmation":false}}
	}`), &ws); err != nil {
		t.Fatalf("unmarshal WorkspaceEdit failed: %v", err)
	}
	if len(ws.Changes) != 1 || len(ws.ChangeAnnotations) != 1 {
		t.Fatalf("unexpected WorkspaceEdit: %+v", ws)
	}

	var folder protocol.WorkspaceFolder
	if err := json.Unmarshal([]byte(`{"uri":"file:///workspace","name":"workspace"}`), &folder); err != nil {
		t.Fatalf("unmarshal WorkspaceFolder failed: %v", err)
	}
	if folder.Name != "workspace" {
		t.Fatalf("unexpected WorkspaceFolder: %+v", folder)
	}
}

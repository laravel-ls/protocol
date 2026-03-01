package protocol_test

import (
	"encoding/json"
	"testing"

	"github.com/laravel-ls/protocol"
)

func Test_DocumentCodeAction_StructsUnmarshalValidJSON(t *testing.T) {
	var params protocol.CodeActionParams
	if err := json.Unmarshal([]byte(`{
		"textDocument":{"uri":"file:///tmp/main.go"},
		"range":{"start":{"line":1,"character":0},"end":{"line":1,"character":4}},
		"context":{"diagnostics":[{"range":{"start":{"line":1,"character":0},"end":{"line":1,"character":4}},"message":"problem"}],"only":["quickfix"],"triggerKind":1}
	}`), &params); err != nil {
		t.Fatalf("unmarshal CodeActionParams failed: %v", err)
	}
	if len(params.Context.Diagnostics) != 1 {
		t.Fatalf("unexpected CodeActionParams diagnostics: %+v", params.Context.Diagnostics)
	}

	var action protocol.CodeAction
	if err := json.Unmarshal([]byte(`{"title":"Fix issue","kind":"quickfix","isPreferred":true,"disabled":{"reason":"not applicable"}}`), &action); err != nil {
		t.Fatalf("unmarshal CodeAction failed: %v", err)
	}
	if action.Disabled == nil || action.Disabled.Reason != "not applicable" {
		t.Fatalf("unexpected CodeAction disabled state: %+v", action.Disabled)
	}
}

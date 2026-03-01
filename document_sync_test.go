package protocol_test

import (
	"encoding/json"
	"testing"

	"github.com/laravel-ls/protocol"
)

func Test_DocumentSync_StructsUnmarshalValidJSON(t *testing.T) {
	var open protocol.DidOpenTextDocumentParams
	if err := json.Unmarshal([]byte(`{"textDocument":{"uri":"file:///tmp/main.go","languageId":"go","version":1,"text":"package main"}}`), &open); err != nil {
		t.Fatalf("unmarshal DidOpenTextDocumentParams failed: %v", err)
	}
	if open.TextDocument.URI == "" {
		t.Fatalf("unexpected DidOpenTextDocumentParams: %+v", open)
	}

	var change protocol.DidChangeTextDocumentParams
	if err := json.Unmarshal([]byte(`{"textDocument":{"uri":"file:///tmp/main.go","version":2},"contentChanges":[{"range":{"start":{"line":0,"character":0},"end":{"line":0,"character":0}},"text":"x"}]}`), &change); err != nil {
		t.Fatalf("unmarshal DidChangeTextDocumentParams failed: %v", err)
	}
	if len(change.ContentChanges) != 1 {
		t.Fatalf("unexpected DidChangeTextDocumentParams: %+v", change)
	}

	var save protocol.DidSaveTextDocumentParams
	if err := json.Unmarshal([]byte(`{"textDocument":{"uri":"file:///tmp/main.go"},"text":"content"}`), &save); err != nil {
		t.Fatalf("unmarshal DidSaveTextDocumentParams failed: %v", err)
	}
	if save.Text != "content" {
		t.Fatalf("unexpected DidSaveTextDocumentParams: %+v", save)
	}
}

package protocol_test

import (
	"encoding/json"
	"testing"

	"github.com/laravel-ls/protocol"
)

func Test_Document_StructsUnmarshalValidJSON(t *testing.T) {
	var id protocol.TextDocumentIdentifier
	if err := json.Unmarshal([]byte(`{"uri":"file:///tmp/main.go"}`), &id); err != nil {
		t.Fatalf("unmarshal TextDocumentIdentifier failed: %v", err)
	}
	if id.URI != "file:///tmp/main.go" {
		t.Fatalf("unexpected TextDocumentIdentifier: %+v", id)
	}

	var versioned protocol.VersionedTextDocumentIdentifier
	if err := json.Unmarshal([]byte(`{"uri":"file:///tmp/main.go","version":5}`), &versioned); err != nil {
		t.Fatalf("unmarshal VersionedTextDocumentIdentifier failed: %v", err)
	}
	if versioned.Version != 5 {
		t.Fatalf("unexpected VersionedTextDocumentIdentifier: %+v", versioned)
	}

	var item protocol.TextDocumentItem
	if err := json.Unmarshal([]byte(`{"uri":"file:///tmp/main.go","languageId":"go","version":5,"text":"package main"}`), &item); err != nil {
		t.Fatalf("unmarshal TextDocumentItem failed: %v", err)
	}
	if item.LanguageID != protocol.LanguageGo {
		t.Fatalf("unexpected TextDocumentItem language: %s", item.LanguageID)
	}

	var pos protocol.TextDocumentPositionParams
	if err := json.Unmarshal([]byte(`{"textDocument":{"uri":"file:///tmp/main.go"},"position":{"line":3,"character":1}}`), &pos); err != nil {
		t.Fatalf("unmarshal TextDocumentPositionParams failed: %v", err)
	}
	if pos.Position.Line != 3 {
		t.Fatalf("unexpected TextDocumentPositionParams: %+v", pos)
	}

	var edit protocol.TextEdit
	if err := json.Unmarshal([]byte(`{"range":{"start":{"line":1,"character":0},"end":{"line":1,"character":2}},"newText":"hi"}`), &edit); err != nil {
		t.Fatalf("unmarshal TextEdit failed: %v", err)
	}
	if edit.NewText != "hi" {
		t.Fatalf("unexpected TextEdit: %+v", edit)
	}

	var docEdit protocol.TextDocumentEdit
	if err := json.Unmarshal([]byte(`{"textDocument":{"uri":"file:///tmp/main.go","version":5},"edits":[{"range":{"start":{"line":0,"character":0},"end":{"line":0,"character":0}},"newText":"x"}]}`), &docEdit); err != nil {
		t.Fatalf("unmarshal TextDocumentEdit failed: %v", err)
	}
	if len(docEdit.Edits) != 1 {
		t.Fatalf("unexpected TextDocumentEdit edits: %+v", docEdit.Edits)
	}
}

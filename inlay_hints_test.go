package protocol_test

import (
	"encoding/json"
	"testing"

	"github.com/laravel-ls/protocol"
)

func Test_InlayHint_ParamsUnmarshalValidJSON(t *testing.T) {
	data := []byte(`{"textDocument":{"uri":"file:///tmp/main.go"},"range":{"start":{"line":1,"character":2},"end":{"line":1,"character":8}}}`)

	var params protocol.InlayHintParams
	if err := json.Unmarshal(data, &params); err != nil {
		t.Fatalf("unmarshal InlayHintParams failed: %v", err)
	}

	if params.TextDocument.URI != "file:///tmp/main.go" {
		t.Fatalf("unexpected textDocument URI: %q", params.TextDocument.URI)
	}

	if params.Range.Start.Line != 1 || params.Range.End.Character != 8 {
		t.Fatalf("unexpected range: %+v", params.Range)
	}
}

func Test_InlayHintLabel_UnmarshalString(t *testing.T) {
	var label protocol.InlayHintLabel
	if err := json.Unmarshal([]byte(`"x:"`), &label); err != nil {
		t.Fatalf("unmarshal InlayHintLabel string failed: %v", err)
	}

	if label.String == nil || *label.String != "x:" {
		t.Fatalf("expected string label 'x:', got %+v", label)
	}

	if label.Parts != nil {
		t.Fatalf("expected parts to be nil for string label")
	}
}

func Test_InlayHintLabel_UnmarshalParts(t *testing.T) {
	data := []byte(`[{"value":"name"},{"value":":","tooltip":"separator"}]`)

	var label protocol.InlayHintLabel
	if err := json.Unmarshal(data, &label); err != nil {
		t.Fatalf("unmarshal InlayHintLabel parts failed: %v", err)
	}

	if len(label.Parts) != 2 {
		t.Fatalf("expected 2 label parts, got %d", len(label.Parts))
	}

	if label.Parts[1].Tooltip == nil || label.Parts[1].Tooltip.String == nil || *label.Parts[1].Tooltip.String != "separator" {
		t.Fatalf("expected second part tooltip string to be set, got %+v", label.Parts[1].Tooltip)
	}
}

func Test_InlayHintTooltip_UnmarshalMarkupContent(t *testing.T) {
	data := []byte(`{"kind":"markdown","value":"**hint**"}`)

	var tooltip protocol.InlayHintTooltip
	if err := json.Unmarshal(data, &tooltip); err != nil {
		t.Fatalf("unmarshal InlayHintTooltip markup failed: %v", err)
	}

	if tooltip.MarkupContent == nil || tooltip.MarkupContent.Kind != protocol.MarkupKindMarkdown {
		t.Fatalf("expected markdown tooltip, got %+v", tooltip)
	}
}

func Test_InlayHint_UnmarshalTypedObject(t *testing.T) {
	data := []byte(`{
		"position": {"line": 2, "character": 10},
		"label": [{"value":"x"},{"value":":","tooltip":{"kind":"plaintext","value":"type separator"}}],
		"kind": 1,
		"textEdits": [{"range":{"start":{"line":2,"character":8},"end":{"line":2,"character":9}},"newText":"value"}],
		"tooltip": "inferred type",
		"paddingLeft": true,
		"paddingRight": false,
		"data": {"id": 42}
	}`)

	var hint protocol.InlayHint
	if err := json.Unmarshal(data, &hint); err != nil {
		t.Fatalf("unmarshal InlayHint failed: %v", err)
	}

	if hint.Position.Line != 2 || hint.Position.Character != 10 {
		t.Fatalf("unexpected position: %+v", hint.Position)
	}

	if hint.Kind == nil || *hint.Kind != protocol.InlayHintKindType {
		t.Fatalf("expected kind=InlayHintKindType, got %+v", hint.Kind)
	}

	if hint.Tooltip == nil || hint.Tooltip.String == nil || *hint.Tooltip.String != "inferred type" {
		t.Fatalf("expected tooltip string 'inferred type', got %+v", hint.Tooltip)
	}

	if len(hint.TextEdits) != 1 {
		t.Fatalf("expected 1 text edit, got %d", len(hint.TextEdits))
	}
}

func Test_InlayHintResponse_UnmarshalArrayAndNull(t *testing.T) {
	var list protocol.InlayHintResponse
	if err := json.Unmarshal([]byte(`[{"position":{"line":0,"character":1},"label":"x:"}]`), &list); err != nil {
		t.Fatalf("unmarshal InlayHintResponse array failed: %v", err)
	}

	if list.Null {
		t.Fatalf("expected non-null response for array")
	}

	if len(list.Hints) != 1 {
		t.Fatalf("expected 1 hint, got %d", len(list.Hints))
	}

	var nullRes protocol.InlayHintResponse
	if err := json.Unmarshal([]byte(`null`), &nullRes); err != nil {
		t.Fatalf("unmarshal InlayHintResponse null failed: %v", err)
	}

	if !nullRes.Null {
		t.Fatalf("expected null response flag to be true")
	}
}

func Test_InlayHintResponse_MarshalArrayAndNull(t *testing.T) {
	response := protocol.InlayHintResponse{
		Hints: []protocol.InlayHint{{
			Position: protocol.Position{Line: 0, Character: 0},
			Label: func() protocol.InlayHintLabel {
				label := "a"
				return protocol.InlayHintLabel{String: &label}
			}(),
		}},
	}

	data, err := json.Marshal(response)
	if err != nil {
		t.Fatalf("marshal InlayHintResponse array failed: %v", err)
	}

	if string(data) == "null" {
		t.Fatalf("expected array JSON, got null")
	}

	nullData, err := json.Marshal(protocol.InlayHintResponse{Null: true})
	if err != nil {
		t.Fatalf("marshal InlayHintResponse null failed: %v", err)
	}

	if string(nullData) != "null" {
		t.Fatalf("expected null JSON, got %s", string(nullData))
	}
}

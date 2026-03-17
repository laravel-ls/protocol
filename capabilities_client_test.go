package protocol_test

import (
	"encoding/json"
	"testing"

	"github.com/laravel-ls/protocol"
)

func Test_CapabilitiesClient_UnmarshalValidJSON(t *testing.T) {
	data := []byte(`{"window":{"showDocument":{"support":true}}}`)

	var decoded protocol.ClientCapabilities
	if err := json.Unmarshal(data, &decoded); err != nil {
		t.Fatalf("unmarshal failed: %v", err)
	}

	if decoded.Window == nil {
		t.Fatalf("expected window capabilities to be present")
	}

	if decoded.Window.ShowDocument == nil {
		t.Fatalf("expected showDocument capabilities to be present")
	}

	if !decoded.Window.ShowDocument.Support {
		t.Fatalf("expected showDocument support to be true")
	}
}

func Test_CapabilitiesClient_UnmarshalEmptyObject(t *testing.T) {
	var decoded protocol.ClientCapabilities
	if err := json.Unmarshal([]byte(`{}`), &decoded); err != nil {
		t.Fatalf("unmarshal failed: %v", err)
	}

	if decoded.Window != nil {
		t.Fatalf("expected window capabilities to be nil for empty object")
	}
}

func Test_CapabilitiesClient_MarshalOmitEmptyWindow(t *testing.T) {
	original := protocol.ClientCapabilities{}

	data, err := json.Marshal(original)
	if err != nil {
		t.Fatalf("marshal failed: %v", err)
	}

	var payload map[string]any
	if err := json.Unmarshal(data, &payload); err != nil {
		t.Fatalf("unmarshal failed: %v", err)
	}

	if _, ok := payload["window"]; ok {
		t.Fatalf("expected window to be omitted, got payload: %s", string(data))
	}
}

func Test_CapabilitiesClient_UnmarshalTopLevelCapabilities(t *testing.T) {
	data := []byte(`{
		"workspace": {"applyEdit": true},
		"textDocument": {"hover": {"dynamicRegistration": true}},
		"notebookDocument": {"synchronization": {"dynamicRegistration": true}},
		"general": {"positionEncodings": ["utf-16", "utf-8"]},
		"experimental": {"myFeature": true},
		"window": {"showDocument": {"support": true}}
	}`)

	var decoded protocol.ClientCapabilities
	if err := json.Unmarshal(data, &decoded); err != nil {
		t.Fatalf("unmarshal failed: %v", err)
	}

	if decoded.Workspace == nil {
		t.Fatalf("expected workspace capabilities to be present")
	}

	if decoded.TextDocument == nil {
		t.Fatalf("expected textDocument capabilities to be present")
	}

	if decoded.NotebookDocument == nil {
		t.Fatalf("expected notebookDocument capabilities to be present")
	}

	if decoded.General == nil {
		t.Fatalf("expected general capabilities to be present")
	}

	if decoded.Experimental == nil {
		t.Fatalf("expected experimental capabilities to be present")
	}

	if decoded.Window == nil || decoded.Window.ShowDocument == nil || !decoded.Window.ShowDocument.Support {
		t.Fatalf("expected window.showDocument.support to be true")
	}
}

func Test_CapabilitiesClient_MarshalOmitEmptyTopLevelFields(t *testing.T) {
	original := protocol.ClientCapabilities{}

	data, err := json.Marshal(original)
	if err != nil {
		t.Fatalf("marshal failed: %v", err)
	}

	var payload map[string]any
	if err := json.Unmarshal(data, &payload); err != nil {
		t.Fatalf("unmarshal failed: %v", err)
	}

	for _, key := range []string{"workspace", "textDocument", "notebookDocument", "general", "experimental", "window"} {
		if _, ok := payload[key]; ok {
			t.Fatalf("expected %q to be omitted, got payload: %s", key, string(data))
		}
	}
}

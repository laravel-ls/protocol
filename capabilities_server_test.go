package protocol_test

import (
	"encoding/json"
	"testing"

	"github.com/laravel-ls/protocol"
)

func Test_CapabilitiesServer_UnmarshalValidJSON(t *testing.T) {
	data := []byte(`{
		"positionEncoding": "utf-16",
		"textDocumentSync": 2,
		"completionProvider": {
			"triggerCharacters": [".", ":"],
			"resolveProvider": true,
			"completionItem": {
				"labelDetailsSupport": true
			}
		},
		"hoverProvider": true,
		"definitionProvider": true,
		"codeActionProvider": true,
		"signatureHelpProvider": {
			"triggerCharacters": ["(", ","],
			"retriggerCharacters": [")"]
		},
		"diagnosticProvider": {
			"interFileDependencies": true,
			"workspaceDiagnostics": true
		}
	}`)

	var decoded protocol.ServerCapabilities
	if err := json.Unmarshal(data, &decoded); err != nil {
		t.Fatalf("unmarshal failed: %v", err)
	}

	positionEncoding := protocol.PositionEncodingKindUTF16
	if decoded.PositionEncoding == nil || *decoded.PositionEncoding != positionEncoding {
		t.Fatalf("expected positionEncoding=%q, got %+v", positionEncoding, decoded.PositionEncoding)
	}

	if decoded.CompletionProvider == nil {
		t.Fatalf("expected completionProvider to be present")
	}

	if decoded.CompletionProvider.ResolveProvider == nil || !*decoded.CompletionProvider.ResolveProvider {
		t.Fatalf("expected completionProvider.resolveProvider=true")
	}

	if decoded.CompletionProvider.CompletionItem == nil {
		t.Fatalf("expected completionProvider.completionItem to be present")
	}

	if decoded.CompletionProvider.CompletionItem.LabelDetailsSupport == nil || !*decoded.CompletionProvider.CompletionItem.LabelDetailsSupport {
		t.Fatalf("expected completionProvider.completionItem.labelDetailsSupport=true")
	}

	textDocumentSync, ok := decoded.TextDocumentSync.(float64)
	if !ok || int(textDocumentSync) != int(protocol.TextDocumentSyncKindIncremental) {
		t.Fatalf("expected textDocumentSync=%d, got %#v", protocol.TextDocumentSyncKindIncremental, decoded.TextDocumentSync)
	}

	diagnosticProvider, ok := decoded.DiagnosticProvider.(map[string]any)
	if !ok {
		t.Fatalf("expected diagnosticProvider object, got %#v", decoded.DiagnosticProvider)
	}

	if inter, ok := diagnosticProvider["interFileDependencies"].(bool); !ok || !inter {
		t.Fatalf("expected diagnosticProvider.interFileDependencies=true, got %#v", diagnosticProvider["interFileDependencies"])
	}

	if ws, ok := diagnosticProvider["workspaceDiagnostics"].(bool); !ok || !ws {
		t.Fatalf("expected diagnosticProvider.workspaceDiagnostics=true, got %#v", diagnosticProvider["workspaceDiagnostics"])
	}
}

func Test_CapabilitiesServer_MarshalOmitEmptyFields(t *testing.T) {
	original := protocol.ServerCapabilities{}

	data, err := json.Marshal(original)
	if err != nil {
		t.Fatalf("marshal failed: %v", err)
	}

	var payload map[string]any
	if err := json.Unmarshal(data, &payload); err != nil {
		t.Fatalf("unmarshal failed: %v", err)
	}

	if len(payload) != 0 {
		t.Fatalf("expected empty JSON object from zero-value capabilities, got: %s", string(data))
	}
}

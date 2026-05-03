package protocol_test

import (
	"encoding/json"
	"testing"

	"github.com/laravel-ls/protocol"
)

func Test_DocumentDiagnostic_StructsUnmarshalValidJSON(t *testing.T) {
	var params protocol.DocumentDiagnosticParams
	if err := json.Unmarshal([]byte(`{"textDocument":{"uri":"file:///tmp/main.go"},"identifier":"go","previousResultId":"r1"}`), &params); err != nil {
		t.Fatalf("unmarshal DocumentDiagnosticParams failed: %v", err)
	}
	if params.Identifier != "go" {
		t.Fatalf("unexpected DocumentDiagnosticParams: %+v", params)
	}

	var report protocol.DocumentDiagnosticReport
	if err := json.Unmarshal([]byte(`{"kind":"full","resultId":"r2","items":[{"range":{"start":{"line":0,"character":0},"end":{"line":0,"character":1}},"message":"m"}]}`), &report); err != nil {
		t.Fatalf("unmarshal DocumentDiagnosticReport failed: %v", err)
	}
	if report.Full == nil || report.Full.Kind != "full" {
		t.Fatalf("unexpected DocumentDiagnosticReport: %+v", report)
	}
}

func Test_PublishDiagnosticsParams_UnmarshalValidJSON(t *testing.T) {
	var params protocol.PublishDiagnosticsParams
	if err := json.Unmarshal([]byte(`{
		"uri":"file:///tmp/main.go",
		"version":2,
		"diagnostics":[{"range":{"start":{"line":1,"character":1},"end":{"line":1,"character":3}},"message":"example diagnostic"}]
	}`), &params); err != nil {
		t.Fatalf("unmarshal PublishDiagnosticsParams failed: %v", err)
	}

	if params.URI != "file:///tmp/main.go" {
		t.Fatalf("unexpected URI: %q", params.URI)
	}

	if params.Version != 2 {
		t.Fatalf("unexpected Version: %d", params.Version)
	}

	if len(params.Diagnostics) != 1 || params.Diagnostics[0].Message != "example diagnostic" {
		t.Fatalf("unexpected Diagnostics: %+v", params.Diagnostics)
	}
}

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

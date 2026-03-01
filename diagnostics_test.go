package protocol_test

import (
	"encoding/json"
	"testing"

	"github.com/laravel-ls/protocol"
)

func Test_Diagnostics_StructsUnmarshalValidJSON(t *testing.T) {
	var diag protocol.Diagnostic
	if err := json.Unmarshal([]byte(`{
		"range":{"start":{"line":1,"character":1},"end":{"line":1,"character":3}},
		"severity":1,
		"code":"E100",
		"codeDescription":{"href":"https://example.com/E100"},
		"source":"golangci-lint",
		"message":"example diagnostic",
		"tags":[1],
		"relatedInformation":[{"location":{"uri":"file:///tmp/main.go","range":{"start":{"line":0,"character":0},"end":{"line":0,"character":1}}},"message":"related"}],
		"data":{"k":"v"}
	}`), &diag); err != nil {
		t.Fatalf("unmarshal Diagnostic failed: %v", err)
	}
	if diag.Message != "example diagnostic" || diag.CodeDescription == nil {
		t.Fatalf("unexpected Diagnostic: %+v", diag)
	}
}

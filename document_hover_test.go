package protocol_test

import (
	"encoding/json"
	"testing"

	"github.com/laravel-ls/protocol"
)

func Test_DocumentHover_ParamsUnmarshalValidJSON(t *testing.T) {
	var hover protocol.HoverParams
	if err := json.Unmarshal([]byte(`{"textDocument":{"uri":"file:///tmp/main.go"},"position":{"line":1,"character":1}}`), &hover); err != nil {
		t.Fatalf("unmarshal HoverParams failed: %v", err)
	}
	if hover.TextDocument.URI == "" {
		t.Fatalf("unexpected HoverParams: %+v", hover)
	}
}

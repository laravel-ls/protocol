package protocol_test

import (
	"encoding/json"
	"testing"

	"github.com/laravel-ls/protocol"
)

func Test_DocumentCompletion_ParamsUnmarshalValidJSON(t *testing.T) {
	var completion protocol.CompletionParams
	if err := json.Unmarshal([]byte(`{"textDocument":{"uri":"file:///tmp/main.go"},"position":{"line":1,"character":1},"context":{"triggerKind":2,"triggerCharacter":"."}}`), &completion); err != nil {
		t.Fatalf("unmarshal CompletionParams failed: %v", err)
	}
	if completion.Context == nil || completion.Context.TriggerCharacter != "." {
		t.Fatalf("unexpected CompletionParams context: %+v", completion.Context)
	}
}

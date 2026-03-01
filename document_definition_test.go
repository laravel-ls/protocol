package protocol_test

import (
	"encoding/json"
	"testing"

	"github.com/laravel-ls/protocol"
)

func Test_DocumentDefinition_ParamsUnmarshalValidJSON(t *testing.T) {
	var def protocol.DefinitionParams
	if err := json.Unmarshal([]byte(`{"textDocument":{"uri":"file:///tmp/main.go"},"position":{"line":1,"character":1}}`), &def); err != nil {
		t.Fatalf("unmarshal DefinitionParams failed: %v", err)
	}
	if def.TextDocument.URI != "file:///tmp/main.go" {
		t.Fatalf("unexpected DefinitionParams: %+v", def)
	}
}

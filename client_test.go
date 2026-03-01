package protocol_test

import (
	"encoding/json"
	"testing"

	"github.com/laravel-ls/protocol"
)

func Test_Client_StructsUnmarshalValidJSON(t *testing.T) {
	var clientInfo protocol.ClientInfo
	if err := json.Unmarshal([]byte(`{"name":"nvim","version":"0.10"}`), &clientInfo); err != nil {
		t.Fatalf("unmarshal ClientInfo failed: %v", err)
	}
	if clientInfo.Name != "nvim" {
		t.Fatalf("unexpected ClientInfo: %+v", clientInfo)
	}

	var command protocol.Command
	if err := json.Unmarshal([]byte(`{"title":"Fix","command":"app.fix","arguments":["a",1]}`), &command); err != nil {
		t.Fatalf("unmarshal Command failed: %v", err)
	}
	if command.Command != "app.fix" || len(command.Arguments) != 2 {
		t.Fatalf("unexpected Command: %+v", command)
	}
}

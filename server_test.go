package protocol_test

import (
	"encoding/json"
	"testing"

	"github.com/laravel-ls/protocol"
)

func Test_Server_InfoUnmarshalValidJSON(t *testing.T) {
	var serverInfo protocol.ServerInfo
	if err := json.Unmarshal([]byte(`{"name":"my-ls","version":"0.0.1"}`), &serverInfo); err != nil {
		t.Fatalf("unmarshal ServerInfo failed: %v", err)
	}
	if serverInfo.Name != "my-ls" {
		t.Fatalf("unexpected ServerInfo: %+v", serverInfo)
	}
}

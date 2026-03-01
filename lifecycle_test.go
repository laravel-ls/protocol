package protocol_test

import (
	"encoding/json"
	"testing"

	"github.com/laravel-ls/protocol"
)

func Test_Lifecycle_StructsUnmarshalValidJSON(t *testing.T) {
	var initParams protocol.InitializeParams
	if err := json.Unmarshal([]byte(`{
		"processId":123,
		"clientInfo":{"name":"vscode","version":"1.0"},
		"locale":"en-US",
		"rootPath":"/workspace",
		"rootUri":"file:///workspace",
		"initializationOptions":{"feature":true},
		"capabilities":{"window":{"showDocument":{"support":true}}},
		"trace":"messages",
		"workspaceFolders":[{"uri":"file:///workspace","name":"workspace"}]
	}`), &initParams); err != nil {
		t.Fatalf("unmarshal InitializeParams failed: %v", err)
	}
	if initParams.ClientInfo == nil || initParams.ClientInfo.Name != "vscode" {
		t.Fatalf("unexpected InitializeParams clientInfo: %+v", initParams.ClientInfo)
	}
	if initParams.Capabilities.Window == nil || initParams.Capabilities.Window.ShowDocument == nil || !initParams.Capabilities.Window.ShowDocument.Support {
		t.Fatalf("unexpected InitializeParams capabilities: %+v", initParams.Capabilities)
	}

	var initResult protocol.InitializeResult
	if err := json.Unmarshal([]byte(`{"capabilities":{"hoverProvider":true},"serverInfo":{"name":"ls","version":"0.1"}}`), &initResult); err != nil {
		t.Fatalf("unmarshal InitializeResult failed: %v", err)
	}
	if initResult.ServerInfo == nil || initResult.ServerInfo.Name != "ls" {
		t.Fatalf("unexpected InitializeResult: %+v", initResult)
	}

	var cancel protocol.CancelParams
	if err := json.Unmarshal([]byte(`{"id":42}`), &cancel); err != nil {
		t.Fatalf("unmarshal CancelParams failed: %v", err)
	}
	if cancel.Id.String() != "42" {
		t.Fatalf("unexpected CancelParams id: %s", cancel.Id.String())
	}
}

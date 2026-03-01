package protocol_test

import (
	"encoding/json"
	"testing"

	"github.com/laravel-ls/protocol"
)

func Test_Window_StructsUnmarshalValidJSON(t *testing.T) {
	var show protocol.ShowMessageParams
	if err := json.Unmarshal([]byte(`{"type":2,"message":"Heads up"}`), &show); err != nil {
		t.Fatalf("unmarshal ShowMessageParams failed: %v", err)
	}
	if show.Type != protocol.MessageTypeWarning || show.Message != "Heads up" {
		t.Fatalf("unexpected ShowMessageParams: %+v", show)
	}

	var request protocol.ShowMessageRequestParams
	if err := json.Unmarshal([]byte(`{"type":1,"message":"Retry?","actions":[{"title":"Retry"}]}`), &request); err != nil {
		t.Fatalf("unmarshal ShowMessageRequestParams failed: %v", err)
	}
	if len(request.Actions) != 1 || request.Actions[0].Title != "Retry" {
		t.Fatalf("unexpected ShowMessageRequestParams actions: %+v", request.Actions)
	}

	var log protocol.LogMessageParams
	if err := json.Unmarshal([]byte(`{"type":4,"message":"log line"}`), &log); err != nil {
		t.Fatalf("unmarshal LogMessageParams failed: %v", err)
	}
	if log.Type != protocol.MessageTypeLog {
		t.Fatalf("unexpected LogMessageParams type: %v", log.Type)
	}

	var showDoc protocol.ShowDocumentParams
	if err := json.Unmarshal([]byte(`{"uri":"file:///tmp/a.go","takeFocus":true,"external":false,"selection":{"start":{"line":1,"character":2},"end":{"line":1,"character":6}}}`), &showDoc); err != nil {
		t.Fatalf("unmarshal ShowDocumentParams failed: %v", err)
	}
	if showDoc.URI != "file:///tmp/a.go" || showDoc.Selection == nil {
		t.Fatalf("unexpected ShowDocumentParams: %+v", showDoc)
	}

	var result protocol.ShowDocumentResult
	if err := json.Unmarshal([]byte(`{"success":true}`), &result); err != nil {
		t.Fatalf("unmarshal ShowDocumentResult failed: %v", err)
	}
	if !result.Success {
		t.Fatalf("unexpected ShowDocumentResult: %+v", result)
	}
}

package protocol_test

import (
	"encoding/json"
	"testing"

	"github.com/laravel-ls/protocol"
)

func Test_Progress_StructsUnmarshalValidJSON(t *testing.T) {
	var workDone protocol.WorkDoneProgressParams
	if err := json.Unmarshal([]byte(`{"workDoneToken":"wd-1"}`), &workDone); err != nil {
		t.Fatalf("unmarshal WorkDoneProgressParams failed: %v", err)
	}
	if workDone.WorkDoneToken == nil {
		t.Fatalf("expected workDoneToken to be set")
	}

	var partial protocol.PartialResultParams
	if err := json.Unmarshal([]byte(`{"partialResultToken":99}`), &partial); err != nil {
		t.Fatalf("unmarshal PartialResultParams failed: %v", err)
	}
	if partial.PartialResultToken == nil {
		t.Fatalf("expected partialResultToken to be set")
	}

	var create protocol.WorkDoneProgressCreateParams
	if err := json.Unmarshal([]byte(`{"token":"create-token"}`), &create); err != nil {
		t.Fatalf("unmarshal WorkDoneProgressCreateParams failed: %v", err)
	}

	var cancel protocol.WorkDoneProgressCancelParams
	if err := json.Unmarshal([]byte(`{"token":7}`), &cancel); err != nil {
		t.Fatalf("unmarshal WorkDoneProgressCancelParams failed: %v", err)
	}
}

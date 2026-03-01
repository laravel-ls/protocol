package protocol_test

import (
	"encoding/json"
	"testing"

	"github.com/laravel-ls/protocol"
)

func Test_DocumentChangeOperation_StructsUnmarshalValidJSON(t *testing.T) {
	var op protocol.ResourceOperation
	if err := json.Unmarshal([]byte(`{"kind":"create","annotationId":"a1"}`), &op); err != nil {
		t.Fatalf("unmarshal ResourceOperation failed: %v", err)
	}
	if op.Kind != "create" {
		t.Fatalf("unexpected ResourceOperation: %+v", op)
	}

	var create protocol.CreateFile
	if err := json.Unmarshal([]byte(`{"kind":"create","uri":"file:///tmp/new.go","options":{"overwrite":true}}`), &create); err != nil {
		t.Fatalf("unmarshal CreateFile failed: %v", err)
	}
	if create.Options == nil || !create.Options.Overwrite {
		t.Fatalf("unexpected CreateFile: %+v", create)
	}

	var rename protocol.RenameFile
	if err := json.Unmarshal([]byte(`{"kind":"rename","oldUri":"file:///tmp/a.go","newUri":"file:///tmp/b.go","options":{"ignoreIfExists":true}}`), &rename); err != nil {
		t.Fatalf("unmarshal RenameFile failed: %v", err)
	}
	if rename.Options == nil || !rename.Options.IgnoreIfExists {
		t.Fatalf("unexpected RenameFile: %+v", rename)
	}

	var del protocol.DeleteFile
	if err := json.Unmarshal([]byte(`{"kind":"delete","uri":"file:///tmp/old.go","options":{"recursive":true,"ignoreIfNotExists":true}}`), &del); err != nil {
		t.Fatalf("unmarshal DeleteFile failed: %v", err)
	}
	if del.Options == nil || !del.Options.Recursive || !del.Options.IgnoreIfNotExists {
		t.Fatalf("unexpected DeleteFile: %+v", del)
	}
}

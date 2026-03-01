package protocol_test

import (
	"encoding/json"
	"testing"

	"github.com/laravel-ls/protocol"
)

func Test_Base_PositionJSONRoundTrip(t *testing.T) {
	original := protocol.Position{Line: 12, Character: 34}

	data, err := json.Marshal(original)
	if err != nil {
		t.Fatalf("marshal failed: %v", err)
	}

	var decoded protocol.Position
	if err := json.Unmarshal(data, &decoded); err != nil {
		t.Fatalf("unmarshal failed: %v", err)
	}

	if decoded.Line != original.Line || decoded.Character != original.Character {
		t.Fatalf("expected %+v, got %+v", original, decoded)
	}
}

func Test_Base_LocationLinkJSONRoundTrip(t *testing.T) {
	original := protocol.LocationLink{
		OriginSelectionRange: &protocol.Range{
			Start: protocol.Position{Line: 1, Character: 2},
			End:   protocol.Position{Line: 1, Character: 5},
		},
		TargetURI: protocol.DocumentURI("file:///tmp/target.go"),
		TargetRange: protocol.Range{
			Start: protocol.Position{Line: 10, Character: 0},
			End:   protocol.Position{Line: 20, Character: 0},
		},
		TargetSelectionRange: protocol.Range{
			Start: protocol.Position{Line: 12, Character: 4},
			End:   protocol.Position{Line: 12, Character: 10},
		},
	}

	data, err := json.Marshal(original)
	if err != nil {
		t.Fatalf("marshal failed: %v", err)
	}

	var decoded protocol.LocationLink
	if err := json.Unmarshal(data, &decoded); err != nil {
		t.Fatalf("unmarshal failed: %v", err)
	}

	if decoded.TargetURI != original.TargetURI {
		t.Fatalf("expected target URI %q, got %q", original.TargetURI, decoded.TargetURI)
	}

	if decoded.OriginSelectionRange == nil {
		t.Fatalf("expected origin selection range to be present")
	}

	if decoded.TargetRange.Start.Line != original.TargetRange.Start.Line || decoded.TargetRange.End.Line != original.TargetRange.End.Line {
		t.Fatalf("expected target range %+v, got %+v", original.TargetRange, decoded.TargetRange)
	}
}

func Test_Base_ChangeAnnotationJSONRoundTrip(t *testing.T) {
	needsConfirmation := true
	original := protocol.ChangeAnnotation{
		Label:             "Refactor imports",
		NeedsConfirmation: &needsConfirmation,
		Description:       "Reorders and removes unused imports",
	}

	data, err := json.Marshal(original)
	if err != nil {
		t.Fatalf("marshal failed: %v", err)
	}

	var decoded protocol.ChangeAnnotation
	if err := json.Unmarshal(data, &decoded); err != nil {
		t.Fatalf("unmarshal failed: %v", err)
	}

	if decoded.Label != original.Label {
		t.Fatalf("expected label %q, got %q", original.Label, decoded.Label)
	}

	if decoded.Description != original.Description {
		t.Fatalf("expected description %q, got %q", original.Description, decoded.Description)
	}

	if decoded.NeedsConfirmation == nil || *decoded.NeedsConfirmation != *original.NeedsConfirmation {
		t.Fatalf("expected needsConfirmation=%v, got %+v", *original.NeedsConfirmation, decoded.NeedsConfirmation)
	}
}

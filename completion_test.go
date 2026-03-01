package protocol_test

import (
	"encoding/json"
	"testing"

	"github.com/laravel-ls/protocol"
)

func Test_Completion_StructsUnmarshalValidJSON(t *testing.T) {
	var item protocol.CompletionItem
	if err := json.Unmarshal([]byte(`{"label":"fmt.Println","kind":3,"detail":"func","documentation":"prints","insertText":"fmt.Println($1)","textEditText":"fmt.Println","commitCharacters":";"}`), &item); err != nil {
		t.Fatalf("unmarshal CompletionItem failed: %v", err)
	}
	if item.Label != "fmt.Println" {
		t.Fatalf("unexpected CompletionItem: %+v", item)
	}

	var list protocol.CompletionList
	if err := json.Unmarshal([]byte(`{"isIncomplete":false,"items":[{"label":"x"}]}`), &list); err != nil {
		t.Fatalf("unmarshal CompletionList failed: %v", err)
	}
	if len(list.Items) != 1 {
		t.Fatalf("unexpected CompletionList: %+v", list)
	}
}

package protocol

import (
	"encoding/json"
	"errors"
)

const (
	// MethodTextDocumentInlayHint method name of `textDocument/inlayHint`.
	MethodTextDocumentInlayHint = "textDocument/inlayHint"
)

// InlayHintParams - Parameters for a `textDocument/inlayHint` request.
//
// @since 3.17.0
//
// See https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification/#inlayHintParams
type InlayHintParams struct {
	WorkDoneProgressParams
	PartialResultParams

	// The text document.
	TextDocument TextDocumentIdentifier `json:"textDocument"`

	// The visible range for which inlay hints should be computed.
	Range Range `json:"range"`
}

// InlayHintKind - The kind of an inlay hint.
//
// @since 3.17.0
//
// See https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification/#inlayHintKind
type InlayHintKind int

const (
	// InlayHintKindType - An inlay hint that shows a type annotation.
	InlayHintKindType InlayHintKind = 1

	// InlayHintKindParameter - An inlay hint that shows a parameter name.
	InlayHintKindParameter InlayHintKind = 2
)

// InlayHintTooltip can be a plain string or a MarkupContent object.
//
// @since 3.17.0
type InlayHintTooltip struct {
	String        *string
	MarkupContent *MarkupContent
}

func (t InlayHintTooltip) MarshalJSON() ([]byte, error) {
	if t.String != nil {
		return json.Marshal(*t.String)
	}
	if t.MarkupContent != nil {
		return json.Marshal(t.MarkupContent)
	}
	return []byte("null"), nil
}

func (t *InlayHintTooltip) UnmarshalJSON(data []byte) error {
	*t = InlayHintTooltip{}

	if string(data) == "null" {
		return nil
	}

	var str string
	if err := json.Unmarshal(data, &str); err == nil {
		t.String = &str
		return nil
	}

	var markup MarkupContent
	if err := json.Unmarshal(data, &markup); err == nil && markup.Kind != "" {
		t.MarkupContent = &markup
		return nil
	}

	return errors.New("invalid InlayHintTooltip: not string, MarkupContent, or null")
}

// InlayHintLabelPart - A segment of an inlay hint label.
//
// @since 3.17.0
//
// See https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification/#inlayHintLabelPart
type InlayHintLabelPart struct {
	// The mandatory label value.
	Value string `json:"value"`

	// The tooltip text or markup shown when hovering this label part.
	Tooltip *InlayHintTooltip `json:"tooltip,omitempty"`

	// A source location for this label part.
	Location *Location `json:"location,omitempty"`

	// A command associated with this label part.
	Command *Command `json:"command,omitempty"`
}

// InlayHintLabel can be either a string or a list of label parts.
//
// @since 3.17.0
type InlayHintLabel struct {
	String *string
	Parts  []InlayHintLabelPart
}

func (l InlayHintLabel) MarshalJSON() ([]byte, error) {
	if l.String != nil {
		return json.Marshal(*l.String)
	}
	if l.Parts != nil {
		return json.Marshal(l.Parts)
	}
	return nil, errors.New("one of InlayHintLabel.String or InlayHintLabel.Parts needs to be set")
}

func (l *InlayHintLabel) UnmarshalJSON(data []byte) error {
	*l = InlayHintLabel{}

	var str string
	if err := json.Unmarshal(data, &str); err == nil {
		l.String = &str
		return nil
	}

	var parts []InlayHintLabelPart
	if err := json.Unmarshal(data, &parts); err == nil {
		l.Parts = parts
		return nil
	}

	return errors.New("invalid InlayHintLabel: not string or []InlayHintLabelPart")
}

// InlayHint represents an inlay hint item.
//
// @since 3.17.0
//
// See https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification/#inlayHint
type InlayHint struct {
	// The position of this hint.
	Position Position `json:"position"`

	// The label of this hint. A human readable string or an array of label parts.
	Label InlayHintLabel `json:"label"`

	// The kind of this hint.
	Kind *InlayHintKind `json:"kind,omitempty"`

	// Optional text edits that are performed when accepting this hint.
	TextEdits []TextEdit `json:"textEdits,omitempty"`

	// The tooltip text when hovering over this hint.
	Tooltip *InlayHintTooltip `json:"tooltip,omitempty"`

	// Render padding before this hint.
	PaddingLeft *bool `json:"paddingLeft,omitempty"`

	// Render padding after this hint.
	PaddingRight *bool `json:"paddingRight,omitempty"`

	// A data entry field preserved between a hint request and resolve request.
	Data LSPAny `json:"data,omitempty"`
}

// InlayHintResponse - Result for a `textDocument/inlayHint` request.
//
// It is either an array of `InlayHint` or `null`.
//
// @since 3.17.0
//
// See https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification/#inlayHint
type InlayHintResponse struct {
	Hints []InlayHint
	Null  bool
}

func (r InlayHintResponse) MarshalJSON() ([]byte, error) {
	if r.Null {
		return []byte("null"), nil
	}
	if r.Hints == nil {
		return []byte("null"), nil
	}
	return json.Marshal(r.Hints)
}

func (r *InlayHintResponse) UnmarshalJSON(data []byte) error {
	*r = InlayHintResponse{}

	if string(data) == "null" {
		r.Null = true
		return nil
	}

	var hints []InlayHint
	if err := json.Unmarshal(data, &hints); err == nil {
		r.Hints = hints
		return nil
	}

	return errors.New("invalid inlay hint response: not null or []InlayHint")
}

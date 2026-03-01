package protocol

const (
	MethodWindowLogMessage         = "window/logMessage"
	MethodWindowShowMessage        = "window/showMessage"
	MethodWindowShowMessageRequest = "window/showMessageRequest"
	MethodWindowShowDocument       = "window/showDocument"
)

// ShowMessageParams - The parameters of a show message notification.
//
// See https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification/#showMessageParams
type ShowMessageParams struct {
	// The message type. See `MessageType`.
	Type MessageType `json:"type"`

	// The actual message.
	Message string `json:"message"`
}

// ShowMessageRequestParams - The parameters of a show message request.
//
// See https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification/#showMessageRequestParams
type ShowMessageRequestParams struct {
	ShowMessageParams

	// The message action items to present.
	Actions []MessageActionItem `json:"actions,omitempty"`
}

// MessageActionItem - A response item for a show message request.
//
// See https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification/#messageActionItem
type MessageActionItem struct {
	// A short title like 'Retry', 'Open Log' etc.
	Title string `json:"title"`
}

// LogMessageParams - The parameters of a log message notification.
//
// See https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification/#logMessageParams
type LogMessageParams struct {
	ShowMessageParams
}

// ShowDocumentParams - Parameters for a show document request.
//
// See https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification/#showDocumentParams
//
// @since 3.16.0
type ShowDocumentParams struct {
	// The URI to show.
	URI DocumentURI `json:"uri"`

	// Indicates if the editor should take focus.
	TakeFocus *bool `json:"takeFocus,omitempty"`

	// Indicates if the URI should be opened in an external program.
	External *bool `json:"external,omitempty"`

	// An optional selection range if the document is a text document.
	Selection *Range `json:"selection,omitempty"`
}

// ShowDocumentResult - Result of a show document request.
//
// See https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification/#showDocumentResult
//
// @since 3.16.0
type ShowDocumentResult struct {
	// Indicates whether the document was shown.
	Success bool `json:"success"`
}

// MessageType - The message type.
//
// See https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification/#messageType
type MessageType int

const (
	// An error message.
	MessageTypeError MessageType = 1

	// A warning message.
	MessageTypeWarning MessageType = 2

	// An information message.
	MessageTypeInfo MessageType = 3

	// A log message.
	MessageTypeLog MessageType = 4
)

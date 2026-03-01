package protocol

// ShowDocumentClientCapabilities - Client capabilities for the show document request.
//
// See https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification/#showDocumentClientCapabilities
//
// @since 3.16.0
type ShowDocumentClientCapabilities struct {
	// The client has support for the show document request.
	Support bool `json:"support"`
}

// WindowClientCapabilities - Client capabilities for window features.
//
// See https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification/#windowClientCapabilities
type WindowClientCapabilities struct {
	// It indicates whether the client supports the `window/showDocument` request.
	//
	// @since 3.16.0
	ShowDocument *ShowDocumentClientCapabilities `json:"showDocument,omitempty"`
}

// ClientCapabilities defines the capabilities of the client (e.g., editor or IDE).
// It tells the language server what features the client supports.
type ClientCapabilities struct {
	// Window specific client capabilities.
	Window *WindowClientCapabilities `json:"window,omitempty"`
}

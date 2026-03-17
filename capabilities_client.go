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

// ShowMessageRequestClientCapabilities - Capabilities for `window/showMessageRequest`.
//
// @since 3.16.0
type ShowMessageRequestClientCapabilities struct {
	// Capabilities specific to the message action item.
	MessageActionItem *ShowMessageRequestClientCapabilitiesMessageActionItem `json:"messageActionItem,omitempty"`
}

// ShowMessageRequestClientCapabilitiesMessageActionItem - Capabilities specific to message action items.
//
// @since 3.16.0
type ShowMessageRequestClientCapabilitiesMessageActionItem struct {
	// Whether the client supports additional attributes that are preserved
	// and sent back to the server in the response.
	AdditionalPropertiesSupport bool `json:"additionalPropertiesSupport,omitempty"`
}

// WindowClientCapabilities - Client capabilities for window features.
//
// See https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification/#windowClientCapabilities
type WindowClientCapabilities struct {
	// Whether client supports handling progress notifications from the server.
	WorkDoneProgress bool `json:"workDoneProgress,omitempty"`

	// Capabilities for the show message request.
	//
	// @since 3.16.0
	ShowMessage *ShowMessageRequestClientCapabilities `json:"showMessage,omitempty"`

	// It indicates whether the client supports the `window/showDocument` request.
	//
	// @since 3.16.0
	ShowDocument *ShowDocumentClientCapabilities `json:"showDocument,omitempty"`
}

// ResourceOperationKind - Resource operation kinds in workspace edits.
type ResourceOperationKind string

const (
	// ResourceOperationCreate indicates support for creating files/folders.
	ResourceOperationCreate ResourceOperationKind = "create"
	// ResourceOperationRename indicates support for renaming files/folders.
	ResourceOperationRename ResourceOperationKind = "rename"
	// ResourceOperationDelete indicates support for deleting files/folders.
	ResourceOperationDelete ResourceOperationKind = "delete"
)

// FailureHandlingKind - Strategies for handling workspace edit failures.
type FailureHandlingKind string

const (
	// FailureHandlingAbort applies the edit up to the first failure and aborts.
	FailureHandlingAbort FailureHandlingKind = "abort"
	// FailureHandlingTransactional applies all edits transactionally.
	FailureHandlingTransactional FailureHandlingKind = "transactional"
	// FailureHandlingTextOnlyTransactional applies text edits transactionally.
	FailureHandlingTextOnlyTransactional FailureHandlingKind = "textOnlyTransactional"
	// FailureHandlingUndo rolls back already-applied edits on failure.
	FailureHandlingUndo FailureHandlingKind = "undo"
)

// DynamicRegistrationClientCapabilities captures capabilities with `dynamicRegistration`.
type DynamicRegistrationClientCapabilities struct {
	// DynamicRegistration indicates whether the client supports dynamic
	// registration for the corresponding request/notification.
	DynamicRegistration bool `json:"dynamicRegistration,omitempty"`
}

type (
	DidChangeConfigurationClientCapabilities   = DynamicRegistrationClientCapabilities
	DidChangeWatchedFilesClientCapabilities    = DynamicRegistrationClientCapabilities
	ExecuteCommandClientCapabilities           = DynamicRegistrationClientCapabilities
	ReferenceClientCapabilities                = DynamicRegistrationClientCapabilities
	DocumentHighlightClientCapabilities        = DynamicRegistrationClientCapabilities
	CodeLensClientCapabilities                 = DynamicRegistrationClientCapabilities
	DocumentColorClientCapabilities            = DynamicRegistrationClientCapabilities
	DocumentFormattingClientCapabilities       = DynamicRegistrationClientCapabilities
	DocumentRangeFormattingClientCapabilities  = DynamicRegistrationClientCapabilities
	DocumentOnTypeFormattingClientCapabilities = DynamicRegistrationClientCapabilities
	SelectionRangeClientCapabilities           = DynamicRegistrationClientCapabilities
	CallHierarchyClientCapabilities            = DynamicRegistrationClientCapabilities
	LinkedEditingRangeClientCapabilities       = DynamicRegistrationClientCapabilities
	MonikerClientCapabilities                  = DynamicRegistrationClientCapabilities
	TypeHierarchyClientCapabilities            = DynamicRegistrationClientCapabilities
	InlineValueClientCapabilities              = DynamicRegistrationClientCapabilities
)

// WorkspaceEditClientCapabilities - Client capabilities for workspace edits.
type WorkspaceEditClientCapabilities struct {
	// DocumentChanges indicates support for versioned document changes.
	DocumentChanges bool `json:"documentChanges,omitempty"`
	// ResourceOperations is the set of resource operations supported in
	// `WorkspaceEdit.documentChanges`.
	ResourceOperations []ResourceOperationKind `json:"resourceOperations,omitempty"`
	// FailureHandling describes how the client handles failures in
	// workspace edit application.
	FailureHandling FailureHandlingKind `json:"failureHandling,omitempty"`
	// NormalizesLineEndings indicates whether the client normalizes line endings
	// when applying text edits.
	NormalizesLineEndings bool `json:"normalizesLineEndings,omitempty"`
	// ChangeAnnotationSupport describes support for change annotations in
	// workspace edits.
	ChangeAnnotationSupport *WorkspaceEditClientCapabilitiesChangeAnnotationSupport `json:"changeAnnotationSupport,omitempty"`
}

// WorkspaceEditClientCapabilitiesChangeAnnotationSupport - Capabilities for change annotations.
type WorkspaceEditClientCapabilitiesChangeAnnotationSupport struct {
	// GroupsOnLabel indicates whether the client can present change annotations
	// grouped by their label.
	GroupsOnLabel bool `json:"groupsOnLabel,omitempty"`
}

// WorkspaceSymbolClientCapabilitiesSymbolKind - Supported workspace symbol kinds.
type WorkspaceSymbolClientCapabilitiesSymbolKind struct {
	// ValueSet is the symbol kind values the client supports.
	ValueSet []uint32 `json:"valueSet,omitempty"`
}

// WorkspaceSymbolClientCapabilitiesTagSupport - Supported workspace symbol tags.
type WorkspaceSymbolClientCapabilitiesTagSupport struct {
	// ValueSet is the symbol tag values the client supports.
	ValueSet []uint32 `json:"valueSet,omitempty"`
}

// WorkspaceSymbolClientCapabilities - Workspace symbol capabilities.
type WorkspaceSymbolClientCapabilities struct {
	// DynamicRegistration indicates whether workspace symbol support can be
	// dynamically registered.
	DynamicRegistration bool `json:"dynamicRegistration,omitempty"`
	// SymbolKind describes the symbol kinds the client supports.
	SymbolKind *WorkspaceSymbolClientCapabilitiesSymbolKind `json:"symbolKind,omitempty"`
	// TagSupport describes supported `SymbolTag` values.
	TagSupport *WorkspaceSymbolClientCapabilitiesTagSupport `json:"tagSupport,omitempty"`
	// ResolveSupport lists properties that can be resolved lazily.
	ResolveSupport *WorkspaceSymbolClientCapabilitiesResolveSupport `json:"resolveSupport,omitempty"`
}

// WorkspaceSymbolClientCapabilitiesResolveSupport - Workspace symbol resolve support capabilities.
type WorkspaceSymbolClientCapabilitiesResolveSupport struct {
	// Properties are workspace symbol properties that can be resolved lazily.
	Properties []string `json:"properties,omitempty"`
}

// FileOperationClientCapabilities - File operation capabilities.
//
// @since 3.16.0
type FileOperationClientCapabilities struct {
	// DidCreate indicates support for `workspace/didCreateFiles` notifications.
	DidCreate bool `json:"didCreate,omitempty"`
	// WillCreate indicates support for `workspace/willCreateFiles` requests.
	WillCreate bool `json:"willCreate,omitempty"`
	// DidRename indicates support for `workspace/didRenameFiles` notifications.
	DidRename bool `json:"didRename,omitempty"`
	// WillRename indicates support for `workspace/willRenameFiles` requests.
	WillRename bool `json:"willRename,omitempty"`
	// DidDelete indicates support for `workspace/didDeleteFiles` notifications.
	DidDelete bool `json:"didDelete,omitempty"`
	// WillDelete indicates support for `workspace/willDeleteFiles` requests.
	WillDelete bool `json:"willDelete,omitempty"`
}

// SemanticTokensWorkspaceClientCapabilities - Workspace semantic tokens capabilities.
type SemanticTokensWorkspaceClientCapabilities struct {
	// RefreshSupport indicates support for `workspace/semanticTokens/refresh`.
	RefreshSupport bool `json:"refreshSupport,omitempty"`
}

// CodeLensWorkspaceClientCapabilities - Workspace code lens capabilities.
type CodeLensWorkspaceClientCapabilities struct {
	// RefreshSupport indicates support for `workspace/codeLens/refresh`.
	RefreshSupport bool `json:"refreshSupport,omitempty"`
}

// InlineValueWorkspaceClientCapabilities - Workspace inline value capabilities.
type InlineValueWorkspaceClientCapabilities struct {
	// RefreshSupport indicates support for `workspace/inlineValue/refresh`.
	RefreshSupport bool `json:"refreshSupport,omitempty"`
}

// InlayHintWorkspaceClientCapabilities - Workspace inlay hint capabilities.
type InlayHintWorkspaceClientCapabilities struct {
	// RefreshSupport indicates support for `workspace/inlayHint/refresh`.
	RefreshSupport bool `json:"refreshSupport,omitempty"`
}

// DiagnosticWorkspaceClientCapabilities - Workspace diagnostics capabilities.
type DiagnosticWorkspaceClientCapabilities struct {
	// RefreshSupport indicates support for `workspace/diagnostic/refresh`.
	RefreshSupport bool `json:"refreshSupport,omitempty"`
}

// WorkspaceClientCapabilities - Workspace specific client capabilities.
//
// See https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification/#workspaceClientCapabilities
type WorkspaceClientCapabilities struct {
	// ApplyEdit indicates whether the client accepts workspace/applyEdit requests.
	ApplyEdit bool `json:"applyEdit,omitempty"`
	// WorkspaceEdit describes supported `WorkspaceEdit` features.
	WorkspaceEdit *WorkspaceEditClientCapabilities `json:"workspaceEdit,omitempty"`
	// DidChangeConfiguration describes dynamic registration support for
	// `workspace/didChangeConfiguration`.
	DidChangeConfiguration *DidChangeConfigurationClientCapabilities `json:"didChangeConfiguration,omitempty"`
	// DidChangeWatchedFiles describes dynamic registration support for
	// `workspace/didChangeWatchedFiles`.
	DidChangeWatchedFiles *DidChangeWatchedFilesClientCapabilities `json:"didChangeWatchedFiles,omitempty"`
	// Symbol describes support for `workspace/symbol`.
	Symbol *WorkspaceSymbolClientCapabilities `json:"symbol,omitempty"`
	// ExecuteCommand describes dynamic registration support for
	// `workspace/executeCommand`.
	ExecuteCommand *ExecuteCommandClientCapabilities `json:"executeCommand,omitempty"`
	// WorkspaceFolders indicates whether workspace folder support is available.
	WorkspaceFolders bool `json:"workspaceFolders,omitempty"`
	// Configuration indicates support for `workspace/configuration` requests.
	Configuration bool `json:"configuration,omitempty"`
	// SemanticTokens describes workspace semantic token refresh support.
	SemanticTokens *SemanticTokensWorkspaceClientCapabilities `json:"semanticTokens,omitempty"`
	// CodeLens describes workspace code lens refresh support.
	CodeLens *CodeLensWorkspaceClientCapabilities `json:"codeLens,omitempty"`
	// FileOperations describes file operation notification/request support.
	FileOperations *FileOperationClientCapabilities `json:"fileOperations,omitempty"`
	// InlineValue describes workspace inline value refresh support.
	InlineValue *InlineValueWorkspaceClientCapabilities `json:"inlineValue,omitempty"`
	// InlayHint describes workspace inlay hint refresh support.
	InlayHint *InlayHintWorkspaceClientCapabilities `json:"inlayHint,omitempty"`
	// Diagnostics describes workspace diagnostic refresh support.
	Diagnostics *DiagnosticWorkspaceClientCapabilities `json:"diagnostics,omitempty"`
}

// TextDocumentSyncClientCapabilities - Synchronization capabilities.
type TextDocumentSyncClientCapabilities struct {
	// DynamicRegistration indicates support for dynamic registration.
	DynamicRegistration bool `json:"dynamicRegistration,omitempty"`
	// WillSave indicates support for before-save notifications.
	WillSave bool `json:"willSave,omitempty"`
	// WillSaveWaitUntil indicates support for before-save requests that can
	// return text edits.
	WillSaveWaitUntil bool `json:"willSaveWaitUntil,omitempty"`
	// DidSave indicates support for save notifications.
	DidSave bool `json:"didSave,omitempty"`
}

// CompletionItemTagSupportClientCapabilities - Completion item tag support.
type CompletionItemTagSupportClientCapabilities struct {
	// ValueSet is the `CompletionItemTag` values supported by the client.
	ValueSet []CompletionItemTag `json:"valueSet,omitempty"`
}

// CompletionItemResolveSupportClientCapabilities - Completion item resolve support.
type CompletionItemResolveSupportClientCapabilities struct {
	// Properties are completion item properties that can be lazily resolved.
	Properties []string `json:"properties,omitempty"`
}

// InsertTextModeSupportClientCapabilities - Insert text mode support.
type InsertTextModeSupportClientCapabilities struct {
	// ValueSet is the supported insert text modes.
	ValueSet []InsertTextMode `json:"valueSet,omitempty"`
}

// CompletionItemClientCapabilities - Capabilities for completion items.
type CompletionItemClientCapabilities struct {
	// SnippetSupport indicates whether snippets are supported.
	SnippetSupport bool `json:"snippetSupport,omitempty"`
	// CommitCharactersSupport indicates support for commit characters.
	CommitCharactersSupport bool `json:"commitCharactersSupport,omitempty"`
	// DocumentationFormat lists preferred documentation formats.
	DocumentationFormat []MarkupKind `json:"documentationFormat,omitempty"`
	// DeprecatedSupport indicates support for deprecated completion items.
	DeprecatedSupport bool `json:"deprecatedSupport,omitempty"`
	// PreselectSupport indicates support for preselected completion items.
	PreselectSupport bool `json:"preselectSupport,omitempty"`
	// TagSupport describes supported completion item tags.
	TagSupport *CompletionItemTagSupportClientCapabilities `json:"tagSupport,omitempty"`
	// InsertReplaceSupport indicates support for insert/replace edits.
	InsertReplaceSupport bool `json:"insertReplaceSupport,omitempty"`
	// ResolveSupport describes lazily resolvable completion item properties.
	ResolveSupport *CompletionItemResolveSupportClientCapabilities `json:"resolveSupport,omitempty"`
	// InsertTextModeSupport describes supported insert text modes.
	InsertTextModeSupport *InsertTextModeSupportClientCapabilities `json:"insertTextModeSupport,omitempty"`
	// LabelDetailsSupport indicates support for completion item label details.
	LabelDetailsSupport bool `json:"labelDetailsSupport,omitempty"`
}

// CompletionItemKindClientCapabilities - Supported completion item kinds.
type CompletionItemKindClientCapabilities struct {
	// ValueSet is the completion item kinds the client supports.
	ValueSet []CompletionItemKind `json:"valueSet,omitempty"`
}

// CompletionListClientCapabilities - Completion list capabilities.
//
// @since 3.17.0
type CompletionListClientCapabilities struct {
	// ItemDefaults lists supported defaults in completion list responses.
	ItemDefaults []string `json:"itemDefaults,omitempty"`
}

// CompletionClientCapabilities - Completion request capabilities.
type CompletionClientCapabilities struct {
	// DynamicRegistration indicates support for dynamic registration.
	DynamicRegistration bool `json:"dynamicRegistration,omitempty"`
	// CompletionItem describes completion item-specific capabilities.
	CompletionItem *CompletionItemClientCapabilities `json:"completionItem,omitempty"`
	// CompletionItemKind describes supported completion item kinds.
	CompletionItemKind *CompletionItemKindClientCapabilities `json:"completionItemKind,omitempty"`
	// ContextSupport indicates support for additional completion context.
	ContextSupport bool `json:"contextSupport,omitempty"`
	// InsertTextMode is the default insertion mode used by the client.
	InsertTextMode InsertTextMode `json:"insertTextMode,omitempty"`
	// CompletionList describes completion list-level capabilities.
	CompletionList *CompletionListClientCapabilities `json:"completionList,omitempty"`
}

// HoverClientCapabilities - Hover request capabilities.
type HoverClientCapabilities struct {
	// DynamicRegistration indicates support for dynamic registration.
	DynamicRegistration bool `json:"dynamicRegistration,omitempty"`
	// ContentFormat lists preferred markup formats for hover content.
	ContentFormat []MarkupKind `json:"contentFormat,omitempty"`
}

// ParameterInformationClientCapabilities - Parameter information capabilities.
type ParameterInformationClientCapabilities struct {
	// LabelOffsetSupport indicates support for tuple label offsets.
	LabelOffsetSupport bool `json:"labelOffsetSupport,omitempty"`
}

// SignatureInformationClientCapabilities - Signature information capabilities.
type SignatureInformationClientCapabilities struct {
	// DocumentationFormat lists preferred formats for signature docs.
	DocumentationFormat []MarkupKind `json:"documentationFormat,omitempty"`
	// ParameterInformation describes parameter info capabilities.
	ParameterInformation *ParameterInformationClientCapabilities `json:"parameterInformation,omitempty"`
	// ActiveParameterSupport indicates support for active parameter info.
	ActiveParameterSupport bool `json:"activeParameterSupport,omitempty"`
}

// SignatureHelpClientCapabilities - Signature help capabilities.
type SignatureHelpClientCapabilities struct {
	// DynamicRegistration indicates support for dynamic registration.
	DynamicRegistration bool `json:"dynamicRegistration,omitempty"`
	// SignatureInformation describes signature information support.
	SignatureInformation *SignatureInformationClientCapabilities `json:"signatureInformation,omitempty"`
	// ContextSupport indicates support for signature help context.
	ContextSupport bool `json:"contextSupport,omitempty"`
}

// DefinitionClientCapabilities - Definition request capabilities.
type DefinitionClientCapabilities struct {
	// DynamicRegistration indicates support for dynamic registration.
	DynamicRegistration bool `json:"dynamicRegistration,omitempty"`
	// LinkSupport indicates support for link-based definitions.
	LinkSupport bool `json:"linkSupport,omitempty"`
}

type (
	DeclarationClientCapabilities    = DefinitionClientCapabilities
	TypeDefinitionClientCapabilities = DefinitionClientCapabilities
	ImplementationClientCapabilities = DefinitionClientCapabilities
)

// DocumentSymbolClientCapabilities - Document symbol capabilities.
type DocumentSymbolClientCapabilities struct {
	// DynamicRegistration indicates support for dynamic registration.
	DynamicRegistration bool `json:"dynamicRegistration,omitempty"`
	// SymbolKind lists symbol kinds supported by the client.
	SymbolKind *DocumentSymbolClientCapabilitiesSymbolKind `json:"symbolKind,omitempty"`
	// HierarchicalDocumentSymbolSupport indicates support for hierarchical
	// symbols in the document symbol response.
	HierarchicalDocumentSymbolSupport bool `json:"hierarchicalDocumentSymbolSupport,omitempty"`
	// TagSupport lists supported `SymbolTag` values.
	TagSupport *DocumentSymbolClientCapabilitiesTagSupport `json:"tagSupport,omitempty"`
	// LabelSupport indicates whether labels for symbols are supported.
	LabelSupport bool `json:"labelSupport,omitempty"`
}

type (
	DocumentSymbolClientCapabilitiesSymbolKind = WorkspaceSymbolClientCapabilitiesSymbolKind
	DocumentSymbolClientCapabilitiesTagSupport = WorkspaceSymbolClientCapabilitiesTagSupport
)

// CodeActionKindClientCapabilities - Supported code action kinds.
type CodeActionKindClientCapabilities struct {
	// ValueSet is the `CodeActionKind` values supported by the client.
	ValueSet []CodeActionKind `json:"valueSet,omitempty"`
}

// CodeActionLiteralSupportClientCapabilities - Literal support for code actions.
type CodeActionLiteralSupportClientCapabilities struct {
	// CodeActionKind describes supported code action kinds.
	CodeActionKind CodeActionKindClientCapabilities `json:"codeActionKind"`
}

// CodeActionResolveSupportClientCapabilities - Code action resolve support.
type CodeActionResolveSupportClientCapabilities struct {
	// Properties are code action properties that can be lazily resolved.
	Properties []string `json:"properties,omitempty"`
}

// CodeActionClientCapabilities - Code action request capabilities.
type CodeActionClientCapabilities struct {
	// DynamicRegistration indicates support for dynamic registration.
	DynamicRegistration bool `json:"dynamicRegistration,omitempty"`
	// CodeActionLiteralSupport describes literal code action support.
	CodeActionLiteralSupport *CodeActionLiteralSupportClientCapabilities `json:"codeActionLiteralSupport,omitempty"`
	// IsPreferredSupport indicates support for the `isPreferred` flag.
	IsPreferredSupport bool `json:"isPreferredSupport,omitempty"`
	// DisabledSupport indicates support for disabled code actions.
	DisabledSupport bool `json:"disabledSupport,omitempty"`
	// DataSupport indicates support for preserving code action data between
	// request and resolve.
	DataSupport bool `json:"dataSupport,omitempty"`
	// ResolveSupport describes lazily resolvable code action properties.
	ResolveSupport *CodeActionResolveSupportClientCapabilities `json:"resolveSupport,omitempty"`
	// HonorsChangeAnnotations indicates whether UI respects change annotations
	// in incoming workspace edits.
	HonorsChangeAnnotations bool `json:"honorsChangeAnnotations,omitempty"`
}

// RenameClientCapabilities - Rename request capabilities.
type RenameClientCapabilities struct {
	// DynamicRegistration indicates support for dynamic registration.
	DynamicRegistration bool `json:"dynamicRegistration,omitempty"`
	// PrepareSupport indicates support for prepare rename requests.
	PrepareSupport bool `json:"prepareSupport,omitempty"`
	// PrepareSupportDefaultBehavior indicates the default behavior when
	// prepare support is enabled.
	PrepareSupportDefaultBehavior uint32 `json:"prepareSupportDefaultBehavior,omitempty"`
	// HonorsChangeAnnotations indicates whether rename applies change
	// annotation semantics in UI.
	HonorsChangeAnnotations bool `json:"honorsChangeAnnotations,omitempty"`
}

// PublishDiagnosticsTagSupportClientCapabilities - Diagnostic tag support.
type PublishDiagnosticsTagSupportClientCapabilities struct {
	// ValueSet is the `DiagnosticTag` values supported by the client.
	ValueSet []DiagnosticTag `json:"valueSet,omitempty"`
}

// PublishDiagnosticsClientCapabilities - Publish diagnostics capabilities.
type PublishDiagnosticsClientCapabilities struct {
	// RelatedInformation indicates support for related diagnostic locations.
	RelatedInformation bool `json:"relatedInformation,omitempty"`
	// TagSupport lists supported diagnostic tag values.
	TagSupport *PublishDiagnosticsTagSupportClientCapabilities `json:"tagSupport,omitempty"`
	// VersionSupport indicates whether diagnostics can be versioned.
	VersionSupport bool `json:"versionSupport,omitempty"`
	// CodeDescriptionSupport indicates support for code descriptions.
	CodeDescriptionSupport bool `json:"codeDescriptionSupport,omitempty"`
	// DataSupport indicates support for preserving diagnostic data between
	// publish and code action requests.
	DataSupport bool `json:"dataSupport,omitempty"`
}

// FoldingRangeKindClientCapabilities - Folding range kind support.
type FoldingRangeKindClientCapabilities struct {
	// ValueSet is the folding range kinds supported by the client.
	ValueSet []string `json:"valueSet,omitempty"`
}

// FoldingRangeClientCapabilities - Folding range capabilities.
type FoldingRangeClientCapabilities struct {
	// DynamicRegistration indicates support for dynamic registration.
	DynamicRegistration bool `json:"dynamicRegistration,omitempty"`
	// RangeLimit is the maximum number of folding ranges preferred.
	RangeLimit uint32 `json:"rangeLimit,omitempty"`
	// LineFoldingOnly indicates whether only whole-line folds are supported.
	LineFoldingOnly bool `json:"lineFoldingOnly,omitempty"`
	// FoldingRangeKind describes supported folding range kinds.
	FoldingRangeKind *FoldingRangeClientCapabilitiesFoldingRangeKind `json:"foldingRangeKind,omitempty"`
	// FoldingRange contains additional folding range specific support.
	FoldingRange *FoldingRangeClientCapabilitiesFoldingRange `json:"foldingRange,omitempty"`
}

type FoldingRangeClientCapabilitiesFoldingRangeKind = FoldingRangeKindClientCapabilities

type FoldingRangeClientCapabilitiesFoldingRange struct {
	// CollapsedText indicates support for custom collapsed text in ranges.
	CollapsedText bool `json:"collapsedText,omitempty"`
}

// TokenFormat - Semantic tokens format kinds.
type TokenFormat string

const (
	// TokenFormatRelative denotes relative token encoding.
	TokenFormatRelative TokenFormat = "relative"
)

// SemanticTokensRequestsClientCapabilities - Semantic token request modes.
type SemanticTokensRequestsClientCapabilities struct {
	// Range indicates support for range semantic token requests.
	Range LSPAny `json:"range,omitempty"`
	// Full indicates support for full document semantic token requests.
	Full LSPAny `json:"full,omitempty"`
}

type SemanticTokensClientCapabilitiesRequests = SemanticTokensRequestsClientCapabilities

// SemanticTokensClientCapabilities - Semantic token capabilities.
type SemanticTokensClientCapabilities struct {
	// DynamicRegistration indicates support for dynamic registration.
	DynamicRegistration bool `json:"dynamicRegistration,omitempty"`
	// Requests describes range/full request support.
	Requests SemanticTokensClientCapabilitiesRequests `json:"requests"`
	// TokenTypes is the complete set of supported token types.
	TokenTypes []string `json:"tokenTypes"`
	// TokenModifiers is the complete set of supported token modifiers.
	TokenModifiers []string `json:"tokenModifiers"`
	// Formats lists supported semantic token encodings.
	Formats []TokenFormat `json:"formats"`
	// OverlappingTokenSupport indicates support for overlapping tokens.
	OverlappingTokenSupport bool `json:"overlappingTokenSupport,omitempty"`
	// MultilineTokenSupport indicates support for multiline tokens.
	MultilineTokenSupport bool `json:"multilineTokenSupport,omitempty"`
	// ServerCancelSupport indicates support for server-side cancellation.
	ServerCancelSupport bool `json:"serverCancelSupport,omitempty"`
	// AugmentsSyntaxTokens indicates whether semantic tokens augment syntax
	// tokens instead of replacing them.
	AugmentsSyntaxTokens bool `json:"augmentsSyntaxTokens,omitempty"`
}

// InlayHintClientCapabilities - Inlay hint capabilities.
//
// @since 3.17.0
type InlayHintClientCapabilities struct {
	// DynamicRegistration indicates support for dynamic registration.
	DynamicRegistration bool `json:"dynamicRegistration,omitempty"`
	// ResolveSupport lists inlay hint properties that can be resolved lazily.
	ResolveSupport *InlayHintClientCapabilitiesResolveSupport `json:"resolveSupport,omitempty"`
}

type InlayHintClientCapabilitiesResolveSupport = WorkspaceSymbolClientCapabilitiesResolveSupport

// DiagnosticClientCapabilities - Pull diagnostic capabilities.
//
// @since 3.17.0
type DiagnosticClientCapabilities struct {
	// DynamicRegistration indicates support for dynamic registration.
	DynamicRegistration bool `json:"dynamicRegistration,omitempty"`
	// RelatedDocumentSupport indicates whether diagnostics can be reported
	// with related document data.
	RelatedDocumentSupport bool `json:"relatedDocumentSupport,omitempty"`
}

// TextDocumentClientCapabilities - Text document specific client capabilities.
//
// See https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification/#textDocumentClientCapabilities
type TextDocumentClientCapabilities struct {
	Synchronization *TextDocumentSyncClientCapabilities `json:"synchronization,omitempty"`

	// Capabilities specific to the `textDocument/completion` request.
	Completion *CompletionClientCapabilities `json:"completion,omitempty"`

	// Capabilities specific to the `textDocument/hover` request.
	Hover *HoverClientCapabilities `json:"hover,omitempty"`

	// Capabilities specific to the `textDocument/signatureHelp` request.
	SignatureHelp *SignatureHelpClientCapabilities `json:"signatureHelp,omitempty"`

	// Capabilities specific to the `textDocument/declaration` request.
	//
	// @since 3.14.0
	Declaration *DeclarationClientCapabilities `json:"declaration,omitempty"`

	// Capabilities specific to the `textDocument/definition` request.
	Definition *DefinitionClientCapabilities `json:"definition,omitempty"`

	// Capabilities specific to the `textDocument/typeDefinition` request.
	//
	// @since 3.6.0
	TypeDefinition *TypeDefinitionClientCapabilities `json:"typeDefinition,omitempty"`

	// Capabilities specific to the `textDocument/implementation` request.
	//
	// @since 3.6.0
	Implementation *ImplementationClientCapabilities `json:"implementation,omitempty"`

	// Capabilities specific to the `textDocument/references` request.
	References *ReferenceClientCapabilities `json:"references,omitempty"`

	// Capabilities specific to the `textDocument/documentHighlight` request.
	DocumentHighlight *DocumentHighlightClientCapabilities `json:"documentHighlight,omitempty"`

	// Capabilities specific to the `textDocument/documentSymbol` request.
	DocumentSymbol *DocumentSymbolClientCapabilities `json:"documentSymbol,omitempty"`

	// Capabilities specific to the `textDocument/codeAction` request.
	CodeAction *CodeActionClientCapabilities `json:"codeAction,omitempty"`

	// Capabilities specific to the `textDocument/codeLens` request.
	CodeLens *CodeLensClientCapabilities `json:"codeLens,omitempty"`

	// Capabilities specific to the `textDocument/documentLink` request.
	DocumentLink *DocumentLinkClientCapabilities `json:"documentLink,omitempty"`

	// Capabilities specific to the `textDocument/documentColor` and the
	// `textDocument/colorPresentation` request.
	//
	// @since 3.6.0
	ColorProvider *DocumentColorClientCapabilities `json:"colorProvider,omitempty"`

	// Capabilities specific to the `textDocument/formatting` request.
	Formatting *DocumentFormattingClientCapabilities `json:"formatting,omitempty"`

	// Capabilities specific to the `textDocument/rangeFormatting` request.
	RangeFormatting *DocumentRangeFormattingClientCapabilities `json:"rangeFormatting,omitempty"`

	// Capabilities specific to the `textDocument/onTypeFormatting` request.
	OnTypeFormatting *DocumentOnTypeFormattingClientCapabilities `json:"onTypeFormatting,omitempty"`

	// Capabilities specific to the `textDocument/rename` request.
	Rename *RenameClientCapabilities `json:"rename,omitempty"`

	// Capabilities specific to the `textDocument/publishDiagnostics`
	// notification.
	PublishDiagnostics *PublishDiagnosticsClientCapabilities `json:"publishDiagnostics,omitempty"`

	// Capabilities specific to the `textDocument/foldingRange` request.
	//
	// @since 3.10.0
	FoldingRange *FoldingRangeClientCapabilities `json:"foldingRange,omitempty"`

	// Capabilities specific to the `textDocument/selectionRange` request.
	//
	// @since 3.15.0
	SelectionRange *SelectionRangeClientCapabilities `json:"selectionRange,omitempty"`

	// Capabilities specific to the `textDocument/linkedEditingRange` request.
	//
	// @since 3.16.0
	LinkedEditingRange *LinkedEditingRangeClientCapabilities `json:"linkedEditingRange,omitempty"`

	// Capabilities specific to the various call hierarchy requests.
	//
	// @since 3.16.0
	CallHierarchy *CallHierarchyClientCapabilities `json:"callHierarchy,omitempty"`

	// Capabilities specific to the various semantic token requests.
	//
	// @since 3.16.0
	SemanticTokens *SemanticTokensClientCapabilities `json:"semanticTokens,omitempty"`

	// Capabilities specific to the `textDocument/moniker` request.
	//
	// @since 3.16.0
	Moniker *MonikerClientCapabilities `json:"moniker,omitempty"`

	// Capabilities specific to the various type hierarchy requests.
	//
	// @since 3.17.0
	TypeHierarchy *TypeHierarchyClientCapabilities `json:"typeHierarchy,omitempty"`

	// Capabilities specific to the `textDocument/inlineValue` request.
	//
	// @since 3.17.0
	InlineValue *InlineValueClientCapabilities `json:"inlineValue,omitempty"`

	// Capabilities specific to the `textDocument/inlayHint` request.
	//
	// @since 3.17.0
	InlayHint *InlayHintClientCapabilities `json:"inlayHint,omitempty"`

	// Capabilities specific to the diagnostic pull model.
	//
	// @since 3.17.0
	Diagnostic *DiagnosticClientCapabilities `json:"diagnostic,omitempty"`
}

type DocumentLinkClientCapabilities struct {
	// DynamicRegistration indicates support for dynamic registration.
	DynamicRegistration bool `json:"dynamicRegistration,omitempty"`
	// TooltipSupport indicates support for the tooltip field on document links.
	TooltipSupport bool `json:"tooltipSupport,omitempty"`
}

// NotebookDocumentSyncClientCapabilities - Notebook synchronization capabilities.
//
// @since 3.17.0
type NotebookDocumentSyncClientCapabilities struct {
	// DynamicRegistration indicates support for dynamic registration.
	DynamicRegistration bool `json:"dynamicRegistration,omitempty"`
	// ExecutionSummarySupport indicates support for notebook cell execution
	// summary synchronization.
	ExecutionSummarySupport bool `json:"executionSummarySupport,omitempty"`
}

// NotebookDocumentClientCapabilities - Notebook document specific client capabilities.
//
// See https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification/#notebookDocumentClientCapabilities
//
// @since 3.17.0
type NotebookDocumentClientCapabilities struct {
	// Synchronization describes notebook synchronization support.
	Synchronization *NotebookDocumentSyncClientCapabilities `json:"synchronization,omitempty"`
}

// StaleRequestSupportClientCapabilities - Stale request support details.
//
// @since 3.17.0
type StaleRequestSupportClientCapabilities struct {
	// Cancel indicates whether stale requests should be canceled.
	Cancel bool `json:"cancel,omitempty"`
	// RetryOnContentModified lists request methods that should be retried when
	// content changed.
	RetryOnContentModified []string `json:"retryOnContentModified,omitempty"`
}

// RegularExpressionsClientCapabilities - Regular expression engine details.
//
// @since 3.16.0
type RegularExpressionsClientCapabilities struct {
	// Engine is the regular expression engine name (for example `ECMAScript`).
	Engine string `json:"engine"`
	// Version is the engine version when available.
	Version string `json:"version,omitempty"`
}

// MarkdownClientCapabilities - Markdown parser details.
//
// @since 3.16.0
type MarkdownClientCapabilities struct {
	// Parser is the markdown parser name.
	Parser string `json:"parser"`
	// Version is the parser version when available.
	Version string `json:"version,omitempty"`
	// AllowedTags are HTML tags allowed in markdown rendering.
	AllowedTags []string `json:"allowedTags,omitempty"`
}

// GeneralClientCapabilities - General client capabilities.
//
// See https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification/#generalClientCapabilities
type GeneralClientCapabilities struct {
	// StaleRequestSupport describes support for stale request handling.
	StaleRequestSupport *StaleRequestSupportClientCapabilities `json:"staleRequestSupport,omitempty"`
	// RegularExpressions describes regex engine capabilities.
	RegularExpressions *RegularExpressionsClientCapabilities `json:"regularExpressions,omitempty"`
	// Markdown describes markdown parser capabilities.
	Markdown *MarkdownClientCapabilities `json:"markdown,omitempty"`
	// PositionEncodings lists position encodings supported by the client.
	PositionEncodings []PositionEncodingKind `json:"positionEncodings,omitempty"`
}

// ClientCapabilities defines the capabilities of the client (e.g., editor or IDE).
// It tells the language server what features the client supports.
//
// See https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification/#clientCapabilities
type ClientCapabilities struct {
	// Workspace specific client capabilities.
	Workspace *WorkspaceClientCapabilities `json:"workspace,omitempty"`

	// Text document specific client capabilities.
	TextDocument *TextDocumentClientCapabilities `json:"textDocument,omitempty"`

	// Notebook document specific client capabilities.
	//
	// @since 3.17.0
	NotebookDocument *NotebookDocumentClientCapabilities `json:"notebookDocument,omitempty"`

	// Window specific client capabilities.
	Window *WindowClientCapabilities `json:"window,omitempty"`

	// General client capabilities.
	General *GeneralClientCapabilities `json:"general,omitempty"`

	// Experimental client capabilities. The value can be any JSON type.
	Experimental LSPAny `json:"experimental,omitempty"`
}

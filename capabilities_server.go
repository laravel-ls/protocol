package protocol

// ServerCapabilities defines the capabilities of the language server.
//
// See https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification/#serverCapabilities
type ServerCapabilities struct {
	// The position encoding the server picked from the encodings offered by the
	// client via the client capability `general.positionEncodings`.
	//
	// If the client didn't provide any position encodings the only valid value
	// that a server can return is `utf-16`.
	//
	// If omitted it defaults to `utf-16`.
	//
	// @since 3.17.0
	PositionEncoding *PositionEncodingKind `json:"positionEncoding,omitempty"`

	// Defines how text documents are synced.
	//
	// This can be either:
	// - `TextDocumentSyncOptions`, or
	// - `TextDocumentSyncKind` (for backwards compatibility).
	TextDocumentSync any `json:"textDocumentSync,omitempty"` // TextDocumentSyncOptions | TextDocumentSyncKind

	// Defines notebook document synchronization support.
	//
	// @since 3.17.0
	NotebookDocumentSync any `json:"notebookDocumentSync,omitempty"` // NotebookDocumentSyncOptions | NotebookDocumentSyncRegistrationOptions

	// The server provides completion support.
	CompletionProvider *CompletionOptions `json:"completionProvider,omitempty"`

	// The server provides hover support.
	HoverProvider any `json:"hoverProvider,omitempty"` // bool | HoverOptions

	// The server provides signature help support.
	SignatureHelpProvider *SignatureHelpOptions `json:"signatureHelpProvider,omitempty"`

	// The server provides goto declaration support.
	//
	// @since 3.14.0
	DeclarationProvider any `json:"declarationProvider,omitempty"` // bool | DeclarationOptions | DeclarationRegistrationOptions

	// The server provides goto definition support.
	DefinitionProvider any `json:"definitionProvider,omitempty"` // bool | DefinitionOptions

	// The server provides goto type definition support.
	//
	// @since 3.6.0
	TypeDefinitionProvider any `json:"typeDefinitionProvider,omitempty"` // bool | TypeDefinitionOptions | TypeDefinitionRegistrationOptions

	// The server provides goto implementation support.
	//
	// @since 3.6.0
	ImplementationProvider any `json:"implementationProvider,omitempty"` // bool | ImplementationOptions | ImplementationRegistrationOptions

	// The server provides find references support.
	ReferencesProvider any `json:"referencesProvider,omitempty"` // bool | ReferenceOptions

	// The server provides document highlight support.
	DocumentHighlightProvider any `json:"documentHighlightProvider,omitempty"` // bool | DocumentHighlightOptions

	// The server provides document symbol support.
	DocumentSymbolProvider any `json:"documentSymbolProvider,omitempty"` // bool | DocumentSymbolOptions

	// The server provides code action support.
	CodeActionProvider any `json:"codeActionProvider,omitempty"` // bool | CodeActionOptions

	// The server provides code lens support.
	CodeLensProvider *CodeLensOptions `json:"codeLensProvider,omitempty"`

	// The server provides document link support.
	DocumentLinkProvider *DocumentLinkOptions `json:"documentLinkProvider,omitempty"`

	// The server provides color provider support.
	//
	// @since 3.6.0
	ColorProvider any `json:"colorProvider,omitempty"` // bool | DocumentColorOptions | DocumentColorRegistrationOptions

	// The server provides document formatting support.
	DocumentFormattingProvider any `json:"documentFormattingProvider,omitempty"` // bool | DocumentFormattingOptions

	// The server provides document range formatting support.
	DocumentRangeFormattingProvider any `json:"documentRangeFormattingProvider,omitempty"` // bool | DocumentRangeFormattingOptions

	// The server provides document on type formatting support.
	DocumentOnTypeFormattingProvider *DocumentOnTypeFormattingOptions `json:"documentOnTypeFormattingProvider,omitempty"`

	// The server provides rename support.
	RenameProvider any `json:"renameProvider,omitempty"` // bool | RenameOptions

	// The server provides folding range support.
	//
	// @since 3.10.0
	FoldingRangeProvider any `json:"foldingRangeProvider,omitempty"` // bool | FoldingRangeOptions | FoldingRangeRegistrationOptions

	// The server provides execute command support.
	ExecuteCommandProvider *ExecuteCommandOptions `json:"executeCommandProvider,omitempty"`

	// The server provides selection range support.
	//
	// @since 3.15.0
	SelectionRangeProvider any `json:"selectionRangeProvider,omitempty"` // bool | SelectionRangeOptions | SelectionRangeRegistrationOptions

	// The server provides linked editing range support.
	//
	// @since 3.16.0
	LinkedEditingRangeProvider any `json:"linkedEditingRangeProvider,omitempty"` // bool | LinkedEditingRangeOptions | LinkedEditingRangeRegistrationOptions

	// The server provides call hierarchy support.
	//
	// @since 3.16.0
	CallHierarchyProvider any `json:"callHierarchyProvider,omitempty"` // bool | CallHierarchyOptions | CallHierarchyRegistrationOptions

	// The server provides semantic tokens support.
	//
	// @since 3.16.0
	SemanticTokensProvider any `json:"semanticTokensProvider,omitempty"` // SemanticTokensOptions | SemanticTokensRegistrationOptions

	// The server provides moniker support.
	//
	// @since 3.16.0
	MonikerProvider any `json:"monikerProvider,omitempty"` // bool | MonikerOptions | MonikerRegistrationOptions

	// The server provides type hierarchy support.
	//
	// @since 3.17.0
	TypeHierarchyProvider any `json:"typeHierarchyProvider,omitempty"` // bool | TypeHierarchyOptions | TypeHierarchyRegistrationOptions

	// The server provides inline value support.
	//
	// @since 3.17.0
	InlineValueProvider any `json:"inlineValueProvider,omitempty"` // bool | InlineValueOptions | InlineValueRegistrationOptions

	// The server provides inlay hint support.
	//
	// @since 3.17.0
	InlayHintProvider any `json:"inlayHintProvider,omitempty"` // bool | InlayHintOptions | InlayHintRegistrationOptions

	// The server has support for pull model diagnostics.
	//
	// @since 3.17.0
	DiagnosticProvider any `json:"diagnosticProvider,omitempty"` // DiagnosticOptions | DiagnosticRegistrationOptions

	// The server provides workspace symbol support.
	//
	// @since 3.17.0
	WorkspaceSymbolProvider any `json:"workspaceSymbolProvider,omitempty"` // bool | WorkspaceSymbolOptions

	// Workspace-specific server capabilities.
	Workspace *WorkspaceServerCapabilities `json:"workspace,omitempty"`

	// Experimental server capabilities. The value can be any JSON type.
	Experimental *LSPAny `json:"experimental,omitempty"`
}

// TextDocumentSyncKind defines how the host (editor) should sync document
// changes to the language server.
//
// See https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification/#textDocumentSyncKind
type TextDocumentSyncKind int

const (
	// Documents should not be synced at all.
	TextDocumentSyncKindNone TextDocumentSyncKind = 0

	// Documents are synced by always sending the full content of the document.
	TextDocumentSyncKindFull TextDocumentSyncKind = 1

	// Documents are synced by sending the full content on open and close. After
	// that only incremental updates to the document are sent.
	TextDocumentSyncKindIncremental TextDocumentSyncKind = 2
)

// TextDocumentSyncOptions defines detailed text document synchronization
// options.
//
// See https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification/#textDocumentSyncOptions
type TextDocumentSyncOptions struct {
	OpenClose         bool                 `json:"openClose,omitempty"`
	Change            TextDocumentSyncKind `json:"change,omitempty"`
	WillSave          bool                 `json:"willSave,omitempty"`
	WillSaveWaitUntil bool                 `json:"willSaveWaitUntil,omitempty"`
	Save              any                  `json:"save,omitempty"` // bool | SaveOptions
}

// SaveOptions options for save notifications.
type SaveOptions struct {
	IncludeText bool `json:"includeText,omitempty"`
}

// CompletionOptions used during `initialize` or server capabilities
// registration.
//
// See https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification/#completionOptions
type CompletionOptions struct {
	WorkDoneProgressOptions

	// The additional characters that should trigger completion.
	TriggerCharacters []string `json:"triggerCharacters,omitempty"`

	// Characters that commit a completion item.
	//
	// @since 3.2.0
	AllCommitCharacters []string `json:"allCommitCharacters,omitempty"`

	// The server supports completion item resolve.
	ResolveProvider *bool `json:"resolveProvider,omitempty"`

	// Completion item-specific server capabilities.
	//
	// @since 3.17.0
	CompletionItem *CompletionOptionsCompletionItem `json:"completionItem,omitempty"`
}

// CompletionOptionsCompletionItem capabilities specific to completion items.
//
// @since 3.17.0
type CompletionOptionsCompletionItem struct {
	// The server supports `CompletionItemLabelDetails` in resolve responses.
	LabelDetailsSupport *bool `json:"labelDetailsSupport,omitempty"`
}

// Backward compatible alias.
type CompletionItemCapability = CompletionOptionsCompletionItem

// SignatureHelpOptions server capabilities for signature help requests.
//
// See https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification/#signatureHelpOptions
type SignatureHelpOptions struct {
	TriggerCharacters   []string `json:"triggerCharacters,omitempty"`
	RetriggerCharacters []string `json:"retriggerCharacters,omitempty"`
}

// WorkDoneProgressOptions indicates if a request supports work-done progress
// reporting.
type WorkDoneProgressOptions struct {
	WorkDoneProgress bool `json:"workDoneProgress,omitempty"`
}

// HoverOptions describes hover support details.
type HoverOptions = WorkDoneProgressOptions

// DeclarationOptions describes declaration support details.
//
// @since 3.14.0
type DeclarationOptions = WorkDoneProgressOptions

// DefinitionOptions describes definition support details.
type DefinitionOptions = WorkDoneProgressOptions

// TypeDefinitionOptions describes type definition support details.
//
// @since 3.6.0
type TypeDefinitionOptions = WorkDoneProgressOptions

// ImplementationOptions describes implementation support details.
//
// @since 3.6.0
type ImplementationOptions = WorkDoneProgressOptions

// ReferenceOptions describes references support details.
type ReferenceOptions = WorkDoneProgressOptions

// DocumentHighlightOptions describes document highlight support details.
type DocumentHighlightOptions = WorkDoneProgressOptions

// DocumentSymbolOptions describes document symbol support details.
type DocumentSymbolOptions = WorkDoneProgressOptions

// CodeActionOptions server capability for code actions.
type CodeActionOptions struct {
	WorkDoneProgressOptions

	CodeActionKinds []CodeActionKind `json:"codeActionKinds,omitempty"`
	ResolveProvider bool             `json:"resolveProvider,omitempty"`
}

// CodeLensOptions server capability for code lenses.
type CodeLensOptions struct {
	ResolveProvider bool `json:"resolveProvider,omitempty"`
}

// DocumentLinkOptions server capability for document links.
type DocumentLinkOptions struct {
	ResolveProvider bool `json:"resolveProvider,omitempty"`
}

// DocumentColorOptions describes document color support details.
//
// @since 3.6.0
type DocumentColorOptions = WorkDoneProgressOptions

// DocumentFormattingOptions describes formatting support details.
type DocumentFormattingOptions = WorkDoneProgressOptions

// DocumentRangeFormattingOptions describes range formatting support details.
type DocumentRangeFormattingOptions = WorkDoneProgressOptions

// DocumentOnTypeFormattingOptions server capability for on-type formatting.
type DocumentOnTypeFormattingOptions struct {
	FirstTriggerCharacter string   `json:"firstTriggerCharacter"`
	MoreTriggerCharacter  []string `json:"moreTriggerCharacter,omitempty"`
}

// RenameOptions server capability for rename.
type RenameOptions struct {
	WorkDoneProgressOptions
	PrepareProvider bool `json:"prepareProvider,omitempty"`
}

// FoldingRangeOptions describes folding range support details.
//
// @since 3.10.0
type FoldingRangeOptions = WorkDoneProgressOptions

// ExecuteCommandOptions server capability for execute command.
type ExecuteCommandOptions struct {
	WorkDoneProgressOptions
	Commands []string `json:"commands"`
}

// SelectionRangeOptions describes selection range support details.
//
// @since 3.15.0
type SelectionRangeOptions = WorkDoneProgressOptions

// LinkedEditingRangeOptions describes linked editing range support details.
//
// @since 3.16.0
type LinkedEditingRangeOptions = WorkDoneProgressOptions

// CallHierarchyOptions describes call hierarchy support details.
//
// @since 3.16.0
type CallHierarchyOptions = WorkDoneProgressOptions

// SemanticTokensLegend describes token types and token modifiers a server uses.
//
// @since 3.16.0
type SemanticTokensLegend struct {
	TokenTypes     []string `json:"tokenTypes"`
	TokenModifiers []string `json:"tokenModifiers"`
}

// SemanticTokensOptionsFull indicates support for full document semantic tokens.
//
// @since 3.16.0
type SemanticTokensOptionsFull struct {
	Delta bool `json:"delta,omitempty"`
}

// SemanticTokensOptions server capability for semantic tokens.
//
// @since 3.16.0
type SemanticTokensOptions struct {
	WorkDoneProgressOptions
	Legend SemanticTokensLegend `json:"legend"`
	Range  any                  `json:"range,omitempty"` // bool | map[string]any
	Full   any                  `json:"full,omitempty"`  // bool | SemanticTokensOptionsFull
}

// MonikerOptions describes moniker support details.
//
// @since 3.16.0
type MonikerOptions = WorkDoneProgressOptions

// TypeHierarchyOptions describes type hierarchy support details.
//
// @since 3.17.0
type TypeHierarchyOptions = WorkDoneProgressOptions

// InlineValueOptions describes inline value support details.
//
// @since 3.17.0
type InlineValueOptions = WorkDoneProgressOptions

// InlayHintOptions server capability for inlay hints.
//
// @since 3.17.0
type InlayHintOptions struct {
	WorkDoneProgressOptions
	ResolveProvider bool `json:"resolveProvider,omitempty"`
}

// DiagnosticOptions registration options to configure pull diagnostics.
//
// See https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification/#diagnosticOptions
//
// @since 3.17.0
type DiagnosticOptions struct {
	WorkDoneProgressOptions

	Identifier            string `json:"identifier,omitempty"`
	InterFileDependencies bool   `json:"interFileDependencies"`
	WorkspaceDiagnostics  bool   `json:"workspaceDiagnostics"`
}

// WorkspaceSymbolOptions server capability for workspace symbols.
//
// @since 3.17.0
type WorkspaceSymbolOptions struct {
	WorkDoneProgressOptions
	ResolveProvider bool `json:"resolveProvider,omitempty"`
}

// WorkspaceServerCapabilities workspace-level server capabilities.
type WorkspaceServerCapabilities struct {
	WorkspaceFolders *WorkspaceFoldersServerCapabilities `json:"workspaceFolders,omitempty"`
	FileOperations   *FileOperationOptions               `json:"fileOperations,omitempty"`
}

// WorkspaceFoldersServerCapabilities server support for workspace folders.
type WorkspaceFoldersServerCapabilities struct {
	Supported           bool `json:"supported,omitempty"`
	ChangeNotifications any  `json:"changeNotifications,omitempty"` // string | bool
}

// FileOperationOptions server support for file operation requests/notifications.
//
// @since 3.16.0
type FileOperationOptions struct {
	DidCreate  *FileOperationRegistrationOptions `json:"didCreate,omitempty"`
	WillCreate *FileOperationRegistrationOptions `json:"willCreate,omitempty"`
	DidRename  *FileOperationRegistrationOptions `json:"didRename,omitempty"`
	WillRename *FileOperationRegistrationOptions `json:"willRename,omitempty"`
	DidDelete  *FileOperationRegistrationOptions `json:"didDelete,omitempty"`
	WillDelete *FileOperationRegistrationOptions `json:"willDelete,omitempty"`
}

// FileOperationRegistrationOptions registration options for file operations.
//
// @since 3.16.0
type FileOperationRegistrationOptions struct {
	Filters []FileOperationFilter `json:"filters"`
}

// FileOperationFilter describes a filter for file operation events.
//
// @since 3.16.0
type FileOperationFilter struct {
	Scheme  string               `json:"scheme,omitempty"`
	Pattern FileOperationPattern `json:"pattern"`
}

// FileOperationPatternKind describes whether a file operation pattern applies
// to files or folders.
//
// @since 3.16.0
type FileOperationPatternKind string

const (
	FileOperationPatternKindFile   FileOperationPatternKind = "file"
	FileOperationPatternKindFolder FileOperationPatternKind = "folder"
)

// FileOperationPatternOptions options for file operation patterns.
//
// @since 3.16.0
type FileOperationPatternOptions struct {
	IgnoreCase bool `json:"ignoreCase,omitempty"`
}

// FileOperationPattern glob pattern for file operations.
//
// @since 3.16.0
type FileOperationPattern struct {
	Glob    string                       `json:"glob"`
	Matches FileOperationPatternKind     `json:"matches,omitempty"`
	Options *FileOperationPatternOptions `json:"options,omitempty"`
}

// Registration option aliases with spec names.
type (
	DeclarationRegistrationOptions          = TextDocumentRegistrationOptions
	TypeDefinitionRegistrationOptions       = TextDocumentRegistrationOptions
	ImplementationRegistrationOptions       = TextDocumentRegistrationOptions
	ColorProviderRegistrationOptions        = TextDocumentRegistrationOptions
	FoldingRangeRegistrationOptions         = TextDocumentRegistrationOptions
	SelectionRangeRegistrationOptions       = TextDocumentRegistrationOptions
	LinkedEditingRangeRegistrationOptions   = TextDocumentRegistrationOptions
	CallHierarchyRegistrationOptions        = TextDocumentRegistrationOptions
	MonikerRegistrationOptions              = TextDocumentRegistrationOptions
	TypeHierarchyRegistrationOptions        = TextDocumentRegistrationOptions
	InlineValueRegistrationOptions          = TextDocumentRegistrationOptions
	InlayHintRegistrationOptions            = TextDocumentRegistrationOptions
	SemanticTokensRegistrationOptions       = TextDocumentRegistrationOptions
	DiagnosticRegistrationOptions           = TextDocumentRegistrationOptions
	NotebookDocumentSyncRegistrationOptions = TextDocumentRegistrationOptions
)

// TextDocumentRegistrationOptions registration options for text document scoped
// capabilities.
type TextDocumentRegistrationOptions struct {
	DocumentSelector LSPAny `json:"documentSelector,omitempty"`
}

// NotebookDocumentSyncOptions server options for notebook synchronization.
//
// @since 3.17.0
type NotebookDocumentSyncOptions struct {
	NotebookSelector []NotebookDocumentSyncOptionsNotebookSelector `json:"notebookSelector"`
	Save             bool                                          `json:"save,omitempty"`
}

// NotebookDocumentSyncOptionsNotebookSelector selects matching notebooks and
// cell languages.
//
// @since 3.17.0
type NotebookDocumentSyncOptionsNotebookSelector struct {
	Notebook any                                                `json:"notebook,omitempty"` // string | NotebookDocumentFilter
	Cells    []NotebookDocumentSyncOptionsNotebookSelectorCells `json:"cells,omitempty"`
}

// NotebookDocumentSyncOptionsNotebookSelectorCells cell language selector.
//
// @since 3.17.0
type NotebookDocumentSyncOptionsNotebookSelectorCells struct {
	Language string `json:"language"`
}

// NotebookDocumentFilter selects notebooks by type, scheme and pattern.
//
// @since 3.17.0
type NotebookDocumentFilter struct {
	NotebookType string `json:"notebookType,omitempty"`
	Scheme       string `json:"scheme,omitempty"`
	Pattern      string `json:"pattern,omitempty"`
}

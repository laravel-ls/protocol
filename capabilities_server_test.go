package protocol_test

import (
	"encoding/json"
	"testing"

	"github.com/laravel-ls/protocol"
)

func Test_CapabilitiesServer_UnmarshalValidJSON(t *testing.T) {
	data := []byte(`{
		"positionEncoding": "utf-16",
		"textDocumentSync": 2,
		"completionProvider": {
			"triggerCharacters": [".", ":"],
			"resolveProvider": true,
			"completionItem": {
				"labelDetailsSupport": true
			}
		},
		"hoverProvider": true,
		"definitionProvider": true,
		"codeActionProvider": true,
		"signatureHelpProvider": {
			"triggerCharacters": ["(", ","],
			"retriggerCharacters": [")"]
		},
		"diagnosticProvider": {
			"interFileDependencies": true,
			"workspaceDiagnostics": true
		}
	}`)

	var decoded protocol.ServerCapabilities
	if err := json.Unmarshal(data, &decoded); err != nil {
		t.Fatalf("unmarshal failed: %v", err)
	}

	positionEncoding := protocol.PositionEncodingKindUTF16
	if decoded.PositionEncoding == nil || *decoded.PositionEncoding != positionEncoding {
		t.Fatalf("expected positionEncoding=%q, got %+v", positionEncoding, decoded.PositionEncoding)
	}

	if decoded.CompletionProvider == nil {
		t.Fatalf("expected completionProvider to be present")
	}

	if decoded.CompletionProvider.ResolveProvider == nil || !*decoded.CompletionProvider.ResolveProvider {
		t.Fatalf("expected completionProvider.resolveProvider=true")
	}

	if decoded.CompletionProvider.CompletionItem == nil {
		t.Fatalf("expected completionProvider.completionItem to be present")
	}

	if decoded.CompletionProvider.CompletionItem.LabelDetailsSupport == nil || !*decoded.CompletionProvider.CompletionItem.LabelDetailsSupport {
		t.Fatalf("expected completionProvider.completionItem.labelDetailsSupport=true")
	}

	textDocumentSync, ok := decoded.TextDocumentSync.(float64)
	if !ok || int(textDocumentSync) != int(protocol.TextDocumentSyncKindIncremental) {
		t.Fatalf("expected textDocumentSync=%d, got %#v", protocol.TextDocumentSyncKindIncremental, decoded.TextDocumentSync)
	}

	diagnosticProvider, ok := decoded.DiagnosticProvider.(map[string]any)
	if !ok {
		t.Fatalf("expected diagnosticProvider object, got %#v", decoded.DiagnosticProvider)
	}

	if inter, ok := diagnosticProvider["interFileDependencies"].(bool); !ok || !inter {
		t.Fatalf("expected diagnosticProvider.interFileDependencies=true, got %#v", diagnosticProvider["interFileDependencies"])
	}

	if ws, ok := diagnosticProvider["workspaceDiagnostics"].(bool); !ok || !ws {
		t.Fatalf("expected diagnosticProvider.workspaceDiagnostics=true, got %#v", diagnosticProvider["workspaceDiagnostics"])
	}
}

func Test_CapabilitiesServer_MarshalOmitEmptyFields(t *testing.T) {
	original := protocol.ServerCapabilities{}

	data, err := json.Marshal(original)
	if err != nil {
		t.Fatalf("marshal failed: %v", err)
	}

	var payload map[string]any
	if err := json.Unmarshal(data, &payload); err != nil {
		t.Fatalf("unmarshal failed: %v", err)
	}

	if len(payload) != 0 {
		t.Fatalf("expected empty JSON object from zero-value capabilities, got: %s", string(data))
	}
}

func Test_CapabilitiesServer_UnmarshalExtendedFields(t *testing.T) {
	data := []byte(`{
		"textDocumentSync": {
			"openClose": true,
			"change": 2,
			"save": {"includeText": true}
		},
		"notebookDocumentSync": {
			"notebookSelector": [
				{
					"notebook": {"notebookType": "jupyter-notebook"},
					"cells": [{"language": "python"}]
				}
			],
			"save": true
		},
		"hoverProvider": {"workDoneProgress": true},
		"declarationProvider": true,
		"typeDefinitionProvider": {"workDoneProgress": true},
		"implementationProvider": true,
		"referencesProvider": {"workDoneProgress": true},
		"documentHighlightProvider": true,
		"documentSymbolProvider": {"workDoneProgress": true},
		"codeActionProvider": {"codeActionKinds": ["quickfix"], "resolveProvider": true},
		"codeLensProvider": {"resolveProvider": true},
		"documentLinkProvider": {"resolveProvider": true},
		"colorProvider": true,
		"documentFormattingProvider": true,
		"documentRangeFormattingProvider": true,
		"documentOnTypeFormattingProvider": {"firstTriggerCharacter": ";", "moreTriggerCharacter": [","]},
		"renameProvider": {"prepareProvider": true},
		"foldingRangeProvider": true,
		"executeCommandProvider": {"commands": ["laravel.run"], "workDoneProgress": true},
		"selectionRangeProvider": {"workDoneProgress": true},
		"linkedEditingRangeProvider": true,
		"callHierarchyProvider": true,
		"semanticTokensProvider": {
			"legend": {"tokenTypes": ["class"], "tokenModifiers": ["declaration"]},
			"full": true
		},
		"monikerProvider": true,
		"typeHierarchyProvider": true,
		"inlineValueProvider": true,
		"inlayHintProvider": {"resolveProvider": true},
		"workspaceSymbolProvider": {"resolveProvider": true},
		"workspace": {
			"workspaceFolders": {"supported": true, "changeNotifications": "workspaceFolders"},
			"fileOperations": {
				"didCreate": {
					"filters": [
						{"scheme": "file", "pattern": {"glob": "**/*.go"}}
					]
				}
			}
		}
	}`)

	var decoded protocol.ServerCapabilities
	if err := json.Unmarshal(data, &decoded); err != nil {
		t.Fatalf("unmarshal failed: %v", err)
	}

	if decoded.CodeLensProvider == nil || !decoded.CodeLensProvider.ResolveProvider {
		t.Fatalf("expected codeLensProvider.resolveProvider=true")
	}

	if decoded.DocumentLinkProvider == nil || !decoded.DocumentLinkProvider.ResolveProvider {
		t.Fatalf("expected documentLinkProvider.resolveProvider=true")
	}

	if decoded.DocumentOnTypeFormattingProvider == nil || decoded.DocumentOnTypeFormattingProvider.FirstTriggerCharacter != ";" {
		t.Fatalf("expected documentOnTypeFormattingProvider.firstTriggerCharacter=';' got %#v", decoded.DocumentOnTypeFormattingProvider)
	}

	if decoded.ExecuteCommandProvider == nil || len(decoded.ExecuteCommandProvider.Commands) != 1 || decoded.ExecuteCommandProvider.Commands[0] != "laravel.run" {
		t.Fatalf("expected executeCommandProvider.commands to contain laravel.run")
	}

	if decoded.Workspace == nil || decoded.Workspace.WorkspaceFolders == nil || !decoded.Workspace.WorkspaceFolders.Supported {
		t.Fatalf("expected workspace.workspaceFolders.supported=true")
	}

	tds, ok := decoded.TextDocumentSync.(map[string]any)
	if !ok {
		t.Fatalf("expected textDocumentSync object, got %#v", decoded.TextDocumentSync)
	}

	if openClose, ok := tds["openClose"].(bool); !ok || !openClose {
		t.Fatalf("expected textDocumentSync.openClose=true, got %#v", tds["openClose"])
	}

	hover, ok := decoded.HoverProvider.(map[string]any)
	if !ok {
		t.Fatalf("expected hoverProvider object, got %#v", decoded.HoverProvider)
	}

	if wdp, ok := hover["workDoneProgress"].(bool); !ok || !wdp {
		t.Fatalf("expected hoverProvider.workDoneProgress=true, got %#v", hover["workDoneProgress"])
	}

	if decl, ok := decoded.DeclarationProvider.(bool); !ok || !decl {
		t.Fatalf("expected declarationProvider=true, got %#v", decoded.DeclarationProvider)
	}

	if sem, ok := decoded.SemanticTokensProvider.(map[string]any); !ok || sem["legend"] == nil {
		t.Fatalf("expected semanticTokensProvider.legend, got %#v", decoded.SemanticTokensProvider)
	}
}

func Test_CapabilitiesServer_MarshalExtendedFields(t *testing.T) {
	resolve := true
	original := protocol.ServerCapabilities{
		PositionEncoding: &[]protocol.PositionEncodingKind{protocol.PositionEncodingKindUTF16}[0],
		TextDocumentSync: protocol.TextDocumentSyncOptions{
			OpenClose: true,
			Change:    protocol.TextDocumentSyncKindIncremental,
			Save:      protocol.SaveOptions{IncludeText: true},
		},
		HoverProvider: protocol.HoverOptions{WorkDoneProgress: true},
		CompletionProvider: &protocol.CompletionOptions{
			ResolveProvider: &resolve,
			CompletionItem: &protocol.CompletionOptionsCompletionItem{
				LabelDetailsSupport: &resolve,
			},
		},
		DeclarationProvider: protocol.DeclarationOptions{WorkDoneProgress: true},
		DefinitionProvider:  true,
		CodeActionProvider: protocol.CodeActionOptions{
			CodeActionKinds: []protocol.CodeActionKind{protocol.CodeActionQuickFix},
			ResolveProvider: true,
		},
		CodeLensProvider: &protocol.CodeLensOptions{ResolveProvider: true},
		DocumentLinkProvider: &protocol.DocumentLinkOptions{
			ResolveProvider: true,
		},
		DocumentOnTypeFormattingProvider: &protocol.DocumentOnTypeFormattingOptions{
			FirstTriggerCharacter: ";",
			MoreTriggerCharacter:  []string{","},
		},
		RenameProvider: protocol.RenameOptions{PrepareProvider: true},
		ExecuteCommandProvider: &protocol.ExecuteCommandOptions{
			Commands: []string{"laravel.run"},
		},
		SemanticTokensProvider: protocol.SemanticTokensOptions{
			Legend: protocol.SemanticTokensLegend{
				TokenTypes:     []string{"class"},
				TokenModifiers: []string{"declaration"},
			},
			Full: true,
		},
		DiagnosticProvider: protocol.DiagnosticOptions{
			InterFileDependencies: true,
			WorkspaceDiagnostics:  true,
		},
		WorkspaceSymbolProvider: protocol.WorkspaceSymbolOptions{ResolveProvider: true},
		Workspace: &protocol.WorkspaceServerCapabilities{
			WorkspaceFolders: &protocol.WorkspaceFoldersServerCapabilities{
				Supported:           true,
				ChangeNotifications: "workspaceFolders",
			},
			FileOperations: &protocol.FileOperationOptions{
				DidCreate: &protocol.FileOperationRegistrationOptions{
					Filters: []protocol.FileOperationFilter{{
						Scheme: "file",
						Pattern: protocol.FileOperationPattern{
							Glob: "**/*.go",
						},
					}},
				},
			},
		},
	}

	data, err := json.Marshal(original)
	if err != nil {
		t.Fatalf("marshal failed: %v", err)
	}

	var payload map[string]any
	if err := json.Unmarshal(data, &payload); err != nil {
		t.Fatalf("unmarshal failed: %v", err)
	}

	if _, ok := payload["textDocumentSync"]; !ok {
		t.Fatalf("expected textDocumentSync in payload: %s", string(data))
	}

	if _, ok := payload["hoverProvider"]; !ok {
		t.Fatalf("expected hoverProvider in payload: %s", string(data))
	}

	if _, ok := payload["declarationProvider"]; !ok {
		t.Fatalf("expected declarationProvider in payload: %s", string(data))
	}

	if _, ok := payload["codeLensProvider"]; !ok {
		t.Fatalf("expected codeLensProvider in payload: %s", string(data))
	}

	if _, ok := payload["documentOnTypeFormattingProvider"]; !ok {
		t.Fatalf("expected documentOnTypeFormattingProvider in payload: %s", string(data))
	}

	if _, ok := payload["semanticTokensProvider"]; !ok {
		t.Fatalf("expected semanticTokensProvider in payload: %s", string(data))
	}

	workspace, ok := payload["workspace"].(map[string]any)
	if !ok {
		t.Fatalf("expected workspace object, got %#v", payload["workspace"])
	}

	if _, ok := workspace["workspaceFolders"]; !ok {
		t.Fatalf("expected workspace.workspaceFolders in payload: %s", string(data))
	}

	fileOps, ok := workspace["fileOperations"].(map[string]any)
	if !ok {
		t.Fatalf("expected workspace.fileOperations object, got %#v", workspace["fileOperations"])
	}

	if _, ok := fileOps["didCreate"]; !ok {
		t.Fatalf("expected workspace.fileOperations.didCreate in payload: %s", string(data))
	}
}

package protocol

// Range of the LSP code space
const (
	LspReservedErrorRangeStart = -32899
	LspReservedErrorRangeEnd   = -32800
)

// LSP Specific error codes used with JSON RPC.
const (
	// Error code indicating that a server received a notification or
	// request before the server received the `initialize` request.
	//
	// For backwards compatibility these codes are not in the LSP range.
	RPCServerNotInitialized int64 = -32002
	RPCUnknownErrorCode     int64 = -32001

	// A request failed but it was syntactically correct, e.g the
	// method name was known and the parameters were valid. The error
	// message should contain human readable information about why
	// the request failed.
	//
	// @since 3.17.0
	RPCRequestFailed int64 = -32803

	// The server detected that the content of a document got
	// modified outside normal conditions. A server should
	// NOT send this error code if it detects a content change
	// in its unprocessed messages. The result even computed
	// on an older state might still be useful for the client.
	//
	// If a client decides that a result is not of any use anymore
	// the client should cancel the request.
	RPCContentModified int64 = -32801

	// The client has canceled a request and a server has detected
	// the cancel.
	RPCRequestCancelled int64 = -32800
)

func IsLspRPCErrorCode(code int64) bool {
	if code >= LspReservedErrorRangeStart && code <= LspReservedErrorRangeEnd {
		return true
	}

	if code == RPCServerNotInitialized || code == RPCUnknownErrorCode {
		return true
	}

	return false
}

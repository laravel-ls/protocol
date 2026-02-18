# Language Server Protocol for golang

Language Server Protocol implementation in Go.

## Design notes

The package does not include JSON RPC handling as that is better left to other packages.

## Example

Example code using [github.com/sourcegraph/jsonrpc2](https://github.com/sourcegraph/jsonrpc2)

```go
package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net"
	"os"

	"github.com/laravel-ls/protocol"
	"github.com/sourcegraph/jsonrpc2"
)

func handleInitialize(params protocol.InitializeParams) (protocol.InitializeResult, error) {
	// Initialization logic
	// Read more at: https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification/#initialize
	return protocol.InitializeResult{
		Capabilities: protocol.ServerCapabilities{
			TextDocumentSync: protocol.TextDocumentSyncKindFull,
		},
		ServerInfo: &protocol.ServerInfo{
			Name:    "My LSP Server",
			Version: "1.0.0",
		},
	}, nil
}

func handler(ctx context.Context, conn *jsonrpc2.Conn, req *jsonrpc2.Request) (any, error) {
	switch req.Method {
	case protocol.MethodInitialize:
		var params protocol.InitializeParams
		if err := json.Unmarshal(*req.Params, &params); err != nil {
			return nil, err
		}
		return handleInitialize(params)
	// ... add more methods
	default:
		// Respond with a method not found error
		return nil, &jsonrpc2.Error{
			Code:    jsonrpc2.CodeMethodNotFound,
			Message: fmt.Sprintf("Method %s not found", req.Method),
		}
	}
}

func handleConn(ctx context.Context, conn io.ReadWriteCloser) error {
	stream := jsonrpc2.NewBufferedStream(conn, jsonrpc2.VSCodeObjectCodec{})
	rpc := jsonrpc2.NewConn(ctx, stream, jsonrpc2.HandlerWithError(handler))

	select {
	case <-ctx.Done():
		return fmt.Errorf("context closed")
	case <-rpc.DisconnectNotify():
		return nil
	}
}

func main() {
	// example TCP server.
	ctx := context.Background()
	logger := log.New(os.Stderr, "[tcp-lsp] ", log.LstdFlags|log.Lmicroseconds)

	addr := "127.0.0.1:4389"
	l, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatal(err)
	}
	logger.Println("Listening on", addr)

	for {
		conn, err := l.Accept()
		if err != nil {
			logger.Println("accept error:", err)
			continue
		}

		logger.Println("Client connected:", conn.RemoteAddr())

		go func(c net.Conn) {
			defer c.Close()
			handleConn(ctx, c)
		}(conn)
	}
}
```

You can test your server with the following command on linux:

```bash
printf 'Content-Length: 58\r\n\r\n{"jsonrpc":"2.0","id":1,"method":"initialize","params":{}}' | nc 127.0.0.1 4389
```

# Author

Henrik Hautakoski <henrik@shufflingpixels.com>

# Language Server Protocol for golang

Language Server Protocol implementation in Go.

## Design notes

The package does not include JSON RPC handling as that is better left to other packages.

## Example

Example code using [github.com/sourcegraph/jsonrpc2](https://github.com/sourcegraph/jsonrpc2)

```go

import (
    "fmt"
    "io"
    "context"
    "encoding/json"
    "github.com/sourcegraph/jsonrpc2"
)

func handler(ctx context.Context, conn *jsonrpc2.Conn, req *jsonrpc2.Request) (any, error) {

    switch req.Method {
    case protocol.MethodInitialize:
        var params protocol.InitializeParams
        if err := json.Unmarshal(*req.Params, &params); err != nil {
            return nil, err
        }
    // handle initialize
    case protocol.MethodTextDocumentDidChange:
        var params protocol.DidChangeTextDocumentParams
        if err := json.Unmarshal(*req.Params, &params); err != nil {
            return nil, err
        }
    // handle did change
    default:
        // Respond with a method not found error
        return nil, &jsonrpc2.Error{
            Code:    jsonrpc2.CodeMethodNotFound,
            Message: fmt.Sprintf("Method %s not found", req.Method),
        }
    }
    return nil, nil
}

func (s Server) Run(ctx context.Context, conn io.ReadWriteCloser) error {
    stream := jsonrpc2.NewBufferedStream(conn, jsonrpc2.VSCodeObjectCodec{})
    rpc := jsonrpc2.NewConn(ctx, stream, jsonrpc2.HandlerWithError(s.dispatch))

    select {
    case <-ctx.Done():
        return fmt.Errorf("context closed")
    case <-rpc.DisconnectNotify():
        return nil
    }
}
```

# Author

Henrik Hautakoski <henrik@shufflingpixels.com>

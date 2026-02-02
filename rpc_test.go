package protocol_test

import (
	"testing"

	"github.com/laravel-ls/protocol"
)

func Test_IsLspRPCErrorCode(t *testing.T) {
	if protocol.IsLspRPCErrorCode(5000) {
		t.Errorf("5000 is not a valid code, but function returned true")
	}

	if protocol.IsLspRPCErrorCode(-9000) {
		t.Errorf("-9000 is not a valid code, but function returned true")
	}

	if protocol.IsLspRPCErrorCode(-32899) == false {
		t.Errorf("-32899 is a valid code, but function returned false")
	}

	if protocol.IsLspRPCErrorCode(-32840) == false {
		t.Errorf("-32840 is a valid code, but function returned false")
	}

	if protocol.IsLspRPCErrorCode(-32800) == false {
		t.Errorf("-32800 is a valid code, but function returned false")
	}

	// Exception
	if protocol.IsLspRPCErrorCode(-32001) == false {
		t.Errorf("-32001 is a valid code, but function returned false")
	}

	if protocol.IsLspRPCErrorCode(-32002) == false {
		t.Errorf("-32002 is a valid code, but function returned false")
	}
}

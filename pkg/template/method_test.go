package template

import "testing"

func TestMethodRender(t *testing.T) {
	m := Method{
		ServiceName: "TestService",
		MethodName:  "TestMethod",
		InputType:   "TestRequest",
		OutputType:  "TestResponse",
		RequestType: "TestRPCRequest",
		Errorf:      "TestErrorf",
	}

	expected := `
func (x *RPCRedisTestService) handleTestMethod(req TestRPCRequest) (any, error) {
	var rpcReq TestRequest

	err := req.ParseParams(&rpcReq)
	if err != nil {
		return nil, TestErrorf("error parsing request: %v", err)
	}

	return x.service.TestMethod(req.Context(), &rpcReq)
}
`

	result, err := m.Render()
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if result != expected {
		t.Errorf("Rendered result does not match expected:\nExpected:\n%s\n\nGot:\n%s", expected, result)
	}
}

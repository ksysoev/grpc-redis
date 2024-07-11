package gen

import (
	"testing"

	"google.golang.org/protobuf/compiler/protogen"
)

func TestGenerateMethod(t *testing.T) {
	plugin := &protogen.Plugin{}
	g := plugin.NewGeneratedFile("test.go", "test")
	method := &protogen.Method{
		Parent: &protogen.Service{
			GoName: "TestService",
		},
		GoName: "TestMethod",
		Input: &protogen.Message{
			GoIdent: protogen.GoIdent{
				GoName:       "TestInput",
				GoImportPath: "github.com/ksysoev/grpc-redis/pkg/gen",
			},
		},
		Output: &protogen.Message{
			GoIdent: protogen.GoIdent{
				GoName:       "TestOutput",
				GoImportPath: "github.com/ksysoev/grpc-redis/pkg/gen",
			},
		},
	}

	expectedMethod := `
func (x *RPCRedisTestService) handleTestMethod(req redis_rpc.Request) (any, error) {
	var rpcReq gen.TestInput

	err := req.ParseParams(&rpcReq)
	if err != nil {
		return nil, fmt.Errorf("error parsing request: %v", err)
	}

	return x.service.TestMethod(req.Context(), &rpcReq)
}
`

	renderedMethod, err := generateMethod(g, method)
	if err != nil {
		t.Fatalf("failed to generate method: %v", err)
	}

	if renderedMethod != expectedMethod {
		t.Fatalf("invalid method code, expected:\n%s\ngot:\n%s", expectedMethod, renderedMethod)
	}
}

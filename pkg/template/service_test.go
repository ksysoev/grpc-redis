package template

import "testing"

func TestRender(t *testing.T) {
	s := &Service{
		ServiceName:  "TestService",
		FullName:     "com.example.TestService",
		RedisClient:  "redisClient",
		RPCServer:    "rpcServer",
		RPCServerOpt: "ServerOption",
		NewRPCServer: "newRPCServer",
		NewUUID:      "newUUID",
		Methods:      []string{"Method1", "Method2"},
	}

	expected := `
// TestService is the server API for com.example.TestService
type RPCRedisTestService struct {
	rpcSever *rpcServer
	service  *TestServiceService
}

func NewRedisTestService(redis *redisClient, grpcService *TestServiceService, opts ...ServerOption) *RPCRedisTestService {
	rpcServer := newRPCServer(redis, "com.example.TestService", "TestServiceGroup", newUUID().String(), opts...)
	service := &RPCRedisTestService{
		rpcSever: rpcServer,
		service:  grpcService,
	}

	// Register handlers
	rpcServer.AddHandler("Method1", service.handleMethod1)
	rpcServer.AddHandler("Method2", service.handleMethod2)

	return service
}

func (x *RPCRedisTestService) Serve() error {
	return x.rpcSever.Run()
}

func (x *RPCRedisTestService) Close() {
	x.rpcSever.Close()
}
`

	result, err := s.Render()
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if result != expected {
		t.Errorf("Rendered result does not match expected:\nExpected:\n%s\n\nGot:\n%s", expected, result)
	}
}

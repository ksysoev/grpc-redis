package template

import (
	"strings"
	"text/template"
)

const serviceTemplate = `
// {{.ServiceName}} is the server API for {{.FullName}}
type RPCRedis{{.ServiceName}} struct {
    rpcSever *{{.RPCServer}}
	service  *{{.ServiceName}}Service
}

func NewRedis{{.ServiceName}}(redis *{{.RedisClient}}, grpcService *{{.ServiceName}}Service) *RPCRedis{{.ServiceName}} {
	rpcServer := {{.NewRPCServer}}(redis, "{{.FullName}}", "{{.ServiceName}}Group", {{.NewUUID}}().String())
	service := &RPCRedis{{.ServiceName}}{
		rpcSever: rpcServer,
		service:  grpcService,
	}

	// Register handlers
	{{- range .Methods}}
	rpcServer.AddHandler("{{.}}", service.handle{{.}})
	{{- end}}

	return service
}

func (x *RPCRedis{{.ServiceName}}) Serve() error {
	return x.rpcSever.Run()
}

func (x *RPCRedis{{.ServiceName}}) Close() {
	x.rpcSever.Close()
}
`

var tmplService *template.Template

func init() {
	tmplService = template.Must(template.New("service").Parse(serviceTemplate))
}

// Service represents a service that provides various methods.
type Service struct {
	ServiceName  string   // The name of the service.
	FullName     string   // The full name of the service.
	RedisClient  string   // The Redis client used by the service.
	RPCServer    string   // The RPC server used by the service.
	NewRPCServer string   // The new RPC server used by the service.
	NewUUID      string   // The new UUID used by the service.
	Methods      []string // The list of methods provided by the service.
}

// Render renders the service template and returns the rendered string.
// It uses the tmplService template and populates it with the data from the Service struct.
// If an error occurs during rendering, it returns an empty string and the error.
func (s Service) Render() (string, error) {
	var buf strings.Builder
	if err := tmplService.Execute(&buf, s); err != nil {
		return "", err
	}
	return buf.String(), nil
}

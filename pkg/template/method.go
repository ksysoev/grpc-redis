package template

import (
	"strings"
	"text/template"
)

const methodTemplate = `
func (x *RPCRedis{{.ServiceName}}) handle{{.MethodName}}(req {{.RequestType}}) (any, error) {
	var rpcReq {{.InputType}}

	err := req.ParseParams(&rpcReq)
	if err != nil {
		return nil, {{.Errorf}}("error parsing request: %v", err)
	}

	return x.service.{{.MethodName}}(req.Context(), &rpcReq)
}
`

var tmplMethod *template.Template

func init() {
	tmplMethod = template.Must(template.New("method").Parse(methodTemplate))
}

// Method represents a gRPC method.
type Method struct {
	ServiceName string // The name of the service that the method belongs to.
	MethodName  string // The name of the method.
	InputType   string // The type of the input message.
	OutputType  string // The type of the output message.
	RequestType string // The type of the request message.
	Errorf      string // The error format string.
}

// Render renders the Method struct into a string representation.
// It uses the tmplService template to execute the rendering process.
// Returns the rendered string and any error encountered during rendering.
func (s *Method) Render() (string, error) {
	var buf strings.Builder
	if err := tmplMethod.Execute(&buf, s); err != nil {
		return "", err
	}

	return buf.String(), nil
}

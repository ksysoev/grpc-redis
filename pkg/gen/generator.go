package gen

import (
	"fmt"

	tmpl "github.com/ksysoev/grpc-redis/pkg/template"
	"google.golang.org/protobuf/compiler/protogen"
)

// fileName is the name of the generated file.
const fileName = "_grpc-redis.pb.go"

// Generate generates the code for the given protocol buffer file.
// It takes a protogen.Plugin and a protogen.File as input and returns an error if any.
// The generated code includes the file header, services, and methods.
// For each service, it generates the service definition and for each method, it generates the method definition.
// If a method is a streaming client or a streaming server, it is skipped.
// The generated code is written to the specified file.
func Generate(gen *protogen.Plugin, file *protogen.File) error {
	filename := file.GeneratedFilenamePrefix + fileName
	g := gen.NewGeneratedFile(filename, file.GoImportPath)

	fileHeader, err := generateFileHeader(file)
	if err != nil {
		return fmt.Errorf("error generating file %s: %v", filename, err)
	}

	g.P(fileHeader)

	for _, service := range file.Services {
		generatedService, err := generateService(g, service)
		if err != nil {
			return fmt.Errorf("error generating file %s: %v", filename, err)
		}

		g.P(generatedService)

		for _, method := range service.Methods {
			if method.Desc.IsStreamingClient() || method.Desc.IsStreamingServer() {
				continue
			}

			generatedMethod, err := generateMethod(g, method)
			if err != nil {
				return fmt.Errorf("error generating file %s: %v", filename, err)
			}

			g.P(generatedMethod)
		}
	}

	return nil
}

// generateFileHeader generates the header content for the generated file based on the provided protogen.File.
// It returns the rendered file header as a string.
// An error is returned if there is any issue generating the file header.
func generateFileHeader(file *protogen.File) (string, error) {
	tmplFile := tmpl.File{
		PackageName: string(file.GoPackageName),
	}

	fileHeaderRender, err := tmplFile.RenderHeader()
	if err != nil {
		return "", fmt.Errorf("error generating file header: %v", err)
	}

	return fileHeaderRender, nil
}

// generateService generates the code for a service based on the provided protogen.GeneratedFile and protogen.Service.
// It returns the rendered service code as a string and an error if there was any issue generating the code.
func generateService(g *protogen.GeneratedFile, service *protogen.Service) (string, error) {
	methods := make([]string, 0, len(service.Methods))
	for _, method := range service.Methods {
		methods = append(methods, method.GoName)
	}

	tmplService := tmpl.Service{
		ServiceName:  service.GoName,
		FullName:     string(service.Desc.FullName()),
		RedisClient:  g.QualifiedGoIdent(protogen.GoIdent{GoName: "Client", GoImportPath: "github.com/redis/go-redis/v9"}),
		RPCServer:    g.QualifiedGoIdent(protogen.GoIdent{GoName: "Server", GoImportPath: "github.com/ksysoev/redis-rpc"}),
		NewRPCServer: g.QualifiedGoIdent(protogen.GoIdent{GoName: "NewServer", GoImportPath: "github.com/ksysoev/redis-rpc"}),
		NewUUID:      g.QualifiedGoIdent(protogen.GoIdent{GoName: "New", GoImportPath: "github.com/google/uuid"}),
		Methods:      methods,
	}

	svcRender, err := tmplService.Render()
	if err != nil {
		return "", fmt.Errorf("error generating service %s: %v", service.GoName, err)
	}

	return svcRender, nil
}

// generateMethod generates a method for the given protogen.GeneratedFile and protogen.Method.
// It returns the rendered method as a string and an error if there was any issue during generation.
func generateMethod(g *protogen.GeneratedFile, method *protogen.Method) (string, error) {
	tmplMethod := tmpl.Method{
		ServiceName: method.Parent.GoName,
		MethodName:  method.GoName,
		InputType:   g.QualifiedGoIdent(method.Input.GoIdent),
		OutputType:  g.QualifiedGoIdent(method.Output.GoIdent),
		RequestType: g.QualifiedGoIdent(protogen.GoIdent{GoName: "Request", GoImportPath: "github.com/ksysoev/redis-rpc"}),
		Errorf:      g.QualifiedGoIdent(protogen.GoIdent{GoName: "Errorf", GoImportPath: "fmt"}),
	}

	methodRender, err := tmplMethod.Render()
	if err != nil {
		return "", fmt.Errorf("error generating method %s: %v", method.GoName, err)
	}

	return methodRender, nil
}

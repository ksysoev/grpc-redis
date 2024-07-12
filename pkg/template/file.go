package template

import (
	"strings"
	"text/template"
)

const fileHeaderTemplate = `
// Code generated by protoc-gen-rpc-redis. DO NOT EDIT.

package {{ .PackageName }}
`

var tmplFileHeader *template.Template

func init() {
	tmplFileHeader = template.Must(template.New("fileHeader").Parse(fileHeaderTemplate))
}

type File struct {
	PackageName string
}

// Render renders the file template and returns the rendered string.
// It executes the tmplFileHeader template with the given File object and writes the result to a string builder.
// If an error occurs during template execution, it returns an empty string and the error.
func (f File) RenderHeader() (string, error) {
	var buf strings.Builder
	if err := tmplFileHeader.Execute(&buf, f); err != nil {
		return "", err
	}
	return buf.String(), nil
}

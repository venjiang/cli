package golang

import (
	"bytes"
	_ "embed"
	"text/template"
)

//go:embed templates/main_rx.tmpl
var MainFuncRxTmpl []byte

//go:embed templates/main_raw_bytes.tmpl
var MainFuncRawBytesTmpl []byte

//go:embed templates/init.tmpl
var InitFuncTmpl []byte

// Context defines context for the template
type Context struct {
	// Name of the servcie
	Name string
	// ZipperAddrs is the addresses of the zipper server
	ZipperAddrs []string
	// Client credential
	Credential string
}

// RenderTmpl renders the template with the given context
func RenderTmpl(tpl string, ctx *Context) ([]byte, error) {
	t := template.Must(template.New("tpl").Parse(tpl))
	buf := bytes.NewBuffer([]byte{})
	err := t.Execute(buf, ctx)
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

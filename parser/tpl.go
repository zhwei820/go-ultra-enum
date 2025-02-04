package parser

const Header = `// Code generated by go-ultra-enum; DO NOT EDIT.
	package {{.Name}}

`
const Tpl = `type {{.NewName | LcFirst}}Enum struct{
{{ range $e := .Elements}}
{{$e.Name}} {{$.NewName | LcFirst}}Ele{{ end}}
}


type {{.NewName | LcFirst}}Ele struct {
    Name string
    value {{.Tpe}}
    Description string
}

func (x *{{.NewName | LcFirst}}Ele ) Val() {{.Tpe}} {
	return x.value
}

var (
    {{.NewName}}Enum = {{.NewName | LcFirst }}Enum{
        {{- range $e := .Elements}}
            {{.Name}} : {{$.NewName | LcFirst}}Ele{Name: "{{.Name}}", value: {{if eq .Tpe "string"}}"{{.Value}}"{{else}}{{.Value}}{{end}}, Description: "{{.Description}}"},
        {{- end}}
    }
    
)
`

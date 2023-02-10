package sql2go

// https://github.com/go-programming-tour-book/tour/blob/master/internal/sql2struct/template.go
const strcutTpl = `
type {{.TableName | ToCamelCase}} struct {
	{{range .Columns}}	{{ $length := len .Comment}} {{ if gt $length 0 }}// {{.Comment}} {{else}}// {{.Name}} {{ end }}
		{{ $typeLen := len .Type }} {{ if gt $typeLen 0 }}{{.Name | ToCamelCase}}	{{.Type}}	{{.Tag}}{{ else }}{{.Name}}{{ end }}
	{{end}}}

func (model {{.TableName | ToCamelCase}}) TableName() string {
return "{{.TableName}}"
}`

// 结构体模板对象；
type StructTemplate struct {
	structTpl string
}

func NewStructTemplate() *StructTemplate {
	return &StructTemplate{structTpl: strcutTpl}
}

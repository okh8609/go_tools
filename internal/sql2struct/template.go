package sql2struct

import (
	"fmt"
	"os"
	"text/template"

	"github.com/okh8609/go_tools/internal/word"
)

var DBType2GoType = map[string]string{ // mysql-type : golang-type
	"int":        "int32",
	"tinyint":    "int8",
	"smallint":   "int",
	"mediumint":  "int64",
	"bigint":     "int64",
	"bit":        "int",
	"bool":       "bool",
	"enum":       "string",
	"set":        "string",
	"varchar":    "string",
	"char":       "string",
	"tinytext":   "string",
	"mediumtext": "string",
	"text":       "string",
	"longtext":   "string",
	"blob":       "string",
	"tinyblob":   "string",
	"mediumblob": "string",
	"longblob":   "string",
	"date":       "time.Time",
	"datetime":   "time.Time",
	"timestamp":  "time.Time",
	"time":       "time.Time",
	"float":      "float64",
	"double":     "float64",
}

const templateStr = `

type {{.TableName | ToCamelCase}} struct {
	{{range .Columns}}
		{{ $length := len .Comment}} {{ if gt $length 0 }}// {{.Comment}} {{else}}// {{.Name}} {{ end }}
		{{ $typeLen := len .Type }} {{ if gt $typeLen 0 }}{{.Name | ToCamelCase}}	{{.Type}}	{{.Tag}}{{ else }}{{.Name}}{{ end }}
	{{end}}}

func (tt {{.TableName | ToCamelCase}}) TableName() string {
	return "{{.TableName}}"
}

`

type StructMember struct {
	Name    string
	Type    string
	Tag     string
	Comment string
}

func Generate(tableName string, tableMembers []*TableMember) error {
	// 準備 要放入template的物件清單
	var structMembers []*StructMember
	for _, mbr := range tableMembers {
		structMembers = append(structMembers, &StructMember{
			Name:    mbr.Name,
			Type:    DBType2GoType[mbr.DataType],
			Tag:     fmt.Sprintf("`json:\"%s\"`", mbr.Name),
			Comment: mbr.Comment,
		})
	}

	// 放入template 並輸出
	tt := template.Must(template.New("sql2struct").Funcs(template.FuncMap{"ToCamelCase": word.Underscore_To_UpperCamelCase}).Parse(templateStr))

	ttData := struct {
		TableName string
		Columns   []*StructMember
	}{
		TableName: tableName,
		Columns:   structMembers,
	}

	err := tt.Execute(os.Stdout, ttData)
	return err
}

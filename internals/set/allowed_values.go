package set

import (
	"fmt"
	"sort"
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"

	"github.com/iostrovok/kibana-fields-generator/internals/face"
)

func AllowedValues(fieldName string, allowedValues []*face.AllowedValue) string {
	if len(allowedValues) == 0 {
		return ""
	}

	def := make([]string, 0)
	val := make([]string, 0)
	for i := range allowedValues {
		e := allowedValues[i]
		out := ""
		f := varNameReg.Split(e.Name, -1)
		for j := range f {
			switch f[j] {
			case "id":
				f[j] = "ID"
			default:
				f[j] = cases.Title(language.English, cases.NoLower).String(f[j])
			}
		}

		out += strings.Join(f, "")

		def = append(def, fmt.Sprintf("%s string // %s\n", out, strings.Replace(e.Description, "\n", " ", -1)))
		val = append(val, fmt.Sprintf("%s : `%s`,\n", out, e.Name))
	}

	sort.Strings(def)
	sort.Strings(val)

	typeName := fieldName + "AllowedType"
	valuesName := fieldName + "AllowedValues"

	return `
type ` + typeName + ` struct {
` + strings.Join(def, "") + `
}

var ` + valuesName + ` ` + typeName + ` = ` + typeName + `{
` + strings.Join(val, "") + `
}
`
}

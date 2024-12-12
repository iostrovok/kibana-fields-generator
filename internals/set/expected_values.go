package set

import (
	"fmt"
	"sort"
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func ExpectedValues(fieldName string, expectedValues []string) string {
	if len(expectedValues) == 0 {
		return ""
	}

	def := make([]string, 0)
	val := make([]string, 0)
	for i := range expectedValues {
		e := expectedValues[i]
		out := ""
		f := varNameReg.Split(e, -1)
		for j := range f {
			switch f[j] {
			case "id":
				f[j] = "ID"
			default:
				f[j] = cases.Title(language.English, cases.NoLower).String(f[j])
			}
		}
		out += strings.Join(f, "")
		def = append(def, fmt.Sprintf("%s string\n", out))
		val = append(val, fmt.Sprintf("%s : `%s`,\n", out, e))
	}

	sort.Strings(def)
	sort.Strings(val)

	typeName := fieldName + "ExpectedType"
	valuesName := fieldName + "ExpectedValues"

	return `
type ` + typeName + ` struct {
` + strings.Join(def, "") + `
}

var ` + valuesName + ` ` + typeName + ` = ` + typeName + `{
` + strings.Join(val, "") + `
}
`
}

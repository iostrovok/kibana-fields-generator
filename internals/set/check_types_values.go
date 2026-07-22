package set

import (
	"fmt"
	"sort"
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func toCamelCase(s string) string {
	out := ""
	for _, v := range strings.Split(s, "_") {
		out += cases.Title(language.English, cases.NoLower).String(v)
	}
	return out
}

func CheckValues(constData []*ConstData) ([]string, string, error) {
	if len(constData) == 0 {
		return nil, "", nil
	}

	def := make([]string, 0)
	fieldSet := make([]string, 0)
	for i := range constData {
		e := constData[i]

		f := varNameReg.Split(e.Name, -1)
		for j := range f {
			if strings.ToLower(f[j]) == "id" {
				f[j] = "ID"
			} else {
				f[j] = cases.Title(language.English, cases.NoLower).String(f[j])
			}
		}

		t := toCamelCase(e.Type)

		switch strings.ToLower(t) {
		case "ip":
			t = "IP"
		}

		fieldSet = append(fieldSet, t)
		def = append(def, fmt.Sprintf("%s %s\n", strings.Join(f, ""), "fields."+t))
	}

	sort.Strings(def)

	return fieldSet, "type TypesType struct {\n" + strings.Join(def, "") + "\n}\nvar Types TypesType = TypesType{}\n", nil
}

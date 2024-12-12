package set

import (
	"fmt"
	"sort"
	"strings"

	"github.com/pkg/errors"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func CheckValues(constData []*ConstData) (string, error) {
	if len(constData) == 0 {
		return "", nil
	}

	def := make([]string, 0)
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

		t := ""
		switch e.Type {
		case "boolean":
			t = "fields.Boolean"
		case "constant_keyword":
			t = "fields.ConstantKeyWord"
		case "date":
			t = "fields.Date"
		case "flattened":
			t = "fields.Flattened"
		case "float":
			t = "fields.Float"
		case "geo_point":
			t = "fields.GeoPoint"
		case "ip":
			t = "fields.IP"
		case "keyword":
			t = "fields.KeyWord"
		case "long":
			t = "fields.Long"
		case "match_only_text":
			t = "fields.MatchOnlyText"
		case "nested":
			t = "fields.Nested"
		case "object":
			t = "fields.Object"
		case "scaled_float":
			t = "fields.Float"
		case "wildcard":
			t = "fields.Wildcard"
		default:
			return "", errors.New("unknown type: " + e.Type)
		}

		def = append(def, fmt.Sprintf("%s %s\n", strings.Join(f, ""), t))
	}

	sort.Strings(def)

	return "type TypesType struct {\n" + strings.Join(def, "") + "\n}\nvar Types TypesType = TypesType{}\n", nil
}

package set

import (
	"fmt"
	"go/format"
	"os"
	"path"
	"regexp"
	"sort"
	"strings"

	"github.com/pkg/errors"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"

	"github.com/iostrovok/kibana-fields-generator/internals/face"
)

var varNameReg = regexp.MustCompile(`[^a-zA-Z0-9]+`)

type ConstData struct {
	Name, Value, Comment, Type string
}

func RunOneSet(setName string, set *face.Set, outPutPath, template string) (err error) {
	constsData := make([]*ConstData, 0)
	fieldsData := make([]string, 0)
	expectedAndAllowedValues := make([]string, 0)

	if err := os.MkdirAll(outPutPath, 0755); err != nil {
		return errors.WithStack(err)
	}

	for fieldName := range set.Fields {
		field := set.Fields[fieldName]
		name := getNameFromField(setName, field.FlatName)
		description := field.Description
		if len(field.Short) > 10 {
			description = field.Short
		}
		constsData = append(constsData, &ConstData{
			Name:    name,
			Value:   field.FlatName,
			Comment: strings.Replace(description, "\n", " ", -1),
			Type:    field.Type,
		})
		fieldsData = append(fieldsData, name)

		if e := ExpectedValues(name, field.ExpectedValues); e != "" {
			expectedAndAllowedValues = append(expectedAndAllowedValues, e)
		}
		if e := AllowedValues(name, field.AllowedValues); e != "" {
			expectedAndAllowedValues = append(expectedAndAllowedValues, e)
		}
	}

	typesChecker, err := CheckValues(constsData)
	if err != nil {
		return err
	}

	sort.Strings(fieldsData)
	sort.Slice(constsData, func(i, j int) bool {
		return constsData[i].Name < constsData[j].Name
	})
	sort.Slice(expectedAndAllowedValues, func(i, j int) bool {
		return expectedAndAllowedValues[i] < expectedAndAllowedValues[j]
	})

	consts := make([]string, len(constsData))
	for i := range constsData {
		consts[i] = fmt.Sprintf("%s fields.Field = \"%s\" // %s\n", constsData[i].Name, constsData[i].Value, constsData[i].Comment)
	}

	pkgName := checkDefaultGoNames(setName)
	template = strings.Replace(template, "<PACKAGE_NAME>", pkgName, 1)
	template = strings.Replace(template, "<CONST>", strings.Join(consts, ""), 1)
	template = strings.Replace(template, "<FIELDS>", strings.Join(fieldsData, ",\n")+",", 1)
	template = strings.Replace(template, "<TYPES_CHECKER>", typesChecker, 1)
	template = strings.Replace(template, "<EXPECTED_AND_ALLOWED_VALUES>", strings.Join(expectedAndAllowedValues, ""), 1)

	formatTemplate, err := format.Source([]byte(template))
	if err != nil {
		return errors.WithStack(err)
	}

	dir := path.Join(outPutPath, pkgName)
	fmt.Printf("save %s to %s\n", setName, dir)
	if err := checkDirFile(dir); err != nil {
		return errors.WithStack(err)
	}

	fileName := path.Join(outPutPath, pkgName, pkgName+".go")
	return errors.WithStack(os.WriteFile(fileName, formatTemplate, 0755))
}

func checkDefaultGoNames(setName string) string {
	if setName == "package" {
		return "pkg"
	}

	if setName == "interface" {
		return "iface"
	}

	return setName
}

func checkDirFile(dir string) error {
	// for test

	_, err := os.ReadDir(dir)
	if os.IsExist(err) || err == nil {
		return nil
	}

	if os.IsNotExist(err) {
		return errors.WithStack(os.Mkdir(dir, 0755))
	}

	return errors.WithStack(err)
}

func skipSet(set *face.Set, field *face.Field) bool {
	switch field.Type {
	case "nested", "object":
		// need to set up internals of this types
		if set.Name != "base" {
			return true
		}
	}

	return false
}

func getNameFromField(setName, field string) string {
	out := ""

	//fields := strings.Split(field, ".")
	fields := varNameReg.Split(field, -1)

	for i := range fields {
		if strings.ToLower(fields[i]) == strings.ToLower(setName) {
			continue
		}

		f := strings.Split(fields[i], "_")
		for i := range f {
			switch f[i] {
			case "id":
				f[i] = "ID"
			default:
				f[i] = cases.Title(language.English, cases.NoLower).String(f[i])
			}
		}

		out += strings.Join(f, "")
	}

	return out
}

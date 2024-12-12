package set

import (
	"go/format"
	"os"
	"path"
	"sort"
	"strings"

	"github.com/pkg/errors"
)

func SaveCheckFile(packagePath, outPutPath, template string, sets []string) (err error) {
	if err := os.MkdirAll(outPutPath, 0755); err != nil {
		return errors.WithStack(err)
	}

	outCheckPutPath := path.Join(outPutPath, "check")
	if err := os.MkdirAll(outCheckPutPath, 0755); err != nil {
		return errors.WithStack(err)
	}

	sort.Strings(sets)

	imports := make([]string, 0)
	allFields := make([]string, 0)
	for i := range sets {
		name := checkDefaultGoNames(sets[i])

		allFields = append(allFields, "AllFields = append(AllFields, "+name+".Fields...)\n")
		imports = append(imports, "\""+packagePath+"/"+name+"\"\n")
	}

	template = strings.Replace(template, "<IMPORTS>", strings.Join(imports, ""), 1)
	template = strings.Replace(template, "<ALL_FIELDS>", strings.Join(allFields, ""), 1)

	byteTemplate, err := format.Source([]byte(template))
	if err != nil {
		return err
	}

	fileName := path.Join(outCheckPutPath, "check.go")
	return os.WriteFile(fileName, byteTemplate, 0755)
}

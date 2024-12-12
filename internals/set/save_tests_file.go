package set

import (
	"go/format"
	"os"
	"path"
	"sort"
	"strings"
)

func SaveTestFile(packagePath, outPutPath, template string, sets []string) (err error) {
	sort.Strings(sets)

	imports := make([]string, 0)
	for i := range sets {
		imports = append(imports, "_ \""+packagePath+"/"+checkDefaultGoNames(sets[i])+"\"\n")
	}

	template = strings.Replace(template, "<IMPORTS>", strings.Join(imports, ""), 1)
	byteTemplate, err := format.Source([]byte(template))
	if err != nil {
		return err
	}

	fileName := path.Join(outPutPath, "syntax_test.go")
	return os.WriteFile(fileName, byteTemplate, 0755)
}

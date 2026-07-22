package set

import (
	"fmt"
	"go/format"
	"os"
	"path"
	"sort"
	"strings"

	"github.com/pkg/errors"
)

func SaveTestFile(packagePath, outPutPath, template, pkg string, sets []string) (err error) {
	sort.Strings(sets)

	imports := make([]string, 0)
	for i := range sets {
		imports = append(imports, "_ \""+packagePath+"/"+checkDefaultGoNames(sets[i])+"\"\n")
	}

	template = strings.Replace(template, "<IMPORTS>", strings.Join(imports, ""), 1)
	template = strings.Replace(template, "<PACKAGE>", pkg, 1)
	byteTemplate, err := format.Source([]byte(template))
	if err != nil {
		return err
	}

	fileName := path.Join(outPutPath, "syntax_test.go")
	return os.WriteFile(fileName, byteTemplate, 0755)
}

func RemoveTestFile(outPutPath string) error {
	fileName := path.Join(outPutPath, "syntax_test.go")

	fmt.Printf("RemoveTestFile fileName: %s\n", fileName)

	if err := os.RemoveAll(fileName); err != nil {
		return errors.WithStack(err)
	}

	return nil
}

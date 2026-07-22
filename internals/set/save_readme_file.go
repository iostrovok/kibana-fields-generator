package set

import (
	"fmt"
	"os"
	"path"
	"strings"

	"github.com/pkg/errors"
)

const ReadmeFile = "README.md"

func SaveReadmeFile(outPutPath, template, version string) error {
	if err := os.MkdirAll(outPutPath, 0755); err != nil {
		return errors.WithStack(err)
	}

	template = strings.ReplaceAll(template, "<VERSION>", version)
	fileName := path.Join(outPutPath, ReadmeFile)
	fmt.Printf("fileName: %s", fileName)

	return os.WriteFile(fileName, []byte(template), 0777)
}

func RemoveReadmeFile(outPutPath string) error {
	fileName := path.Join(outPutPath, ReadmeFile)

	fmt.Printf("RemoveReadmeFile fileName: %s\n", fileName)

	if err := os.RemoveAll(fileName); err != nil {
		return errors.WithStack(err)
	}

	return nil
}

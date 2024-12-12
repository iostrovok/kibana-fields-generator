package set

import (
	"os"
	"path"
	"strings"

	"github.com/pkg/errors"
)

func SaveReadmeFile(outPutPath, template, version string) error {
	if err := os.MkdirAll(outPutPath, 0755); err != nil {
		return errors.WithStack(err)
	}

	template = strings.ReplaceAll(template, "<VERSION>", version)
	fileName := path.Join(outPutPath, "README.md")

	return os.WriteFile(fileName, []byte(template), 0555)
}

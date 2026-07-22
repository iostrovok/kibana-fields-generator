package set

import (
	"fmt"
	"go/format"
	"os"
	"path"
	"strings"

	"github.com/pkg/errors"
)

const FieldsGoFile = "fields.go"

func SaveFieldsFile(outPutPath, template, fields string) error {
	if err := os.MkdirAll(outPutPath, 0755); err != nil {
		return errors.WithStack(err)
	}

	template = strings.ReplaceAll(template, "<ALL_FIELDS>", fields)
	byteTemplate, err := format.Source([]byte(template))
	if err != nil {
		return err
	}

	fileName := path.Join(outPutPath, FieldsGoFile)
	fmt.Printf("fileName: %s", fileName)
	return os.WriteFile(fileName, []byte(byteTemplate), 0777)
}

func RemoveFieldsFile(outPutPath string) error {
	fileName := path.Join(outPutPath, FieldsGoFile)
	fmt.Printf("RemoveFieldsFile fileName: %s\n", fileName)

	if err := os.RemoveAll(fileName); err != nil {
		return errors.WithStack(err)
	}

	return nil
}

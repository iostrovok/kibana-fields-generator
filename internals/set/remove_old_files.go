package set

import (
	"github.com/pkg/errors"
)

func RemoveAll(outPutPath, outPackagePutPath string) error {
	if err := RemoveReadmeFile(outPutPath); err != nil {
		return errors.WithStack(err)
	}

	if err := RemoveFieldsFile(outPutPath); err != nil {
		return errors.WithStack(err)
	}

	if err := RemoveTestFile(outPutPath); err != nil {
		return errors.WithStack(err)
	}

	if err := RemoveCheckFile(outPutPath); err != nil {
		return errors.WithStack(err)
	}

	if err := RemoveAllSets(outPackagePutPath); err != nil {
		return errors.WithStack(err)
	}

	return nil
}

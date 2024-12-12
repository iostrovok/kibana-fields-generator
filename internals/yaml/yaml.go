package yaml

import (
	"io"
	"net/http"
	"os"

	"gopkg.in/yaml.v3"

	"github.com/iostrovok/kibana-fields-generator/internals/face"
)

func Load(url string) (map[string]*face.Set, error) {
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()
	return ReadYML(res.Body)
}

func Read(fileName string) (map[string]*face.Set, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}

	out := map[string]*face.Set{}

	decoder := yaml.NewDecoder(file)
	decoder.KnownFields(true)
	if err := decoder.Decode(&out); err != nil {
		return nil, err
	}

	return out, nil
}

func ReadYML(file io.Reader) (map[string]*face.Set, error) {
	out := map[string]*face.Set{}

	decoder := yaml.NewDecoder(file)
	decoder.KnownFields(true)
	if err := decoder.Decode(&out); err != nil {
		return nil, err
	}

	return out, nil
}

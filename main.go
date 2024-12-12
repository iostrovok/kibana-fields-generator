package main

import (
	"embed"
	"flag"
	"fmt"
	"sort"
	"strings"
	"time"

	"github.com/iostrovok/kibana-fields-generator/internals/face"
	"github.com/iostrovok/kibana-fields-generator/internals/set"
	"github.com/iostrovok/kibana-fields-generator/internals/yaml"
)

const (
	importPath     = "github.com/iostrovok/kibana-fields"
	source         = "https://raw.githubusercontent.com/elastic/ecs/<VERSION>/generated/ecs/ecs_nested.yml"
	defaultVersion = "8.12"
)

var (
	savePath      = flag.String("path", "", `path for saving generated files`)
	version       = flag.String("version", "8.12", `version of ecs`)
	pkgImportPath = importPath + "/x"
	pkgSavePath   = ""
)

//go:embed templates/template.txt
//go:embed templates/syntax_test.txt
//go:embed templates/check.txt
//go:embed templates/README.md
var FILES embed.FS

type Set struct {
	Indexed       bool
	FieldSet      string
	Field         string
	Type          string
	Level         string
	Normalization string
	Example       string
	Description   string
}

func init() {
	flag.Parse()
	pkgSavePath = *savePath + "/x"
}

func main() {
	fmt.Printf("\nLoad source yaml file\n")
	url := strings.Replace(source, "<VERSION>", *version, -1)
	defer func() {
		fmt.Printf("\ndone for\n'%s'\n\n", url)
	}()
	ecsNested, err := yaml.Load(url)
	if err != nil {
		fmt.Printf("ERROR: %+v\n", err)
		return
	}

	fmt.Printf("\nLoad templtes\n")
	fmt.Printf("\nurl: %s\n", url)
	time.Sleep(1 * time.Second)

	template, err := FILES.ReadFile("templates/template.txt")
	if err != nil {
		fmt.Printf("ERROR: %+v\n", err)
		return
	}

	testTemplate, err := FILES.ReadFile("templates/syntax_test.txt")
	if err != nil {
		fmt.Printf("ERROR: %+v\n", err)
		return
	}

	checkTemplate, err := FILES.ReadFile("templates/check.txt")
	if err != nil {
		fmt.Printf("ERROR: %+v\n", err)
		return
	}

	readTemplate, err := FILES.ReadFile("templates/README.md")
	if err != nil {
		fmt.Printf("ERROR: %+v\n", err)
		return
	}

	fmt.Printf("\nRun set files...\n")
	setName := sortSetNames(ecsNested)
	for _, setName := range setName {
		fmt.Printf("\nstart '%s'\n", setName)
		if err := set.RunOneSet(setName, ecsNested[setName], pkgSavePath, string(template)); err != nil {
			fmt.Printf("ERROR: %+v\n", err)
			return
		}
	}

	fmt.Printf("\nRun test files...\n")
	if err := set.SaveTestFile(pkgImportPath, *savePath, string(testTemplate), setName); err != nil {
		fmt.Printf("ERROR: %+v\n", err)
		return
	}

	fmt.Printf("\nRun check files...\n")
	if err := set.SaveCheckFile(pkgImportPath, *savePath, string(checkTemplate), setName); err != nil {
		fmt.Printf("ERROR: %+v\n", err)
		return
	}

	fmt.Printf("\nRun README files...\n")
	if err := set.SaveReadmeFile(*savePath, string(readTemplate), *version); err != nil {
		fmt.Printf("ERROR: %+v\n", err)
		return
	}
}

func sortSetNames(set map[string]*face.Set) []string {
	out := make([]string, 0)
	for setName := range set {
		out = append(out, setName)
	}
	sort.Strings(out)
	return out
}

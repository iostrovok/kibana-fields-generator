package face

type Set struct {
	Beta        string            `yaml:"beta" json:"beta"`
	Name        string            `yaml:"name" json:"name"`
	Prefix      string            `yaml:"prefix" json:"prefix"`
	Title       string            `yaml:"title" json:"title"`
	Type        string            `yaml:"type" json:"type"`
	Group       int               `yaml:"group" json:"group"`
	Footnote    string            `yaml:"footnote" json:"footnote"`
	Description string            `yaml:"description" json:"description"`
	Short       string            `yaml:"short" json:"short"`
	Fields      map[string]*Field `yaml:"fields" json:"fields"`
	Nestings    []string          `yaml:"nestings" json:"nestings"`
	ReusedHere  []*ReusedHere     `yaml:"reused_here" json:"reused_here"`
	Root        bool              `yaml:"root" json:"root"`
	Reusable    *Reusable         `yaml:"reusable" json:"reusable"`
}

type Reusable struct {
	Expected []*Expected `yaml:"expected" json:"expected"`
	TopLevel bool        `yaml:"top_level" json:"top_level"`
}

type Expected struct {
	Full          string   `yaml:"full" json:"full"`
	As            string   `yaml:"as" json:"as"`
	At            string   `yaml:"at" json:"at"`
	Beta          string   `yaml:"beta" json:"beta"`
	ShortOverride string   `yaml:"short_override" json:"short_override"`
	Normalize     []string `yaml:"normalize" json:"normalize"`
}

type ReusedHere struct {
	Beta       string   `yaml:"beta" json:"beta"`
	Full       string   `yaml:"full" json:"full"`
	Normalize  []string `yaml:"normalize" json:"normalize"`
	SchemaName string   `yaml:"schema_name" json:"schema_name"`
	Short      string   `yaml:"short" json:"short"`
}

type MultiFields struct {
	FlatName string `yaml:"flat_name" json:"flat_name"`
	Name     string `yaml:"name" json:"name"`
	Type     string `yaml:"type" json:"type"`
}

type AllowedValue struct {
	Beta               string   `yaml:"beta" json:"beta"`
	Name               string   `yaml:"name" json:"name"`
	Description        string   `yaml:"description" json:"description"`
	ExpectedEventTypes []string `yaml:"expected_event_types" json:"expected_event_types"`
}

type Field struct {
	AllowedValues    []*AllowedValue `yaml:"allowed_values" json:"allowed_values"`
	Beta             string          `yaml:"beta" json:"beta"`
	DashedName       string          `yaml:"dashed_name" json:"dashed_name"`
	Description      string          `yaml:"description" json:"description"`
	DocValues        bool            `yaml:"doc_values" json:"doc_values"`
	Index            bool            `yaml:"index" json:"index"`
	Required         bool            `yaml:"required" json:"required"`
	InputFormat      string          `yaml:"input_format" json:"input_format"`
	ObjectType       string          `yaml:"object_type" json:"object_type"`
	OutputFormat     string          `yaml:"output_format" json:"output_format"`
	OutputPrecision  int             `yaml:"output_precision" json:"output_precision"`
	Example          string          `yaml:"example" json:"example"`
	ExpectedValues   []string        `yaml:"expected_values" json:"expected_values"`
	FlatName         string          `yaml:"flat_name" json:"flat_name"`
	Format           string          `yaml:"format" json:"format"`
	IgnoreAbove      int             `yaml:"ignore_above" json:"ignore_above"`
	Level            string          `yaml:"level" json:"level"`
	MultiFields      []*MultiFields  `yaml:"multi_fields" json:"multi_fields"`
	Name             string          `yaml:"name" json:"name"`
	Normalize        []string        `yaml:"normalize" json:"normalize"`
	ScalingFactor    int             `yaml:"scaling_factor" json:"scaling_factor"`
	Pattern          string          `yaml:"pattern" json:"pattern"`
	OriginalFieldset string          `yaml:"original_fieldset" json:"original_fieldset"`
	Short            string          `yaml:"short" json:"short"`
	Type             string          `yaml:"type" json:"type"`
}

package tsclient

type labelDefType int

const metricNameLabel = "__name__"

const (
	dynamicLabel labelDefType = iota
	commonLabel
)

type definitionLabel struct {
	name      string
	labelType labelDefType

	// only when labelType is commonLabel
	commonValue string

	// only when labelType is dynamicLabel
	i int
}

// DefinitionOpts can be changed with opt setters
type DefinitionOpts struct {
	// CommonLabels is a set of static key-value labels
	CommonLabels map[string]string

	// MeasuredLabelKeys is a list of label keys whose values will be specified
	// at run-time
	MeasuredLabelKeys []string
}

// Definition holds the description of a metric.
type Definition struct {
	name         string
	sortedLabels []definitionLabel
}

// DefinitionOpt is an option initializer for metric registration.
type DefinitionOpt func(*DefinitionOpts)

// WithCommonLabels includes common labels
func WithCommonLabels(labels map[string]string) DefinitionOpt {
	_ = "STUB: not implemented"
	return *new(DefinitionOpt)
}

// WithMeasuredLabels includes labels
func WithMeasuredLabels(labelKeys ...string) DefinitionOpt {
	_ = "STUB: not implemented"
	return *new(DefinitionOpt)
}

// NewDefinition returns a new definition
func NewDefinition(name string, opts ...DefinitionOpt) *Definition {
	_ = "STUB: not implemented"
	return nil
}

// NewDefinitionFromMap returns a new definition with common labels for each given value, a "__name__" key must also be present
func NewDefinitionFromMap(m map[string]string) *Definition { _ = "STUB: not implemented"; return nil }

// GetLFM returns an lfm corresponding to a definition
func GetLFM(def *Definition, labels []string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// ParseMetricDelimited parses a delimited message
func ParseMetricDelimited(s string) (map[string]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ConvertLFMMapToPrometheusEncodedName converts a metric in the form map[string]string{"__name__":"sonar_cpu","__source__":"user", "cpu":"cpu0"}
// to sonar_cpu{__source__="user",cpu="cpu0",host_id="899669",mode="iowait",user_id="208897"}
func ConvertLFMMapToPrometheusEncodedName(lfm map[string]string) string {
	_ = "STUB: not implemented"
	return ""
}

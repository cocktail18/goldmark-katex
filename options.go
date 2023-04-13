package katex

type Output string

const (
	OutputHtml          Output = "html"
	OutputMathml        Output = "mathml"
	OutputHtmlAndMathml Output = "htmlAndMathml"
)

type Options struct {
	DisplayMode      bool        `json:"displayMode"`
	Output           Output      `json:"output"`
	Leqno            bool        `json:"leqno"`
	Fleqn            bool        `json:"fleqn"`
	ThrowOnError     bool        `json:"throwOnError"`
	ErrorColor       string      `json:"errorColor"`
	MinRuleThickness float32     `json:"minRuleThickness"`
	ColorIsTextColor bool        `json:"colorIsTextColor"`
	MaxSize          int         `json:"maxSize"`
	MaxExpand        int         `json:"maxExpand"`
	Strict           interface{} `json:"strict"`
	Trust            interface{} `json:"trust"`
	GlobalGroup      bool        `json:"globalGroup"`
}

package templatefuncs

import "text/template"

// TemplateFuncProvider defines the interface for external function providers.
type TemplateFuncProvider interface {
	GetFunctions() template.FuncMap
}

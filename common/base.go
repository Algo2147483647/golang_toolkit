package common

import (
	"bytes"
	"encoding/json"
	"fmt"
	"html/template"
	"strings"
)

// Contains checks if an element is in a slice, works with any type T
// Returns true if the element is found, false otherwise.
func Contains[T comparable](slice []T, element T) bool {
	for _, v := range slice {
		if v == element {
			return true
		}
	}
	return false
}

// ReplaceTemplateWithMap replaces placeholders in a template string with values from a map.
func ReplaceTemplateWithMap(templateStr string, data map[string]interface{}) (string, error) {
	// Parse the template with our custom functions
	tmpl, err := template.New("base").Funcs(GetFuncMap()).Parse(templateStr)
	if err != nil {
		return "", fmt.Errorf("failed to parse template: %v", err)
	}

	// Execute the template
	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, data); err != nil {
		return "", fmt.Errorf("failed to execute template: %v", err)
	}

	return buf.String(), nil
}

// GetFuncMap returns a map of functions that can be used in templates
func GetFuncMap() template.FuncMap {
	return template.FuncMap{
		"json": func(v interface{}) (string, error) {
			b, err := json.Marshal(v)
			if err != nil {
				return "", err
			}
			return string(b), nil
		},
		"jsonPretty": func(v interface{}) (string, error) {
			b, err := json.MarshalIndent(v, "", "  ")
			if err != nil {
				return "", err
			}
			return string(b), nil
		},
		"toString": func(v interface{}) string {
			return fmt.Sprintf("%v", v)
		},
		"toLower":   strings.ToLower,
		"toUpper":   strings.ToUpper,
		"trim":      strings.TrimSpace,
		"contains":  strings.Contains,
		"hasPrefix": strings.HasPrefix,
		"hasSuffix": strings.HasSuffix,
		"split":     strings.Split,
		"join":      strings.Join,
	}
}

package common

import (
	"context"
	"encoding/json"
	"fmt"
	"html/template"
	"strings"
	"time"
)

func TemplateParse(ctx context.Context, templateStr string, params map[string]interface{}, funcMap template.FuncMap) (string, error) {
	if params == nil {
		params = make(map[string]interface{})
	}

	tmpl, err := template.New("template").Funcs(funcMap).Parse(templateStr)
	if err != nil {
		return "", fmt.Errorf("failed to parse template: %w", err)
	}

	var buf strings.Builder
	if err := tmpl.Execute(&buf, params); err != nil {
		return "", fmt.Errorf("failed to execute template: %w", err)
	}

	return buf.String(), nil
}

func GetBasicFuncMap() template.FuncMap {
	return template.FuncMap{
		"now":       time.Now,
		"toLower":   strings.ToLower,
		"toUpper":   strings.ToUpper,
		"trim":      strings.TrimSpace,
		"contains":  strings.Contains,
		"hasPrefix": strings.HasPrefix,
		"hasSuffix": strings.HasSuffix,
		"split":     strings.Split,
		"join":      strings.Join,

		"formatTime": func(t time.Time, layout string) string {
			return t.Format(layout)
		},

		"add": func(a, b interface{}) (interface{}, error) {
			switch a := a.(type) {
			case int:
				switch b := b.(type) {
				case int:
					return a + b, nil
				case float64:
					return float64(a) + b, nil
				}
			case float64:
				switch b := b.(type) {
				case int:
					return a + float64(b), nil
				case float64:
					return a + b, nil
				}
			}
			return nil, fmt.Errorf("unsupported types for addition")
		},

		"hasKey": func(m map[string]interface{}, key string) bool {
			_, ok := m[key]
			return ok
		},

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

		"jsonParse": func(s string) (map[string]interface{}, error) {
			var result map[string]interface{}
			if err := json.Unmarshal([]byte(s), &result); err != nil {
				return nil, err
			}
			return result, nil
		},

		"jsonParseArray": func(s string) ([]interface{}, error) {
			var result []interface{}
			if err := json.Unmarshal([]byte(s), &result); err != nil {
				return nil, err
			}
			return result, nil
		},

		"jsonValid": func(v interface{}) bool {
			_, err := json.Marshal(v)
			return err == nil
		},
	}
}

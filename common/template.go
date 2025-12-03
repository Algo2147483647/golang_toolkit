package common

import (
	"context"
	"encoding/json"
	"fmt"
	"html/template"
	"strings"
	textTemplate "text/template"
	"time"
)

func TemplateParse(ctx context.Context, template string, params map[string]interface{}, funcMap textTemplate.FuncMap) (string, error) {
	if params == nil {
		params = make(map[string]interface{})
	}

	tmpl, err := textTemplate.New("template").Funcs(funcMap).Parse(template)
	if err != nil {
		return "", fmt.Errorf("failed to parse template: %w", err)
	}

	var buf strings.Builder
	if err := tmpl.Execute(&buf, params); err != nil {
		return "", fmt.Errorf("failed to execute template: %w", err)
	}

	return buf.String(), nil
}

func GetBasicTemplateFuncMap() template.FuncMap {
	return textTemplate.FuncMap{
		"now":     time.Now,
		"toUpper": strings.ToUpper,
		"toLower": strings.ToLower,
		"trim":    strings.TrimSpace,
		"join":    strings.Join,
		"split":   strings.Split,

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

		// 将任意值转换为格式化的 JSON 字符串（带缩进）
		"jsonPretty": func(v interface{}) (string, error) {
			b, err := json.MarshalIndent(v, "", "  ")
			if err != nil {
				return "", err
			}
			return string(b), nil
		},

		// 将 JSON 字符串解析为 map[string]interface{}
		"jsonParse": func(s string) (map[string]interface{}, error) {
			var result map[string]interface{}
			if err := json.Unmarshal([]byte(s), &result); err != nil {
				return nil, err
			}
			return result, nil
		},

		// 将 JSON 字符串解析为 slice
		"jsonParseArray": func(s string) ([]interface{}, error) {
			var result []interface{}
			if err := json.Unmarshal([]byte(s), &result); err != nil {
				return nil, err
			}
			return result, nil
		},

		// 检查一个值是否可以转换为有效的 JSON
		"jsonValid": func(v interface{}) bool {
			_, err := json.Marshal(v)
			return err == nil
		},
	}
}

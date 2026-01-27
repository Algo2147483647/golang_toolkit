package rule_engine

type Environment struct {
	Variables map[string]ValueIf
	Functions map[string]func(...ValueIf) (ValueIf, error)
}

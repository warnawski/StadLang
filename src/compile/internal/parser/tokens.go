package parser

var (
	KeywordsMap = map[string]bool{
		"constant":   true,
		"func":       true,
		"typeobject": true,
	}
	ScopeTypeMap = map[string]bool{
		"public":  true,
		"private": true,
	}

	TypeMap = map[string]bool{
		"bool":   true,
		"int":    true,
		"float":  true,
		"string": true,
		"void":   true,
	}
	DelimiterMap = map[string]bool{
		"{":  true,
		"}":  true,
		"[":  true,
		"]":  true,
		"->": true,
		":=": true,
		":":  true,
	}
	CommentMap = map[string]bool{
		"//": true,
	}
)

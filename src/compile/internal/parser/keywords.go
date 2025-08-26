package parser

var keywords = map[string]bool{
	"private":    true,
	"public":     true,
	"constant":   true,
	"func":       true,
	"return":     true,
	"for":        true,
	"if":         true,
	"else":       true,
	"elif":       true,
	"switch":     true,
	"case":       true,
	"default":    true,
	"typeobject": true,
}

var operators = map[string]bool{
	"+":  true,
	"-":  true,
	"*":  true,
	"/":  true,
	"%":  true,
	"=":  true,
	">":  true,
	">=": true,
	"<":  true,
	"<=": true,
	":=": true,
	"==": true,
	"|":  true,
	"&":  true,
	"&&": true,
	"||": true,
	"->": true,
	":":  true,
}

var types = map[string]bool{
	"bool":   true,
	"int":    true,
	"string": true,
	"float":  true,
	"null":   true,
	"void":   true,
	"map":    true,
	"array":  true,
}

var delimiters = map[rune]bool{
	'(': true,
	')': true,
	':': true,
	';': true,
	',': true,
}

var braces = map[rune]bool{
	'{': true,
	'}': true,
}

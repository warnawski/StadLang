package parser

import "fmt"

func newToken(v string, line int, col int) Token {

	var ttype int

	switch {
	case keywords[v]:
		ttype = T_KEYWORD
	case types[v]:
		ttype = T_TYPE
	case operators[v]:
		ttype = T_OPERATOR
	case delimiters[rune(v[0])]:
		ttype = T_DELIMITER
	case len(v) == 1 && braces[rune(v[0])]:
		if v == "{" {
			ttype = T_LBRACE
		} else if v == "}" {
			ttype = T_RBRACE
		}
	default:
		ttype = T_IDENNTIFIER

	}

	return Token{
		Type:   ttype,
		Values: v,
		Line:   line,
		Column: col,
	}
}

func Tokenizator(inputCode string) []Token {

	r := []rune(inputCode)

	var buff []rune
	var tokens []Token

	var line = 1
	var col = 1

	var op2 = map[rune]bool{
		'=': true,
		'&': true,
		'>': true,
		'|': true,
	}

	for s := 0; s < len(r); s++ {
		fmt.Println(string(r[s]))

		switch r[s] {

		case '/':
			if peak(s, r) {
				if r[s+1] == '/' {

					s++

					for s < len(r) && r[s] != '\n' {
						s++
					}
					line++
					col = 1

				} else if r[s+1] == '*' {

					s++

					for s < len(r) && r[s] != '*' && r[s+1] != '/' {

						if r[s] == '\n' {
							line++
							col = 1
						}
						s++
					}

				} else {

					tokens = append(tokens, newToken(string('/'), line, col))
					col++

				}
			}

		case '\n', ' ':
			if r[s] == '\n' {

				flushBuffer(&buff, &tokens, col, line)

				line++
				col = 1
			} else {

				flushBuffer(&buff, &tokens, col, line)

				col++
			}

		case '+', '-', '*', '=', '|', '&', ':':

			if peak(s, r) {

				if op2[r[s+1]] {
					flushBuffer(&buff, &tokens, col, line)
					tokens = append(tokens, newToken(string([]rune{r[s], r[s+1]}), line, col))
					buff = buff[:0]

					s += 2
					col += 2
				} else {
					flushBuffer(&buff, &tokens, col, line)
					tokens = append(tokens, newToken(string(r[s]), line, col))
					buff = buff[:0]

					s++
					col++

				}
			}

		case '(', ')', ',', ';':

			flushBuffer(&buff, &tokens, col, line)
			col++

		case '{', '}':
			flushBuffer(&buff, &tokens, col, line)
			tokens = append(tokens, newToken(string(r[s]), line, col))
			col++

		default:
			buff = append(buff, r[s])
			tokens = append(tokens, newToken(string(r[s]), line, col))
			col++
		}

	}

	return tokens
}

/* //Function for checked next symbol on EOF */
func peak(x int, slice []rune) bool {

	if x+1 < len(slice) {
		return true
	}

	return false
}

func flushBuffer(buffer *[]rune, list *[]Token, column int, line int) {
	if len(*buffer) == 0 {
		return
	}
	*list = append(*list, newToken(string(*buffer), column, line))
	*buffer = (*buffer)[:0]
}

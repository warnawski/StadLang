package parser

const (
	Keywords   = "Keywords"
	ScopeType  = "ScopeType"
	Type       = "Type"
	Identifier = "Identifier"
	Operator   = "Operator"

	Int    = "Int"
	String = "String"

	Delimiter = "Delimiter"

	Comment = "Comment"
)

type Token struct {
	Type  string
	Value string
	Line  int
	Col   int
}

func newToken(v string, c int, l int) Token {

	t := typeChecker(v)

	return Token{
		Type:  t,
		Value: v,
		Line:  l,
		Col:   c,
	}
}

func typeChecker(v string) string {
	return ""
}

func tokenizationV1(code string) []Token {

	mutstr := []rune(code)

	var buff []rune
	var column int = 1
	var line int = 1

	for x, _ := range mutstr {

		if x == '\n' {
			line++
			column = 1
		} else {
			column++
		}

		if mutstr[x] != ' ' {
			buff = append(buff, mutstr[x])
		} else {
			tkn := make([]rune, len(buff))
			copy(tkn, buff)

			token := string(tkn)
			newToken(token, column, line)
			buff = buff[:0]
		}

	}

	return nil
}

func TokenizationV2(code string) []Token {

	mutstr := []rune(code)
	var buff []rune
	var column int = 1
	var line int = 1
	var tokenList []Token

	//Запускаем цикл по всему коду
	for x := 0; x < len(mutstr); x++ {

		//Условие для проверки на однострочные комментарии
		if x+1 < len(mutstr) && mutstr[x] == '/' && mutstr[x+1] == '/' {

			x += 2
			//Запускаем цикл, пока не закончится комментарий
			for x < len(mutstr) && mutstr[x] != '\n' {
				x++
			}
			line++
			column = 1
		}

		if mutstr[x] == '\n' {
			line++
			column = 1
		} else {
			column++
		}

		if mutstr[x] != ' ' {
			buff = append(buff, mutstr[x])
		} else {
			tkn := newToken(string(buff), column, line)
			tokenList = append(tokenList, tkn)
			buff = buff[:0]
		}

		switch mutstr[x] {

		case '+':
		}

	}
	return tokenList
}

func TokenizationV3(code string) []Token {

	mutstr := []rune(code)
	var buff []rune
	var column = 1
	var line = 1
	var tokenList []Token

	for x := 0; x < len(mutstr); x++ {

		//Запускаем проверку символов
		switch mutstr[x] {

		//Проверка на структуры комментариев
		case '/':

			//Защита от случаев, если всего 1 / и конец кода
			if x+1 < len(mutstr) {

				//Проверка на однострочный комментарий
				if mutstr[x+1] == '/' {
					x += 2

					for x < len(mutstr) && mutstr[x] != '\n' {
						x++
					}
					line++
					column = 1
					continue

					//Проверка на двустрочный комментарий
				} else if mutstr[x+1] == '*' {
					x += 2

					for x < len(mutstr) && mutstr[x] != '*' && mutstr[x+1] != '/' {
						if mutstr[x] == '\n' {
							line++
							column = 1
						}
						x++
					}
					x++
					continue
				}
			}
			tokenList = append(tokenList, newToken(string(mutstr[x]), column, line))
			column++

		//Проверка на разделители токенов и операторы
		case ' ', '\t', '\n', '=', '(', ')', '{', '}', '+', '-', ':':

			//Список двусимвольных операторов
			var op_2 = []string{
				":=", "+=", "-=", "*=", "==", "!=", ">=", "<=", "->", "&&", "||", "**",
			}
			var op_1 = []string{
				"+", "-", "*", "|", "%", "=", "<", ">", "!", "&", "|",
			}
			//При слитном написании
			if len(buff) > 0 {
				tokenList = append(tokenList, newToken(string(buff), column, line))
				buff = buff[:0]
			}
			//Проверка двусимвольных операторов
			for o := 0; o < len(op_2); o++ {

				if x+1 < len(mutstr) && op_2[o] == string([]rune{mutstr[x], mutstr[x+1]}) {
					tokenList = append(tokenList, newToken(string([]rune{mutstr[x], mutstr[x+1]}), column, line))
					x++
					column += 2
					buff = buff[:0]
					break
				}
			}

			for o := 0; o < len(op_1); o++ {

				if op_1[o] == string(mutstr[x]) {
					tokenList = append(tokenList, newToken(string(mutstr[x]), column, line))
					column++
					buff = buff[:0]
					break
				}
			}

			/*
				if len(buff) > 0 {

					for _, o := range op {

						if o == mutstr[x]+mutstr[x+1] {

							tokenList = append(tokenList, newToken(string(buff), column, line))
							buff = buff[:0]

							operator := mutstr[x]+mutstr[x+1]

							buff = append(buff, operator)
							tokenList = append(tokenList, newToken(string(buff), column, line))
							buff = buff[:0]
						}
					}

					} else {

				}
			*/

			if mutstr[x] == '\n' {
				line++
				column = 1
			} else {
				column++
			}

		default:
			buff = append(buff, mutstr[x])
			column++

		}
	}
	return tokenList
}

func TokenizationV4(code string) []Token {

	return nil
}

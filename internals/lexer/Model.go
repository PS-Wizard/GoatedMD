package lexer

type Token struct {
	Type    string
	Level   int
	Text    string
	Context any // Idk just incase i might need something?
}

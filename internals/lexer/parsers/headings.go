package parsers

import (
	"strings"

	"github.com/PS-Wizard/GoatedMD/internals/lexer"
)

func ParseHeadings(line string) *lexer.Token {

	line = strings.TrimSpace(line)
	if strings.HasPrefix(line, "#") {
		level := 0
		for i := 0; i <= len(line) && line[i] == '#'; i++ {
			level++
		}

		if len(line) > level && line[level] == ' ' {
			content := strings.TrimSpace(line[level:])

			return &lexer.Token{
				Type:    "Heading",
				Level:   level,
				Text:    content,
				Context: nil,
			}

		}
	}
	return nil
}

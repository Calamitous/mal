package reader

import (
	"fmt"
	"regexp"

	"types"
)

func Read_str(input string) string {
	fmt.Println(tokenizer(input))
	return "ASdf"
}

type TokenReader struct {
	tokens   []string
	position int
}

func (r TokenReader) peek() string {
	return r.tokens[r.position]
}

func (r TokenReader) next() string {
	char := r.tokens[r.position]
	r.position += 1
	return char
}

func tokenizer(input string) []string {
	var token_regex = regexp.MustCompile(`[\s,]*(~@|[\[\]{}()'` + "`" + `~^@]|"(?:\\.|[^\\"])*"|;.*|[^\s\[\]{}('"` + "`" + `,;)]*)`)
	var results []string

	for _, group := range token_regex.FindAllStringSubmatch(input, -1) {
		fmt.Println(group)
		if (group[1] == "") || (group[1][0] == ';') {
			continue
		}
		results = append(results, group[1])
	}

	return results
}

func read_form(tr TokenReader) {
	if tr.peek() == "(" {
		read_list(tr)
	} else {
		read_atom(tr)
	}
}

func read_list(tr TokenReader) types.List {
	if tr.peek() != ")" {
		read_form(tr)
	}
	var t []types.MalType
	for i := range tr.tokens {
		t = append(t, types.MalType(tr.tokens[i]))
	}
	return types.List{
		Val: t,
	}
}

func read_atom(tr TokenReader) {
	print("hi")
	tr.tokens = append(tr.tokens, tr.next())
}

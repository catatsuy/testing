package mt

import (
	"bufio"
	"io"
	"strings"
)

type MT struct {
	Author   string
	Title    string
	Basename string
}

func (m *MT) Parse(r io.Reader) {
	scanner := bufio.NewScanner(r)

	for scanner.Scan() {
		ss := strings.Split(scanner.Text(), ": ")
		if len(ss) <= 1 {
			continue
		}
		key, value := ss[0], ss[1]

		switch key {
		case "AUTHOR":
			m.Author = value
			break
		case "TITLE":
			m.Title = value
			break
		case "BASENAME":
			m.Basename = value
			break
		}
	}
}

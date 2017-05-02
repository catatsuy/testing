package mt

import (
	"bufio"
	"io"
	"strings"
)

type MT struct {
	Author string
}

func (m *MT) Parse(r io.Reader) {
	scanner := bufio.NewScanner(r)

	for scanner.Scan() {
		ss := strings.Split(scanner.Text(), ": ")
		if len(ss) <= 1 {
			continue
		}
		if ss[0] == "AUTHOR" {
			m.Author = ss[1]
		}
	}
}

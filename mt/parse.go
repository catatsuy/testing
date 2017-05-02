package mt

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"

	"github.com/pkg/errors"
)

type MT struct {
	Author        string
	Title         string
	Basename      string
	Status        string
	AllowComments int
}

func Parse(r io.Reader) (*MT, error) {
	m := &MT{}

	scanner := bufio.NewScanner(r)

	var err error

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
		case "STATUS":
			if value == "Draft" || value == "Publish" || value == "Future" {
				m.Status = value
			} else {
				return nil, fmt.Errorf("STATUS column is allowed only Draft or Publish or Future. Got %s", value)
			}
			break
		case "ALLOW COMMENTS":
			m.AllowComments, err = strconv.Atoi(value)
			if err != nil {
				return nil, errors.Wrap(err, "ALLOW COMMENTS column is allowed only 0 or 1")
			}
			if m.AllowComments != 0 && m.AllowComments != 1 {
				return nil, fmt.Errorf("ALLOW COMMENTS column is allowed only 0 or 1. Got %d", m.AllowComments)
			}
			break
		}
	}

	return m, nil
}

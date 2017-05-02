package mt

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"

	"github.com/pkg/errors"
)

// Default
const (
	// If it is not inialized, AllowComments is -1
	DefaultAllowComments = -1
)

// Movable Type Import Format
type MT struct {
	Author   string
	Title    string
	Basename string
	Status   string

	// 0 or 1. If it is not inialized DefaultAllowComments.
	AllowComments int
}

// NewMT creates MT.
func NewMT() *MT {
	return &MT{
		AllowComments: DefaultAllowComments,
	}
}

// Parse creates MT struct from io.Reader
func Parse(r io.Reader) (*MT, error) {
	m := NewMT()

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

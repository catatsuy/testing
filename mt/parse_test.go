package mt_test

import (
	"bytes"
	"testing"

	. "github.com/catatsuy/testing/mt"
)

func TestParse(t *testing.T) {
	buf := bytes.NewBufferString(`AUTHOR: catatsuy
TITLE: ポエム
BASENAME: poem
STATUS: Publish
ALLOW COMMENTS: 1
CONVERT BREAKS: 0
DATE: 04/22/2017 20:41:58
CATEGORY: ポエム
CATEGORY: 技術系
-----
BODY:
<p>bodybody</p>
-----
EXTENDED BODY:
<p>extended body body body</p>
-----
`)

	expected := &MT{
		Author: "catatsuy",
		Title:  "ポエム",
	}

	m := &MT{}
	m.Parse(buf)

	if m.Author != expected.Author {
		t.Errorf("expected author column %s; got %s", expected.Author, m.Author)
	}
	if m.Title != expected.Title {
		t.Errorf("expected title column %s; got %s", expected.Title, m.Title)
	}
}

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

	m := &MT{}
	m.Parse(buf)

	if m.Author != "catatsuy" {
		t.Errorf("Cannot parse author column: %s", m.Author)
	}
}

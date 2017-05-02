package mt_test

import (
	"bytes"
	"reflect"
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
		Author:   "catatsuy",
		Title:    "ポエム",
		Basename: "poem",
	}

	m := &MT{}
	m.Parse(buf)

	if !reflect.DeepEqual(m, expected) {
		t.Errorf("Error parsing, expected %q; got %q", expected, m)
	}
}

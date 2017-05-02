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
		Author:        "catatsuy",
		Title:         "ポエム",
		Basename:      "poem",
		Status:        "Publish",
		AllowComments: 1,
	}

	m, err := Parse(buf)

	if err != nil {
		t.Fatalf("got error %q", err)
	}

	if !reflect.DeepEqual(m, expected) {
		t.Errorf("Error parsing, expected %q; got %q", expected, m)
	}
}

func TestParseStatusNotAllowed(t *testing.T) {
	buf := bytes.NewBufferString(`STATUS: Published`)

	_, err := Parse(buf)

	if err == nil || err.Error() != "STATUS column is allowed only Draft or Publish or Future. Got Published" {
		t.Errorf("Error parsing, got %q", err)
	}
}

func TestNewMT(t *testing.T) {
	m := NewMT()

	if m.AllowComments != DefaultAllowComments {
		t.Errorf("By default, AllowComments is -1")
	}
}

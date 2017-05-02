package mt_test

import (
	"bytes"
	"io"
	"reflect"
	"testing"
	"time"

	. "github.com/catatsuy/testing/mt"
)

func TestParse(t *testing.T) {
	buf := bytes.NewBufferString(`AUTHOR: catatsuy
TITLE: ポエム
BASENAME: poem
STATUS: Publish
ALLOW COMMENTS: 1
ALLOW PINGS: 1
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
		AllowPings:    1,
		ConvertBreaks: "0",
		Date:          time.Date(2017, time.April, 22, 20, 41, 58, 0, time.UTC),
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

func TestParseDate(t *testing.T) {
	var featuretests = []struct {
		buf io.Reader
		t   time.Time
	}{
		{
			bytes.NewBufferString(`DATE: 04/22/2017 08:41:58 PM`),
			time.Date(2017, time.April, 22, 20, 41, 58, 0, time.UTC),
		},
		{
			bytes.NewBufferString(`DATE: 04/22/2017 08:41:58 AM`),
			time.Date(2017, time.April, 22, 8, 41, 58, 0, time.UTC),
		},
		{
			bytes.NewBufferString(`DATE: 04/22/2017 20:41:58`),
			time.Date(2017, time.April, 22, 20, 41, 58, 0, time.UTC),
		},
	}

	for _, ft := range featuretests {
		m, err := Parse(ft.buf)

		if err != nil {
			t.Fatalf("got error %q", err)
		}

		if m.Date != ft.t {
			t.Errorf("m.Date got %v; want %v", m.Date)
		}
	}
}

func TestNewMT(t *testing.T) {
	m := NewMT()

	if m.AllowComments != DefaultAllowComments {
		t.Errorf("By default, AllowComments is %d, got %d", DefaultAllowComments, m.AllowComments)
	}

	if m.AllowPings != DefaultAllowPings {
		t.Errorf("By default, AllowComments is %d, got %d", DefaultAllowPings, m.AllowPings)
	}
}

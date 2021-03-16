package bridge

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMetadata(t *testing.T) {

	meta := []struct {
		Key string
		Val string
	}{
		{
			Key: "SERVICE_TAGS",
			Val: "braid,base",
		},
		{
			Key: "SERVICE_14201_NAME",
			Val: "base",
		},
		{
			Key: "SERVICE_14202_CHECK_HTTP",
			Val: "/health",
		},
	}

	metalst := []string{}
	for _, v := range meta {
		metalst = append(metalst, v.Key+"="+v.Val)
	}
	metadataImpl(metalst, "14201")
}

func TestEscapedComma(t *testing.T) {
	cases := []struct {
		Tag      string
		Expected []string
	}{
		{
			Tag:      "",
			Expected: []string{},
		},
		{
			Tag:      "foobar",
			Expected: []string{"foobar"},
		},
		{
			Tag:      "foo,bar",
			Expected: []string{"foo", "bar"},
		},
		{
			Tag:      "foo\\,bar",
			Expected: []string{"foo,bar"},
		},
		{
			Tag:      "foo,bar\\,baz",
			Expected: []string{"foo", "bar,baz"},
		},
		{
			Tag:      "\\,foobar\\,",
			Expected: []string{",foobar,"},
		},
		{
			Tag:      ",,,,foo,,,bar,,,",
			Expected: []string{"foo", "bar"},
		},
		{
			Tag:      ",,,,",
			Expected: []string{},
		},
		{
			Tag:      ",,\\,,",
			Expected: []string{","},
		},
	}

	for _, c := range cases {
		results := recParseEscapedComma(c.Tag)
		sort.Strings(c.Expected)
		sort.Strings(results)
		assert.EqualValues(t, c.Expected, results)
	}
}

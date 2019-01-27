package lab

import (
	"reflect"
	"testing"
)

func TestParseContent(t *testing.T) {
	cases := []struct {
		Content  string
		Expected CreateMergeRequestOptions
	}{
		{"Foo\nBar", CreateMergeRequestOptions{Title: "Foo", Description: "Bar"}},
		{"Foo\n\n\nBar\n", CreateMergeRequestOptions{Title: "Foo", Description: "Bar"}},
	}

	for _, tc := range cases {
		opts := CreateMergeRequestOptions{}
		opts.ParseContent(tc.Content)

		if !reflect.DeepEqual(opts, tc.Expected) {
			t.Errorf("ParseContent(%s) = %+v, expected %+vs", tc.Content, opts, tc.Expected)
		}
	}
}

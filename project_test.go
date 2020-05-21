package lab

import (
	"net/url"
	"reflect"
	"testing"
)

func TestProjectURL(t *testing.T) {
	cases := []struct {
		URL      string
		Expected RemoteProject
	}{
		{"https://gitlab.com/joshbohde/lab", RemoteProject{Name: "origin", Host: "gitlab.com", Path: "joshbohde/lab"}},
		{"https://gitlab.com/joshbohde/lab.git", RemoteProject{Name: "origin", Host: "gitlab.com", Path: "joshbohde/lab"}},
	}

	for _, tc := range cases {
		u, err := url.ParseRequestURI(tc.URL)
		if err != nil {
			t.Errorf("Error parsing %s: %s = ", tc.URL, err)
		}

		actual := ParseRemoteProject("origin", *u)

		if !reflect.DeepEqual(actual, tc.Expected) {
			t.Errorf("ProjectUrl(%s) = %+v, expected %+vs", tc.URL, actual, tc.Expected)
		}
	}
}

package git

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGit(t *testing.T) {
	git := New()

	t.Run("RevList", func(t *testing.T) {
		cases := []struct {
			Base     string
			Merge    string
			Expected []string
		}{
			{"4d8e32b7aaaa972591f1aedada9fa7f80f56b675", "4d8e32b7aaaa972591f1aedada9fa7f80f56b675", []string{}},
			{"4d8e32b7aaaa972591f1aedada9fa7f80f56b675", "21180cedbed57968c7a510cc8e35af5546b309e8", []string{"21180cedbed57968c7a510cc8e35af5546b309e8"}},
			{"4d8e32b7aaaa972591f1aedada9fa7f80f56b675", "8f1bfab449ea5cc3e453827bae811c2146ce10c4", []string{"8f1bfab449ea5cc3e453827bae811c2146ce10c4", "2327b71470ea1ff1bb4a7e124c231db54ee1f5c0", "eab9a19249659467dd2fcec34587f17292e4a250", "98ba2955afff8c8e534d509631eb31116a96dba5", "37d36f74ee151627206655560cd327cf69c2aff0", "757b875b21dfaa5a57bebfb724973f0c41d9311f", "1d984a145d513085aafd685ed9ad73cc3a42fa6d", "1a0607f0d56a27bcec306f3dc0000d4612512109", "76b4ec8efe59b4a86aa581db82e3a53c24db0d48", "9476697ba32895fe18104a396526ec49d68cf090", "12a2a64dfa8aed200e9ebe7b2e59df57832bc895", "b1544579e489b785ab1dd61c2cc67c778294eb1e", "1ab70f2dae45ae5b6f23d87af28dc3f59b7aac05", "21180cedbed57968c7a510cc8e35af5546b309e8"}},
		}

		for _, tc := range cases {
			actual, err := git.RevList(tc.Base, tc.Merge)

			assert.NoError(t, err)
			assert.Equal(t, tc.Expected, actual)
		}

	})

	t.Run("CommitMessage", func(t *testing.T) {
		t.Run("good", func(t *testing.T) {
			cases := []struct {
				Ref      string
				Expected string
			}{
				{"4d8e32b7aaaa972591f1aedada9fa7f80f56b675", "Initial commit.\n\n"},
				{"21180cedbed57968c7a510cc8e35af5546b309e8", "Initial code.\n\n"},
			}

			for _, tc := range cases {
				actual, err := git.CommitMessage(tc.Ref)

				assert.NoError(t, err)
				assert.Equal(t, tc.Expected, actual)
			}
		})
		t.Run("missing", func(t *testing.T) {
			actual, err := git.CommitMessage("alskdjflaksjef")

			assert.Error(t, err, "Received unexpected error:\nexit status 128")
			assert.Empty(t, actual)
		})

	})
}

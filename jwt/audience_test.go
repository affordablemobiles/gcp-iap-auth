package jwt

import (
	"strings"
	"testing"
)

func TestAudiences(t *testing.T) {
	testTable := []struct {
		name string
		aud  string
		err  string
	}{
		{
			name: "misc: not enough slashes",
			aud:  "/projects/1234",
			err:  "must follow the format",
		},
		{
			name: "app engine: valid",
			aud:  "/projects/1234/apps/fake-project-id",
			err:  "",
		},
		{
			name: "app engine: missing service details",
			aud:  "/projects/1234/",
			err:  "is missing service details",
		},
		{
			name: "global: valid",
			aud:  "/projects/1234/global/backendServices/1234",
			err:  "",
		},
		{
			name: "global: missing service details",
			aud:  "/projects/1234/",
			err:  "is missing service details",
		},
	}

	for _, tc := range testTable {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			_, err := ParseAudience(tc.aud)

			switch {
			case err == nil && tc.err == "":
				// noop
			case err != nil && tc.err == "":
				t.Error("expected no error, got error:", err)
			case err == nil && tc.err != "":
				t.Error("expected error, got no error:", tc.err)
			case err != nil && tc.err != "":
				if !strings.Contains(err.Error(), tc.err) {
					t.Error("unexpected error:", err)
				}
			}

		})
	}
}

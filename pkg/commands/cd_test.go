package commands

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCalculateDir(t *testing.T) {
	testcases := []struct {
		cwd       string
		dir       string
		expected  string
		errstring string
	}{
		{
			cwd:       "",
			dir:       "testme",
			expected:  "testme",
			errstring: "",
		},
		{
			cwd:       "testme",
			dir:       "hello",
			expected:  "testme/hello",
			errstring: "",
		},
		{
			cwd:       "testme",
			dir:       "../",
			expected:  "",
			errstring: "",
		},
		{
			cwd:       "testme/nested-1",
			dir:       "../",
			expected:  "testme",
			errstring: "",
		},
		{
			cwd:       "testme/nested-1",
			dir:       "../../",
			expected:  "",
			errstring: "",
		},
		{
			cwd:       "testme/nested-1",
			dir:       "/some-root-dir",
			expected:  "/some-root-dir",
			errstring: "",
		},
		{
			cwd:       "testme",
			dir:       "../../",
			expected:  "",
			errstring: "cd: no such file or directory: ../../",
		},
	}

	for _, tc := range testcases {
		t.Run(fmt.Sprintf("adding %s to cwd", tc.dir), func(t *testing.T) {
			fmt.Println(tc.cwd, "->", tc.dir)
			ctx := NewContext(context.Background())
			MustFromContext(ctx).cwd = tc.cwd

			output, err := calculateDir(ctx, tc.dir)
			fmt.Println("error", err)
			if tc.errstring == "" {
				require.Nil(t, err)
			} else {
				require.Equal(t, err.Error(), tc.errstring)
			}

			require.Equal(t, output, tc.expected)
		})
	}
}

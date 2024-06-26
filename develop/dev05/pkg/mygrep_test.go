package mygrep

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGrep(t *testing.T) {

	testCases := []struct {
		arg []string
	}{
		{
			arg: []string{"-n", "Games", "input"},
		},
		{
			arg: []string{"-A", "1", "Games", "input"},
		},
		{
			arg: []string{"-A", "4", "Games", "input"},
		},
		{
			arg: []string{"-B", "3", "Games", "input"},
		},
		{
			arg: []string{"-C", "5", "Games", "input"},
		},
		{
			arg: []string{"-C", "5", "-v", "Games", "input"},
		},
		{
			arg: []string{"-C", "5", "-i", "gAmes", "input"},
		},
		{
			arg: []string{"-C", "5", "-i", "-v", "gAmes", "input"},
		},
		{
			arg: []string{"-C", "5", "-i", "-v", "-c", "gAmes", "input"},
		},
		{
			arg: []string{"-C", "5", "-i", "-v", "-c", "1", "input"},
		},
	}

	for _, v := range testCases {
		t.Run("mygrep test", func(t *testing.T) {
			var cfg GrepCfg

			require.NoError(t, cfg.ParseConfig(v.arg))

			require.NoError(t, cfg.Run())
		})
	}
}

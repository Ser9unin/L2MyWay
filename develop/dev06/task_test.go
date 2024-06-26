package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCut(t *testing.T) {

	testCases := []struct {
		arg []string
	}{
		{
			arg: []string{"-d", "|"},
		},
		{
			arg: []string{"-d", "|", "-f", "2-4"},
		},
		{
			arg: []string{"-d", "512"},
		},
	}

	for _, v := range testCases {
		t.Run("mycut test", func(t *testing.T) {
			var cfg Config

			require.NoError(t, cfg.ParseConfig(v.arg))
			fmt.Println(cfg.d, cfg.f, cfg.s)

			data, err := os.Open("input")
			if err != nil {
				log.Fatal(err)
			}

			cfg.reader = data

			if err != nil && err != io.EOF {
				log.Fatal(err)
			}

			require.NoError(t, cfg.Run())
		})
	}
}

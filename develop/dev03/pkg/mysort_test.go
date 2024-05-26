package mysort

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSort(t *testing.T) {
	testCases := []struct {
		caseName string
		cfg      SortCfg
		data     []string
		want     []string
	}{
		{
			caseName: "normal",
			cfg:      SortCfg{},
			data: []string{
				"Standing on one's head at job interviews forms a lasting impression.",
				"The chic gangster liked to start the day with a pink scarf.",
				"He kept telling himself that one day it would all somehow make sense.",
				"Cats are good pets, for they are clean and are not noisy.",
				"I am my aunt's sister's daughter.",
			},
			want: []string{
				"Cats are good pets, for they are clean and are not noisy.",
				"He kept telling himself that one day it would all somehow make sense.",
				"I am my aunt's sister's daughter.",
				"Standing on one's head at job interviews forms a lasting impression.",
				"The chic gangster liked to start the day with a pink scarf.",
			},
		},
		{
			caseName: "not numeric order",
			cfg:      SortCfg{},
			data:     []string{"1", "5", "13", "23", "11", "21", "31"},
			want:     []string{"1", "11", "13", "21", "23", "31", "5"},
		},
		{
			caseName: "reverse order",
			cfg: SortCfg{
				isReverse: true,
			},
			data: []string{"1", "5", "13", "23", "11", "21", "31"},
			want: []string{"5", "31", "23", "21", "13", "11", "1"},
		},
		{
			caseName: "delete duplicate",
			cfg: SortCfg{
				delDuplicate: true,
			},
			data: []string{"1", "1", "5", "13", "23", "11", "11", "21", "31", "31", "31"},
			want: []string{"1", "11", "13", "21", "23", "31", "5"},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.caseName, func(t *testing.T) {
			got := mysort(&tC.cfg, tC.data)
			assert.Equal(t, tC.want, got)
		})
	}
}

func TestSortColumns(t *testing.T) {
	testCases := []struct {
		caseName string
		cfg      SortCfg
		data     []string
		want     []string
	}{
		{
			caseName: "by 2nd column",
			cfg: SortCfg{
				column: 2,
			},
			data: []string{
				"Standing on one's head at job interviews forms a lasting impression.",
				"The chic gangster liked to start the day with a pink scarf.",
				"He kept telling himself that one day it would all somehow make sense.",
				"Cats are good pets, for they are clean and are not noisy.",
				"I am my aunt's sister's daughter.",
			},
			want: []string{
				"I am my aunt's sister's daughter.",
				"Cats are good pets, for they are clean and are not noisy.",
				"The chic gangster liked to start the day with a pink scarf.",
				"He kept telling himself that one day it would all somehow make sense.",
				"Standing on one's head at job interviews forms a lasting impression.",
			},
		},
		{
			caseName: "by column out of range",
			cfg: SortCfg{
				column: 200,
			},
			data: []string{
				"Standing on one's head at job interviews forms a lasting impression.",
				"The chic gangster liked to start the day with a pink scarf.",
				"He kept telling himself that one day it would all somehow make sense.",
				"Cats are good pets, for they are clean and are not noisy.",
				"I am my aunt's sister's daughter.",
			},
			want: []string{
				"Cats are good pets, for they are clean and are not noisy.",
				"He kept telling himself that one day it would all somehow make sense.",
				"I am my aunt's sister's daughter.",
				"Standing on one's head at job interviews forms a lasting impression.",
				"The chic gangster liked to start the day with a pink scarf.",
			},
		},
		{
			caseName: "numbers numeric order",
			cfg: SortCfg{
				column:    1,
				isNumeric: true,
			},
			data: []string{"5", "23", "1", "21", "31", "13", "11"},
			want: []string{"1", "5", "11", "13", "21", "23", "31"},
		},
		{
			caseName: "by 2nd column, in reverse",
			cfg: SortCfg{
				column:    2,
				isReverse: true,
			},
			data: []string{
				"Standing on one's head at job interviews forms a lasting impression.",
				"The chic gangster liked to start the day with a pink scarf.",
				"He kept telling himself that one day it would all somehow make sense.",
				"Cats are good pets, for they are clean and are not noisy.",
				"I am my aunt's sister's daughter.",
			},
			want: []string{
				"Standing on one's head at job interviews forms a lasting impression.",
				"He kept telling himself that one day it would all somehow make sense.",
				"The chic gangster liked to start the day with a pink scarf.",
				"Cats are good pets, for they are clean and are not noisy.",
				"I am my aunt's sister's daughter.",
			},
		},
		{
			caseName: "delete duplicate",
			cfg: SortCfg{
				column:       1,
				delDuplicate: true,
			},
			data: []string{"1", "1", "5", "13", "23", "11", "11", "21", "31", "31", "31"},
			want: []string{"1", "11", "13", "21", "23", "31", "5"},
		},
		{
			caseName: "numeric sort, but column starts with letter",
			cfg: SortCfg{
				column:    1,
				isNumeric: true,
			},
			data: []string{
				"d1",
				"ad5",
				"asbv13",
				"sfg23",
				"fa11",
				"gh21",
				"31",
			},
			want: []string{
				"31",
				"ad5",
				"asbv13",
				"d1",
				"fa11",
				"gh21",
				"sfg23",
			},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.caseName, func(t *testing.T) {
			got := sortColumns(&tC.cfg, tC.data)
			assert.Equal(t, tC.want, got)
		})
	}
}

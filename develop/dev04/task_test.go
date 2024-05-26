package anagram

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestDictionary(t *testing.T) {
	testCase := []struct {
		caseDescription string
		in              []string
		length          int
		want            map[string][]string
	}{
		{
			caseDescription: "normal",
			in:              []string{"alo", "aa", "столик", "столик", "столик", "aa", "пятка", "фыв", "bb", "выф", "пятак", "слиток", "листок", "тяпка"},
			length:          2,
			want: map[string][]string{
				"пятка":  {"пятак", "тяпка"},
				"столик": {"листок", "слиток"},
			},
		},
		{
			caseDescription: "some_letters",
			in:              []string{"абвг", "абвгд", "абввг", "абвг", "абввг"},
			length:          0,
			want:            map[string][]string{},
		},
		{
			caseDescription: "empty input",
			in:              []string{},
			length:          0,
			want:            map[string][]string{},
		},
		{
			caseDescription: "unicode",
			in: []string{
				"航合投職羽張123", "羽12張3航合投職", "2張1合職3羽航投", "21羽合張投3航職", "投張羽合2職31航", "航張職2投13合羽", "航張職2投13合羽", "航張職2投13合羽",
			},
			length: 1,
			want: map[string][]string{
				"航合投職羽張123": {"21羽合張投3航職", "2張1合職3羽航投", "投張羽合2職31航", "羽12張3航合投職", "航張職2投13合羽"},
			},
		},
	}

	for _, v := range testCase {
		t.Run(v.caseDescription, func(t *testing.T) {
			res := dictionary(v.in)
			t.Log("check", v.in, "\n\toutput dictionary: ", res)
			if len(res) != v.length {
				t.Errorf("expected length: %d, got : %d", v.length, len(res))
			}
			for _, val := range res {
				if len(val) == 0 {
					t.Error(val, " len == 0")
				}
				if !sort.StringsAreSorted(val) {
					t.Error(val, " is not sorted")
				}
			}

			require.Equal(t, v.length, len(res))
			require.Equal(t, v.want, res)
		})
	}
}

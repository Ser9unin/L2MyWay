package mysort

import (
	"sort"
	"strings"
)

func MySort(cfg SortCfg, data []string) {
	if cfg.column == 1 && !cfg.isNumeric {
		data = mysort(&cfg, data)
		writeToOutput(data)
	}

	data = sortColumns(&cfg, data)
	writeToOutput(data)

}

func mysort(cfg *SortCfg, data []string) []string {
	if cfg.delDuplicate {
		data = delDuplicate(data)
	}
	if cfg.isReverse {
		sort.Sort(sort.Reverse(sort.StringSlice(data)))
	} else {
		sort.Strings(data)
	}

	return data
}

func sortColumns(cfg *SortCfg, data []string) []string {
	if cfg.delDuplicate {
		data = delDuplicate(data)
	}

	t := stringTable{
		data:      make([][]string, 0, len(data)),
		column:    cfg.column - 1,
		isNumeric: cfg.isNumeric,
	}

	for _, v := range data {
		t.data = append(t.data, strings.Fields(v))
	}

	if cfg.isReverse {
		sort.Sort(sort.Reverse(t))
	} else {
		sort.Sort(t)
	}

	for i, v := range t.data {
		data[i] = strings.Join(v, " ")
	}

	return data
}

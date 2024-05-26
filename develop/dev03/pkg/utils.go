package mysort

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

type SortCfg struct {
	column       int  // отвечает флагу -k — указание колонки для сортировки
	isNumeric    bool // отвечает флагу -n — сортировать по числовому значению
	isReverse    bool // отвечает флагу -r — сортировать в обратном порядке
	delDuplicate bool // отвечает флагу -u — не выводить повторяющиеся строки
	isMonth      bool // отвечает флагу -M — сортировать по названию месяца
	ignoreB      bool // отвечает флагу -b — игнорировать хвостовые пробелы
	sortCheck    bool // отвечает флагу -c — проверять отсортированы ли данные
	suffixN      bool // отвечает флагу -h — сортировать по числовому значению с учётом суффиксов
}

func (cfg *SortCfg) ParseConfig(args []string) (string, error) {
	flags := flag.NewFlagSet("mysort", flag.ContinueOnError)
	flags.IntVar(&cfg.column, "k", 1, "sort by column")
	flags.BoolVar(&cfg.isNumeric, "n", false, "sort by string numeric value")
	flags.BoolVar(&cfg.isReverse, "r", false, "sort in reverse order")
	flags.BoolVar(&cfg.delDuplicate, "u", false, "remove duplicates")
	flags.BoolVar(&cfg.isMonth, "M", false, "sort by month name")
	flags.BoolVar(&cfg.ignoreB, "b", false, "ignore suffix spaces")
	flags.BoolVar(&cfg.sortCheck, "c", false, "check if sorted")
	flags.BoolVar(&cfg.suffixN, "h", false, "sort by numeric including suffix value")

	if err := flags.Parse(args); err != nil {
		flags.Usage()
		return "", err
	}

	filepath := flags.Arg(0)

	return filepath, nil
}

func GetDatafromFile(filepath string) ([]string, error) {
	// Открываем файл для чтения
	file, err := os.Open(filepath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "can't open file %s: %v\n", filepath, err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lines := make([]string, 0)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return lines, err
	}

	return lines, nil
}

type stringTable struct {
	data      [][]string
	column    int
	isNumeric bool
}

func (t stringTable) Len() int {
	return len(t.data)
}

// объявляем и описываем метод Less что бы отвечать интерфейсу функции sort из встроенной библиотеки sort
func (t stringTable) Less(i, j int) bool {
	col := t.column

	//
	if col > len(t.data[i])-1 || col > len(t.data[j]) {
		col = 0
	}

	if t.isNumeric {
		n1 := trimNonNumber(t.data[i][col])
		n2 := trimNonNumber(t.data[j][col])

		i1, err := strconv.Atoi(n1)
		if err != nil {
			return (t.data[i][col] < t.data[j][col])
		}
		j1, err := strconv.Atoi(n2)
		if err != nil {
			return (t.data[i][col] < t.data[j][col])
		}

		return i1 < j1
	}
	return (t.data[i][col] < t.data[j][col])
}

// объявляем и описываем метод Swap что бы отвечать интерфейсу функции sort из встроенной библиотеки sort
func (t stringTable) Swap(i, j int) {
	t.data[i], t.data[j] = t.data[j], t.data[i]
}

// удаляет дубликаты если при запуске mysort выставлен флаг -u
func delDuplicate(data []string) []string {
	exists := make(map[string]struct{}, len(data))
	res := make([]string, 0, len(data))
	for _, v := range data {
		if _, ok := exists[v]; ok {
			continue
		}
		res = append(res, v)
		exists[v] = struct{}{}
	}

	return res
}

// пишем в stdOut
func writeToOutput(data []string) {
	for _, v := range data {
		fmt.Fprintf(os.Stdout, "%s\n", v)
	}
}

// trimNonNumber удаляет руны не содержащие числа начиная с конца строки
func trimNonNumber(str string) string {
	return strings.TrimRightFunc(str, func(r rune) bool {
		return !unicode.IsNumber(r)
	})
}

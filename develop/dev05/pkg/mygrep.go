package mygrep

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"regexp"
)

/*
=== Утилита grep ===

Реализовать утилиту фильтрации (man grep)

Поддержать флаги:
-A - "after" печатать +N строк после совпадения
-B - "before" печатать +N строк до совпадения
-C - "context" (A+B) печатать ±N строк вокруг совпадения
-c - "count" (количество строк)
-i - "ignore-case" (игнорировать регистр)
-v - "invert" (вместо совпадения, исключать)
-F - "fixed", точное совпадение со строкой, не паттерн
-n - "line num", печатать номер строки

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

type GrepCfg struct {
	after           int
	before          int
	context         int
	count           bool
	ignoreCase      bool
	inVert          bool
	fixed           bool
	lineNum         bool
	sequenceToCheck string
	reader          io.ReadCloser
	data            []string
}

func (cfg *GrepCfg) ParseConfig(args []string) error {
	flags := flag.NewFlagSet("mygrep", flag.ContinueOnError)
	flags.IntVar(&cfg.after, "A", 0, " \"after\" print +N lines after match")
	flags.IntVar(&cfg.before, "B", 0, " \"before\" prtin +N lines before match")
	flags.IntVar(&cfg.context, "C", 0, " \"context\" prtin N lines before and after match")
	flags.BoolVar(&cfg.count, "c", false, " \"count\" - number of lines")
	flags.BoolVar(&cfg.ignoreCase, "i", false, " \"ignore case\" - ignores letter case")
	flags.BoolVar(&cfg.inVert, "v", false, " \"invert\" - all except match")
	flags.BoolVar(&cfg.fixed, "F", false, " \"fixed\" - precise match with string")
	flags.BoolVar(&cfg.lineNum, "n", false, " \"line num \" - number of line")

	if err := flags.Parse(args); err != nil {
		flags.Usage()
		return err
	}

	if cfg.after == 0 {
		cfg.after = cfg.context
	}
	if cfg.before == 0 {
		cfg.before = cfg.context
	}

	cfg.sequenceToCheck = flags.Arg(0)
	// если флаг -F - "fixed", точное совпадение со строкой, не паттерн, заключаем проверяемую последовательность между ^$
	// что бы regexp сравнивал последовательность целиком
	if cfg.fixed {
		cfg.sequenceToCheck = fmt.Sprintf("^%s$", cfg.sequenceToCheck)
	}

	// если флаг -i - "ignore-case" (игнорировать регистр), добавляем в начало последовательности между (?i)
	// что бы regexp игнорировал регистр
	if cfg.ignoreCase {
		cfg.sequenceToCheck = "(?i)" + cfg.sequenceToCheck
	}

	// если читаем из STDIN то задаём его в качестве cfg.reader
	stat, _ := os.Stdin.Stat()
	if (stat.Mode() & os.ModeCharDevice) == 0 {
		cfg.reader = os.Stdin
		return nil
	}

	// если читаем из файла то задаём его в качестве cfg.reader
	file, err := os.Open(flags.Arg(1))
	if err != nil {
		fmt.Fprintf(os.Stderr, "can't open file %s: %v\n", flags.Arg(0), err)
		return err
	}
	cfg.reader = file

	return nil
}

func (cfg *GrepCfg) Run() error {
	defer cfg.reader.Close()
	scanner := bufio.NewScanner(cfg.reader)

	for scanner.Scan() {
		cfg.data = append(cfg.data, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	r, err := regexp.Compile(cfg.sequenceToCheck)
	if err != nil {
		return err
	}

	matched, count := Match(cfg, r)
	// если установлен флаг -c count то печатаем количество совпадений
	if cfg.count {
		fmt.Println(count)
	}

	PrintResult(cfg, matched)

	return nil
}

func Match(cfg *GrepCfg, r *regexp.Regexp) (matched []int, count int) {
	matched = make([]int, 0, len(cfg.data)/3)

	for i, v := range cfg.data {
		if r.MatchString(v) && !cfg.inVert || !r.MatchString(v) && cfg.inVert {
			// если установлен флаг -c count то считаем количество совпадений
			if cfg.count {
				count++
			}
			// в matched заносим строки которые совпали или наоборот не совпали в зависимости от условия inVert
			matched = append(matched, i)
		}
	}

	return matched, count
}

func PrintResult(cfg *GrepCfg, matched []int) {

	for _, v := range matched {
		// решил разделить совпадения, так будет понятно где начинается и кончается контекст при обнаружении совпадения
		fmt.Println("-----")

		if cfg.before != 0 {
			b := v - cfg.before
			if b < 0 {
				b = 0
			}
			for b != v {
				if cfg.lineNum {
					fmt.Printf("%d %s\n", b+1, cfg.data[b])
				} else {
					fmt.Printf("%s\n", cfg.data[b])
				}
				b++
			}
		}

		if cfg.lineNum {
			fmt.Printf("%d %s\n", v+1, cfg.data[v])
		} else {
			fmt.Printf("%s\n", cfg.data[v])
		}

		if cfg.after != 0 {
			a := v
			end := v + cfg.after
			if end > len(cfg.data) {
				end = len(cfg.data)
			}
			for a < end {
				if cfg.lineNum {
					fmt.Printf("%d %s\n", a+1, cfg.data[a])
				} else {
					fmt.Printf("%s\n", cfg.data[a])
				}
				a++
			}
		}
	}
}

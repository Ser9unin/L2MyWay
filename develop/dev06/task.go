package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

/*
=== Утилита cut ===

Принимает STDIN, разбивает по разделителю (TAB) на колонки, выводит запрошенные

Поддержать флаги:
-f - "fields" - выбрать поля (колонки)
-d - "delimiter" - использовать другой разделитель
-s - "separated" - только строки с разделителем

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

type Config struct {
	d       string
	f       []int
	s       bool
	reader  io.ReadCloser
	tillEnd bool
}

func (cfg *Config) ParseConfig(args []string) error {
	flags := flag.NewFlagSet("mycut", flag.ExitOnError)
	flags.BoolVar(&cfg.s, "s", false, "-s - only strings with delimiters")
	flags.StringVar(&cfg.d, "d", "\t", "-d - set new delimiter")
	flags.Func("f", "select only these fields", cfg.ParseFields)

	if err := flags.Parse(args); err != nil {
		flags.Usage()
		return err
	}

	stat, _ := os.Stdin.Stat()
	if (stat.Mode() & os.ModeCharDevice) == 0 {
		cfg.reader = os.Stdin
		return nil
	}

	return nil
}

func (cfg *Config) ParseFields(s string) error {
	if strings.Contains(s, ",") {
		f := strings.Split(s, ",")
		for _, v := range f {
			i, err := strconv.Atoi(v)
			if err != nil {
				return err
			}
			cfg.f = append(cfg.f, i-1)
		}
		return nil
	}

	if strings.Contains(s, "-") {
		var from, to int
		var err error
		if len(s) == 3 {
			from, err = strconv.Atoi(string(s[0]))
			if err != nil {
				return err
			}
			to, err = strconv.Atoi(string(s[2]))
			if err != nil {
				return err
			}
		}

		if len(s) == 2 {
			if s[0] == '-' {
				from = 1
				to, err = strconv.Atoi(string(s[1]))
				if err != nil {
					return err
				}
			}

			if s[1] == '-' {
				cfg.tillEnd = true
				from, err = strconv.Atoi(string(s[0]))
				if err != nil {
					return err
				}
				to = from
			}
		}

		for ; from <= to; from++ {
			cfg.f = append(cfg.f, from-1)
		}

		return nil
	}

	n, err := strconv.Atoi(s)
	if err != nil {
		return err
	}
	cfg.f = append(cfg.f, n-1)

	return nil
}

func (cfg *Config) Run() error {
	scanner := bufio.NewScanner(cfg.reader)

	for scanner.Scan() {
		res := strings.Split(scanner.Text(), cfg.d)
		l := len(res)

		if l == 1 && cfg.s {
			continue
		}

		if l == 1 && !cfg.s {
			fmt.Println(res[0])
			continue
		}

		if cfg.tillEnd {
			for i := cfg.f[0]; i < l; i++ {
				fmt.Printf("%s%s", res[i], cfg.d)
			}
			fmt.Println()
			continue
		}

		for _, v := range cfg.f {
			if v < l {
				fmt.Printf("%s%s", res[v], cfg.d)
			}
		}
		fmt.Println()
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	return nil
}

func main() {
	args := os.Args[:1]

	cfg := Config{}
	err := cfg.ParseConfig(args)
	if err != nil {
		log.Fatal(err)
	}

	err = cfg.Run()
	if err != nil {
		log.Fatal(err)
	}
}

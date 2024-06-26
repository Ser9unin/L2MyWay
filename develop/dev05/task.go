package main

import (
	"log"
	"os"

	mygrep "github.com/Ser9unin/L2MyWay/develop/dev05/pkg"
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

func main() {
	args := os.Args[1:]
	var config mygrep.GrepCfg

	err := config.ParseConfig(args)
	if err != nil {
		log.Fatal(err)
	}

	err = config.Run()
	if err != nil {
		log.Fatal(err)
	}
}

package main

import (
	"fmt"
	"log"
	"os"

	mywget "github.com/Ser9unin/L2MyWay/develop/dev09/pkg/wget"
)

/*
=== Утилита wget ===

# Реализовать утилиту wget с возможностью скачивать сайты целиком

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/
func main() {
	args := os.Args[:1]

	wgetApp := mywget.Mywget{}

	err := wgetApp.ParseConfig(args)
	if err != nil {
		log.Fatal(err)
	}

	err = wgetApp.Run()

	if err != nil {
		fmt.Fprintf(os.Stderr, "Runtime error: %v\n", err)
		os.Exit(1)
	}

	os.Exit(0)
}

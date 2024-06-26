Что выведет программа? Объяснить вывод программы. Объяснить как работают defer’ы и их порядок вызовов.

```go
package main

import (
	"fmt"
)


func test() (x int) {
	defer func() {
		x++
	}()
	x = 1
	return
}


func anotherTest() int {
	var x int
	defer func() {
		x++
	}()
	x = 1
	return x
}


func main() {
	fmt.Println(test())
	fmt.Println(anotherTest())
}
```

Ответ:
```
2
1
Правила работы defer:

1. Аргументы отложенной функции оцениваются при вычислении оператора defer:
func a() {
    i := 0
    defer fmt.Println(i) // выведет 0
    i++
    return
}
2. Отложенные вызовы функций выполняются в порядке LIFO после возврата из основной функции;
3. Отложенные функции могут считывать и присваивать именованные возвращаемые значения возвращаемой функции.

Функция test попадает под третье правило, поэтому функции defer удается увеличить переменную.

А в функции anotherTest, в функциональном литерале defer переменная x увеличивается с 0 на 1 (однако это никак не влияет на саму x в основной функции, x просто вычисляется), затем в основной функции переменной x присваивается 1 и возвращается. defer не имеет доступа к переменной x

```

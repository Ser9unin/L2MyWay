Что выведет программа? Объяснить вывод программы.

```go
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func asChan(vs ...int) <-chan int {
	c := make(chan int)

	go func() {
		for _, v := range vs {
			c <- v
			time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
		}

		close(c)
	}()
	return c
}

func merge(a, b <-chan int) <-chan int {
	c := make(chan int)
	go func() {
		for {
			select {
			case v := <-a:
				c <- v
			case v := <-b:
				c <- v
			}
		}
	}()
	return c
}

func main() {

	a := asChan(1, 3, 5, 7)
	b := asChan(2, 4 ,6, 8)
	c := merge(a, b )
	for v := range c {
		fmt.Println(v)
	}
}
```

Ответ:
```
Сначала будет выведено некоторое количество чисел переданных в виде каналов a и b
вывод будет в случайном порядке, так как порядок выполнения case внутри select не определен
select{
	case:
	case:
}

потом будет идти поток нулей, так как канал "c" объявленный в main не закрыт и чтение из него продолжается, но данные в него больше не поступают,
так как после передачи всех значений в канал "с" объявленный в asChan этот канал закрывается, таким образом в функции merge, тут:
for {
	select {
	case v := <-a:
		c <- v
	case v := <-b:
		c <- v
	}
}

мы ведем бесконечное чтение из закрытого канала в открытый канал.
```

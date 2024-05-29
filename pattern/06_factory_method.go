package pattern

import "fmt"

/*
	Реализовать паттерн «фабричный метод».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Factory_method_pattern
*/

// Фабричный метод — порождающий паттерн проектирования, решает проблему создания различных продуктов,
// без указания конкретных классов продуктов.

/*
	ипользуется когда заране неизвестны типы и зависимости с которыми прийдется работать
	плюсы: Не нужно выстраивать бизнес логику отталкиваясь от того с каким типом прийдется работать

	Использование:
	1) когда заранее неизвестны типы и зависимости объектов, с которыми должен работать код
	2) когда нужна возможность пользователям расширять части фреймворка или библиотеки
	3) когда нужно экономить системные ресурсы, повторно используя уже созданные объекты, вместо порождения новых

	+:
	1) избавляет классы от привязки к конкретным классам продуктов
	2) выделяет код производства продуктов в одно место, упрощая поддержку кода
	3) упрощает добавление новых продуктов в программу
	4) реализует принцип открытости/закрытости

	-:
	1) может привести к созданию больших параллельных иерархий классов, так как для каждого класса
	продукта надо создать свой подкласс создателя
	2) может появиться "божественный конструктор"
*/

// Интерфейс продукта фабрики
type carI interface {
	getInfo() string
}

// у классов производимых фабрикой могут быть отличия, в данном случае каждый тип авто описан структурой с разными совйствами
type policeCar struct {
	color     string
	hp        int
	clearance int
	lightbar  bool
}

func newPoliceCar() carI {
	return &policeCar{
		color:     "Синие полоски",
		hp:        220,
		clearance: 8,
		lightbar:  true,
	}
}

func (pc *policeCar) getInfo() string {
	println("Это наша корова, и мы её доим.")
	return fmt.Sprintf("Цвет: %s, Кони: %d, Дорожный просвет: %d, Люстра: %v", pc.color, pc.hp, pc.clearance, pc.lightbar)
}

type sportCar struct {
	color     string
	hp        int
	clearance int
	zoomzoom  bool
}

func newSportCar() carI {
	return &sportCar{
		color:     "Красный",
		hp:        500,
		clearance: 4,
		zoomzoom:  true,
	}
}

func (sc *sportCar) getInfo() string {
	println("Очень злая гоночная тачка. Врум, Врум!")
	return fmt.Sprintf("Цвет: %s, Кони: %d, Дорожный просвет: %d, \"ВРУМ! ВРУМ\": %v", sc.color, sc.hp, sc.clearance, sc.zoomzoom)
}

type defaultCar struct {
	color     string
	hp        int
	clearance int
	doors     int
	wheels    int
}

func newDefaultCar() carI {
	return &defaultCar{
		color:     "Унылый серый",
		hp:        150,
		clearance: 12,
		doors:     4,
		wheels:    4,
	}
}
func (dc *defaultCar) getInfo() string {
	println("Непримечательное корыто")
	return fmt.Sprintf("Цвет: %s, Кони: %d, Дорожный просвет: %d, Двери: %d, Колеса: %d", dc.color, dc.hp, dc.clearance, dc.doors)
}

// Производство машин
func factory(typeCar string) carI {
	switch typeCar {
	case "sport car":
		return newSportCar()
	case "police car":
		return newPoliceCar()
	default:
		return newDefaultCar()
	}
}

func main() {
	car1 := factory("sport car")
	car2 := factory("police car")
	car3 := factory("def")

	println(car1.getInfo())
	println(car2.getInfo())
	println(car3.getInfo())
}

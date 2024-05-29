package pattern

/*
	Реализовать паттерн «посетитель».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Visitor_pattern
*/

/*
	Посетитель - поведенческий паттерн, позволяющий добавлять поведение в структуру без ее изменения.
	Использование:
	1) когда нужно выполнить какую-то операцию над всеми элементами сложной структуры объектов (дерево)
	2) когда не нужно засорять классы несвязанными операциями
	3) когда новое поведение имеет смысл только для некоторых классов из существующей иерархии

	+:
	1) упрощает добавление операций, работающих со сложными структурами объектов
	2) объединение родственных операций в одном классе
	3) посетитель может накапливать состояние при обходе структуры элементов

	-:
	1) паттерн не оправдан, если иерархия элементов часто меняется
	2) может привести к нарушению инкапсуляции элементов
*/

// описываем интерфейсом Element, который принимает посетителя
type Element interface {
	Accept(visitor Visitor)
}

// описываем конкретный элемент А
type CertainElementA struct{}

func NewCertainElementA() *CertainElementA {
	return &CertainElementA{}
}

// описываем метод Accept тем самым наш элемент отвечает интерфейсу Element
func (ce *CertainElementA) Accept(visitor Visitor) {
	visitor.VisitCertainElementA(ce)
}

// описываем конкретный элемент B
type CertainElementB struct{}

func NewCertainElementB() *CertainElementB {
	return &CertainElementB{}
}

// описываем метод Accept тем самым наш элемент отвечает интерфейсу Element
func (ce *CertainElementB) Accept(visitor Visitor) {
	visitor.VisitCertainElementB(ce)
}

// описываем интерфейс Visitor в котором указываем какие элементы он будет посещать
type Visitor interface {
	VisitCertainElementA(ce *CertainElementA)
	VisitCertainElementB(ce *CertainElementB)
}

// описываем конкретного посетителя
type CertainVisitor struct{}

func NewCertainVisitor() *CertainVisitor {
	return &CertainVisitor{}
}

// описываем метод для посещения элемента A
func (cv *CertainVisitor) VisitCertainElementA(ce *CertainElementA) {
	println("Посетил элемент A")
}

func (cv *CertainVisitor) VisitCertainElementB(ce *CertainElementB) {
	println("Посетил элемент B")
}

func main() {
	visitor := NewCertainVisitor()

	elementA := NewCertainElementA()
	elementA.Accept(visitor)

	elementB := NewCertainElementB()
	elementB.Accept(visitor)
}

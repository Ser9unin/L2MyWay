package pattern

/*
	Реализовать паттерн «строитель».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Builder_pattern
*/
/*
Строитель — это порождающий паттерн проектирования, используется для построения сложных объектов покомпонентно,
дает возможность создания немного отличающихся в значениях, но одинаковых в конструкции объектов.
	Использование:
	1) построение сложного объекта от его представления

	+:
	1) позволяет изменить внутреннее представление продукта
	2) инкапсулирует код для построения и представления
	3) обеспечивает контроль за этапами процесса строительства

	-:
	1) для каждого типа продукта должен быть создан отдельный строитель
	2) классы строителя должны быть изменяемыми
	3) может затруднить внедрение зависимостей
*/
type Valve struct {
	Name string
	DN   int
	Kvs  float64
}
type ValveBuilderI interface {
	Name(val string) ValveBuilderI
	DN(val int) ValveBuilderI
	Kvs(val float64) ValveBuilderI

	Build() Valve
}

type valveBuilder struct {
	name string
	dN   int
	kvs  float64
}

func NewValveBuilder() ValveBuilderI {
	return valveBuilder{}.Name("Radiator valve").DN(15).Kvs(1.6)
}

func (v valveBuilder) Name(val string) ValveBuilderI {
	v.name = val
	return v
}

func (v valveBuilder) DN(val int) ValveBuilderI {
	v.dN = val
	return v
}

func (v valveBuilder) Kvs(val float64) ValveBuilderI {
	v.kvs = val
	return v
}

func (v valveBuilder) Build() Valve {
	return Valve{
		Name: v.name,
		DN:   v.dN,
		Kvs:  v.kvs,
	}
}

func main() {
	valveBuilder := NewValveBuilder()
	//default valve
	defaultValve := valveBuilder.Build()
	println(defaultValve)

	// new High flow valve
	highFlowValve := valveBuilder.Name("High flow valve").DN(150).Kvs(5000).Build()
	println(highFlowValve)
}

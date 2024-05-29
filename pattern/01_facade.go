package pattern

/*
	Реализовать паттерн «фасад».
Объяснить применимость паттерна, его плюсы и минусы,а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Facade_pattern
*/

/*
	Фасад - структурный паттерн, реализует простой доступ к сложной системе
	+:
	1) изолирует клиентов от компонентов сложной подсистемы

	-:
	1) рискует стать антипаттерном "божественный объект"
*/

func main() {
	driver := driver{}
	driver.StartEngine()
}

// фасад это водитель с функцией завести двигатель
type driver struct{}

func (d *driver) StartEngine() {
	println("Водитель вставил ключ в зажигание и повернул")
	battery := battery{}
	starter := starter{}

	battery.SupplyPower()
	starter.RotateСrankshaft()
}

// после поворота ключа в зажигании в автомобиле происходит множество взаимодействий между компонентами о которых не знает клиент который попросил водителя завести авто
type battery struct{}

func (b *battery) SupplyPower() {
	b.SupplyPowerToStarter()
	b.SupplyPowerToInterruptor()
}

func (b *battery) SupplyPowerToInterruptor() {
	println("аккумулятор подаёт питание на прерыватель")
	interruptor := interruptor{}
	interruptor.DistributePowerToSparkPlugs()
}

func (b *battery) SupplyPowerToStarter() {
	println("аккумулятор подаёт питание на стартер")
}

type interruptor struct {
}

func (i *interruptor) DistributePowerToSparkPlugs() {
	println("прерыватель распределяет питание между свечами зажигания")
}

type starter struct{}

func (s *starter) RotateСrankshaft() {
	println("стартер крутит коленвал")
	crankshaft := crankshaft{}
	crankshaft.PushPullPistons()
}

type crankshaft struct{}

func (c *crankshaft) PushPullPistons() {
	println("коленвал толкает поршни")
}

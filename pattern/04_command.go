package pattern

/*
	Реализовать паттерн «комманда».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Command_pattern
*/

/*
	Команда - поведенческий паттерн, в котором запросы или операции являются отдельными объектами
	Использование:
	1) когда вы хотите параметризировать объекты выполняемым действием
	2) когда вы хотите ставить операции в очередь, выполнять их по расписанию или передовать по сети
	3) когда вам нужна операция отмены

	+:
	1) убирает прямую зависимость между объектами, вызывающими операции, и объектами, которые их непосредственно выполняют
	2) позволяет реализовать простую отмену и повтор операций
	3) позволяет реализовать отложенный запуск операций
	4) позволяет собирать сложные команды из простых
	5) реализует принцип открытости/закрытости

	-:
	1) усложняет код программы из-за введения множества дополнительных классов
*/

type Command interface {
	Execute()
}

type Receiver struct {
	Name string
}

func (r *Receiver) Act() {
	println(r.Name, " выполняет действие")
}

type CertainCommand struct {
	receiver *Receiver
}

func (c *CertainCommand) Execute() {
	c.receiver.Act()
}

type Invoker struct {
	command Command
}

func (i *Invoker) SetCommand(command Command) {
	i.command = command
}

func (i *Invoker) ExecuteCommand() {
	i.command.Execute()
}

func main() {
	// создаём получателя, команду и вызываюшего
	receiver := &Receiver{Name: "Получатель 1"}
	command := &CertainCommand{receiver: receiver}
	invoker := &Invoker{}

	// задаём и запускаем команду
	invoker.SetCommand(command)
	invoker.ExecuteCommand() // вывод Получатель 1 выполняет действие
}

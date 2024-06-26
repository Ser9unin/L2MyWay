package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"net"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

/*
=== Взаимодействие с ОС ===

Необходимо реализовать собственный шелл

встроенные команды: cd/pwd/echo/kill/ps
поддержать fork/exec команды
конвеер на пайпах

Реализовать утилиту netcat (nc) клиент
принимать данные из stdin и отправлять в соединение (tcp/udp)
Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func main() {

	myshell, _ := strconv.Atoi(os.Getenv("MYSHELL"))
	os.Setenv("MYSHELL", strconv.Itoa(myshell+1))

	for {
		fmt.Print("shell: ")

		// Читаю ввод
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()

		// Если ввод содержит пайп
		if strings.Contains(scanner.Text(), "|") {
			// Выполняю папйплайн и если были ошибки вывожу их
			piplineinput := strings.Split(scanner.Text(), "|")
			errs := executePipeline(piplineinput)
			for i := range errs {
				fmt.Fprintln(os.Stderr, errs[i])
			}
		} else {
			// Иначе выполняю либо билт-ин команду либо из $PATH
			commandStrSet := strings.Fields(scanner.Text())
			err := execute(commandStrSet)

			if err != nil {
				fmt.Fprintln(os.Stderr, err)
			}
		}

	}
}

func executePipeline(cmds []string) (errs []error) {
	var err error
	// Кол-во комманд
	length := len(cmds)

	// Создаю массив под нужное кол-во комманд
	cmd := make([]*exec.Cmd, length)

	// Итерирусь по аргументу функции И разбиваю комманду на имя и аргументы
	// Создаю для каждой комманды объект *exec.Cmd
	for i := 0; i < length; i++ {
		in1 := strings.Fields(cmds[i])

		cmd[i] = exec.Command(in1[0], in1[1:]...)
	}
	// Итерируюсь по массиву *exec.Cmd
	for i := 0; i < length; i++ {
		// Если комманда первая, то ее Stdout будет писать в stdin след комманды
		// Stdin остается без изменений
		if i == 0 {
			cmd[i].Stdin = os.Stdin
			cmd[i].Stdout, err = cmd[i+1].StdinPipe()

			if err != nil {
				errs = append(errs, err)
			}
			//Если комманда в промежутке между первой и последней
			// то ее Stdout будет писать в stdin след комманды
			// а Stdin был определен на предыдущей итерации
		} else if i != length-1 {
			cmd[i].Stdout, err = cmd[i+1].StdinPipe()
			if err != nil {
				errs = append(errs, err)
			}
			// Если комманда последняя, то она выводит результат в Stdout
		} else {
			cmd[i].Stdout = os.Stdout
		}

		cmd[i].Stderr = os.Stderr

		// Заупускаю комманду
		err = cmd[i].Start()
		if err != nil {
			errs = append(errs, err)
		}
	}

	// Жду завершения всех комманд
	for i := 0; i < length; i++ {
		err = cmd[i].Wait()
		if err != nil {
			log.Fatal(err)
		}
	}
	return
}

func execute(in []string) (err error) {

	countArg := len(in)

	if countArg == 0 {
		return
	}

	switch in[0] {
	case "pwd":
		err = pwd()
	case "cd":
		if countArg == 1 {
			err = cd(os.Getenv("HOME"))
		} else {
			err = cd(in[1])
		}
	case "echo":
		echo(in[1:])
	case "kill":
		err = kill(in[1:])
	case "nc":
		err = netcat(in[1:])
	case "exit":
		fmt.Println("exiting from the shell")
		os.Exit(0)
	default:
		err = executeFromPath(in)
	}
	return
}

func executeFromPath(args []string) error {
	cmd := exec.Command(args[0], args[1:]...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		return err
	}

	return nil
}

func cd(arg string) (err error) {
	if err = os.Chdir(arg); err != nil {
		return
	}

	os.Setenv("OLDPWD", os.Getenv("PWD"))

	path, err := os.Getwd()
	if err != nil {
		return
	}
	os.Setenv("PWD", path)

	return
}

func pwd() (err error) {
	path, err := os.Getwd()
	if err != nil {
		return
	}

	fmt.Println(path)

	return
}

func echo(args []string) {
	var flag bool = true

	for _, v := range args {
		if v == "-n" {
			flag = false
		} else {
			fmt.Print(v)
		}
	}
	if flag {
		fmt.Print("\n")
	}
}

func kill(args []string) error {
	if len(args) == 0 {
		return errors.New("kill: not enough arguments")
	}

	for _, v := range args {
		pid, err := strconv.Atoi(v)
		if err != nil {
			return errors.New("kill: illegal pid: " + v)
		}
		p, err := os.FindProcess(pid)
		if err != nil {
			return err
		}
		err = p.Kill()
		if err != nil {
			return err
		}
	}
	return nil
}

func netcat(args []string) error {
	if len(args) < 3 {
		return errors.New("too few args (type, ip, port)")
	}
	//
	conn, err := net.Dial(args[0], args[1]+":"+args[2])
	if err != nil {
		return err
	}
	defer conn.Close()

	// Читаем сообщение со стандартного ввода
	reader := bufio.NewReader(os.Stdin)
	message, _ := reader.ReadString('\n')

	_, err = conn.Write([]byte(message))

	if err != nil {
		return err
	}
	return nil
}

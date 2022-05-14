package main

import (
	"fmt"
	"os"

	"github.com/U/command"
)

func main() {
	//Обработка флагов
	fl, err := command.Finit()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	//Создание массива строк с выбором метода ввода

	W := command.Inoutchoose()

	// Обработка и вывод

	command.Out(W, fl)

}

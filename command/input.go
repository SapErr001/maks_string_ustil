package command

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Instr() []string {
	var str []string

	in := bufio.NewScanner(os.Stdin)

	for in.Text() != "end" {
		if err := in.Err(); err != nil {
			fmt.Fprintln(os.Stderr, "Ошибка ввода:", err)
		}
		in.Scan()
		str = append(str, in.Text())
	}
	//fmt.Println(str)
	return str[:len(str)-1]
}

func Infile() []string {

	file, err := os.Open(os.Args[len(os.Args)-1])

	if err != nil {
		return Instr()
	} else {
		scanner := bufio.NewScanner(file)

		var str []string

		defer file.Close()
		for scanner.Scan() {
			str = append(str, scanner.Text())
		}
		//fmt.Println(str)
		return str
	}
}

var Nout int

func Inoutchoose() []string {

	if len(os.Args) == 1 {
		return Instr()
	} else {

		return Infile()
	}
}

func Outfile(a []Uniqstr, f Fl) {

	file, err := os.Create(os.Args[len(os.Args)-Nout])

	if err != nil {
		fmt.Println("Unable to create file:", err)
		os.Exit(1)
	}
	defer file.Close()

	for i := 0; i < len(a); i++ {

		//fmt.Fprintln(file, a[i].num, " ", a[i].str)
		if *f.Usec {

			fmt.Fprintln(file, a[i].num, " ", a[i].str)
		}
		if *f.Used {
			fmt.Fprintln(file, a[i].str)
		}
		if *f.Useu {
			if a[i].num == 1 {
				fmt.Fprintln(file, a[i].str)
			}
		}
		if !*f.Usec && !*f.Used && !*f.Useu {

			fmt.Fprintln(file, a[i].num, " ", a[i].str)
		}
	}

}

func FlagsCount() bool {
	for i := 0; i < len(os.Args); i++ {
		if strings.Count(strings.Join(os.Args[0:], ""), "Out.txt") == 1 &&
			strings.Count(strings.Join(os.Args[0:], ""), ".txt") == 1 {
			Nout = 1
			return true
		}

		if strings.Count(strings.Join(os.Args[0:], ""), "Out.txt") == 1 &&
			strings.Count(strings.Join(os.Args[0:], ""), ".txt") == 2 {
			Nout = 2
			return true
		}
	}

	return false
}

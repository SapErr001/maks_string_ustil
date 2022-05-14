package command

import (
	"errors"
	"flag"
)

type Fl struct {
	Usec   *bool
	Used   *bool
	Useu   *bool
	Usei   *bool
	Usef   *bool
	Uses   *bool
	Useint *int
}

func Sw(a *bool) int {
	var q int
	if *a == true {
		q = 1
	} else {
		q = 0
	}
	return q
}

func Finit() (a Fl, err error) {

	var Flinit Fl
	a = Flinit
	Flinit.Usec = flag.Bool("c", false, "help message for flagname")
	Flinit.Used = flag.Bool("d", false, "help message for flagname")
	Flinit.Useu = flag.Bool("u", false, "help message for flagname")
	Flinit.Usei = flag.Bool("i", false, "help message for flagname")
	Flinit.Usef = flag.Bool("f", false, "help message for flagname")
	Flinit.Uses = flag.Bool("s", false, "help message for flagname")
	Flinit.Useint = flag.Int("Num", 0, "help message for flagname")
	flag.Parse()

	if Sw(Flinit.Usec)+Sw(Flinit.Used)+Sw(Flinit.Useu) > 1 {

		err = errors.New("Ошибка ввода флагов")

	}
	return Flinit, err
}

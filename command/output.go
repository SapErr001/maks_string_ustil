package command

import (
	"fmt"
)

func Chout(a []Uniqstr, f Fl) bool {
	if FlagsCount() {
		Outfile(a, f)
		return true
	}
	return false
}

func Cout(a []Uniqstr) {

	for i := 0; i < len(a); i++ {

		fmt.Println(a[i].num, " ", a[i].str)
	}
}
func Dout(a []Uniqstr) {

	for i := 0; i < len(a); i++ {
		fmt.Println(a[i].str)
	}
}

func Uout(a []Uniqstr) {

	for i := 0; i < len(a); i++ {
		if a[i].num == 1 {
			fmt.Println(a[i].str)
		}

	}
}

var Nstr []string

func ChFlags(str []string, f Fl) []Numnum {
	var nstr []string

	if *f.Usef {
		nstr = Slicefield(str, f)

	}
	if *f.Uses {
		nstr = Slicestr(str, f)

	}
	if *f.Usei {
		nstr = Lowerstr(str)

	}
	if !*f.Usef && !*f.Uses && !*f.Usei {

		return Findsamestrnumber(str)
	}
	return Findsamestrnumber(nstr)
}
func Out(str []string, f Fl) {
	var r []Numnum
	r = ChFlags(str, f)

	if *f.Usec {
		if Chout(Newsamestr(str, r), f) == false {
			Cout(Newsamestr(str, r))
		} else {
			Chout(Newsamestr(str, r), f)
		}
	}
	if *f.Used {
		if Chout(Newsamestr(str, r), f) == false {
			Dout(Newsamestr(str, r))
		} else {
			Chout(Newsamestr(str, r), f)
		}
	}
	if *f.Useu {

		if Chout(Newsamestr(str, r), f) == false {
			Uout(Newsamestr(str, r))
		} else {
			Chout(Newsamestr(str, r), f)
		}
	}
	if !*f.Usec && !*f.Used && !*f.Useu {

		if Chout(Newsamestr(str, r), f) == false {
			Cout(Newsamestr(str, r))
		} else {
			Chout(Newsamestr(str, r), f)
		}
	}
}

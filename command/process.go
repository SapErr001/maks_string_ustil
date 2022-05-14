package command

import (
	"strings"
)

type Uniqstr struct {
	num int
	str string
}

func Lowerstr(a []string) []string {
	var lownewslicestr []string
	for _, v := range a {

		lownewslicestr = append(lownewslicestr, strings.ToLower(v))

	}
	return lownewslicestr

}

func Slicefield(a []string, f Fl) []string {
	var Fieldnewstr []string
	for _, v := range a {

		s := strings.Fields(v)

		if len(s) < *f.Useint {
			Fieldnewstr = append(Fieldnewstr, " ")
		} else if len(s) > 0 && v != " " {
			Fieldnewstr = append(Fieldnewstr, strings.Join(s[*f.Useint:], " "))
		} else {
			Fieldnewstr = append(Fieldnewstr, v)
		}
	}
	return Fieldnewstr
}

func Slicestr(a []string, f Fl) []string {
	var newslicestr []string

	for _, v := range a {

		if len(v) > 0 {
			newslicestr = append(newslicestr, v[*f.Useint:len(v)])
		} else {
			newslicestr = append(newslicestr, v)
		}
	}
	return newslicestr
}

func Findsamestr(a []string) []Uniqstr {
	var Uniq []Uniqstr
	for _, v := range a {
		fl := 0
		if len(Uniq) == 0 {
			Uniq = append(Uniq, Uniqstr{0, ""})
			Uniq[0].str = v
			Uniq[0].num++
			//Uniq = append(Uniq, Uniqstr{0, ""})
		} else {
			for i := 0; i < len(Uniq); i++ {

				if Uniq[i].str == v {

					Uniq[i].num++
					fl++
					break
				}

			}
			if fl == 0 {
				Uniq = append(Uniq, Uniqstr{1, v})
			}

		}
	}
	return Uniq
}

type Numnum struct {
	Ncol     int 	//номер строки
	Nsamestr int	//количество повторений данного номера строки
}

func Findsamestrnumber(a []string) []Numnum {

	var Uniq []Uniqstr
	var Numstr []Numnum
	for n, v := range a {
		fl := 0
		if len(Uniq) == 0 {
			Uniq = append(Uniq, Uniqstr{0, ""})
			Uniq[0].str = v
			Uniq[0].num++

			Numstr = append(Numstr, Numnum{n, Uniq[0].num})
			//Uniq = append(Uniq, Uniqstr{0, ""})
		} else {
			for i := 0; i < len(Uniq); i++ {

				if Uniq[i].str == v {

					Uniq[i].num++
					fl++
					Numstr[i].Nsamestr++
					break
				}

			}
			if fl == 0 {
				Uniq = append(Uniq, Uniqstr{1, v})

				Numstr = append(Numstr, Numnum{n, 1})
			}

		}
	}
	return Numstr
}

func Newsamestr(a []string, m []Numnum) []Uniqstr {
	var Un []Uniqstr
	for n, v := range a {
		for i := 0; i < len(m); i++ {
			if n == m[i].Ncol {
				Un = append(Un, Uniqstr{m[i].Nsamestr, v})
			}
		}
	}
	return Un
}

package main

import (
	"fmt"
	"os"
	"sort"
)

type Ascendente []string

func (a Ascendente) Len() int           { return len(a) }
func (a Ascendente) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a Ascendente) Less(i, j int) bool { return a[i] < a[j] }

type Descendente []string

func (a Descendente) Len() int           { return len(a) }
func (a Descendente) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a Descendente) Less(i, j int) bool { return a[i] > a[j] }

func main() {
	var n int
	var t_str string
	fmt.Print("Ingrese el numero de strings: ")
	fmt.Scanln(&n)
	sl := make([]string, n)
	for x := 0; x < n; x++ {
		fmt.Print("Ingrese: ")
		fmt.Scanln(&t_str)
		sl[x] = t_str
	}

	sort.Sort(Ascendente(sl))

	asc_file, err := os.Create("ascendente.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer asc_file.Close()

	for x := 0; x < n; x++ {
		asc_file.WriteString(sl[x])
		asc_file.WriteString("\n")
	}

	sort.Sort(Descendente(sl))

	desc_file, err := os.Create("descendente.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer desc_file.Close()

	for x := 0; x < n; x++ {
		desc_file.WriteString(sl[x])
		desc_file.WriteString("\n")
	}
}

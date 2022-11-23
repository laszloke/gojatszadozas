package main

import "fmt"

func main() {
		fmt.Println("Hello vilag")

		x := map[string][]string{
			"Marci":  []string{"alukalas", "ugralas", "szaladgalas"},
			"Frocsi": []string{"vakarozas", "morgolodas", "bepisiles", "kaja lopas"},
		}
		fmt.Println(x)

		fmt.Println("feltoltes utan")
	for i, v := range x {
		fmt.Println(i)
		for j, z := range v {
			fmt.Println("\t", j, z)
		}
	}

	x["Tamas"] = []string{"fetrenges", "hisztizes", "pizza rendeles","ejjeli uzengetes","uszas","teszt"}
	fmt.Println("uj elem hozza adva")
	for i, v := range x {
		fmt.Println(i)
		for j, z := range v {
			fmt.Println("\t", j, z)
		}
	}

	fmt.Println("torles ")
	delete(x, "Tamas")
	delete(x, "Marci")
	for i, v := range x {
		fmt.Println(i)
		for j, z := range v {
			fmt.Println("\t", j, z)
		}
	}
}

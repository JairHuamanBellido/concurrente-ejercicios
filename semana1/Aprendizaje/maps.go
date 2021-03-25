package main

import "fmt"

func main() {

	// Creación de un mapa
	teams := map[string][]string{
		"Orcas":   []string{"Fred", "Ralph", "Bijou"},
		"Lions":   []string{"Sarah", "Peter", "Billie"},
		"Kittens": []string{"Waldo", "Raul", "Ze"},
	}

	// Acceder a una propiedad del mapa
	fmt.Println(teams["Orcas"])

	teams["Dogs"] = []string{"Cookie", "Kaity"}
	fmt.Println(teams["Dogs"])

	// Borrar propiedad
	delete(teams, "Dogs")
	fmt.Println(teams)

	// Mapas como SET

	inSet := map[int]bool{}
	vals := []int{1, 2, 3, 4, 4, 4, 5}

	for _, v := range vals {
		inSet[v] = true
	}

	fmt.Println(len(inSet))
	// Verificar si el valor 2 se encuentra en el SET
	fmt.Println(inSet[2])

	intSet := map[int]struct{}{}
	valsfoo := []int{5, 10, 2, 5, 8, 7, 3, 9, 1, 2, 10}
	for _, v := range valsfoo {
		intSet[v] = struct{}{}
	}

	// Validar si el 5 esta dentro del SET
	if _, ok := intSet[5]; ok {
		fmt.Println("5 is in the set")
	}

}

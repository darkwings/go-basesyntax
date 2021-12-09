package maps

import "fmt"

func MapsDemo() {

	// Inizializzazione con make
	myMap := make(map[string]int)
	myMap["alice"] = 31
	myMap["bob"] = 32

	// Iterazione con range
	for name, age := range myMap {
		fmt.Printf("%s\t%d\n", name, age)
	}

	// Versione alternativa per inizializzazione
	myMap2 := map[string]int{
		"alice":   31,
		"charlie": 34,
	}
	fmt.Println(myMap2["alice"])

	// funzione builtin per rimuovere
	delete(myMap2, "alice")

	// Se leggo un elemento che non esiste ottengo il valore 0 del tipo indicato come value
	fmt.Println(myMap["aaa"]) // 0, il valore di default di un int

	// Verifica della effettiva non esistenza
	_, ok := myMap["aaa"]
	if !ok {
		fmt.Println("Non esiste")
	}
}

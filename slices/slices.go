package slices

import "fmt"

func SlicesDemo() {
	// slice come parte di un array
	// in questo caso mantiene un riferimento all'array originale, per cui
	// manipolare la slice si riflette nella manipolazione dell'array sottostante
	months := [...]string{1: "January", 2: "February", 3: "March", 4: "April", 5: "May", 6: "June", 7: "July", 8: "August", 9: "September", 10: "October", 11: "November", 12: "December"}
	q2 := months[4:7]
	summer := months[6:9]
	fmt.Println(q2)
	fmt.Println(summer)

	// Very inefficient!!!
	for _, s := range summer {
		for _, q := range q2 {
			if s == q {
				fmt.Printf("%s appears in both\n", s)
			}
		}
	}

	// slice creata ad hoc
	// Vi sono funzioni builtin come append() che permettono di aggiungere elementi ad una slice
	var runes []rune // rune è un int32 di fatto, rappresenta un carattere unicode
	for _, r := range "Hello, BF" {
		runes = append(runes, r)
	}
	fmt.Printf("%q\n", runes)

}

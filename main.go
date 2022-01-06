package main

import (
	"fmt"
	"go-basesyntax/interfaces"
	"go-basesyntax/singleton"
)

// Questo metodo viene eseguito prima di tutto il resto
func init() {
	fmt.Println("init main")
}

func main() {
	fmt.Println("main")
	interfaces.DemoCast()

	manager1 := singleton.GetManagerInstance()
	manager1.DoWork()
	manager2 := singleton.GetManagerInstance() // won't log creation, manager2 is actually the same as manager
	manager2.DoWork()
}

package interfaces

import "fmt"

func init() {
	fmt.Println("init interface")
}

type MyType interface {
	WriteSomething()
}

type OfficialMyTypeImpl struct{}

func (m OfficialMyTypeImpl) WriteSomething() {
	fmt.Println("OfficialMyTypeImpl OK")
}

type FakeMyTypeImpl struct{}

func (m FakeMyTypeImpl) WriteSomething() {
	fmt.Printf("FakeMyTypeImpl")
}

func DemoCast() {

	var myType MyType // VS Code segnala warning ma dobbiamo mostrare il 'cast' per cui ignoriamo

	// Assegno alla variabile myType un oggetto OfficialMyTypeImpl
	myType = OfficialMyTypeImpl{}

	// ...

	// x.(T)
	// - se T è un tipo "concrete", l'asserzione verifica che il dynamic type di x sia identico a T
	// - se T è un interfaccia, l'asserzione verifica che il dynamic type di x soddisfa T.
	// Nell'esecuzione seguente, l'effetto è quello di assegnare a offMyType l'implementazione
	// "concreta" di myType, ossia OfficialMyTypeImpl. Di fatto, abbiamo implementato un cast.
	// Se il cast fallisce, si ha un panic
	offMyType, ok := myType.(OfficialMyTypeImpl)
	if !ok {
		panic("cast failed")
	}
	offMyType.WriteSomething() // scrive OfficialMyTypeImpl
}

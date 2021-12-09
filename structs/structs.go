package structs

import (
	"fmt"
	"time"
)

type Employee struct {
	ID            int
	Name, Address string
	DoB           time.Time
	Position      string
	Salary        int
	ManagerID     int
	evaluation    string // Private (package) member
}

type Point struct{ X, Y int }

func StructsDemo() {
	// Inizializzazione
	p := Point{1, 2}
	fmt.Println(p)

	// Inizializzazione con parametri nominali
	e := Employee{ID: 1, Salary: 120} // i field non usati sono inizializzati al loro valore 0, l'ordine è facoltativo
	fmt.Println(e)
	StructsUpdateSalary(&e, 140)
	fmt.Println(e) // Si vedrà che il Salary è passato a 260

	// Dato che è molto frequente passare il puntatore di una struct, è possibile definirlo direttamente
	pp := &Employee{ID: 2, Salary: 250, Name: "Frank"}
	fmt.Println(pp)

	pp.IncrementSalary(34)
	fmt.Println(pp)
}

// In Go il passaggio di parametri è by-value. Nel caso delle struct è preferibile
// passare il puntatore, in modo da poter operare direttamente sull'oggetto
func StructsUpdateSalary(e *Employee, inc int) {
	e.Salary += inc
}

// Metodo aggiunto ad una struct
func (e *Employee) IncrementSalary(inc int) {
	e.Salary += inc
}

// Idem come sopra, per rendere il passaggio di dati più efficiente
// ma non modifichiamo la struct originale.
func StructsGetIncrementedSalary(e *Employee, inc int) int {
	return e.Salary + inc
}

// Metodo abbastanza inutile, ma permette di vedere come accedere ad un membro
// 'privato' della struct.
// Go da questo punto di vista è semplice. Se qualcosa ha nome che inizia con
// lettera minuscola, è visibile solo all'interno del package
func (e *Employee) SetEvaluation(eval string) {
	e.evaluation = eval
}

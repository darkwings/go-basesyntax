package inputflags

import (
	"flag"
	"fmt"
)

// Vediamo ora come gestire al meglio i flag da linea di comando

type Celsius float64
type Fahrenheit float64

type celsiusFlag struct{ Celsius }

// La funzione String() si applica al tipo Celsius (così come si fa con una Struct)
func (c Celsius) String() string {
	return fmt.Sprintf("%g°C", c)
}

func (f Fahrenheit) String() string {
	return fmt.Sprintf("%g°F", f)
}

// Conversione da gradi Fahrenheit a Celsius
func FToC(f Fahrenheit) Celsius { return Celsius((f - 32) * 5 / 9) }

func (f *celsiusFlag) Set(s string) error {
	var unit string
	var value float64
	fmt.Sscanf(s, "%f%s", &value, &unit) // no error check needed
	switch unit {
	case "C", "°C":
		f.Celsius = Celsius(value)
		return nil
	case "F", "°F":
		f.Celsius = FToC(Fahrenheit(value))
		return nil
	}
	return fmt.Errorf("invalid temperature %q", s)
}

// A questo punto "celsiusFlag" ha i metodi Set(string) e Get (ereditato da Celsius)
// Di fatto soddisfa l'interfaccia flag.Value, che permette di eseguire parsing
// da linea di comando

func CelsiusFlag(name string, value Celsius, usage string) *Celsius {
	f := celsiusFlag{value}
	flag.CommandLine.Var(&f, name, usage)
	return &f.Celsius
}

/*
Il programma main() sarà qualcosa del genere

import (
	"flag"
	"fmt"

	"frank.com/go/basesyntax/inputflags"
)

// Definisce il flag temp che può essere letto da linea di comando
// Capitolo 7.4 del libro "The GO Programming Language"
var temp = inputflags.CelsiusFlag("temp", 20.0, "the temperature")

func main() {

	flag.Parse()
	fmt.Println(*temp)
}

===

Una tipica sessione di lavoro sarà

$ go build
$ ./basesyntax
20°C
$ ./basesyntax -temp -18C
-18°C
$ ./basesyntax -temp 212°F
100°C
*/

package iota

import "fmt"

func IotaSample() {
	type Weekday int
	const (
		Sunday Weekday = iota
		Monday
		Tuesday
		Wednesday
		Thursday
		Friday
		Saturday
	)
	fmt.Println(Monday)
	fmt.Println(1 == Monday) // Weekday alla fine è un int

	// Posso usare anche lo shift, in modo che i vari possibili valori siano
	// bitwise shifted (e si possano fare cose interessanti come mostrato nelle
	// func bitwiseOp64() e bitwiseOp() di seguito)
	type Flags uint
	const (
		FlagUp Flags = 1 << iota // is up
		FlagBroadcast
		FlagLoopback
		FlagPointToPoint
		FlagMulticast
	)
	fmt.Println(FlagUp)
	fmt.Println(FlagBroadcast)
	fmt.Println(FlagLoopback)
	fmt.Printf("%08b\n", FlagUp)        // 00000001
	fmt.Printf("%08b\n", FlagBroadcast) // 00000010
	fmt.Printf("%08b\n", FlagLoopback)  // 00000100
}

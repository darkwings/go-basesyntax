package structs

// Membro inaccessibile. Quando si vuole una cosa del genere,
// non si può fare altro che usare una struct
type Counter struct{ n int }

func (c *Counter) N() int     { return c.n }
func (c *Counter) Increment() { c.n++ }
func (c *Counter) Reset()     { c.n = 0 }

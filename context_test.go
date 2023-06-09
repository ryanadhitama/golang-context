package golang_context

import (
	"fmt"
	"testing"
	"context"
)

func TestContext(t *testing.T) {
	background := context.Background()
	fmt.Println(background)

	todo := context.TODO()
	fmt.Println(todo)
}

// Context with value
func TestContextWithValue(t *testing.T) {
	a := context.Background()
	b := context.WithValue(a, "b", "B")
	c := context.WithValue(a, "c", "C")
	d := context.WithValue(b, "d", "D")
	e := context.WithValue(b, "e", "E")
	f := context.WithValue(c, "f", "F")
	g := context.WithValue(c, "g", "G")

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(f)
	fmt.Println(g)

	fmt.Println(b.Value("b"))
	fmt.Println(f.Value("c"))
}
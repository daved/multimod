package main

import "fmt"
import "github.com/daved/multimod/a"
import "github.com/daved/multimod/b"
import "github.com/daved/multimod/internal/c"

func main() {
	fns := []func() string{
		a.Version,
		b.Version,
		c.ZeroOneZero,
	}

	for _, fn := range fns {
		fmt.Println(fn())
	}
}

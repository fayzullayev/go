package main

import (
	"fmt"
	"strings"
)

type str string

func (s *str) Upper() {
	*s = str(strings.ToUpper(string(*s)))
}

func (s *str) Log() {
	fmt.Println(*s)
}
func main() {
	var name str = str("mirodil")
	name.Log()
	name.Upper()
	name.Log()
}

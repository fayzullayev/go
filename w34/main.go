//https://golangbot.com/pointers/

package main

import (
	"fmt"
)

func mutate(s []rune) string {
	s[0] = 'a'
	return string(s)
}
func main() {
	h := "😙ello"

	fmt.Println(mutate([]rune(h)))
}

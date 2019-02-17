package main

import (
	"fmt"
	"strings"
)

func replaceBlank(source string) {
	ret := strings.Replace(source, " ", "20%", -1)
	fmt.Println(ret)
}

func main() {
	s := " a b c "
	replaceBlank(s)
}

package main

import (
	"fmt"
	"strings"
)

func main() {
	var str1 string = "This is nieda"
	fmt.Printf("%s", strings.Split(str1, " "))
}

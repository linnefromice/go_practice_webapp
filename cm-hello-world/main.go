package main

import (
	"fmt"
	"github.com/linnefromice/go_practice_webapp/cm-hello-world/subpkg"
)

func main() {
	fmt.Println(subpkg.Hello())
	fmt.Println(subpkg.Golang())
	fmt.Println(Goodbye())
}

package main

import (
	"fmt"

	"github.com/tonkatsu7/leftpad"
)

func main() {
	fmt.Println(leftpad.Format("hello", 15))
}

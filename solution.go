package main

import (
	"fmt"

	"github.com/kyokomi/emoji/v2"
)

func main() {
	a := emoji.Sprint(":world_map:")
	fmt.Println("Hello " + a[:len(a)-1] + "!")
}

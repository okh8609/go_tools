package greeting

import (
	"fmt"

	"github.com/fatih/color"
)

func Say(s string) {
	fmt.Println(s)
}

func SayWithColor(s string) {
	color.Red(s)
}

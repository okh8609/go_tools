package main

import (
	"github.com/okh8609/go_tools/greeting"
	"github.com/okh8609/go_tools/my_flag"
)

func main() {
	greeting.Say("Hello")
	greeting.SayWithColor("World")
	my_flag.Fooo()
	my_flag.Fooo2()

}

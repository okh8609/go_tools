package my_flag

import (
	"flag"
	"log"
)

func Fooo() {
	var name string
	flag.StringVar(&name, "name", "golang", "info...")
	flag.StringVar(&name, "n", "golang", "info...")
	flag.Parse()
	log.Printf("name: %s", name)
}

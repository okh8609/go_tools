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

func Fooo2() {

	flag.Parse()

	go_cmd := flag.NewFlagSet("go", flag.ExitOnError)
	py_cmd := flag.NewFlagSet("py", flag.ExitOnError)

	var name string
	go_cmd.StringVar(&name, "n", "GO~", "info")
	py_cmd.StringVar(&name, "i", "PY~", "info")

	args := flag.Args()
	if len(args) <= 0 {
		return
	}

	switch args[0] {
	case "go":
		_ = go_cmd.Parse(args[1:])
	case "py":
		_ = py_cmd.Parse(args[1:])
	}

	log.Printf("name: %s", name)
}

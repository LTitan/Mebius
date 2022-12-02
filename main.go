package main

import (
	"github.com/LTitan/Mebius/pkg/options"
)

func main() {
	root := options.NewRootCommand()
	root.ExecuteOrDie()
}

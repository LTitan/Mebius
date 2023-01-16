package main

import (
	"github.com/LTitan/Mebius/pkg/controllers"
	"github.com/LTitan/Mebius/pkg/options"
)

func main() {
	root := options.NewRootCommand()
	controllers.RegisterFramework(root)
	root.ExecuteOrDie()
}

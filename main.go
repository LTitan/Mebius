package main

import (
	"github.com/LTitan/Mebius/pkg/controllers"
	"github.com/LTitan/Mebius/pkg/options"
	"github.com/LTitan/Mebius/pkg/server"
)

func main() {
	root := options.NewRootCommand()
	controllers.NewFramework(root).RegisterCommand()
	server.NewServer(root).RegisterCommand()
	root.ExecuteOrDie()
}

package main

import (
	"github.com/LTitan/Mebius/pkg/controllers"
	"github.com/LTitan/Mebius/pkg/factory"

	"github.com/LTitan/Mebius/pkg/gateway"
	"github.com/LTitan/Mebius/pkg/options"
	"github.com/LTitan/Mebius/pkg/server"
)

func main() {
	root := options.NewRootCommand()
	applications := []factory.Application{
		controllers.NewFramework(root),
		server.NewServer(root),
		gateway.NewGateway(root),
	}
	for i := range applications {
		applications[i].RegisterCommand()
	}
	root.ExecuteOrDie()
}

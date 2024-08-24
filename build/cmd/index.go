package main

import (
	"github.com/outofforest/build"
	"github.com/outofforest/buildgo"

	me "build"
)

var commands = map[string]build.Command{
	"setup": {Fn: me.Setup, Description: "Installs tools required by development environment"},
}

func init() {
	buildgo.AddCommands(commands)
	build.RegisterCommands(commands)
}

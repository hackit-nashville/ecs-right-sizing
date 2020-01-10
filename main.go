package main

import (
	"github.com/hackit-nashville/ecs-right-sizing/cmd"
)

// Version of cli. Overwritten during build
var Version = "development"

func main() {
	cmd.Execute(Version)
}

package main

import (
	"hackit-ecs-right-sizing/cmd"
)

// Version of cli. Overwritten during build
var Version = "development"

func main() {
	cmd.Execute(Version)
}

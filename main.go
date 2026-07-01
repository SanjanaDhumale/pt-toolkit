package main

import (
	"github.com/SanjanaDhumale/pt-toolkit/cmd"
	"github.com/SanjanaDhumale/pt-toolkit/internal/config"
)

func main() {
	config.Init()
	cmd.Execute()
}
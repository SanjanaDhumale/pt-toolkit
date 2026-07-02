package services

import (
	"fmt"

	"github.com/SanjanaDhumale/pt-toolkit/internal/plugin"
)

func ListPlugins() {

	fmt.Println("===================================")
	fmt.Println("JMeter Plugin Manager")
	fmt.Println("===================================")

	plugin.List()
}
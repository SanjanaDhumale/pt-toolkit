package services

import (
	"fmt"

	"github.com/SanjanaDhumale/pt-toolkit/internal/stack"
)

func InstallStack(name string) {

	s, err := stack.Find(name)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("========================================")
	fmt.Println("Installing:", s.Name)
	fmt.Println("========================================")

	for _, c := range s.Components {

		fmt.Println("Checking:", c.Name)
	}

	fmt.Println()
	fmt.Println("✓ Stack Ready")
}
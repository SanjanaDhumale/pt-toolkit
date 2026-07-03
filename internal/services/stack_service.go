package services

import (
	"fmt"

	"github.com/SanjanaDhumale/pt-toolkit/internal/stack"
)

func ListStacks() {

	fmt.Println("==========================================")
	fmt.Println(" Available Performance Engineering Stacks")
	fmt.Println("==========================================")
	fmt.Println()

	stacks := stack.GetStacks()

	for _, s := range stacks {

		fmt.Println("------------------------------------------")
		fmt.Println("📦", s.Name)
		fmt.Println()

		fmt.Println("Components")

		for _, c := range s.Components {
			fmt.Println("  ✓", c.Name())
		}

		fmt.Println()
	}
}
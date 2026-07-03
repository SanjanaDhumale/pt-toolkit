package ui

import (
	"fmt"
	"strings"
)

const width = 70

func Header(title string) {

	line := strings.Repeat("═", width)

	fmt.Println()
	fmt.Println("╔" + line + "╗")
	fmt.Printf("║ %-70s ║\n", title)
	fmt.Println("╚" + line + "╝")
}

func Section(title string) {

	fmt.Println()
	fmt.Printf("▶ %s\n", title)
	fmt.Println(strings.Repeat("─", 45))
}

func Step(step string) {
	fmt.Printf("➜ %s\n", step)
}

func Success(msg string) {
	fmt.Printf("   ✅ %s\n", msg)
}

func Error(msg string) {
	fmt.Printf("   ❌ %s\n", msg)
}

func Warning(msg string) {
	fmt.Printf("   ⚠️  %s\n", msg)
}

func Info(msg string) {
	fmt.Printf("   ℹ️  %s\n", msg)
}

func Blank() {
	fmt.Println()
}
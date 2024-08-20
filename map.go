package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

func clearScreen() {
	switch runtime.GOOS {
	case "linux", "darwin":
		cmd := exec.Command("clear") // Linux and macOS
		cmd.Stdout = os.Stdout
		cmd.Run()
	case "windows":
		cmd := exec.Command("cmd", "/c", "cls") // Windows
		cmd.Stdout = os.Stdout
		cmd.Run()
	default:
		fmt.Print("\033[H\033[2J") // ANSI escape codes for other platforms
	}
}

func printMap() {
	printStatus()
	fmt.Println(strings.Repeat(" ", 150))
	fmt.Println(strings.Repeat(" ", 150))
	fmt.Println(strings.Repeat(" ", 150))
	fmt.Println(strings.Repeat(" ", 150))
	fmt.Print(strings.Repeat(" ", 50), strings.Repeat("-", 50), strings.Repeat(" ", 50), "\n")
	fmt.Print(strings.Repeat(" ", 50), "|", strings.Repeat(".", 48), "|", "\n")
	fmt.Print(strings.Repeat(" ", 50), "|", strings.Repeat(".", 48), "|", "\n")
	fmt.Print(strings.Repeat(" ", 50), "|", strings.Repeat(".", 48), "|", "\n")
	fmt.Print(strings.Repeat(" ", 50), "|", strings.Repeat(".", 48), "|", "\n")
	fmt.Print(strings.Repeat(" ", 50), "|", strings.Repeat(".", 48), "|", "\n")
	fmt.Print(strings.Repeat(" ", 50), "|", strings.Repeat(".", 48), "|", "\n")
	fmt.Print(strings.Repeat(" ", 50), "|", strings.Repeat(".", 48), "|", "\n")
	fmt.Print(strings.Repeat(" ", 50), "|", strings.Repeat(".", 48), "|", "\n")
	fmt.Print(strings.Repeat(" ", 50), "|", strings.Repeat(".", 48), "|", "\n")
	fmt.Print(strings.Repeat(" ", 50), "|", strings.Repeat(".", 48), "|", "\n")
	fmt.Print(strings.Repeat(" ", 50), strings.Repeat("-", 50), strings.Repeat(" ", 50), "\n")
	fmt.Println(strings.Repeat(" ", 150))
	fmt.Println(strings.Repeat(" ", 150))
	fmt.Println(strings.Repeat(" ", 150))
	fmt.Println(strings.Repeat(" ", 150))
	printStats()
	fmt.Println()
}

func printStatus() {
	fmt.Println("4 pieces of gold")
}

func printStats() {
	fmt.Printf("Level: %v\t  ", 1)
	fmt.Printf("HP: %v(%v)\t  ", 10, 20)
	fmt.Printf("Str: %v\t  ", 10)
	fmt.Printf("Exp: %v\t  ", 15)
}

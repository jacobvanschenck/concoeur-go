package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
)

func printLine(s string) {
	fmt.Print(s)
	fmt.Print("\n")
	fmt.Print("\r\033[K")
}

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

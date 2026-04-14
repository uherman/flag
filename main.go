package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		printUsage()
		os.Exit(1)
	}

	arg := os.Args[1]
	name := strings.TrimLeft(arg, "-")

	if name == "" || name == "help" || name == "h" {
		printUsage()
		return
	}
	if name == "list" || name == "l" {
		printList()
		return
	}
	if name == "all" || name == "a" {
		printAll()
		return
	}

	f := LookupFlag(name)
	if f == nil {
		fmt.Fprintf(os.Stderr, "Unknown flag: %s\nRun 'flag -list' to see available flags.\n", name)
		os.Exit(1)
	}

	width := 60
	height := width * f.Ratio[0] / f.Ratio[1]
	if height%2 != 0 {
		height++
	}

	canvas := NewCanvas(width, height)
	f.Draw(canvas)
	fmt.Print(canvas.Render())
}

func printUsage() {
	fmt.Println("Usage: flag -<country>")
	fmt.Println()
	fmt.Println("Options:")
	fmt.Println("  -all, -a     Print all flags")
	fmt.Println("  -list, -l    List all available flags")
	fmt.Println("  -help, -h    Show this help")
	fmt.Println()
	fmt.Println("Examples:")
	fmt.Println("  flag -sweden")
	fmt.Println("  flag -usa")
	fmt.Println("  flag -jp")
}

func printAll() {
	width := 60
	for _, f := range Flags {
		height := width * f.Ratio[0] / f.Ratio[1]
		if height%2 != 0 {
			height++
		}
		canvas := NewCanvas(width, height)
		f.Draw(canvas)
		fmt.Printf("  %s\n", f.Name)
		fmt.Print(canvas.Render())
		fmt.Println()
	}
}

func printList() {
	fmt.Println("Available flags:")
	fmt.Println()
	for _, f := range Flags {
		aliases := strings.Join(f.Aliases, ", ")
		fmt.Printf("  %-15s  (%s)\n", f.Name, aliases)
	}
}

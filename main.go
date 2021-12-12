package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/danicat/simpleansi"
)

var maze []string

func loadMaze(file string) error {
	f, err := os.Open(file)
	if err != nil {
		return err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		maze = append(maze, line)
	}

	return nil
}

func draw() {
	for _, line := range maze {
		fmt.Println(line)
	}
}

func prepareTerminal() {
	cbTerm := exec.Command("stty", "cbreak", "-echo")
	cbTerm.Stdin = os.Stdin

	err := cbTerm.Run()
	if err != nil {
		log.Fatalf("unable to activate cbreak mode: %v", err)
	}
}

func restoreTerminal() {
	cookedTerm := exec.Command("stty", "-cbreak", "echo")
	cookedTerm.Stdin = os.Stdin

	err := cookedTerm.Run()
	if err != nil {
		log.Fatalf("unable to restore cooked mode: %v", err)
	}
}

func readInput() (string, error) {
	buffer := make([]byte, 100)

	cnt, err := os.Stdin.Read(buffer)
	if err != nil {
		return "", err
	}

	if cnt == 1 && buffer[0] == 0x1b {
		return "ESC", nil
	}

	return "", nil
}

func main() {
	err := loadMaze("maze.txt")
	if err != nil {
		log.Fatalf("failed to load maze: %v", err)
	}

	prepareTerminal()
	defer restoreTerminal()

	// game loop
	for {
		// update screen
		simpleansi.ClearScreen()
		draw()

		// process input
		input, err := readInput()
		if err != nil {
			log.Print("error reading input:", err)
			break
		}

		// process movement
		// process collisions
		// check game over

		if input == "ESC" {
			break
		}

		// repeat
	}
}

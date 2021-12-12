package main

import "testing"

func ExampleDraw() {
	maze = []string{
		"####",
		"#  #",
		"####",
	}

	draw()
	//output:
	//####
	//#  #
	//####
}

func TestLoadMaze(t *testing.T) {
	err := loadMaze("maze.txt")
	if err != nil {
		t.Fatalf("failed to load maze: %v", err)
	}
}

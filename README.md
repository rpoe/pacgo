# Step 02: Examples and testing

In this lesson you will learn how to:

- Test your go code using `go test`
- Write testable examples

## Task 01: Testing the draw function with an Example

Create a file called `main_test.go` with the following content:

```go
package main

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
```

Run `go test` to run the tests. 

Try also `go test -v` and `go test -cover`.

## Task 02: Testing the loadMaze function

Add a new test to `main_test.go`:

```go
func TestLoadMaze(t *testing.T) {
	err := loadMaze("maze.txt")
	if err != nil {
		t.Fatalf("failed to load maze: %v", err)
	}
}
```

Noticed any changes in `go test -v` and `go test -cover`?

## Next step

Congratulations, step 02 is complete!

Run the following command on your terminal to go to step 03:

```sh
$ git checkout step03
```
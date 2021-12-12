# Step 01: The Game Loop

In this lesson you will learn how to:

- Declare and initialize variables
- Use for loops
- Work with functions
- Read data from files

## Overview

We've got the basics covered, now it's time to get this game started!

First, we are going to setup our game loop, which will run all the routines necessary for the game to work.

Next we are going to read the maze data. We have a file called `maze.txt` that's a text representation of the maze (you can open it in a text editor if you like). You may assume that:

```
- # represents a wall
- . represents a dot
- P represents the player
- G represents the ghosts (enemies)
- X represents the power up pills
```

We will be working with this representation for the most part of the tutorial, but don't worry, we will get to emojis soon enough!

## Task 01: The Game Loop

One thing that is very particular of game programs is that we have something called the `game loop`, which is basically an infinite loop that runs all the operations necessary for a game - over and over - like reading inputs, updating the screen, controlling the AI, and so on...

In order to create our game loop we are going to use the `for` statement as described on the previous lesson:

```go
package main

import "fmt"

func main() {
    for {
        fmt.Println("I'm the game loop")
        break
    }
}
```

### For loops and iterations

As in many other programing languages, in Go you can control the number of iterations in a for loop. In order to do that we need to use a variable:

```go
package main

import "fmt"

func main() {
    for i := 0; i < 5; i++ {
        fmt.Println("I'm the game loop")
    }
}
```

This form of the `for` loop allows you to control the number of iterations. The first part (before the first `;`) is the initialization part. The `:=` operator allows you to declare and initialize a variable at the same time. In this case we are initializing `i` to `0`. The compiler automatically assigns the type of `i` to `int` based on the value after `:=`. That's the main benefit of using this operator.

The middle part is the `condition` that is evaluated in every iteration. The loop will run again if `condition` is evaluated to `true`.

Finally, the last part is the increment/decrement part where you can modify the iteration variable according to your needs. In this case we are just incrementing it by 1 each time.

In Go, both the initializer and the increment/decrement part can be ommitted, but then you need to control the loop conditions somewhere else:

```go
package main

import "fmt"

func main() {
    var i int
    for i < 5 {
        fmt.Println("I'm the game loop")
        i += 1
    }
}
```

As you might have noticed in the example above, you can also declare (and optionally initialize) a variable in Go with the keyword `var`:

```go
var a int
var a = 0
a := 0
```
All three options above produce the same result - an integer variable called `a` with initial value `0`. Please note that in the first example `var a int`, you are not specifying the initial value, so Go will apply the default zero value for the given data type.

But when creating a variable inside a for loop only the third option is allowed. Also any variable created in the `for` statement will only valid inside the `for` loop.

```go
for i := 0; i < 5; i++ {
    fmt.Println("I'm the game loop")
}
fmt.Println(i) // won't compile - i is not defined here
```

### For loop with range

Another type of `for` loop is the `range` loop, as we are going to see in the next task.

## Task 02: Load the Maze

Let's start by reading the `maze.txt` file.

We are going to use the function `Open` from the `os` package to open it, and a scanner object from the buffered IO package (`bufio`) to read it to memory (to a global variable called `maze`). Finally we need to release the file handler by calling `os.Close`.

All that comes together as the code below:

```go
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
```

Now let's break it down and see what's going on.

Please note that you need to import `bufio` and `os` packages as shown below:

```go
import "bufio"
import "os"
```

Alternatively, since you already have one import (`fmt`), you can add it as a list:

```go
import (
    "bufio"
    "fmt"
    "os"
)
```

The `os.Open()` function returns a pair of values: a file and an error. Returning multiple values from a function is a common pattern in Go, specially for returning errors.

```go
f, err := os.Open(file)
```

The `:=` operator is an assignment operator, but with the special property that it automatically infers the type of the variable(s) based on the value(s) on the right hand side.

Keep in mind that Go is a strongly typed language, but that nice feature saves us the trouble of specifying the type when it's possible to infer it.

In the case above, Go automatically infers the type for both `f` and `err` variables.

When a function returns an error it is a common pattern to check the error immediately afterwards:

```go
    f, err := os.Open(file)
    if err != nil {
        // do something with err
        log.Print("...")
        return
    }
```

Note: It is a good practice to keep the "happy path" aligned to the left, and the sad path to the right (i.e., terminating the function early).

`nil` in Go means no value is assigned to a variable.

The `if` statement executes a statement if the condition is true. It can optionally have an initialization clause just like the `for` statement, and an `else` clause that runs if the condition is false. Please keep in mind that the scope of the variable created will just be the if statement body. For example:

```go
// optional initialization clause
if foo := rand.Intn(2); foo == 0 {
    fmt.Print(foo) // foo is valid here
} else {
    fmt.Print(foo) // and here
}
// but you can't use foo here!
```

Another interesting aspect of the `loadMaze` code is the use of the `defer` keyword. It basically says to call the function after `defer` at the end of the current function. It is very useful for cleanup purposes and in this case we are using it to close the file we've just opened:

```go
func loadMaze(file) error {
    f, err := os.Open(file)
    // omitted error handling
    defer f.Close() // puts f.Close() in the call stack

    // rest of the code

    return nil
    // f.Close is called implicitly
}
```

The next part of the code just reads the file line by line and appends it to the maze slice:

```go
    scanner := bufio.NewScanner(f)
    for scanner.Scan() {
        line := scanner.Text()
        maze = append(maze, line)
    }
```

A scanner is a very convenient way to read a file. `scanner.Scan()` will return true while there is something to be read from the file, and `scanner.Text()` will return the next line of input.

The `append` built-in function is responsible for adding a new element to the `maze` slice.

## Task 03: Updating the Screen

Once we have the maze file loaded into memory we need to print it to the screen.

One way to do that is to iterate over each entry in the `maze` slice and print it. This can be conveniently done with the `range` operator:

```go
func printScreen() {
    for _, line := range maze {
        fmt.Println(line)
    }
}
```

Please note that we are using the `:=` assignment operator to initialize two values: the underscore (_) and the `line` variable. The underscore is just a placeholder for where the compiler would expect a variable name. Using the underscore means that we are ignoring that value.

In the case of the `range` operator, the first return value is the index of the element, starting from zero. The second return value is the value itself.

If we did not write the underscore character to ignore the first value, the range operator would return just the index (and not the value). For example:

```go
for idx := range maze {
    fmt.Println(idx)
}
```

Since in this case we only care about the content and not the index, we can safely ignore the index by assigning it to the underscore.

## Task 03: Updating the game loop

Now that we have both `loadMaze` and `printScreen` functions, we should update the `main` function to initialize the maze and print it on the game loop. See the code below:

```go
func main() {
    // initialise game

    // load resources
    err := loadMaze("maze01.txt")
    if err != nil {
        log.Println("failed to load maze:", err)
        return
    }

    // game loop
    for {
        // update screen
        printScreen()

        // process input

        // process movement

        // process collisions

        // check game over

        // Temp: break infinite loop
        break

        // repeat
    }
}
```

Like always we are keeping the happy path to the left, so if the `loadMaze` function fails we use `log.Println` to log it and then `return` to terminate the program execution. Since we are using a new package, `log`, please make sure it is added to the import section:

```go
import (
    "bufio"
    "fmt"
    "log"
    "os"
)
```

Some IDEs, like `vscode`, can be configured to do this automatically for you.

Note: one could also use `log.Fatalln` for the same effect, but we need to make sure that any deferred calls are executed before exiting the `main` function, and functions in the `log.Fatal` family skip deferred function calls by calling `os.Exit(1)` internally. We don't have any deffered calls in the main function yet, but we will add one in the next chapter.

Now that we've finished the game loop modifications we can run the program with `go run` or compile it with `go build` and run it as a standalone program.

```sh
go run main.go
```

You should see the maze printed to the terminal.

## Next step

Congratulations, step 01 is complete!

Run the following command on your terminal to go to step 02:

```sh
$ git checkout step02
```
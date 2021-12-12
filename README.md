# Step 00: Hello (Game) World

In this lesson you will learn how to:

- Run your first `go` program
- Work with go modules
- Build (compile) your first program

## Task 01: main.go

As in any development language tutorial, we are going to start by running a "Hello world" type of program. This is what it looks like in go:

```go
package main

import "fmt"

func main() {
    fmt.Println("hello, world!")
}
```
Create a file named `main.go` and copy and paste the code above.

### Running your first Go program

Now that we have a `main.go` file (`.go` is the file extension for Go source code), you can run it by using the command line tool `go`. Just type:

```sh
$ go run main.go
hello, world!
```

Note: although it might look like the `go` command line tool is interpreting the code we just wrote, it is actually compiling `main.go` to a temporary binary file and running it under the hood.

If you run code with `go run` just make sure to include all the needed files in the same command line, as it won't automatically add any files that are not specified.

For example, in order to use `go run` with two files `a.go` and `b.go` you should type:

```sh
$ go run a.go b.go
```

## Task 02: Building your first Go program

Now that we know how to run a file with `go run`, we are going to learn how to build an executable for you program. No surprisingly, the command for doing that is `go build`. But before we can use `go build`, we need to setup our Go modules.

Go modules is how a go program track its dependencies. It also gives you ways to manage the version of your releases. 

Note: The full use of go modules is out of the scope of this tutorial. If you want to learn more about `go module`, please refer to the official docs: https://go.dev/ref/mod

In order to initialize a go module you need to give it a name. So in your console type:

```sh
$ go mod init github.com/your-user-name/pacgo
```

Note: replace `your-user-name` with yours :)

After running that program you should see a new file called `go.mod` in your current directory.

Now you will be able to `build` your executable with `go build`. The executable name will be the last name in it's module name. In the case above is `pacgo`.

```sh
$ go build
$ ./pacgo
hello, world!
```

### Understanding the program

Now let's have a close look of what the program does.

First line is the `package` name. Every valid Go program must have one of these. Also, if you want to make a **runnable** program you must have a package called `main` and a `main` function, which will be the entry point of the program.

We are creating an executable program so we are using `package main` on the top of the file.

Next are the `import` statements. You use those to make code in other packages accessible to your program.

Finally the `main` function. You define function in Go with the keyword `func` followed by its name, its parameters in between a pair of parentheses, followed by the return value and finally the function body, which is contained in a pair of curly brackets `{}`. For example:

```go
func main() {
    // I'm a function body
}
```

This is a function called `main` which takes no parameters and returns nothing. That's the proper definition of a main function in Go, as opposed to the definitions in other languages where the main function may take the command line arguments and/or return an integer value.

In the game main function we have some comments (any text after `//` is a comment) that are acting as placeholders for the actual game code. We'll use those to drive each modification in an orderly way.

The most important concept in a game is what is called the game loop. That's basically an infinite loop where all the game mechanics are processed.

A loop in Go is defined with the keyword `for`. The `for` loop can have an initializer, a condition and a post processing step, but all of those are optional. If you omit all you have a loop that never ends:

```go
for {
    // I never end
}
```

We can exit an infinite loop with a `break` statement. For example. the code below has the same effect as our `hello, world!` application written in task 01.

```go
func main() {
    for {
        fmt.Println("hello, world")
        break
    }
}
```

## Next step

Congratulations, step 00 is complete!

Run the following command to go to step 01:

```sh
$ git checkout step01
```

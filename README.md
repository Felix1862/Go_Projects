# Project List
   ## Project 1: Hello World
### Folder: hello-world/

This is the simplest possible Go program. It prints "Hello World" to the terminal to confirm your Go environment is working correctly.


 How to Run
 ```bash
    cd hello-world
    go run main.go
```
----

## Project 2: Inner & Outer Block
Folder: inner-outer-block/

This Go program demonstrates how **variable scoping** works between outer and inner code blocks.

What It Demonstrates
- A variable city is declared in the main (outer) block.
- A variable country is declared in an inner block.
- The city variable is accessible in both blocks.
- The country variable is **only accessible** inside the inner block.


  Sample Output
  ```bash
      $ go run main.go
      UK
      London
      London
  ```

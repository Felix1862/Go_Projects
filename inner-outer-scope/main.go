// inner_outer_scope.go
//
// This program demonstrates how variable scope works in Go using nested code blocks.
// - 'city' is declared in the outer block and is accessible everywhere within main()
// - 'country' is declared inside an inner block and is only accessible within that block
//
// This helps beginners understand lexical scoping and visibility rules in Go.

package main

import "fmt"

func main() {
    // Declare a variable `city` in the outer block
    city := "London"

    // Begin inner block
    {
        // Declare a new variable `country` only visible within this block
        country := "UK"

        // Print the `country` variable (inner scope)
        fmt.Println(country)

        // Access the `city` variable from the outer block (still in scope)
        fmt.Println(city)
    }
    // End inner block

    // Outside the inner block, `country` no longer exists
    // But we can still access `city` here
    fmt.Println(city)
}

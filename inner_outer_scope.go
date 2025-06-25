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

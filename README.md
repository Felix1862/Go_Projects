# Go_Projects
# Go Simple Projects

This repository is a collection of beginner-friendly Go projects.  
Each folder contains a single Go program with clear comments and explanations to help learners understand Go fundamentals.

---

## Project 1: Hello World

### File: `hello-world/main.go`

This is the simplest possible Go program. It prints **"Hello World"** to the terminal to confirm your Go environment is working.

### How to Run

Make sure you have Go installed. You can verify it with:

```bash
go version

---
## Project 2: Inner & Outet Block 
This simple Go program demonstrates **variable scoping** using inner and outer code blocks.

## What It Does

- Declares a variable `city` in the main (outer) block.
- Creates an inner block where a new variable `country` is defined.
- Shows that:
  - `city` is accessible from both inner and outer blocks.
  - `country` is **only accessible** within the inner block.

## Output

```bash
$ go run main.go
UK
London
London

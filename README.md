# Monkey Programming Language Interpreter

A Go implementation of the Monkey programming language interpreter, developed as part of learning compiler construction through ["Writing an Interpreter in Go"](https://interpreterbook.com/) by Thorsten Ball.

## Features

- **Data Types**:
  - Integers, Booleans, Strings
  - Arrays and Hashes (dictionaries)
- **Language Features**:
  - Variable bindings
  - Arithmetic expressions
  - Conditional statements (if/else)
  - Return statements
  - First-class and higher-order functions
  - Closures
  - Recursion
- **Built-in Functions**:
  - String manipulation (`len`, split, concatenation)
  - Array operations (`push`, `first`, `last`, `rest`)
- **REPL Support**

## Installation

1. Ensure Go (1.20+) is installed
2. Clone repository:
   ```bash
   git clone https://github.com/hamidoujand/interpreter.git

3. Build and run:
   ```bash
   cd interpreter
   go build -o monkey
    ./monkey


## Example Programs

1. Fibonacci Sequence:
    ```js
    let fibonacci = fn(n) {
    if (n <= 1) {
        return n;
    } else {
        fibonacci(n-1) + fibonacci(n-2);
    }
    };
    fibonacci(10); // Returns 55


## Acknowledgements
Based on concepts from ["Writing an Interpreter in Go"](https://interpreterbook.com/) by Thorsten Ball."
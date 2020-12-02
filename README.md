# The Donut Programming Language

The donut programming language is being developed as a part of an attempt to understand how programming languages work.

The task at hand is to build a tree walking interpreter. Such interpreters parse the source code, build an abstract syntax tree (AST) and then evaulate this tree. We would have to build our own lexer, parser and tree representation. We'll examine in depth what an AST is, how to build this tree, how to evaulate it and how to extend our language with data structure and built-in functions.

# Installation

```
go get go.ajitem.com/donut
```

# Feature list 

The donut programming language has the following desired feature list:
* C-like syntax
* variable bindings
* integers, booleans and arithmetic expressions
* built-in function (standard library)
* first-class and higher-order functions
* closures
* string data structure
* an array data structure
* a hash data structure (map)

# Sample Syntax

```
    let age = 1;
    let name = "Donut";
    let result = 2 * (10 / 5)
    
    let array = [1, 2, 3, 4, 5];
    array[2] // => 3
    
    let object = {"name": "Ajitem", "age": 29}
    object["name"] // => "Ajitem"
    
    let sum = func(a, b) { return a + b };
    let diff = func(a, b) { a - b };
    
    let factorial = func(x) {
        if (n >= 1) {
            return n * factorial(x - 1);
        } else {
            return 1;
        }
    };
    
    let twice = func(f, x) {
        return f(f(x));
    }
    
    let increment = func(x) {
        return x + 1;
    }
    
    twice(increment, 3); // => 5
```

# REPL

The REPL now supports parsing a program.

```
Hello ajitem! Welcome to The Donut Programming Language!
Awaiting input...
🍩 let x = 1 + 3 * 2
let x = (1 + (3 * 2));
🍩 let y = 3 + 5 * 2 - 1 / 3
let y = ((3 + (5 * 2)) - (1 / 3));
🍩 let y 12 * 2
        expected next to be =, got Integer instead
🍩 
```

# Next Steps

The next step here is to evaluate the parsed program :)

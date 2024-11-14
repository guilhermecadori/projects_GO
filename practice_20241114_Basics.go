// Basics
/*
	Packages
		Go applications are organized in packages. A package is a collection of source files 
		located in the same directory. All source files in a directory must share the same package name. 
		It is conventional for the package name to be the last directory in the import path. 
		For example, the files in the "math/rand" package begin with the statement package rand. 
		When a package is imported, only entities (functions, types, variables, constants) whose 
		names start with a capital letter can be used / accessed. The recommended style of naming 
		in Go is that identifiers will be named using camelCase, except for those meant to be 
		accessible across packages which should be PascalCase.

	Variables
		Go is statically-typed, which means all variables must have a defined type at compile-time.
		Variables can be defined by explicitly specifying a type:
			var explicit int // Explicitly typed

		You can also use an initializer, and the compiler will assign the variable type to match the type 
		of the initializer.
			implicit := 10   // Implicitly typed as an int

		Once declared, variables can be assigned values using the = operator. Once declared, a variable's type 
		can never change.
			count := 1 // Assign initial value
			count = 2  // Update to new value
			count = false // This throws a compiler error due to assigning a non `int` type
	
	Constants
		Constants hold a piece of data just like variables, but their value cannot change during the execution 
		of the program. Constants are defined using the const keyword and can be numbers, characters,
		strings or booleans:
			const Age = 21 // Defines a numeric constant 'Age' with the value of 21

	Functions
		Go functions accept zero or more parameters. Parameters must be explicitly typed, there is no type inference.
		Values are returned from functions using the return keyword.
		A function is invoked by specifying the function name and passing arguments for each of the 
		function's parameters.
*/

// Example:
package main

// Hello is a public function.
func Hello (name string) string {
    return hi(name)
}

// hi is a private function.
func hi (name string) string {
    return "hi " + name
}






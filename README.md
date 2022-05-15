# Iterator

## Introduction

This project is an implementation of iterators for go using generics. I started this project because I like the 
functional map/filter/reduce pattern that is offered by many languages. 

Go is a simple language - meaning that it is lean in features - but it offers functional principles in the form of 
closures.

Go did not support generics until version 1.18. Go 1.18.x is at the time of writing (2022-05-06) the latest release version.  

I was wondering if Go would now allow us to write a good type safe iterator framework. 

I also wanted to evaluate the Gherkin/Cucumber BDD test framework [Godog](https://github.com/cucumber/godog).

## Usage

Please take a look at the examples in [iterators_test.go](iterators_test.go).

## Conclusions

### Generics

Go does not support Generics well enough to allow me to design an intuitive framework.

   * Go does not allow for Generic methods. These are required for chaining the operations like: 
     ```go 
      FromSlice(sliceOfStructs).Map(toString).Filter(oneCharacterStrings).Reduce(mostCommonCharacter)
     ``` 
   * Go is not able to infer types in cases where a generic struct is passed to a generic interface. This is
     unfortunate because in Go it is expected for functions to accept interfaces and return structs. 
     ```go 
     // This will not compile.
     iter := FromSlice([]int{1,2,3,4})
     Map(iter, toString)
     // This will compile, because we specify the generic type of int when calling the generic Map function.
     iter := FromSlice([]int{1,2,3,4})
     Map[int](iter, toString)
     ```
     
There are workarounds possible for the second issue, such as:
   * Returning the generic interface from all functions and methods.
   * Create a method in each type that returns a generic interface from the struct. 

Returning a generic interface means a struct can only implement that interface closes a lot of doors for future 
expansion. 

Creating a method in each type that returns a generic interface can make the library harder to use, and is not a lot 
better than just specifying the generic type of the receiving function.

Such a workaround could look like this: 

```go
     iter := FromSlice([]int{1,2,3,4})
     Map(iter.I(), toString)
``` 

Instead of: 

```go
    iter := FromSlice([]int{1,2,3,4})
    Map[int](iter, toString)
```

Also, because there is no support for generic methods means no generic fluent API like constructs are possible.

So, I think the current design is the best design for this library yet and future versions of the Go compiler could 
have improvements on generics which makes this library automatically friendlier to use. 

### Godog

[Godog](https://github.com/cucumber/godog) is awesome. 

* It is possible to run Godog tests with `go test` by using the subtest functionality of Go. So it 
  will integrate in pipelines already configured for pure go tests.
* It allows you to write tests in readable language. ([gherkin](https://cucumber.io/docs/gherkin/reference/) DSL)
* It allows you to reuse test code which makes extending the tests easier the more tests you have written.
* It generates skeleton code for the features that you have written.

Godog does have some small issues
 
* Sometimes it does not generate the correct regular expression in the step.   
  In some cases these need small manual modifications. Not a big issues, because it is easy to modify by hand. 
* The captures in the step are passed as arguments to a function. Godog does not support all build-in types (yet?).
* There is no 1.0.0 release yet.

It is a very usable library. The small issues I encountered where easy to workaround.


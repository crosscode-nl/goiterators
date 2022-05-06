# Iterators

## Introduction

This project is en implementation of iterators for go using generics. I started this project because I like the 
functional map/filter/reduce pattern that is offered by many languages. 

Go is a simple language - meaning that it is lean in features - but it offers functional principles in the form of 
closures.

Go did not support generics until version 1.18. Go 1.18.x is at the time of writing (2022-05-06) the latest release version.  

I was wondering if Go would now allow us to write a good type safe iterator framework. 

Next thing I wanted to test is the Gherkin/Cucumber BDD test framework [Godog](https://github.com/cucumber/godog).  

## Conclusion

### Generics

Go does not support Generics well enough to allow for an intuitive framework.

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
     
There are workarounds possible, such as retuning the generic interface from all functions. However, no support 
for generic methods means no generic fluent API like constructs are possible. 

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


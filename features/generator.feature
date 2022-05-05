Feature: Generators return a value generating iterator
  Scenario: Generator returns a GeneratingIterator that calls the provided GeneratorFunc repeat count times.
    Given a GeneratorFunc that returns the count and repeat concatenated with a comma.
    And a repeat value of 3
    When Generate() is called
    Then calling Next() until false is returned should return the following strings:
    | 0,3 |
    | 1,3 |
    | 2,3 |

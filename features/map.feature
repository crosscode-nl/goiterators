Feature: Map modifies items item the iteration to a new value and/or type
  A valid Iterable and functioning iterator is returned when Filter is called

  Scenario: A slice with 3 items returns an Iterable that returns exactly 3 items when FromSlice is called
    Given an Iterable with the following values:
      | 1 |
      | 2 |
      | 3 |
    And a predicate that only selects odd numbers
    When Filter is called
    Then Next() returns true 2 times and then returns false

  Scenario: A slice with 3 items returns an Iterable that returns the provided items when FromSlice is called
    Given an Iterable with the following values:
      | 1 |
      | 2 |
      | 3 |
    And a predicate that only selects odd numbers
    When Filter is called
    Then Get() after Next() should return:
      | 1 |
      | 3 |

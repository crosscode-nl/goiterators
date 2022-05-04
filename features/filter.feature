Feature: Filter removes items from the iteration that do not match the predicate
  A valid Iterable and functioning iterator is returned when Filter is called

  Scenario: An Iterable with 3 items returns a FilterIterator when Filter is called
    Given an Iterable with the following values:
      | 1 |
      | 2 |
      | 3 |
    And a predicate that only selects odd numbers
    When Filter is called
    Then a FilterIterator is returned

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

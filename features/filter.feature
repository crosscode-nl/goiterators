Feature: Filter removes items from the iteration that do not match the predicate
  A valid Iterable and functioning iterator is returned when Filter is called

  Scenario: An Iterable with int 1,2, & 3 items returns exactly 2 numbers
    with a predicate that selects only odd numbers
    Given an Iterable with the following values:
      | 1 |
      | 2 |
      | 3 |
    And a predicate that only selects odd numbers
    When Filter is called
    Then Next() returns true 2 times and then returns false

  Scenario: An Iterable with int 1,2, & 3 items returns 1 and 3 when filtered
    with a predicate that selects only odd numbers

    Given an Iterable with the following values:
      | 1 |
      | 2 |
      | 3 |
    And a predicate that only selects odd numbers
    When Filter is called
    Then calling Next() until false is returned should return the following integers:
      | 1 |
      | 3 |

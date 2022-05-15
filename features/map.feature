Feature: Map modifies items item the iteration to a new value and/or type
  A valid Iterable and functioning iterator is returned when Filter is called

  Scenario: An Iterable with int 1,2, & 3 items returns exactly 3 values when Map is called
    with a map function that multiples the values and converts the int to a string, prefixed with test
    Given an Iterable with the following values:
      | 1 |
      | 2 |
      | 3 |
    And a map function that multiples the values and converts the int to a string, prefixed with test
    When Map is called
    Then Next() returns true 3 times and then returns false

  Scenario: An Iterable with int 1,2, & 3 items returns test1, test2, test3 when Map is called
  with a map function that multiples the values and converts the int to a string, prefixed with test
    Given an Iterable with the following values:
      | 1 |
      | 2 |
      | 3 |
    And a map function that multiples the values and converts the int to a string, prefixed with test
    When Map is called
    Then calling Next() until false is returned should return the following strings:
      | test2 |
      | test4 |
      | test6 |

  Scenario: MapIterator handles errors in source iterator
    Given an Iterable in an error state
    And a map function that multiples the values and converts the int to a string, prefixed with test
    When Map is called
    Then Error() of string iterator returns an error

    Given an Iterable with the following values:
      | 1 |
      | 2 |
      | 3 |
    And a map function that multiples the values and converts the int to a string, prefixed with test
    When Map is called
    Then Error() of string iterator returns nil


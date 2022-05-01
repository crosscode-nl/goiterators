Feature: FromSlice returns an iterator
  A valid Iterator and functioning iterator is returned when FromSlice is called

  Scenario: A slice with 3 items returns a SliceIterator when FromSlice is called
    Given a slice with the following values by list
      | 1 |
      | 2 |
      | 3 |
    When FromSlice is called
    Then a SliceIterator is returned

  Scenario: A slice with 3 items returns an Iterable that returns exactly 3 items when FromSlice is called
    Given a slice with the following values by list
      | 1 |
      | 2 |
      | 3 |
    When FromSlice is called
    Then Next() returns true 3 times and then returns false

  Scenario: A slice with 3 items returns an Iterable that returns the provided items when FromSlice is called
    Given a slice with the following values by list
      | 1 |
      | 2 |
      | 3 |
    When FromSlice is called
    Then Get() after Next() should return:
      | 1 |
      | 2 |
      | 3 |





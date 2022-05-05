Feature: FromSlice returns an iterator
  A valid Iterator and functioning iterator is returned when FromSlice is called

  Scenario: A SliceIterator with correct values in fields is returned when FromSlice is called
    Given a slice with the following values:
      | 1 |
      | 2 |
      | 3 |
    When FromSlice is called
    Then a SliceIterator is returned with .values containing:
      | 1 |
      | 2 |
      | 3 |
    Then a SliceIterator is returned with .idx containing -1
    Then a SliceIterator is returned with .error containing nil
    Then a SliceIterator is returned with .reverse containing false

  Scenario: A slice with 3 items returns an Iterable that returns exactly 3 items when FromSlice is called
    Given a slice with the following values:
      | 1 |
      | 2 |
      | 3 |
    When FromSlice is called
    Then Next() returns true 3 times and then returns false

  Scenario: A slice with 3 items returns an Iterable that returns the provided items when FromSlice is called
    Given a slice with the following values:
      | 1 |
      | 2 |
      | 3 |
    When FromSlice is called
    Then calling Next() until false is returned should return the following integers:
      | 1 |
      | 2 |
      | 3 |





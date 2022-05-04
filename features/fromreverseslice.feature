Feature: FromReverseSlice returns an reverse iterator
  A valid Iterable and functioning iterator is returned when FromSlice is called

  Scenario: A SliceIterator with correct values in fields is returned when FromReverseSlice is called
    Given a slice with the following values:
      | 1 |
      | 2 |
      | 3 |
    When FromReverseSlice is called
    Then a SliceIterator is returned with .values containing:
      | 1 |
      | 2 |
      | 3 |
    Then a SliceIterator is returned with .idx containing -1
    Then a SliceIterator is returned with .error containing nil
    Then a SliceIterator is returned with .reverse containing true

  Scenario: A slice with 3 items returns a SliceIterator when FromReverseSlice is called
    Given a slice with the following values:
      | 1 |
      | 2 |
      | 3 |
    When FromReverseSlice is called
    Then a SliceIterator is returned

  Scenario: A slice with 3 items returns an Iterable that returns exactly 3 items when FromReverseSlice is called
    Given a slice with the following values:
      | 1 |
      | 2 |
      | 3 |
    When FromReverseSlice is called
    Then Next() returns true 3 times and then returns false

  Scenario: A slice with 3 items returns an Iterable that returns the provided items in reverse when FromReverseSlice is called
    Given a slice with the following values:
      | 1 |
      | 2 |
      | 3 |
    When FromReverseSlice is called
    Then Get() after Next() should return:
      | 3 |
      | 2 |
      | 1 |





Feature: ToSlice renders an Iterable to a slice

  Scenario: An Iterable with int 1,2, & 3 items returns a slice with 1, 2 and 3
    Given an Iterable with the following values:
      | 1 |
      | 2 |
      | 3 |
    When ToSlice is called
    Then a slice is returned with the following values:
      | 1 |
      | 2 |
      | 3 |
